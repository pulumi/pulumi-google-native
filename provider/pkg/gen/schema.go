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
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gedex/inflector"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	"google.golang.org/api/discovery/v1"
)

// Note - this needs to be kept in sync with the layout in the SDK package
const goBasePath = "github.com/pulumi/pulumi-google-native/sdk/go/google"

// PulumiSchema will generate a Pulumi schema for the given Google Cloud discovery documents.
func PulumiSchema() (*schema.PackageSpec, *resources.CloudAPIMetadata, error) {
	pkg := schema.PackageSpec{
		Name:        "google-native",
		DisplayName: "Google Cloud Native",
		Description: "A native Pulumi package for creating and managing Google Cloud resources.",
		License:     "Apache-2.0",
		Keywords:    []string{"pulumi", "google cloud", "category/cloud", "kind/native"},
		Homepage:    "https://pulumi.com",
		Publisher:   "Pulumi",
		Repository:  "https://github.com/pulumi/pulumi-google-native",
		Config: schema.ConfigSpec{
			Variables: map[string]schema.PropertySpec{
				"project": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The default project to manage resources in. If another project is specified on a resource, it will take precedence.",
				},
				"region": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The default region to manage resources in. If another region is specified on a regional resource, it will take precedence.",
				},
				"zone": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "The default zone to manage resources in. Generally, this zone should be within the default region you specified. If another zone is specified on a zonal resource, it will take precedence.",
				},
				// Google Partner Name for User-Agent.
				"partnerName": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "A Google Partner Name to facilitate partner resource usage attribution.",
				},
				"disablePartnerName": {
					TypeSpec:    schema.TypeSpec{Type: "boolean"},
					Description: "This will disable the Pulumi Partner Name which is used if a custom `partnerName` isn't specified.",
				},
				"appendUserAgent": {
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Description: "Additional user-agent string to append to the default one (<prod_name>/<ver>).",
				},
			},
		},
		Provider: schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Description: "The provider type for the Google Cloud package.",
				Type:        "object",
				Properties: map[string]schema.PropertySpec{
					"project": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "The default project to manage resources in. If another project is specified on a resource, it will take precedence.",
					},
					"region": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "The default region to manage resources in. If another region is specified on a regional resource, it will take precedence.",
					},
					"zone": {
						TypeSpec:    schema.TypeSpec{Type: "string"},
						Description: "The default zone to manage resources in. Generally, this zone should be within the default region you specified. If another zone is specified on a zonal resource, it will take precedence.",
					},
				},
			},
			InputProperties: map[string]schema.PropertySpec{
				"project": {
					TypeSpec: schema.TypeSpec{Type: "string"},
					DefaultInfo: &schema.DefaultSpec{
						Environment: []string{
							"GOOGLE_PROJECT",
							"GOOGLE_CLOUD_PROJECT",
							"GCLOUD_PROJECT",
							"CLOUDSDK_CORE_PROJECT",
						},
					},
					Description: "The default project to manage resources in. If another project is specified on a resource, it will take precedence.",
				},
				"region": {
					TypeSpec: schema.TypeSpec{Type: "string"},
					DefaultInfo: &schema.DefaultSpec{
						Environment: []string{
							"GOOGLE_REGION",
							"GCLOUD_REGION",
							"CLOUDSDK_COMPUTE_REGION",
						},
					},
					Description: "The default region to manage resources in. If another region is specified on a regional resource, it will take precedence.",
				},
				"zone": {
					TypeSpec: schema.TypeSpec{Type: "string"},
					DefaultInfo: &schema.DefaultSpec{
						Environment: []string{
							"GOOGLE_ZONE",
							"GCLOUD_ZONE",
							"CLOUDSDK_COMPUTE_ZONE",
						},
					},
					Description: "The default zone to manage resources in. Generally, this zone should be within the default region you specified. If another zone is specified on a zonal resource, it will take precedence.",
				},
				// Google Partner Name for User-Agent.
				"partnerName": {
					TypeSpec: schema.TypeSpec{Type: "string"},
					DefaultInfo: &schema.DefaultSpec{
						Environment: []string{
							"GOOGLE_PARTNER_NAME",
						},
					},
					Description: "A Google Partner Name to facilitate partner resource usage attribution.",
				},
				"disablePartnerName": {
					TypeSpec: schema.TypeSpec{Type: "boolean"},
					DefaultInfo: &schema.DefaultSpec{
						Environment: []string{
							"GOOGLE_DISABLE_PARTNER_NAME",
						},
					},
					Description: "This will disable the Pulumi Partner Name which is used if a custom `partnerName` isn't specified.",
				},
				"appendUserAgent": {
					TypeSpec: schema.TypeSpec{Type: "string"},
					DefaultInfo: &schema.DefaultSpec{
						Environment: []string{
							"GOOGLE_APPEND_USER_AGENT",
						},
					},
					Description: "Additional user-agent string to append to the default one (<prod_name>/<ver>).",
				},
			},
		},
		Types:     map[string]schema.ComplexTypeSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Functions: map[string]schema.FunctionSpec{},
		Language:  map[string]schema.RawMessage{},
	}
	metadata := resources.CloudAPIMetadata{
		Types:     map[string]resources.CloudAPIType{},
		Resources: map[string]resources.CloudAPIResource{},
		Functions: map[string]resources.CloudAPIFunction{},
	}

	err := genMixins(&pkg)
	if err != nil {
		return nil, nil, err
	}

	csharpNamespaces := map[string]string{
		"google-native": "GoogleNative",
	}
	javaPackages := map[string]string{}
	pythonModuleNames := map[string]string{}
	golangImportAliases := map[string]string{}

	var fileNames []string
	root := path.Join(".", "discovery")
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
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
		javaPackages[module] = fmt.Sprintf("%s.%s", document.Name, document.Version)
		pythonModuleNames[module] = module
		golangImportAliases[filepath.Join(goBasePath, module)] = document.Name

		res, ops, err := findResources(fileName, document.Resources, document.Schemas)
		if err != nil {
			return nil, nil, err
		}

		for _, typeName := range codegen.SortedKeys(res) {
			// Generate the resource itself.
			err := gen.genResource(typeName, res[typeName], ops)
			if err != nil {
				return nil, nil, err
			}
			// Generate a getXyz function for each Xyz resource.
			err = gen.genFunction(typeName, res[typeName].getMethod)
			if err != nil {
				return nil, nil, err
			}
		}

		for _, additionalGetMethods := range findAdditionalGetFunctions(fileName, document) {
			err = gen.genFunction(additionalGetMethods.ReturnTypeName, additionalGetMethods.Method)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	for _, version := range []string{"v1", "v1beta1"} {
		pkg.Functions[fmt.Sprintf("google-native:container/%s:Cluster/getKubeconfig", version)] = schema.FunctionSpec{
			Description: "Generate a kubeconfig for cluster authentication." +
				"\n\nThe kubeconfig generated is automatically stringified for ease of use with the pulumi/kubernetes" +
				" provider.\nThe kubeconfig uses the new `gke-gcloud-auth-plugin` authentication plugin as" +
				" recommended by Google.\n\n" +
				"See for more details:\n" +
				`- https://cloud.google.com/blog/products/containers-kubernetes/kubectl-auth-changes-in-gke`,
			Inputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"__self__": {
						TypeSpec: schema.TypeSpec{
							Ref: fmt.Sprintf("#/resources/google-native:container/%s:Cluster", version),
						},
					},
				},
				Required: []string{"__self__"},
			},
			Outputs: &schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					"kubeconfig": {
						TypeSpec: schema.TypeSpec{
							Type: "string",
						},
					},
				},
				Required: []string{
					"kubeconfig",
				},
			},
		}
	}

	pkg.Language["go"] = rawMessage(map[string]interface{}{
		"importBasePath":               goBasePath,
		"packageImportAliases":         golangImportAliases,
		"liftSingleValueMethodReturns": true,
		"respectSchemaVersion":         true,
	})
	pkg.Language["nodejs"] = rawMessage(map[string]interface{}{
		"dependencies": map[string]string{
			"@pulumi/pulumi": "^3.0.0",
		},
		"readme": `The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud
programs. This provider uses the Google Cloud REST API directly and therefore provides full access to Google Cloud.
The provider is currently in public preview and is not recommended for production deployments yet. Breaking changes
will be introduced in minor version releases.`,
		"liftSingleValueMethodReturns": true,
		"respectSchemaVersion":         true,
	})

	pkg.Language["python"] = rawMessage(map[string]interface{}{
		"moduleNameOverrides": pythonModuleNames,
		"requires": map[string]string{
			"pulumi": ">=3.0.0,<4.0.0",
		},
		"usesIOClasses": true,
		"readme": `The native Google Cloud Provider for Pulumi lets you provision Google Cloud resources in your cloud
programs. This provider uses the Google Cloud REST API directly and therefore provides full access to Google Cloud.
The provider is currently in public preview and is not recommended for production deployments yet. Breaking changes
will be introduced in minor version releases.`,
		"liftSingleValueMethodReturns": true,
		"pyproject": map[string]bool{
			"enabled": true,
		},
		"respectSchemaVersion": true,
	})

	pkg.Language["csharp"] = rawMessage(map[string]interface{}{
		"packageReferences": map[string]string{
			"Pulumi": "3.*",
		},
		"namespaces":                   csharpNamespaces,
		"liftSingleValueMethodReturns": true,
		"respectSchemaVersion":         true,
	})

	pkg.Language["java"] = rawMessage(map[string]interface{}{
		"packages":                     javaPackages,
		"liftSingleValueMethodReturns": true,
		"respectSchemaVersion":         true,
	})

	return &pkg, &metadata, nil
}

