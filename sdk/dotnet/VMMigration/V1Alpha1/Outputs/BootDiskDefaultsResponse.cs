// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1.Outputs
{

    /// <summary>
    /// BootDiskDefaults hold information about the boot disk of a VM.
    /// </summary>
    [OutputType]
    public sealed class BootDiskDefaultsResponse
    {
        /// <summary>
        /// Optional. Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux operating system running within the instance. If not specified, the server chooses a default device name to apply to this disk, in the form persistent-disk-x, where x is a number assigned by Google Compute Engine. This field is only applicable for persistent disks.
        /// </summary>
        public readonly string DeviceName;
        /// <summary>
        /// Optional. The name of the disk.
        /// </summary>
        public readonly string DiskName;
        /// <summary>
        /// Optional. The type of disk provisioning to use for the VM.
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// Optional. The encryption to apply to the boot disk.
        /// </summary>
        public readonly Outputs.EncryptionResponse Encryption;
        /// <summary>
        /// The image to use when creating the disk.
        /// </summary>
        public readonly Outputs.DiskImageDefaultsResponse Image;

        [OutputConstructor]
        private BootDiskDefaultsResponse(
            string deviceName,

            string diskName,

            string diskType,

            Outputs.EncryptionResponse encryption,

            Outputs.DiskImageDefaultsResponse image)
        {
            DeviceName = deviceName;
            DiskName = diskName;
            DiskType = diskType;
            Encryption = encryption;
            Image = image;
        }
    }
}
