// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Outputs
{

    /// <summary>
    /// The data for a Dialogflow intent. Represents a detected intent in the conversation, e.g. MAKES_PROMISE.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1DialogflowIntentResponse
    {
        /// <summary>
        /// The human-readable name of the intent.
        /// </summary>
        public readonly string DisplayName;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1DialogflowIntentResponse(string displayName)
        {
            DisplayName = displayName;
        }
    }
}