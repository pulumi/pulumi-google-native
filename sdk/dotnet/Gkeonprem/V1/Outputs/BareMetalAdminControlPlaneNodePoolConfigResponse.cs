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
    /// BareMetalAdminControlPlaneNodePoolConfig specifies the control plane node pool configuration. We have a control plane specific node pool config so that we can flexible about supporting control plane specific fields in the future.
    /// </summary>
    [OutputType]
    public sealed class BareMetalAdminControlPlaneNodePoolConfigResponse
    {
        /// <summary>
        /// The generic configuration for a node pool running the control plane.
        /// </summary>
        public readonly Outputs.BareMetalNodePoolConfigResponse NodePoolConfig;

        [OutputConstructor]
        private BareMetalAdminControlPlaneNodePoolConfigResponse(Outputs.BareMetalNodePoolConfigResponse nodePoolConfig)
        {
            NodePoolConfig = nodePoolConfig;
        }
    }
}