// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Outputs
{

    /// <summary>
    /// Specifies the cluster network configuration.
    /// </summary>
    [OutputType]
    public sealed class BareMetalNetworkConfigResponse
    {
        /// <summary>
        /// Enables the use of advanced Anthos networking features, such as Bundled Load Balancing with BGP or the egress NAT gateway. Setting configuration for advanced networking features will automatically set this flag.
        /// </summary>
        public readonly bool AdvancedNetworking;
        /// <summary>
        /// Configuration for island mode CIDR. In an island-mode network, nodes have unique IP addresses, but pods don't have unique addresses across clusters. This doesn't cause problems because pods in one cluster never directly communicate with pods in another cluster. Instead, there are gateways that mediate between a pod in one cluster and a pod in another cluster.
        /// </summary>
        public readonly Outputs.BareMetalIslandModeCidrConfigResponse IslandModeCidr;
        /// <summary>
        /// Configuration for multiple network interfaces.
        /// </summary>
        public readonly Outputs.BareMetalMultipleNetworkInterfacesConfigResponse MultipleNetworkInterfacesConfig;
        /// <summary>
        /// Configuration for SR-IOV.
        /// </summary>
        public readonly Outputs.BareMetalSrIovConfigResponse SrIovConfig;

        [OutputConstructor]
        private BareMetalNetworkConfigResponse(
            bool advancedNetworking,

            Outputs.BareMetalIslandModeCidrConfigResponse islandModeCidr,

            Outputs.BareMetalMultipleNetworkInterfacesConfigResponse multipleNetworkInterfacesConfig,

            Outputs.BareMetalSrIovConfigResponse srIovConfig)
        {
            AdvancedNetworking = advancedNetworking;
            IslandModeCidr = islandModeCidr;
            MultipleNetworkInterfacesConfig = multipleNetworkInterfacesConfig;
            SrIovConfig = srIovConfig;
        }
    }
}