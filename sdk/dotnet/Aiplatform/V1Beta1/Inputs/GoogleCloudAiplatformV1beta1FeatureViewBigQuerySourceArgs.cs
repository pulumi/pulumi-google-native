// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Inputs
{

    public sealed class GoogleCloudAiplatformV1beta1FeatureViewBigQuerySourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("entityIdColumns", required: true)]
        private InputList<string>? _entityIdColumns;

        /// <summary>
        /// Columns to construct entity_id / row keys. Start by supporting 1 only.
        /// </summary>
        public InputList<string> EntityIdColumns
        {
            get => _entityIdColumns ?? (_entityIdColumns = new InputList<string>());
            set => _entityIdColumns = value;
        }

        /// <summary>
        /// The BigQuery view URI that will be materialized on each sync trigger based on FeatureView.SyncConfig.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public GoogleCloudAiplatformV1beta1FeatureViewBigQuerySourceArgs()
        {
        }
        public static new GoogleCloudAiplatformV1beta1FeatureViewBigQuerySourceArgs Empty => new GoogleCloudAiplatformV1beta1FeatureViewBigQuerySourceArgs();
    }
}