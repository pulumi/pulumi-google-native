// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2.Inputs
{

    /// <summary>
    /// The specification for personalization.
    /// </summary>
    public sealed class GoogleCloudRetailV2SearchRequestPersonalizationSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to Mode.AUTO.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.Retail.V2.GoogleCloudRetailV2SearchRequestPersonalizationSpecMode>? Mode { get; set; }

        public GoogleCloudRetailV2SearchRequestPersonalizationSpecArgs()
        {
        }
        public static new GoogleCloudRetailV2SearchRequestPersonalizationSpecArgs Empty => new GoogleCloudRetailV2SearchRequestPersonalizationSpecArgs();
    }
}