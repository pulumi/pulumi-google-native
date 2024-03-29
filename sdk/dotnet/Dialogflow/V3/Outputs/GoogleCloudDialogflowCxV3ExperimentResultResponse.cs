// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// The inference result which includes an objective metric to optimize and the confidence interval.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3ExperimentResultResponse
    {
        /// <summary>
        /// The last time the experiment's stats data was updated. Will have default value if stats have never been computed for this experiment.
        /// </summary>
        public readonly string LastUpdateTime;
        /// <summary>
        /// Version variants and metrics.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3ExperimentResultVersionMetricsResponse> VersionMetrics;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3ExperimentResultResponse(
            string lastUpdateTime,

            ImmutableArray<Outputs.GoogleCloudDialogflowCxV3ExperimentResultVersionMetricsResponse> versionMetrics)
        {
            LastUpdateTime = lastUpdateTime;
            VersionMetrics = versionMetrics;
        }
    }
}