func genMixins(pkg *schema.PackageSpec) error {
	// Mixin 'getClientConfig' to read current configuration values.
	if _, has := pkg.Functions["google-native:authorization:getClientConfig"]; has {
		return errors.New("Invoke 'google-native:authorization:getClientConfig' is already defined")
	}
	pkg.Functions["google-native:authorization:getClientConfig"] = schema.FunctionSpec{
		Description: "Use this function to access the current configuration of the native Google provider.",
		Outputs: &schema.ObjectTypeSpec{
			Description: "Configuration values returned by getClientConfig.",
			Properties: map[string]schema.PropertySpec{
				"project": {
					Description: "Project of the current user.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				"region": {
					Description: "Google Cloud region",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				"zone": {
					Description: "Google Cloud zone",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
			Type:     "object",
			Required: []string{"project", "region"},
		},
	}

	// Mixin 'getClientToken' to obtain an Google auth token.
	if _, has := pkg.Functions["google-native:authorization:getClientToken"]; has {
		return errors.New("Invoke 'google-native:authorization:getClientToken' is already defined")
	}
	pkg.Functions["google-native:authorization:getClientToken"] = schema.FunctionSpec{
		Description: "Use this function to get an Google authentication token for the current login context.",
		Outputs: &schema.ObjectTypeSpec{
			Description: "Configuration values returned by getClientToken.",
			Properties: map[string]schema.PropertySpec{
				"accessToken": {
					Description: "The OAuth2 access token used by the client to authenticate against the Google Cloud API.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Secret:      true,
				},
				"expiry": {
					Description: "Expiry is the optional expiration time of the access token. If zero, TokenSource implementations will reuse the same token forever and RefreshToken or equivalent mechanisms for that TokenSource will not be used.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
				"refreshToken": {
					Description: "RefreshToken is a token that's used by the application (as opposed to the user) to refresh the access token if it expires.",
					TypeSpec:    schema.TypeSpec{Type: "string"},
					Secret:      true,
				},
				"tokenType": {
					Description: `The type of auth token. Possible types are "Bearer", "MAC", "Basic".`,
					TypeSpec:    schema.TypeSpec{Type: "string"},
				},
			},
			Type:     "object",
			Required: []string{"accessToken", "tokenType"},
		},
	}

	return nil
}

var titleReplacer = strings.NewReplacer(" ", "", "-", "")
var versionReplacer = strings.NewReplacer("alpha", "Alpha", "beta", "Beta", "v", "V")

func csharpNamespace(document *discovery.RestDescription) string {
	moduleName := strings.Title(document.Name)
	if v, ok := csharpNamespaceOverrides[moduleName]; ok {
		return v
	}

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

func (g *packageGenerator) fullPath(method *discovery.RestMethod, preferPath bool) string {
	if method == nil {
		return ""
	}
	var pathURL string
	if !preferPath && len(method.FlatPath) > 0 {
		pathURL = resources.AssembleURL(g.rest.RootUrl, g.rest.BasePath, method.FlatPath)
	} else {
		pathURL = resources.AssembleURL(g.rest.RootUrl, g.rest.BasePath, method.Path)
	}

	queryParams := url.Values{}
	for param, details := range method.Parameters {
		if isIgnoredQueryParam(param) {
			continue
		}
		if details.Location != "query" || !isRequired(details) {
			continue
		}
		queryParams.Add(param, "{"+param+"}")
	}
	if len(queryParams) > 0 {
		var err error
		pathURL, err = url.QueryUnescape(pathURL + "?" + queryParams.Encode())
		if err != nil {
			return ""
		}
	}
	return pathURL
}

func (g *packageGenerator) genResource(typeName string, dd discoveryDocumentResource,
	ops map[string]*operation) error {
	resourceTok := g.genToken(typeName)

	g.pkg.Types["google-native:iam/v1:Condition"] = iamCondition

	properties := map[string]schema.PropertySpec{}
	requiredProperties := codegen.NewStringSet()
	inputProperties := map[string]schema.PropertySpec{}
	requiredInputProperties := codegen.NewStringSet()

	resourcePath := func(method *discovery.RestMethod) string {
		if method == nil {
			return ""
		}
		var pathURL string
		if len(method.FlatPath) > 0 {
			pathURL = resources.AssembleURL(g.rest.RootUrl, g.rest.BasePath, method.FlatPath)
		} else {
			pathURL = resources.AssembleURL(g.rest.RootUrl, g.rest.BasePath, method.Path)
		}

		return pathURL
	}
	methodPath := func(method *discovery.RestMethod) string {
		if method == nil {
			return ""
		}
		if len(method.FlatPath) > 0 {
			return method.FlatPath
		}

		return method.Path
	}
	createPath := resourcePath(dd.createMethod)

	resourceMeta := resources.CloudAPIResource{
		RootURL: g.rest.BaseUrl,
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				Endpoint: resources.CloudAPIEndpoint{
					Template: g.fullPath(dd.createMethod, false),
				},
				Verb: dd.createMethod.HttpMethod,
			},
		},
		Delete: resources.CloudAPIOperation{},
		Read: resources.CloudAPIOperation{
			Verb: dd.getMethod.HttpMethod,
		},
		Update: resources.UpdateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{},
		},
	}

	// Populate the resource name key for `getIamPolicy` and `setIamPolicy` methods.
	if strings.HasSuffix(typeName, "IamPolicy") {
		switch g.mod {
		case "storage/v1":
			resourceMeta.IamResourceName = "bucket"
		default:
			resourceMeta.IamResourceName = "resource"
		}
	}

	patternParams := codegen.NewStringSet()

	for _, name := range codegen.SortedKeys(dd.createMethod.Parameters) {
		param := dd.createMethod.Parameters[name]
		if isIgnoredQueryParam(name) {
			continue
		}
		if param.Location != "query" {
			continue
		}

		p := resources.CloudAPIResourceParam{
			Name: name,
			Kind: "query",
		}
		sdkName := ToLowerCamel(name)
		if sdkName != name {
			p.SdkName = sdkName
		}
		queryType := param.Type
		if queryType == "" {
			queryType = "string"
		}
		resourceMeta.Create.Endpoint.Values = append(resourceMeta.Create.Endpoint.Values, p)
		patternParams.Add(sdkName)

		inputProperties[sdkName] = schema.PropertySpec{
			Description:      param.Description,
			TypeSpec:         schema.TypeSpec{Type: queryType},
			ReplaceOnChanges: isRequired(param),
		}
		if isRequired(param) {
			requiredInputProperties.Add(sdkName)
			requiredProperties.Add(sdkName)
		}
		properties[sdkName] = inputProperties[sdkName]
	}

	subMatches := pathRegex.FindAllStringSubmatch(createPath, -1)
	for _, names := range subMatches {
		if len(names) < 2 {
			return errors.Errorf("failed to match create path %q", createPath)
		}

		name := names[1]
		sdkName := apiParamNameToSdkName(name)
		inputProperties[sdkName] = schema.PropertySpec{
			TypeSpec:         schema.TypeSpec{Type: "string"},
			ReplaceOnChanges: true, // All path parameters should trigger replace-on-changes.
		}
		properties[sdkName] = inputProperties[sdkName]
		p := resources.CloudAPIResourceParam{
			Name: name,
			Kind: "path",
		}
		if sdkName != name {
			p.SdkName = sdkName
		}
		resourceMeta.Create.Endpoint.Values = append(resourceMeta.Create.Endpoint.Values, p)
		patternParams.Add(sdkName)
		requiredInputProperties.Add(sdkName)
		requiredProperties.Add(sdkName)
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

		bodyBag, err := g.genProperties(typeName, &createRequest, flatten, false, patternParams)
		if err != nil {
			return err
		}

		isOutput := func(desc string) bool {
			lowerDesc := strings.ToLower(desc)
			return isReadOnly(lowerDesc) ||
				!(strings.Contains(lowerDesc, "values include") ||
					strings.Contains(lowerDesc, "value must be specified"))
		}
		for name, prop := range bodyBag.specs {
			// If the create request has a status field, lets skip it from being marked as a required input.
			if dd.createMethod.Request.Ref == dd.getMethod.Response.Ref && name == "status" && isOutput(prop.
				Description) {
				continue
			}
			inputProperties[name] = prop
		}
		for name := range bodyBag.requiredSpecs {
			requiredInputProperties.Add(name)
		}
		resourceMeta.Create.SDKProperties = bodyBag.properties

		if op, ok := ops[dd.createMethod.Response.Ref]; ok {
			g.setOperationsBaseURL(&resourceMeta.Create.CloudAPIOperation, "", op)
		} else {
			response := g.rest.Schemas[dd.createMethod.Response.Ref]
			for propName, prop := range response.Properties {
				if op, ok := ops[prop.Ref]; ok {
					g.setOperationsBaseURL(&resourceMeta.Create.CloudAPIOperation, propName, op)
					break
				}
			}
		}

		if dd.updateMethod != nil {
			resourceMeta.Update.Verb = dd.updateMethod.HttpMethod
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

			resourceMeta.Update.SDKProperties = map[string]resources.CloudAPIProperty{}
			updateBag, err := g.genProperties(typeName, &updateRequest, updateFlatten, false, nil)
			if err != nil {
				return err
			}

			for name, param := range dd.updateMethod.Parameters {
				if param.Format == "google-fieldmask" && isRequired(param) {
					contract.Assert(param.Location == "query")
					resourceMeta.Update.UpdateMask.QueryParamName = name
				}
			}

			for name, value := range updateBag.properties {
				if _, has := bodyBag.properties[name]; has {
					resourceMeta.Update.SDKProperties[name] = value
				} else {
					if value.Format == "google-fieldmask" && value.Required {
						resourceMeta.Update.UpdateMask.BodyPropertyName = name
					} else {
						fmt.Printf("unknown update property %s: %s.%s\n", resourceTok, dd.updateMethod.Request.Ref, name)
					}
				}
			}
		}
	}

	if dd.getMethod.Response != nil {
		response := g.rest.Schemas[dd.getMethod.Response.Ref]
		responseBag, err := g.genProperties(typeName, &response, "", true, nil)
		if err != nil {
			return err
		}

		if err = mergo.Merge(&properties, responseBag.specs); err != nil {
			return err
		}
		if err = mergo.Merge(&requiredProperties, responseBag.requiredSpecs); err != nil {
			return err
		}
		g.escapeCSharpNames(typeName, properties)

		// Decide how the provider will determine the ID of a created resource.
		// Option 1: it's directly returned from the API in the "self" property.
		for _, p := range []string{"selfLink", "self"} {
			if _, has := response.Properties[p]; has {
				resourceMeta.IDProperty = p
				resourceMeta.Read.Endpoint.SelfLinkProperty = p
				resourceMeta.Update.Endpoint.SelfLinkProperty = p
				resourceMeta.Delete.Endpoint.SelfLinkProperty = p
				break
			}
		}

		// Option 2: the provider has to manually build it from the GET method path.
		if resourceMeta.IDProperty == "" {
			var idVals []resources.CloudAPIResourceParam
			queryValNames := codegen.NewStringSet()
			idPath := methodPath(dd.getMethod)
			queryParams := url.Values{}
			for param, details := range dd.getMethod.Parameters {
				// TODO: this may need to be changed to isRequired(details)
				if details.Location != "query" || !details.Required {
					continue
				}
				queryValNames.Add(param)
				queryParams.Add(param, "{"+param+"}")
			}
			if len(queryParams) > 0 {
				idPath, err = url.QueryUnescape(idPath + "?" + queryParams.Encode())
				if err != nil {
					fmt.Printf("Failed to unescape ID params for resource: %s: %v\n", resourceTok, err)
					return nil
				}
			}

			vals, err := g.buildIdParams(typeName, idPath, inputProperties, &response)
			if err != nil {
				fmt.Printf("Failed to build ID params for resource %s: %v\n", resourceTok, err)
				return nil
			}
			// TODO: This logic could be extracted into something more generic. Currently, the value computation
			//       is always based on the GET method, so this may not be universally correct.
			for _, k := range codegen.SortedKeys(vals) {
				v := resources.CloudAPIResourceParam{
					Name:    k,
					SdkName: vals[k],
					Kind:    "path",
				}
				if queryValNames.Has(k) {
					v.Kind = "query"
				}
				idVals = append(idVals, v)
			}
			if p := g.fullPath(dd.getMethod, false); len(p) > 0 {
				resourceMeta.Read.Endpoint.Template = p
				resourceMeta.Read.Endpoint.Values = idVals
			}
			if p := g.fullPath(dd.updateMethod, false); len(p) > 0 {
				resourceMeta.Update.Endpoint.Template = p
				resourceMeta.Update.Endpoint.Values = idVals
			}
			if p := g.fullPath(dd.deleteMethod, false); len(p) > 0 {
				resourceMeta.Delete.Endpoint.Template = p
				resourceMeta.Delete.Endpoint.Values = idVals
			}
			resourceMeta.IDPath = idPath
			resourceMeta.IDParams = vals
		}
	}

	if dd.updateMethod != nil {
		for name, value := range dd.updateMethod.Parameters {
			if value.Format == "google-fieldmask" && isRequired(value) {
				resourceMeta.Update.Endpoint.Values = append(resourceMeta.Update.Endpoint.Values,
					resources.CloudAPIResourceParam{
						Name:    name,
						SdkName: name,
						Kind:    value.Location,
					})
			}
		}

		if dd.updateMethod.Response != nil && dd.updateMethod.Response.Ref != "" {
			if op, ok := ops[dd.updateMethod.Response.Ref]; ok {
				g.setOperationsBaseURL(&resourceMeta.Update.CloudAPIOperation, "", op)
			}
		}
	}

	if d := dd.deleteMethod; d != nil {
		if v := d.HttpMethod; len(v) > 0 {
			resourceMeta.Delete.Verb = v
		} else {
			resourceMeta.Delete.Verb = "DELETE"
		}

		if dd.deleteMethod.Response != nil && dd.deleteMethod.Response.Ref != "" {
			if op, ok := ops[dd.deleteMethod.Response.Ref]; ok {
				g.setOperationsBaseURL(&resourceMeta.Delete, "", op)
			}
		}
	}

	// Detect resources that support media upload and mark them as such.
	if dd.createMethod.MediaUpload != nil && dd.createMethod.MediaUpload.Protocols != nil &&
		dd.createMethod.MediaUpload.Protocols.Simple != nil {
		resourceMeta.Create.Endpoint.Template = resources.AssembleURL(
			g.rest.RootUrl, dd.createMethod.MediaUpload.Protocols.Simple.Path)
		resourceMeta.AssetUpload = true
		inputProperties["source"] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{
				Ref: "pulumi.json#/Asset",
			},
		}
	}

	if strings.Contains(dd.createMethod.Description, "multipart/form-data") && dd.createMethod.Request.Ref == "GoogleApiHttpBody" {
		resourceMeta.FormDataUpload = resources.FormDataUpload{
			FormFields: map[string]resources.CloudAPIProperty{
				"data": {},
				"file": {},
			},
		}
		inputProperties["file"] = schema.PropertySpec{
			Description: "File to upload.",
			TypeSpec: schema.TypeSpec{
				Ref: "pulumi.json#/Asset",
			},
		}
	}

	description := dd.createMethod.Description

	// Apply auto-naming.
	autoNameable := !strings.HasSuffix(typeName, "IamPolicy") && !autonameExcludes.Has(resourceTok)
	if autoNameable {
		namePattern, err := namePropertyPattern(inputProperties)
		if err == nil {
			requiredInputProperties.Delete("name")
			resourceMeta.Create.Autoname.FieldName = namePattern
		} else if name, ok := autonameOverrides[fmt.Sprintf("google-native:%s:%s", g.mod, typeName)]; ok {
			requiredInputProperties.Delete(name)
			resourceMeta.Create.Autoname.FieldName = name
		} else {
			description += "\nAuto-naming is currently not supported for this resource."
		}
	}

	// Apply auto-project and auto-location population.
	requiredInputProperties.Delete("project")
	requiredInputProperties.Delete("location")
	requiredInputProperties.Delete("zone")

	resourceSpec := schema.ResourceSpec{
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: description,
			Type:        "object",
			Properties:  properties,
			Required:    requiredProperties.SortedValues(),
		},
		Methods:         map[string]string{},
		InputProperties: inputProperties,
		RequiredInputs:  requiredInputProperties.SortedValues(),
	}

	if md, ok := metadataOverrides[resourceTok]; ok {
		if err := mergo.Merge(&resourceMeta, md, mergo.WithOverride); err != nil {
			return fmt.Errorf("failed to merge metadata for resource: %q", resourceTok)
		}
	}

	switch resourceTok {
	case "google-native:container/v1:Cluster":
		resourceSpec.Methods["getKubeconfig"] = "google-native:container/v1:Cluster/getKubeconfig"
	case "google-native:container/v1beta1:Cluster":
		resourceSpec.Methods["getKubeconfig"] = "google-native:container/v1beta1:Cluster/getKubeconfig"
	}

	// For resources with a `setIamPolicy` method defined, also generate Binding and Member resources to provide
	// more granular alternatives to overwriting the entire policy.
	if dd.hasIAMOverlays {
		resourceMeta := deepcopy.Copy(resourceMeta).(resources.CloudAPIResource)
		// Delete for policy is the same operation as an update.
		resourceMeta.Delete.Endpoint = resourceMeta.Update.Endpoint
		resourceMeta.Delete.SDKProperties = resourceMeta.Update.SDKProperties
		resourceMeta.Delete.Verb = resourceMeta.Update.Verb

		// Override default descriptions for overlay properties if a corresponding description is present in the schema.
		// Note: the layout of these properties is inconsistent across resources, so we check multiple alternatives.
		// It doesn't look like these Policy/Binding types are currently populated in the pkg.Types, so we manually
		// parse the schema maps to get this information for now.
		binding := deepcopy.Copy(iamBindingSpec).(schema.ResourceSpec)
		binding.Description = resourceSpec.Description
		member := deepcopy.Copy(iamMemberSpec).(schema.ResourceSpec)
		member.Description = resourceSpec.Description

		bindingMembers := binding.Properties["members"]
		bindingRole := binding.Properties["role"]
		memberMember := member.Properties["member"]
		memberRole := member.Properties["role"]
		if policy, ok := g.rest.Schemas["Policy"]; ok {
			if bindings, ok := policy.Properties["bindings"]; ok {
				if bindings.Items != nil && bindings.Items.Ref == "Binding" {
					if bindingSchema, ok := g.rest.Schemas["Binding"]; ok {
						if members, ok := bindingSchema.Properties["members"]; ok {
							bindingMembers.Description = members.Description
							binding.Properties["members"] = bindingMembers
							memberMember.Description = members.Description
							member.Properties["member"] = memberMember
						}
						if role, ok := bindingSchema.Properties["role"]; ok {
							bindingRole.Description = role.Description
							binding.Properties["role"] = bindingRole
							memberRole.Description = role.Description
							member.Properties["role"] = memberRole
						}
					}
				}
				if bindings.Items != nil {
					if members, ok := bindings.Items.Properties["members"]; ok {
						bindingMembers.Description = members.Description
						binding.Properties["members"] = bindingMembers
						memberMember.Description = members.Description
						member.Properties["member"] = memberMember
					}
					if role, ok := bindings.Items.Properties["role"]; ok {
						bindingRole.Description = role.Description
						binding.Properties["role"] = bindingRole
						memberRole.Description = role.Description
						member.Properties["role"] = memberRole
					}
				}
			}
		}

		// Add the IamBinding resource overlay
		iamBindingToken := strings.TrimSuffix(resourceTok, "Policy") + "Binding"
		g.pkg.Resources[iamBindingToken] = binding
		g.metadata.Resources[iamBindingToken] = resourceMeta

		// Add the IamMember resource overlay
		iamMemberToken := strings.TrimSuffix(resourceTok, "Policy") + "Member"
		g.pkg.Resources[iamMemberToken] = member
		g.metadata.Resources[iamMemberToken] = resourceMeta
	}

	if resourceMeta.Delete.Undefined() {
		resourceSpec.ObjectTypeSpec.Description +=
			"\nNote - this resource's API doesn't support deletion. When deleted, the resource will persist\n" +
				"on Google Cloud even though it will be deleted from Pulumi state."
	}

	g.pkg.Resources[resourceTok] = resourceSpec
	g.metadata.Resources[resourceTok] = resourceMeta

	return nil
}

