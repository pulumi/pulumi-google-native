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
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
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

func TestCalculateObjectDiff(t *testing.T) {
	diff := &resource.ObjectDiff{
		Adds: resource.PropertyMap{
			"p3":            {V: "newkey"},
			"dots.addition": {V: "added"},
		},
		Deletes: resource.PropertyMap{
			"p2":            {V: "iamdeleted"},
			"dots.deletion": {V: "deleted"},
		},
		Updates: map[resource.PropertyKey]resource.ValueDiff{
			"p1": {
				Old: resource.PropertyValue{V: "oldvalue"},
				New: resource.PropertyValue{V: "newvalue"},
			},
			"dots.update": {
				Old: resource.PropertyValue{V: "oldvalue"},
				New: resource.PropertyValue{V: "newvalue"},
			},
			"nested": {
				Object: &resource.ObjectDiff{
					Updates: map[resource.PropertyKey]resource.ValueDiff{
						"deep": {
							Object: &resource.ObjectDiff{
								Updates: map[resource.PropertyKey]resource.ValueDiff{
									"nested.update/field": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
									"whitespaced field": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
									`"quotedfield"`: {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
									"regular": {
										Old: resource.PropertyValue{V: "oldvalue"},
										New: resource.PropertyValue{V: "newvalue"},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	d := differ{}
	actual := d.calculateObjectDiff(diff, "", "")

	assert.Equal(t, rpc.PropertyDiff_ADD, actual["p3"].Kind)
	assert.Equal(t, rpc.PropertyDiff_ADD, actual[`["dots.addition"]`].Kind)
	assert.Equal(t, rpc.PropertyDiff_DELETE, actual["p2"].Kind)
	assert.Equal(t, rpc.PropertyDiff_DELETE, actual[`["dots.deletion"]`].Kind)
	assert.Equal(t, rpc.PropertyDiff_UPDATE, actual["p1"].Kind)
	assert.Equal(t, rpc.PropertyDiff_UPDATE, actual[`["dots.update"]`].Kind)
	assert.Equal(t, rpc.PropertyDiff_UPDATE, actual[`nested.deep["nested.update/field"]`].Kind)
	assert.Equal(t, rpc.PropertyDiff_UPDATE, actual[`nested.deep["whitespaced field"]`].Kind)
	assert.Equal(t, rpc.PropertyDiff_UPDATE, actual[`nested.deep["\"quotedfield\""]`].Kind)
	assert.Equal(t, rpc.PropertyDiff_UPDATE, actual[`nested.deep.regular`].Kind)
}
