// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a view over log entries in a log bucket. A bucket may contain a maximum of 30 views.
 * Auto-naming is currently not supported for this resource.
 */
export class BillingAccountBucketView extends pulumi.CustomResource {
    /**
     * Get an existing BillingAccountBucketView resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BillingAccountBucketView {
        return new BillingAccountBucketView(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:logging/v2:BillingAccountBucketView';

    /**
     * Returns true if the given object is an instance of BillingAccountBucketView.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BillingAccountBucketView {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BillingAccountBucketView.__pulumiType;
    }

    public readonly billingAccountId!: pulumi.Output<string>;
    public readonly bucketId!: pulumi.Output<string>;
    /**
     * The creation timestamp of the view.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Describes this view.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
     */
    public readonly filter!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The last update timestamp of the view.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Required. A client-assigned identifier such as "my-view". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
     */
    public readonly viewId!: pulumi.Output<string>;

    /**
     * Create a BillingAccountBucketView resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BillingAccountBucketViewArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.billingAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingAccountId'");
            }
            if ((!args || args.bucketId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketId'");
            }
            if ((!args || args.viewId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'viewId'");
            }
            resourceInputs["billingAccountId"] = args ? args.billingAccountId : undefined;
            resourceInputs["bucketId"] = args ? args.bucketId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["viewId"] = args ? args.viewId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["billingAccountId"] = undefined /*out*/;
            resourceInputs["bucketId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["filter"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["viewId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["billingAccountId", "bucketId", "location", "viewId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(BillingAccountBucketView.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BillingAccountBucketView resource.
 */
export interface BillingAccountBucketViewArgs {
    billingAccountId: pulumi.Input<string>;
    bucketId: pulumi.Input<string>;
    /**
     * Describes this view.
     */
    description?: pulumi.Input<string>;
    /**
     * Filter that restricts which log entries in a bucket are visible in this view.Filters are restricted to be a logical AND of ==/!= of any of the following: originating project/folder/organization/billing account. resource type log idFor example:SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout")
     */
    filter?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the view.For example:projects/my-project/locations/global/buckets/my-bucket/views/my-view
     */
    name?: pulumi.Input<string>;
    /**
     * Required. A client-assigned identifier such as "my-view". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods.
     */
    viewId: pulumi.Input<string>;
}