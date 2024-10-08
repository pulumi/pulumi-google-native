// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Auxiliary services configuration for a Cluster.
    /// </summary>
    [OutputType]
    public sealed class AuxiliaryServicesConfigResponse
    {
        /// <summary>
        /// Optional. The Hive Metastore configuration for this workload.
        /// </summary>
        public readonly Outputs.MetastoreConfigResponse MetastoreConfig;
        /// <summary>
        /// Optional. The Spark History Server configuration for the workload.
        /// </summary>
        public readonly Outputs.SparkHistoryServerConfigResponse SparkHistoryServerConfig;

        [OutputConstructor]
        private AuxiliaryServicesConfigResponse(
            Outputs.MetastoreConfigResponse metastoreConfig,

            Outputs.SparkHistoryServerConfigResponse sparkHistoryServerConfig)
        {
            MetastoreConfig = metastoreConfig;
            SparkHistoryServerConfig = sparkHistoryServerConfig;
        }
    }
}
