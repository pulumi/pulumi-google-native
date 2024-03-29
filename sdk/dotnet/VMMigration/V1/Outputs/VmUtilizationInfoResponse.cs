// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1.Outputs
{

    /// <summary>
    /// Utilization information of a single VM.
    /// </summary>
    [OutputType]
    public sealed class VmUtilizationInfoResponse
    {
        /// <summary>
        /// Utilization metrics for this VM.
        /// </summary>
        public readonly Outputs.VmUtilizationMetricsResponse Utilization;
        /// <summary>
        /// The VM's ID in the source.
        /// </summary>
        public readonly string VmId;
        /// <summary>
        /// The description of the VM in a Source of type Vmware.
        /// </summary>
        public readonly Outputs.VmwareVmDetailsResponse VmwareVmDetails;

        [OutputConstructor]
        private VmUtilizationInfoResponse(
            Outputs.VmUtilizationMetricsResponse utilization,

            string vmId,

            Outputs.VmwareVmDetailsResponse vmwareVmDetails)
        {
            Utilization = utilization;
            VmId = vmId;
            VmwareVmDetails = vmwareVmDetails;
        }
    }
}
