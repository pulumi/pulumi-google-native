// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkServices.V1Beta1.Inputs
{

    /// <summary>
    /// Option to specify health based failover behavior. This is not related to Network load balancer FailoverPolicy.
    /// </summary>
    public sealed class ServiceLbPolicyFailoverConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The percentage threshold that a load balancer will begin to send traffic to failover backends. If the percentage of endpoints in a MIG/NEG is smaller than this value, traffic would be sent to failover backends if possible. This field should be set to a value between 1 and 99. The default value is 50 for Global external HTTP(S) load balancer (classic) and Proxyless service mesh, and 70 for others.
        /// </summary>
        [Input("failoverHealthThreshold")]
        public Input<int>? FailoverHealthThreshold { get; set; }

        public ServiceLbPolicyFailoverConfigArgs()
        {
        }
        public static new ServiceLbPolicyFailoverConfigArgs Empty => new ServiceLbPolicyFailoverConfigArgs();
    }
}