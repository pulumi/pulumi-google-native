// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Outputs
{

    /// <summary>
    /// The setting used to configure ClientGateways. It is adding routes to the client's routing table after the connection is established.
    /// </summary>
    [OutputType]
    public sealed class DestinationRouteResponse
    {
        /// <summary>
        /// The network address of the subnet for which the packet is routed to the ClientGateway.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// The network mask of the subnet for which the packet is routed to the ClientGateway.
        /// </summary>
        public readonly string Netmask;

        [OutputConstructor]
        private DestinationRouteResponse(
            string address,

            string netmask)
        {
            Address = address;
            Netmask = netmask;
        }
    }
}