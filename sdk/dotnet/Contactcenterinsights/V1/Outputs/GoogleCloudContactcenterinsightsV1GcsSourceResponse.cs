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
    /// A Cloud Storage source of conversation data.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1GcsSourceResponse
    {
        /// <summary>
        /// Cloud Storage URI that points to a file that contains the conversation audio.
        /// </summary>
        public readonly string AudioUri;
        /// <summary>
        /// Immutable. Cloud Storage URI that points to a file that contains the conversation transcript.
        /// </summary>
        public readonly string TranscriptUri;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1GcsSourceResponse(
            string audioUri,

            string transcriptUri)
        {
            AudioUri = audioUri;
            TranscriptUri = transcriptUri;
        }
    }
}