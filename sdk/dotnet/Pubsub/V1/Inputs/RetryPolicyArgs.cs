// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Pubsub.V1.Inputs
{

    /// <summary>
    /// A policy that specifies how Cloud Pub/Sub retries message delivery. Retry delay will be exponential based on provided minimum and maximum backoffs. https://en.wikipedia.org/wiki/Exponential_backoff. RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message. Retry Policy is implemented on a best effort basis. At times, the delay between consecutive deliveries may not match the configuration. That is, delay can be more or less than configured backoff.
    /// </summary>
    public sealed class RetryPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds.
        /// </summary>
        [Input("maximumBackoff")]
        public Input<string>? MaximumBackoff { get; set; }

        /// <summary>
        /// The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
        /// </summary>
        [Input("minimumBackoff")]
        public Input<string>? MinimumBackoff { get; set; }

        public RetryPolicyArgs()
        {
        }
    }
}