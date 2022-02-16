// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    /// <summary>
    /// External table registered by Dataplex. Dataplex publishes data discovered from an asset into multiple other systems (BigQuery, DPMS) in form of tables. We call them "external tables". External tables are also synced into the Data Catalog. This message contains pointers to those external tables (fully qualified name, resource name et cetera) within the Data Catalog.
    /// </summary>
    public sealed class GoogleCloudDatacatalogV1DataplexExternalTableArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Data Catalog entry representing the external table.
        /// </summary>
        [Input("dataCatalogEntry")]
        public Input<string>? DataCatalogEntry { get; set; }

        /// <summary>
        /// Fully qualified name (FQN) of the external table.
        /// </summary>
        [Input("fullyQualifiedName")]
        public Input<string>? FullyQualifiedName { get; set; }

        /// <summary>
        /// Google Cloud resource name of the external table.
        /// </summary>
        [Input("googleCloudResource")]
        public Input<string>? GoogleCloudResource { get; set; }

        /// <summary>
        /// Service in which the external table is registered.
        /// </summary>
        [Input("system")]
        public Input<Pulumi.GoogleNative.DataCatalog.V1.GoogleCloudDatacatalogV1DataplexExternalTableSystem>? System { get; set; }

        public GoogleCloudDatacatalogV1DataplexExternalTableArgs()
        {
        }
    }
}