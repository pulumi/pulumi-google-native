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
    /// Dataproc job config.
    /// </summary>
    [OutputType]
    public sealed class JobPlacementResponse
    {
        /// <summary>
        /// Optional. Cluster labels to identify a cluster where the job will be submitted.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ClusterLabels;
        /// <summary>
        /// The name of the cluster where the job will be submitted.
        /// </summary>
        public readonly string ClusterName;
        /// <summary>
        /// A cluster UUID generated by the Dataproc service when the job is submitted.
        /// </summary>
        public readonly string ClusterUuid;

        [OutputConstructor]
        private JobPlacementResponse(
            ImmutableDictionary<string, string> clusterLabels,

            string clusterName,

            string clusterUuid)
        {
            ClusterLabels = clusterLabels;
            ClusterName = clusterName;
            ClusterUuid = clusterUuid;
        }
    }
}
