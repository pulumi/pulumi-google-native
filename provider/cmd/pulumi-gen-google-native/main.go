// Copyright 2016-2021, Pulumi Corporation.

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
		err := copyJsonFile(item.DiscoveryRestUrl, outDir, fileName)
		if err != nil {
			return errors.Wrapf(err, "writing %s", fileName)
		}
	}

	return nil
}

func copyJsonFile(url, outDir, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var ifce interface{}
	err = json.Unmarshal(body, &ifce)
	if err != nil {
		return err
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
		return nodejsgen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "python":
		return pythongen.GeneratePackage(toolDescription, ppkg, extraFiles)
	case "go":
		return gogen.GeneratePackage(toolDescription, ppkg)
	case "dotnet":
		return dotnetgen.GeneratePackage(toolDescription, ppkg, extraFiles)
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
