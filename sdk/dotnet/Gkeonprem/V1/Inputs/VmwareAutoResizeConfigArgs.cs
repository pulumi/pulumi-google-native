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
    /// Represents auto resizing configurations for the VMware user cluster.
    /// </summary>
    public sealed class VmwareAutoResizeConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable controle plane node auto resizing.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public VmwareAutoResizeConfigArgs()
        {
        }
        public static new VmwareAutoResizeConfigArgs Empty => new VmwareAutoResizeConfigArgs();
    }
}