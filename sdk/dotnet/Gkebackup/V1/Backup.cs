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
    /// Creates a Backup for the given BackupPlan.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:gkebackup/v1:Backup")]
    public partial class Backup : Pulumi.CustomResource
    {
        /// <summary>
        /// If True, all namespaces were included in the Backup.
        /// </summary>
        [Output("allNamespaces")]
        public Output<bool> AllNamespaces { get; private set; } = null!;

        /// <summary>
        /// The client-provided short name for the Backup resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of Backups in this BackupPlan
        /// </summary>
        [Output("backupId")]
        public Output<string?> BackupId { get; private set; } = null!;

        [Output("backupPlanId")]
        public Output<string> BackupPlanId { get; private set; } = null!;

        /// <summary>
        /// Information about the GKE cluster from which this Backup was created.
        /// </summary>
        [Output("clusterMetadata")]
        public Output<Outputs.ClusterMetadataResponse> ClusterMetadata { get; private set; } = null!;

        /// <summary>
        /// Completion time of the Backup
        /// </summary>
        [Output("completeTime")]
        public Output<string> CompleteTime { get; private set; } = null!;

        /// <summary>
        /// The size of the config backup in bytes.
        /// </summary>
        [Output("configBackupSizeBytes")]
        public Output<string> ConfigBackupSizeBytes { get; private set; } = null!;

        /// <summary>
        /// Whether or not the Backup contains Kubernetes Secrets. Controlled by the parent BackupPlan's include_secrets value.
        /// </summary>
        [Output("containsSecrets")]
        public Output<bool> ContainsSecrets { get; private set; } = null!;

        /// <summary>
        /// Whether or not the Backup contains volume data. Controlled by the parent BackupPlan's include_volume_data value.
        /// </summary>
        [Output("containsVolumeData")]
        public Output<bool> ContainsVolumeData { get; private set; } = null!;

        /// <summary>
        /// The timestamp when this Backup resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Minimum age for this Backup (in days). If this field is set to a non-zero value, the Backup will be "locked" against deletion (either manual or automatic deletion) for the number of days provided (measured from the creation time of the Backup). MUST be an integer value between 0-90 (inclusive). Defaults to parent BackupPlan's backup_delete_lock_days setting and may only be increased (either at creation time or in a subsequent update).
        /// </summary>
        [Output("deleteLockDays")]
        public Output<int> DeleteLockDays { get; private set; } = null!;

        /// <summary>
        /// The time at which an existing delete lock will expire for this backup (calculated from create_time + delete_lock_days).
        /// </summary>
        [Output("deleteLockExpireTime")]
        public Output<string> DeleteLockExpireTime { get; private set; } = null!;

        /// <summary>
        /// User specified descriptive string for this Backup.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The customer managed encryption key that was used to encrypt the Backup's artifacts. Inherited from the parent BackupPlan's encryption_key value.
        /// </summary>
        [Output("encryptionKey")]
        public Output<Outputs.EncryptionKeyResponse> EncryptionKey { get; private set; } = null!;

        /// <summary>
        /// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a backup from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform backup updates in order to avoid race conditions: An `etag` is returned in the response to `GetBackup`, and systems are expected to put that etag in the request to `UpdateBackup` or `DeleteBackup` to ensure that their change will be applied to the same version of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A set of custom labels supplied by user.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// This flag indicates whether this Backup resource was created manually by a user or via a schedule in the BackupPlan. A value of True means that the Backup was created manually.
        /// </summary>
        [Output("manual")]
        public Output<bool> Manual { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name of the Backup. projects/*/locations/*/backupPlans/*/backups/*
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The total number of Kubernetes Pods contained in the Backup.
        /// </summary>
        [Output("podCount")]
        public Output<int> PodCount { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The total number of Kubernetes resources included in the Backup.
        /// </summary>
        [Output("resourceCount")]
        public Output<int> ResourceCount { get; private set; } = null!;

        /// <summary>
        /// The age (in days) after which this Backup will be automatically deleted. Must be an integer value &gt;= 0: - If 0, no automatic deletion will occur for this Backup. - If not 0, this must be &gt;= delete_lock_days. Once a Backup is created, this value may only be increased. Defaults to the parent BackupPlan's backup_retain_days value.
        /// </summary>
        [Output("retainDays")]
        public Output<int> RetainDays { get; private set; } = null!;

        /// <summary>
        /// The time at which this Backup will be automatically deleted (calculated from create_time + retain_days).
        /// </summary>
        [Output("retainExpireTime")]
        public Output<string> RetainExpireTime { get; private set; } = null!;

        /// <summary>
        /// If set, the list of ProtectedApplications whose resources were included in the Backup.
        /// </summary>
        [Output("selectedApplications")]
        public Output<Outputs.NamespacedNamesResponse> SelectedApplications { get; private set; } = null!;

        /// <summary>
        /// If set, the list of namespaces that were included in the Backup.
        /// </summary>
        [Output("selectedNamespaces")]
        public Output<Outputs.NamespacesResponse> SelectedNamespaces { get; private set; } = null!;

        /// <summary>
        /// The total size of the Backup in bytes = config backup size + sum(volume backup sizes)
        /// </summary>
        [Output("sizeBytes")]
        public Output<string> SizeBytes { get; private set; } = null!;

        /// <summary>
        /// Current state of the Backup
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Human-readable description of why the backup is in the current `state`.
        /// </summary>
        [Output("stateReason")]
        public Output<string> StateReason { get; private set; } = null!;

        /// <summary>
        /// Server generated global unique identifier of [UUID4](https://en.wikipedia.org/wiki/Universally_unique_identifier)
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The timestamp when this Backup resource was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The total number of volume backups contained in the Backup.
        /// </summary>
        [Output("volumeCount")]
        public Output<int> VolumeCount { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("google-native:gkebackup/v1:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:gkebackup/v1:Backup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, options);
        }
    }

    public sealed class BackupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client-provided short name for the Backup resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of Backups in this BackupPlan
        /// </summary>
        [Input("backupId")]
        public Input<string>? BackupId { get; set; }

        [Input("backupPlanId", required: true)]
        public Input<string> BackupPlanId { get; set; } = null!;

        /// <summary>
        /// Minimum age for this Backup (in days). If this field is set to a non-zero value, the Backup will be "locked" against deletion (either manual or automatic deletion) for the number of days provided (measured from the creation time of the Backup). MUST be an integer value between 0-90 (inclusive). Defaults to parent BackupPlan's backup_delete_lock_days setting and may only be increased (either at creation time or in a subsequent update).
        /// </summary>
        [Input("deleteLockDays")]
        public Input<int>? DeleteLockDays { get; set; }

        /// <summary>
        /// User specified descriptive string for this Backup.
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
        /// The age (in days) after which this Backup will be automatically deleted. Must be an integer value &gt;= 0: - If 0, no automatic deletion will occur for this Backup. - If not 0, this must be &gt;= delete_lock_days. Once a Backup is created, this value may only be increased. Defaults to the parent BackupPlan's backup_retain_days value.
        /// </summary>
        [Input("retainDays")]
        public Input<int>? RetainDays { get; set; }

        public BackupArgs()
        {
        }
    }
}