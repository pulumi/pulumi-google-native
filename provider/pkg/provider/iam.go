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
	"regexp"
	"sort"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pulumi/pulumi-google-native/provider/pkg/googleclient"
	"github.com/pulumi/pulumi-google-native/provider/pkg/resources"
	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

type iamPolicyArgs struct {
	Bindings []*iamPolicyBinding `pulumi:"bindings,omitempty"`
}
type iamPolicyBinding struct {
	Condition *iamPolicyCondition `pulumi:"condition,omitempty"`
	Members   []string            `pulumi:"members,omitempty"`
	Role      string              `pulumi:"role,omitempty"`
}
type iamPolicyCondition struct {
	Description string `pulumi:"description,omitempty"`
	Expression  string `pulumi:"expression,omitempty"`
	Location    string `pulumi:"location,omitempty"`
	Title       string `pulumi:"title,omitempty"`
}

var iamOverlayRegex = regexp.MustCompile(`Iam(Member|Binding)$`)

func isIAMOverlay(urn resource.URN) bool {
	return iamOverlayRegex.MatchString(urn.Type().String())
}

func isIAMMember(urn resource.URN) bool {
	return strings.HasSuffix(urn.Type().String(), "IamMember")
}

func isIAMBinding(urn resource.URN) bool {
	return strings.HasSuffix(urn.Type().String(), "IamBinding")
}

// normalizeIamOverlayInputs transforms the IAM overlay inputs into the same shape as the underlying IAMPolicy resource.
func normalizeIamOverlayInputs(
	metadata resources.CloudAPIResource,
	urn resource.URN,
	inputs resource.PropertyMap,
) (*iamPolicyBinding, resource.PropertyMap, error) {
	contract.Assertf(metadata.IamResourceName != "", "missing metadata.IamResourceName for %s", urn)

	values := inputs.Copy()

	// Transform "name" -> API resource name
	if values["name"].HasValue() {
		values[resource.PropertyKey(metadata.IamResourceName)] = values["name"]
	}
	delete(values, "name")

	// Transform member (string) -> members ([]string)
	if isIAMMember(urn) {
		values[resource.PropertyKey("members")] = resource.NewArrayProperty(
			[]resource.PropertyValue{
				values[resource.PropertyKey("member")],
			})
		delete(values, "member")
	}

	inputsMap := values.Mappable()
	var newBinding iamPolicyBinding
	err := mapstructure.Decode(inputsMap, &newBinding)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding input IamMember: %w", err)
	}

	// These properties are captured under "bindings"
	delete(values, "condition")
	delete(values, "members")
	delete(values, "role")

	// Always use version 3 to support conditions
	values[resource.PropertyKey("version")] = resource.NewNumberProperty(3)

	return &newBinding, values, nil
}

// iamPolicyVersionArg is the query parameter required to use policy conditions (version 3)
var iamPolicyVersionArg = map[string]string{"optionsRequestedPolicyVersion": "3"}

func getIAMState(
	metadata resources.CloudAPIResource,
	inputs resource.PropertyMap,
	client *googleclient.GoogleClient,
) (*iamPolicyArgs, error) {
	uri, err := buildURL(metadata.Read.Endpoint, inputs, iamPolicyVersionArg)
	if err != nil {
		return nil, err
	}

	stateMap, err := retryRequest(client, metadata.Read.Verb, uri, "", nil)
	if err != nil {
		if strings.Contains(err.Error(), `Unknown name "optionsRequestedPolicyVersion"`) {
			uri, err = buildURL(metadata.Read.Endpoint, inputs, nil)
			if err != nil {
				return nil, err
			}
			stateMap, err = retryRequest(client, metadata.Read.Verb, uri, "", nil)
			if err != nil {
				return nil, fmt.Errorf("error fetching existing IAM Policy: %w", err)
			}
		} else {
			return nil, fmt.Errorf("error fetching existing IAM Policy: %w", err)
		}
	}
	var stateStruct iamPolicyArgs
	err = mapstructure.Decode(stateMap, &stateStruct)
	if err != nil {
		return nil, fmt.Errorf("error decoding existing IAM Policy: %w", err)
	}

	return &stateStruct, nil
}

