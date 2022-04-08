package gen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gedex/inflector"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/stretchr/testify/require"
	"google.golang.org/api/discovery/v1"
)

func TestCoverageHeuristics(t *testing.T) {
	var fileNames []string
	root := path.Join(".", "discovery")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileNames = append(fileNames, path)
		}
		return nil
	})
	require.NoError(t, err)
	for _, fileName := range fileNames {
		document, err := readDiscoveryDocument(fileName)
		require.NoError(t, err)

		if document.Name == "" {
			continue
		}
		module := fmt.Sprintf("%s/%s", document.Name, document.Version)

		res, err := findResourcesTest(fileName, document.Resources)
		require.NoError(t, err)
	}
}

// findResources builds a map of resources for a given Discovery Document resource. Map keys are resource names and
// values are pointers to CRUD REST methods.
func findResourcesTest(
	docName string,
	rest map[string]discovery.RestResource,
) (map[string]discoveryDocumentResource, error) {
	result := map[string]discoveryDocumentResource{}
	add := func(typeName string, dd discoveryDocumentResource) error {
		return addFoundResource(result, typeName, dd)
	}
	err := findResourcesTestImpl(docName, "", rest, add)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func findResourcesTestImpl(docName, parentName string, rest map[string]discovery.RestResource,
	handleFunc func(string, discoveryDocumentResource) error) error {
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
			case "patch":
				updateMethod = &restMethod
			case "update", "replaceService" /*special case for Cloud Run*/ :
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
						getMethod:    &getIamPolicy,
						updateMethod: &restMethod,
					}
					err := handleFunc(typeName, dd)
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
				return errors.Errorf(
					"get method %q is not supported: %s (%s)", getMethod.HttpMethod, typeName, docName)
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
				getMethod:    getMethod,
				updateMethod: updateMethod,
				deleteMethod: deleteMethod,
			}
			err := handleFunc(typeName, dd)
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
		err := findResourcesImpl(docName, newParent, res.Resources, handleFunc)
		if err != nil {
			return err
		}
	}
	return nil
}
