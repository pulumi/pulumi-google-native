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

//go:build nodejs || all
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
)

func TestGKE(t *testing.T) {
	clusterDiffs := codegen.NewStringSet("resourceLabels")
	nodepoolDiffs := codegen.NewStringSet("config")
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:           filepath.Join(getCwd(t), "gke-ts", "step1"),
			SkipRefresh:   true,
			RunUpdateTest: false,
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "gke-ts", "step2"),
					Additive: true,
					ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
						foundClusterUpdate, foundNodepoolUpdate := false, false
						for _, event := range stack.Events {
							if res := event.ResOutputsEvent; res != nil && res.Metadata.Op == apitype.OpUpdate {
								if res.Metadata.Type == "google-native:container/v1:Cluster" {
									for _, field := range res.Metadata.Diffs {
										if clusterDiffs.Has(field) {
											foundClusterUpdate = true
										}
									}
								}
								if res.Metadata.Type == "google-native:container/v1:NodePool" {
									for _, field := range res.Metadata.Diffs {
										if nodepoolDiffs.Has(field) {
											foundNodepoolUpdate = true
										}
									}
								}
							}
						}
						assert.True(t, foundClusterUpdate, "should have seen a cluster update")
						assert.True(t, foundNodepoolUpdate, "should have seen a nodepool update")
						t.Logf("outputs: %+v", stack.Outputs)
					},
				},
			},
		})
	integration.ProgramTest(t, &test)
}

func TestCloudRunTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "cloudrun-ts", "step1"),
			SkipRefresh: true,
			EditDirs: []integration.EditDir{
				{Dir: filepath.Join(getCwd(t), "cloudrun-ts", "step2"), Additive: true},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestFunctionsTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "functions-ts"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				assertHTTPHelloWorld(t, stack.Outputs["functionUrl"].(string), nil)
			},
			SkipRefresh: true,
		})

	integration.ProgramTest(t, &test)
}

func TestPubSubTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "pubsub-ts"),
			SkipRefresh: true,
		})

	integration.ProgramTest(t, &test)
}

func TestIAMServiceAccountTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "iam-serviceaccount-ts", "step1"),
			SkipRefresh: true,
			EditDirs: []integration.EditDir{
				{Dir: filepath.Join(getCwd(t), "iam-serviceaccount-ts", "step2"), Additive: true},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestWebserverTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "webserver-ts"),
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				assertHTTPMatchesContent(t, stack.Outputs["instanceIP"].(string), "Hello, World!\n", nil)
			},
			SkipRefresh: true,
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "webserver-ts", "step2"),
					Additive: true,
					ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
						assertHTTPMatchesContent(t, stack.Outputs["instanceIP"].(string), "Hello, World!\n", nil)
						assert.Equal(t, stack.Outputs["tag"].(string), "test")
					},
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestStorageTransferTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:                  filepath.Join(getCwd(t), "storagetransfer-ts"),
			ExpectRefreshChanges: true,
		})
	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/google-native",
		},
	})

	return baseJS
}

func TestProjectIamPolicyTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "project-iam-policy-ts"),
			SkipRefresh: true,
			SkipUpdate:  true,
			Quick:       true,
		})

	integration.ProgramTest(t, &test)
}

func TestBucketIamTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "bucket-iam-ts"),
			SkipRefresh: true,
			SkipUpdate:  true,
			Quick:       true,
		})

	integration.ProgramTest(t, &test)
}

func TestGetClientConfigTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "get-client-config-ts"),
			SkipRefresh: true,
			Quick:       true,
			ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
				clientConfig := stack.Outputs["clientConfig"].(map[string]interface{})
				token := stack.Outputs["token"].(map[string]interface{})
				assert.True(t, len(clientConfig) > 0)
				assert.True(t, len(token) > 0)
			},
		})

	integration.ProgramTest(t, &test)
}
