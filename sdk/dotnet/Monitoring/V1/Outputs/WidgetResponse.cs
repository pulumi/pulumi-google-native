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
    /// Widget contains a single dashboard component and configuration of how to present the component in the dashboard.
    /// </summary>
    [OutputType]
    public sealed class WidgetResponse
    {
        /// <summary>
        /// A chart of alert policy data.
        /// </summary>
        public readonly Outputs.AlertChartResponse AlertChart;
        /// <summary>
        /// A blank space.
        /// </summary>
        public readonly Outputs.EmptyResponse Blank;
        /// <summary>
        /// A widget that shows a stream of logs.
        /// </summary>
        public readonly Outputs.LogsPanelResponse LogsPanel;
        /// <summary>
        /// A scorecard summarizing time series data.
        /// </summary>
        public readonly Outputs.ScorecardResponse Scorecard;
        /// <summary>
        /// A raw string or markdown displaying textual content.
        /// </summary>
        public readonly Outputs.TextResponse Text;
        /// <summary>
        /// A widget that displays time series data in a tabular format.
        /// </summary>
        public readonly Outputs.TimeSeriesTableResponse TimeSeriesTable;
        /// <summary>
        /// Optional. The title of the widget.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// A chart of time series data.
        /// </summary>
        public readonly Outputs.XyChartResponse XyChart;

        [OutputConstructor]
        private WidgetResponse(
            Outputs.AlertChartResponse alertChart,

            Outputs.EmptyResponse blank,

            Outputs.LogsPanelResponse logsPanel,

            Outputs.ScorecardResponse scorecard,

            Outputs.TextResponse text,

            Outputs.TimeSeriesTableResponse timeSeriesTable,

            string title,

            Outputs.XyChartResponse xyChart)
        {
            AlertChart = alertChart;
            Blank = blank;
            LogsPanel = logsPanel;
            Scorecard = scorecard;
            Text = text;
            TimeSeriesTable = timeSeriesTable;
            Title = title;
            XyChart = xyChart;
        }
    }
}