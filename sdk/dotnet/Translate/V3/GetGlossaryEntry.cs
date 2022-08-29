// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Translate.V3
{
    public static class GetGlossaryEntry
    {
        /// <summary>
        /// Gets a single glossary entry by the given id.
        /// </summary>
        public static Task<GetGlossaryEntryResult> InvokeAsync(GetGlossaryEntryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGlossaryEntryResult>("google-native:translate/v3:getGlossaryEntry", args ?? new GetGlossaryEntryArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a single glossary entry by the given id.
        /// </summary>
        public static Output<GetGlossaryEntryResult> Invoke(GetGlossaryEntryInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGlossaryEntryResult>("google-native:translate/v3:getGlossaryEntry", args ?? new GetGlossaryEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlossaryEntryArgs : global::Pulumi.InvokeArgs
    {
        [Input("glossaryEntryId", required: true)]
        public string GlossaryEntryId { get; set; } = null!;

        [Input("glossaryId", required: true)]
        public string GlossaryId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGlossaryEntryArgs()
        {
        }
        public static new GetGlossaryEntryArgs Empty => new GetGlossaryEntryArgs();
    }

    public sealed class GetGlossaryEntryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("glossaryEntryId", required: true)]
        public Input<string> GlossaryEntryId { get; set; } = null!;

        [Input("glossaryId", required: true)]
        public Input<string> GlossaryId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGlossaryEntryInvokeArgs()
        {
        }
        public static new GetGlossaryEntryInvokeArgs Empty => new GetGlossaryEntryInvokeArgs();
    }


    [OutputType]
    public sealed class GetGlossaryEntryResult
    {
        /// <summary>
        /// Describes the glossary entry.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The resource name of the entry. Format: "projects/*/locations/*/glossaries/*/glossaryEntries/*"
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Used for an unidirectional glossary.
        /// </summary>
        public readonly Outputs.GlossaryTermsPairResponse TermsPair;
        /// <summary>
        /// Used for an equivalent term sets glossary.
        /// </summary>
        public readonly Outputs.GlossaryTermsSetResponse TermsSet;

        [OutputConstructor]
        private GetGlossaryEntryResult(
            string description,

            string name,

            Outputs.GlossaryTermsPairResponse termsPair,

            Outputs.GlossaryTermsSetResponse termsSet)
        {
            Description = description;
            Name = name;
            TermsPair = termsPair;
            TermsSet = termsSet;
        }
    }
}