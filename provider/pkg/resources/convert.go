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

package resources

import (
	"reflect"
	"strings"
)

// SdkShapeConverter providers functions to convert between HTTP request/response shapes and
// Pulumi SDK shapes (with flattening, renaming, etc.).
type SdkShapeConverter struct {
	Types map[string]CloudAPIType
}

type convertPropValues func(
	props map[string]CloudAPIProperty,
	inputs, state map[string]interface{},
) map[string]interface{}

func (k *SdkShapeConverter) convertPropValue(
	prop *CloudAPIProperty,
	value, state interface{},
	convertMap convertPropValues,
) interface{} {
	if value == nil {
		return nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}
		stateMap, ok := state.(map[string]interface{})
		if !ok {
			stateMap = map[string]interface{}{}
		}

		if strings.HasPrefix(prop.Ref, "#/types/") {
			typeName := strings.TrimPrefix(prop.Ref, "#/types/")
			typ, ok := k.Types[typeName]
			if !ok {
				return value
			}
			return convertMap(typ.Properties, valueMap, stateMap)
		}

		if prop.AdditionalProperties != nil {
			result := map[string]interface{}{}
			for key, item := range valueMap {
				result[key] = k.convertPropValue(prop.AdditionalProperties, item, stateMap[key], convertMap)
			}
			return result
		}

		return value
	case reflect.Slice, reflect.Array:
		if prop.Items == nil {
			return value
		}

		result := make([]interface{}, 0)
		s := reflect.ValueOf(value)

		stateSlice, ok := state.([]interface{})
		if !ok {
			stateSlice = make([]interface{}, s.Len())
		}

		for i := 0; i < s.Len(); i++ {
			var stateItem interface{}
			if len(stateSlice) > i {
				stateItem = stateSlice[i]
			}
			result = append(result, k.convertPropValue(prop.Items, s.Index(i).Interface(), stateItem, convertMap))
		}
		return result
	}
	return value
}

// SdkPropertiesToRequestBody converts a map of SDK properties to JSON request body to be sent to an HTTP API.
func (k *SdkShapeConverter) SdkPropertiesToRequestBody(props map[string]CloudAPIProperty,
	inputs map[string]interface{}, state map[string]interface{}) map[string]interface{} {

	body := map[string]interface{}{}
	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		parent := body
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}

		values := inputs
		if prop.CopyFromOutputs && state != nil {
			values = state
		}

		if value, has := values[sdkName]; has {
			if prop.Flatten {
				child := k.convertPropValue(&p, value, state[sdkName], k.SdkPropertiesToRequestBody)
				asMap := child.(map[string]interface{})
				for k, v := range asMap {
					parent[k] = v
				}
				continue
			}
			if prop.Container != "" {
				if v, has := body[prop.Container].(map[string]interface{}); has {
					parent = v
				} else {
					parent = map[string]interface{}{}
					body[prop.Container] = parent
				}
			}
			parent[name] = k.convertPropValue(&p, value, state[sdkName], k.SdkPropertiesToRequestBody)
		}
	}
	return body
}
