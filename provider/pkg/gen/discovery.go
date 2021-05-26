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
	add := func(typeName string, dd discoveryDocumentResource) {
		addFoundResource(result, typeName, dd)
	}
	err := findResourcesImpl(docName, "", rest, add)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findResourcesImpl(docName, parentName string, rest map[string]discovery.RestResource,
	add func(string, discoveryDocumentResource)) error {
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
					add(typeName, dd)
				}
			case "delete", "dropDatabase" /*special case for Spanner Database*/ :
				deleteMethod = &restMethod
			}
		}

		typeName := fmt.Sprintf("%s%s", parentName, name)
		switch {
		case createMethod != nil && getMethod != nil:
			if getMethod.HttpMethod != "GET" {
				return errors.Errorf("get method %q is not supported: %s (%s)", getMethod.HttpMethod, typeName, docName)
			}
			dd := discoveryDocumentResource{
				createMethod: createMethod,
				getMethod: getMethod,
				updateMethod: updateMethod,
				deleteMethod: deleteMethod,
			}
			add(typeName, dd)
		case strings.Contains(typeName, "Operation"):
			// Operation variants don't need to be made rest.
			continue
		case createMethod != nil && getMethod == nil && listMethod != nil:
			// List methods return collections, so the shape of response wouldn't match the SDK map.
			// Skip them for now but emit a note.
			fmt.Printf("List methods are not supported: %s (%s), skipping.\n", typeName, docName)
		case (deleteMethod != nil || updateMethod != nil) && len(postMethods) > 0:
			// It can be useful to look at this output and look for potential missing rest with unexpected
			// HTTP POST method names.
			fmt.Printf("No create method for resource: %s (%s), skipping.\n", typeName, docName)
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

func addFoundResource(resourceMap map[string]discoveryDocumentResource,typeName string, dd discoveryDocumentResource) {
	for _, param := range dd.createMethod.Parameters {
		if param.Location == "path" && isDeprecated(param.Description) {
			// If a path parameter is deprecated, the URL is effectively deprecated, so skip this resource.
			return
		}
	}

	// TODO: resolve conflicts so that we never override resources with the same name
	resourceMap[typeName] = dd
}
