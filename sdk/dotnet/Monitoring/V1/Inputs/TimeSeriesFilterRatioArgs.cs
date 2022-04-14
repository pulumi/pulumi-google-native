// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Inputs
{

    /// <summary>
    /// A pair of time series filters that define a ratio computation. The output time series is the pair-wise division of each aligned element from the numerator and denominator time series.
    /// </summary>
    public sealed class TimeSeriesFilterRatioArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The denominator of the ratio.
        /// </summary>
        [Input("denominator")]
        public Input<Inputs.RatioPartArgs>? Denominator { get; set; }

        /// <summary>
        /// The numerator of the ratio.
        /// </summary>
        [Input("numerator")]
        public Input<Inputs.RatioPartArgs>? Numerator { get; set; }

        /// <summary>
        /// Ranking based time series filter.
        /// </summary>
        [Input("pickTimeSeriesFilter")]
        public Input<Inputs.PickTimeSeriesFilterArgs>? PickTimeSeriesFilter { get; set; }

        /// <summary>
        /// Apply a second aggregation after the ratio is computed.
        /// </summary>
        [Input("secondaryAggregation")]
        public Input<Inputs.AggregationArgs>? SecondaryAggregation { get; set; }

        /// <summary>
        /// Statistics based time series filter. Note: This field is deprecated and completely ignored by the API.
        /// </summary>
        [Input("statisticalTimeSeriesFilter")]
        public Input<Inputs.StatisticalTimeSeriesFilterArgs>? StatisticalTimeSeriesFilter { get; set; }

        public TimeSeriesFilterRatioArgs()
        {
        }
    }
}