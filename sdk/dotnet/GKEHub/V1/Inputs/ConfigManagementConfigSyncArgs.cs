// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Inputs
{

    /// <summary>
    /// Configuration for Config Sync
    /// </summary>
    public sealed class ConfigManagementConfigSyncArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true to allow the vertical scaling. Defaults to false which disallows vertical scaling. This field is deprecated.
        /// </summary>
        [Input("allowVerticalScale")]
        public Input<bool>? AllowVerticalScale { get; set; }

        /// <summary>
        /// Enables the installation of ConfigSync. If set to true, ConfigSync resources will be created and the other ConfigSync fields will be applied if exist. If set to false, all other ConfigSync fields will be ignored, ConfigSync resources will be deleted. If omitted, ConfigSync resources will be managed depends on the presence of the git or oci field.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Git repo configuration for the cluster.
        /// </summary>
        [Input("git")]
        public Input<Inputs.ConfigManagementGitConfigArgs>? Git { get; set; }

        /// <summary>
        /// The Email of the Google Cloud Service Account (GSA) used for exporting Config Sync metrics to Cloud Monitoring and Cloud Monarch when Workload Identity is enabled. The GSA should have the Monitoring Metric Writer (roles/monitoring.metricWriter) IAM role. The Kubernetes ServiceAccount `default` in the namespace `config-management-monitoring` should be bound to the GSA. This field is required when automatic Feature management is enabled.
        /// </summary>
        [Input("metricsGcpServiceAccountEmail")]
        public Input<string>? MetricsGcpServiceAccountEmail { get; set; }

        /// <summary>
        /// OCI repo configuration for the cluster
        /// </summary>
        [Input("oci")]
        public Input<Inputs.ConfigManagementOciConfigArgs>? Oci { get; set; }

        /// <summary>
        /// Set to true to enable the Config Sync admission webhook to prevent drifts. If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.
        /// </summary>
        [Input("preventDrift")]
        public Input<bool>? PreventDrift { get; set; }

        /// <summary>
        /// Specifies whether the Config Sync Repo is in "hierarchical" or "unstructured" mode.
        /// </summary>
        [Input("sourceFormat")]
        public Input<string>? SourceFormat { get; set; }

        public ConfigManagementConfigSyncArgs()
        {
        }
        public static new ConfigManagementConfigSyncArgs Empty => new ConfigManagementConfigSyncArgs();
    }
}