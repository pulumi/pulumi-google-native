// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkSecurity.V1
{
    public static class GetOrganizationAddressGroup
    {
        /// <summary>
        /// Gets details of a single address group.
        /// </summary>
        public static Task<GetOrganizationAddressGroupResult> InvokeAsync(GetOrganizationAddressGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganizationAddressGroupResult>("google-native:networksecurity/v1:getOrganizationAddressGroup", args ?? new GetOrganizationAddressGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single address group.
        /// </summary>
        public static Output<GetOrganizationAddressGroupResult> Invoke(GetOrganizationAddressGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganizationAddressGroupResult>("google-native:networksecurity/v1:getOrganizationAddressGroup", args ?? new GetOrganizationAddressGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganizationAddressGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressGroupId", required: true)]
        public string AddressGroupId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("organizationId", required: true)]
        public string OrganizationId { get; set; } = null!;

        public GetOrganizationAddressGroupArgs()
        {
        }
        public static new GetOrganizationAddressGroupArgs Empty => new GetOrganizationAddressGroupArgs();
    }

    public sealed class GetOrganizationAddressGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressGroupId", required: true)]
        public Input<string> AddressGroupId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public GetOrganizationAddressGroupInvokeArgs()
        {
        }
        public static new GetOrganizationAddressGroupInvokeArgs Empty => new GetOrganizationAddressGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganizationAddressGroupResult
    {
        /// <summary>
        /// Capacity of the Address Group
        /// </summary>
        public readonly int Capacity;
        /// <summary>
        /// The timestamp when the resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Optional. Free-text description of the resource.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Optional. List of items.
        /// </summary>
        public readonly ImmutableArray<string> Items;
        /// <summary>
        /// Optional. Set of label tags associated with the AddressGroup resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Name of the AddressGroup resource. It matches pattern `projects/*/locations/{location}/addressGroups/`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Server-defined fully-qualified URL for this resource.
        /// </summary>
        public readonly string SelfLink;
        /// <summary>
        /// The type of the Address Group. Possible values are "IPv4" or "IPV6".
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The timestamp when the resource was updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetOrganizationAddressGroupResult(
            int capacity,

            string createTime,

            string description,

            ImmutableArray<string> items,

            ImmutableDictionary<string, string> labels,

            string name,

            string selfLink,

            string type,

            string updateTime)
        {
            Capacity = capacity;
            CreateTime = createTime;
            Description = description;
            Items = items;
            Labels = labels;
            Name = name;
            SelfLink = selfLink;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
