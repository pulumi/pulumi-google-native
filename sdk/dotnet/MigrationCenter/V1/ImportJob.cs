// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1
{
    /// <summary>
    /// Creates an import job.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:migrationcenter/v1:ImportJob")]
    public partial class ImportJob : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Reference to a source.
        /// </summary>
        [Output("assetSource")]
        public Output<string> AssetSource { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the import job was completed.
        /// </summary>
        [Output("completeTime")]
        public Output<string> CompleteTime { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the import job was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. User-friendly display name. Maximum length is 256 characters.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The report with the results of running the import job.
        /// </summary>
        [Output("executionReport")]
        public Output<Outputs.ExecutionReportResponse> ExecutionReport { get; private set; } = null!;

        /// <summary>
        /// Required. ID of the import job.
        /// </summary>
        [Output("importJobId")]
        public Output<string> ImportJobId { get; private set; } = null!;

        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The full name of the import job.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// The state of the import job.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the import job was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The report with the validation results of the import job.
        /// </summary>
        [Output("validationReport")]
        public Output<Outputs.ValidationReportResponse> ValidationReport { get; private set; } = null!;


        /// <summary>
        /// Create a ImportJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImportJob(string name, ImportJobArgs args, CustomResourceOptions? options = null)
            : base("google-native:migrationcenter/v1:ImportJob", name, args ?? new ImportJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImportJob(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:migrationcenter/v1:ImportJob", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "importJobId",
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ImportJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImportJob Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImportJob(name, id, options);
        }
    }

    public sealed class ImportJobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Reference to a source.
        /// </summary>
        [Input("assetSource", required: true)]
        public Input<string> AssetSource { get; set; } = null!;

        /// <summary>
        /// Optional. User-friendly display name. Maximum length is 256 characters.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Required. ID of the import job.
        /// </summary>
        [Input("importJobId", required: true)]
        public Input<string> ImportJobId { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        public ImportJobArgs()
        {
        }
        public static new ImportJobArgs Empty => new ImportJobArgs();
    }
}