// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1Beta1
{
    /// <summary>
    /// Create a new trigger in a particular project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:eventarc/v1beta1:Trigger")]
    public partial class Trigger : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Destination specifies where the events should be sent to.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.DestinationResponse> Destination { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent only on create requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Optional. User labels attached to the triggers that can be used to group resources.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        /// </summary>
        [Output("matchingCriteria")]
        public Output<ImmutableArray<Outputs.MatchingCriteriaResponse>> MatchingCriteria { get; private set; } = null!;

        /// <summary>
        /// The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// In order to deliver messages, Eventarc may use other GCP products as transport intermediary. This field contains a reference to that transport intermediary. This information can be used for debugging purposes.
        /// </summary>
        [Output("transport")]
        public Output<Outputs.TransportResponse> Transport { get; private set; } = null!;

        /// <summary>
        /// The last-modified time.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Trigger resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Trigger(string name, TriggerArgs args, CustomResourceOptions? options = null)
            : base("google-native:eventarc/v1beta1:Trigger", name, args ?? new TriggerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Trigger(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:eventarc/v1beta1:Trigger", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Trigger resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Trigger Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Trigger(name, id, options);
        }
    }

    public sealed class TriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination specifies where the events should be sent to.
        /// </summary>
        [Input("destination", required: true)]
        public Input<Inputs.DestinationArgs> Destination { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User labels attached to the triggers that can be used to group resources.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("matchingCriteria", required: true)]
        private InputList<Inputs.MatchingCriteriaArgs>? _matchingCriteria;

        /// <summary>
        /// null The criteria by which events are filtered. Only events that match with this criteria will be sent to the destination.
        /// </summary>
        public InputList<Inputs.MatchingCriteriaArgs> MatchingCriteria
        {
            get => _matchingCriteria ?? (_matchingCriteria = new InputList<Inputs.MatchingCriteriaArgs>());
            set => _matchingCriteria = value;
        }

        /// <summary>
        /// The resource name of the trigger. Must be unique within the location on the project and must in `projects/{project}/locations/{location}/triggers/{trigger}` format.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. The IAM service account email associated with the trigger. The service account represents the identity of the trigger. The principal who calls this API must have `iam.serviceAccounts.actAs` permission in the service account. See https://cloud.google.com/iam/docs/understanding-service-accounts?hl=en#sa_common for more information. For Cloud Run destinations, this service account is used to generate identity tokens when invoking the service. See https://cloud.google.com/run/docs/triggering/pubsub-push#create-service-account for information on how to invoke authenticated Cloud Run services. In order to create Audit Log triggers, the service account should also have 'eventarc.events.receiveAuditLogV1Written' permission.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// Required. The user-provided ID to be assigned to the trigger.
        /// </summary>
        [Input("triggerId", required: true)]
        public Input<string> TriggerId { get; set; } = null!;

        /// <summary>
        /// Required. If set, validate the request and preview the review, but do not actually post it.
        /// </summary>
        [Input("validateOnly", required: true)]
        public Input<string> ValidateOnly { get; set; } = null!;

        public TriggerArgs()
        {
        }
    }
}