func (g *packageGenerator) setOperationsBaseURL(cloudOp *resources.CloudAPIOperation,
	embeddedOperationFieldName string,
	op *operation) {
	if cloudOp.Operations == nil {
		cloudOp.Operations = &resources.Operations{}
	}
	if _, has := op.schema.Properties["selfLink"]; has {
		cloudOp.Operations.SelfLinkProperty = "selfLink"
	} else {
		operationParamToSdkName := func(name string, param discovery.JsonSchema) string {
			switch name {
			case "operation":
				if strings.Contains(param.Description, "(ID of the operation)") {
					return "id"
				}
			case "managedZone":
				return "name"
			default:
				if strings.HasSuffix(name, "Id") {
					name = name[:len(name)-2]
					return inflector.Singularize(name)
				}
			}
			return name
		}
		// Bunch of operations' rest methods have parameters that are not correctly
		// referenced in flatPaths but are in paths.
		cloudOp.Operations.Template = g.fullPath(op.restMethod, true)
		paramNames := codegen.NewStringSet()
		for name := range op.restMethod.Parameters {
			paramNames.Add(name)
		}
		for _, name := range paramNames.SortedValues() {
			value := op.restMethod.Parameters[name]
			cloudOp.Operations.Values = append(cloudOp.Operations.Values,
				resources.CloudAPIResourceParam{
					Name:    name,
					SdkName: operationParamToSdkName(name, value),
					Kind:    value.Location,
				})
		}
	}
	cloudOp.Operations.EmbeddedOperationField = embeddedOperationFieldName
}

