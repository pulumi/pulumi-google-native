// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// Key-value pair to store extra metadata.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudApigeeV1AttributeResponse
    {
        /// <summary>
        /// API key of the attribute.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Value of the attribute.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GoogleCloudApigeeV1AttributeResponse(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
