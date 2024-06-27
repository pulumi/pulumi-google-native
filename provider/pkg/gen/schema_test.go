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

package gen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jtacoma/uritemplates"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v3/codegen"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"

	"github.com/stretchr/testify/require"
	"google.golang.org/api/discovery/v1"
)

var root = filepath.Join("..", "..", "..")

// This is a validation test to make sure that the metadata generated at schema generation stands-up
// to the assumptions made for operation URL resolution at runtime in the provider.
// NOTE this will need to be kept up-to-date with the provider.go implementation but this framework gives
// valuable confidence in correctness around metadata generation without destabilizing the provider at runtime.
func TestMetadata_Operations(t *testing.T) {
	cwd, err := os.Getwd()
	require.NoError(t, err)
	t.Logf("Current directory: %s", cwd)
	var fileNames []string
	require.NoError(t, filepath.Walk(filepath.Join(root, "discovery"), func(path string, info os.FileInfo,
		err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	}))

	metadata, err := loadMetadata()
	require.NoError(t, err)

	pkgSpec, err := loadSchema()
	require.NoError(t, err)

	for _, fileName := range fileNames {
		document, err := readDiscoveryDocument(fileName)
		require.NoError(t, err)

		if document.Name == "" {
			continue
		}
		module := fmt.Sprintf("%s/%s", document.Name, document.Version)
		res, ops, err := findResources(fileName, document.Resources, document.Schemas)
		require.NoError(t, err)

		resourcesForModule := codegen.NewStringSet()
		for tok := range metadata.Resources {
			if strings.Contains(tok, module) {
				resourcesForModule.Add(tok)
			}
		}

		for _, tok := range resourcesForModule.SortedValues() {
			mdResource := metadata.Resources[tok]
			resource := pkgSpec.Resources[tok]
			validateOpTemplate := func(resourceOperationType string, op *resources.Operations,
				ddResource *discoveryDocumentResource) {
				if op == nil || op.Template == "" {
					return
				}
				tmpl, err := uritemplates.Parse(op.Template)
				require.NoError(t, err)
				paramNames := codegen.NewStringSet(tmpl.Names()...)

				// For each operation type/verb, find the corresponding operation body as referenced in
				// the discovery docs.
				var referencedOp *operation
				switch resourceOperationType {
				case "CREATE":
					if ddResource.createMethod != nil && ddResource.createMethod.Response.Ref != "" {
						if createOp, ok := ops[ddResource.createMethod.Response.Ref]; ok {
							referencedOp = createOp
						}
					}
				case "UPDATE":
					if ddResource.updateMethod != nil && ddResource.updateMethod.Response.Ref != "" {
						if updateOp, ok := ops[ddResource.updateMethod.Response.Ref]; ok {
							referencedOp = updateOp
						}
					}
				case "DELETE":
					if ddResource.deleteMethod != nil && ddResource.deleteMethod.Response.Ref != "" {
						if deleteOp, ok := ops[ddResource.deleteMethod.Response.Ref]; ok {
							referencedOp = deleteOp
						}
					}
				default:
					assert.Failf(t, "unexpected resourceOperationType: %q", resourceOperationType)
				}

				for _, v := range op.Values {
					if v.Kind == "query" {
						continue
					}
					assert.Containsf(t, paramNames, v.Name, "[%s] [%s] %s.%s not found", resourceOperationType,
						op.Template, tok, v.Name)
					paramNames.Delete(v.Name)
					expectedName := v.Name
					if v.SdkName != "" {
						expectedName = v.SdkName
					}

					found := false
					// First check if the param is in the operation response body.
					if referencedOp != nil {
						for name := range referencedOp.schema.Properties {
							if name == expectedName {
								found = true
								break
							}
						}

						if found {
							continue
						}
					}

					// Next look in the resource input properties.
					for name := range resource.InputProperties {
						if name == expectedName {
							found = true
							break
						}
					}

					if found {
						continue
					}

					// Finally check in the resource properties.
					for name := range resource.Properties {
						if name == expectedName {
							found = true
							break
						}
					}

					assert.Truef(t, found, "[%s] [%s] %s (sdkname: %s) not found in %s", resourceOperationType, op.Template,
						v.Name,
						expectedName,
						tok)
				}

				assert.Emptyf(t, paramNames, "[%s] [%s] superfluous parameters in template: %+v for %q",
					resourceOperationType, op.Template, paramNames, tok)
			}
			resourceName := strings.Split(tok, ":")[2]
			if ddResource, ok := res[resourceName]; ok {
				validateOpTemplate("CREATE", mdResource.Create.Operations, &ddResource)
				validateOpTemplate("UPDATE", mdResource.Update.Operations, &ddResource)
				validateOpTemplate("DELETE", mdResource.Delete.Operations, &ddResource)
			}
		}

	}
}

// loadMetadata deserializes the provided compressed json byte array into a CloudAPIMetadata struct.
func loadMetadata() (*resources.CloudAPIMetadata, error) {
	var resourceMap resources.CloudAPIMetadata

	bytes, err := os.Open(filepath.Join(root, "provider", "cmd", "pulumi-resource-google-native",
		"metadata.json"))
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(bytes).Decode(&resourceMap); err != nil {
		return nil, errors.Wrap(err, "unmarshalling resource map")
	}
	return &resourceMap, nil
}

func loadSchema() (*schema.PackageSpec, error) {
	var pkg schema.PackageSpec

	bytes, err := os.Open(filepath.Join(root, "provider", "cmd", "pulumi-resource-google-native",
		"schema.json"))
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(bytes).Decode(&pkg); err != nil {
		return nil, errors.Wrap(err, "unmarshalling pkgspec")
	}
	return &pkg, nil
}

type TypeSpecTestCase struct {
	name     string
	schema   *discovery.JsonSchema
	expected *schema.TypeSpec
}

func Test_genTypeSpec(t *testing.T) {
	cases := []TypeSpecTestCase{
		{
			name: "map[string, string]",
			schema: &discovery.JsonSchema{
				Type: "object",
				AdditionalProperties: &discovery.JsonSchema{
					Type: "string",
				},
			},
			expected: &schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: &schema.TypeSpec{Type: "string"},
			},
		},
	}
	g := packageGenerator{
		pkg:          nil,
		metadata:     nil,
		rest:         nil,
		mod:          "foo",
		visitedTypes: codegen.NewStringSet(),
		docName:      "foo",
	}

	for _, tt := range cases {
		out, err := g.genTypeSpec("foo", "bar", tt.schema, true)
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, out)
	}
}
