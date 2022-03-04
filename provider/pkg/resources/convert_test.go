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
					SdkName:   "v4",
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
				"fingerprint": {
					CopyFromOutputs: true,
				},
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
			Create: CloudAPIOperation{
				SDKProperties: map[string]CloudAPIProperty{
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
						Ref:       "#/types/google-native:testing:More",
					},
					"tags":         {},
					"untypedArray": {},
					"untypedDict": {
						Ref: "pulumi.json#/Any",
					},
				},
			},
			Update: CloudAPIOperation{
				SDKProperties: map[string]CloudAPIProperty{
					"more": {
						Container: "properties",
						Ref:       "#/types/google-native:testing:More",
					},
					"tags": {},
					"tagsFingerprint": {
						CopyFromOutputs: true,
					},
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
var sampleSdkState = map[string]interface{}{
	"more": map[string]interface{}{
		"items": []interface{}{
			map[string]interface{}{"Aaa": "111", "bbb": "333"},
			map[string]interface{}{"Aaa": "222", "fingerprint": "fitems[1]"},
		},
		"itemsMap": map[string]interface{}{
			"key1": map[string]interface{}{"fingerprint": "fkey1"},
			"key2": map[string]interface{}{"Aaa": "666"},
		},
	},
	"name":            "MyResource",
	"tagsFingerprint": "ABCDEF123456",
}
var sampleAPIUpdatePackage = map[string]interface{}{
	"properties": map[string]interface{}{
		"more": map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"aaa": "111", "ccc": map[string]interface{}{"bbb": "333"}},
				map[string]interface{}{"aaa": "222", "fingerprint": "fitems[1]"},
			},
			"itemsMap": map[string]interface{}{
				"key1": map[string]interface{}{"aaa": "444", "ccc": map[string]interface{}{"bbb": "555"}, "fingerprint": "fkey1"},
				"key2": map[string]interface{}{"aaa": "666"},
			},
		},
	},
	"tags": map[string]interface{}{
		"createdBy":   "admin",
		"application": "dashboard",
	},
	"tagsFingerprint": "ABCDEF123456",
}

func TestSdkPropertiesToRequestBody(t *testing.T) {
	bodyProperties := resourceMap.Resources["r1"].Create.SDKProperties
	data := c.SdkPropertiesToRequestBody(bodyProperties, sampleSdkProps, nil)
	assert.Equal(t, sampleAPIPackage, data)
}

func TestSdkPropertiesToRequestBodyWithState(t *testing.T) {
	// State has fewer elements in collections and different values to test
	// that converter is prepared for this kind of changes.
	state := map[string]interface{}{
		"name":      "MyResource",
		"threshold": 456,
		"structure": map[string]interface{}{
			"v1": "value1old",
			"v4": false,
		},
		"p1": "prop1old",
		"p2": "prop2old",
		"p3": "prop3old",
		"more": map[string]interface{}{
			"items": []interface{}{
				map[string]interface{}{"bbb": "333"},
			},
			"itemsMap": map[string]interface{}{
				"key1": map[string]interface{}{"Aaa": "444"},
			},
		},
		"tags": map[string]interface{}{
			"application": "another",
		},
		"untypedArray": []interface{}{
			map[string]interface{}{"key1": "value2"},
		},
		"untypedDict": map[string]interface{}{
			"key1": "value1",
		},
	}
	bodyProperties := resourceMap.Resources["r1"].Create.SDKProperties
	data := c.SdkPropertiesToRequestBody(bodyProperties, sampleSdkProps, state)
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
	bodyProperties := resourceMap.Resources["r1"].Create.SDKProperties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, emptyCollectionData, nil)
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
	bodyProperties := resourceMap.Resources["r1"].Create.SDKProperties
	actualBody := c.SdkPropertiesToRequestBody(bodyProperties, nilCollectionData, nil)
	assert.Equal(t, expectedBody, actualBody)
}

func TestSdkPropertiesToRequestBodyCopyOutputs(t *testing.T) {
	bodyProperties := resourceMap.Resources["r1"].Update.SDKProperties
	data := c.SdkPropertiesToRequestBody(bodyProperties, sampleSdkProps, sampleSdkState)
	assert.Equal(t, sampleAPIUpdatePackage, data)
}
