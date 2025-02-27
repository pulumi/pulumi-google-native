// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Configuration options for private clusters.
    /// </summary>
    public sealed class PrivateClusterConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the master's internal IP address is used as the cluster endpoint.
        /// </summary>
        [Input("enablePrivateEndpoint")]
        public Input<bool>? EnablePrivateEndpoint { get; set; }

        /// <summary>
        /// Whether nodes have internal IP addresses only. If enabled, all nodes are given only RFC 1918 private addresses and communicate with the master via private networking.
        /// </summary>
        [Input("enablePrivateNodes")]
        public Input<bool>? EnablePrivateNodes { get; set; }

        /// <summary>
        /// Controls master global access settings.
        /// </summary>
        [Input("masterGlobalAccessConfig")]
        public Input<Inputs.PrivateClusterMasterGlobalAccessConfigArgs>? MasterGlobalAccessConfig { get; set; }

        /// <summary>
        /// The IP range in CIDR notation to use for the hosted master network. This range will be used for assigning internal IP addresses to the master or set of masters, as well as the ILB VIP. This range must not overlap with any other ranges in use within the cluster's network.
        /// </summary>
        [Input("masterIpv4CidrBlock")]
        public Input<string>? MasterIpv4CidrBlock { get; set; }

        /// <summary>
        /// Subnet to provision the master's private endpoint during cluster creation. Specified in projects/*/regions/*/subnetworks/* format.
        /// </summary>
        [Input("privateEndpointSubnetwork")]
        public Input<string>? PrivateEndpointSubnetwork { get; set; }

        public PrivateClusterConfigArgs()
        {
        }
        public static new PrivateClusterConfigArgs Empty => new PrivateClusterConfigArgs();
    }
}
