// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1.Inputs
{

    /// <summary>
    /// BackupConfig defines the configuration of Backups created via this BackupPlan.
    /// </summary>
    public sealed class BackupConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If True, include all namespaced resources
        /// </summary>
        [Input("allNamespaces")]
        public Input<bool>? AllNamespaces { get; set; }

        /// <summary>
        /// This defines a customer managed encryption key that will be used to encrypt the "config" portion (the Kubernetes resources) of Backups created via this plan. Default (empty): Config backup artifacts will not be encrypted.
        /// </summary>
        [Input("encryptionKey")]
        public Input<Inputs.EncryptionKeyArgs>? EncryptionKey { get; set; }

        /// <summary>
        /// This flag specifies whether Kubernetes Secret resources should be included when they fall into the scope of Backups. Default: False
        /// </summary>
        [Input("includeSecrets")]
        public Input<bool>? IncludeSecrets { get; set; }

        /// <summary>
        /// This flag specifies whether volume data should be backed up when PVCs are included in the scope of a Backup. Default: False
        /// </summary>
        [Input("includeVolumeData")]
        public Input<bool>? IncludeVolumeData { get; set; }

        /// <summary>
        /// If set, include just the resources referenced by the listed ProtectedApplications.
        /// </summary>
        [Input("selectedApplications")]
        public Input<Inputs.NamespacedNamesArgs>? SelectedApplications { get; set; }

        /// <summary>
        /// If set, include just the resources in the listed namespaces.
        /// </summary>
        [Input("selectedNamespaces")]
        public Input<Inputs.NamespacesArgs>? SelectedNamespaces { get; set; }

        public BackupConfigArgs()
        {
        }
    }
}