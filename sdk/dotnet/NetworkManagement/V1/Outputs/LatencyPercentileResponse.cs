// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1.Outputs
{

    /// <summary>
    /// Latency percentile rank and value.
    /// </summary>
    [OutputType]
    public sealed class LatencyPercentileResponse
    {
        /// <summary>
        /// percent-th percentile of latency observed, in microseconds. Fraction of percent/100 of samples have latency lower or equal to the value of this field.
        /// </summary>
        public readonly string LatencyMicros;
        /// <summary>
        /// Percentage of samples this data point applies to.
        /// </summary>
        public readonly int Percent;

        [OutputConstructor]
        private LatencyPercentileResponse(
            string latencyMicros,

            int percent)
        {
            LatencyMicros = latencyMicros;
            Percent = percent;
        }
    }
}
