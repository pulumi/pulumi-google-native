// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Vmmigration.V1
{
    public static class GetCutoverJob
    {
        /// <summary>
        /// Gets details of a single CutoverJob.
        /// </summary>
        public static Task<GetCutoverJobResult> InvokeAsync(GetCutoverJobArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCutoverJobResult>("google-native:vmmigration/v1:getCutoverJob", args ?? new GetCutoverJobArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single CutoverJob.
        /// </summary>
        public static Output<GetCutoverJobResult> Invoke(GetCutoverJobInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCutoverJobResult>("google-native:vmmigration/v1:getCutoverJob", args ?? new GetCutoverJobInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCutoverJobArgs : global::Pulumi.InvokeArgs
    {
        [Input("cutoverJobId", required: true)]
        public string CutoverJobId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public string MigratingVmId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("sourceId", required: true)]
        public string SourceId { get; set; } = null!;

        public GetCutoverJobArgs()
        {
        }
        public static new GetCutoverJobArgs Empty => new GetCutoverJobArgs();
    }

    public sealed class GetCutoverJobInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("cutoverJobId", required: true)]
        public Input<string> CutoverJobId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("migratingVmId", required: true)]
        public Input<string> MigratingVmId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public GetCutoverJobInvokeArgs()
        {
        }
        public static new GetCutoverJobInvokeArgs Empty => new GetCutoverJobInvokeArgs();
    }


    [OutputType]
    public sealed class GetCutoverJobResult
    {
        /// <summary>
        /// Details of the target VM in Compute Engine.
        /// </summary>
        public readonly Outputs.ComputeEngineTargetDetailsResponse ComputeEngineTargetDetails;
        /// <summary>
        /// The time the cutover job was created (as an API call, not when it was actually created in the target).
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The time the cutover job had finished.
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Provides details for the errors that led to the Cutover Job's state.
        /// </summary>
        public readonly Outputs.StatusResponse Error;
        /// <summary>
        /// The name of the cutover job.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current progress in percentage of the cutover job.
        /// </summary>
        public readonly int ProgressPercent;
        /// <summary>
        /// State of the cutover job.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// A message providing possible extra details about the current state.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The time the state was last updated.
        /// </summary>
        public readonly string StateTime;
        /// <summary>
        /// The cutover steps list representing its progress.
        /// </summary>
        public readonly ImmutableArray<Outputs.CutoverStepResponse> Steps;

        [OutputConstructor]
        private GetCutoverJobResult(
            Outputs.ComputeEngineTargetDetailsResponse computeEngineTargetDetails,

            string createTime,

            string endTime,

            Outputs.StatusResponse error,

            string name,

            int progressPercent,

            string state,

            string stateMessage,

            string stateTime,

            ImmutableArray<Outputs.CutoverStepResponse> steps)
        {
            ComputeEngineTargetDetails = computeEngineTargetDetails;
            CreateTime = createTime;
            EndTime = endTime;
            Error = error;
            Name = name;
            ProgressPercent = progressPercent;
            State = state;
            StateMessage = stateMessage;
            StateTime = stateTime;
            Steps = steps;
        }
    }
}