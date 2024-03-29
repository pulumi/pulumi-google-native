// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Inputs
{

    /// <summary>
    /// Specifies config to enable/disable auto repair. The cluster-health-controller is deployed only if Enabled is true.
    /// </summary>
    public sealed class VmwareAutoRepairConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether auto repair is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public VmwareAutoRepairConfigArgs()
        {
        }
        public static new VmwareAutoRepairConfigArgs Empty => new VmwareAutoRepairConfigArgs();
    }
}
