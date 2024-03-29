// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// Batch compute resources associated with the task.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1TaskInfrastructureSpecBatchComputeResourcesResponse
    {
        /// <summary>
        /// Optional. Total number of job executors. Executor Count should be between 2 and 100. Default=2
        /// </summary>
        public readonly int ExecutorsCount;
        /// <summary>
        /// Optional. Max configurable executors. If max_executors_count &gt; executors_count, then auto-scaling is enabled. Max Executor Count should be between 2 and 1000. Default=1000
        /// </summary>
        public readonly int MaxExecutorsCount;

        [OutputConstructor]
        private GoogleCloudDataplexV1TaskInfrastructureSpecBatchComputeResourcesResponse(
            int executorsCount,

            int maxExecutorsCount)
        {
            ExecutorsCount = executorsCount;
            MaxExecutorsCount = maxExecutorsCount;
        }
    }
}
