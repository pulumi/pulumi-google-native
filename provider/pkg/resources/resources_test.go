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

func TestCalculateResourceId_IdProperty(t *testing.T) {
	endpoint := CloudAPIEndpoint{
		SelfLinkProperty: "selfLink",
	}
	inputs := map[string]interface{}{
		"name": "foo",
	}
	expected := "https://myapi.google.com/v1/myresource/foo"
	outputs := map[string]interface{}{
		"name":     "foo",
		"selfLink": "https://myapi.google.com/v1/myresource/foo",
	}
	actual, err := endpoint.URI(inputs, outputs)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestEndpoint_URI(t *testing.T) {
	endpoint := CloudAPIEndpoint{
		Template: "/v1/myparent/{parentsId}/myresource/{myresourcesId}",
		Values: []CloudAPIResourceParam{
			{
				Name:    "parentsId",
				SdkName: "parentId",
			},
			{
				Name:    "myresourcesId",
				SdkName: "reference.name",
			},
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
	actual, err := endpoint.URI(inputs, outputs)
	assert.NoError(t, err)
	assert.Equal(t, "/v1/myparent/foo/myresource/bar", actual)
}

func TestEndpoint_URI_RFC_6570(t *testing.T) {
	endpoint := CloudAPIEndpoint{
		Template: "/v1/myparent/{+name}",
		Values: []CloudAPIResourceParam{
			{
				Name:    "name",
				SdkName: "name",
			},
		},
	}
	inputs := map[string]interface{}{
		"parentId": "foo",
	}
	outputs := map[string]interface{}{
		"name": "foo/myresource/bar",
	}
	actual, err := endpoint.URI(inputs, outputs)
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

	a, has := EvalPropertyValue(values, "a")
	assert.True(t, has)
	assert.Equal(t, "123", a)

	_, has = EvalPropertyValue(values, "unknown")
	assert.False(t, has)

	d, has := EvalPropertyValue(values, "b.c.d")
	assert.True(t, has)
	assert.Equal(t, "456", d)

	_, has = EvalPropertyValue(values, "b.unknown.d")
	assert.False(t, has)

	_, has = EvalPropertyValue(values, "notastring")
	assert.False(t, has)
}
