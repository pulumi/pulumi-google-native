// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// ParameterMetadataEnumOption specifies the option shown in the enum form.
    /// </summary>
    [OutputType]
    public sealed class ParameterMetadataEnumOptionResponse
    {
        /// <summary>
        /// Optional. The description to display for the enum option.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. The label to display for the enum option.
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// The value of the enum option.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ParameterMetadataEnumOptionResponse(
            string description,

            string label,

            string value)
        {
            Description = description;
            Label = label;
            Value = value;
        }
    }
}