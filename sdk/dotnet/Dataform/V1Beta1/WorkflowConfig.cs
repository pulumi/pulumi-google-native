// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataform.V1Beta1
{
    /// <summary>
    /// Creates a new WorkflowConfig in a given Repository.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:dataform/v1beta1:WorkflowConfig")]
    public partial class WorkflowConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Optional. Optional schedule (in cron format) for automatic execution of this workflow config.
        /// </summary>
        [Output("cronSchedule")]
        public Output<string> CronSchedule { get; private set; } = null!;

        /// <summary>
        /// Optional. If left unset, a default InvocationConfig will be used.
        /// </summary>
        [Output("invocationConfig")]
        public Output<Outputs.InvocationConfigResponse> InvocationConfig { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The workflow config's name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Records of the 10 most recent scheduled execution attempts, ordered in in descending order of `execution_time`. Updated whenever automatic creation of a workflow invocation is triggered by cron_schedule.
        /// </summary>
        [Output("recentScheduledExecutionRecords")]
        public Output<ImmutableArray<Outputs.ScheduledExecutionRecordResponse>> RecentScheduledExecutionRecords { get; private set; } = null!;

        /// <summary>
        /// The name of the release config whose release_compilation_result should be executed. Must be in the format `projects/*/locations/*/repositories/*/releaseConfigs/*`.
        /// </summary>
        [Output("releaseConfig")]
        public Output<string> ReleaseConfig { get; private set; } = null!;

        [Output("repositoryId")]
        public Output<string> RepositoryId { get; private set; } = null!;

        /// <summary>
        /// Optional. Specifies the time zone to be used when interpreting cron_schedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
        /// </summary>
        [Output("timeZone")]
        public Output<string> TimeZone { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for the workflow config, which will become the final component of the workflow config's resource name.
        /// </summary>
        [Output("workflowConfigId")]
        public Output<string> WorkflowConfigId { get; private set; } = null!;


        /// <summary>
        /// Create a WorkflowConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkflowConfig(string name, WorkflowConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:dataform/v1beta1:WorkflowConfig", name, args ?? new WorkflowConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkflowConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:dataform/v1beta1:WorkflowConfig", name, null, MakeResourceOptions(options, id))
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
                    "repositoryId",
                    "workflowConfigId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing WorkflowConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkflowConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkflowConfig(name, id, options);
        }
    }

    public sealed class WorkflowConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Optional schedule (in cron format) for automatic execution of this workflow config.
        /// </summary>
        [Input("cronSchedule")]
        public Input<string>? CronSchedule { get; set; }

        /// <summary>
        /// Optional. If left unset, a default InvocationConfig will be used.
        /// </summary>
        [Input("invocationConfig")]
        public Input<Inputs.InvocationConfigArgs>? InvocationConfig { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the release config whose release_compilation_result should be executed. Must be in the format `projects/*/locations/*/repositories/*/releaseConfigs/*`.
        /// </summary>
        [Input("releaseConfig", required: true)]
        public Input<string> ReleaseConfig { get; set; } = null!;

        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        /// <summary>
        /// Optional. Specifies the time zone to be used when interpreting cron_schedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
        /// </summary>
        [Input("timeZone")]
        public Input<string>? TimeZone { get; set; }

        /// <summary>
        /// Required. The ID to use for the workflow config, which will become the final component of the workflow config's resource name.
        /// </summary>
        [Input("workflowConfigId", required: true)]
        public Input<string> WorkflowConfigId { get; set; } = null!;

        public WorkflowConfigArgs()
        {
        }
        public static new WorkflowConfigArgs Empty => new WorkflowConfigArgs();
    }
}