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
	"testing"

	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestGetInputsFromState(t *testing.T) {
	res := resources.CloudAPIResource{
		Create: resources.CreateAPIOperation{
			CloudAPIOperation: resources.CloudAPIOperation{
				SDKProperties: map[string]resources.CloudAPIProperty{
					"p1": {},
					"p2": {
						Container: "c1",
					},
					"p3": {
						SdkName: "p3sdk",
					},
				},
			},
		},
	}
	state := resource.NewPropertyMapFromMap(map[string]interface{}{
		"p1": "v1",
		"c1": map[string]interface{}{
			"p2": "v2",
		},
		"p3":      123.456,
		"output1": "should-be-ignored",
	})
	actual := getInputsFromState(res, state).Mappable()
	expected := map[string]interface{}{
		"p1":    "v1",
		"p2":    "v2",
		"p3sdk": 123.456,
	}
	assert.Equal(t, expected, actual)
}

func TestApplyDiff(t *testing.T) {
	state := resource.PropertyMap{
		"p1": {V: "oldvalue"},
		"p2": {V: "iamdeleted"},
	}
	diff := resource.ObjectDiff{
		Adds: resource.PropertyMap{
			"p3": {V: "newkey"},
		},
		Deletes: resource.PropertyMap{
			"p2": {V: "iamdeleted"},
		},
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"p1": {
				Old: resource.PropertyValue{V: "oldvalue"},
				New: resource.PropertyValue{V: "newvalue"},
			},
		},
	}
	actual := applyDiff(state, &diff)
	expected := resource.PropertyMap{
		"p1": {V: "newvalue"},
		"p3": {V: "newkey"},
	}
	assert.Equal(t, expected, actual)
}
