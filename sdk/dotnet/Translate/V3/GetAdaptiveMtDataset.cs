// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3
{
    public static class GetAdaptiveMtDataset
    {
        /// <summary>
        /// Gets the Adaptive MT dataset.
        /// </summary>
        public static Task<GetAdaptiveMtDatasetResult> InvokeAsync(GetAdaptiveMtDatasetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAdaptiveMtDatasetResult>("google-native:translate/v3:getAdaptiveMtDataset", args ?? new GetAdaptiveMtDatasetArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the Adaptive MT dataset.
        /// </summary>
        public static Output<GetAdaptiveMtDatasetResult> Invoke(GetAdaptiveMtDatasetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAdaptiveMtDatasetResult>("google-native:translate/v3:getAdaptiveMtDataset", args ?? new GetAdaptiveMtDatasetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAdaptiveMtDatasetArgs : global::Pulumi.InvokeArgs
    {
        [Input("adaptiveMtDatasetId", required: true)]
        public string AdaptiveMtDatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAdaptiveMtDatasetArgs()
        {
        }
        public static new GetAdaptiveMtDatasetArgs Empty => new GetAdaptiveMtDatasetArgs();
    }

    public sealed class GetAdaptiveMtDatasetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("adaptiveMtDatasetId", required: true)]
        public Input<string> AdaptiveMtDatasetId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAdaptiveMtDatasetInvokeArgs()
        {
        }
        public static new GetAdaptiveMtDatasetInvokeArgs Empty => new GetAdaptiveMtDatasetInvokeArgs();
    }


    [OutputType]
    public sealed class GetAdaptiveMtDatasetResult
    {
        /// <summary>
        /// Timestamp when this dataset was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The name of the dataset to show in the interface. The name can be up to 32 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscores (_), and ASCII digits 0-9.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The number of examples in the dataset.
        /// </summary>
        public readonly int ExampleCount;
        /// <summary>
        /// The resource name of the dataset, in form of `projects/{project-number-or-id}/locations/{location_id}/adaptiveMtDatasets/{dataset_id}`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The BCP-47 language code of the source language.
        /// </summary>
        public readonly string SourceLanguageCode;
        /// <summary>
        /// The BCP-47 language code of the target language.
        /// </summary>
        public readonly string TargetLanguageCode;
        /// <summary>
        /// Timestamp when this dataset was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAdaptiveMtDatasetResult(
            string createTime,

            string displayName,

            int exampleCount,

            string name,

            string sourceLanguageCode,

            string targetLanguageCode,

            string updateTime)
        {
            CreateTime = createTime;
            DisplayName = displayName;
            ExampleCount = exampleCount;
            Name = name;
            SourceLanguageCode = sourceLanguageCode;
            TargetLanguageCode = targetLanguageCode;
            UpdateTime = updateTime;
        }
    }
}
