// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DiscoveryEngine.V1Alpha.Inputs
{

    /// <summary>
    /// Defines a conversation message.
    /// </summary>
    public sealed class GoogleCloudDiscoveryengineV1alphaConversationMessageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search reply.
        /// </summary>
        [Input("reply")]
        public Input<Inputs.GoogleCloudDiscoveryengineV1alphaReplyArgs>? Reply { get; set; }

        /// <summary>
        /// User text input.
        /// </summary>
        [Input("userInput")]
        public Input<Inputs.GoogleCloudDiscoveryengineV1alphaTextInputArgs>? UserInput { get; set; }

        public GoogleCloudDiscoveryengineV1alphaConversationMessageArgs()
        {
        }
        public static new GoogleCloudDiscoveryengineV1alphaConversationMessageArgs Empty => new GoogleCloudDiscoveryengineV1alphaConversationMessageArgs();
    }
}
