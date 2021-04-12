// Copyright 2016-2021, Pulumi Corporation.

package gen

import (
	"encoding/json"
	"fmt"
	"github.com/gedex/inflector"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-gcp-native/provider/pkg/discovery"
	"github.com/pulumi/pulumi-gcp-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-gcp-native/sdk/go/gcp"

// PulumiSchema will generate a Pulumi schema for the given Google Cloud discovery documents.
func PulumiSchema() (*schema.PackageSpec, *resources.CloudAPIMetadata, error) {
	pkg := schema.PackageSpec{
		Name:        "gcp-native",
		Description: "A native Pulumi package for creating and managing Google Cloud resources.",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "google cloud"},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-gcp-native",
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
		"gcp-native": "GcpNative",
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

		if document.Name == "" {
			continue
		}

		module := fmt.Sprintf("%s/%s", document.Name, document.Version)
		gen := packageGenerator{
			pkg:          &pkg,
			metadata:     &metadata,
			rest:         document,
			mod:          module,
			visitedTypes: codegen.NewStringSet(),
			docName:      fileName,
		}
		csharpNamespaces[document.Name] = csharpNamespace(document)
		csharpNamespaces[module] = csharpVersionedNamespace(document)
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
			"@pulumi/pulumi": "^3.0.0-alpha.0",
		},
		"readme": `TODO`,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=3.0.0a1,<4.0.0",
		},
		"usesIOClasses": true,
		"readme":        `TODO`,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "3.*-*",
			"System.Collections.Immutable": "1.6.0",
		},
		"namespaces": csharpNamespaces,
	})

	return &pkg, &metadata, nil
}

var titleReplacer = strings.NewReplacer(" ", "", "-", "")
var versionReplacer = strings.NewReplacer("alpha", "Alpha", "beta", "Beta", "v", "V")

func csharpNamespace(document *discovery.RestDescription) string {
	moduleName := strings.Title(document.Name)
	title := titleReplacer.Replace(document.Title)
	idx := strings.Index(strings.ToLower(title), document.Name)
	if idx >= 0 {
		moduleName = title[idx : idx+len(document.Name)]
	}
	return moduleName
}
func csharpVersionedNamespace(document *discovery.RestDescription) string {
	moduleName := csharpNamespace(document)
	version := versionReplacer.Replace(document.Version)
	return fmt.Sprintf("%s.%s", moduleName, version)
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
	docName      string
}

