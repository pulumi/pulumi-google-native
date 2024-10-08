// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// Statistics and anomalies generated by Model Monitoring.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesResponse
    {
        /// <summary>
        /// Number of anomalies within all stats.
        /// </summary>
        public readonly int AnomalyCount;
        /// <summary>
        /// Deployed Model ID.
        /// </summary>
        public readonly string DeployedModelId;
        /// <summary>
        /// A list of historical Stats and Anomalies generated for all Features.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesFeatureHistoricStatsAnomaliesResponse> FeatureStats;
        /// <summary>
        /// Model Monitoring Objective those stats and anomalies belonging to.
        /// </summary>
        public readonly string Objective;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesResponse(
            int anomalyCount,

            string deployedModelId,

            ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1ModelMonitoringStatsAnomaliesFeatureHistoricStatsAnomaliesResponse> featureStats,

            string objective)
        {
            AnomalyCount = anomalyCount;
            DeployedModelId = deployedModelId;
            FeatureStats = featureStats;
            Objective = objective;
        }
    }
}
