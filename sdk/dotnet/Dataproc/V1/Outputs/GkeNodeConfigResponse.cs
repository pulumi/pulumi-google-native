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
    /// Parameters that describe cluster nodes.
    /// </summary>
    [OutputType]
    public sealed class GkeNodeConfigResponse
    {
        /// <summary>
        /// Optional. A list of hardware accelerators (https://cloud.google.com/compute/docs/gpus) to attach to each node.
        /// </summary>
        public readonly ImmutableArray<Outputs.GkeNodePoolAcceleratorConfigResponse> Accelerators;
        /// <summary>
        /// Optional. The Customer Managed Encryption Key (CMEK) (https://cloud.google.com/kubernetes-engine/docs/how-to/using-cmek) used to encrypt the boot disk attached to each node in the node pool. Specify the key using the following format: projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}
        /// </summary>
        public readonly string BootDiskKmsKey;
        /// <summary>
        /// Optional. The number of local SSD disks to attach to the node, which is limited by the maximum number of disks allowable per zone (see Adding Local SSDs (https://cloud.google.com/compute/docs/disks/local-ssd)).
        /// </summary>
        public readonly int LocalSsdCount;
        /// <summary>
        /// Optional. The name of a Compute Engine machine type (https://cloud.google.com/compute/docs/machine-types).
        /// </summary>
        public readonly string MachineType;
        /// <summary>
        /// Optional. Minimum CPU platform (https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform) to be used by this instance. The instance may be scheduled on the specified or a newer CPU platform. Specify the friendly names of CPU platforms, such as "Intel Haswell"` or Intel Sandy Bridge".
        /// </summary>
        public readonly string MinCpuPlatform;
        /// <summary>
        /// Optional. Whether the nodes are created as legacy preemptible VM instances (https://cloud.google.com/compute/docs/instances/preemptible). Also see Spot VMs, preemptible VM instances without a maximum lifetime. Legacy and Spot preemptible nodes cannot be used in a node pool with the CONTROLLER role or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role).
        /// </summary>
        public readonly bool Preemptible;
        /// <summary>
        /// Optional. Whether the nodes are created as Spot VM instances (https://cloud.google.com/compute/docs/instances/spot). Spot VMs are the latest update to legacy preemptible VMs. Spot VMs do not have a maximum lifetime. Legacy and Spot preemptible nodes cannot be used in a node pool with the CONTROLLER role or in the DEFAULT node pool if the CONTROLLER role is not assigned (the DEFAULT node pool will assume the CONTROLLER role).
        /// </summary>
        public readonly bool Spot;

        [OutputConstructor]
        private GkeNodeConfigResponse(
            ImmutableArray<Outputs.GkeNodePoolAcceleratorConfigResponse> accelerators,

            string bootDiskKmsKey,

            int localSsdCount,

            string machineType,

            string minCpuPlatform,

            bool preemptible,

            bool spot)
        {
            Accelerators = accelerators;
            BootDiskKmsKey = bootDiskKmsKey;
            LocalSsdCount = localSsdCount;
            MachineType = machineType;
            MinCpuPlatform = minCpuPlatform;
            Preemptible = preemptible;
            Spot = spot;
        }
    }
}
