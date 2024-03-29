// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// AdditionalPodNetworkConfig is the configuration for additional pod networks within the NodeNetworkConfig message
    /// </summary>
    [OutputType]
    public sealed class AdditionalPodNetworkConfigResponse
    {
        /// <summary>
        /// The maximum number of pods per node which use this pod network
        /// </summary>
        public readonly Outputs.MaxPodsConstraintResponse MaxPodsPerNode;
        /// <summary>
        /// The name of the secondary range on the subnet which provides IP address for this pod range
        /// </summary>
        public readonly string SecondaryPodRange;
        /// <summary>
        /// Name of the subnetwork where the additional pod network belongs
        /// </summary>
        public readonly string Subnetwork;

        [OutputConstructor]
        private AdditionalPodNetworkConfigResponse(
            Outputs.MaxPodsConstraintResponse maxPodsPerNode,

            string secondaryPodRange,

            string subnetwork)
        {
            MaxPodsPerNode = maxPodsPerNode;
            SecondaryPodRange = secondaryPodRange;
            Subnetwork = subnetwork;
        }
    }
}
