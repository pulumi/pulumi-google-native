// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta.Outputs
{

    /// <summary>
    /// An ephemeral directory which won't persist across workstation sessions. It is freshly created on every workstation start operation.
    /// </summary>
    [OutputType]
    public sealed class EphemeralDirectoryResponse
    {
        /// <summary>
        /// An EphemeralDirectory backed by a Compute Engine persistent disk.
        /// </summary>
        public readonly Outputs.GcePersistentDiskResponse GcePd;
        /// <summary>
        /// Location of this directory in the running workstation.
        /// </summary>
        public readonly string MountPath;

        [OutputConstructor]
        private EphemeralDirectoryResponse(
            Outputs.GcePersistentDiskResponse gcePd,

            string mountPath)
        {
            GcePd = gcePd;
            MountPath = mountPath;
        }
    }
}