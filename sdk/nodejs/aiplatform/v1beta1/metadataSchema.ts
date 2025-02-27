// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a MetadataSchema.
 * Auto-naming is currently not supported for this resource.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class MetadataSchema extends pulumi.CustomResource {
    /**
     * Get an existing MetadataSchema resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): MetadataSchema {
        return new MetadataSchema(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:aiplatform/v1beta1:MetadataSchema';

    /**
     * Returns true if the given object is an instance of MetadataSchema.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MetadataSchema {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MetadataSchema.__pulumiType;
    }

    /**
     * Timestamp when this MetadataSchema was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the Metadata Schema
     */
    public readonly description!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The {metadata_schema} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}/metadataSchemas/{metadataschema}` If not provided, the MetadataStore's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all MetadataSchemas in the parent Location. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting MetadataSchema.)
     */
    public readonly metadataSchemaId!: pulumi.Output<string | undefined>;
    public readonly metadataStoreId!: pulumi.Output<string>;
    /**
     * The resource name of the MetadataSchema.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The raw YAML string representation of the MetadataSchema. The combination of [MetadataSchema.version] and the schema name given by `title` in [MetadataSchema.schema] must be unique within a MetadataStore. The schema is defined as an OpenAPI 3.0.2 [MetadataSchema Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)
     */
    public readonly schema!: pulumi.Output<string>;
    /**
     * The type of the MetadataSchema. This is a property that identifies which metadata types will use the MetadataSchema.
     */
    public readonly schemaType!: pulumi.Output<string>;
    /**
     * The version of the MetadataSchema. The version's format must match the following regular expression: `^[0-9]+.+.+$`, which would allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.
     */
    public readonly schemaVersion!: pulumi.Output<string>;

    /**
     * Create a MetadataSchema resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetadataSchemaArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.metadataStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'metadataStoreId'");
            }
            if ((!args || args.schema === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schema'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["metadataSchemaId"] = args ? args.metadataSchemaId : undefined;
            resourceInputs["metadataStoreId"] = args ? args.metadataStoreId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["schema"] = args ? args.schema : undefined;
            resourceInputs["schemaType"] = args ? args.schemaType : undefined;
            resourceInputs["schemaVersion"] = args ? args.schemaVersion : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["metadataSchemaId"] = undefined /*out*/;
            resourceInputs["metadataStoreId"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["schema"] = undefined /*out*/;
            resourceInputs["schemaType"] = undefined /*out*/;
            resourceInputs["schemaVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "metadataStoreId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(MetadataSchema.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a MetadataSchema resource.
 */
export interface MetadataSchemaArgs {
    /**
     * Description of the Metadata Schema
     */
    description?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The {metadata_schema} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}/metadataSchemas/{metadataschema}` If not provided, the MetadataStore's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all MetadataSchemas in the parent Location. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting MetadataSchema.)
     */
    metadataSchemaId?: pulumi.Input<string>;
    metadataStoreId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The raw YAML string representation of the MetadataSchema. The combination of [MetadataSchema.version] and the schema name given by `title` in [MetadataSchema.schema] must be unique within a MetadataStore. The schema is defined as an OpenAPI 3.0.2 [MetadataSchema Object](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/3.0.2.md#schemaObject)
     */
    schema: pulumi.Input<string>;
    /**
     * The type of the MetadataSchema. This is a property that identifies which metadata types will use the MetadataSchema.
     */
    schemaType?: pulumi.Input<enums.aiplatform.v1beta1.MetadataSchemaSchemaType>;
    /**
     * The version of the MetadataSchema. The version's format must match the following regular expression: `^[0-9]+.+.+$`, which would allow to order/compare different versions. Example: 1.0.0, 1.0.1, etc.
     */
    schemaVersion?: pulumi.Input<string>;
}
