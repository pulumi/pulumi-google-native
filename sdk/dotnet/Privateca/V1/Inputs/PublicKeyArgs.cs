// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Privateca.V1.Inputs
{

    /// <summary>
    /// A PublicKey describes a public key.
    /// </summary>
    public sealed class PublicKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The format of the public key.
        /// </summary>
        [Input("format", required: true)]
        public Input<Pulumi.GoogleNative.Privateca.V1.PublicKeyFormat> Format { get; set; } = null!;

        /// <summary>
        /// A public key. The padding and encoding must match with the `KeyFormat` value specified for the `format` field.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public PublicKeyArgs()
        {
        }
        public static new PublicKeyArgs Empty => new PublicKeyArgs();
    }
}
