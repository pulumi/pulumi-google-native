// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Vmmigration.V1.Outputs
{

    /// <summary>
    /// ReplicationCycle contains information about the current replication cycle status.
    /// </summary>
    [OutputType]
    public sealed class ReplicationCycleResponse
    {
        /// <summary>
        /// The cycle's ordinal number.
        /// </summary>
        public readonly int CycleNumber;
        /// <summary>
        /// The time the replication cycle has ended.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Provides details on the state of the cycle in case of an error.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The identifier of the ReplicationCycle.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current progress in percentage of this cycle. Was replaced by 'steps' field, which breaks down the cycle progression more accurately.
        /// </summary>
        public readonly int ProgressPercent;
        /// <summary>
        /// The time the replication cycle has started.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// State of the ReplicationCycle.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The cycle's steps list representing its progress.
        /// </summary>
        public readonly ImmutableArray<Outputs.CycleStepResponse> Steps;
        /// <summary>
        /// The accumulated duration the replication cycle was paused.
        /// </summary>
        public readonly string TotalPauseDuration;

        [OutputConstructor]
        private ReplicationCycleResponse(
            int cycleNumber,

            string endTime,

            Outputs.StatusResponse error,

            string name,

            int progressPercent,

            string startTime,

            string state,

            ImmutableArray<Outputs.CycleStepResponse> steps,

            string totalPauseDuration)
        {
            CycleNumber = cycleNumber;
            EndTime = endTime;
            Error = error;
            Name = name;
            ProgressPercent = progressPercent;
            StartTime = startTime;
            State = state;
            Steps = steps;
            TotalPauseDuration = totalPauseDuration;
        }
    }
}