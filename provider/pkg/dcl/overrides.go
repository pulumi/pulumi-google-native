package dcl

import (
	"fmt"
	"path/filepath"
	"strings"

	"google.golang.org/api/discovery/v1"

	"github.com/pulumi/pulumi-google-native/provider/pkg/utils"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/vertexai"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const goBasePath = "github.com/pulumi/pulumi-google-native/sdk/go/google"

func AddDCLResources(pkg *schema.PackageSpec, metadata *resources.CloudAPIMetadata,
	pythonModuleNames, goImportAliases, csharpNamespaces, javaPackages map[string]string) error {
	return addVertexAIResources(pkg, metadata, pythonModuleNames, goImportAliases, csharpNamespaces, javaPackages)
}

func csharpVersionedNamespace(moduleName string) string {
	split := strings.Split(moduleName, "/")
	module, version := split[0], split[1]

	version = versionReplacer.Replace(version)
	return fmt.Sprintf("%s.%s", module, version)
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

func addVertexAIResources(
	pkg *schema.PackageSpec,
	metadata *resources.CloudAPIMetadata,
	pythonModuleNames,
	goImportAliases,
	csharpNamespaces,
	javaPackages map[string]string,
) error {
	schemas := map[string]*dcl.Schema{
		"google-native:vertexai/v1:MetadataSchema":       vertexai.DCLMetadataSchemaSchema(),
		"google-native:vertexai/v1:MetadataStore":        vertexai.DCLMetadataStoreSchema(),
		"google-native:vertexai/v1:Endpoint":             vertexai.DCLEndpointSchema(),
		"google-native:vertexai/v1:ModelDeployment":      vertexai.DCLModelDeploymentSchema(),
		"google-native:vertexai/v1:Model":                vertexai.DCLModelSchema(),
		"google-native:vertexai/v1:EndpointTrafficSplit": vertexai.DCLEndpointTrafficSplitSchema(),
	}
	for tok, schema := range schemas {
		if err := processDCLSchema(tok, schema, pkg, metadata); err != nil {
			return err
		}
	}

	const moduleName = "vertexai/v1"
	csharpNamespaces[moduleName] = csharpVersionedNamespace(moduleName)
	javaPackages[moduleName] = strings.Replace(moduleName, "/", ".", 1)
	pythonModuleNames[moduleName] = moduleName
	goImportAliases[filepath.Join(goBasePath, moduleName)] = "vertexai"

	return nil
}

func processDCLSchema(tok string, dclSchema *dcl.Schema, pkg *schema.PackageSpec,
	metadata *resources.CloudAPIMetadata) error {

	var registerPropertyType func(string, *dcl.Property) (*schema.TypeSpec, error)
	registerPropertyType = func(modName string, prop *dcl.Property) (*schema.TypeSpec, error) {
		switch prop.TypeEnum() {
		case dcl.EnumType:
			if len(prop.Enum) > 0 && !prop.ReadOnly {
				contract.Assertf(prop.GoType != "",
					"Property %q in %q for tok: %q doesn't have a GoType as expected for string enums",
					prop.GoName, modName, tok)
				tok := fmt.Sprintf("google-native:%s:%s", modName, prop.GoType)
				referencedTypeName := fmt.Sprintf("#/types/%s", tok)
				if _, has := pkg.Types[tok]; has {
					return &schema.TypeSpec{
						Ref: referencedTypeName,
					}, nil
				}
				enumSpec := schema.ComplexTypeSpec{
					Enum: []schema.EnumValueSpec{},
					ObjectTypeSpec: schema.ObjectTypeSpec{
						Description: prop.Description,
						Type:        "string",
					},
				}

				values := codegen.NewStringSet()
				for _, val := range prop.Enum {
					values.Add(val)
					enumVal := schema.EnumValueSpec{
						Value: val,
						Name:  utils.ToUpperCamel(val),
					}
					enumSpec.Enum = append(enumSpec.Enum, enumVal)
				}
				pkg.Types[tok] = enumSpec
				return &schema.TypeSpec{
					Ref: referencedTypeName,
				}, nil
			}
		case dcl.StringType, dcl.ReferenceType:
			return &schema.TypeSpec{Type: "string"}, nil
		case dcl.BooleanType:
			return &schema.TypeSpec{Type: "boolean"}, nil
		case dcl.IntegerType:
			return &schema.TypeSpec{Type: "integer"}, nil
		case dcl.DoubleType:
			return &schema.TypeSpec{Type: "number"}, nil
		case dcl.MapType:
			contract.Assertf(prop.AdditionalProperties != nil,
				"Expect %q in %q for tok: %q to have additionalProperties for map type", prop.GoName, modName, tok)
			ref, err := registerPropertyType(modName, prop.AdditionalProperties)
			if err != nil {
				return nil, err
			}
			return &schema.TypeSpec{
				Type:                 "object",
				AdditionalProperties: ref,
			}, nil
		case dcl.ObjectType:
			goType := prop.GoType
			if goType == "" {
				goType = prop.GoName
			}
			tok := fmt.Sprintf("google-native:%s:%s", modName, goType)
			props := map[string]schema.PropertySpec{}
			for n, p := range prop.Properties {
				ref, err := registerPropertyType(modName, p)
				if err != nil {
					return nil, err
				}
				props[n] = schema.PropertySpec{
					TypeSpec:         *ref,
					Description:      p.Description,
					ReplaceOnChanges: p.Immutable,
				}
			}
			pkg.Types[tok] = schema.ComplexTypeSpec{
				ObjectTypeSpec: schema.ObjectTypeSpec{
					Description: prop.Description,
					Properties:  props,
					Type:        "object",
					Required:    prop.Required,
				},
			}
			referencedTypeName := fmt.Sprintf("#/types/%s", tok)
			return &schema.TypeSpec{
				Ref: referencedTypeName,
			}, nil
		case dcl.ArrayType:
			contract.Assertf(prop.Items != nil,
				"Property %q in %q for tok: %q doesn't have an Item field as expected for arrays", prop.GoName,
				modName, tok)
			ref, err := registerPropertyType(modName, prop.Items)
			if err != nil {
				return nil, err
			}
			return &schema.TypeSpec{
				Type:  "array",
				Items: ref,
			}, nil
		}
		return nil, fmt.Errorf("unexpected type: %q for tok: %q", prop.Type, tok)
	}

	processComponent := func(moduleName string, component *dcl.Component) error {
		dclRes := component.SchemaProperty
		contract.Assertf(dclRes.Type == "object", "%q schema type should be object, got: %q", tok, dclRes.Type)
		resSpec := schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: dclRes.Description,
				Type:        "object",
				Properties:  map[string]schema.PropertySpec{},
			},
			InputProperties: map[string]schema.PropertySpec{},
			RequiredInputs:  dclRes.Required,
		}

		for prop, dclProp := range dclRes.Properties {
			ref, err := registerPropertyType(moduleName, dclProp)
			if err != nil {
				return err
			}
			resSpec.Properties[prop] = schema.PropertySpec{
				TypeSpec:         *ref,
				Description:      dclProp.Description,
				ReplaceOnChanges: dclProp.Immutable,
			}

			if !dclProp.ReadOnly {
				resSpec.InputProperties[prop] = resSpec.Properties[prop]
			}
		}
		pkg.Resources[tok] = resSpec
		metadata.Resources[tok] = resources.CloudAPIResource{DCLManaged: true}
		return nil
	}

	t, err := tokens.ParseTypeToken(tok)
	if err != nil {
		return err
	}
	moduleName := t.Module().Name().String()
	for _, component := range dclSchema.Components.Schemas {
		if err := processComponent(moduleName, component); err != nil {
			return err
		}
	}
	return nil
}
