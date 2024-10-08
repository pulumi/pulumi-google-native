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
    /// Provides compatibility information for various metadata stores.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1EntityCompatibilityStatusResponse
    {
        /// <summary>
        /// Whether this entity is compatible with BigQuery.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusCompatibilityResponse Bigquery;
        /// <summary>
        /// Whether this entity is compatible with Hive Metastore.
        /// </summary>
        public readonly Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusCompatibilityResponse HiveMetastore;

        [OutputConstructor]
        private GoogleCloudDataplexV1EntityCompatibilityStatusResponse(
            Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusCompatibilityResponse bigquery,

            Outputs.GoogleCloudDataplexV1EntityCompatibilityStatusCompatibilityResponse hiveMetastore)
        {
            Bigquery = bigquery;
            HiveMetastore = hiveMetastore;
        }
    }
}
