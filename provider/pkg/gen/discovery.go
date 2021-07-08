// Copyright 2016-2021, Pulumi Corporation.

package gen

import (
	"fmt"
	"github.com/gedex/inflector"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"google.golang.org/api/discovery/v1"
	"strings"
)

// discoveryDocumentResource is a combination of REST methods that should be treated as
// a fully capable Pulumi resource for the generation purpose.
type discoveryDocumentResource struct {
	createMethod, getMethod, updateMethod, deleteMethod *discovery.RestMethod
}

// findResources builds a map of resources for a given Discovery Document resource. Map keys are resource names and
// values are pointers to CRUD REST methods.
func findResources(docName string, rest map[string]discovery.RestResource) (map[string]discoveryDocumentResource, error) {
	result := map[string]discoveryDocumentResource{}
	add := func(typeName string, dd discoveryDocumentResource) error {
		return addFoundResource(result, typeName, dd)
	}
	err := findResourcesImpl(docName, "", rest, add)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findResourcesImpl(docName, parentName string, rest map[string]discovery.RestResource,
	add func(string, discoveryDocumentResource) error) error {
	var postMethods []discovery.RestMethod
	for _, resourceName := range codegen.SortedKeys(rest) {
		res := rest[resourceName]
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
				typeName := fmt.Sprintf("%s%sIamPolicy", parentName, name)
				if getIamPolicy, has := res.Methods["getIamPolicy"]; has {
					dd := discoveryDocumentResource{
						createMethod: &restMethod,
						getMethod: &getIamPolicy,
						updateMethod: &restMethod,
					}
					err := add(typeName, dd)
					if err != nil {
						return err
					}
				}
			case "delete", "dropDatabase" /*special case for Spanner Database*/ :
				deleteMethod = &restMethod
			}
		}

		typeName := name
		switch {
		case createMethod != nil && getMethod != nil:
			if getMethod.HttpMethod != "GET" {
				return errors.Errorf("get method %q is not supported: %s (%s)", getMethod.HttpMethod, typeName, docName)
			}
			path := createMethod.FlatPath
			if path == "" {
				path = createMethod.Path
			}
			if override, has := resourceNameByPathOverrides[path]; has {
				if override == "" {
					return nil
				}
				typeName = override
			}
			ref := createMethod.Response.Ref
			if createMethod.Request != nil {
				ref = createMethod.Request.Ref
			}
			if override, has := resourceNameByTypeOverrides[ref]; has {
				if override == "" {
					return nil
				}
				typeName = override
			}

			dd := discoveryDocumentResource{
				createMethod: createMethod,
				getMethod: getMethod,
				updateMethod: updateMethod,
				deleteMethod: deleteMethod,
			}
			err := add(typeName, dd)
			if err != nil {
				return err
			}
		case strings.Contains(typeName, "Operation"):
			// Operation variants don't need to be made rest.
			continue
		case createMethod != nil && getMethod == nil && listMethod != nil:
			// List methods return collections, so the shape of response wouldn't match the SDK map.
			// Skip them for now but emit a note.
			fmt.Printf("List methods are not supported: %s (%s), skipping.\n", typeName, docName)
		case (deleteMethod != nil || updateMethod != nil) && len(postMethods) > 0:
			// No create method. We should simply skip these resources
			// as discussed in https://github.com/pulumi/pulumi-google-native/issues/84
		}

		var newParent string
		switch name {
		case "Project", "Location", "Zone", "Global":
			newParent = parentName
		default:
			newParent = parentName + name
		}
		err := findResourcesImpl(docName, newParent, res.Resources, add)
		if err != nil {
			return err
		}
	}
	return nil
}

func addFoundResource(resourceMap map[string]discoveryDocumentResource, typeName string, dd discoveryDocumentResource) error {
	for _, param := range dd.createMethod.Parameters {
		if param.Location == "path" && isDeprecated(param.Description) {
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

	// Check if two conflicting resources represent the same data type. Otherwise, error - we don't support this scenario.
	if existing.createMethod.Response.Ref != dd.createMethod.Response.Ref {
		return errors.Errorf("%q conflict: %q (%q) vs %q (%q)", typeName, existing.createMethod.Response.Ref, existing.createMethod.FlatPath, dd.createMethod.Response.Ref, dd.createMethod.FlatPath)
	}

	// Parse the paths of both resources to extract the parameters.
	existingParams := findApiParams(existing)
	newParams := findApiParams(dd)

	// Check if one set of parameters is preferred over the other, pick the preferred one if so.
	if preferParams(existingParams, newParams) {
		return nil
	}
	if preferParams(newParams, existingParams) {
		resourceMap[typeName] = dd
		return nil
	}

	return errors.Errorf("%q incompatible params: %q vs %q", typeName, existing.createMethod.FlatPath, dd.createMethod.FlatPath)
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

func findApiParams(dd discoveryDocumentResource) codegen.StringSet {
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
