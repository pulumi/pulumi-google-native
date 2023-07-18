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
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
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

// calculateDetailedDiff produced a property diff for a given object diff and a resource definition. It inspects
// the schema of the resource to find out if the requested diff can be performed in-place or requires a replacement.
func calculateDetailedDiff(resource *resources.CloudAPIResource, types map[string]resources.CloudAPIType,
	diff *resource.ObjectDiff) map[string]*rpc.PropertyDiff {
	replaceKeys := codegen.NewStringSet()

	populateReplaceKeys := func(params []resources.CloudAPIResourceParam,
		properties map[string]resources.CloudAPIProperty) {
		for _, p := range params {
			// All path parameters that are part of the resource create URI automatically cause a replacement.
			if p.Kind == "path" {
				name := p.Name
				if p.SdkName != "" {
					name = p.SdkName
				}
				replaceKeys.Add(name)
			}
		}

		// Search for forceNew on top-level resource properties
		findForceNew("", properties, replaceKeys)
		for propName, prop := range properties {
			// Object types.
			if prop.Ref != "" {
				typName := strings.TrimPrefix(prop.Ref, "#/types/")
				if typ, has := types[typName]; has {
					findForceNew(propName+".", typ.Properties, replaceKeys)
				}
			}
			// Arrays of objects.
			if prop.Items != nil && prop.Items.Ref != "" {
				typName := strings.TrimPrefix(prop.Items.Ref, "#/types/")
				if typ, has := types[typName]; has {
					findForceNew(propName+"[].", typ.Properties, replaceKeys)
				}
			}
		}
	}

	populateReplaceKeys(resource.Create.Endpoint.Values, resource.Create.SDKProperties)
	populateReplaceKeys(resource.Update.Endpoint.Values, resource.Update.SDKProperties)

	d := differ{replaceKeys: replaceKeys}
	return d.calculateObjectDiff(diff, "", "")
}

func findForceNew(base string, props map[string]resources.CloudAPIProperty, replaceKeys codegen.StringSet) {
	for propName, prop := range props {
		if prop.ForceNew {
			name := propName
			if prop.SdkName != "" {
				name = prop.SdkName
			}
			replaceKeys.Add(base + name)
		}
	}
}

type differ struct {
	replaceKeys codegen.StringSet
}

func (d *differ) calculateObjectDiff(diff *resource.ObjectDiff, diffBase, replaceBase string) map[string]*rpc.PropertyDiff {
	detailedDiff := map[string]*rpc.PropertyDiff{}

	// asResourcePath converts a given base and a property key to a wellformed resource path key.
	asResourcePath := func(base string, p resource.PropertyKey) string {
		s := string(p)

		// If the field name contains characters that needs escaping then the property key should be presented
		// in square brackets and quotes.
		// Example:
		// If base is `config.labels.` and the field name is `app.kubernetes.io/name` then the resource key will
		// be `config.labels["app.kubernetes.io/name"]`
		if strings.ContainsAny(s, `". `) {
			// Remove the trailing dot for the base
			base = strings.TrimSuffix(base, ".")

			// Escape double quotes
			s = strings.ReplaceAll(s, `"`, `\"`)

			s = fmt.Sprintf(`["%s"]`, s)
		}

		return base + s
	}

	for k, v := range diff.Updates {
		key := asResourcePath(diffBase, k)
		replaceKey := asResourcePath(replaceBase, k)
		subDiff := d.calculateValueDiff(&v, key, replaceKey)
		for sk, sv := range subDiff {
			detailedDiff[sk] = sv
		}
	}
	for k := range diff.Adds {
		diffKey := asResourcePath(diffBase, k)
		replaceKey := asResourcePath(replaceBase, k)
		kind := rpc.PropertyDiff_ADD
		if d.replaceKeys.Has(replaceKey) {
			kind = rpc.PropertyDiff_ADD_REPLACE
		}
		detailedDiff[diffKey] = &rpc.PropertyDiff{Kind: kind}
	}
	for k := range diff.Deletes {
		diffKey := asResourcePath(diffBase, k)
		replaceKey := asResourcePath(replaceBase, k)
		kind := rpc.PropertyDiff_DELETE
		if d.replaceKeys.Has(replaceKey) {
			kind = rpc.PropertyDiff_DELETE_REPLACE
		}
		detailedDiff[diffKey] = &rpc.PropertyDiff{Kind: kind}
	}

	return detailedDiff
}

