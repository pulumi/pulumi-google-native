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
    /// Define the Connectors target endpoint.
    /// </summary>
    [OutputType]
    public sealed class DestinationConfigResponse
    {
        /// <summary>
        /// The destinations for the key.
        /// </summary>
        public readonly ImmutableArray<Outputs.DestinationResponse> Destinations;
        /// <summary>
        /// The key is the destination identifier that is supported by the Connector.
        /// </summary>
        public readonly string Key;

        [OutputConstructor]
        private DestinationConfigResponse(
            ImmutableArray<Outputs.DestinationResponse> destinations,

            string key)
        {
            Destinations = destinations;
            Key = key;
        }
    }
}