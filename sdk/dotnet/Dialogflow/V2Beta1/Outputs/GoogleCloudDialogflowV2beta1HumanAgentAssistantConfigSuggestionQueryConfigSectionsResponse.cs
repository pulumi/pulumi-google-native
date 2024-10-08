// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2Beta1.Outputs
{

    /// <summary>
    /// Custom sections to return when requesting a summary of a conversation. This is only supported when `baseline_model_version` == '2.0'. Supported features: CONVERSATION_SUMMARIZATION, CONVERSATION_SUMMARIZATION_VOICE.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionQueryConfigSectionsResponse
    {
        /// <summary>
        /// The selected sections chosen to return when requesting a summary of a conversation. A duplicate selected section will be treated as a single selected section. If section types are not provided, the default will be {SITUATION, ACTION, RESULT}.
        /// </summary>
        public readonly ImmutableArray<string> SectionTypes;

        [OutputConstructor]
        private GoogleCloudDialogflowV2beta1HumanAgentAssistantConfigSuggestionQueryConfigSectionsResponse(ImmutableArray<string> sectionTypes)
        {
            SectionTypes = sectionTypes;
        }
    }
}
