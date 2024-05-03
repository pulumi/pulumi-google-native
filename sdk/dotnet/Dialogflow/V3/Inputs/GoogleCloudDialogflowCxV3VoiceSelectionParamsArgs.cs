// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Description of which voice to use for speech synthesis.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3VoiceSelectionParamsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The name of the voice. If not set, the service will choose a voice based on the other parameters such as language_code and ssml_gender. For the list of available voices, please refer to [Supported voices and languages](https://cloud.google.com/text-to-speech/docs/voices).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Optional. The preferred gender of the voice. If not set, the service will choose a voice based on the other parameters such as language_code and name. Note that this is only a preference, not requirement. If a voice of the appropriate gender is not available, the synthesizer substitutes a voice with a different gender rather than failing the request.
        /// </summary>
        [Input("ssmlGender")]
        public Input<Pulumi.GoogleNative.Dialogflow.V3.GoogleCloudDialogflowCxV3VoiceSelectionParamsSsmlGender>? SsmlGender { get; set; }

        public GoogleCloudDialogflowCxV3VoiceSelectionParamsArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3VoiceSelectionParamsArgs Empty => new GoogleCloudDialogflowCxV3VoiceSelectionParamsArgs();
    }
}