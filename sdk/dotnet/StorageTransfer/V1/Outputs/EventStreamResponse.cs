// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Outputs
{

    /// <summary>
    /// Specifies the Event-driven transfer options. Event-driven transfers listen to an event stream to transfer updated files.
    /// </summary>
    [OutputType]
    public sealed class EventStreamResponse
    {
        /// <summary>
        /// Specifies the data and time at which Storage Transfer Service stops listening for events from this stream. After this time, any transfers in progress will complete, but no new transfers are initiated.
        /// </summary>
        public readonly string EventStreamExpirationTime;
        /// <summary>
        /// Specifies the date and time that Storage Transfer Service starts listening for events from this stream. If no start time is specified or start time is in the past, Storage Transfer Service starts listening immediately.
        /// </summary>
        public readonly string EventStreamStartTime;
        /// <summary>
        /// Specifies a unique name of the resource such as AWS SQS ARN in the form 'arn:aws:sqs:region:account_id:queue_name', or Pub/Sub subscription resource name in the form 'projects/{project}/subscriptions/{sub}'.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private EventStreamResponse(
            string eventStreamExpirationTime,

            string eventStreamStartTime,

            string name)
        {
            EventStreamExpirationTime = eventStreamExpirationTime;
            EventStreamStartTime = eventStreamStartTime;
            Name = name;
        }
    }
}