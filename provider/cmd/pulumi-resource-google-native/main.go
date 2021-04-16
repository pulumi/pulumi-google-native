// Copyright 2016-2021, Pulumi Corporation.

package main

import (
	"github.com/pulumi/pulumi-google-native/provider/pkg/provider"
	"github.com/pulumi/pulumi-google-native/provider/pkg/version"
)

var providerName = "google-native"

func main() {
	provider.Serve(providerName, version.Version, pulumiSchema, cloudApiResources)
}