func (g *packageGenerator) genFunction(typeName string, getMethod *discovery.RestMethod) error {
	resourceTok := g.genToken("get" + typeName)

	inputProperties := map[string]schema.PropertySpec{}
	requiredInputProperties := codegen.NewStringSet()

	getPath := getMethod.Path
	if getMethod.FlatPath != "" {
		getPath = getMethod.FlatPath
	}

	httpMethod := getMethod.HttpMethod
	if httpMethod == "" {
		httpMethod = "GET"
	}
	functionMeta := resources.CloudAPIFunction{
		URL: resources.CloudAPIEndpoint{
			Template: resources.AssembleURL(g.rest.BaseUrl, getPath),
		},
		Verb: httpMethod,
	}

	for _, name := range codegen.SortedKeys(getMethod.Parameters) {
		param := getMethod.Parameters[name]
		if param.Location != "query" {
			continue
		}

		p := resources.CloudAPIResourceParam{
			Name: name,
			Kind: "query",
		}
		sdkName := ToLowerCamel(name)
		if sdkName != name {
			p.SdkName = sdkName
		}
		functionMeta.URL.Values = append(functionMeta.URL.Values, p)
		queryType := param.Type
		if queryType == "" {
			queryType = "string"
		}
		inputProperties[sdkName] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{Type: queryType},
		}
		if isRequired(param) {
			requiredInputProperties.Add(sdkName)
		}
	}

	subMatches := pathRegex.FindAllStringSubmatch(getPath, -1)
	for _, names := range subMatches {
		if len(names) < 2 {
			return errors.Errorf("failed to match get path %q", getPath)
		}

		name := names[1]
		sdkName := apiParamNameToSdkName(name)
		inputProperties[sdkName] = schema.PropertySpec{
			TypeSpec: schema.TypeSpec{Type: "string"},
		}
		p := resources.CloudAPIResourceParam{
			Name: name,
			Kind: "path",
		}
		if sdkName != name {
			p.SdkName = sdkName
		}
		functionMeta.URL.Values = append(functionMeta.URL.Values, p)
		requiredInputProperties.Add(sdkName)
	}

	properties := map[string]schema.PropertySpec{}
	requiredProperties := codegen.NewStringSet()
	if getMethod.Response != nil {
		response := g.rest.Schemas[getMethod.Response.Ref]
		responseBag, err := g.genProperties(typeName, &response, "", true, nil)
		if err != nil {
			return err
		}
		properties = responseBag.specs
		requiredProperties = responseBag.requiredSpecs
	}

	// Apply auto-project population.
	requiredInputProperties.Delete("project")

	functionSpec := schema.FunctionSpec{
		Description: getMethod.Description,
		Inputs: &schema.ObjectTypeSpec{
			Type:       "object",
			Properties: inputProperties,
			Required:   requiredInputProperties.SortedValues(),
		},
		Outputs: &schema.ObjectTypeSpec{
			Type:       "object",
			Properties: properties,
			Required:   requiredProperties.SortedValues(),
		},
	}
	g.pkg.Functions[resourceTok] = functionSpec
	g.metadata.Functions[resourceTok] = functionMeta
	return nil
}

