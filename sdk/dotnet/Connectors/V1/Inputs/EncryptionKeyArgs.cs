// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Inputs
{

    /// <summary>
    /// Encryption Key value.
    /// </summary>
    public sealed class EncryptionKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The [KMS key name] with which the content of the Operation is encrypted. The expected format: `projects/*/locations/*/keyRings/*/cryptoKeys/*`. Will be empty string if google managed.
        /// </summary>
        [Input("kmsKeyName")]
        public Input<string>? KmsKeyName { get; set; }

        /// <summary>
        /// Type.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Connectors.V1.EncryptionKeyType>? Type { get; set; }

        public EncryptionKeyArgs()
        {
        }
        public static new EncryptionKeyArgs Empty => new EncryptionKeyArgs();
    }
}