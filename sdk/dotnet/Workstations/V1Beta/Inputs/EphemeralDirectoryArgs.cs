// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta.Inputs
{

    /// <summary>
    /// An ephemeral directory which won't persist across workstation sessions. It is freshly created on every workstation start operation.
    /// </summary>
    public sealed class EphemeralDirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An EphemeralDirectory backed by a Compute Engine persistent disk.
        /// </summary>
        [Input("gcePd")]
        public Input<Inputs.GcePersistentDiskArgs>? GcePd { get; set; }

        /// <summary>
        /// Location of this directory in the running workstation.
        /// </summary>
        [Input("mountPath", required: true)]
        public Input<string> MountPath { get; set; } = null!;

        public EphemeralDirectoryArgs()
        {
        }
        public static new EphemeralDirectoryArgs Empty => new EphemeralDirectoryArgs();
    }
}