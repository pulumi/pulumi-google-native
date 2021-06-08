// Copyright 2016-2021, Pulumi Corporation.

package gen

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"google.golang.org/api/discovery/v1"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-google-native/sdk/go/google"

// PulumiSchema will generate a Pulumi schema for the given Google Cloud discovery documents.
func PulumiSchema() (*schema.PackageSpec, *resources.CloudAPIMetadata, error) {
	pkg := schema.PackageSpec{
		Name:        "google-native",
		Description: "A native Pulumi package for creating and managing Google Cloud resources.",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "google cloud"},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-google-native",
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
		"google-native": "GoogleNative",
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

		res, err := findResources(fileName, document.Resources)
		if err != nil {
			return nil, nil, err
		}

		for _, typeName := range codegen.SortedKeys(res) {
			err := gen.genResource(typeName, res[typeName])
			if err != nil {
				return nil, nil, err
			}
		}
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":       goBasePath,
		"packageImportAliases": golangImportAliases,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^3.0.0",
		},
		"readme": `The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud
programs. This provider uses the Google Cloud REST API directly and therefore provides full access to Google Cloud.
The provider is currently in public preview and is not recommended for production deployments yet. Breaking changes
will be introduced in minor version releases.`,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
		"usesIOClasses": true,
		"readme":        `The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud
programs. This provider uses the Google Cloud REST API directly and therefore provides full access to Google Cloud.
The provider is currently in public preview and is not recommended for production deployments yet. Breaking changes
will be introduced in minor version releases.`,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi":                       "3.*",
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

var pathRegex = regexp.MustCompile(`{([A-Za-z0-9]*)}`)

func (g *packageGenerator) genToken(typeName string) string {
	return fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, typeName)
}

