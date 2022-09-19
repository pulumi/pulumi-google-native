// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Outputs
{

    [OutputType]
    public sealed class DataMaskingStatisticsResponse
    {
        /// <summary>
        /// [Preview] Whether any accessed data was protected by data masking. The actual evaluation is done by accessStats.masked_field_count &gt; 0. Since this is only used for the discovery_doc generation purpose, as long as the type (boolean) matches, client library can leverage this. The actual evaluation of the variable is done else-where.
        /// </summary>
        public readonly bool DataMaskingApplied;

        [OutputConstructor]
        private DataMaskingStatisticsResponse(bool dataMaskingApplied)
        {
            DataMaskingApplied = dataMaskingApplied;
        }
    }
}