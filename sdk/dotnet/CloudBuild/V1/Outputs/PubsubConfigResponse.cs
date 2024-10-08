// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// PubsubConfig describes the configuration of a trigger that creates a build whenever a Pub/Sub message is published.
    /// </summary>
    [OutputType]
    public sealed class PubsubConfigResponse
    {
        /// <summary>
        /// Service account that will make the push request.
        /// </summary>
        public readonly string ServiceAccountEmail;
        /// <summary>
        /// Potential issues with the underlying Pub/Sub subscription configuration. Only populated on get requests.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Name of the subscription. Format is `projects/{project}/subscriptions/{subscription}`.
        /// </summary>
        public readonly string Subscription;
        /// <summary>
        /// The name of the topic from which this subscription is receiving messages. Format is `projects/{project}/topics/{topic}`.
        /// </summary>
        public readonly string Topic;

        [OutputConstructor]
        private PubsubConfigResponse(
            string serviceAccountEmail,

            string state,

            string subscription,

            string topic)
        {
            ServiceAccountEmail = serviceAccountEmail;
            State = state;
            Subscription = subscription;
            Topic = topic;
        }
    }
}
