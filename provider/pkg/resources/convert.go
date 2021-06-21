// Copyright 2016-2021, Pulumi Corporation.

package resources

import (
	"reflect"
	"strings"
)

const body = "body"

// SdkShapeConverter providers functions to convert between HTTP request/response shapes and
// Pulumi SDK shapes (with flattening, renaming, etc.).
type SdkShapeConverter struct {
	Types map[string]CloudAPIType
}

type convertPropValues func(props map[string]CloudAPIProperty, values map[string]interface{}) map[string]interface{}

func (k *SdkShapeConverter) convertPropValue(prop *CloudAPIProperty, value interface{}, convertMap convertPropValues) interface{} {
	if value == nil {
		return nil
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.Map:
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return value
		}

		if strings.HasPrefix(prop.Ref, "#/types/") {
			typeName := strings.TrimPrefix(prop.Ref, "#/types/")
			typ, ok := k.Types[typeName]
			if !ok {
				return value
			}
			return convertMap(typ.Properties, valueMap)
		}

		if prop.AdditionalProperties != nil {
			result := map[string]interface{}{}
			for key, item := range valueMap {
				result[key] = k.convertPropValue(prop.AdditionalProperties, item, convertMap)
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
		for i := 0; i < s.Len(); i++ {
			result = append(result, k.convertPropValue(prop.Items, s.Index(i).Interface(), convertMap))
		}
		return result
	}
	return value
}

// SdkPropertiesToRequestBody converts a map of SDK properties to JSON request body to be sent to an HTTP API.
func (k *SdkShapeConverter) SdkPropertiesToRequestBody(props map[string]CloudAPIProperty,
	values map[string]interface{}) map[string]interface{} {

	body := map[string]interface{}{}
	for name, prop := range props {
		p := prop // https://github.com/golang/go/wiki/CommonMistakes#using-reference-to-loop-iterator-variable
		parent := body
		sdkName := name
		if prop.SdkName != "" {
			sdkName = prop.SdkName
		}
		if value, has := values[sdkName]; has {
			if prop.Container != "" {
				if v, has := body[prop.Container].(map[string]interface{}); has {
					parent = v
				} else {
					parent = map[string]interface{}{}
					body[prop.Container] = parent
				}
			}
			parent[name] = k.convertPropValue(&p, value, k.SdkPropertiesToRequestBody)
		}
	}
	return body
}
