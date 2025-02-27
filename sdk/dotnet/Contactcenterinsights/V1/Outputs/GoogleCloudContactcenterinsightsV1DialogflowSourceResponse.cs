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
    /// A Dialogflow source of conversation data.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1DialogflowSourceResponse
    {
        /// <summary>
        /// Cloud Storage URI that points to a file that contains the conversation audio.
        /// </summary>
        public readonly string AudioUri;
        /// <summary>
        /// The name of the Dialogflow conversation that this conversation resource is derived from. Format: projects/{project}/locations/{location}/conversations/{conversation}
        /// </summary>
        public readonly string DialogflowConversation;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1DialogflowSourceResponse(
            string audioUri,

            string dialogflowConversation)
        {
            AudioUri = audioUri;
            DialogflowConversation = dialogflowConversation;
        }
    }
}
