// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1
{
    /// <summary>
    /// Create a Job.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:batch/v1:Job")]
    public partial class Job : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Compute resource allocation for all TaskGroups in the Job.
        /// </summary>
        [Output("allocationPolicy")]
        public Output<Outputs.AllocationPolicyResponse> AllocationPolicy { get; private set; } = null!;

        /// <summary>
        /// When the Job was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// ID used to uniquely identify the Job within its parent scope. This field should contain at most 63 characters and must start with lowercase characters. Only lowercase characters, numbers and '-' are accepted. The '-' character cannot be the first or the last one. A system generated ID will be used if the field is not set. The job.name field in the request will be ignored and the created resource name of the Job will be "{parent}/jobs/{job_id}".
        /// </summary>
        [Output("jobId")]
        public Output<string?> JobId { get; private set; } = null!;

        /// <summary>
        /// Labels for the Job. Labels could be user provided or system generated. For example, "labels": { "department": "finance", "environment": "test" } You can assign up to 64 labels. [Google Compute Engine label restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) apply. Label names that start with "goog-" or "google-" are reserved.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Log preservation policy for the Job.
        /// </summary>
        [Output("logsPolicy")]
        public Output<Outputs.LogsPolicyResponse> LogsPolicy { get; private set; } = null!;

        /// <summary>
        /// Job name. For example: "projects/123456/locations/us-central1/jobs/job01".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Notification configurations.
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableArray<Outputs.JobNotificationResponse>> Notifications { get; private set; } = null!;

        /// <summary>
        /// Priority of the Job. The valid value range is [0, 100). A job with higher priority value is more likely to run earlier if all other requirements are satisfied.
        /// </summary>
        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// Job status. It is read only for users.
        /// </summary>
        [Output("status")]
        public Output<Outputs.JobStatusResponse> Status { get; private set; } = null!;

        /// <summary>
        /// TaskGroups in the Job. Only one TaskGroup is supported now.
        /// </summary>
        [Output("taskGroups")]
        public Output<ImmutableArray<Outputs.TaskGroupResponse>> TaskGroups { get; private set; } = null!;

        /// <summary>
        /// A system generated unique ID (in UUID4 format) for the Job.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The last time the Job was updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("google-native:batch/v1:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:batch/v1:Job", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
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
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Job(name, id, options);
        }
    }

    public sealed class JobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Compute resource allocation for all TaskGroups in the Job.
        /// </summary>
        [Input("allocationPolicy")]
        public Input<Inputs.AllocationPolicyArgs>? AllocationPolicy { get; set; }

        /// <summary>
        /// ID used to uniquely identify the Job within its parent scope. This field should contain at most 63 characters and must start with lowercase characters. Only lowercase characters, numbers and '-' are accepted. The '-' character cannot be the first or the last one. A system generated ID will be used if the field is not set. The job.name field in the request will be ignored and the created resource name of the Job will be "{parent}/jobs/{job_id}".
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels for the Job. Labels could be user provided or system generated. For example, "labels": { "department": "finance", "environment": "test" } You can assign up to 64 labels. [Google Compute Engine label restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) apply. Label names that start with "goog-" or "google-" are reserved.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Log preservation policy for the Job.
        /// </summary>
        [Input("logsPolicy")]
        public Input<Inputs.LogsPolicyArgs>? LogsPolicy { get; set; }

        [Input("notifications")]
        private InputList<Inputs.JobNotificationArgs>? _notifications;

        /// <summary>
        /// Notification configurations.
        /// </summary>
        public InputList<Inputs.JobNotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.JobNotificationArgs>());
            set => _notifications = value;
        }

        /// <summary>
        /// Priority of the Job. The valid value range is [0, 100). A job with higher priority value is more likely to run earlier if all other requirements are satisfied.
        /// </summary>
        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        [Input("taskGroups", required: true)]
        private InputList<Inputs.TaskGroupArgs>? _taskGroups;

        /// <summary>
        /// TaskGroups in the Job. Only one TaskGroup is supported now.
        /// </summary>
        public InputList<Inputs.TaskGroupArgs> TaskGroups
        {
            get => _taskGroups ?? (_taskGroups = new InputList<Inputs.TaskGroupArgs>());
            set => _taskGroups = value;
        }

        public JobArgs()
        {
        }
        public static new JobArgs Empty => new JobArgs();
    }
}