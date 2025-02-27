// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Inputs
{

    /// <summary>
    /// Represents the spec of disk options.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1DiskSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size in GB of the boot disk (default is 100GB).
        /// </summary>
        [Input("bootDiskSizeGb")]
        public Input<int>? BootDiskSizeGb { get; set; }

        /// <summary>
        /// Type of the boot disk (default is "pd-ssd"). Valid values: "pd-ssd" (Persistent Disk Solid State Drive) or "pd-standard" (Persistent Disk Hard Disk Drive).
        /// </summary>
        [Input("bootDiskType")]
        public Input<string>? BootDiskType { get; set; }

        public GoogleCloudAiplatformV1DiskSpecArgs()
        {
        }
        public static new GoogleCloudAiplatformV1DiskSpecArgs Empty => new GoogleCloudAiplatformV1DiskSpecArgs();
    }
}
