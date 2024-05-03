// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Per-instance properties to be set on individual instances. To be extended in the future.
    /// </summary>
    public sealed class BulkInsertInstanceResourcePerInstancePropertiesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the hostname of the instance. More details in: https://cloud.google.com/compute/docs/instances/custom-hostname-vm#naming_convention
        /// </summary>
        [Input("hostname")]
        public Input<string>? Hostname { get; set; }

        /// <summary>
        /// This field is only temporary. It will be removed. Do not use it.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public BulkInsertInstanceResourcePerInstancePropertiesArgs()
        {
        }
        public static new BulkInsertInstanceResourcePerInstancePropertiesArgs Empty => new BulkInsertInstanceResourcePerInstancePropertiesArgs();
    }
}