// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// If set, the detailed data profiles will be persisted to the location of your choice whenever updated.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2ExportResponse
    {
        /// <summary>
        /// Store all table and column profiles in an existing table or a new table in an existing dataset. Each re-generation will result in a new row in BigQuery.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2BigQueryTableResponse ProfileTable;

        [OutputConstructor]
        private GooglePrivacyDlpV2ExportResponse(Outputs.GooglePrivacyDlpV2BigQueryTableResponse profileTable)
        {
            ProfileTable = profileTable;
        }
    }
}