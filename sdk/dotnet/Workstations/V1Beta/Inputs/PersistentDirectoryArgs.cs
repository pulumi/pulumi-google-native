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
    /// A directory to persist across workstation sessions.
    /// </summary>
    public sealed class PersistentDirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A PersistentDirectory backed by a Compute Engine persistent disk.
        /// </summary>
        [Input("gcePd")]
        public Input<Inputs.GceRegionalPersistentDiskArgs>? GcePd { get; set; }

        /// <summary>
        /// Location of this directory in the running workstation.
        /// </summary>
        [Input("mountPath")]
        public Input<string>? MountPath { get; set; }

        public PersistentDirectoryArgs()
        {
        }
        public static new PersistentDirectoryArgs Empty => new PersistentDirectoryArgs();
    }
}