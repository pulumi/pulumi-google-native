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
    /// Specifies the VIP config for the VMware user cluster load balancer.
    /// </summary>
    [OutputType]
    public sealed class VmwareVipConfigResponse
    {
        /// <summary>
        /// The VIP which you previously set aside for the Kubernetes API of this cluster.
        /// </summary>
        public readonly string ControlPlaneVip;
        /// <summary>
        /// The VIP which you previously set aside for ingress traffic into this cluster.
        /// </summary>
        public readonly string IngressVip;

        [OutputConstructor]
        private VmwareVipConfigResponse(
            string controlPlaneVip,

            string ingressVip)
        {
            ControlPlaneVip = controlPlaneVip;
            IngressVip = ingressVip;
        }
    }
}