// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1
{
    public static class GetTimeSeries
    {
        /// <summary>
        /// Gets a TensorboardTimeSeries.
        /// </summary>
        public static Task<GetTimeSeriesResult> InvokeAsync(GetTimeSeriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTimeSeriesResult>("google-native:aiplatform/v1:getTimeSeries", args ?? new GetTimeSeriesArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a TensorboardTimeSeries.
        /// </summary>
        public static Output<GetTimeSeriesResult> Invoke(GetTimeSeriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTimeSeriesResult>("google-native:aiplatform/v1:getTimeSeries", args ?? new GetTimeSeriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTimeSeriesArgs : global::Pulumi.InvokeArgs
    {
        [Input("experimentId", required: true)]
        public string ExperimentId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("runId", required: true)]
        public string RunId { get; set; } = null!;

        [Input("tensorboardId", required: true)]
        public string TensorboardId { get; set; } = null!;

        [Input("timeSeriesId", required: true)]
        public string TimeSeriesId { get; set; } = null!;

        public GetTimeSeriesArgs()
        {
        }
        public static new GetTimeSeriesArgs Empty => new GetTimeSeriesArgs();
    }

    public sealed class GetTimeSeriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("experimentId", required: true)]
        public Input<string> ExperimentId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("runId", required: true)]
        public Input<string> RunId { get; set; } = null!;

        [Input("tensorboardId", required: true)]
        public Input<string> TensorboardId { get; set; } = null!;

        [Input("timeSeriesId", required: true)]
        public Input<string> TimeSeriesId { get; set; } = null!;

        public GetTimeSeriesInvokeArgs()
        {
        }
        public static new GetTimeSeriesInvokeArgs Empty => new GetTimeSeriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetTimeSeriesResult
    {
        /// <summary>
        /// Timestamp when this TensorboardTimeSeries was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of this TensorboardTimeSeries.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User provided name of this TensorboardTimeSeries. This value should be unique among all TensorboardTimeSeries resources belonging to the same TensorboardRun resource (parent resource).
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Used to perform a consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Scalar, Tensor, or Blob metadata for this TensorboardTimeSeries.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1TensorboardTimeSeriesMetadataResponse Metadata;
        /// <summary>
        /// Name of the TensorboardTimeSeries.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Data of the current plugin, with the size limited to 65KB.
        /// </summary>
        public readonly string PluginData;
        /// <summary>
        /// Immutable. Name of the plugin this time series pertain to. Such as Scalar, Tensor, Blob
        /// </summary>
        public readonly string PluginName;
        /// <summary>
        /// Timestamp when this TensorboardTimeSeries was last updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Immutable. Type of TensorboardTimeSeries value.
        /// </summary>
        public readonly string ValueType;

        [OutputConstructor]
        private GetTimeSeriesResult(
            string createTime,

            string description,

            string displayName,

            string etag,

            Outputs.GoogleCloudAiplatformV1TensorboardTimeSeriesMetadataResponse metadata,

            string name,

            string pluginData,

            string pluginName,

            string updateTime,

            string valueType)
        {
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            Etag = etag;
            Metadata = metadata;
            Name = name;
            PluginData = pluginData;
            PluginName = pluginName;
            UpdateTime = updateTime;
            ValueType = valueType;
        }
    }
}