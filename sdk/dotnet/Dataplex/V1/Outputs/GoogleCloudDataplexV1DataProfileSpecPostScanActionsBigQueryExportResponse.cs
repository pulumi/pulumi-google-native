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
    /// The configuration of BigQuery export post scan action.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataProfileSpecPostScanActionsBigQueryExportResponse
    {
        /// <summary>
        /// Optional. The BigQuery table to export DataProfileScan results to. Format: //bigquery.googleapis.com/projects/PROJECT_ID/datasets/DATASET_ID/tables/TABLE_ID
        /// </summary>
        public readonly string ResultsTable;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataProfileSpecPostScanActionsBigQueryExportResponse(string resultsTable)
        {
            ResultsTable = resultsTable;
        }
    }
}