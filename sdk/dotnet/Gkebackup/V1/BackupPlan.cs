// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1
{
    /// <summary>
    /// Creates a new BackupPlan in a given location.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:gkebackup/v1:BackupPlan")]
    public partial class BackupPlan : Pulumi.CustomResource
    {
        /// <summary>
        /// Defines the configuration of Backups created via this BackupPlan.
        /// </summary>
        [Output("backupConfig")]
        public Output<Outputs.BackupConfigResponse> BackupConfig { get; private set; } = null!;

        /// <summary>
        /// Defines a schedule for automatic Backup creation via this BackupPlan.
        /// </summary>
        [Output("backupSchedule")]
        public Output<Outputs.ScheduleResponse> BackupSchedule { get; private set; } = null!;

        /// <summary>
        /// Immutable. The source cluster from which Backups will be created via this BackupPlan. Valid formats: - projects/*/locations/*/clusters/* - projects/*/zones/*/clusters/*
        /// </summary>
        [Output("cluster")]
        public Output<string> Cluster { get; private set; } = null!;

        /// <summary>
        /// The timestamp when this BackupPlan resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any new Backups from being created via this BackupPlan (including scheduled Backups). Default: False
        /// </summary>
        [Output("deactivated")]
        public Output<bool> Deactivated { get; private set; } = null!;

        /// <summary>
        /// User specified descriptive string for this BackupPlan.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup plan from overwriting each other. It is strongly suggested that systems make use of the 'etag' in the read-modify-write cycle to perform BackupPlan updates in order to avoid race conditions: An `etag` is returned in the response to `GetBackupPlan`, and systems are expected to put that etag in the request to `UpdateBackupPlan` or `DeleteBackupPlan` to ensure that their change will be applied to the same version of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A set of custom labels supplied by user.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// The full name of the BackupPlan resource. Format: projects/*/locations/*/backupPlans/*
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of Kubernetes Pods backed up in the last successful Backup created via this BackupPlan.
        /// </summary>
        [Output("protectedPodCount")]
        public Output<int> ProtectedPodCount { get; private set; } = null!;

        /// <summary>
        /// RetentionPolicy governs lifecycle of Backups created under this plan.
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.RetentionPolicyResponse> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The timestamp when this BackupPlan resource was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a BackupPlan resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupPlan(string name, BackupPlanArgs args, CustomResourceOptions? options = null)
            : base("google-native:gkebackup/v1:BackupPlan", name, args ?? new BackupPlanArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupPlan(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gkebackup/v1:BackupPlan", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BackupPlan resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupPlan Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BackupPlan(name, id, options);
        }
    }

    public sealed class BackupPlanArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines the configuration of Backups created via this BackupPlan.
        /// </summary>
        [Input("backupConfig")]
        public Input<Inputs.BackupConfigArgs>? BackupConfig { get; set; }

        /// <summary>
        /// Required. The client-provided short name for the BackupPlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of BackupPlans in this location
        /// </summary>
        [Input("backupPlanId", required: true)]
        public Input<string> BackupPlanId { get; set; } = null!;

        /// <summary>
        /// Defines a schedule for automatic Backup creation via this BackupPlan.
        /// </summary>
        [Input("backupSchedule")]
        public Input<Inputs.ScheduleArgs>? BackupSchedule { get; set; }

        /// <summary>
        /// Immutable. The source cluster from which Backups will be created via this BackupPlan. Valid formats: - projects/*/locations/*/clusters/* - projects/*/zones/*/clusters/*
        /// </summary>
        [Input("cluster", required: true)]
        public Input<string> Cluster { get; set; } = null!;

        /// <summary>
        /// This flag indicates whether this BackupPlan has been deactivated. Setting this field to True locks the BackupPlan such that no further updates will be allowed (except deletes), including the deactivated field itself. It also prevents any new Backups from being created via this BackupPlan (including scheduled Backups). Default: False
        /// </summary>
        [Input("deactivated")]
        public Input<bool>? Deactivated { get; set; }

        /// <summary>
        /// User specified descriptive string for this BackupPlan.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of custom labels supplied by user.
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
        /// RetentionPolicy governs lifecycle of Backups created under this plan.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.RetentionPolicyArgs>? RetentionPolicy { get; set; }

        public BackupPlanArgs()
        {
        }
    }
}