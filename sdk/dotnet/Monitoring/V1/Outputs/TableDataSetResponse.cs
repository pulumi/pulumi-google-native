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
    /// Groups a time series query definition with table options.
    /// </summary>
    [OutputType]
    public sealed class TableDataSetResponse
    {
        /// <summary>
        /// Optional. The lower bound on data point frequency for this data set, implemented by specifying the minimum alignment period to use in a time series query For example, if the data is published once every 10 minutes, the min_alignment_period should be at least 10 minutes. It would not make sense to fetch and align data at one minute intervals.
        /// </summary>
        public readonly string MinAlignmentPeriod;
        /// <summary>
        /// Optional. Table display options for configuring how the table is rendered.
        /// </summary>
        public readonly Outputs.TableDisplayOptionsResponse TableDisplayOptions;
        /// <summary>
        /// Optional. A template string for naming TimeSeries in the resulting data set. This should be a string with interpolations of the form ${label_name}, which will resolve to the label's value i.e. "${resource.labels.project_id}."
        /// </summary>
        public readonly string TableTemplate;
        /// <summary>
        /// Fields for querying time series data from the Stackdriver metrics API.
        /// </summary>
        public readonly Outputs.TimeSeriesQueryResponse TimeSeriesQuery;

        [OutputConstructor]
        private TableDataSetResponse(
            string minAlignmentPeriod,

            Outputs.TableDisplayOptionsResponse tableDisplayOptions,

            string tableTemplate,

            Outputs.TimeSeriesQueryResponse timeSeriesQuery)
        {
            MinAlignmentPeriod = minAlignmentPeriod;
            TableDisplayOptions = tableDisplayOptions;
            TableTemplate = tableTemplate;
            TimeSeriesQuery = timeSeriesQuery;
        }
    }
}
