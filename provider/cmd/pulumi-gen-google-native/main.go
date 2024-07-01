// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"google.golang.org/api/discovery/v1"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-google-native/provider/pkg/gen"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	dotnetgen "github.com/pulumi/pulumi/pkg/v3/codegen/dotnet"
	gogen "github.com/pulumi/pulumi/pkg/v3/codegen/go"
	nodejsgen "github.com/pulumi/pulumi/pkg/v3/codegen/nodejs"
	pythongen "github.com/pulumi/pulumi/pkg/v3/codegen/python"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tools"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {
	languages := os.Args[1]

	version := ""
	if len(os.Args) == 3 {
		version = os.Args[2]
	}

	pkgSpec, meta, err := gen.PulumiSchema()
	if err != nil {
		panic(err)
	}

	for _, language := range strings.Split(languages, ",") {
		switch language {
		case "discovery":
			dir := path.Join(".", "discovery")
			if err = emitDiscoveryFiles(dir); err != nil {
				break
			}
		case "schema":
			dir := path.Join(".", "provider", "cmd", "pulumi-resource-google-native")
			if err = emitSchema(*pkgSpec, version, dir, "main", true); err != nil {
				break
			}
			// Also, emit the resource metadata for the provider.
			if err = emitMetadata(meta, dir, "main", true); err != nil {
				break
			}
		default:
			dir := path.Join(".", "sdk", language)
			pkgSpec.Version = version
			err = emitPackage(pkgSpec, language, dir)
		}
		if err != nil {
			panic(fmt.Sprintf("%+v", err))
		}
	}
}

func emitDiscoveryFiles(outDir string) error {
	resp, err := http.Get("https://discovery.googleapis.com/discovery/v1/apis?parameters")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	byteValue, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var list discovery.DirectoryList
	err = json.Unmarshal(byteValue, &list)
	if err != nil {
		return err
	}

	// TODO: this one item is not found in the directory list - figure out why.
	// https://github.com/pulumi/pulumi-google-native/issues/69
	list.Items = append(list.Items, &discovery.DirectoryListItems{
		DiscoveryRestUrl: "https://vpcaccess.googleapis.com/$discovery/rest?version=v1",
		Title:            "Cloud VPC Access",
		Id:               "vpcaccess:v1",
	})

	private := regexp.MustCompile("v[0-9]+p[0-9]+[a-z0-9]+")

	for _, item := range list.Items {
		// TODO: this is arbitrary - find a better way?
		// https://github.com/pulumi/pulumi-google-native/issues/69
		if !strings.HasPrefix(item.DocumentationLink, "https://cloud.google.com") &&
			!strings.HasPrefix(item.Title, "Cloud") &&
			!strings.HasPrefix(item.Title, "Firebase") &&
			!strings.HasPrefix(item.Title, "Identity Toolkit") &&
			!strings.HasPrefix(item.Title, "Compute") {
			continue
		}

		if private.MatchString(item.Version) {
			// Skip private previews
			continue
		}

		fileName := fmt.Sprintf("%s.json", strings.ReplaceAll(item.Id, ":", "_"))
		err := copyJSONFile(item.DiscoveryRestUrl, outDir, fileName)
		if err != nil {
			return errors.Wrapf(err, "writing %s", fileName)
		}
	}

	return nil
}

func copyJSONFile(url, outDir, path string) error {
	// nolint: gosec
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ifce := map[string]interface{}{}
	err = json.Unmarshal(body, &ifce)
	if err != nil {
		return err
	}

	if rev, ok := ifce["revision"].(string); ok {
		f, err := os.Open(filepath.Join(outDir, path))
		if err == nil {
			existing := map[string]interface{}{}
			if err := json.NewDecoder(f).Decode(&existing); err != nil {
				return err
			}
			existingRev, ok := existing["revision"].(string)
			if ok && existingRev > rev {
				// We sometimes get older revisions from discovery calls. Ignore if this is the case.
				fmt.Printf("Existing revision for %s (%s) is newer than retrieved (%s), skipping.\n", path,
					existingRev, rev)
				return nil
			}
		} else if !os.IsNotExist(err) {
			return err
		}
	}
	// Marshaller will write JSON nodes in alphanumeric order which makes the files "stable",
	// and we can track the history more easily.
	output, err := json.MarshalIndent(ifce, "", "  ")
	if err != nil {
		return err
	}

	return emitFile(outDir, path, output)
}

