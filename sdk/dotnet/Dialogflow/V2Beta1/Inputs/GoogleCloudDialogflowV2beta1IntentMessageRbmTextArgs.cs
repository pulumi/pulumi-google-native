// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Inputs
{

    /// <summary>
    /// Rich Business Messaging (RBM) text response with suggestions.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2beta1IntentMessageRbmTextArgs : Pulumi.ResourceArgs
    {
        [Input("rbmSuggestion")]
        private InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestionArgs>? _rbmSuggestion;

        /// <summary>
        /// Optional. One or more suggestions to show to the user.
        /// </summary>
        public InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestionArgs> RbmSuggestion
        {
            get => _rbmSuggestion ?? (_rbmSuggestion = new InputList<Inputs.GoogleCloudDialogflowV2beta1IntentMessageRbmSuggestionArgs>());
            set => _rbmSuggestion = value;
        }

        /// <summary>
        /// Text sent and displayed to the user.
        /// </summary>
        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        public GoogleCloudDialogflowV2beta1IntentMessageRbmTextArgs()
        {
        }
    }
}