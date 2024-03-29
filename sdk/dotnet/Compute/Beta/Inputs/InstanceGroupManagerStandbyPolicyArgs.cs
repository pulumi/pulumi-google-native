// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    public sealed class InstanceGroupManagerStandbyPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("initialDelaySec")]
        public Input<int>? InitialDelaySec { get; set; }

        /// <summary>
        /// Defines behaviour of using instances from standby pool to resize MIG.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.Compute.Beta.InstanceGroupManagerStandbyPolicyMode>? Mode { get; set; }

        public InstanceGroupManagerStandbyPolicyArgs()
        {
        }
        public static new InstanceGroupManagerStandbyPolicyArgs Empty => new InstanceGroupManagerStandbyPolicyArgs();
    }
}
