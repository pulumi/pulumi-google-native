// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2Alpha1.Inputs
{

    /// <summary>
    /// Defines the policy of the QueuedRequest.
    /// </summary>
    public sealed class QueueingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A relative time after which resources may be created.
        /// </summary>
        [Input("validAfterDuration")]
        public Input<string>? ValidAfterDuration { get; set; }

        /// <summary>
        /// An absolute time at which resources may be created.
        /// </summary>
        [Input("validAfterTime")]
        public Input<string>? ValidAfterTime { get; set; }

        /// <summary>
        /// An absolute time interval within which resources may be created.
        /// </summary>
        [Input("validInterval")]
        public Input<Inputs.IntervalArgs>? ValidInterval { get; set; }

        /// <summary>
        /// A relative time after which resources should not be created. If the request cannot be fulfilled by this time the request will be failed.
        /// </summary>
        [Input("validUntilDuration")]
        public Input<string>? ValidUntilDuration { get; set; }

        /// <summary>
        /// An absolute time after which resources should not be created. If the request cannot be fulfilled by this time the request will be failed.
        /// </summary>
        [Input("validUntilTime")]
        public Input<string>? ValidUntilTime { get; set; }

        public QueueingPolicyArgs()
        {
        }
        public static new QueueingPolicyArgs Empty => new QueueingPolicyArgs();
    }
}