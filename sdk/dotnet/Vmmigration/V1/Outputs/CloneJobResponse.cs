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
    /// CloneJob describes the process of creating a clone of a MigratingVM to the requested target based on the latest successful uploaded snapshots. While the migration cycles of a MigratingVm take place, it is possible to verify the uploaded VM can be started in the cloud, by creating a clone. The clone can be created without any downtime, and it is created using the latest snapshots which are already in the cloud. The cloneJob is only responsible for its work, not its products, which means once it is finished, it will never touch the instance it created. It will only delete it in case of the CloneJob being cancelled or upon failure to clone.
    /// </summary>
    [OutputType]
    public sealed class CloneJobResponse
    {
        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        public readonly Outputs.ComputeEngineTargetDetailsResponse ComputeEngineTargetDetails;
        /// <summary>
        /// The time the clone job was created (as an API call, not when it was actually created in the target).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time the clone job was ended.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Provides details for the errors that led to the Clone Job's state.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The name of the clone.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// State of the clone job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The time the state was last updated.
        /// </summary>
        public readonly string StateTime;
        /// <summary>
        /// The clone steps list representing its progress.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloneStepResponse> Steps;

        [OutputConstructor]
        private CloneJobResponse(
            Outputs.ComputeEngineTargetDetailsResponse computeEngineTargetDetails,

            string createTime,

            string endTime,

            Outputs.StatusResponse error,

            string name,

            string state,

            string stateTime,

            ImmutableArray<Outputs.CloneStepResponse> steps)
        {
            ComputeEngineTargetDetails = computeEngineTargetDetails;
            CreateTime = createTime;
            EndTime = endTime;
            Error = error;
            Name = name;
            State = state;
            StateTime = stateTime;
            Steps = steps;
        }
    }
}