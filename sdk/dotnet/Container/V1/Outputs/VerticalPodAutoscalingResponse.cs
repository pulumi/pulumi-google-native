// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// VerticalPodAutoscaling contains global, per-cluster information required by Vertical Pod Autoscaler to automatically adjust the resources of pods controlled by it.
    /// </summary>
    [OutputType]
    public sealed class VerticalPodAutoscalingResponse
    {
        /// <summary>
        /// Enables vertical pod autoscaling.
        /// </summary>
        public readonly bool Enabled;

        [OutputConstructor]
        private VerticalPodAutoscalingResponse(bool enabled)
        {
            Enabled = enabled;
        }
    }
}