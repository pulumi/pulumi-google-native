// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Inputs
{

    /// <summary>
    /// Autoscaling Policy config associated with the cluster.
    /// </summary>
    public sealed class AutoscalingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The autoscaling policy used by the cluster.Only resource names including projectid and location (region) are valid. Examples: https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id] projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]Note that the policy must be in the same project and Dataproc region.
        /// </summary>
        [Input("policyUri")]
        public Input<string>? PolicyUri { get; set; }

        public AutoscalingConfigArgs()
        {
        }
        public static new AutoscalingConfigArgs Empty => new AutoscalingConfigArgs();
    }
}
