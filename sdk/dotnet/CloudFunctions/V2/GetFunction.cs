// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2
{
    public static class GetFunction
    {
        /// <summary>
        /// Returns a function with the given name from the requested project.
        /// </summary>
        public static Task<GetFunctionResult> InvokeAsync(GetFunctionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFunctionResult>("google-native:cloudfunctions/v2:getFunction", args ?? new GetFunctionArgs(), options.WithDefaults());

        /// <summary>
        /// Returns a function with the given name from the requested project.
        /// </summary>
        public static Output<GetFunctionResult> Invoke(GetFunctionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFunctionResult>("google-native:cloudfunctions/v2:getFunction", args ?? new GetFunctionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFunctionArgs : global::Pulumi.InvokeArgs
    {
        [Input("functionId", required: true)]
        public string FunctionId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetFunctionArgs()
        {
        }
        public static new GetFunctionArgs Empty => new GetFunctionArgs();
    }

    public sealed class GetFunctionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("functionId", required: true)]
        public Input<string> FunctionId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetFunctionInvokeArgs()
        {
        }
        public static new GetFunctionInvokeArgs Empty => new GetFunctionInvokeArgs();
    }


    [OutputType]
    public sealed class GetFunctionResult
    {
        /// <summary>
        /// Describes the Build step of the function that builds a container from the given source.
        /// </summary>
        public readonly Outputs.BuildConfigResponse BuildConfig;
        /// <summary>
        /// User-provided description of a function.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Describe whether the function is gen1 or gen2.
        /// </summary>
        public readonly string Environment;
        /// <summary>
        /// An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
        /// </summary>
        public readonly Outputs.EventTriggerResponse EventTrigger;
        /// <summary>
        /// Labels associated with this Cloud Function.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
        /// </summary>
        public readonly Outputs.ServiceConfigResponse ServiceConfig;
        /// <summary>
        /// State of the function.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// State Messages for this Cloud Function.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudFunctionsV2StateMessageResponse> StateMessages;
        /// <summary>
        /// The last update timestamp of a Cloud Function.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetFunctionResult(
            Outputs.BuildConfigResponse buildConfig,

            string description,

            string environment,

            Outputs.EventTriggerResponse eventTrigger,

            ImmutableDictionary<string, string> labels,

            string name,

            Outputs.ServiceConfigResponse serviceConfig,

            string state,

            ImmutableArray<Outputs.GoogleCloudFunctionsV2StateMessageResponse> stateMessages,

            string updateTime)
        {
            BuildConfig = buildConfig;
            Description = description;
            Environment = environment;
            EventTrigger = eventTrigger;
            Labels = labels;
            Name = name;
            ServiceConfig = serviceConfig;
            State = state;
            StateMessages = stateMessages;
            UpdateTime = updateTime;
        }
    }
}