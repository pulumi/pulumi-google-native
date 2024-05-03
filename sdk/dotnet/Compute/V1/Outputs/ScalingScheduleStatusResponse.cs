// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class ScalingScheduleStatusResponse
    {
        /// <summary>
        /// The last time the scaling schedule became active. Note: this is a timestamp when a schedule actually became active, not when it was planned to do so. The timestamp is in RFC3339 text format.
        /// </summary>
        public readonly string LastStartTime;
        /// <summary>
        /// The next time the scaling schedule is to become active. Note: this is a timestamp when a schedule is planned to run, but the actual time might be slightly different. The timestamp is in RFC3339 text format.
        /// </summary>
        public readonly string NextStartTime;
        /// <summary>
        /// The current state of a scaling schedule.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private ScalingScheduleStatusResponse(
            string lastStartTime,

            string nextStartTime,

            string state)
        {
            LastStartTime = lastStartTime;
            NextStartTime = nextStartTime;
            State = state;
        }
    }
}