// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Inputs
{

    /// <summary>
    /// The `RepairRolloutRule` automation rule will automatically repair a failed `Rollout`.
    /// </summary>
    public sealed class RepairRolloutRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the rule. This id must be unique in the `Automation` resource to which this rule belongs. The format is `a-z{0,62}`.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("jobs")]
        private InputList<string>? _jobs;

        /// <summary>
        /// Optional. Jobs to repair. Proceeds only after job name matched any one in the list, or for all jobs if unspecified or empty. The phase that includes the job must match the phase ID specified in `source_phase`. This value must consist of lower-case letters, numbers, and hyphens, start with a letter and end with a letter or a number, and have a max length of 63 characters. In other words, it must match the following regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
        /// </summary>
        public InputList<string> Jobs
        {
            get => _jobs ?? (_jobs = new InputList<string>());
            set => _jobs = value;
        }

        [Input("repairModes", required: true)]
        private InputList<Inputs.RepairModeArgs>? _repairModes;

        /// <summary>
        /// Defines the types of automatic repair actions for failed jobs.
        /// </summary>
        public InputList<Inputs.RepairModeArgs> RepairModes
        {
            get => _repairModes ?? (_repairModes = new InputList<Inputs.RepairModeArgs>());
            set => _repairModes = value;
        }

        [Input("sourcePhases")]
        private InputList<string>? _sourcePhases;

        /// <summary>
        /// Optional. Phases within which jobs are subject to automatic repair actions on failure. Proceeds only after phase name matched any one in the list, or for all phases if unspecified. This value must consist of lower-case letters, numbers, and hyphens, start with a letter and end with a letter or a number, and have a max length of 63 characters. In other words, it must match the following regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
        /// </summary>
        public InputList<string> SourcePhases
        {
            get => _sourcePhases ?? (_sourcePhases = new InputList<string>());
            set => _sourcePhases = value;
        }

        public RepairRolloutRuleArgs()
        {
        }
        public static new RepairRolloutRuleArgs Empty => new RepairRolloutRuleArgs();
    }
}