func inputsForIAMOverlayCreate(
	metadata resources.CloudAPIResource,
	urn resource.URN,
	inputs resource.PropertyMap,
	client *googleclient.GoogleClient,
) (resource.PropertyMap, error) {
	newBinding, iamPolicyInputs, err := normalizeIamOverlayInputs(metadata, urn, inputs)
	if err != nil {
		return nil, err
	}

	stateStruct, err := getIAMState(metadata, iamPolicyInputs, client)
	if err != nil {
		return nil, err
	}

	// For IamBinding type, drop any existing bindings matching the role.
	if isIAMBinding(urn) {
		stateStruct.Bindings = filterBindingsWithRoleAndCondition(
			stateStruct.Bindings, newBinding.Role, newBinding.Condition)
	}
	stateStruct.Bindings = mergeIAMBindings(append(stateStruct.Bindings, newBinding))

	statePM := resource.NewPropertyMap(stateStruct)
	iamPolicyInputs[resource.PropertyKey("bindings")] = statePM["bindings"]

	return iamPolicyInputs, nil
}

func inputsForIAMOverlayUpdate(
	metadata resources.CloudAPIResource,
	urn resource.URN,
	inputs resource.PropertyMap,
	oldInputs resource.PropertyMap,
	client *googleclient.GoogleClient,
) (resource.PropertyMap, error) {
	newBinding, iamPolicyInputs, err := normalizeIamOverlayInputs(metadata, urn, inputs)
	if err != nil {
		return nil, err
	}

	oldBinding, _, err := normalizeIamOverlayInputs(metadata, urn, oldInputs)
	if err != nil {
		return nil, err
	}

	stateStruct, err := getIAMState(metadata, iamPolicyInputs, client)
	if err != nil {
		return nil, err
	}

	if isIAMBinding(urn) {
		// For IamBinding type, drop any existing bindings matching the role.
		stateStruct.Bindings = filterBindingsWithRoleAndCondition(
			stateStruct.Bindings, newBinding.Role, newBinding.Condition)
		stateStruct.Bindings = mergeIAMBindings(append(stateStruct.Bindings, newBinding))
	} else {
		stateStruct.Bindings = subtractFromBindings(stateStruct.Bindings, oldBinding)
		stateStruct.Bindings = mergeIAMBindings(append(stateStruct.Bindings, newBinding))
	}

	statePM := resource.NewPropertyMap(stateStruct)
	iamPolicyInputs[resource.PropertyKey("bindings")] = statePM["bindings"]
	iamPolicyInputs[resource.PropertyKey("version")] = resource.NewNumberProperty(3)

	return iamPolicyInputs, nil
}

func inputsForIAMOverlayDelete(
	metadata resources.CloudAPIResource,
	urn resource.URN,
	inputs resource.PropertyMap,
	client *googleclient.GoogleClient,
) (resource.PropertyMap, error) {
	newBinding, iamPolicyInputs, err := normalizeIamOverlayInputs(metadata, urn, inputs)
	if err != nil {
		return nil, err
	}

	stateStruct, err := getIAMState(metadata, iamPolicyInputs, client)
	if err != nil {
		return nil, err
	}

	// For IamBinding type, drop any existing bindings matching the role.
	if isIAMBinding(urn) {
		stateStruct.Bindings = filterBindingsWithRoleAndCondition(
			stateStruct.Bindings, newBinding.Role, newBinding.Condition)
	}
	stateStruct.Bindings = subtractFromBindings(stateStruct.Bindings, newBinding)

	statePM := resource.NewPropertyMap(stateStruct)
	iamPolicyInputs[resource.PropertyKey("bindings")] = statePM["bindings"]

	return iamPolicyInputs, nil
}

// Note: This following functions were adapted from https://github.com/hashicorp/terraform-provider-google-beta/blob/v4.34.0/google-beta/iam.go

// Flattens a list of Bindings so each role+condition has a single Binding with combined members
func mergeIAMBindings(bindings []*iamPolicyBinding) []*iamPolicyBinding {
	bm := mapFromIamBindingList(bindings)
	return listFromIamBindingMap(bm)
}

type conditionKey struct {
	Description string
	Expression  string
	Title       string
}

func conditionKeyFromCondition(condition *iamPolicyCondition) conditionKey {
	if condition == nil {
		return conditionKey{}
	}
	return conditionKey{condition.Description, condition.Expression, condition.Title}
}

func (k conditionKey) Empty() bool {
	return k == conditionKey{}
}

func (k conditionKey) String() string {
	return fmt.Sprintf("%s/%s/%s", k.Title, k.Description, k.Expression)
}

type iamBindingKey struct {
	Role      string
	Condition conditionKey
}

