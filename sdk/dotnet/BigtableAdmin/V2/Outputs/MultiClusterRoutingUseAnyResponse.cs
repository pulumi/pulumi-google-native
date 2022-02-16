// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// Read/write requests are routed to the nearest cluster in the instance, and will fail over to the nearest cluster that is available in the event of transient errors or delays. Clusters in a region are considered equidistant. Choosing this option sacrifices read-your-writes consistency to improve availability.
    /// </summary>
    [OutputType]
    public sealed class MultiClusterRoutingUseAnyResponse
    {
        /// <summary>
        /// The set of clusters to route to. The order is ignored; clusters will be tried in order of distance. If left empty, all clusters are eligible.
        /// </summary>
        public readonly ImmutableArray<string> ClusterIds;

        [OutputConstructor]
        private MultiClusterRoutingUseAnyResponse(ImmutableArray<string> clusterIds)
        {
            ClusterIds = clusterIds;
        }
    }
}