func (d *differ) calculateValueDiff(v *resource.ValueDiff, diffBase, replaceBase string) map[string]*rpc.PropertyDiff {
	detailedDiff := map[string]*rpc.PropertyDiff{}
	switch {
	case v.Object != nil:
		subDiff := d.calculateObjectDiff(v.Object, diffBase+".", replaceBase+".")
		for sk, sv := range subDiff {
			if sv.Kind == rpc.PropertyDiff_UPDATE && d.replaceKeys.Has(replaceBase) {
				// If the parent property causes a replacement, all child properties cause a replacement.
				sv.Kind = rpc.PropertyDiff_UPDATE_REPLACE
			}
			detailedDiff[sk] = sv
		}
	case v.Array != nil:
		for idx, item := range v.Array.Updates {
			key := fmt.Sprintf("%s[%d]", diffBase, idx)
			subDiff := d.calculateValueDiff(&item, key, replaceBase+"[]")
			for sk, sv := range subDiff {
				detailedDiff[sk] = sv
			}
		}
		for idx := range v.Array.Adds {
			key := fmt.Sprintf("%s[%d]", diffBase, idx)
			detailedDiff[key] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_ADD}
		}
		for idx := range v.Array.Deletes {
			key := fmt.Sprintf("%s[%d]", diffBase, idx)
			detailedDiff[key] = &rpc.PropertyDiff{Kind: rpc.PropertyDiff_DELETE}
		}
	default:
		kind := rpc.PropertyDiff_UPDATE
		if d.replaceKeys.Has(replaceBase) {
			kind = rpc.PropertyDiff_UPDATE_REPLACE
		}
		detailedDiff[diffBase] = &rpc.PropertyDiff{Kind: kind}
	}
	return detailedDiff
}

func diffWithFlattenedKeys(olds, news resource.PropertyMap, detailedDiff map[string]*rpc.PropertyDiff) (map[string]rpc.PropertyDiff_Kind,
	error) {
	expandedKeys := map[string]rpc.PropertyDiff_Kind{}
	for k, diff := range detailedDiff {
		propPath, err := resource.ParsePropertyPath(k)
		if err != nil {
			return nil, err
		}
		switch diff.Kind {
		case rpc.PropertyDiff_ADD, rpc.PropertyDiff_ADD_REPLACE:
			val, ok := propPath.Get(resource.NewObjectProperty(news))
			if !ok {
				return nil, fmt.Errorf("failed to lookup value at path %q on new inputs", k)
			}
			addNestedKeys(val, k, diff.Kind, expandedKeys)
		case rpc.PropertyDiff_DELETE, rpc.PropertyDiff_DELETE_REPLACE:
			val, ok := propPath.Get(resource.NewObjectProperty(olds))
			if !ok {
				return nil, fmt.Errorf("failed to lookup value at path %q on old inputs", k)
			}
			addNestedKeys(val, k, diff.Kind, expandedKeys)
		case rpc.PropertyDiff_UPDATE, rpc.PropertyDiff_UPDATE_REPLACE:
			expandedKeys[k] = diff.Kind
		}
	}
	return expandedKeys, nil
}

func addNestedKeys(val resource.PropertyValue, currPath string, kind rpc.PropertyDiff_Kind,
	expandedKeys map[string]rpc.PropertyDiff_Kind) {
	if val.IsNull() {
		return
	} else if val.IsBool() || val.IsNumber() || val.IsString() || val.IsAsset() || val.IsArchive() {
		expandedKeys[currPath] = kind
		return
	} else if val.IsArray() {
		for i, e := range val.ArrayValue() {
			addNestedKeys(e, fmt.Sprintf("%s[%d]", currPath, i), kind, expandedKeys)
		}
		return
	} else if val.IsObject() {
		propMap := val.ObjectValue()
		for k, v := range propMap {
			addNestedKeys(v, fmt.Sprintf("%s.%s", currPath, k), kind, expandedKeys)
		}
		return
	}
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