// buildIdParams creates a map of parameters that are needed to build resource IDs from an ID path template.
// Keys are API parameter names, values are SDK property names (that may be equal to keys, or not).
func (g *packageGenerator) buildIdParams(
	typeName string,
	idPath string,
	inputProperties map[string]schema.PropertySpec,
	response *discovery.JsonSchema,
) (map[string]string, error) {
	result := map[string]string{}

	u, err := url.Parse(idPath)
	if err != nil {
		return nil, err
	}

	// First look for substitutable references in the 'path' part of idPath
	subMatches := pathRegex.FindAllStringSubmatch(u.RawPath, -1)
	for idx, names := range subMatches {
		if len(names) < 2 {
			return nil, errors.Errorf("failed to match id path %q", u.RawPath)
		}

		name := names[1]

		// If the property is already defined in the input args, add its SDK name.
		sdkName := apiParamNameToSdkName(name)
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
			return nil, errors.Errorf("property %q not found in %q", name, u.RawPath)
		}

		// First, check if we have an explicit annotation for which property to use for name resolution.
		key := fmt.Sprintf("%s:%s.%s", g.mod, typeName, name)
		if v, has := resourceNamePropertyOverrides[key]; has {
			if !g.schemaContainsProperty(response, v) {
				return nil, errors.Errorf("property %q not found in response schema of %q", v, key)
			}
			result[name] = v
			// Return the result because we get here only for the last match.
			continue
		}

		// Then, check if there is a property called "name".
		if _, has := response.Properties["name"]; has {
			result[name] = "name"
			continue
		}

		// Give up if none is present.
		return nil, errors.Errorf("property %q not found in %q", name, u.RawPath)
	}

	// Next handle the 'query' part of idPath
	subMatches = pathRegex.FindAllStringSubmatch(u.RawQuery, -1)
	for _, names := range subMatches {
		if len(names) < 2 {
			return nil, errors.Errorf("failed to match id path %q", u.RawPath)
		}

		name := names[1]

		// If the property is already defined in the input args, add its SDK name.
		sdkName := apiParamNameToSdkName(name)
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

		// Give up if none is present.
		return nil, errors.Errorf("property %q not found in %q", name, u.RawQuery)
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

func (g *packageGenerator) genProperties(typeName string, typeSchema *discovery.JsonSchema, flatten string,
	isOutput bool, patternParams codegen.StringSet) (*propertyBag, error) {
	result := propertyBag{
		specs:         map[string]schema.PropertySpec{},
		requiredSpecs: codegen.NewStringSet(),
		properties:    map[string]resources.CloudAPIProperty{},
	}
	for _, name := range codegen.SortedKeys(typeSchema.Properties) {
		// Consult the unsupported properties map to see if we should
		// skip this property altogether.
		if overrides, ok := unsupportedPropertiesOverrides[typeName]; ok {
			if contains(overrides, name) {
				log.Printf("Skipping unsupported property %s in type %s", name, typeName)
				continue
			}
		}

		value := typeSchema.Properties[name]
		sdkName := apiPropNameToSdkName(typeName, name)

		readOnly := value.ReadOnly || isReadOnly(value.Description)
		if !isOutput && readOnly {
			continue
		}
		if isOutput && name == "id" {
			continue
		}

		prop := value

		if name == flatten {
			subtypeSchema := g.rest.Schemas[prop.Ref]
			sub, err := g.genProperties(typeName, &subtypeSchema, "", isOutput, patternParams)
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

		copyFromOutput := g.shouldCopyFromOutput(name) && !isOutput
		if !copyFromOutput {
			if isRequired(prop) || isOutput {
				result.requiredSpecs.Add(sdkName)
			}

			p := schema.PropertySpec{
				Description: clearDescription(prop.Description),
				TypeSpec:    *typeSpec,
			}
			if isDeprecated(prop.Description) {
				p.DeprecationMessage = prop.Description
			}
			result.specs[sdkName] = p
		}

		apiProp := resources.CloudAPIProperty{
			Ref:      typeSpec.Ref,
			Format:   prop.Format,
			Required: isRequired(prop),
			ForceNew: !isOutput && isImmutable(prop.Description),

			Items:                g.itemTypeToProperty(typeSpec.Items),
			AdditionalProperties: g.itemTypeToProperty(typeSpec.AdditionalProperties),
			CopyFromOutputs:      copyFromOutput,
			Pattern:              propertyPattern(sdkName, prop.Description, patternParams),
		}
		if name != sdkName {
			apiProp.SdkName = sdkName
		}
		result.properties[name] = apiProp
	}
	return &result, nil
}

func (g *packageGenerator) shouldCopyFromOutput(propName string) bool {
	mod := strings.Split(g.mod, "/")[0]
	switch mod {
	case "compute", "container", "deploymentmanager":
		return strings.Contains(strings.ToLower(propName), "fingerprint")
	default:
		// Some other modules contain properties called "fingerprint" or similar but that need to be
		// populated manually by users, e.g. to provide a hash of a file or a certificate. Therefore,
		// we maintain the opt-in list of modules above and treat the rest as normal inputs.
		return false
	}
}

func (g *packageGenerator) itemTypeToProperty(typ *schema.TypeSpec) *resources.CloudAPIProperty {
	if typ == nil || typ.Ref == "pulumi.json#/Any" {
		return nil
	}

	return &resources.CloudAPIProperty{
		Ref:                  typ.Ref,
		Items:                g.itemTypeToProperty(typ.Items),
		AdditionalProperties: g.itemTypeToProperty(typ.AdditionalProperties),
	}
}

func (g *packageGenerator) genTypeSpec(typeName, propName string, prop *discovery.JsonSchema, isOutput bool) (*schema.TypeSpec, error) {
	switch {
	case prop.Items != nil:
		items, err := g.genTypeSpec(typeName, propName+"Item", prop.Items, isOutput)
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
		typePropName := fmt.Sprintf(`%s%s`, typeName, strings.Title(propName))
		schemaName := typePropName
		if isOutput {
			schemaName += "Response"
		}
		tok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, schemaName)
		if _, has := g.rest.Schemas[schemaName]; has {
			return nil, errors.Errorf("properties type name %q conflicts with a schema type", tok)
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)
		bag, err := g.genProperties(typePropName, prop, "", isOutput, nil)
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
		g.metadata.Types[tok] = resources.CloudAPIType{
			Properties: bag.properties,
		}
		return &schema.TypeSpec{
			Type: "object",
			Ref:  referencedTypeName,
		}, nil
	case prop.Type == "object" && prop.AdditionalProperties != nil:
		// The prop is a map with a string key and a simple value
		// if it doesn't have a ref.
		if prop.AdditionalProperties.Ref == "" {
			switch prop.AdditionalProperties.Type {
			case "any":
				return &schema.TypeSpec{
					Type:                 "object",
					AdditionalProperties: &schema.TypeSpec{Ref: "pulumi.json#/Any"},
				}, nil
			case "array":
				typeSpec, err := g.genTypeSpec(propName, propName, prop.AdditionalProperties.Items, isOutput)
				if err != nil {
					return nil, err
				}

				return &schema.TypeSpec{
					Type: "object",
					AdditionalProperties: &schema.TypeSpec{
						Type:  "array",
						Items: typeSpec,
					},
				}, nil
			case "nil":
				return nil, errors.New(fmt.Sprintf("nil is not a valid array element type: %v", prop))
			case "object":
				return nil, errors.New(fmt.Sprintf("object is not a valid array element type: %v", prop))
			default:
				return &schema.TypeSpec{
					Type: "object",
					AdditionalProperties: &schema.TypeSpec{
						Type: prop.AdditionalProperties.Type,
					},
				}, nil
			}
		}

		// Otherwise, the value in-turn is a complex type.
		typePropName := fmt.Sprintf(`%s%s`, typeName, ToUpperCamel(propName))
		refTypeSpec, err := g.genTypeSpec(typePropName, propName, prop.AdditionalProperties, isOutput)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error generating type spec for $ref in additional properties %v", err))
		}

		return &schema.TypeSpec{
			Type:                 "object",
			AdditionalProperties: refTypeSpec,
		}, nil
	case len(prop.Enum) > 0 && !isOutput:
		return g.genEnumType(typeName, propName, prop)
	case prop.Type != "":
		return &schema.TypeSpec{Type: prop.Type}, nil
	case prop.Ref != "":
		schemaName := prop.Ref
		tok := fmt.Sprintf(`%s:%s:%s`, g.pkg.Name, g.mod, schemaName)
		if isOutput {
			tok += "Response"
		}
		referencedTypeName := fmt.Sprintf("#/types/%s", tok)

		if !g.visitedTypes.Has(tok) {
			g.visitedTypes.Add(tok)

			typeSchema := g.rest.Schemas[schemaName]
			bag, err := g.genProperties(schemaName, &typeSchema, "", isOutput, nil)
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
			g.metadata.Types[tok] = resources.CloudAPIType{
				Properties: bag.properties,
			}
		}

		return &schema.TypeSpec{
			Type: "object",
			Ref:  referencedTypeName,
		}, nil
	}
	return nil, errors.New("unknown type")
}

