// Copyright 2016-2021, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var resourceMap = &CloudAPIMetadata{
	Types: map[string]CloudAPIType{
		"google-native:testing:Structure": {
			Properties: map[string]CloudAPIProperty{
				"v1": {},
				"v2": {},
				"v3-odd": {
					SdkName: "v3",
				},
				"v4-nested": {
					SdkName:    "v4",
					Container: "props",
				},
			},
		},
		"google-native:testing:More": {
			Properties: map[string]CloudAPIProperty{
				"items": {
					Items: &CloudAPIProperty{
						Ref: "#/types/google-native:testing:MoreItem",
					},
				},
				"itemsMap": {
					AdditionalProperties: &CloudAPIProperty{
						Ref: "#/types/google-native:testing:MoreItem",
					},
				},
			},
		},
		"google-native:testing:MoreItem": {
			Properties: map[string]CloudAPIProperty{
				"aaa": {
					SdkName: "Aaa",
				},
				"bbb": {
					Container: "ccc",
				},
			},
		},
	},
	Resources: map[string]CloudAPIResource{
		"r1": {
			CreateProperties: map[string]CloudAPIProperty{
				"name": {},
				"x-threshold": {
					SdkName: "threshold",
				},
				"structure": {
					Ref: "#/types/google-native:testing:Structure",
				},
				"p1": {
					Container: "properties",
				},
				"p2": {
					Container: "properties",
				},
				"more": {
					Container: "properties",
					Ref:        "#/types/google-native:testing:More",
				},
				"tags":         {},
				"untypedArray": {},
				"untypedDict": {
					Ref: "pulumi.json#/Any",
				},
			},
		},
	},
}

var c = SdkShapeConverter{Types: resourceMap.Types}

var sampleAPIPackage = map[string]interface{}{
	"name":        "MyResource",
	"x-threshold": 123,
	"structure": map[string]interface{}{
		"v1":     "value1",
		"v2":     2,
		"v3-odd": "odd-value",
		"props": map[string]interface{}{
			"v4-nested": true,
		},
	},
	"properties": map[string]interface{}{
		"p1": "prop1",
		"p2": "prop2",
		"more": map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"aaa": "111", "ccc": map[string]interface{}{"bbb": "333"}},
				map[string]interface{}{"aaa": "222"},
			},
			"itemsMap": map[string]interface{}{
				"key1": map[string]interface{}{"aaa": "444", "ccc": map[string]interface{}{"bbb": "555"}},
				"key2": map[string]interface{}{"aaa": "666"},
			},
		},
	},
	"tags": map[string]interface{}{
		"createdBy":   "admin",
		"application": "dashboard",
	},
	"untypedArray": []interface{}{
		map[string]interface{}{"key1": "value1"},
		map[string]interface{}{"key1": "value2"},
	},
	"untypedDict": map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	},
}
var sampleSdkProps = map[string]interface{}{
	"name":      "MyResource",
	"threshold": 123,
	"structure": map[string]interface{}{
		"v1": "value1",
		"v2": 2,
		"v3": "odd-value",
		"v4": true,
	},
	"p1": "prop1",
	"p2": "prop2",
	"p3": "prop3",
	"more": map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{"Aaa": "111", "bbb": "333"},
			map[string]interface{}{"Aaa": "222"},
		},
		"itemsMap": map[string]interface{}{
			"key1": map[string]interface{}{"Aaa": "444", "bbb": "555"},
			"key2": map[string]interface{}{"Aaa": "666"},
		},
	},
	"tags": map[string]interface{}{
		"createdBy":   "admin",
		"application": "dashboard",
	},
	"untypedArray": []interface{}{
		map[string]interface{}{"key1": "value1"},
		map[string]interface{}{"key1": "value2"},
	},
	"untypedDict": map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	},
}

func TestSdkPropertiesToRequestBody(t *testing.T) {
	bodyProperties := resourceMap.Resources["r1"].CreateProperties
	data := c.SdkPropertiesToRequestBody(bodyProperties, sampleSdkProps)
	assert.Equal(t, sampleAPIPackage, data)
}

func TestSdkPropertiesToRequestBodyEmptyCollections(t *testing.T) {
	var emptyCollectionData = map[string]interface{}{
		"more": map[string]interface{}{
			"items":    make([]interface{}, 0),
			"itemsMap": make(map[string]interface{}),
		},
	}
	var expectedBody = map[string]interface{}{
		"properties": map[string]interface{}{
			"more": map[string]interface{}{
				"items":    make([]interface{}, 0),
				"itemsMap": make(map[string]interface{}),
			},
		},
	}
	bodyProperties := resourceMap.Resources["r1"].CreateProperties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, emptyCollectionData)
	assert.Equal(t, expectedBody, actualBody)
}

func TestSdkPropertiesToRequestBodyNilCollections(t *testing.T) {
	var nilCollectionData = map[string]interface{}{
		"more": map[string]interface{}{
			"items":    nil,
			"itemsMap": nil,
		},
	}
	var expectedBody = map[string]interface{}{
		"properties": map[string]interface{}{
			"more": map[string]interface{}{
				"items":    nil,
				"itemsMap": nil,
			},
		},
	}
	bodyProperties := resourceMap.Resources["r1"].CreateProperties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, nilCollectionData)
	assert.Equal(t, expectedBody, actualBody)
}
