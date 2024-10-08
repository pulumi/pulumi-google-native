// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// Stats for the requested dimensions: QPS, duration, and error/warning rate
    /// </summary>
    public sealed class EnterpriseCrmEventbusStatsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dimensions that these stats have been aggregated on.
        /// </summary>
        [Input("dimensions")]
        public Input<Inputs.EnterpriseCrmEventbusStatsDimensionsArgs>? Dimensions { get; set; }

        /// <summary>
        /// Average duration in seconds.
        /// </summary>
        [Input("durationInSeconds")]
        public Input<double>? DurationInSeconds { get; set; }

        /// <summary>
        /// Average error rate.
        /// </summary>
        [Input("errorRate")]
        public Input<double>? ErrorRate { get; set; }

        /// <summary>
        /// Queries per second.
        /// </summary>
        [Input("qps")]
        public Input<double>? Qps { get; set; }

        /// <summary>
        /// Average warning rate.
        /// </summary>
        [Input("warningRate")]
        public Input<double>? WarningRate { get; set; }

        public EnterpriseCrmEventbusStatsArgs()
        {
        }
        public static new EnterpriseCrmEventbusStatsArgs Empty => new EnterpriseCrmEventbusStatsArgs();
    }
}
