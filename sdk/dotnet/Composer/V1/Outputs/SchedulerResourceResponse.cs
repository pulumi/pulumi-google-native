// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1.Outputs
{

    /// <summary>
    /// Configuration for resources used by Airflow schedulers.
    /// </summary>
    [OutputType]
    public sealed class SchedulerResourceResponse
    {
        /// <summary>
        /// Optional. The number of schedulers.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// Optional. CPU request and limit for a single Airflow scheduler replica.
        /// </summary>
        public readonly double Cpu;
        /// <summary>
        /// Optional. Memory (GB) request and limit for a single Airflow scheduler replica.
        /// </summary>
        public readonly double MemoryGb;
        /// <summary>
        /// Optional. Storage (GB) request and limit for a single Airflow scheduler replica.
        /// </summary>
        public readonly double StorageGb;

        [OutputConstructor]
        private SchedulerResourceResponse(
            int count,

            double cpu,

            double memoryGb,

            double storageGb)
        {
            Count = count;
            Cpu = cpu;
            MemoryGb = memoryGb;
            StorageGb = storageGb;
        }
    }
}