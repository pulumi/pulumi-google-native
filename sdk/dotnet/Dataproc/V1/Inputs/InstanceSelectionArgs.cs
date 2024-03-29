// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Defines machines types and a rank to which the machines types belong.
    /// </summary>
    public sealed class InstanceSelectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("machineTypes")]
        private InputList<string>? _machineTypes;

        /// <summary>
        /// Optional. Full machine-type names, e.g. "n1-standard-16".
        /// </summary>
        public InputList<string> MachineTypes
        {
            get => _machineTypes ?? (_machineTypes = new InputList<string>());
            set => _machineTypes = value;
        }

        /// <summary>
        /// Optional. Preference of this instance selection. Lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        public InstanceSelectionArgs()
        {
        }
        public static new InstanceSelectionArgs Empty => new InstanceSelectionArgs();
    }
}
