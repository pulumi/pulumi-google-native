// Copyright 2016-2020, Pulumi Corporation.
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
	"github.com/gedex/inflector"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-cloud/provider/pkg/discovery"
	"github.com/pulumi/pulumi-google-cloud/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/pulumi/pulumi/pkg/v2/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-google-cloud/sdk/go/google-cloud"

// PulumiSchema will generate a Pulumi schema for the given Google Cloud discovery documents.
func PulumiSchema() (*schema.PackageSpec, *resources.CloudAPIMetadata, error) {
	pkg := schema.PackageSpec{
		Name:        "google-cloud",
		Description: "A Next Generation Pulumi package for creating and managing Google Cloud resources.",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "google cloud"},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-google-cloud",
		Config:      schema.ConfigSpec{},
		Provider: schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The provider type for the Google Cloud package.",
				Type:        "object",
			},
			InputProperties: map[string]schema.PropertySpec{},
		},
		Types:     map[string]schema.ComplexTypeSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Language:  map[string]json.RawMessage{},
	}
	metadata := resources.CloudAPIMetadata{
		Resources: map[string]resources.CloudAPIResource{},
	}

	csharpNamespaces := map[string]string{
		"google-cloud": "GoogleCloud",
	}
	pythonModuleNames := map[string]string{}
	golangImportAliases := map[string]string{}

	var fileNames []string
	root := path.Join(".", "discovery")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	for _, fileName := range fileNames {
		document, err := readDiscoveryDocument(fileName)
		if err != nil {
			return nil, nil, err
		}

		module := fmt.Sprintf("%s/%s", document.Name, document.Version)
		gen := packageGenerator{
			pkg:          &pkg,
			metadata:     &metadata,
			rest:         document,
			mod:          module,
			visitedTypes: codegen.NewStringSet(),
		}
		csharpNamespaces[module] = fmt.Sprintf("%s.%s", strings.Title(document.Name), strings.Title(document.Version))
		pythonModuleNames[module] = module
		golangImportAliases[filepath.Join(goBasePath, module)] = document.Name

		err = gen.findResources("", document.Resources)
		if err != nil {
			return nil, nil, err
		}
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":       goBasePath,
		"packageImportAliases": golangImportAliases,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^2.0.0",
		},
		"readme": `TODO`,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=2.11.2,<3.0.0",
		},
		"usesIOClasses": true,
		"readme":        `TODO`,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "2.*",
			"System.Collections.Immutable": "1.6.0",
		},
		"namespaces": csharpNamespaces,
	})

	return &pkg, &metadata, nil
}

func readDiscoveryDocument(fileName string) (*discovery.RestDescription, error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var rest discovery.RestDescription
	err = json.Unmarshal(byteValue, &rest)
	if err != nil {
		return nil, err
	}
	return &rest, nil
}

type packageGenerator struct {
	pkg          *schema.PackageSpec
	metadata     *resources.CloudAPIMetadata
	rest         *discovery.RestDescription
	mod          string
	visitedTypes codegen.StringSet
}

func (g *packageGenerator) findResources(parent string, resources map[string]discovery.RestResource) error {
	for name, res := range resources {
		name = ToUpperCamel(inflector.Singularize(name))
		var createMethod, deleteMethod *discovery.RestMethod
		for methodName, value := range res.Methods {
			restMethod := value
			switch methodName {
			case "create", "insert":
				createMethod = &restMethod
			case "setIamPolicy":
				typeName := fmt.Sprintf("%s%sIamPolicy", parent, name)
				err := g.genResource(typeName, &restMethod, nil)
				if err != nil {
					return err
				}
			case "delete":
				deleteMethod = &restMethod
			}
		}

		if createMethod != nil {
			typeName := fmt.Sprintf("%s%s", parent, name)
			err := g.genResource(typeName, createMethod, deleteMethod)
			if err != nil {
				return err
			}
		}

		var newParent string
		switch name {
		case "Project", "Location", "Zone", "Global":
			newParent = parent
		default:
			newParent = parent + name
		}
		err := g.findResources(newParent, res.Resources)
		if err != nil {
			return err
		}
	}
	return nil
}

var pathRegex = regexp.MustCompile(`{([A-Za-z0-9]*)}`)

