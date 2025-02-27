// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new ArchiveDeployment.
 */
export class ArchiveDeployment extends pulumi.CustomResource {
    /**
     * Get an existing ArchiveDeployment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ArchiveDeployment {
        return new ArchiveDeployment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:ArchiveDeployment';

    /**
     * Returns true if the given object is an instance of ArchiveDeployment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ArchiveDeployment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ArchiveDeployment.__pulumiType;
    }

    /**
     * The time at which the Archive Deployment was created in milliseconds since the epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly environmentId!: pulumi.Output<string>;
    /**
     * Input only. The Google Cloud Storage signed URL returned from GenerateUploadUrl and used to upload the Archive zip file.
     */
    public readonly gcsUri!: pulumi.Output<string>;
    /**
     * User-supplied key-value pairs used to organize ArchiveDeployments. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the Archive Deployment in the following format: `organizations/{org}/environments/{env}/archiveDeployments/{id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A reference to the LRO that created this Archive Deployment in the following format: `organizations/{org}/operations/{id}`
     */
    public /*out*/ readonly operation!: pulumi.Output<string>;
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * The time at which the Archive Deployment was updated in milliseconds since the epoch.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a ArchiveDeployment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ArchiveDeploymentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["gcsUri"] = args ? args.gcsUri : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["operation"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["environmentId"] = undefined /*out*/;
            resourceInputs["gcsUri"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["operation"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["environmentId", "organizationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ArchiveDeployment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ArchiveDeployment resource.
 */
export interface ArchiveDeploymentArgs {
    environmentId: pulumi.Input<string>;
    /**
     * Input only. The Google Cloud Storage signed URL returned from GenerateUploadUrl and used to upload the Archive zip file.
     */
    gcsUri?: pulumi.Input<string>;
    /**
     * User-supplied key-value pairs used to organize ArchiveDeployments. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the Archive Deployment in the following format: `organizations/{org}/environments/{env}/archiveDeployments/{id}`.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
