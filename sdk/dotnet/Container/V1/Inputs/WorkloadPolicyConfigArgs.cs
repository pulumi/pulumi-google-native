// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// WorkloadPolicyConfig is the configuration of workload policy for autopilot clusters.
    /// </summary>
    public sealed class WorkloadPolicyConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If true, workloads can use NET_ADMIN capability.
        /// </summary>
        [Input("allowNetAdmin")]
        public Input<bool>? AllowNetAdmin { get; set; }

        public WorkloadPolicyConfigArgs()
        {
        }
        public static new WorkloadPolicyConfigArgs Empty => new WorkloadPolicyConfigArgs();
    }
}
