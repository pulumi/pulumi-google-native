// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQueryDataTransfer.V1.Outputs
{

    /// <summary>
    /// Options customizing the data transfer schedule.
    /// </summary>
    [OutputType]
    public sealed class ScheduleOptionsResponse
    {
        /// <summary>
        /// If true, automatic scheduling of data transfer runs for this configuration will be disabled. The runs can be started on ad-hoc basis using StartManualTransferRuns API. When automatic scheduling is disabled, the TransferConfig.schedule field will be ignored.
        /// </summary>
        public readonly bool DisableAutoScheduling;
        /// <summary>
        /// Defines time to stop scheduling transfer runs. A transfer run cannot be scheduled at or after the end time. The end time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Specifies time to start scheduling transfer runs. The first run will be scheduled at or after the start time according to a recurrence pattern defined in the schedule string. The start time can be changed at any moment. The time when a data transfer can be trigerred manually is not limited by this option.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private ScheduleOptionsResponse(
            bool disableAutoScheduling,

            string endTime,

            string startTime)
        {
            DisableAutoScheduling = disableAutoScheduling;
            EndTime = endTime;
            StartTime = startTime;
        }
    }
}
