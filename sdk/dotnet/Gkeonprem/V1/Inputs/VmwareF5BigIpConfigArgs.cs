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
    /// Represents configuration parameters for an F5 BIG-IP load balancer.
    /// </summary>
    public sealed class VmwareF5BigIpConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The load balancer's IP address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The preexisting partition to be used by the load balancer. This partition is usually created for the admin cluster for example: 'my-f5-admin-partition'.
        /// </summary>
        [Input("partition")]
        public Input<string>? Partition { get; set; }

        /// <summary>
        /// The pool name. Only necessary, if using SNAT.
        /// </summary>
        [Input("snatPool")]
        public Input<string>? SnatPool { get; set; }

        public VmwareF5BigIpConfigArgs()
        {
        }
        public static new VmwareF5BigIpConfigArgs Empty => new VmwareF5BigIpConfigArgs();
    }
}