// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Outputs
{

    /// <summary>
    /// Dataset template used for dynamic dataset creation.
    /// </summary>
    [OutputType]
    public sealed class DatasetTemplateResponse
    {
        /// <summary>
        /// If supplied, every created dataset will have its name prefixed by the provided value. The prefix and name will be separated by an underscore. i.e. _.
        /// </summary>
        public readonly string DatasetIdPrefix;
        /// <summary>
        /// Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table. The BigQuery Service Account associated with your project requires access to this encryption key. i.e. projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{cryptoKey}. See https://cloud.google.com/bigquery/docs/customer-managed-encryption for more information.
        /// </summary>
        public readonly string KmsKeyName;
        /// <summary>
        /// The geographic location where the dataset should reside. See https://cloud.google.com/bigquery/docs/locations for supported locations.
        /// </summary>
        public readonly string Location;

        [OutputConstructor]
        private DatasetTemplateResponse(
            string datasetIdPrefix,

            string kmsKeyName,

            string location)
        {
            DatasetIdPrefix = datasetIdPrefix;
            KmsKeyName = kmsKeyName;
            Location = location;
        }
    }
}
