// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contentwarehouse.V1
{
    /// <summary>
    /// Creates a document schema.
    /// </summary>
    [GoogleNativeResourceType("google-native:contentwarehouse/v1:DocumentSchema")]
    public partial class DocumentSchema : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the document schema is created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Schema description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the schema given by the user. Must be unique per project.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Document Type, true refers the document is a folder, otherwise it is a typical document.
        /// </summary>
        [Output("documentIsFolder")]
        public Output<bool> DocumentIsFolder { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Document details.
        /// </summary>
        [Output("propertyDefinitions")]
        public Output<ImmutableArray<Outputs.GoogleCloudContentwarehouseV1PropertyDefinitionResponse>> PropertyDefinitions { get; private set; } = null!;

        /// <summary>
        /// The time when the document schema is last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentSchema resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentSchema(string name, DocumentSchemaArgs args, CustomResourceOptions? options = null)
            : base("google-native:contentwarehouse/v1:DocumentSchema", name, args ?? new DocumentSchemaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentSchema(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:contentwarehouse/v1:DocumentSchema", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DocumentSchema resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentSchema Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DocumentSchema(name, id, options);
        }
    }

    public sealed class DocumentSchemaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Schema description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the schema given by the user. Must be unique per project.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Document Type, true refers the document is a folder, otherwise it is a typical document.
        /// </summary>
        [Input("documentIsFolder")]
        public Input<bool>? DocumentIsFolder { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("propertyDefinitions")]
        private InputList<Inputs.GoogleCloudContentwarehouseV1PropertyDefinitionArgs>? _propertyDefinitions;

        /// <summary>
        /// Document details.
        /// </summary>
        public InputList<Inputs.GoogleCloudContentwarehouseV1PropertyDefinitionArgs> PropertyDefinitions
        {
            get => _propertyDefinitions ?? (_propertyDefinitions = new InputList<Inputs.GoogleCloudContentwarehouseV1PropertyDefinitionArgs>());
            set => _propertyDefinitions = value;
        }

        public DocumentSchemaArgs()
        {
        }
        public static new DocumentSchemaArgs Empty => new DocumentSchemaArgs();
    }
}