// genEnumType generates the enum type.
func (g *packageGenerator) genEnumType(typeName, propName string, prop *discovery.JsonSchema) (*schema.TypeSpec, error) {
	if prop.Type != "string" {
		return nil, errors.Errorf("string-based enum expected but found %q", prop.Type)
	}

	enumName := typeName + ToUpperCamel(propName)
	tok := fmt.Sprintf("%s:%s:%s", g.pkg.Name, g.mod, enumName)

	enumSpec := &schema.ComplexTypeSpec{
		Enum: []schema.EnumValueSpec{},
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: prop.Description,
			Type:        "string",
		},
	}

	values := codegen.NewStringSet()
	for idx, val := range prop.Enum {
		values.Add(val)
		enumVal := schema.EnumValueSpec{
			Value:       val,
			Name:        ToUpperCamel(val),
			Description: prop.EnumDescriptions[idx],
		}
		enumSpec.Enum = append(enumSpec.Enum, enumVal)
	}

	// Make sure that the type name we composed doesn't clash with another type
	// already defined in the schema earlier. The same enum does show up in multiple
	// places of specs, so we want to error only if they a) have the same name
	// b) the list of values does not match.
	if other, ok := g.pkg.Types[tok]; ok {
		same := len(enumSpec.Enum) == len(other.Enum)
		for _, val := range other.Enum {
			same = same && values.Has(val.Value.(string))
		}
		if !same {
			return nil, errors.Errorf("duplicate enum %q", tok)
		}
	}

	g.pkg.Types[tok] = *enumSpec

	referencedTypeName := fmt.Sprintf("#/types/%s", tok)
	return &schema.TypeSpec{
		Ref: referencedTypeName,
	}, nil
}

