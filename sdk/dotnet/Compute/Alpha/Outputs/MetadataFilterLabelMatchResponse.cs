// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// MetadataFilter label name value pairs that are expected to match corresponding labels presented as metadata to the load balancer.
    /// </summary>
    [OutputType]
    public sealed class MetadataFilterLabelMatchResponse
    {
        /// <summary>
        /// Name of metadata label. The name can have a maximum length of 1024 characters and must be at least 1 character long.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the label must match the specified value. value can have a maximum length of 1024 characters.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private MetadataFilterLabelMatchResponse(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}