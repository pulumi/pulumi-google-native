// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1.Outputs
{

    /// <summary>
    /// A condition that applies the override if a line item SKU is found in the SKU group.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudChannelV1SkuGroupConditionResponse
    {
        /// <summary>
        /// Specifies a SKU group (https://cloud.google.com/skus/sku-groups). Resource name of SKU group. Format: accounts/{account}/skuGroups/{sku_group}. Example: "accounts/C01234/skuGroups/3d50fd57-3157-4577-a5a9-a219b8490041".
        /// </summary>
        public readonly string SkuGroup;

        [OutputConstructor]
        private GoogleCloudChannelV1SkuGroupConditionResponse(string skuGroup)
        {
            SkuGroup = skuGroup;
        }
    }
}