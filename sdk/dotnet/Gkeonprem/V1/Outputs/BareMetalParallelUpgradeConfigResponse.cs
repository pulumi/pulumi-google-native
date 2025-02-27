// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Outputs
{

    /// <summary>
    /// BareMetalParallelUpgradeConfig defines the parallel upgrade settings for worker node pools.
    /// </summary>
    [OutputType]
    public sealed class BareMetalParallelUpgradeConfigResponse
    {
        /// <summary>
        /// The maximum number of nodes that can be upgraded at once.
        /// </summary>
        public readonly int ConcurrentNodes;
        /// <summary>
        /// The minimum number of nodes that should be healthy and available during an upgrade. If set to the default value of 0, it is possible that none of the nodes will be available during an upgrade.
        /// </summary>
        public readonly int MinimumAvailableNodes;

        [OutputConstructor]
        private BareMetalParallelUpgradeConfigResponse(
            int concurrentNodes,

            int minimumAvailableNodes)
        {
            ConcurrentNodes = concurrentNodes;
            MinimumAvailableNodes = minimumAvailableNodes;
        }
    }
}
