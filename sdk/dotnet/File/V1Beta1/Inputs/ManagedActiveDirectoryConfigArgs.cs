// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.File.V1Beta1.Inputs
{

    /// <summary>
    /// ManagedActiveDirectoryConfig contains all the parameters for connecting to Managed Active Directory.
    /// </summary>
    public sealed class ManagedActiveDirectoryConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The computer name is used as a prefix to the mount remote target. Example: if the computer_name is `my-computer`, the mount command will look like: `$mount -o vers=4,sec=krb5 my-computer.filestore.:`.
        /// </summary>
        [Input("computer")]
        public Input<string>? Computer { get; set; }

        /// <summary>
        /// Fully qualified domain name.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        public ManagedActiveDirectoryConfigArgs()
        {
        }
        public static new ManagedActiveDirectoryConfigArgs Empty => new ManagedActiveDirectoryConfigArgs();
    }
}