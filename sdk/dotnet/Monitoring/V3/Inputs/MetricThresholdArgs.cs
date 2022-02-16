// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// A condition type that compares a collection of time series against a threshold.
    /// </summary>
    public sealed class MetricThresholdArgs : Pulumi.ResourceArgs
    {
        [Input("aggregations")]
        private InputList<Inputs.AggregationArgs>? _aggregations;

        /// <summary>
        /// Specifies the alignment of data points in individual time series as well as how to combine the retrieved time series together (such as when aggregating multiple streams on each resource to a single stream for each resource or when aggregating streams across all members of a group of resources). Multiple aggregations are applied in the order specified.This field is similar to the one in the ListTimeSeries request (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list). It is advisable to use the ListTimeSeries method when debugging this field.
        /// </summary>
        public InputList<Inputs.AggregationArgs> Aggregations
        {
            get => _aggregations ?? (_aggregations = new InputList<Inputs.AggregationArgs>());
            set => _aggregations = value;
        }

        /// <summary>
        /// The comparison to apply between the time series (indicated by filter and aggregation) and the threshold (indicated by threshold_value). The comparison is applied on each time series, with the time series on the left-hand side and the threshold on the right-hand side.Only COMPARISON_LT and COMPARISON_GT are supported currently.
        /// </summary>
        [Input("comparison")]
        public Input<Pulumi.GoogleNative.Monitoring.V3.MetricThresholdComparison>? Comparison { get; set; }

        [Input("denominatorAggregations")]
        private InputList<Inputs.AggregationArgs>? _denominatorAggregations;

        /// <summary>
        /// Specifies the alignment of data points in individual time series selected by denominatorFilter as well as how to combine the retrieved time series together (such as when aggregating multiple streams on each resource to a single stream for each resource or when aggregating streams across all members of a group of resources).When computing ratios, the aggregations and denominator_aggregations fields must use the same alignment period and produce time series that have the same periodicity and labels.
        /// </summary>
        public InputList<Inputs.AggregationArgs> DenominatorAggregations
        {
            get => _denominatorAggregations ?? (_denominatorAggregations = new InputList<Inputs.AggregationArgs>());
            set => _denominatorAggregations = value;
        }

        /// <summary>
        /// A filter (https://cloud.google.com/monitoring/api/v3/filters) that identifies a time series that should be used as the denominator of a ratio that will be compared with the threshold. If a denominator_filter is specified, the time series specified by the filter field will be used as the numerator.The filter must specify the metric type and optionally may contain restrictions on resource type, resource labels, and metric labels. This field may not exceed 2048 Unicode characters in length.
        /// </summary>
        [Input("denominatorFilter")]
        public Input<string>? DenominatorFilter { get; set; }

        /// <summary>
        /// The amount of time that a time series must violate the threshold to be considered failing. Currently, only values that are a multiple of a minute--e.g., 0, 60, 120, or 300 seconds--are supported. If an invalid value is given, an error will be returned. When choosing a duration, it is useful to keep in mind the frequency of the underlying time series data (which may also be affected by any alignments specified in the aggregations field); a good duration is long enough so that a single outlier does not generate spurious alerts, but short enough that unhealthy states are detected and alerted on quickly.
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// A condition control that determines how metric-threshold conditions are evaluated when data stops arriving.
        /// </summary>
        [Input("evaluationMissingData")]
        public Input<Pulumi.GoogleNative.Monitoring.V3.MetricThresholdEvaluationMissingData>? EvaluationMissingData { get; set; }

        /// <summary>
        /// A filter (https://cloud.google.com/monitoring/api/v3/filters) that identifies which time series should be compared with the threshold.The filter is similar to the one that is specified in the ListTimeSeries request (https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.timeSeries/list) (that call is useful to verify the time series that will be retrieved / processed). The filter must specify the metric type and the resource type. Optionally, it can specify resource labels and metric labels. This field must not exceed 2048 Unicode characters in length.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// A value against which to compare the time series.
        /// </summary>
        [Input("thresholdValue")]
        public Input<double>? ThresholdValue { get; set; }

        /// <summary>
        /// The number/percent of time series for which the comparison must hold in order for the condition to trigger. If unspecified, then the condition will trigger if the comparison is true for any of the time series that have been identified by filter and aggregations, or by the ratio, if denominator_filter and denominator_aggregations are specified.
        /// </summary>
        [Input("trigger")]
        public Input<Inputs.TriggerArgs>? Trigger { get; set; }

        public MetricThresholdArgs()
        {
        }
    }
}