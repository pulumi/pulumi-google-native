// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BlockchainNodeEngine.V1.Outputs
{

    /// <summary>
    /// Contains endpoint information through which to interact with a blockchain node.
    /// </summary>
    [OutputType]
    public sealed class EndpointInfoResponse
    {
        /// <summary>
        /// The assigned URL for the node JSON-RPC API endpoint.
        /// </summary>
        public readonly string JsonRpcApiEndpoint;
        /// <summary>
        /// The assigned URL for the node WebSockets API endpoint.
        /// </summary>
        public readonly string WebsocketsApiEndpoint;

        [OutputConstructor]
        private EndpointInfoResponse(
            string jsonRpcApiEndpoint,

            string websocketsApiEndpoint)
        {
            JsonRpcApiEndpoint = jsonRpcApiEndpoint;
            WebsocketsApiEndpoint = websocketsApiEndpoint;
        }
    }
}