// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Outputs
{

    /// <summary>
    /// The configuration for exporting to BigQuery.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudHealthcareV1FhirBigQueryDestinationResponse
    {
        /// <summary>
        /// BigQuery URI to an existing dataset, up to 2000 characters long, in the format `bq://projectId.bqDatasetId`.
        /// </summary>
        public readonly string DatasetUri;
        /// <summary>
        /// If this flag is `TRUE`, all tables are deleted from the dataset before the new exported tables are written. If the flag is not set and the destination dataset contains tables, the export call returns an error. If `write_disposition` is specified, this parameter is ignored. force=false is equivalent to write_disposition=WRITE_EMPTY and force=true is equivalent to write_disposition=WRITE_TRUNCATE.
        /// </summary>
        public readonly bool Force;
        /// <summary>
        /// The configuration for the exported BigQuery schema.
        /// </summary>
        public readonly Outputs.SchemaConfigResponse SchemaConfig;
        /// <summary>
        /// Determines if existing data in the destination dataset is overwritten, appended to, or not written if the tables contain data. If a write_disposition is specified, the `force` parameter is ignored.
        /// </summary>
        public readonly string WriteDisposition;

        [OutputConstructor]
        private GoogleCloudHealthcareV1FhirBigQueryDestinationResponse(
            string datasetUri,

            bool force,

            Outputs.SchemaConfigResponse schemaConfig,

            string writeDisposition)
        {
            DatasetUri = datasetUri;
            Force = force;
            SchemaConfig = schemaConfig;
            WriteDisposition = writeDisposition;
        }
    }
}
