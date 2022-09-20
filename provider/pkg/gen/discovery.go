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
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/gedex/inflector"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"google.golang.org/api/discovery/v1"
)

type discoveryCRUDMethods struct {
	createMethod, getMethod, updateMethod, deleteMethod, listMethod *discovery.RestMethod
}

// discoveryDocumentResource is a combination of REST methods that should be treated as
// a fully capable Pulumi resource for the generation purpose.
type discoveryDocumentResource struct {
	discoveryCRUDMethods
	resourceName   string
	hasIAMOverlays bool
}

// operation encapsulates information to check for delayed operations' status
// in response to resource CRUD operations.
type operation struct {
	restMethod *discovery.RestMethod
	schema     *discovery.JsonSchema
}

// findResources builds a map of resources for a given Discovery Document resource. Map keys are resource names and
// values are pointers to CRUD REST methods.
func findResources(
	docName string,
	rest map[string]discovery.RestResource,
	schemas map[string]discovery.JsonSchema,
) (map[string]discoveryDocumentResource, map[string]*operation, error) {
	resources := map[string]discoveryDocumentResource{}
	operations := map[string]*operation{}
	addResource := func(typeName string, dd discoveryDocumentResource) error {
		if err := addFoundResource(resources, typeName, dd); err != nil {
			return fmt.Errorf("failed to add resource for %q: %w", docName, err)
		}
		return nil
	}
	addOperation := func(typeName string, dd discoveryDocumentResource) error {
		if strings.HasSuffix(docName, "appengine_v1.json") ||
			strings.HasSuffix(docName, "appengine_v1alpha.json") ||
			strings.HasSuffix(docName, "appengine_v1beta.json") {
			dd.getMethod.Path = strings.ReplaceAll(dd.getMethod.Path, `{appsId}/operations/{operationsId}`, `{+name}`)
			dd.getMethod.FlatPath = strings.ReplaceAll(dd.getMethod.FlatPath, `{appsId}/operations/{operationsId}`,
				`{+name}`)
			dd.getMethod.Parameters = map[string]discovery.JsonSchema{
				"name": {Description: "`name` encapsulates both `appsId` and `operationsId`", Location: "path",
					Required: true,
					Type:     "string"},
			}
		}
		return addFoundOperation(operations, typeName, dd, schemas)
	}
	err := findResourcesImpl(docName, "", rest, addResource, addOperation)
	if err != nil {
		return nil, nil, err
	}
	return resources, operations, nil
}

