// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Outputs
{

    /// <summary>
    /// Enum type definition.
    /// </summary>
    [OutputType]
    public sealed class EnumResponse
    {
        /// <summary>
        /// Enum value definitions.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnumValueResponse> Enumvalue;
        /// <summary>
        /// Enum type name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Protocol buffer options.
        /// </summary>
        public readonly ImmutableArray<Outputs.OptionResponse> Options;
        /// <summary>
        /// The source context.
        /// </summary>
        public readonly Outputs.SourceContextResponse SourceContext;
        /// <summary>
        /// The source syntax.
        /// </summary>
        public readonly string Syntax;

        [OutputConstructor]
        private EnumResponse(
            ImmutableArray<Outputs.EnumValueResponse> enumvalue,

            string name,

            ImmutableArray<Outputs.OptionResponse> options,

            Outputs.SourceContextResponse sourceContext,

            string syntax)
        {
            Enumvalue = enumvalue;
            Name = name;
            Options = options;
            SourceContext = sourceContext;
            Syntax = syntax;
        }
    }
}