// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Configuration for an App Engine network endpoint group (NEG). The service is optional, may be provided explicitly or in the URL mask. The version is optional and can only be provided explicitly or in the URL mask when service is present. Note: App Engine service must be in the same project and located in the same region as the Serverless NEG.
    /// </summary>
    public sealed class NetworkEndpointGroupAppEngineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional serving service. The service name is case-sensitive and must be 1-63 characters long. Example value: "default", "my-service".
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// A template to parse service and version fields from a request URL. URL mask allows for routing to multiple App Engine services without having to create multiple Network Endpoint Groups and backend services. For example, the request URLs "foo1-dot-appname.appspot.com/v1" and "foo1-dot-appname.appspot.com/v2" can be backed by the same Serverless NEG with URL mask "&lt;service&gt;-dot-appname.appspot.com/&lt;version&gt;". The URL mask will parse them to { service = "foo1", version = "v1" } and { service = "foo1", version = "v2" } respectively.
        /// </summary>
        [Input("urlMask")]
        public Input<string>? UrlMask { get; set; }

        /// <summary>
        /// Optional serving version. The version name is case-sensitive and must be 1-100 characters long. Example value: "v1", "v2".
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NetworkEndpointGroupAppEngineArgs()
        {
        }
        public static new NetworkEndpointGroupAppEngineArgs Empty => new NetworkEndpointGroupAppEngineArgs();
    }
}
