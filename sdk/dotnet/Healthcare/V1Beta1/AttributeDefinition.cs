// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1
{
    /// <summary>
    /// Creates a new Attribute definition in the parent consent store.
    /// </summary>
    [GoogleNativeResourceType("google-native:healthcare/v1beta1:AttributeDefinition")]
    public partial class AttributeDefinition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Possible values for the attribute. The number of allowed values must not exceed 500. An empty list is invalid. The list can only be expanded after creation.
        /// </summary>
        [Output("allowedValues")]
        public Output<ImmutableArray<string>> AllowedValues { get; private set; } = null!;

        /// <summary>
        /// Required. The ID of the Attribute definition to create. The string must match the following regex: `_a-zA-Z{0,255}` and must not be a reserved keyword within the Common Expression Language as listed on https://github.com/google/cel-spec/blob/master/doc/langdef.md.
        /// </summary>
        [Output("attributeDefinitionId")]
        public Output<string> AttributeDefinitionId { get; private set; } = null!;

        /// <summary>
        /// The category of the attribute. The value of this field cannot be changed after creation.
        /// </summary>
        [Output("category")]
        public Output<string> Category { get; private set; } = null!;

        /// <summary>
        /// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
        /// </summary>
        [Output("consentDefaultValues")]
        public Output<ImmutableArray<string>> ConsentDefaultValues { get; private set; } = null!;

        [Output("consentStoreId")]
        public Output<string> ConsentStoreId { get; private set; } = null!;

        /// <summary>
        /// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
        /// </summary>
        [Output("dataMappingDefaultValue")]
        public Output<string> DataMappingDefaultValue { get; private set; } = null!;

        [Output("datasetId")]
        public Output<string> DatasetId { get; private set; } = null!;

        /// <summary>
        /// Optional. A description of the attribute.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a AttributeDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttributeDefinition(string name, AttributeDefinitionArgs args, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:AttributeDefinition", name, args ?? new AttributeDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttributeDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:healthcare/v1beta1:AttributeDefinition", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "attributeDefinitionId",
                    "consentStoreId",
                    "datasetId",
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
        /// Get an existing AttributeDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttributeDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AttributeDefinition(name, id, options);
        }
    }

    public sealed class AttributeDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedValues", required: true)]
        private InputList<string>? _allowedValues;

        /// <summary>
        /// Possible values for the attribute. The number of allowed values must not exceed 500. An empty list is invalid. The list can only be expanded after creation.
        /// </summary>
        public InputList<string> AllowedValues
        {
            get => _allowedValues ?? (_allowedValues = new InputList<string>());
            set => _allowedValues = value;
        }

        /// <summary>
        /// Required. The ID of the Attribute definition to create. The string must match the following regex: `_a-zA-Z{0,255}` and must not be a reserved keyword within the Common Expression Language as listed on https://github.com/google/cel-spec/blob/master/doc/langdef.md.
        /// </summary>
        [Input("attributeDefinitionId", required: true)]
        public Input<string> AttributeDefinitionId { get; set; } = null!;

        /// <summary>
        /// The category of the attribute. The value of this field cannot be changed after creation.
        /// </summary>
        [Input("category", required: true)]
        public Input<Pulumi.GoogleNative.Healthcare.V1Beta1.AttributeDefinitionCategory> Category { get; set; } = null!;

        [Input("consentDefaultValues")]
        private InputList<string>? _consentDefaultValues;

        /// <summary>
        /// Optional. Default values of the attribute in Consents. If no default values are specified, it defaults to an empty value.
        /// </summary>
        public InputList<string> ConsentDefaultValues
        {
            get => _consentDefaultValues ?? (_consentDefaultValues = new InputList<string>());
            set => _consentDefaultValues = value;
        }

        [Input("consentStoreId", required: true)]
        public Input<string> ConsentStoreId { get; set; } = null!;

        /// <summary>
        /// Optional. Default value of the attribute in User data mappings. If no default value is specified, it defaults to an empty value. This field is only applicable to attributes of the category `RESOURCE`.
        /// </summary>
        [Input("dataMappingDefaultValue")]
        public Input<string>? DataMappingDefaultValue { get; set; }

        [Input("datasetId", required: true)]
        public Input<string> DatasetId { get; set; } = null!;

        /// <summary>
        /// Optional. A description of the attribute.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource name of the Attribute definition, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/consentStores/{consent_store_id}/attributeDefinitions/{attribute_definition_id}`. Cannot be changed after creation.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        public AttributeDefinitionArgs()
        {
        }
        public static new AttributeDefinitionArgs Empty => new AttributeDefinitionArgs();
    }
}
