// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    /// <summary>
    /// The set of all usage signals that Data Catalog stores. Note: Usually, these signals are updated daily. In rare cases, an update may fail but will be performed again on the next day.
    /// </summary>
    public sealed class GoogleCloudDatacatalogV1UsageSignalArgs : global::Pulumi.ResourceArgs
    {
        [Input("commonUsageWithinTimeRange")]
        private InputMap<string>? _commonUsageWithinTimeRange;

        /// <summary>
        /// Common usage statistics over each of the predefined time ranges. Supported time ranges are `{"24H", "7D", "30D", "Lifetime"}`.
        /// </summary>
        public InputMap<string> CommonUsageWithinTimeRange
        {
            get => _commonUsageWithinTimeRange ?? (_commonUsageWithinTimeRange = new InputMap<string>());
            set => _commonUsageWithinTimeRange = value;
        }

        /// <summary>
        /// Favorite count in the source system.
        /// </summary>
        [Input("favoriteCount")]
        public Input<string>? FavoriteCount { get; set; }

        /// <summary>
        /// The end timestamp of the duration of usage statistics.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public GoogleCloudDatacatalogV1UsageSignalArgs()
        {
        }
        public static new GoogleCloudDatacatalogV1UsageSignalArgs Empty => new GoogleCloudDatacatalogV1UsageSignalArgs();
    }
}