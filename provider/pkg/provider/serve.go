// Copyright 2016-2021, Pulumi Corporation.

package provider

import (
	"github.com/pulumi/pulumi/pkg/v2/resource/provider"
	"github.com/pulumi/pulumi/sdk/v2/go/common/util/cmdutil"
	rpc "github.com/pulumi/pulumi/sdk/v2/proto/go"
)

// Serve launches the gRPC server for the resource provider.
func Serve(providerName, version string, schemaBytes []byte, metadataBytes []byte) {
	// Start gRPC service.
	err := provider.Main(providerName, func(host *provider.HostClient) (rpc.ResourceProviderServer, error) {
		return makeProvider(host, providerName, version, schemaBytes, metadataBytes)
	})
	if err != nil {
		cmdutil.ExitError(err.Error())
	}
}
