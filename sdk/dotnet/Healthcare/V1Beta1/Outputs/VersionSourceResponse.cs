// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// Describes a selector for extracting and matching an MSH field to a value.
    /// </summary>
    [OutputType]
    public sealed class VersionSourceResponse
    {
        /// <summary>
        /// The field to extract from the MSH segment. For example, "3.1" or "18[1].1".
        /// </summary>
        public readonly string MshField;
        /// <summary>
        /// The value to match with the field. For example, "My Application Name" or "2.3".
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private VersionSourceResponse(
            string mshField,

            string value)
        {
            MshField = mshField;
            Value = value;
        }
    }
}