// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Inputs
{

    /// <summary>
    /// Specifies the locad balancer config for the VMware user cluster.
    /// </summary>
    public sealed class VmwareLoadBalancerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration for F5 Big IP typed load balancers.
        /// </summary>
        [Input("f5Config")]
        public Input<Inputs.VmwareF5BigIpConfigArgs>? F5Config { get; set; }

        /// <summary>
        /// Manually configured load balancers.
        /// </summary>
        [Input("manualLbConfig")]
        public Input<Inputs.VmwareManualLbConfigArgs>? ManualLbConfig { get; set; }

        /// <summary>
        /// Configuration for MetalLB typed load balancers.
        /// </summary>
        [Input("metalLbConfig")]
        public Input<Inputs.VmwareMetalLbConfigArgs>? MetalLbConfig { get; set; }

        /// <summary>
        /// The VIPs used by the load balancer.
        /// </summary>
        [Input("vipConfig")]
        public Input<Inputs.VmwareVipConfigArgs>? VipConfig { get; set; }

        public VmwareLoadBalancerConfigArgs()
        {
        }
        public static new VmwareLoadBalancerConfigArgs Empty => new VmwareLoadBalancerConfigArgs();
    }
}