// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    public static class GetExclusion
    {
        /// <summary>
        /// Gets the description of an exclusion in the _Default sink.
        /// </summary>
        public static Task<GetExclusionResult> InvokeAsync(GetExclusionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExclusionResult>("google-native:logging/v2:getExclusion", args ?? new GetExclusionArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the description of an exclusion in the _Default sink.
        /// </summary>
        public static Output<GetExclusionResult> Invoke(GetExclusionInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetExclusionResult>("google-native:logging/v2:getExclusion", args ?? new GetExclusionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExclusionArgs : Pulumi.InvokeArgs
    {
        [Input("exclusionId", required: true)]
        public string ExclusionId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetExclusionArgs()
        {
        }
    }

    public sealed class GetExclusionInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("exclusionId", required: true)]
        public Input<string> ExclusionId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetExclusionInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetExclusionResult
    {
        /// <summary>
        /// The creation timestamp of the exclusion.This field may not be present for older exclusions.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. A description of this exclusion.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
        /// </summary>
        public readonly bool Disabled;
        /// <summary>
        /// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity&lt;ERROR sample(insertId, 0.99)
        /// </summary>
        public readonly string Filter;
        /// <summary>
        /// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The last update timestamp of the exclusion.This field may not be present for older exclusions.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetExclusionResult(
            string createTime,

            string description,

            bool disabled,

            string filter,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            Disabled = disabled;
            Filter = filter;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}