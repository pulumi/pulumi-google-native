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
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCalculateResourceId_IdProperty(t *testing.T) {
	res := resources.CloudAPIResource{
		RootURL:    "https://myapi.google.com",
		IDProperty: "selfLink",
	}
	inputs := map[string]interface{}{
		"name": "foo",
	}
	expected := "/v1/myresource/foo"
	outputs := map[string]interface{}{
		"name":     "foo",
		"selfLink": "https://myapi.google.com" + expected,
	}
	actual, err := calculateResourceID(res, inputs, outputs)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestCalculateResourceId_IdPath(t *testing.T) {
	res := resources.CloudAPIResource{
		RootURL: "https://myapi.google.com",
		IDPath:  "/v1/myparent/{parentsId}/myresource/{myresourcesId}",
		IDParams: map[string]string{
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
	actual, err := calculateResourceID(res, inputs, outputs)
	assert.NoError(t, err)
	assert.Equal(t, "/v1/myparent/foo/myresource/bar", actual)
}

func TestGetDefaultName_Generated(t *testing.T) {
	urn := resource.URN("urn:pulumi:dev::test::test-provider:testModule:TestResource::myName")
	olds := resource.NewPropertyMapFromMap(map[string]interface{}{})
	news := resource.NewPropertyMapFromMap(map[string]interface{}{
		"project":  "p01",
		"location": "west",
		"ignoreMe": 1.2,
	})
	actual := getDefaultName(urn, "projects/{project}/locations/{location}/things/{name}", olds, news)
	expectedPrefix := "projects/p01/locations/west/things/myName-"
	assert.True(t, strings.HasPrefix(actual.StringValue(), expectedPrefix))
	assert.Equal(t, len(expectedPrefix)+7, len(actual.StringValue()))
}

func TestGetDefaultName_OldApplied(t *testing.T) {
	urn := resource.URN("urn:pulumi:dev::test::test-provider:testModule:TestResource::myName")
	fixedName := "projects/p01/locations/west/things/anotherName"
	olds := resource.NewPropertyMapFromMap(map[string]interface{}{
		"name": fixedName,
	})
	news := resource.NewPropertyMapFromMap(map[string]interface{}{
		"project":  "p01",
		"location": "west",
		"ignoreMe": 1.2,
	})
	actual := getDefaultName(urn, "projects/{project}/locations/{location}/things/{name}", olds, news)
	assert.Equal(t, fixedName, actual.StringValue())
}

func TestApplyPropertyPattern_Applied(t *testing.T) {
	name := resource.PropertyKey("machineType")
	prop := resources.CloudAPIProperty{Pattern: "zones/{zone}/machineTypes/{machineType}"}
	news := resource.NewPropertyMapFromMap(map[string]interface{}{
		"zone":        "uswest-1a",
		"machineType": "f1-micro",
		"ignoreMe":    1.2,
	})
	actual, ok := applyPropertyPattern(name, prop, news)
	assert.True(t, ok)
	assert.Equal(t, "zones/uswest-1a/machineTypes/f1-micro", actual.StringValue())
}

func TestApplyPropertyPattern_UserValue(t *testing.T) {
	name := resource.PropertyKey("machineType")
	prop := resources.CloudAPIProperty{Pattern: "zones/{zone}/machineTypes/{machineType}"}
	news := resource.NewPropertyMapFromMap(map[string]interface{}{
		"zone":        "uswest-1a", // this is ignored
		"machineType": "zones/uswest-2a/machineTypes/f1-micro",
	})
	actual, ok := applyPropertyPattern(name, prop, news)
	assert.False(t, ok)
	assert.Nil(t, actual)
}

func TestApplyPropertyPattern_UnknownValue(t *testing.T) {
	name := resource.PropertyKey("machineType")
	prop := resources.CloudAPIProperty{Pattern: "zones/{zone}/machineTypes/{machineType}"}
	news := resource.NewPropertyMapFromMap(map[string]interface{}{
		"zone":        "uswest-1a",
		"machineType": resource.NewOutputProperty(resource.Output{Element: resource.NewStringProperty("")}),
	})
	actual, ok := applyPropertyPattern(name, prop, news)
	assert.False(t, ok)
	assert.Nil(t, actual)
}

func TestApplyPropertyPattern_UnknownParameter(t *testing.T) {
	name := resource.PropertyKey("machineType")
	prop := resources.CloudAPIProperty{Pattern: "zones/{zone}/machineTypes/{machineType}"}
	news := resource.NewPropertyMapFromMap(map[string]interface{}{
		"zone":        resource.NewOutputProperty(resource.Output{Element: resource.NewStringProperty("")}),
		"machineType": "f1-micro",
	})
	actual, ok := applyPropertyPattern(name, prop, news)
	assert.False(t, ok)
	assert.Nil(t, actual)
}