func findResourcesImpl(docName, parentName string, rest map[string]discovery.RestResource,
	addResource func(string, discoveryDocumentResource) error,
	addOperation func(string, discoveryDocumentResource) error) error {
	var postMethods []discovery.RestMethod

	candidateResources := map[string]discoveryDocumentResource{}

	type pair struct {
		typeName     string
		resourceName string
	}
	type pairs []pair

	var sortedKeys pairs

	for _, resourceName := range codegen.SortedKeys(rest) {
		res := rest[resourceName]
		name := ToUpperCamel(inflector.Singularize(resourceName))

		var createMethod, getMethod, updateMethod, deleteMethod, listMethod *discovery.RestMethod
		sortedMethods := codegen.SortedKeys(res.Methods)
		for _, methodName := range sortedMethods {
			restMethod := res.Methods[methodName]
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
			case "patch":
				updateMethod = &restMethod
			case "update", "replaceService", /*special case for Cloud Run*/
				"partialUpdateInstance" /*special case for bigtable*/ :
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
				typeName := fmt.Sprintf("%s%sIamPolicy", parentName, name)
				if getIamPolicy, has := res.Methods["getIamPolicy"]; has {
					dd := discoveryDocumentResource{
						discoveryCRUDMethods: discoveryCRUDMethods{
							createMethod: &restMethod,
							getMethod:    &getIamPolicy,
							updateMethod: &restMethod,
							listMethod:   listMethod,
						},
						resourceName:   resourceName,
						hasIAMOverlays: true,
					}
					err := addResource(typeName, dd)
					if err != nil {
						return err
					}
				}
			case "delete", "dropDatabase" /*special case for Spanner Database*/ :
				deleteMethod = &restMethod
			}
		}

		if strings.Contains(strings.ToLower(name), "operation") {
			if getMethod != nil {
				operationGetPath := getMethod.FlatPath
				if operationGetPath == "" {
					operationGetPath = getMethod.Path
				}
				if override, has := resourceNameByPathOverrides[fmt.Sprintf("%s:%s", docName, operationGetPath)]; has {
					if override == "" {
						continue
					}
					name = override
				}
			}
			if err := addOperation(name, discoveryDocumentResource{
				discoveryCRUDMethods: discoveryCRUDMethods{
					createMethod: createMethod,
					deleteMethod: deleteMethod,
					updateMethod: updateMethod,
					getMethod:    getMethod,
					listMethod:   listMethod,
				},
				resourceName: resourceName,
			}); err != nil {
				log.Printf("Failed to add operation %q in file %q: %s", name, docName, err)
				continue
			}
		} else {
			sortedKeys = append(sortedKeys, pair{typeName: name, resourceName: resourceName})
			candidateResources[name] = discoveryDocumentResource{
				discoveryCRUDMethods: discoveryCRUDMethods{
					createMethod: createMethod,
					deleteMethod: deleteMethod,
					updateMethod: updateMethod,
					getMethod:    getMethod,
					listMethod:   listMethod,
				},
				resourceName: resourceName,
			}
		}
	}

	sort.SliceStable(sortedKeys, func(i, j int) bool {
		return sortedKeys[i].resourceName < sortedKeys[j].resourceName
	})

	for _, pair := range sortedKeys {
		typeName := pair.typeName
		methods := candidateResources[typeName]
		switch {
		case methods.createMethod != nil && methods.getMethod != nil:
			if methods.getMethod.HttpMethod != "GET" {
				return errors.Errorf(
					"get method %q is not supported: %s (%s)", methods.getMethod.HttpMethod, typeName, docName)
			}
			path := methods.createMethod.FlatPath
			if path == "" {
				path = methods.createMethod.Path
			}
			if override, has := resourceNameByPathOverrides[fmt.Sprintf("%s:%s", docName, path)]; has {
				if override == "" {
					continue
				}
				typeName = override
			}
			ref := methods.createMethod.Response.Ref
			if methods.createMethod.Request != nil {
				ref = methods.createMethod.Request.Ref
			}
			if override, has := resourceNameByTypeOverrides[ref]; has {
				if override == "" {
					continue
				}
				typeName = override
			}

			dd := discoveryDocumentResource{
				discoveryCRUDMethods: discoveryCRUDMethods{
					createMethod: methods.createMethod,
					getMethod:    methods.getMethod,
					updateMethod: methods.updateMethod,
					deleteMethod: methods.deleteMethod,
				},
				resourceName: methods.resourceName,
			}
			err := addResource(typeName, dd)
			if err != nil {
				return err
			}
		case methods.createMethod != nil && methods.getMethod == nil && methods.listMethod != nil:
			// List methods return collections, so the shape of response wouldn't match the SDK map.
			// Skip them for now but emit a note.
			fmt.Printf("List methods are not supported: %s (%s), skipping.\n", typeName, docName)
		case (methods.deleteMethod != nil || methods.updateMethod != nil) && len(postMethods) > 0:
			// No create method. We should simply skip these resources
			// as discussed in https://github.com/pulumi/pulumi-google-native/issues/84
		}

		var newParent string
		switch typeName {
		case "Project", "Location", "Zone", "Global":
			newParent = parentName
		default:
			newParent = parentName + typeName
		}
		res := rest[methods.resourceName]
		err := findResourcesImpl(docName, newParent, res.Resources, addResource, addOperation)
		if err != nil {
			return err
		}
	}
	return nil
}

