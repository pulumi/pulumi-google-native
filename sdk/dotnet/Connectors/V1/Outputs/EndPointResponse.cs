// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Outputs
{

    /// <summary>
    /// Endpoint message includes details of the Destination endpoint.
    /// </summary>
    [OutputType]
    public sealed class EndPointResponse
    {
        /// <summary>
        /// The URI of the Endpoint.
        /// </summary>
        public readonly string EndpointUri;
        /// <summary>
        /// List of Header to be added to the Endpoint.
        /// </summary>
        public readonly ImmutableArray<Outputs.HeaderResponse> Headers;

        [OutputConstructor]
        private EndPointResponse(
            string endpointUri,

            ImmutableArray<Outputs.HeaderResponse> headers)
        {
            EndpointUri = endpointUri;
            Headers = headers;
        }
    }
}