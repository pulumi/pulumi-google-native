// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Beta.Outputs
{

    /// <summary>
    /// The price range of all variant Product having the same Product.primary_product_id.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2betaPriceInfoPriceRangeResponse
    {
        /// <summary>
        /// The inclusive Product.pricing_info.original_price internal of all variant Product having the same Product.primary_product_id.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2betaIntervalResponse OriginalPrice;
        /// <summary>
        /// The inclusive Product.pricing_info.price interval of all variant Product having the same Product.primary_product_id.
        /// </summary>
        public readonly Outputs.GoogleCloudRetailV2betaIntervalResponse Price;

        [OutputConstructor]
        private GoogleCloudRetailV2betaPriceInfoPriceRangeResponse(
            Outputs.GoogleCloudRetailV2betaIntervalResponse originalPrice,

            Outputs.GoogleCloudRetailV2betaIntervalResponse price)
        {
            OriginalPrice = originalPrice;
            Price = price;
        }
    }
}