// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1.Outputs
{

    /// <summary>
    /// A Histogram Chart shows a distribution of values into buckets, showing a count of values which fall into a bucket.
    /// </summary>
    [OutputType]
    public sealed class ReportSummaryHistogramChartDataResponse
    {
        /// <summary>
        /// Buckets in the histogram. There will be `n+1` buckets matching `n` lower bounds in the request. The first bucket will be from -infinity to the first bound. Subsequent buckets will be between one bound and the next. The final bucket will be from the final bound to infinity.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportSummaryHistogramChartDataBucketResponse> Buckets;

        [OutputConstructor]
        private ReportSummaryHistogramChartDataResponse(ImmutableArray<Outputs.ReportSummaryHistogramChartDataBucketResponse> buckets)
        {
            Buckets = buckets;
        }
    }
}
