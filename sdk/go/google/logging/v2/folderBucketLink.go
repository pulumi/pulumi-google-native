// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Asynchronously creates a linked dataset in BigQuery which makes it possible to use BigQuery to read the logs stored in the log bucket. A log bucket may currently only contain one link.
// Auto-naming is currently not supported for this resource.
type FolderBucketLink struct {
	pulumi.CustomResourceState

	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset BigQueryDatasetResponseOutput `pulumi:"bigqueryDataset"`
	BucketId        pulumi.StringOutput           `pulumi:"bucketId"`
	// The creation timestamp of the link.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes this link.The maximum length of the description is 8000 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	FolderId    pulumi.StringOutput `pulumi:"folderId"`
	// The resource lifecycle state.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
	LinkId   pulumi.StringOutput `pulumi:"linkId"`
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewFolderBucketLink registers a new resource with the given unique name, arguments, and options.
func NewFolderBucketLink(ctx *pulumi.Context,
	name string, args *FolderBucketLinkArgs, opts ...pulumi.ResourceOption) (*FolderBucketLink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.FolderId == nil {
		return nil, errors.New("invalid value for required argument 'FolderId'")
	}
	if args.LinkId == nil {
		return nil, errors.New("invalid value for required argument 'LinkId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bucketId",
		"folderId",
		"linkId",
		"location",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FolderBucketLink
	err := ctx.RegisterResource("google-native:logging/v2:FolderBucketLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderBucketLink gets an existing FolderBucketLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderBucketLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderBucketLinkState, opts ...pulumi.ResourceOption) (*FolderBucketLink, error) {
	var resource FolderBucketLink
	err := ctx.ReadResource("google-native:logging/v2:FolderBucketLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderBucketLink resources.
type folderBucketLinkState struct {
}

type FolderBucketLinkState struct {
}

func (FolderBucketLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderBucketLinkState)(nil)).Elem()
}

type folderBucketLinkArgs struct {
	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset *BigQueryDataset `pulumi:"bigqueryDataset"`
	BucketId        string           `pulumi:"bucketId"`
	// Describes this link.The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	FolderId    string  `pulumi:"folderId"`
	// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
	LinkId   string  `pulumi:"linkId"`
	Location *string `pulumi:"location"`
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FolderBucketLink resource.
type FolderBucketLinkArgs struct {
	// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
	BigqueryDataset BigQueryDatasetPtrInput
	BucketId        pulumi.StringInput
	// Describes this link.The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	FolderId    pulumi.StringInput
	// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
	LinkId   pulumi.StringInput
	Location pulumi.StringPtrInput
	// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
	Name pulumi.StringPtrInput
}

func (FolderBucketLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderBucketLinkArgs)(nil)).Elem()
}

type FolderBucketLinkInput interface {
	pulumi.Input

	ToFolderBucketLinkOutput() FolderBucketLinkOutput
	ToFolderBucketLinkOutputWithContext(ctx context.Context) FolderBucketLinkOutput
}

func (*FolderBucketLink) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderBucketLink)(nil)).Elem()
}

func (i *FolderBucketLink) ToFolderBucketLinkOutput() FolderBucketLinkOutput {
	return i.ToFolderBucketLinkOutputWithContext(context.Background())
}

func (i *FolderBucketLink) ToFolderBucketLinkOutputWithContext(ctx context.Context) FolderBucketLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderBucketLinkOutput)
}

type FolderBucketLinkOutput struct{ *pulumi.OutputState }

func (FolderBucketLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderBucketLink)(nil)).Elem()
}

func (o FolderBucketLinkOutput) ToFolderBucketLinkOutput() FolderBucketLinkOutput {
	return o
}

func (o FolderBucketLinkOutput) ToFolderBucketLinkOutputWithContext(ctx context.Context) FolderBucketLinkOutput {
	return o
}

// The information of a BigQuery Dataset. When a link is created, a BigQuery dataset is created along with it, in the same project as the LogBucket it's linked to. This dataset will also have BigQuery Views corresponding to the LogViews in the bucket.
func (o FolderBucketLinkOutput) BigqueryDataset() BigQueryDatasetResponseOutput {
	return o.ApplyT(func(v *FolderBucketLink) BigQueryDatasetResponseOutput { return v.BigqueryDataset }).(BigQueryDatasetResponseOutput)
}

func (o FolderBucketLinkOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The creation timestamp of the link.
func (o FolderBucketLinkOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Describes this link.The maximum length of the description is 8000 characters.
func (o FolderBucketLinkOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o FolderBucketLinkOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.FolderId }).(pulumi.StringOutput)
}

// The resource lifecycle state.
func (o FolderBucketLinkOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// Required. The ID to use for the link. The link_id can have up to 100 characters. A valid link_id must only have alphanumeric characters and underscores within it.
func (o FolderBucketLinkOutput) LinkId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.LinkId }).(pulumi.StringOutput)
}

func (o FolderBucketLinkOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the link. The name can have up to 100 characters. A valid link id (at the end of the link name) must only have alphanumeric characters and underscores within it. "projects/[PROJECT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "organizations/[ORGANIZATION_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "billingAccounts/[BILLING_ACCOUNT_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" "folders/[FOLDER_ID]/locations/[LOCATION_ID]/buckets/[BUCKET_ID]/links/[LINK_ID]" For example:`projects/my-project/locations/global/buckets/my-bucket/links/my_link
func (o FolderBucketLinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderBucketLink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderBucketLinkInput)(nil)).Elem(), &FolderBucketLink{})
	pulumi.RegisterOutputType(FolderBucketLinkOutput{})
}