func (g *packageGenerator) genResource(typeName string, createMethod, deleteMethod *discovery.RestMethod) error {
	resourceTok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, typeName)

	inputProperties := map[string]schema.PropertySpec{}
	requiredProperties := codegen.NewStringSet()

	createPath := createMethod.FlatPath
	if createPath == "" {
		createPath = createMethod.Path
	}

	for name, param := range createMethod.Parameters {
		if param.Location != "query" || !param.Required {
			continue
		}
		createPath += fmt.Sprintf("?%s={%[1]s}", name)
	}

	resourceMeta := resources.CloudAPIResource{
		BaseUrl:    g.rest.BaseUrl,
		CreatePath: createPath,
	}

	subMatches := pathRegex.FindAllStringSubmatch(createPath, -1)
	for _, names := range subMatches {
		name := names[1]
		inputProperties[name] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{Type: "string"},
		}
		resourceMeta.CreateParams = append(resourceMeta.CreateParams, name)
		requiredProperties.Add(name)
	}

	if deleteMethod != nil {
		deletePath := deleteMethod.FlatPath
		if deletePath == "" {
			deletePath = deleteMethod.Path
		}

		resourceMeta.DeletePath = deletePath

		subMatches := pathRegex.FindAllStringSubmatch(deletePath, -1)
		for _, names := range subMatches {
			name := names[1]
			if _, has := inputProperties[name]; !has {
				inputProperties[name] = schema.PropertySpec{
					TypeSpec: schema.TypeSpec{Type: "string"},
				}
				requiredProperties.Add(name)
			}
			resourceMeta.DeleteParams = append(resourceMeta.DeleteParams, name)
		}
	}

	if createMethod.Request != nil {
		request := g.rest.Schemas[createMethod.Request.Ref]
		bodyProperties, err := g.genProperties(&request)
		if err != nil {
			return err
		}

		for name, prop := range bodyProperties {
			inputProperties[name] = prop
		}
	}

	if resourceTok == "google-cloud:storage/v1:Object" {
		resourceTok = "google-cloud:storage/v1:BucketObject"
		inputProperties["source"] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{
				Ref: "pulumi.json#/Asset",
			},
		}
	}

	resourceSpec := schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: createMethod.Description,
			Type:        "object",
			// TODO: Output properties
		},
		InputProperties: inputProperties,
		RequiredInputs:  requiredProperties.SortedValues(),
	}
	g.pkg.Resources[resourceTok] = resourceSpec
	g.metadata.Resources[resourceTok] = resourceMeta
	return nil
}

func (g *packageGenerator) genProperties(typeSchema *discovery.JsonSchema) (map[string]schema.PropertySpec, error) {
	result := map[string]schema.PropertySpec{}
	for name, value := range typeSchema.Properties {
		prop := value
		typeSpec, err := g.genTypeSpec(&prop)
		if err != nil {
			return nil, err
		}

		result[name] = schema.PropertySpec{
			Description: prop.Description,
			TypeSpec:    *typeSpec,
		}
	}
	return result, nil
}

func makeValidIdentifier(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}

func (g *packageGenerator) genTypeSpec(prop *discovery.JsonSchema) (*schema.TypeSpec, error) {
	switch {
	case prop.Items != nil:
		items, err := g.genTypeSpec(prop.Items)
		if err != nil {
			return nil, err
		}

		return &schema.TypeSpec{
			Type:  "array",
			Items: items,
		}, nil
	case prop.Type == "any":
		return &schema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
	case prop.Type != "":
		return &schema.TypeSpec{Type: prop.Type}, nil
	case prop.Ref != "":
		tok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, prop.Ref)
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)

		if !g.visitedTypes.Has(tok) {
			g.visitedTypes.Add(tok)

			typeSchema := g.rest.Schemas[prop.Ref]
			properties, err := g.genProperties(&typeSchema)
			if err != nil {
				return nil, err
			}

			g.pkg.Types[tok] = schema.ComplexTypeSpec{
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: typeSchema.Description,
					Type:        "object",
					Properties:  properties,
				},
			}
		}

		return &schema.TypeSpec{
			Type: "object",
			Ref:  referencedTypeName,
		}, nil
	}
	return nil, errors.New("unknown type")
}

func resourceName(restMethod discovery.RestMethod) (name string) {
	// TODO: This is very ad-hoc, find a universal naming approach.
	name = restMethod.Response.Ref
	switch name {
	case "Operation":
		name = restMethod.Request.Ref
		name = strings.TrimPrefix(name, "Create")
		name = strings.TrimSuffix(name, "Request")
	case "Object":
		name = "BucketObject"
	}
	return
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
