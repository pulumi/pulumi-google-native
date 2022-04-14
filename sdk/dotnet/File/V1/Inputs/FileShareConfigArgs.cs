// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.File.V1.Inputs
{

    /// <summary>
    /// File share configuration for the instance.
    /// </summary>
    public sealed class FileShareConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// File share capacity in gigabytes (GB). Cloud Filestore defines 1 GB as 1024^3 bytes.
        /// </summary>
        [Input("capacityGb")]
        public Input<string>? CapacityGb { get; set; }

        /// <summary>
        /// The name of the file share (must be 16 characters or less).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nfsExportOptions")]
        private InputList<Inputs.NfsExportOptionsArgs>? _nfsExportOptions;

        /// <summary>
        /// Nfs Export Options. There is a limit of 10 export options per file share.
        /// </summary>
        public InputList<Inputs.NfsExportOptionsArgs> NfsExportOptions
        {
            get => _nfsExportOptions ?? (_nfsExportOptions = new InputList<Inputs.NfsExportOptionsArgs>());
            set => _nfsExportOptions = value;
        }

        /// <summary>
        /// The resource name of the backup, in the format `projects/{project_number}/locations/{location_id}/backups/{backup_id}`, that this file share has been restored from.
        /// </summary>
        [Input("sourceBackup")]
        public Input<string>? SourceBackup { get; set; }

        public FileShareConfigArgs()
        {
        }
    }
}