func (g *packageGenerator) escapeCSharpNames(typeName string, resourceResponse map[string]schema.PropertySpec) {
	for name, swagger := range resourceResponse {
		// C# doesn't allow properties to have the same name as its containing type.
		if strings.Title(name) == typeName {
			swagger.Language = map[string]schema.RawMessage{
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

// isRequired returns true if the property or a parameter indicates that it is required.
func isRequired(parameter discovery.JsonSchema) bool {
	if parameter.Required {
		return true
	}
	return strings.HasPrefix(parameter.Description, "Required.")
}

// isImmutable returns true if the property or a parameter indicates that it is immutable.
func isImmutable(description string) bool {
	return strings.Contains(description, "Immutable.")
}

// isReadOnly returns true if the description of a property or a parameter signals that it is
// an output-only property.
func isReadOnly(description string) bool {
	lowerDesc := strings.ToLower(description)
	return strings.HasPrefix(lowerDesc, "[output only]") ||
		strings.HasPrefix(lowerDesc, "[output-only]") ||
		strings.HasPrefix(lowerDesc, "output only.") ||
		strings.Contains(lowerDesc, "(output only)") ||
		strings.HasSuffix(lowerDesc, "@outputonly")
}

// clearDescription removes annotations like "output only" from description text.
func clearDescription(description string) string {
	description = strings.TrimPrefix(description, "Required. ")
	description = strings.TrimPrefix(description, "Output only. ")
	description = strings.TrimPrefix(description, "[Output Only] ")
	description = strings.TrimPrefix(description, "[Output-only] ")
	description = strings.TrimPrefix(description, "Output only. ")
	description = strings.TrimSuffix(description, "@OutputOnly")
	return description
}

// isIgnoredQueryParam returns true if a query parameter with a given name
// should be ignored.
func isIgnoredQueryParam(name string) bool {
	// "validateOnly" is a special query parameter that allows a dry-run
	// without actually executing an operation. It's always options and
	// should not be projected to SDKs. See https://google.aip.dev/163#guidance.
	return name == "validateOnly"
}

func rawMessage(v interface{}) schema.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}
