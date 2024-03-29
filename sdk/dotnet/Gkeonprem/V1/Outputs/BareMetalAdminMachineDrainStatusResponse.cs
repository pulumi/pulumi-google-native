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
    /// BareMetalAdminMachineDrainStatus represents the status of bare metal node machines that are undergoing drain operations.
    /// </summary>
    [OutputType]
    public sealed class BareMetalAdminMachineDrainStatusResponse
    {
        /// <summary>
        /// The list of drained machines.
        /// </summary>
        public readonly ImmutableArray<Outputs.BareMetalAdminDrainedMachineResponse> DrainedMachines;
        /// <summary>
        /// The list of draning machines.
        /// </summary>
        public readonly ImmutableArray<Outputs.BareMetalAdminDrainingMachineResponse> DrainingMachines;

        [OutputConstructor]
        private BareMetalAdminMachineDrainStatusResponse(
            ImmutableArray<Outputs.BareMetalAdminDrainedMachineResponse> drainedMachines,

            ImmutableArray<Outputs.BareMetalAdminDrainingMachineResponse> drainingMachines)
        {
            DrainedMachines = drainedMachines;
            DrainingMachines = drainingMachines;
        }
    }
}
