// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V2.Inputs
{

    /// <summary>
    /// Defines the hand off to a live agent, typically on which external agent service provider to connect to a conversation. Currently, this feature is not general available, please contact Google to get access.
    /// </summary>
    public sealed class GoogleCloudDialogflowV2HumanAgentHandoffConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Uses LivePerson (https://www.liveperson.com).
        /// </summary>
        [Input("livePersonConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigLivePersonConfigArgs>? LivePersonConfig { get; set; }

        /// <summary>
        /// Uses Salesforce Live Agent.
        /// </summary>
        [Input("salesforceLiveAgentConfig")]
        public Input<Inputs.GoogleCloudDialogflowV2HumanAgentHandoffConfigSalesforceLiveAgentConfigArgs>? SalesforceLiveAgentConfig { get; set; }

        public GoogleCloudDialogflowV2HumanAgentHandoffConfigArgs()
        {
        }
    }
}