func (g *packageGenerator) findResources(parent string, resources map[string]discovery.RestResource) error {
	var postMethods []discovery.RestMethod
	for _, resourceName := range codegen.SortedKeys(resources) {
		res := resources[resourceName]
		name := ToUpperCamel(inflector.Singularize(resourceName))
		var createMethod, getMethod, listMethod, updateMethod, deleteMethod *discovery.RestMethod
		for methodName, value := range res.Methods {
			restMethod := value
			if restMethod.HttpMethod == "POST" {
				postMethods = append(postMethods, restMethod)
			}
			switch methodName {
			case "create", "insert":
				createMethod = &restMethod
			case "submit", "register":
				if createMethod == nil {
					createMethod = &restMethod
				}
			case "update", "replaceService" /*special case for Cloud Run*/ :
				updateMethod = &restMethod
			case "patch":
				if updateMethod == nil {
					updateMethod = &restMethod
				}
			case "get":
				getMethod = &restMethod
			case "read":
				if getMethod == nil {
					getMethod = &restMethod
				}
			case "list":
				listMethod = &restMethod
			case "setIamPolicy":
				typeName := fmt.Sprintf("%s%sIamPolicy", parent, name)
				if getIamPolicy, has := res.Methods["getIamPolicy"]; has {
					err := g.genResource(typeName, &restMethod, &getIamPolicy, &restMethod, nil)
					if err != nil {
						return err
					}
				}
			case "delete":
				deleteMethod = &restMethod
			}
		}

		typeName := fmt.Sprintf("%s%s", parent, name)
		switch {
		case createMethod != nil && getMethod != nil && getMethod.HttpMethod != "GET":
			// We support read with GET method only. Warn about others.
			fmt.Printf("Get method %q is not supported: %s (%s), skipping.\n", getMethod.HttpMethod, typeName, g.docName)
		case createMethod != nil && getMethod != nil:
			err := g.genResource(typeName, createMethod, getMethod, updateMethod, deleteMethod)
			if err != nil {
				return err
			}
		case strings.Contains(typeName, "Operation"):
			// Operation variants don't need to be made resources.
			continue
		case createMethod != nil && getMethod == nil && listMethod != nil:
			// List methods return collections, so the shape of response wouldn't match the SDK map.
			// Skip them for now but emit a note.
			fmt.Printf("List methods are not supported: %s (%s), skipping.\n", typeName, g.docName)
		case (deleteMethod != nil || updateMethod != nil) && len(postMethods) > 0:
			// It can be useful to look at this output and look for potential missing resources with unexpected
			// HTTP POST method names.
			fmt.Printf("No create method for resource: %s (%s), skipping.\n", typeName, g.docName)
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

func (g *packageGenerator) genToken(typeName string) string {
	return fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, typeName)
}

func (g *packageGenerator) genResource(typeName string, createMethod, getMethod, updateMethod, deleteMethod *discovery.RestMethod) error {
	resourceTok := g.genToken(typeName)

	inputProperties := map[string]schema.PropertySpec{}
	requiredInputProperties := codegen.NewStringSet()

	createPath := createMethod.FlatPath
	if createPath == "" {
		createPath = createMethod.Path
	}

	for name, param := range createMethod.Parameters {
		if param.Location == "path" && strings.HasPrefix(param.Description, "Deprecated.") {
			// If a path parameter is deprecated, the URL is effectively deprecated, so skip this resource.
			return nil
		}
		if param.Location != "query" || !param.Required {
			continue
		}
		createPath += fmt.Sprintf("?%s={%[1]s}", name)
	}

	resourceMeta := resources.CloudAPIResource{
		BaseUrl:    g.rest.BaseUrl,
		CreatePath: createPath,
		CreateVerb: createMethod.HttpMethod,
		NoDelete:   deleteMethod == nil,
	}

	subMatches := pathRegex.FindAllStringSubmatch(createPath, -1)
	for _, names := range subMatches {
		name := names[1]
		inputProperties[name] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{Type: "string"},
		}
		resourceMeta.CreateParams = append(resourceMeta.CreateParams, name)
		requiredInputProperties.Add(name)
	}

	idPath := getMethod.FlatPath
	if idPath == "" {
		idPath = getMethod.Path
	}

	resourceMeta.IdPath = idPath
	subMatches = pathRegex.FindAllStringSubmatch(idPath, -1)
	for _, names := range subMatches {
		name := names[1]
		if _, has := inputProperties[name]; !has {
			inputProperties[name] = schema.PropertySpec{
				TypeSpec: schema.TypeSpec{Type: "string"},
			}
			requiredInputProperties.Add(name)
		}
		resourceMeta.IdParams = append(resourceMeta.IdParams, name)
	}

	if createMethod.Request != nil {
		// If the request type matches the pattern when it contains a property of the type equal to the response
		// type of the GET endpoint, then we want to flatten that property, so that the resource inputs are a superset
		// of the resource outputs. This also helps reconcile the shape of create and update operations.
		var flatten string
		createRequest := g.rest.Schemas[createMethod.Request.Ref]
		if createMethod.Request.Ref != getMethod.Response.Ref {
			for _, name := range codegen.SortedKeys(createRequest.Properties) {
				v := createRequest.Properties[name]
				if v.Ref == getMethod.Response.Ref {
					flatten = name
				}
			}
		}

		bodyBag, err := g.genProperties(&createRequest, flatten, false)
		if err != nil {
			return err
		}

		for name, prop := range bodyBag.specs {
			inputProperties[name] = prop
		}
		for name := range bodyBag.requiredSpecs {
			requiredInputProperties.Add(name)
		}
		resourceMeta.CreateProperties = bodyBag.properties

		if updateMethod != nil {
			resourceMeta.UpdateVerb = updateMethod.HttpMethod
			var updateFlatten string
			updateRequest := g.rest.Schemas[updateMethod.Request.Ref]
			if updateMethod.Request.Ref != getMethod.Response.Ref {
				for _, name := range codegen.SortedKeys(updateRequest.Properties) {
					v := updateRequest.Properties[name]
					if v.Ref == getMethod.Response.Ref {
						updateFlatten = name
					}
				}
			}
			resourceMeta.UpdateProperties = map[string]resources.CloudAPIProperty{}
			updateBag, err := g.genProperties(&updateRequest, updateFlatten, false)
			if err != nil {
				return err
			}

			for name, value := range updateBag.properties {
				if _, has := bodyBag.properties[name]; has {
					resourceMeta.UpdateProperties[name] = value
				} else {
					// TODO: do we need to handle masks?
					if !strings.HasSuffix(name, "Mask") {
						fmt.Printf("unknown update property %s: %s.%s\n", resourceTok, updateMethod.Request.Ref, name)
					}
				}
			}
		}
	}

	properties := map[string]schema.PropertySpec{}
	requiredProperties := codegen.NewStringSet()
	if getMethod.Response != nil {
		response := g.rest.Schemas[getMethod.Response.Ref]
		responseBag, err := g.genProperties(&response, "", true)
		if err != nil {
			return err
		}
		properties = responseBag.specs
		requiredProperties = responseBag.requiredSpecs
		g.escapeCSharpNames(typeName, properties)
	}

	if resourceTok == "gcp-native:storage/v1:Object" {
		resourceTok = "gcp-native:storage/v1:BucketObject"
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
			Properties:  properties,
			Required:    requiredProperties.SortedValues(),
		},
		InputProperties: inputProperties,
		RequiredInputs:  requiredInputProperties.SortedValues(),
	}
	g.pkg.Resources[resourceTok] = resourceSpec
	g.metadata.Resources[resourceTok] = resourceMeta
	return nil
}

