// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    /// <summary>
    /// Common statistics on the entry's usage. They can be set on any system.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1CommonUsageStatsResponse
    {
        /// <summary>
        /// View count in source system.
        /// </summary>
        public readonly string ViewCount;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1CommonUsageStatsResponse(string viewCount)
        {
            ViewCount = viewCount;
        }
    }
}