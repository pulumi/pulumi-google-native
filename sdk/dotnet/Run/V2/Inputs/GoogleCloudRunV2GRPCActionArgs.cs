// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// GRPCAction describes an action involving a GRPC port.
    /// </summary>
    public sealed class GoogleCloudRunV2GRPCActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port number of the gRPC service. Number must be in the range 1 to 65535. If not specified, defaults to 8080.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public GoogleCloudRunV2GRPCActionArgs()
        {
        }
        public static new GoogleCloudRunV2GRPCActionArgs Empty => new GoogleCloudRunV2GRPCActionArgs();
    }
}