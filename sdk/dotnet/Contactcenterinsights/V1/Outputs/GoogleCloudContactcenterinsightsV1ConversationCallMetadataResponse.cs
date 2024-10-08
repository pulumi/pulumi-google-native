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
    /// Call-specific metadata.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1ConversationCallMetadataResponse
    {
        /// <summary>
        /// The audio channel that contains the agent.
        /// </summary>
        public readonly int AgentChannel;
        /// <summary>
        /// The audio channel that contains the customer.
        /// </summary>
        public readonly int CustomerChannel;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1ConversationCallMetadataResponse(
            int agentChannel,

            int customerChannel)
        {
            AgentChannel = agentChannel;
            CustomerChannel = customerChannel;
        }
    }
}
