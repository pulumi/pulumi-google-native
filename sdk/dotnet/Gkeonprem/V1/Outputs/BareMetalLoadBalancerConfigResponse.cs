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
    /// Specifies the load balancer configuration.
    /// </summary>
    [OutputType]
    public sealed class BareMetalLoadBalancerConfigResponse
    {
        /// <summary>
        /// Configuration for BGP typed load balancers. When set network_config.advanced_networking is automatically set to true.
        /// </summary>
        public readonly Outputs.BareMetalBgpLbConfigResponse BgpLbConfig;
        /// <summary>
        /// Manually configured load balancers.
        /// </summary>
        public readonly Outputs.BareMetalManualLbConfigResponse ManualLbConfig;
        /// <summary>
        /// Configuration for MetalLB load balancers.
        /// </summary>
        public readonly Outputs.BareMetalMetalLbConfigResponse MetalLbConfig;
        /// <summary>
        /// Configures the ports that the load balancer will listen on.
        /// </summary>
        public readonly Outputs.BareMetalPortConfigResponse PortConfig;
        /// <summary>
        /// The VIPs used by the load balancer.
        /// </summary>
        public readonly Outputs.BareMetalVipConfigResponse VipConfig;

        [OutputConstructor]
        private BareMetalLoadBalancerConfigResponse(
            Outputs.BareMetalBgpLbConfigResponse bgpLbConfig,

            Outputs.BareMetalManualLbConfigResponse manualLbConfig,

            Outputs.BareMetalMetalLbConfigResponse metalLbConfig,

            Outputs.BareMetalPortConfigResponse portConfig,

            Outputs.BareMetalVipConfigResponse vipConfig)
        {
            BgpLbConfig = bgpLbConfig;
            ManualLbConfig = manualLbConfig;
            MetalLbConfig = metalLbConfig;
            PortConfig = portConfig;
            VipConfig = vipConfig;
        }
    }
}