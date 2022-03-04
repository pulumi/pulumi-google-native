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
)

// getInputsFromState calculates the inputs map given resource state.
func getInputsFromState(res resources.CloudAPIResource, state resource.PropertyMap) resource.PropertyMap {
	inputsMap := map[string]interface{}{}
	stateMap := state.Mappable()
	for name, prop := range res.Create.SDKProperties {
		parent := stateMap
		if prop.Container != "" {
			container, ok := stateMap[prop.Container].(map[string]interface{})
			if !ok {
				continue
			}
			parent = container
		}

		if value, has := parent[name]; has {
			sdkName := name
			if prop.SdkName != "" {
				sdkName = prop.SdkName
			}
			inputsMap[sdkName] = value
		}
	}
	return resource.NewPropertyMapFromMap(inputsMap)
}

// applyDiff produces a new map as a merge of a calculated diff into an existing map of values.
func applyDiff(values resource.PropertyMap, diff *resource.ObjectDiff) resource.PropertyMap {
	if diff == nil {
		return values
	}

	result := resource.PropertyMap{}
	for name, value := range values {
		if !diff.Deleted(name) {
			result[name] = value
		}
	}
	for key, value := range diff.Adds {
		result[key] = value
	}
	for key, value := range diff.Updates {
		result[key] = value.New
	}
	return result
}
