// Copyright 2016-2022, Pulumi Corporation.  All rights reserved.
//go:build python || all
// +build python all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "python", "bin"),
		},
	})

	return basePy
}

func TestBigTableInstance(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	options := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
		Dir:         filepath.Join(cwd, "bigtable-py", "step1"),
		SkipRefresh: true,
		Config: map[string]string{
			"google-native:project":  os.Getenv("GOOGLE_PROJECT"),
			"google-native:location": os.Getenv("GOOGLE_ZONE"),
		},
		EditDirs: []integration.EditDir{
			{
				Dir:      filepath.Join(cwd, "bigtable-py", "step2"),
				Additive: true,
			},
			{
				Dir:      filepath.Join(cwd, "bigtable-py", "step3"),
				Additive: true,
			},
		},
	})
	integration.ProgramTest(t, &options)
}

func TestGetKubeConfigPy(t *testing.T) {
	cwd, err := os.Getwd()
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	options := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
		Dir:         filepath.Join(cwd, "kubeconfig-py"),
		SkipRefresh: true,
		Config: map[string]string{
			"google-native:project":  os.Getenv("GOOGLE_PROJECT"),
			"google-native:location": os.Getenv("GOOGLE_ZONE"),
		},
		RunUpdateTest: false,
	})
	integration.ProgramTest(t, &options)
}
