// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Message to be used to configure custom alerting in the {@code EventConfig} protos for an event.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIntegrationsV1alphaIntegrationAlertConfigResponse
    {
        /// <summary>
        /// The period over which the metric value should be aggregated and evaluated. Format is , where integer should be a positive integer and unit should be one of (s,m,h,d,w) meaning (second, minute, hour, day, week). For an EXPECTED_MIN threshold, this aggregation_period must be lesser than 24 hours.
        /// </summary>
        public readonly string AggregationPeriod;
        /// <summary>
        /// For how many contiguous aggregation periods should the expected min or max be violated for the alert to be fired.
        /// </summary>
        public readonly int AlertThreshold;
        /// <summary>
        /// Set to false by default. When set to true, the metrics are not aggregated or pushed to Monarch for this integration alert.
        /// </summary>
        public readonly bool DisableAlert;
        /// <summary>
        /// Name of the alert. This will be displayed in the alert subject. If set, this name should be unique within the scope of the integration.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Should be specified only for *AVERAGE_DURATION and *PERCENTILE_DURATION metrics. This member should be used to specify what duration value the metrics should exceed for the alert to trigger.
        /// </summary>
        public readonly string DurationThreshold;
        /// <summary>
        /// The type of metric.
        /// </summary>
        public readonly string MetricType;
        /// <summary>
        /// For either events or tasks, depending on the type of alert, count only final attempts, not retries.
        /// </summary>
        public readonly bool OnlyFinalAttempt;
        /// <summary>
        /// The threshold type, whether lower(expected_min) or upper(expected_max), for which this alert is being configured. If value falls below expected_min or exceeds expected_max, an alert will be fired.
        /// </summary>
        public readonly string ThresholdType;
        /// <summary>
        /// The metric value, above or below which the alert should be triggered.
        /// </summary>
        public readonly Outputs.GoogleCloudIntegrationsV1alphaIntegrationAlertConfigThresholdValueResponse ThresholdValue;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaIntegrationAlertConfigResponse(
            string aggregationPeriod,

            int alertThreshold,

            bool disableAlert,

            string displayName,

            string durationThreshold,

            string metricType,

            bool onlyFinalAttempt,

            string thresholdType,

            Outputs.GoogleCloudIntegrationsV1alphaIntegrationAlertConfigThresholdValueResponse thresholdValue)
        {
            AggregationPeriod = aggregationPeriod;
            AlertThreshold = alertThreshold;
            DisableAlert = disableAlert;
            DisplayName = displayName;
            DurationThreshold = durationThreshold;
            MetricType = metricType;
            OnlyFinalAttempt = onlyFinalAttempt;
            ThresholdType = thresholdType;
            ThresholdValue = thresholdValue;
        }
    }
}