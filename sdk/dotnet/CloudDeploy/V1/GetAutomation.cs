// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1
{
    public static class GetAutomation
    {
        /// <summary>
        /// Gets details of a single Automation.
        /// </summary>
        public static Task<GetAutomationResult> InvokeAsync(GetAutomationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAutomationResult>("google-native:clouddeploy/v1:getAutomation", args ?? new GetAutomationArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single Automation.
        /// </summary>
        public static Output<GetAutomationResult> Invoke(GetAutomationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAutomationResult>("google-native:clouddeploy/v1:getAutomation", args ?? new GetAutomationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAutomationArgs : global::Pulumi.InvokeArgs
    {
        [Input("automationId", required: true)]
        public string AutomationId { get; set; } = null!;

        [Input("deliveryPipelineId", required: true)]
        public string DeliveryPipelineId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAutomationArgs()
        {
        }
        public static new GetAutomationArgs Empty => new GetAutomationArgs();
    }

    public sealed class GetAutomationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("automationId", required: true)]
        public Input<string> AutomationId { get; set; } = null!;

        [Input("deliveryPipelineId", required: true)]
        public Input<string> DeliveryPipelineId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAutomationInvokeArgs()
        {
        }
        public static new GetAutomationInvokeArgs Empty => new GetAutomationInvokeArgs();
    }


    [OutputType]
    public sealed class GetAutomationResult
    {
        /// <summary>
        /// Optional. User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. Annotations must meet the following constraints: * Annotations are key/value pairs. * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (`/`). * The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between. * The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots(`.`), not longer than 253 characters in total, followed by a slash (`/`). See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set for more details.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Time at which the automation was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Description of the `Automation`. Max length is 255 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. The weak etag of the `Automation` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be &lt;= 63 characters.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the `Automation`. Format is `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/automations/{automation}`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of Automation rules associated with the Automation resource. Must have at least one rule and limited to 250 rules per Delivery Pipeline. Note: the order of the rules here is not the same as the order of execution.
        /// </summary>
        public readonly ImmutableArray<Outputs.AutomationRuleResponse> Rules;
        /// <summary>
        /// Selected resources to which the automation will be applied.
        /// </summary>
        public readonly Outputs.AutomationResourceSelectorResponse Selector;
        /// <summary>
        /// Email address of the user-managed IAM service account that creates Cloud Deploy release and rollout resources.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Optional. When Suspended, automation is deactivated from execution.
        /// </summary>
        public readonly bool Suspended;
        /// <summary>
        /// Unique identifier of the `Automation`.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// Time at which the automation was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAutomationResult(
            ImmutableDictionary<string, string> annotations,

            string createTime,

            string description,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            ImmutableArray<Outputs.AutomationRuleResponse> rules,

            Outputs.AutomationResourceSelectorResponse selector,

            string serviceAccount,

            bool suspended,

            string uid,

            string updateTime)
        {
            Annotations = annotations;
            CreateTime = createTime;
            Description = description;
            Etag = etag;
            Labels = labels;
            Name = name;
            Rules = rules;
            Selector = selector;
            ServiceAccount = serviceAccount;
            Suspended = suspended;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}