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
    /// A patient identifier and associated type.
    /// </summary>
    [OutputType]
    public sealed class PatientIdResponse
    {
        /// <summary>
        /// ID type. For example, MRN or NHS.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The patient's unique identifier.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private PatientIdResponse(
            string type,

            string value)
        {
            Type = type;
            Value = value;
        }
    }
}