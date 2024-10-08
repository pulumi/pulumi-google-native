// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1.Inputs
{

    /// <summary>
    /// The type of machines to consider when calculating virtual machine migration insights and recommendations. Not all machine types are available in all zones and regions.
    /// </summary>
    public sealed class MachinePreferencesArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedMachineSeries")]
        private InputList<Inputs.MachineSeriesArgs>? _allowedMachineSeries;

        /// <summary>
        /// Compute Engine machine series to consider for insights and recommendations. If empty, no restriction is applied on the machine series.
        /// </summary>
        public InputList<Inputs.MachineSeriesArgs> AllowedMachineSeries
        {
            get => _allowedMachineSeries ?? (_allowedMachineSeries = new InputList<Inputs.MachineSeriesArgs>());
            set => _allowedMachineSeries = value;
        }

        public MachinePreferencesArgs()
        {
        }
        public static new MachinePreferencesArgs Empty => new MachinePreferencesArgs();
    }
}
