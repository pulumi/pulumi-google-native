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
    /// Specifies the VIP config for the bare metal load balancer.
    /// </summary>
    [OutputType]
    public sealed class BareMetalVipConfigResponse
    {
        /// <summary>
        /// The VIP which you previously set aside for the Kubernetes API of this bare metal user cluster.
        /// </summary>
        public readonly string ControlPlaneVip;
        /// <summary>
        /// The VIP which you previously set aside for ingress traffic into this bare metal user cluster.
        /// </summary>
        public readonly string IngressVip;

        [OutputConstructor]
        private BareMetalVipConfigResponse(
            string controlPlaneVip,

            string ingressVip)
        {
            ControlPlaneVip = controlPlaneVip;
            IngressVip = ingressVip;
        }
    }
}