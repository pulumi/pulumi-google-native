// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1
{
    public static class GetGroup
    {
        /// <summary>
        /// Gets the details of a group.
        /// </summary>
        public static Task<GetGroupResult> InvokeAsync(GetGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGroupResult>("google-native:migrationcenter/v1:getGroup", args ?? new GetGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets the details of a group.
        /// </summary>
        public static Output<GetGroupResult> Invoke(GetGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGroupResult>("google-native:migrationcenter/v1:getGroup", args ?? new GetGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetGroupArgs()
        {
        }
        public static new GetGroupArgs Empty => new GetGroupArgs();
    }

    public sealed class GetGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetGroupInvokeArgs()
        {
        }
        public static new GetGroupInvokeArgs Empty => new GetGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetGroupResult
    {
        /// <summary>
        /// The timestamp when the group was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. The description of the group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. User-friendly display name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The name of the group.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The timestamp when the group was last updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetGroupResult(
            string createTime,

            string description,

            string displayName,

            ImmutableDictionary<string, string> labels,

            string name,

            string updateTime)
        {
            CreateTime = createTime;
            Description = description;
            DisplayName = displayName;
            Labels = labels;
            Name = name;
            UpdateTime = updateTime;
        }
    }
}