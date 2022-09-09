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
	"encoding/json"
	"reflect"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
)

// Note: This file was adapted from https://github.com/hashicorp/terraform-provider-google-beta/blob/v4.34.0/google-beta/iam_test.go

func TestIamMergeBindings(t *testing.T) {
	testCases := []struct {
		input  []*iamPolicyBinding
		expect []*iamPolicyBinding
	}{
		// Nothing to merge - return same list
		{
			input:  []*iamPolicyBinding{},
			expect: []*iamPolicyBinding{},
		},
		// No members returns no binding
		{
			input: []*iamPolicyBinding{
				{
					Role: "role-1",
				},
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
		},
		// Nothing to merge - return same list
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
			},
		},
		// Nothing to merge - return same list
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
		},
		// Merge members for role
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-2"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
		},
		// Merge members for role
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-3"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2", "member-3"},
				},
			},
		},
		// Complex merge
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-3", "member-4"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-2", "member-1"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-5"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
				{
					Role:    "empty-role",
					Members: []string{},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2", "member-3", "member-4", "member-5"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1"},
				},
			},
		},
		// Same role+members, different condition
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
			},
		},
		// Same role, same condition
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
				{
					Role:    "role-1",
					Members: []string{"member-3"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2", "member-3"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
			},
		},
		// Different roles, same condition
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
				{
					Role:    "role-2",
					Members: []string{"member-3"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
				{
					Role:    "role-2",
					Members: []string{"member-3"},
					Condition: &iamPolicyCondition{
						Title: "condition-1",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		got := mergeIAMBindings(tc.input)
		if !compareBindings(got, tc.expect) {
			t.Errorf("Unexpected value for mergeIAMBindings(%s).\nActual: %s\nExpected: %s\n",
				debugPrintBindings(tc.input), debugPrintBindings(got), debugPrintBindings(tc.expect))
		}
	}
}

func TestIamFilterBindingsWithRoleAndCondition(t *testing.T) {
	testCases := []struct {
		input          []*iamPolicyBinding
		role           string
		conditionTitle string
		expect         []*iamPolicyBinding
	}{
		// No-op
		{
			input:  []*iamPolicyBinding{},
			role:   "role-1",
			expect: []*iamPolicyBinding{},
		},
		// Remove one binding
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
			role:   "role-1",
			expect: []*iamPolicyBinding{},
		},
		// Remove multiple bindings
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-3"},
				},
			},
			role:   "role-1",
			expect: []*iamPolicyBinding{},
		},
		// Remove multiple bindings and leave some.
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1", "member-3"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-2"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1", "member-2"},
				},
			},
			role: "role-1",
			expect: []*iamPolicyBinding{
				{
					Role:    "role-2",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1", "member-3"},
				},
			},
		},
		// Remove one binding with condition
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
				{
					Role:      "role-1",
					Members:   []string{"member-3", "member-4"},
					Condition: &iamPolicyCondition{Title: "condition-1"},
				},
			},
			role:           "role-1",
			conditionTitle: "condition-1",
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
		},
	}

	for _, tc := range testCases {
		got := filterBindingsWithRoleAndCondition(tc.input, tc.role, &iamPolicyCondition{Title: tc.conditionTitle})
		if !compareBindings(got, tc.expect) {
			t.Errorf("Got unexpected value for removeAllBindingsWithRole(%s, %s).\nActual: %s\nExpected: %s",
				debugPrintBindings(tc.input), tc.role, debugPrintBindings(got), debugPrintBindings(tc.expect))
		}
	}
}

