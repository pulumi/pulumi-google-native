// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1Beta1
{
    public static class GetTaxonomy
    {
        /// <summary>
        /// Gets a taxonomy.
        /// </summary>
        public static Task<GetTaxonomyResult> InvokeAsync(GetTaxonomyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTaxonomyResult>("google-native:datacatalog/v1beta1:getTaxonomy", args ?? new GetTaxonomyArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a taxonomy.
        /// </summary>
        public static Output<GetTaxonomyResult> Invoke(GetTaxonomyInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetTaxonomyResult>("google-native:datacatalog/v1beta1:getTaxonomy", args ?? new GetTaxonomyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTaxonomyArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("taxonomyId", required: true)]
        public string TaxonomyId { get; set; } = null!;

        public GetTaxonomyArgs()
        {
        }
    }

    public sealed class GetTaxonomyInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("taxonomyId", required: true)]
        public Input<string> TaxonomyId { get; set; } = null!;

        public GetTaxonomyInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTaxonomyResult
    {
        /// <summary>
        /// Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        /// </summary>
        public readonly ImmutableArray<string> ActivatedPolicyTypes;
        /// <summary>
        /// Optional. Description of this taxonomy. It must: contain only unicode characters, tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes long when encoded in UTF-8. If not set, defaults to an empty description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User defined name of this taxonomy. It must: contain only unicode letters, numbers, underscores, dashes and spaces; not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8. The taxonomy display name must be unique within an organization.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name of this taxonomy, whose format is: "projects/{project_number}/locations/{location_id}/taxonomies/{id}".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of policy tags contained in this taxonomy.
        /// </summary>
        public readonly int PolicyTagCount;
        /// <summary>
        /// Timestamps about this taxonomy. Only create_time and update_time are used.
        /// </summary>
        public readonly Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse TaxonomyTimestamps;

        [OutputConstructor]
        private GetTaxonomyResult(
            ImmutableArray<string> activatedPolicyTypes,

            string description,

            string displayName,

            string name,

            int policyTagCount,

            Outputs.GoogleCloudDatacatalogV1beta1SystemTimestampsResponse taxonomyTimestamps)
        {
            ActivatedPolicyTypes = activatedPolicyTypes;
            Description = description;
            DisplayName = displayName;
            Name = name;
            PolicyTagCount = policyTagCount;
            TaxonomyTimestamps = taxonomyTimestamps;
        }
    }
}