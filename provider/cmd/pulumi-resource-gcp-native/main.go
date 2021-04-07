// Copyright 2016-2021, Pulumi Corporation.

package main

import (
	"github.com/pulumi/pulumi-gcp-native/provider/pkg/provider"
	"github.com/pulumi/pulumi-gcp-native/provider/pkg/version"
)

var providerName = "gcp-native"

func main() {
	provider.Serve(providerName, version.Version, pulumiSchema, cloudApiResources)
}
