// Copyright 2016-2020, Pulumi Corporation.

package main

import (
	"github.com/pulumi/pulumi-google-cloud/provider/pkg/provider"
	"github.com/pulumi/pulumi-google-cloud/provider/pkg/version"
)

var providerName = "google-cloud"

func main() {
	provider.Serve(providerName, version.Version, nil)
}
