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
    /// CycleStep holds information about a step progress.
    /// </summary>
    [OutputType]
    public sealed class CycleStepResponse
    {
        /// <summary>
        /// The time the cycle step has ended.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Initializing replication step.
        /// </summary>
        public readonly Outputs.InitializingReplicationStepResponse InitializingReplication;
        /// <summary>
        /// Post processing step.
        /// </summary>
        public readonly Outputs.PostProcessingStepResponse PostProcessing;
        /// <summary>
        /// Replicating step.
        /// </summary>
        public readonly Outputs.ReplicatingStepResponse Replicating;
        /// <summary>
        /// The time the cycle step has started.
        /// </summary>
        public readonly string StartTime;

        [OutputConstructor]
        private CycleStepResponse(
            string endTime,

            Outputs.InitializingReplicationStepResponse initializingReplication,

            Outputs.PostProcessingStepResponse postProcessing,

            Outputs.ReplicatingStepResponse replicating,

            string startTime)
        {
            EndTime = endTime;
            InitializingReplication = initializingReplication;
            PostProcessing = postProcessing;
            Replicating = replicating;
            StartTime = startTime;
        }
    }
}