func (g *packageGenerator) genProperties(typeSchema *discovery.JsonSchema, flatten string, isOutput bool) (*propertyBag, error) {
	result := propertyBag{
		specs:         map[string]schema.PropertySpec{},
		requiredSpecs: codegen.NewStringSet(),
		properties:    map[string]resources.CloudAPIProperty{},
	}
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		if strings.Contains(value.Description, "Deprecated.") {
			continue
		}
		if !isOutput && value.ReadOnly {
			continue
		}
		if isOutput && name == "id" {
			continue
		}

		prop := value

		if name == flatten {
			subtypeSchema := g.rest.Schemas[prop.Ref]
			sub, err := g.genProperties(&subtypeSchema, "", isOutput)
			if err != nil {
				return nil, err
			}
			for subName, subVal := range sub.specs {
				result.specs[subName] = subVal
			}
			for subName := range sub.requiredSpecs {
				result.requiredSpecs.Add(subName)
			}
			for subName := range sub.properties {
				result.properties[subName] = resources.CloudAPIProperty{
					Container: name,
				}
			}
			continue
		}

		typeSpec, err := g.genTypeSpec(&prop, isOutput)
		if err != nil {
			return nil, err
		}

		if prop.Required || isOutput {
			result.requiredSpecs.Add(name)
		}

		description := strings.TrimPrefix(prop.Description, "Output only. ")

		result.specs[name] = schema.PropertySpec{
			Description: description,
			TypeSpec:    *typeSpec,
		}
		result.properties[name] = resources.CloudAPIProperty{}
	}
	return &result, nil
}

func (g *packageGenerator) genTypeSpec(prop *discovery.JsonSchema, isOutput bool) (*schema.TypeSpec, error) {
	switch {
	case prop.Items != nil:
		items, err := g.genTypeSpec(prop.Items, isOutput)
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
		if isOutput {
			tok += "Response"
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)

		if !g.visitedTypes.Has(tok) {
			g.visitedTypes.Add(tok)

			typeSchema := g.rest.Schemas[prop.Ref]
			bag, err := g.genProperties(&typeSchema, "", isOutput)
			if err != nil {
				return nil, err
			}

			g.pkg.Types[tok] = schema.ComplexTypeSpec{
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: typeSchema.Description,
					Type:        "object",
					Properties:  bag.specs,
					Required:    bag.requiredSpecs.SortedValues(),
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

func (m *packageGenerator) escapeCSharpNames(typeName string, resourceResponse map[string]schema.PropertySpec) {
	for name, swagger := range resourceResponse {
		// C# doesn't allow properties to have the same name as its containing type.
		if strings.Title(name) == typeName {
			swagger.Language = map[string]json.RawMessage{
				"csharp": rawMessage(map[string]interface{}{
					"name": fmt.Sprintf("%sValue", typeName),
				}),
			}
			resourceResponse[name] = swagger
		}
	}
}

// propertyBag keeps the schema and metadata properties for a single API type.
type propertyBag struct {
	specs         map[string]schema.PropertySpec
	properties    map[string]resources.CloudAPIProperty
	requiredSpecs codegen.StringSet
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