func (g *packageGenerator) genResource(typeName string, dd discoveryDocumentResource) error {
	resourceTok := g.genToken(typeName)

	inputProperties := map[string]schema.PropertySpec{}
	requiredInputProperties := codegen.NewStringSet()

	createPath := dd.createMethod.FlatPath
	if createPath == "" {
		createPath = dd.createMethod.Path
	}

	resourceMeta := resources.CloudAPIResource{
		BaseUrl:    g.rest.BaseUrl,
		CreatePath: createPath,
		CreateVerb: dd.createMethod.HttpMethod,
		NoDelete:   dd.deleteMethod == nil,
	}

	for _, name := range codegen.SortedKeys(dd.createMethod.Parameters) {
		param := dd.createMethod.Parameters[name]
		required := param.Required || strings.HasPrefix(param.Description, "Required.")
		if param.Location != "query" || isDeprecated(param.Description) {
			continue
		}

		p := resources.CloudAPIResourceParam{
			Name:     name,
			Location: "query",
		}
		sdkName := ToLowerCamel(name)
		if sdkName != name {
			p.SdkName = sdkName
		}
		resourceMeta.CreateParams = append(resourceMeta.CreateParams, p)

		inputProperties[sdkName] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{Type: "string"},
		}
		if required {
			requiredInputProperties.Add(sdkName)
		}
	}

	subMatches := pathRegex.FindAllStringSubmatch(createPath, -1)
	for _, names := range subMatches {
		if len(names) < 2 {
			return errors.Errorf("failed to match create path %q", createPath)
		}

		name := names[1]
		sdkName := apiNameToSdkName(name)
		inputProperties[sdkName] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{Type: "string"},
		}
		param := resources.CloudAPIResourceParam{
			Name:     name,
			Location: "path",
		}
		if sdkName != name {
			param.SdkName = sdkName
		}
		resourceMeta.CreateParams = append(resourceMeta.CreateParams, param)
		requiredInputProperties.Add(sdkName)
	}

	if dd.createMethod.Request != nil {
		// If the request type matches the pattern when it contains a property of the type equal to the response
		// type of the GET endpoint, then we want to flatten that property, so that the resource inputs are a superset
		// of the resource outputs. This also helps reconcile the shape of create and update operations.
		var flatten string
		createRequest := g.rest.Schemas[dd.createMethod.Request.Ref]
		if dd.createMethod.Request.Ref != dd.getMethod.Response.Ref {
			for _, name := range codegen.SortedKeys(createRequest.Properties) {
				v := createRequest.Properties[name]
				if v.Ref == dd.getMethod.Response.Ref {
					flatten = name
				}
			}
		}

		bodyBag, err := g.genProperties(typeName, &createRequest, flatten, false)
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

		if dd.updateMethod != nil {
			resourceMeta.UpdateVerb = dd.updateMethod.HttpMethod
			var updateFlatten string
			updateRequest := g.rest.Schemas[dd.updateMethod.Request.Ref]
			if dd.updateMethod.Request.Ref != dd.getMethod.Response.Ref {
				for _, name := range codegen.SortedKeys(updateRequest.Properties) {
					v := updateRequest.Properties[name]
					if v.Ref == dd.getMethod.Response.Ref {
						updateFlatten = name
					}
				}
			}
			resourceMeta.UpdateProperties = map[string]resources.CloudAPIProperty{}
			updateBag, err := g.genProperties(typeName, &updateRequest, updateFlatten, false)
			if err != nil {
				return err
			}

			for name, value := range updateBag.properties {
				if _, has := bodyBag.properties[name]; has {
					resourceMeta.UpdateProperties[name] = value
				} else {
					// TODO: do we need to handle masks?
					if !strings.HasSuffix(name, "Mask") {
						fmt.Printf("unknown update property %s: %s.%s\n", resourceTok, dd.updateMethod.Request.Ref, name)
					}
				}
			}
		}
	}

	properties := map[string]schema.PropertySpec{}
	requiredProperties := codegen.NewStringSet()
	if dd.getMethod.Response != nil {
		response := g.rest.Schemas[dd.getMethod.Response.Ref]
		responseBag, err := g.genProperties(typeName, &response, "", true)
		if err != nil {
			return err
		}
		properties = responseBag.specs
		requiredProperties = responseBag.requiredSpecs
		g.escapeCSharpNames(typeName, properties)

		// Decide how the provider will determine the ID of a created resource.
		// Option 1: it's directly returned from the API in the "self" property.
		for _, p := range []string{"selfLink", "self"} {
			if _, has := response.Properties[p]; has {
				resourceMeta.IdProperty = p
			}
		}

		// Option 2: the provider has to manually build it from the GET method path.
		if resourceMeta.IdProperty == "" {
			idPath := dd.getMethod.FlatPath
			if idPath == "" {
				idPath = dd.getMethod.Path
			}
			v, err := g.buildIdParams(typeName, idPath, inputProperties, &response)
			if err != nil {
				fmt.Printf("Failed to build ID params for resource %s: %v\n", resourceTok, err)
				return nil
			}
			resourceMeta.IdPath = idPath
			resourceMeta.IdParams = v
		}
	}

	if resourceTok == "google-native:storage/v1:Object" {
		resourceTok = "google-native:storage/v1:BucketObject"
		inputProperties["source"] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{
				Ref: "pulumi.json#/Asset",
			},
		}
	}

	resourceSpec := schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: dd.createMethod.Description,
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

// buildIdParams creates a map of parameters that are needed to build resource IDs from an ID path template.
// Keys are API parameter names, values are SDK property names (that may be equal to keys, or not).
func (g *packageGenerator) buildIdParams(typeName string, idPath string, inputProperties map[string]schema.PropertySpec, response *discovery.JsonSchema) (map[string]string, error) {
	result := map[string]string{}

	subMatches := pathRegex.FindAllStringSubmatch(idPath, -1)
	for idx, names := range subMatches {
		if len(names) < 2 {
			return nil, errors.Errorf("failed to match id path %q", idPath)
		}

		name := names[1]

		// If the property is already defined in the input args, add its SDK name.
		sdkName := apiNameToSdkName(name)
		if _, has := inputProperties[sdkName]; has {
			result[name] = sdkName
			continue
		}

		// If the property is already defined in the response, add its API name.
		// Note: API name equals to SDK name for all response properties, currently.
		if _, has := response.Properties[name]; has {
			result[name] = name
			continue
		}

		// The last parameter represents the resource name. It is not a part of the create path,
		// and isn't always modelled as an explicit property in the SDKs. If it's missing in the SDK,
		// it's assigned by the API, and the provider needs to infer its value from the response.
		isLast := idx == len(subMatches)-1
		if !isLast {
			return nil, errors.Errorf("property %q not found in %q", name, idPath)
		}

		// First, check if we have an explicit annotation for which property to use for name resolution.
		key := fmt.Sprintf("%s:%s.%s", g.mod, typeName, name)
		if v, has := resourceNamePropertyOverrides[key]; has {
			if !g.schemaContainsProperty(response, v) {
				return nil, errors.Errorf("property %q not found in response schema of %q", v, key)
			}
			result[name] = v
			// Return the result because we get here only for the last match.
			return result, nil
		}

		// Then, check if there is a property called "name".
		if _, has := response.Properties["name"]; has {
			result[name] = "name"
			// Return the result because we get here only for the last match.
			return result, nil
		}

		// Give up if none is present.
		return nil, errors.Errorf("property %q not found in %q", name, idPath)
	}

	return result, nil
}

