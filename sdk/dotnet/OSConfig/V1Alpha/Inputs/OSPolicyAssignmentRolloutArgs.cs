// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Inputs
{

    /// <summary>
    /// Message to configure the rollout at the zonal level for the OS policy assignment.
    /// </summary>
    public sealed class OSPolicyAssignmentRolloutArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number (or percentage) of VMs per zone to disrupt at any given moment.
        /// </summary>
        [Input("disruptionBudget", required: true)]
        public Input<Inputs.FixedOrPercentArgs> DisruptionBudget { get; set; } = null!;

        /// <summary>
        /// This determines the minimum duration of time to wait after the configuration changes are applied through the current rollout. A VM continues to count towards the `disruption_budget` at least until this duration of time has passed after configuration changes are applied.
        /// </summary>
        [Input("minWaitDuration", required: true)]
        public Input<string> MinWaitDuration { get; set; } = null!;

        public OSPolicyAssignmentRolloutArgs()
        {
        }
    }
}