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
    /// Status of the data scan execution.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataScanExecutionStatusResponse
    {
        /// <summary>
        /// The time when the latest DataScanJob ended.
        /// </summary>
        public readonly string LatestJobEndTime;
        /// <summary>
        /// The time when the latest DataScanJob started.
        /// </summary>
        public readonly string LatestJobStartTime;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataScanExecutionStatusResponse(
            string latestJobEndTime,

            string latestJobStartTime)
        {
            LatestJobEndTime = latestJobEndTime;
            LatestJobStartTime = latestJobStartTime;
        }
    }
}