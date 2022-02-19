// Copyright 2016-2021, Pulumi Corporation.
package provider

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func calculateResourceId(res resources.CloudAPIResource, inputs map[string]interface{}, outputs map[string]interface{}) (string, error) {
	if res.IdProperty != "" {
		v, ok := outputs[res.IdProperty].(string)
		if !ok {
			return "", errors.Errorf("ID property %q not found", res.IdProperty)
		}
		return strings.TrimPrefix(v, res.BaseUrl), nil
	}

	id := res.IdPath
	idParams := res.IdParams
	for name, alias := range idParams {
		var propValue string
		if v, has := evalPropertyValue(inputs, alias); has {
			propValue = v
		} else if v, has := evalPropertyValue(outputs, alias); has {
			propValue = v
		}

		if propValue == "" {
			return "", errors.Errorf("property %q/%q not found", name, alias)
		}

		// The name property can sometimes contain multiple segments. We only care about the last one
		// because we flattened the path while building metadata.
		parts := strings.Split(propValue, "/")
		propValue = parts[len(parts)-1]

		id = strings.Replace(id, fmt.Sprintf("{%s}", name), propValue, 1)
	}
	return id, nil
}

func evalPropertyValue(values map[string]interface{}, path string) (string, bool) {
	current := values
	parts := strings.Split(path, ".")
	for idx, part := range parts {
		value := current[part]
		if idx == len(parts)-1 {
			if str, ok := value.(string); ok {
				return str, true
			}
			return "", false
		}
		childMap, ok := value.(map[string]interface{})
		if !ok {
			return "", false
		}
		current = childMap
	}
	return "", false
}

// buildCreateUrl composes the URL to invoke to create a resource with given inputs.
func buildCreateUrl(res resources.CloudAPIResource, inputs resource.PropertyMap) (string, error) {
	uri := res.ResourceUrl(res.CreatePath)
	return buildUrl(uri, res.CreateParams, inputs)
}

func buildFunctionUrl(res resources.CloudAPIFunction, inputs resource.PropertyMap) (string, error) {
	return buildUrl(res.Url, res.Params, inputs)
}

func buildUrl(uriTemplate string, params []resources.CloudAPIResourceParam, inputs resource.PropertyMap) (string, error) {
	queryMap := map[string]string{}
	uriString := uriTemplate
	for _, param := range params {
		sdkName := param.Name
		if param.SdkName != "" {
			sdkName = param.SdkName
		}
		key := resource.PropertyKey(sdkName)
		if !inputs[key].HasValue() {
			continue
		}

		value := inputs[key].StringValue()
		switch param.Location {
		case "path":
			uriString = strings.Replace(uriString, fmt.Sprintf("{%s}", param.Name), url.PathEscape(value), 1)
		case "query":
			queryMap[param.Name] = value
		default:
			return "", errors.Errorf("unknown param location %q", param.Location)
		}
	}
	uri, err := url.Parse(uriString)
	if err != nil {
		return "", errors.Wrapf(err, "parsing resource URL %q", uriString)
	}
	query := uri.Query()
	for key, value := range queryMap {
		query.Add(key, value)
	}
	uri.RawQuery = query.Encode()
	return uri.String(), nil
}

// getDefaultName generates a random name for a resource based on its URN name, a given pattern,
// and other properties.
func getDefaultName(urn resource.URN, pattern string,
	olds, news resource.PropertyMap) resource.PropertyValue {
	if v, ok := olds[resource.PropertyKey("name")]; ok {
		return v
	}

	name := urn.Name().String()

	// Resource name is URN name + "-" + random suffix.
	random, err := resource.NewUniqueHex(name+"-", 7, 0)
	contract.AssertNoError(err)

	result := strings.Replace(pattern, "{name}", random, 1)
	for key, value := range news {
		if !value.HasValue() || !value.IsString() {
			continue
		}

		part := fmt.Sprintf("{%s}", string(key))
		result = strings.Replace(result, part, value.StringValue(), 1)
	}

	return resource.NewStringProperty(result)
}

// applyPropertyPattern composes a pattern-based string value from the given property map, if the property
// has a pattern attached.
func applyPropertyPattern(name resource.PropertyKey, prop resources.CloudAPIProperty,
	news resource.PropertyMap) (*resource.PropertyValue, bool) {
	if prop.Pattern == "" {
		return nil, false
	}

	// Find the value from the user and check that it's a simple values without a path pattern.
	userValue, ok := news[name]
	if !ok || !userValue.HasValue() || !userValue.IsString() || strings.Contains(userValue.StringValue(), "/") {
		return nil, false
	}

	// Build the pattern-based value from the template and parameters.
	result := prop.Pattern
	for key, value := range news {
		if !value.HasValue() || !value.IsString() {
			continue
		}
		part := fmt.Sprintf("{%s}", string(key))
		result = strings.Replace(result, part, value.StringValue(), 1)
	}

	if strings.Contains(result, "{") {
		// Not all properties resolved, so the result is invalid.
		return nil, false
	}

	newValue := resource.NewStringProperty(result)
	return &newValue, true
}
