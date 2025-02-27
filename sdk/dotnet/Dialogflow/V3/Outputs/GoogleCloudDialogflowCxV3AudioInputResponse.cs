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
    /// Represents the natural speech audio to be processed.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3AudioInputResponse
    {
        /// <summary>
        /// The natural language speech audio to be processed. A single request can contain up to 2 minutes of speech audio data. The transcribed text cannot contain more than 256 bytes. For non-streaming audio detect intent, both `config` and `audio` must be provided. For streaming audio detect intent, `config` must be provided in the first request and `audio` must be provided in all following requests.
        /// </summary>
        public readonly string Audio;
        /// <summary>
        /// Instructs the speech recognizer how to process the speech audio.
        /// </summary>
        public readonly Outputs.GoogleCloudDialogflowCxV3InputAudioConfigResponse Config;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3AudioInputResponse(
            string audio,

            Outputs.GoogleCloudDialogflowCxV3InputAudioConfigResponse config)
        {
            Audio = audio;
            Config = config;
        }
    }
}
