// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    /// <summary>
    /// The PageOptimizationConfig for model training. This determines how many panels to optimize for, and which serving configs to consider for each panel. The purpose of this model is to optimize which ServingConfig to show on which panels in way that optimizes the visitors shopping journey.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaModelPageOptimizationConfigResponse
    {
        /// <summary>
        /// The type of UserEvent this page optimization is shown for. Each page has an associated event type - this will be the corresponding event type for the page that the page optimization model is used on. Supported types: * `add-to-cart`: Products being added to cart. * `detail-page-view`: Products detail page viewed. * `home-page-view`: Homepage viewed * `category-page-view`: Homepage viewed * `shopping-cart-page-view`: User viewing a shopping cart. `home-page-view` only allows models with type `recommended-for-you`. All other page_optimization_event_type allow all Model.types.
        /// </summary>
        public readonly string PageOptimizationEventType;
        /// <summary>
        /// A list of panel configurations. Limit = 5.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRetailV2alphaModelPageOptimizationConfigPanelResponse> Panels;
        /// <summary>
        /// Optional. How to restrict results across panels e.g. can the same ServingConfig be shown on multiple panels at once. If unspecified, default to `UNIQUE_MODEL_RESTRICTION`.
        /// </summary>
        public readonly string Restriction;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaModelPageOptimizationConfigResponse(
            string pageOptimizationEventType,

            ImmutableArray<Outputs.GoogleCloudRetailV2alphaModelPageOptimizationConfigPanelResponse> panels,

            string restriction)
        {
            PageOptimizationEventType = pageOptimizationEventType;
            Panels = panels;
            Restriction = restriction;
        }
    }
}