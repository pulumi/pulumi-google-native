// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    /// <summary>
    /// Fields specific for BigQuery routines.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1BigQueryRoutineSpecResponse
    {
        /// <summary>
        /// Paths of the imported libraries.
        /// </summary>
        public readonly ImmutableArray<string> ImportedLibraries;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1BigQueryRoutineSpecResponse(ImmutableArray<string> importedLibraries)
        {
            ImportedLibraries = importedLibraries;
        }
    }
}
