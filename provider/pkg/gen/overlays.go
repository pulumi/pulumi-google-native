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

package gen

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

// These "overlay" resources are hand-authored rather than being generated directly from the discovery documents. The
// resource consists of a schematized SDK projection, plus a custom implementation in the provider.

var iamBindingSpec = schema.ResourceSpec{
	ObjectTypeSpec: schema.ObjectTypeSpec{
		Description: "TODO",
		Type:        "object",
		Properties: map[string]schema.PropertySpec{
			"condition": {
				Description: "An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.",
				TypeSpec: schema.TypeSpec{
					Ref: "#/types/google-native:iam/v1:Condition",
				},
			},
			"etag": {
				Description: "The etag of the resource's IAM policy.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"members": {
				Description: "Identities that will be granted the privilege in role. Each entry can have one of the following values:\n\n * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.\n * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.\n * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.\n * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.",
				TypeSpec: schema.TypeSpec{
					Type: "array",
					Items: &schema.TypeSpec{
						Type: "string",
					},
				},
			},
			"name": {
				Description: "The name of the resource to manage IAM policies for.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"project": {
				Description: "The project in which the resource belongs. If it is not provided, a default will be supplied.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"role": {
				Description: "The role that should be applied. Only one `IamBinding` can be used per role.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
		},
		Required: []string{
			"etag",
			"members",
			"name",
			"project",
			"role",
		},
	},
	InputProperties: map[string]schema.PropertySpec{
		"condition": {
			Description: "An IAM Condition for a given binding.",
			TypeSpec: schema.TypeSpec{
				Ref: "#/types/google-native:iam/v1:Condition",
			},
		},
		"members": {
			Description: "Identities that will be granted the privilege in role. Each entry can have one of the following values:\n\n * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.\n * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.\n * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.\n * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.",
			TypeSpec: schema.TypeSpec{
				Type: "array",
				Items: &schema.TypeSpec{
					Type: "string",
				},
			},
		},
		"name": {
			Description: "The name of the resource to manage IAM policies for.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"role": {
			Description: "The role that should be applied. Only one `IamBinding` can be used per role.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
	},
	RequiredInputs: []string{
		"members",
		"name",
		"role",
	},
}

var iamMemberSpec = schema.ResourceSpec{
	ObjectTypeSpec: schema.ObjectTypeSpec{
		Description: "TODO",
		Type:        "object",
		Properties: map[string]schema.PropertySpec{
			"condition": {
				Description: "An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.",
				TypeSpec: schema.TypeSpec{
					Ref: "#/types/google-native:iam/v1:Condition",
				},
			},
			"etag": {
				Description: "The etag of the resource's IAM policy.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"member": {
				Description: "Identity that will be granted the privilege in role. The entry can have one of the following values:\n\n * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.\n * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.\n * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.\n * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"name": {
				Description: "The name of the resource to manage IAM policies for.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"project": {
				Description: "The project in which the resource belongs. If it is not provided, a default will be supplied.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
			"role": {
				Description: "The role that should be applied. Only one `IamBinding` can be used per role.",
				TypeSpec: schema.TypeSpec{
					Type: "string",
				},
			},
		},
		Required: []string{
			"etag",
			"member",
			"name",
			"project",
			"role",
		},
	},
	InputProperties: map[string]schema.PropertySpec{
		"condition": {
			Description: "An IAM Condition for a given binding.",
			TypeSpec: schema.TypeSpec{
				Ref: "#/types/google-native:iam/v1:Condition",
			},
		},
		"member": {
			Description: "Identity that will be granted the privilege in role. The entry can have one of the following values:\n\n * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.\n * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.\n * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.\n * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"name": {
			Description: "The name of the resource to manage IAM policies for.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
		"role": {
			Description: "The role that should be applied. Only one `IamBinding` can be used per role.",
			TypeSpec: schema.TypeSpec{
				Type: "string",
			},
		},
	},
	RequiredInputs: []string{
		"member",
		"name",
		"role",
	},
}
