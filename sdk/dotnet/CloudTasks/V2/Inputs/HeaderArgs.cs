// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudTasks.V2.Inputs
{

    /// <summary>
    /// Defines a header message. A header can have a key and a value.
    /// </summary>
    public sealed class HeaderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Key of the header.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The Value of the header.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public HeaderArgs()
        {
        }
        public static new HeaderArgs Empty => new HeaderArgs();
    }
}