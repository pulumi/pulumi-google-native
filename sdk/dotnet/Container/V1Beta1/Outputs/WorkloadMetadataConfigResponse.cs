// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Outputs
{

    /// <summary>
    /// WorkloadMetadataConfig defines the metadata configuration to expose to workloads on the node pool.
    /// </summary>
    [OutputType]
    public sealed class WorkloadMetadataConfigResponse
    {
        /// <summary>
        /// Mode is the configuration for how to expose metadata to workloads running on the node pool.
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// NodeMetadata is the configuration for how to expose metadata to the workloads running on the node.
        /// </summary>
        public readonly string NodeMetadata;

        [OutputConstructor]
        private WorkloadMetadataConfigResponse(
            string mode,

            string nodeMetadata)
        {
            Mode = mode;
            NodeMetadata = nodeMetadata;
        }
    }
}
