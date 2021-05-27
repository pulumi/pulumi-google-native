// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateResourceId_IdProperty(t *testing.T) {
	res := resources.CloudAPIResource{
		BaseUrl:    "https://myapi.google.com",
		IdProperty: "selfLink",
	}
	inputs := map[string]interface{}{
		"name": "foo",
	}
	expected := "/v1/myresource/foo"
	outputs := map[string]interface{}{
		"name":     "foo",
		"selfLink": "https://myapi.google.com" + expected,
	}
	actual, err := calculateResourceId(res, inputs, outputs)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCalculateResourceId_IdPath(t *testing.T) {
	res := resources.CloudAPIResource{
		BaseUrl: "https://myapi.google.com",
		IdPath:  "/v1/myparent/{parentsId}/myresource/{myresourcesId}",
		IdParams: map[string]string{
			"parentsId":     "parentId",
			"myresourcesId": "reference.name",
		},
	}
	inputs := map[string]interface{}{
		"parentId": "foo",
	}
	outputs := map[string]interface{}{
		"reference": map[string]interface{}{
			"name": "myparent/foo/myresource/bar",
		},
	}
	actual, err := calculateResourceId(res, inputs, outputs)
	assert.NoError(t, err)
	assert.Equal(t, "/v1/myparent/foo/myresource/bar", actual)
}

func TestEvalPropertyValue(t *testing.T) {
	values := map[string]interface{}{
		"a": "123",
		"b": map[string]interface{}{
			"c": map[string]interface{}{
				"d": "456",
			},
		},
		"notastring": 789,
	}

	a, has := evalPropertyValue(values, "a")
	assert.True(t, has)
	assert.Equal(t, "123", a)

	_, has = evalPropertyValue(values, "unknown")
	assert.False(t, has)

	d, has := evalPropertyValue(values, "b.c.d")
	assert.True(t, has)
	assert.Equal(t, "456", d)

	_, has = evalPropertyValue(values, "b.unknown.d")
	assert.False(t, has)

	_, has = evalPropertyValue(values, "notastring")
	assert.False(t, has)
}
