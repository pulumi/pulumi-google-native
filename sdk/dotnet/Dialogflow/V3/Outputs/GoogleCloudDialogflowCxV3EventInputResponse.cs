// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// Represents the event to trigger.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3EventInputResponse
    {
        /// <summary>
        /// Name of the event.
        /// </summary>
        public readonly string Event;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3EventInputResponse(string @event)
        {
            Event = @event;
        }
    }
}