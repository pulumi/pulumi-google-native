// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// Either an InstancePolicy or an instance template.
    /// </summary>
    [OutputType]
    public sealed class InstancePolicyOrTemplateResponse
    {
        /// <summary>
        /// Set this field true if users want Batch to help fetch drivers from a third party location and install them for GPUs specified in policy.accelerators or instance_template on their behalf. Default is false.
        /// </summary>
        public readonly bool InstallGpuDrivers;
        /// <summary>
        /// Name of an instance template used to create VMs. Named the field as 'instance_template' instead of 'template' to avoid c++ keyword conflict.
        /// </summary>
        public readonly string InstanceTemplate;
        /// <summary>
        /// InstancePolicy.
        /// </summary>
        public readonly Outputs.InstancePolicyResponse Policy;

        [OutputConstructor]
        private InstancePolicyOrTemplateResponse(
            bool installGpuDrivers,

            string instanceTemplate,

            Outputs.InstancePolicyResponse policy)
        {
            InstallGpuDrivers = installGpuDrivers;
            InstanceTemplate = instanceTemplate;
            Policy = policy;
        }
    }
}