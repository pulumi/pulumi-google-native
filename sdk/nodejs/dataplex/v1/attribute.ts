// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Create a DataAttribute resource.
 * Auto-naming is currently not supported for this resource.
 */
export class Attribute extends pulumi.CustomResource {
    /**
     * Get an existing Attribute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Attribute {
        return new Attribute(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataplex/v1:Attribute';

    /**
     * Returns true if the given object is an instance of Attribute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Attribute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Attribute.__pulumiType;
    }

    /**
     * The number of child attributes present for this attribute.
     */
    public /*out*/ readonly attributeCount!: pulumi.Output<number>;
    /**
     * The time when the DataAttribute was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Specified when applied to data stored on the resource (eg: rows, columns in BigQuery Tables).
     */
    public readonly dataAccessSpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1DataAccessSpecResponse>;
    /**
     * Required. DataAttribute identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the DataTaxonomy.
     */
    public readonly dataAttributeId!: pulumi.Output<string>;
    public readonly dataTaxonomyId!: pulumi.Output<string>;
    /**
     * Optional. Description of the DataAttribute.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. User friendly display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Optional. User-defined labels for the DataAttribute.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The relative resource name of the dataAttribute, of the form: projects/{project_number}/locations/{location_id}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. The ID of the parent DataAttribute resource, should belong to the same data taxonomy. Circular dependency in parent chain is not valid. Maximum depth of the hierarchy allowed is 4. a -> b -> c -> d -> e, depth = 4
     */
    public readonly parentId!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. Specified when applied to a resource (eg: Cloud Storage bucket, BigQuery dataset, BigQuery table).
     */
    public readonly resourceAccessSpec!: pulumi.Output<outputs.dataplex.v1.GoogleCloudDataplexV1ResourceAccessSpecResponse>;
    /**
     * System generated globally unique ID for the DataAttribute. This ID will be different if the DataAttribute is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the DataAttribute was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Optional. Only validate the request, but do not perform mutations. The default is false.
     */
    public readonly validateOnly!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Attribute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttributeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dataAttributeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataAttributeId'");
            }
            if ((!args || args.dataTaxonomyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataTaxonomyId'");
            }
            resourceInputs["dataAccessSpec"] = args ? args.dataAccessSpec : undefined;
            resourceInputs["dataAttributeId"] = args ? args.dataAttributeId : undefined;
            resourceInputs["dataTaxonomyId"] = args ? args.dataTaxonomyId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["parentId"] = args ? args.parentId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["resourceAccessSpec"] = args ? args.resourceAccessSpec : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
            resourceInputs["attributeCount"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["attributeCount"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataAccessSpec"] = undefined /*out*/;
            resourceInputs["dataAttributeId"] = undefined /*out*/;
            resourceInputs["dataTaxonomyId"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parentId"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["resourceAccessSpec"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["validateOnly"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["dataAttributeId", "dataTaxonomyId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Attribute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Attribute resource.
 */
export interface AttributeArgs {
    /**
     * Optional. Specified when applied to data stored on the resource (eg: rows, columns in BigQuery Tables).
     */
    dataAccessSpec?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1DataAccessSpecArgs>;
    /**
     * Required. DataAttribute identifier. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must be between 1-63 characters. * Must end with a number or a letter. * Must be unique within the DataTaxonomy.
     */
    dataAttributeId: pulumi.Input<string>;
    dataTaxonomyId: pulumi.Input<string>;
    /**
     * Optional. Description of the DataAttribute.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Optional. User-defined labels for the DataAttribute.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. The ID of the parent DataAttribute resource, should belong to the same data taxonomy. Circular dependency in parent chain is not valid. Maximum depth of the hierarchy allowed is 4. a -> b -> c -> d -> e, depth = 4
     */
    parentId?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. Specified when applied to a resource (eg: Cloud Storage bucket, BigQuery dataset, BigQuery table).
     */
    resourceAccessSpec?: pulumi.Input<inputs.dataplex.v1.GoogleCloudDataplexV1ResourceAccessSpecArgs>;
    /**
     * Optional. Only validate the request, but do not perform mutations. The default is false.
     */
    validateOnly?: pulumi.Input<boolean>;
}