// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1
{
    /// <summary>
    /// Create a new channel in a particular project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:eventarc/v1:Channel")]
    public partial class Channel : Pulumi.CustomResource
    {
        /// <summary>
        /// The activation token for the channel. The token must be used by the provider to register the channel for publishing.
        /// </summary>
        [Output("activationToken")]
        public Output<string> ActivationToken { get; private set; } = null!;

        /// <summary>
        /// Required. The user-provided ID to be assigned to the channel.
        /// </summary>
        [Output("channelId")]
        public Output<string> ChannelId { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
        /// </summary>
        [Output("cryptoKeyName")]
        public Output<string> CryptoKeyName { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the channel. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/channels/{channel_id}` format.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
        /// </summary>
        [Output("provider")]
        public Output<string> Provider { get; private set; } = null!;

        /// <summary>
        /// The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{project}/topics/{topic_id}`.
        /// </summary>
        [Output("pubsubTopic")]
        public Output<string> PubsubTopic { get; private set; } = null!;

        /// <summary>
        /// The state of a Channel.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Server assigned unique identifier for the channel. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Required. If set, validate the request and preview the review, but do not post it.
        /// </summary>
        [Output("validateOnly")]
        public Output<string> ValidateOnly { get; private set; } = null!;


        /// <summary>
        /// Create a Channel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Channel(string name, ChannelArgs args, CustomResourceOptions? options = null)
            : base("google-native:eventarc/v1:Channel", name, args ?? new ChannelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Channel(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:eventarc/v1:Channel", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Channel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Channel Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Channel(name, id, options);
        }
    }

    public sealed class ChannelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The user-provided ID to be assigned to the channel.
        /// </summary>
        [Input("channelId", required: true)]
        public Input<string> ChannelId { get; set; } = null!;

        /// <summary>
        /// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/*/locations/*/keyRings/*/cryptoKeys/*`.
        /// </summary>
        [Input("cryptoKeyName")]
        public Input<string>? CryptoKeyName { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the channel. Must be unique within the location on the project and must be in `projects/{project}/locations/{location}/channels/{channel_id}` format.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
        /// </summary>
        [Input("provider")]
        public Input<string>? Provider { get; set; }

        /// <summary>
        /// Required. If set, validate the request and preview the review, but do not post it.
        /// </summary>
        [Input("validateOnly", required: true)]
        public Input<string> ValidateOnly { get; set; } = null!;

        public ChannelArgs()
        {
        }
    }
}