// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Logging.V2
{
    public static class GetBillingAccountBucketLink
    {
        /// <summary>
        /// Gets a link.
        /// </summary>
        public static Task<GetBillingAccountBucketLinkResult> InvokeAsync(GetBillingAccountBucketLinkArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBillingAccountBucketLinkResult>("google-native:logging/v2:getBillingAccountBucketLink", args ?? new GetBillingAccountBucketLinkArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a link.
        /// </summary>
        public static Output<GetBillingAccountBucketLinkResult> Invoke(GetBillingAccountBucketLinkInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBillingAccountBucketLinkResult>("google-native:logging/v2:getBillingAccountBucketLink", args ?? new GetBillingAccountBucketLinkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBillingAccountBucketLinkArgs : global::Pulumi.InvokeArgs
    {
        [Input("billingAccountId", required: true)]
        public string BillingAccountId { get; set; } = null!;

        [Input("bucketId", required: true)]
        public string BucketId { get; set; } = null!;

        [Input("linkId", required: true)]
        public string LinkId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        public GetBillingAccountBucketLinkArgs()
        {
        }
        public static new GetBillingAccountBucketLinkArgs Empty => new GetBillingAccountBucketLinkArgs();
    }

    public sealed class GetBillingAccountBucketLinkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("billingAccountId", required: true)]
        public Input<string> BillingAccountId { get; set; } = null!;

        [Input("bucketId", required: true)]
        public Input<string> BucketId { get; set; } = null!;

        [Input("linkId", required: true)]
        public Input<string> LinkId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        public GetBillingAccountBucketLinkInvokeArgs()
        {
        }
        public static new GetBillingAccountBucketLinkInvokeArgs Empty => new GetBillingAccountBucketLinkInvokeArgs();
    }


    [OutputType]
    public sealed class GetBillingAccountBucketLinkResult
    {
        /// <summary>
        /// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
        /// </summary>
        public readonly Outputs.BigQueryDatasetResponse BigqueryDataset;
        /// <summary>
        /// The creation timestamp of the link.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Describes this link.The maximum length of the description is 8000 characters.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The resource lifecycle state.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetBillingAccountBucketLinkResult(
            Outputs.BigQueryDatasetResponse bigqueryDataset,

            string createTime,

            string description,

            string lifecycleState,

            string name)
        {
            BigqueryDataset = bigqueryDataset;
            CreateTime = createTime;
            Description = description;
            LifecycleState = lifecycleState;
            Name = name;
        }
    }
}