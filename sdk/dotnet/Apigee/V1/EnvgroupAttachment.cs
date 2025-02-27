// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates a new attachment of an environment to an environment group.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:EnvgroupAttachment")]
    public partial class EnvgroupAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time at which the environment group attachment was created as milliseconds since epoch.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("envgroupId")]
        public Output<string> EnvgroupId { get; private set; } = null!;

        /// <summary>
        /// ID of the attached environment.
        /// </summary>
        [Output("environment")]
        public Output<string> Environment { get; private set; } = null!;

        /// <summary>
        /// ID of the environment group.
        /// </summary>
        [Output("environmentGroupId")]
        public Output<string> EnvironmentGroupId { get; private set; } = null!;

        /// <summary>
        /// ID of the environment group attachment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;


        /// <summary>
        /// Create a EnvgroupAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvgroupAttachment(string name, EnvgroupAttachmentArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:EnvgroupAttachment", name, args ?? new EnvgroupAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvgroupAttachment(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:EnvgroupAttachment", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "envgroupId",
                    "organizationId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EnvgroupAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvgroupAttachment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EnvgroupAttachment(name, id, options);
        }
    }

    public sealed class EnvgroupAttachmentArgs : global::Pulumi.ResourceArgs
    {
        [Input("envgroupId", required: true)]
        public Input<string> EnvgroupId { get; set; } = null!;

        /// <summary>
        /// ID of the attached environment.
        /// </summary>
        [Input("environment", required: true)]
        public Input<string> Environment { get; set; } = null!;

        /// <summary>
        /// ID of the environment group attachment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        public EnvgroupAttachmentArgs()
        {
        }
        public static new EnvgroupAttachmentArgs Empty => new EnvgroupAttachmentArgs();
    }
}
