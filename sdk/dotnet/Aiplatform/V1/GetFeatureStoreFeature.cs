// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1
{
    public static class GetFeatureStoreFeature
    {
        /// <summary>
        /// Gets details of a single Feature.
        /// </summary>
        public static Task<GetFeatureStoreFeatureResult> InvokeAsync(GetFeatureStoreFeatureArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeatureStoreFeatureResult>("google-native:aiplatform/v1:getFeatureStoreFeature", args ?? new GetFeatureStoreFeatureArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Feature.
        /// </summary>
        public static Output<GetFeatureStoreFeatureResult> Invoke(GetFeatureStoreFeatureInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeatureStoreFeatureResult>("google-native:aiplatform/v1:getFeatureStoreFeature", args ?? new GetFeatureStoreFeatureInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureStoreFeatureArgs : global::Pulumi.InvokeArgs
    {
        [Input("entityTypeId", required: true)]
        public string EntityTypeId { get; set; } = null!;

        [Input("featureId", required: true)]
        public string FeatureId { get; set; } = null!;

        [Input("featurestoreId", required: true)]
        public string FeaturestoreId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetFeatureStoreFeatureArgs()
        {
        }
        public static new GetFeatureStoreFeatureArgs Empty => new GetFeatureStoreFeatureArgs();
    }

    public sealed class GetFeatureStoreFeatureInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("entityTypeId", required: true)]
        public Input<string> EntityTypeId { get; set; } = null!;

        [Input("featureId", required: true)]
        public Input<string> FeatureId { get; set; } = null!;

        [Input("featurestoreId", required: true)]
        public Input<string> FeaturestoreId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetFeatureStoreFeatureInvokeArgs()
        {
        }
        public static new GetFeatureStoreFeatureInvokeArgs Empty => new GetFeatureStoreFeatureInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeatureStoreFeatureResult
    {
        /// <summary>
        /// Only applicable for Vertex AI Feature Store (Legacy). Timestamp when this EntityType was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Description of the Feature.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. Only applicable for Vertex AI Feature Store (Legacy). If not set, use the monitoring_config defined for the EntityType this Feature belongs to. Only Features with type (Feature.ValueType) BOOL, STRING, DOUBLE or INT64 can enable monitoring. If set to true, all types of data monitoring are disabled despite the config on EntityType.
        /// </summary>
        public readonly bool DisableMonitoring;
        /// <summary>
        /// Used to perform a consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Optional. The labels with user-defined metadata to organize your Features. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information on and examples of labels. No more than 64 user labels can be associated with one Feature (System labels are excluded)." System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Only applicable for Vertex AI Feature Store (Legacy). The list of historical stats and anomalies with specified objectives.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudAiplatformV1FeatureMonitoringStatsAnomalyResponse> MonitoringStatsAnomalies;
        /// <summary>
        /// Immutable. Name of the Feature. Format: `projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entity_type}/features/{feature}` `projects/{project}/locations/{location}/featureGroups/{feature_group}/features/{feature}` The last part feature is assigned by the client. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Only applicable for Vertex AI Feature Store (Legacy). Timestamp when this EntityType was most recently updated.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// Immutable. Only applicable for Vertex AI Feature Store (Legacy). Type of Feature value.
        /// </summary>
        public readonly string ValueType;
        /// <summary>
        /// Only applicable for Vertex AI Feature Store. The name of the BigQuery Table/View columnn hosting data for this version. If no value is provided, will use feature_id.
        /// </summary>
        public readonly string VersionColumnName;

        [OutputConstructor]
        private GetFeatureStoreFeatureResult(
            string createTime,

            string description,

            bool disableMonitoring,

            string etag,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<Outputs.GoogleCloudAiplatformV1FeatureMonitoringStatsAnomalyResponse> monitoringStatsAnomalies,

            string name,

            string updateTime,

            string valueType,

            string versionColumnName)
        {
            CreateTime = createTime;
            Description = description;
            DisableMonitoring = disableMonitoring;
            Etag = etag;
            Labels = labels;
            MonitoringStatsAnomalies = monitoringStatsAnomalies;
            Name = name;
            UpdateTime = updateTime;
            ValueType = valueType;
            VersionColumnName = versionColumnName;
        }
    }
}