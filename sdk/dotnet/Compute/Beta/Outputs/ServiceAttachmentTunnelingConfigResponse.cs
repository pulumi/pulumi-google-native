// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Use to configure this PSC connection in tunneling mode. In tunneling mode traffic from consumer to producer will be encapsulated as it crosses the VPC boundary and traffic from producer to consumer will be decapsulated in the same manner.
    /// </summary>
    [OutputType]
    public sealed class ServiceAttachmentTunnelingConfigResponse
    {
        /// <summary>
        /// Specify the encapsulation protocol and what metadata to include in incoming encapsulated packet headers.
        /// </summary>
        public readonly string EncapsulationProfile;
        /// <summary>
        /// How this Service Attachment will treat traffic sent to the tunnel_ip, destined for the consumer network.
        /// </summary>
        public readonly string RoutingMode;

        [OutputConstructor]
        private ServiceAttachmentTunnelingConfigResponse(
            string encapsulationProfile,

            string routingMode)
        {
            EncapsulationProfile = encapsulationProfile;
            RoutingMode = routingMode;
        }
    }
}