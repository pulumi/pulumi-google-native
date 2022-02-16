// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GameServices.V1Beta.Outputs
{

    /// <summary>
    /// Autoscaling config for an Agones fleet.
    /// </summary>
    [OutputType]
    public sealed class ScalingConfigResponse
    {
        /// <summary>
        /// Agones fleet autoscaler spec. Example spec: https://agones.dev/site/docs/reference/fleetautoscaler/
        /// </summary>
        public readonly string FleetAutoscalerSpec;
        /// <summary>
        /// The name of the Scaling Config
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The schedules to which this Scaling Config applies.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScheduleResponse> Schedules;
        /// <summary>
        /// Labels used to identify the game server clusters to which this Agones scaling config applies. A game server cluster is subject to this Agones scaling config if its labels match any of the selector entries.
        /// </summary>
        public readonly ImmutableArray<Outputs.LabelSelectorResponse> Selectors;

        [OutputConstructor]
        private ScalingConfigResponse(
            string fleetAutoscalerSpec,

            string name,

            ImmutableArray<Outputs.ScheduleResponse> schedules,

            ImmutableArray<Outputs.LabelSelectorResponse> selectors)
        {
            FleetAutoscalerSpec = fleetAutoscalerSpec;
            Name = name;
            Schedules = schedules;
            Selectors = selectors;
        }
    }
}