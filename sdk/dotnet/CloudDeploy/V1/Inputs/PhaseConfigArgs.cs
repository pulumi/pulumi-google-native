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
    /// PhaseConfig represents the configuration for a phase in the custom canary deployment.
    /// </summary>
    public sealed class PhaseConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Percentage deployment for the phase.
        /// </summary>
        [Input("percentage", required: true)]
        public Input<int> Percentage { get; set; } = null!;

        /// <summary>
        /// The ID to assign to the `Rollout` phase. This value must consist of lower-case letters, numbers, and hyphens, start with a letter and end with a letter or a number, and have a max length of 63 characters. In other words, it must match the following regex: `^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`.
        /// </summary>
        [Input("phaseId", required: true)]
        public Input<string> PhaseId { get; set; } = null!;

        [Input("profiles")]
        private InputList<string>? _profiles;

        /// <summary>
        /// Skaffold profiles to use when rendering the manifest for this phase. These are in addition to the profiles list specified in the `DeliveryPipeline` stage.
        /// </summary>
        public InputList<string> Profiles
        {
            get => _profiles ?? (_profiles = new InputList<string>());
            set => _profiles = value;
        }

        /// <summary>
        /// Whether to run verify tests after the deployment.
        /// </summary>
        [Input("verify")]
        public Input<bool>? Verify { get; set; }

        public PhaseConfigArgs()
        {
        }
        public static new PhaseConfigArgs Empty => new PhaseConfigArgs();
    }
}