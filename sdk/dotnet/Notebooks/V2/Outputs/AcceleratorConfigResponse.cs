// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Notebooks.V2.Outputs
{

    /// <summary>
    /// An accelerator configuration for a VM instance Definition of a hardware accelerator. Note that there is no check on `type` and `core_count` combinations. TPUs are not supported. See [GPUs on Compute Engine](https://cloud.google.com/compute/docs/gpus/#gpus-list) to find a valid combination.
    /// </summary>
    [OutputType]
    public sealed class AcceleratorConfigResponse
    {
        /// <summary>
        /// Optional. Count of cores of this accelerator.
        /// </summary>
        public readonly string CoreCount;
        /// <summary>
        /// Optional. Type of this accelerator.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AcceleratorConfigResponse(
            string coreCount,

            string type)
        {
            CoreCount = coreCount;
            Type = type;
        }
    }
}