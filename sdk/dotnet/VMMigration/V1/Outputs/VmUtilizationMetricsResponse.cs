// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1.Outputs
{

    /// <summary>
    /// Utilization metrics values for a single VM.
    /// </summary>
    [OutputType]
    public sealed class VmUtilizationMetricsResponse
    {
        /// <summary>
        /// Average CPU usage, percent.
        /// </summary>
        public readonly int CpuAveragePercent;
        /// <summary>
        /// Max CPU usage, percent.
        /// </summary>
        public readonly int CpuMaxPercent;
        /// <summary>
        /// Average disk IO rate, in kilobytes per second.
        /// </summary>
        public readonly string DiskIoRateAverageKbps;
        /// <summary>
        /// Max disk IO rate, in kilobytes per second.
        /// </summary>
        public readonly string DiskIoRateMaxKbps;
        /// <summary>
        /// Average memory usage, percent.
        /// </summary>
        public readonly int MemoryAveragePercent;
        /// <summary>
        /// Max memory usage, percent.
        /// </summary>
        public readonly int MemoryMaxPercent;
        /// <summary>
        /// Average network throughput (combined transmit-rates and receive-rates), in kilobytes per second.
        /// </summary>
        public readonly string NetworkThroughputAverageKbps;
        /// <summary>
        /// Max network throughput (combined transmit-rates and receive-rates), in kilobytes per second.
        /// </summary>
        public readonly string NetworkThroughputMaxKbps;

        [OutputConstructor]
        private VmUtilizationMetricsResponse(
            int cpuAveragePercent,

            int cpuMaxPercent,

            string diskIoRateAverageKbps,

            string diskIoRateMaxKbps,

            int memoryAveragePercent,

            int memoryMaxPercent,

            string networkThroughputAverageKbps,

            string networkThroughputMaxKbps)
        {
            CpuAveragePercent = cpuAveragePercent;
            CpuMaxPercent = cpuMaxPercent;
            DiskIoRateAverageKbps = diskIoRateAverageKbps;
            DiskIoRateMaxKbps = diskIoRateMaxKbps;
            MemoryAveragePercent = memoryAveragePercent;
            MemoryMaxPercent = memoryMaxPercent;
            NetworkThroughputAverageKbps = networkThroughputAverageKbps;
            NetworkThroughputMaxKbps = networkThroughputMaxKbps;
        }
    }
}
