// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a link.
 */
export function getBillingAccountBucketLink(args: GetBillingAccountBucketLinkArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingAccountBucketLinkResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:logging/v2:getBillingAccountBucketLink", {
        "billingAccountId": args.billingAccountId,
        "bucketId": args.bucketId,
        "linkId": args.linkId,
        "location": args.location,
    }, opts);
}

export interface GetBillingAccountBucketLinkArgs {
    billingAccountId: string;
    bucketId: string;
    linkId: string;
    location: string;
}

export interface GetBillingAccountBucketLinkResult {
    /**
     * The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
     */
    readonly bigqueryDataset: outputs.logging.v2.BigQueryDatasetResponse;
    /**
     * The creation timestamp of the link.
     */
    readonly createTime: string;
    /**
     * Describes this link.The maximum length of the description is 8000 characters.
     */
    readonly description: string;
    /**
     * The resource lifecycle state.
     */
    readonly lifecycleState: string;
    /**
     * The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
     */
    readonly name: string;
}
/**
 * Gets a link.
 */
export function getBillingAccountBucketLinkOutput(args: GetBillingAccountBucketLinkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBillingAccountBucketLinkResult> {
    return pulumi.output(args).apply((a: any) => getBillingAccountBucketLink(a, opts))
}

export interface GetBillingAccountBucketLinkOutputArgs {
    billingAccountId: pulumi.Input<string>;
    bucketId: pulumi.Input<string>;
    linkId: pulumi.Input<string>;
    location: pulumi.Input<string>;
}