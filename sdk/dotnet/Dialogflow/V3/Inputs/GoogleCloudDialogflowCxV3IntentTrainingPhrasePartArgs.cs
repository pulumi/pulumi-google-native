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
    /// Represents a part of a training phrase.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3IntentTrainingPhrasePartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parameter used to annotate this part of the training phrase. This field is required for annotated parts of the training phrase.
        /// </summary>
        [Input("parameterId")]
        public Input<string>? ParameterId { get; set; }

        /// <summary>
        /// The text for this part.
        /// </summary>
        [Input("text", required: true)]
        public Input<string> Text { get; set; } = null!;

        public GoogleCloudDialogflowCxV3IntentTrainingPhrasePartArgs()
        {
        }
    }
}