func TestIamSubtractFromBindings(t *testing.T) {
	testCases := []struct {
		input  []*iamPolicyBinding
		remove []*iamPolicyBinding
		expect []*iamPolicyBinding
	}{
		// No-op
		{
			input:  []*iamPolicyBinding{},
			remove: []*iamPolicyBinding{},
			expect: []*iamPolicyBinding{},
		},
		// Empty input should no-op return empty
		{
			input: []*iamPolicyBinding{},
			remove: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
			expect: []*iamPolicyBinding{},
		},
		// Empty removal should return original input
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
			remove: []*iamPolicyBinding{},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
		},
		// Removal not in input should no-op
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-1+"},
				},
			},
			remove: []*iamPolicyBinding{
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-1+"},
				},
			},
		},
		// Same input/remove should return empty
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
			remove: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
			expect: []*iamPolicyBinding{},
		},
		// Single removal
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-2"},
				},
			},
			remove: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-2"},
				},
			},
		},
		// Complex removal
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-2", "member-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
			remove: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-2", "member-4"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-2"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1"},
				},
			},
		},
		// With conditions
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-2", "member-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-1",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1"},
				},
				{
					Role:      "role-2",
					Members:   []string{"member-1"},
					Condition: &iamPolicyCondition{Title: "condition-1"},
				},
			},
			remove: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-2", "member-4"},
				},
				{
					Role:      "role-2",
					Members:   []string{"member-1"},
					Condition: &iamPolicyCondition{Title: "condition-1"},
				},
			},
			expect: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"member-1", "member-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"member-1"},
				},
				{
					Role:    "role-3",
					Members: []string{"member-1"},
				},
			},
		},
	}

	for _, tc := range testCases {
		got := subtractFromBindings(tc.input, tc.remove...)
		if !compareBindings(got, tc.expect) {
			t.Errorf("Unexpected value for subtractFromBindings(%s, %s).\nActual: %s\nExpected: %s\n",
				debugPrintBindings(tc.input), debugPrintBindings(tc.remove), debugPrintBindings(got), debugPrintBindings(tc.expect))
		}
	}
}