func (g *packageGenerator) schemaContainsProperty(schema *discovery.JsonSchema, name string) bool {
	current := *schema
	parts := strings.Split(name, ".")
	for idx, part := range parts {
		value, ok := current.Properties[part]
		if !ok {
			return false
		}
		if idx == len(parts)-1 {
			return true
		}
		current = g.rest.Schemas[value.Ref]
	}
	return false
}

func (g *packageGenerator) genProperties(typeName string, typeSchema *discovery.JsonSchema, flatten string, isOutput bool) (*propertyBag, error) {
	result := propertyBag{
		specs:         map[string]schema.PropertySpec{},
		requiredSpecs: codegen.NewStringSet(),
		properties:    map[string]resources.CloudAPIProperty{},
	}
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		value := typeSchema.Properties[name]
		sdkName := apiNameToSdkName(name)

		if isDeprecated(value.Description) {
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
			sub, err := g.genProperties(typeName, &subtypeSchema, "", isOutput)
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

		typeSpec, err := g.genTypeSpec(typeName, sdkName, &prop, isOutput)
		if err != nil {
			return nil, err
		}

		if prop.Required || isOutput {
			result.requiredSpecs.Add(sdkName)
		}

		description := strings.TrimPrefix(prop.Description, "Output only. ")

		result.specs[sdkName] = schema.PropertySpec{
			Description: description,
			TypeSpec:    *typeSpec,
		}
		apiProp := resources.CloudAPIProperty{}
		if name != sdkName {
			apiProp.SdkName = sdkName
		}
		result.properties[name] = apiProp
	}
	return &result, nil
}

func (g *packageGenerator) genTypeSpec(typeName, propName string, prop *discovery.JsonSchema, isOutput bool) (*schema.TypeSpec, error) {
	switch {
	case prop.Items != nil:
		items, err := g.genTypeSpec(typeName, propName + "Item", prop.Items, isOutput)
		if err != nil {
			return nil, err
		}

		return &schema.TypeSpec{
			Type:  "array",
			Items: items,
		}, nil
	case prop.Type == "any":
		return &schema.TypeSpec{Ref: "pulumi.json#/Any"}, nil
	case prop.Type == "object" && len(prop.Properties) > 0:
		schemaName := fmt.Sprintf(`%s%s`, typeName, strings.Title(propName))
		if isOutput {
			schemaName += "Response"
		}
		tok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, schemaName)
		if _, has := g.rest.Schemas[schemaName]; has {
			return nil, errors.Errorf("properties type name %q conflicts with a schema type", tok)
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)
		bag, err := g.genProperties(typeName, prop, "", isOutput)
		if err != nil {
			return nil, err
		}
		g.pkg.Types[tok] = schema.ComplexTypeSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: prop.Description,
				Type:        "object",
				Properties:  bag.specs,
				Required:    bag.requiredSpecs.SortedValues(),
			},
		}
		return &schema.TypeSpec{
			Type: "object",
			Ref:  referencedTypeName,
		}, nil
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
			bag, err := g.genProperties(typeName, &typeSchema, "", isOutput)
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

// isDeprecated returns true if the description of a property or a parameter signals that it has
// been deprecated.
func isDeprecated(description string) bool {
	// Unfortunately, there's no fixed format for the deprecation message.
	// Also, the word "deprecated" appears in several legit descriptions, so
	// we need that tricky exclusion list below to filter out the actual deprecations.
	lowerDesc := strings.ToLower(description)
	return strings.HasPrefix(lowerDesc, "deprecated") ||
		strings.Contains(lowerDesc, ". deprecated") ||
		strings.Contains(description, ". @Deprecated") ||
		strings.Contains(description, "(Deprecated") ||
		strings.Contains(description, "Deprecated;") ||
		strings.Contains(description, "Deprecated.") ||
		strings.Contains(description, "[Deprecated]") ||
		strings.Contains(description, "Deprecated in favor") ||
		strings.Contains(description, "This field is deprecated")
}

func rawMessage(v interface{}) json.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