// emitSchema writes the Pulumi schema JSON to the 'schema.json' file in the given directory.
func emitSchema(pkgSpec schema.PackageSpec, version, outDir string, goPackageName string, emitJSON bool) error {
	schemaJSON, err := json.MarshalIndent(pkgSpec, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling Pulumi schema")
	}

	// Ensure the spec is stamped with a version.
	pkgSpec.Version = version

	compressedSchema, err := gen.CompressSchema(pkgSpec)
	if err != nil {
		return errors.Wrap(err, "failed to compress schema")
	}
	err = emitFile(outDir, "schema.go", []byte(fmt.Sprintf(`package %s
var pulumiSchema = %#v
`, goPackageName, compressedSchema)))
	if err != nil {
		return errors.Wrap(err, "saving metadata")
	}

	if emitJSON {
		if err := emitFile(outDir, "schema.json", schemaJSON); err != nil {
			return err
		}
	}
	return nil
}

func emitMetadata(metadata *resources.CloudAPIMetadata, outDir string, goPackageName string, emitJSON bool) error {
	compressedMeta := bytes.Buffer{}
	compressedWriter := gzip.NewWriter(&compressedMeta)
	err := json.NewEncoder(compressedWriter).Encode(metadata)
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	if err = compressedWriter.Close(); err != nil {
		return err
	}

	formatted, err := json.MarshalIndent(metadata, "", "    ")
	if err != nil {
		return errors.Wrap(err, "marshaling metadata")
	}

	err = emitFile(outDir, "metadata.go", []byte(fmt.Sprintf(`package %s
var cloudApiResources = %#v
`, goPackageName, compressedMeta.Bytes())))
	if err != nil {
		return err
	}

	if emitJSON {
		err := emitFile(outDir, "metadata.json", formatted)
		if err != nil {
			return err
		}
	}
	return nil
}

func generate(ppkg *schema.Package, language string) (map[string][]byte, error) {
	toolDescription := "the Pulumi SDK Generator"
	extraFiles := map[string][]byte{}
	switch language {
	case "nodejs":
		return nodejsgen.GeneratePackage(toolDescription, ppkg, extraFiles, map[string]string{})
	case "python":
		return pythongen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "go":
		return gogen.GeneratePackage(toolDescription, ppkg)
	case "dotnet":
		return dotnetgen.GeneratePackage(toolDescription, ppkg, extraFiles, map[string]string{})
	}

	return nil, errors.Errorf("unknown language '%s'", language)
}

// emitPackage emits an entire package pack into the configured output directory with the configured settings.
func emitPackage(pkgSpec *schema.PackageSpec, language, outDir string) error {
	ppkg, err := schema.ImportSpec(*pkgSpec, nil)
	if err != nil {
		return errors.Wrap(err, "reading schema")
	}

	files, err := generate(ppkg, language)
	if err != nil {
		return errors.Wrapf(err, "generating %s package", language)
	}

	for f, contents := range files {
		if err := emitFile(outDir, f, contents); err != nil {
			return errors.Wrapf(err, "emitting file %v", f)
		}
	}

	return nil
}

// emitFile creates a file in a given directory and writes the byte contents to it.
func emitFile(outDir, relPath string, contents []byte) error {
	p := path.Join(outDir, relPath)
	if err := tools.EnsureDir(path.Dir(p)); err != nil {
		return errors.Wrap(err, "creating directory")
	}

	f, err := os.Create(p)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer contract.IgnoreClose(f)

	_, err = f.Write(contents)
	return err
}
