// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2
{
    public static class GetServingConfig
    {
        /// <summary>
        /// Gets a ServingConfig. Returns a NotFound error if the ServingConfig does not exist.
        /// </summary>
        public static Task<GetServingConfigResult> InvokeAsync(GetServingConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetServingConfigResult>("google-native:retail/v2:getServingConfig", args ?? new GetServingConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a ServingConfig. Returns a NotFound error if the ServingConfig does not exist.
        /// </summary>
        public static Output<GetServingConfigResult> Invoke(GetServingConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetServingConfigResult>("google-native:retail/v2:getServingConfig", args ?? new GetServingConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetServingConfigArgs : global::Pulumi.InvokeArgs
    {
        [Input("catalogId", required: true)]
        public string CatalogId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("servingConfigId", required: true)]
        public string ServingConfigId { get; set; } = null!;

        public GetServingConfigArgs()
        {
        }
        public static new GetServingConfigArgs Empty => new GetServingConfigArgs();
    }

    public sealed class GetServingConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("servingConfigId", required: true)]
        public Input<string> ServingConfigId { get; set; } = null!;

        public GetServingConfigInvokeArgs()
        {
        }
        public static new GetServingConfigInvokeArgs Empty => new GetServingConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetServingConfigResult
    {
        /// <summary>
        /// Condition boost specifications. If a product matches multiple conditions in the specifications, boost scores from these specifications are all applied and combined in a non-linear way. Maximum number of specifications is 100. Notice that if both ServingConfig.boost_control_ids and SearchRequest.boost_spec are set, the boost conditions from both places are evaluated. If a search request matches multiple boost conditions, the final boost score is equal to the sum of the boost scores from all matched boost conditions. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> BoostControlIds;
        /// <summary>
        /// The human readable serving config display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// How much diversity to use in recommendation model results e.g. `medium-diversity` or `high-diversity`. Currently supported values: * `no-diversity` * `low-diversity` * `medium-diversity` * `high-diversity` * `auto-diversity` If not specified, we choose default based on recommendation model type. Default value: `no-diversity`. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        public readonly string DiversityLevel;
        /// <summary>
        /// What kind of diversity to use - data driven or rule based. If unset, the server behavior defaults to RULE_BASED_DIVERSITY.
        /// </summary>
        public readonly string DiversityType;
        /// <summary>
        /// Condition do not associate specifications. If multiple do not associate conditions match, all matching do not associate controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> DoNotAssociateControlIds;
        /// <summary>
        /// The specification for dynamically generated facets. Notice that only textual facets can be dynamically generated. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2SearchRequestDynamicFacetSpecResponse DynamicFacetSpec;
        /// <summary>
        /// Whether to add additional category filters on the `similar-items` model. If not specified, we enable it by default. Allowed values are: * `no-category-match`: No additional filtering of original results from the model and the customer's filters. * `relaxed-category-match`: Only keep results with categories that match at least one item categories in the PredictRequests's context item. * If customer also sends filters in the PredictRequest, then the results will satisfy both conditions (user given and category match). Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        public readonly string EnableCategoryFilterLevel;
        /// <summary>
        /// Facet specifications for faceted search. If empty, no facets are returned. The ids refer to the ids of Control resources with only the Facet control set. These controls are assumed to be in the same Catalog as the ServingConfig. A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error is returned. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> FacetControlIds;
        /// <summary>
        /// Condition filter specifications. If a product matches multiple conditions in the specifications, filters from these specifications are all applied and combined via the AND operator. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> FilterControlIds;
        /// <summary>
        /// Condition ignore specifications. If multiple ignore conditions match, all matching ignore controls in the list will execute. - Order does not matter. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> IgnoreControlIds;
        /// <summary>
        /// The id of the model in the same Catalog to use at serving time. Currently only RecommendationModels are supported: https://cloud.google.com/retail/recommendations-ai/docs/create-models Can be changed but only to a compatible model (e.g. others-you-may-like CTR to others-you-may-like CVR). Required when solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        public readonly string ModelId;
        /// <summary>
        /// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/servingConfig/*`
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Condition oneway synonyms specifications. If multiple oneway synonyms conditions match, all matching oneway synonyms controls in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> OnewaySynonymsControlIds;
        /// <summary>
        /// The specification for personalization spec. Can only be set if solution_types is SOLUTION_TYPE_SEARCH. Notice that if both ServingConfig.personalization_spec and SearchRequest.personalization_spec are set. SearchRequest.personalization_spec will override ServingConfig.personalization_spec.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2SearchRequestPersonalizationSpecResponse PersonalizationSpec;
        /// <summary>
        /// How much price ranking we want in serving results. Price reranking causes product items with a similar recommendation probability to be ordered by price, with the highest-priced items first. This setting could result in a decrease in click-through and conversion rates. Allowed values are: * `no-price-reranking` * `low-price-reranking` * `medium-price-reranking` * `high-price-reranking` If not specified, we choose default based on model type. Default value: `no-price-reranking`. Can only be set if solution_types is SOLUTION_TYPE_RECOMMENDATION.
        /// </summary>
        public readonly string PriceRerankingLevel;
        /// <summary>
        /// Condition redirect specifications. Only the first triggered redirect action is applied, even if multiple apply. Maximum number of specifications is 1000. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> RedirectControlIds;
        /// <summary>
        /// Condition replacement specifications. - Applied according to the order in the list. - A previously replaced term can not be re-replaced. - Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> ReplacementControlIds;
        /// <summary>
        /// Immutable. Specifies the solution types that a serving config can be associated with. Currently we support setting only one type of solution.
        /// </summary>
        public readonly ImmutableArray<string> SolutionTypes;
        /// <summary>
        /// Condition synonyms specifications. If multiple syonyms conditions match, all matching synonyms control in the list will execute. Order of controls in the list will not matter. Maximum number of specifications is 100. Can only be set if solution_types is SOLUTION_TYPE_SEARCH.
        /// </summary>
        public readonly ImmutableArray<string> TwowaySynonymsControlIds;

        [OutputConstructor]
        private GetServingConfigResult(
            ImmutableArray<string> boostControlIds,

            string displayName,

            string diversityLevel,

            string diversityType,

            ImmutableArray<string> doNotAssociateControlIds,

            Outputs.GoogleCloudRetailV2SearchRequestDynamicFacetSpecResponse dynamicFacetSpec,

            string enableCategoryFilterLevel,

            ImmutableArray<string> facetControlIds,

            ImmutableArray<string> filterControlIds,

            ImmutableArray<string> ignoreControlIds,

            string modelId,

            string name,

            ImmutableArray<string> onewaySynonymsControlIds,

            Outputs.GoogleCloudRetailV2SearchRequestPersonalizationSpecResponse personalizationSpec,

            string priceRerankingLevel,

            ImmutableArray<string> redirectControlIds,

            ImmutableArray<string> replacementControlIds,

            ImmutableArray<string> solutionTypes,

            ImmutableArray<string> twowaySynonymsControlIds)
        {
            BoostControlIds = boostControlIds;
            DisplayName = displayName;
            DiversityLevel = diversityLevel;
            DiversityType = diversityType;
            DoNotAssociateControlIds = doNotAssociateControlIds;
            DynamicFacetSpec = dynamicFacetSpec;
            EnableCategoryFilterLevel = enableCategoryFilterLevel;
            FacetControlIds = facetControlIds;
            FilterControlIds = filterControlIds;
            IgnoreControlIds = ignoreControlIds;
            ModelId = modelId;
            Name = name;
            OnewaySynonymsControlIds = onewaySynonymsControlIds;
            PersonalizationSpec = personalizationSpec;
            PriceRerankingLevel = priceRerankingLevel;
            RedirectControlIds = redirectControlIds;
            ReplacementControlIds = replacementControlIds;
            SolutionTypes = solutionTypes;
            TwowaySynonymsControlIds = twowaySynonymsControlIds;
        }
    }
}