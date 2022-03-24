// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestGKE(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "gke-ts"),
			SkipRefresh: true,
		})

	integration.ProgramTest(t, &test)
}

func TestCloudRunTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "cloudrun-ts", "step1"),
			SkipRefresh: true,
			EditDirs: []integration.EditDir{
				{Dir: filepath.Join(getCwd(t), "cloudrun-ts", "step2")},
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
					Dir:      "step2",
					Additive: true,
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