func TestIamCreateIamBindingsMap(t *testing.T) {
	testCases := []struct {
		input  []*iamPolicyBinding
		expect map[iamBindingKey]codegen.StringSet
	}{
		// Empty
		{
			input:  []*iamPolicyBinding{},
			expect: map[iamBindingKey]codegen.StringSet{},
		},
		// Single role, multiple members
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"user-1", "user-2"},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("user-1", "user-2"),
			},
		},
		// Multiple bindings for the same role
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"user-1", "user-2"},
				},
				{
					Role:    "role-1",
					Members: []string{"user-3"},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("user-1", "user-2", "user-3"),
			},
		},
		// Multiple roles+members
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"user-1", "user-2"},
				},
				{
					Role:    "role-2",
					Members: []string{"user-1"},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("user-1", "user-2"),
				{"role-2", conditionKey{}}: codegen.NewStringSet("user-1"),
			},
		},
		// Multiple roles+members
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"user-1", "user-2"},
				},
				{
					Role:    "role-2",
					Members: []string{"user-1"},
				},
				{
					Role:    "role-1",
					Members: []string{"user-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"user-2"},
				},
				{
					Role:    "role-3",
					Members: []string{"user-3"},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("user-1", "user-2", "user-3"),
				{"role-2", conditionKey{}}: codegen.NewStringSet("user-1", "user-2"),
				{"role-3", conditionKey{}}: codegen.NewStringSet("user-3"),
			},
		},
		// Verify case sensitivity for member strings across multiple roles
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"deleted:serviceAccount:useR-1", "user-2"},
				},
				{
					Role:    "role-2",
					Members: []string{"deleted:user:user-1"},
				},
				{
					Role:    "role-1",
					Members: []string{"serviceAccount:user-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"user-2"},
				},
				{
					Role:    "role-3",
					Members: []string{"user-3"},
				},
				{
					Role:    "role-4",
					Members: []string{"deleted:principal:useR-1"},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("deleted:serviceAccount:user-1", "user-2", "serviceAccount:user-3"),
				{"role-2", conditionKey{}}: codegen.NewStringSet("deleted:user:user-1", "user-2"),
				{"role-3", conditionKey{}}: codegen.NewStringSet("user-3"),
				{"role-4", conditionKey{}}: codegen.NewStringSet("deleted:principal:useR-1"),
			},
		},
		// Verify case sensitivity for member strings across multiple roles
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"principalSet://iam.googleapis.com/projects/1066737951711/locations/global/workloadIdentityPools/example-pool/attribute.aws_role/arn:aws:sts::999999999999:assumed-role/some-eu-central-1-lambdaRole"},
				},
				{
					Role:    "role-2",
					Members: []string{"principal://iam.googleapis.com/projects/1066737951711/locations/global/workloadIdentityPools/example-pool/attribute.aws_role/arn:aws:sts::999999999999:assumed-role/some-eu-central-1-lambdaRole"},
				},
				{
					Role:    "role-1",
					Members: []string{"serviceAccount:useR-3"},
				},
				{
					Role:    "role-2",
					Members: []string{"user-2"},
				},
				{
					Role:    "role-3",
					Members: []string{"user-3"},
				},
				{
					Role:    "role-3",
					Members: []string{"principalHierarchy://iam.googleapis.com/projects/1066737951711/locations/global/workloadIdentityPools"},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("principalSet://iam.googleapis.com/projects/1066737951711/locations/global/workloadIdentityPools/example-pool/attribute.aws_role/arn:aws:sts::999999999999:assumed-role/some-eu-central-1-lambdaRole", "serviceAccount:user-3"),
				{"role-2", conditionKey{}}: codegen.NewStringSet("principal://iam.googleapis.com/projects/1066737951711/locations/global/workloadIdentityPools/example-pool/attribute.aws_role/arn:aws:sts::999999999999:assumed-role/some-eu-central-1-lambdaRole", "user-2"),
				{"role-3", conditionKey{}}: codegen.NewStringSet("principalHierarchy://iam.googleapis.com/projects/1066737951711/locations/global/workloadIdentityPools", "user-3"),
			},
		},
		// Multiple roles+members with conditions
		{
			input: []*iamPolicyBinding{
				{
					Role:    "role-1",
					Members: []string{"user-1", "user-2"},
				},
				{
					Role:    "role-2",
					Members: []string{"user-1"},
					Condition: &iamPolicyCondition{
						Title:       "condition-1",
						Description: "condition-1-desc",
						Expression:  "condition-1-expr",
					},
				},
				{
					Role:    "role-2",
					Members: []string{"user-2"},
					Condition: &iamPolicyCondition{
						Title:       "condition-1",
						Description: "condition-1-desc",
						Expression:  "condition-1-expr",
					},
				},
				{
					Role:    "role-2",
					Members: []string{"user-1"},
					Condition: &iamPolicyCondition{
						Title:       "condition-2",
						Description: "condition-2-desc",
						Expression:  "condition-2-expr",
					},
				},
			},
			expect: map[iamBindingKey]codegen.StringSet{
				{"role-1", conditionKey{}}: codegen.NewStringSet("user-1", "user-2"),
				{
					Role: "role-2",
					Condition: conditionKey{
						Title:       "condition-1",
						Description: "condition-1-desc",
						Expression:  "condition-1-expr",
					},
				}: codegen.NewStringSet("user-1", "user-2"),
				{
					Role: "role-2",
					Condition: conditionKey{
						Title:       "condition-2",
						Description: "condition-2-desc",
						Expression:  "condition-2-expr",
					},
				}: codegen.NewStringSet("user-1"),
			},
		},
	}

	for _, tc := range testCases {
		got := mapFromIamBindingList(tc.input)
		if !reflect.DeepEqual(got, tc.expect) {
			t.Errorf("Unexpected value for mapFromIamBindingList(%s).\nActual: %#v\nExpected: %#v\n",
				debugPrintBindings(tc.input), got, tc.expect)
		}
	}
}

// Util to deref and print bindings
func debugPrintBindings(bs []*iamPolicyBinding) string {
	v, _ := json.MarshalIndent(bs, "", "\t")
	return string(v)
}

func compareBindings(a, b []*iamPolicyBinding) bool {
	aMap := mapFromIamBindingList(a)
	bMap := mapFromIamBindingList(b)
	return reflect.DeepEqual(aMap, bMap)
}