type GetFunction struct {
	ReturnTypeName string
	Method         *discovery.RestMethod
}

func findAdditionalGetFunctions(docName string, document *discovery.RestDescription) []GetFunction {
	switch docName {
	case "discovery/container_v1.json", "discovery/container_v1beta1.json":
		getServerConfigMethod := document.Resources["projects"].Resources["locations"].Methods["getServerConfig"]
		return []GetFunction{
			{
				ReturnTypeName: "ServerConfig",
				Method:         &getServerConfigMethod,
			},
		}
	}
	return nil
}

func addFoundOperation(operationsMap map[string]*operation, typeName string,
	dd discoveryDocumentResource, jsonSchema map[string]discovery.JsonSchema) error {

	for _, method := range []*discovery.RestMethod{
		dd.getMethod, dd.deleteMethod,
	} {
		if method == nil || method.Response == nil || method.Response.Ref == "" {
			continue
		}
		ref := method.Response.Ref
		schema, ok := jsonSchema[ref]
		if !ok {
			return fmt.Errorf("no schema found for %q", ref)
		}
		operationsMap[ref] = &operation{
			restMethod: method,
			schema:     &schema,
		}
		return nil
	}
	return fmt.Errorf("no methods found for Operations type: %q", typeName)
}

func addFoundResource(
	resourceMap map[string]discoveryDocumentResource,
	typeName string,
	dd discoveryDocumentResource,
) error {
	for _, param := range dd.createMethod.Parameters {
		if param.Location == "path" && isDeprecated(param.Description) {
			// TODO: Should this be disabled?
			// If a path parameter is deprecated, the URL is effectively deprecated, so skip this resource.
			return nil
		}
	}

	existing, has := resourceMap[typeName]
	if !has {
		// No conflict - use this resource.
		resourceMap[typeName] = dd
		return nil
	}

	// Check if two conflicting resources represent the same data type. Otherwise, error - we don't support this
	// scenario.
	if existing.createMethod.Response.Ref != dd.createMethod.Response.Ref {
		return errors.Errorf(
			"%q conflict: %q (%q) vs %q (%q)",
			typeName,
			existing.createMethod.Response.Ref,
			existing.createMethod.FlatPath,
			dd.createMethod.Response.Ref,
			dd.createMethod.FlatPath)
	}

	// Parse the paths of both resources to extract the parameters.
	existingParams := findAPIParams(existing)
	newParams := findAPIParams(dd)

	// Check if one set of parameters is preferred over the other, pick the preferred one if so.
	if preferParams(existingParams, newParams) {
		return nil
	}
	if preferParams(newParams, existingParams) {
		resourceMap[typeName] = dd
		return nil
	}

	return errors.Errorf(
		"%q incompatible params: %q vs %q. You might have to override the names in 'resourceNameByPathOverrides' in"+
			" overrides.go", typeName,
		existing.createMethod.FlatPath,
		dd.createMethod.FlatPath)
}

func preferParams(set codegen.StringSet, other codegen.StringSet) bool {
	if set.Contains(other) {
		return true
	}

	diff := other.Subtract(set)
	if len(diff) == 1 && diff.Has("regionId") && set.Has("location") {
		return true
	}

	return false
}

func findAPIParams(dd discoveryDocumentResource) codegen.StringSet {
	path := dd.createMethod.FlatPath
	if path == "" {
		path = dd.createMethod.Path
	}
	params := findParams(path)
	result := codegen.NewStringSet()
	for _, param := range params {
		result.Add(apiParamNameToSdkName(param))
	}
	return result
}

func findParams(path string) (result []string) {
	subMatches := pathRegex.FindAllStringSubmatch(path, -1)
	for _, names := range subMatches {
		if len(names) < 2 {
			panic(fmt.Sprintf("failed to match path parameters for %q", path))
		}
		result = append(result, names[1])
	}
	return
}
