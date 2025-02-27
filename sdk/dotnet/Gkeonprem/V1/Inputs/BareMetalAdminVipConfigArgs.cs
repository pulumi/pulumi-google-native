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
    /// BareMetalAdminVipConfig for bare metal load balancer configurations.
    /// </summary>
    public sealed class BareMetalAdminVipConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VIP which you previously set aside for the Kubernetes API of this bare metal admin cluster.
        /// </summary>
        [Input("controlPlaneVip")]
        public Input<string>? ControlPlaneVip { get; set; }

        public BareMetalAdminVipConfigArgs()
        {
        }
        public static new BareMetalAdminVipConfigArgs Empty => new BareMetalAdminVipConfigArgs();
    }
}
