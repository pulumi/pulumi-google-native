// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Outputs
{

    /// <summary>
    /// Preview: An identifier for an aggregation function. Aggregation functions are SQL functions that group or transform data from multiple points to a single point. This is a preview feature and may be subject to change before final release.
    /// </summary>
    [OutputType]
    public sealed class AggregationFunctionResponse
    {
        /// <summary>
        /// Optional. Parameters applied to the aggregation function. Only used for functions that require them.
        /// </summary>
        public readonly ImmutableArray<Outputs.ParameterResponse> Parameters;
        /// <summary>
        /// The type of aggregation function, must be one of the following: "none" - no function. "percentile" - APPROX_QUANTILES() - 1 parameter numeric value "average" - AVG() "count" - COUNT() "count-distinct" - COUNT(DISTINCT) "count-distinct-approx" - APPROX_COUNT_DISTINCT() "max" - MAX() "min" - MIN() "sum" - SUM()
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AggregationFunctionResponse(
            ImmutableArray<Outputs.ParameterResponse> parameters,

            string type)
        {
            Parameters = parameters;
            Type = type;
        }
    }
}