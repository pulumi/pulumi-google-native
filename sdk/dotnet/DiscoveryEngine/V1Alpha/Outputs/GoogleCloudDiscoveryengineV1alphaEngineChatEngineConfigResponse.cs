// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DiscoveryEngine.V1Alpha.Outputs
{

    /// <summary>
    /// Configurations for a Chat Engine.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigResponse
    {
        /// <summary>
        /// The configurationt generate the Dialogflow agent that is associated to this Engine. Note that these configurations are one-time consumed by and passed to Dialogflow service. It means they cannot be retrieved using EngineService.GetEngine or EngineService.ListEngines API after engine creation.
        /// </summary>
        public readonly Outputs.GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigAgentCreationConfigResponse AgentCreationConfig;
        /// <summary>
        /// The resource name of an exist Dialogflow agent to link to this Chat Engine. Customers can either provide `agent_creation_config` to create agent or provide an agent name that links the agent with the Chat engine. Format: `projects//locations//agents/`. Note that the `dialogflow_agent_to_link` are one-time consumed by and passed to Dialogflow service. It means they cannot be retrieved using EngineService.GetEngine or EngineService.ListEngines API after engine creation. Please use chat_engine_metadata.dialogflow_agent for actual agent association after Engine is created.
        /// </summary>
        public readonly string DialogflowAgentToLink;

        [OutputConstructor]
        private GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigResponse(
            Outputs.GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigAgentCreationConfigResponse agentCreationConfig,

            string dialogflowAgentToLink)
        {
            AgentCreationConfig = agentCreationConfig;
            DialogflowAgentToLink = dialogflowAgentToLink;
        }
    }
}