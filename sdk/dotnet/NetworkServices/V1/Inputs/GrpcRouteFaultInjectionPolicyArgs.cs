// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1.Inputs
{

    /// <summary>
    /// The specification for fault injection introduced into traffic to test the resiliency of clients to destination service failure. As part of fault injection, when clients send requests to a destination, delays can be introduced on a percentage of requests before sending those requests to the destination service. Similarly requests from clients can be aborted by for a percentage of requests.
    /// </summary>
    public sealed class GrpcRouteFaultInjectionPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The specification for aborting to client requests.
        /// </summary>
        [Input("abort")]
        public Input<Inputs.GrpcRouteFaultInjectionPolicyAbortArgs>? Abort { get; set; }

        /// <summary>
        /// The specification for injecting delay to client requests.
        /// </summary>
        [Input("delay")]
        public Input<Inputs.GrpcRouteFaultInjectionPolicyDelayArgs>? Delay { get; set; }

        public GrpcRouteFaultInjectionPolicyArgs()
        {
        }
        public static new GrpcRouteFaultInjectionPolicyArgs Empty => new GrpcRouteFaultInjectionPolicyArgs();
    }
}
