// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInputsFromState(t *testing.T) {
	res := resources.CloudAPIResource{
		CreateProperties: map[string]resources.CloudAPIProperty{
			"p1": {},
			"p2": {
				Container: "c1",
			},
			"p3": {},
		},
	}
	state := resource.NewPropertyMapFromMap(map[string]interface{}{
		"p1": "v1",
		"c1": map[string]interface{}{
			"p2": "v2",
		},
		"p3": 123.456,
		"output1": "should-be-ignored",
	})
	actual := getInputsFromState(res, state).Mappable()
	expected := map[string]interface{}{
		"p1": "v1",
		"p2": "v2",
		"p3": 123.456,
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
