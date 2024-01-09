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

package provider

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func calculateResourceID(
	res resources.CloudAPIResource,
	inputs map[string]interface{},
	outputs map[string]interface{},
) (string, error) {
	if res.IDProperty != "" {
		v, ok := outputs[res.IDProperty].(string)
		if !ok {
			return "", errors.Errorf("ID property %q not found", res.IDProperty)
		}
		return strings.TrimPrefix(v, res.RootURL), nil
	}

	id := res.IDPath
	idParams := res.IDParams
	for name, alias := range idParams {
		var propValue string
		if v, has := resources.EvalPropertyValue(inputs, alias); has {
			propValue = v
		} else if v, has := resources.EvalPropertyValue(outputs, alias); has {
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

// buildCreateURL composes the URL to invoke to create a resource with given inputs.
func buildCreateURL(res resources.CloudAPIResource, inputs resource.PropertyMap) (string, error) {
	return buildURL(res.Create.Endpoint, inputs, nil)
}

func buildFunctionURL(res resources.CloudAPIFunction, inputs resource.PropertyMap) (string, error) {
	return buildURL(res.URL, inputs, nil)
}

func buildURL(
	endpoint resources.CloudAPIEndpoint,
	inputs resource.PropertyMap,
	extraQueryArgs map[string]string,
) (string, error) {
	queryMap := map[string]string{}
	uriString := endpoint.Template
	for _, param := range endpoint.Values {
		sdkName := param.Name
		if param.SdkName != "" {
			sdkName = param.SdkName
		}
		key := resource.PropertyKey(sdkName)
		inputString, found, err := getInputString(inputs, key)
		if !found {
			continue
		}
		if err != nil {
			return "", errors.Wrapf(err, "getting input string for property %q", key)
		}

		switch param.Kind {
		case "path":
			uriString = strings.Replace(uriString, fmt.Sprintf("{%s}", param.Name), url.PathEscape(inputString), 1)
		case "query":
			queryMap[param.Name] = inputString
		default:
			return "", errors.Errorf("unknown param location %q", param.Kind)
		}
	}
	uri, err := url.Parse(uriString)
	if err != nil {
		return "", errors.Wrapf(err, "parsing resource URL %q", uriString)
	}
	uri.Scheme = "https"
	query := uri.Query()
	for key, value := range queryMap {
		query.Set(key, value)
	}
	for key, value := range extraQueryArgs {
		query.Set(key, value)
	}
	uri.RawQuery = query.Encode()
	return uri.String(), nil
}

func getInputString(inputs resource.PropertyMap, key resource.PropertyKey) (string, bool, error) {
	val := inputs[key]
	if !val.HasValue() {
		return "", false, nil
	}
	if val.IsString() {
		return val.StringValue(), true, nil
	}
	if val.IsNumber() {
		return fmt.Sprintf("%v", val.NumberValue()), true, nil
	}
	if val.IsBool() {
		return fmt.Sprintf("%v", val.BoolValue()), true, nil
	}

	return "", true, fmt.Errorf("unexpected input type for property key %q - expected string, number or boolean", key)
}

var autonameFieldRegex = regexp.MustCompile(`^\w+$`) // Only word characters allowed

// getDefaultName generates a random name for a resource based on its URN name, a given pattern,
// and other properties if necessary. If an existing name is derived from previous state,
// the existing name is returned. If getDefaultName generates a new random name,
// then the second return value is true, false otherwise.
func getDefaultName(
	randomSeed []byte,
	urn resource.URN,
	nameKey resource.PropertyKey,
	pattern string,
	olds, news resource.PropertyMap,
) (resource.PropertyValue, bool) {
	previouslyAutonamed := olds.HasValue(autonamed)
	if v, ok := olds[nameKey]; ok && previouslyAutonamed {
		return v, true
	}

	name := urn.Name()

	// Resource name is URN name + "-" + random suffix.
	random, err := resource.NewUniqueName(randomSeed, name+"-", 7, 0, nil)
	contract.AssertNoError(err)

	// Simple field replacement, so just return the autoname.
	if autonameFieldRegex.MatchString(pattern) {
		return resource.NewStringProperty(random), true
	}

	result := strings.Replace(pattern, "{name}", random, 1)
	for key, value := range news {
		if !value.HasValue() || !value.IsString() {
			continue
		}

		part := fmt.Sprintf("{%s}", string(key))
		result = strings.Replace(result, part, value.StringValue(), 1)
	}

	return resource.NewStringProperty(result), true
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
