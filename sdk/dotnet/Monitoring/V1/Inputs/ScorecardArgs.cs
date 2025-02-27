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
    /// A widget showing the latest value of a metric, and how this value relates to one or more thresholds.
    /// </summary>
    public sealed class ScorecardArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Will cause the Scorecard to show only the value, with no indicator to its value relative to its thresholds.
        /// </summary>
        [Input("blankView")]
        public Input<Inputs.EmptyArgs>? BlankView { get; set; }

        /// <summary>
        /// Will cause the scorecard to show a gauge chart.
        /// </summary>
        [Input("gaugeView")]
        public Input<Inputs.GaugeViewArgs>? GaugeView { get; set; }

        /// <summary>
        /// Will cause the scorecard to show a spark chart.
        /// </summary>
        [Input("sparkChartView")]
        public Input<Inputs.SparkChartViewArgs>? SparkChartView { get; set; }

        [Input("thresholds")]
        private InputList<Inputs.ThresholdArgs>? _thresholds;

        /// <summary>
        /// The thresholds used to determine the state of the scorecard given the time series' current value. For an actual value x, the scorecard is in a danger state if x is less than or equal to a danger threshold that triggers below, or greater than or equal to a danger threshold that triggers above. Similarly, if x is above/below a warning threshold that triggers above/below, then the scorecard is in a warning state - unless x also puts it in a danger state. (Danger trumps warning.)As an example, consider a scorecard with the following four thresholds: { value: 90, category: 'DANGER', trigger: 'ABOVE', }, { value: 70, category: 'WARNING', trigger: 'ABOVE', }, { value: 10, category: 'DANGER', trigger: 'BELOW', }, { value: 20, category: 'WARNING', trigger: 'BELOW', } Then: values less than or equal to 10 would put the scorecard in a DANGER state, values greater than 10 but less than or equal to 20 a WARNING state, values strictly between 20 and 70 an OK state, values greater than or equal to 70 but less than 90 a WARNING state, and values greater than or equal to 90 a DANGER state.
        /// </summary>
        public InputList<Inputs.ThresholdArgs> Thresholds
        {
            get => _thresholds ?? (_thresholds = new InputList<Inputs.ThresholdArgs>());
            set => _thresholds = value;
        }

        /// <summary>
        /// Fields for querying time series data from the Stackdriver metrics API.
        /// </summary>
        [Input("timeSeriesQuery", required: true)]
        public Input<Inputs.TimeSeriesQueryArgs> TimeSeriesQuery { get; set; } = null!;

        public ScorecardArgs()
        {
        }
        public static new ScorecardArgs Empty => new ScorecardArgs();
    }
}