// Removes a single role+condition binding from a list of Bindings
func filterBindingsWithRoleAndCondition(
	b []*iamPolicyBinding,
	role string,
	condition *iamPolicyCondition,
) []*iamPolicyBinding {
	bMap := mapFromIamBindingList(b)
	key := iamBindingKey{role, conditionKeyFromCondition(condition)}
	delete(bMap, key)
	return listFromIamBindingMap(bMap)
}

// Removes given role+condition/bound-member pairs from the given Bindings (i.e subtraction).
func subtractFromBindings(
	bindings []*iamPolicyBinding,
	toRemove ...*iamPolicyBinding,
) []*iamPolicyBinding {
	currMap := mapFromIamBindingList(bindings)
	toRemoveMap := mapFromIamBindingList(toRemove)

	for key, removeSet := range toRemoveMap {
		members, ok := currMap[key]
		if !ok {
			continue
		}
		// Remove all removed members
		for m := range removeSet {
			delete(members, m)
		}
		// Remove role+condition from bindings
		if len(members) == 0 {
			delete(currMap, key)
		}
	}

	return listFromIamBindingMap(currMap)
}

func iamMemberIsCaseSensitive(member string) bool {
	// allAuthenticatedUsers and allUsers are special identifiers that are case sensitive. See:
	// https://cloud.google.com/iam/docs/overview#all-authenticated-users
	return strings.Contains(member, "allAuthenticatedUsers") ||
		strings.Contains(member, "allUsers") ||
		strings.HasPrefix(member, "principal:") ||
		strings.HasPrefix(member, "principalHierarchy:") ||
		strings.HasPrefix(member, "principalSet:")
}

// normalizeIamMemberCasing returns the case adjusted value of an iamMember
// this is important as iam will ignore casing unless it is one of the following
// member types: principalSet, principal, principalHierarchy
// members are in <type>:<value> format
// <type> is case-sensitive
// <value> isn't in most cases
// so lowercase the value unless iamMemberIsCaseSensitive and leave the type alone
// since Dec '19 members can be prefixed with "deleted:" to indicate the principal
// has been deleted
func normalizeIamMemberCasing(member string) string {
	var pieces []string
	if strings.HasPrefix(member, "deleted:") {
		pieces = strings.SplitN(member, ":", 3)
		if len(pieces) > 2 && !iamMemberIsCaseSensitive(strings.TrimPrefix(member, "deleted:")) {
			pieces[2] = strings.ToLower(pieces[2])
		}
	} else if !iamMemberIsCaseSensitive(member) {
		pieces = strings.SplitN(member, ":", 2)
		if len(pieces) > 1 {
			pieces[1] = strings.ToLower(pieces[1])
		}
	}

	if len(pieces) > 0 {
		member = strings.Join(pieces, ":")
	}
	return member
}

// Construct map of role to set of members from list of bindings.
func mapFromIamBindingList(bindings []*iamPolicyBinding) map[iamBindingKey]codegen.StringSet {
	bm := map[iamBindingKey]codegen.StringSet{}
	// Get each binding
	for _, b := range bindings {
		members := make(map[string]struct{})
		key := iamBindingKey{b.Role, conditionKeyFromCondition(b.Condition)}
		// Initialize members map
		if _, ok := bm[key]; ok {
			members = bm[key]
		}
		// Get each member (user/principal) for the binding
		for _, m := range b.Members {
			m = normalizeIamMemberCasing(m)
			// Add the member
			members[m] = struct{}{}
		}
		if len(members) > 0 {
			bm[key] = members
		} else {
			delete(bm, key)
		}
	}
	return bm
}

// Return list of Bindings for a map of role to member sets
func listFromIamBindingMap(bm map[iamBindingKey]codegen.StringSet) []*iamPolicyBinding {
	rb := make([]*iamPolicyBinding, 0, len(bm))
	var keys []iamBindingKey
	for k := range bm {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		keyI := keys[i]
		keyJ := keys[j]
		return fmt.Sprintf("%s%s", keyI.Role, keyI.Condition.String()) < fmt.Sprintf("%s%s", keyJ.Role, keyJ.Condition.String())
	})
	for _, key := range keys {
		members := bm[key]
		if len(members) == 0 {
			continue
		}
		b := &iamPolicyBinding{
			Role:    key.Role,
			Members: members.SortedValues(),
		}
		if !key.Condition.Empty() {
			b.Condition = &iamPolicyCondition{
				Description: key.Condition.Description,
				Expression:  key.Condition.Expression,
				Title:       key.Condition.Title,
			}
		}
		rb = append(rb, b)
	}
	return rb
}
