// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    /// <summary>
    /// Scaling based on user-defined schedule. The message describes a single scaling schedule. A scaling schedule changes the minimum number of VM instances an autoscaler can recommend, which can trigger scaling out.
    /// </summary>
    [OutputType]
    public sealed class AutoscalingPolicyScalingScheduleResponse
    {
        /// <summary>
        /// A description of a scaling schedule.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A boolean value that specifies whether a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect. This field is optional, and its value is false by default.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// The duration of time intervals, in seconds, for which this scaling schedule is to run. The minimum allowed value is 300. This field is required.
        /// </summary>
        public readonly int DurationSec;
        /// <summary>
        /// The minimum number of VM instances that the autoscaler will recommend in time intervals starting according to schedule. This field is required.
        /// </summary>
        public readonly int MinRequiredReplicas;
        /// <summary>
        /// The start timestamps of time intervals when this scaling schedule is to provide a scaling signal. This field uses the extended cron format (with an optional year field). The expression can describe a single timestamp if the optional year is set, in which case the scaling schedule runs once. The schedule is interpreted with respect to time_zone. This field is required. Note: These timestamps only describe when autoscaler starts providing the scaling signal. The VMs need additional time to become serving.
        /// </summary>
        public readonly string Schedule;
        /// <summary>
        /// The time zone to use when interpreting the schedule. The value of this field must be a time zone name from the tz database: https://en.wikipedia.org/wiki/Tz_database. This field is assigned a default value of "UTC" if left empty.
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private AutoscalingPolicyScalingScheduleResponse(
            string description,

            bool disabled,

            int durationSec,

            int minRequiredReplicas,

            string schedule,

            string timeZone)
        {
            Description = description;
            Disabled = disabled;
            DurationSec = durationSec;
            MinRequiredReplicas = minRequiredReplicas;
            Schedule = schedule;
            TimeZone = timeZone;
        }
    }
}