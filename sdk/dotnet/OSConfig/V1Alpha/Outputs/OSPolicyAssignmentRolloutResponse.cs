// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Alpha.Outputs
{

    /// <summary>
    /// Message to configure the rollout at the zonal level for the OS policy assignment.
    /// </summary>
    [OutputType]
    public sealed class OSPolicyAssignmentRolloutResponse
    {
        /// <summary>
        /// The maximum number (or percentage) of VMs per zone to disrupt at any given moment.
        /// </summary>
        public readonly Outputs.FixedOrPercentResponse DisruptionBudget;
        /// <summary>
        /// This determines the minimum duration of time to wait after the configuration changes are applied through the current rollout. A VM continues to count towards the `disruption_budget` at least until this duration of time has passed after configuration changes are applied.
        /// </summary>
        public readonly string MinWaitDuration;

        [OutputConstructor]
        private OSPolicyAssignmentRolloutResponse(
            Outputs.FixedOrPercentResponse disruptionBudget,

            string minWaitDuration)
        {
            DisruptionBudget = disruptionBudget;
            MinWaitDuration = minWaitDuration;
        }
    }
}
