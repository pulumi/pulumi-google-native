// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single FeatureView.
func LookupFeatureView(ctx *pulumi.Context, args *LookupFeatureViewArgs, opts ...pulumi.InvokeOption) (*LookupFeatureViewResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeatureViewResult
	err := ctx.Invoke("google-native:aiplatform/v1:getFeatureView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFeatureViewArgs struct {
	FeatureOnlineStoreId string  `pulumi:"featureOnlineStoreId"`
	FeatureViewId        string  `pulumi:"featureViewId"`
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
}

type LookupFeatureViewResult struct {
	// Optional. Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
	BigQuerySource GoogleCloudAiplatformV1FeatureViewBigQuerySourceResponse `pulumi:"bigQuerySource"`
	// Timestamp when this FeatureView was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag string `pulumi:"etag"`
	// Optional. Configures the features from a Feature Registry source that need to be loaded onto the FeatureOnlineStore.
	FeatureRegistrySource GoogleCloudAiplatformV1FeatureViewFeatureRegistrySourceResponse `pulumi:"featureRegistrySource"`
	// Optional. The labels with user-defined metadata to organize your FeatureViews. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information on and examples of labels. No more than 64 user labels can be associated with one FeatureOnlineStore(System labels are excluded)." System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
	Labels map[string]string `pulumi:"labels"`
	// Name of the FeatureView. Format: `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}`
	Name string `pulumi:"name"`
	// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
	SyncConfig GoogleCloudAiplatformV1FeatureViewSyncConfigResponse `pulumi:"syncConfig"`
	// Timestamp when this FeatureView was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupFeatureViewOutput(ctx *pulumi.Context, args LookupFeatureViewOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureViewResult, error) {
			args := v.(LookupFeatureViewArgs)
			r, err := LookupFeatureView(ctx, &args, opts...)
			var s LookupFeatureViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeatureViewResultOutput)
}

type LookupFeatureViewOutputArgs struct {
	FeatureOnlineStoreId pulumi.StringInput    `pulumi:"featureOnlineStoreId"`
	FeatureViewId        pulumi.StringInput    `pulumi:"featureViewId"`
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupFeatureViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureViewArgs)(nil)).Elem()
}

type LookupFeatureViewResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureViewResult)(nil)).Elem()
}

func (o LookupFeatureViewResultOutput) ToLookupFeatureViewResultOutput() LookupFeatureViewResultOutput {
	return o
}

func (o LookupFeatureViewResultOutput) ToLookupFeatureViewResultOutputWithContext(ctx context.Context) LookupFeatureViewResultOutput {
	return o
}

// Optional. Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
func (o LookupFeatureViewResultOutput) BigQuerySource() GoogleCloudAiplatformV1FeatureViewBigQuerySourceResponseOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) GoogleCloudAiplatformV1FeatureViewBigQuerySourceResponse {
		return v.BigQuerySource
	}).(GoogleCloudAiplatformV1FeatureViewBigQuerySourceResponseOutput)
}

// Timestamp when this FeatureView was created.
func (o LookupFeatureViewResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o LookupFeatureViewResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Optional. Configures the features from a Feature Registry source that need to be loaded onto the FeatureOnlineStore.
func (o LookupFeatureViewResultOutput) FeatureRegistrySource() GoogleCloudAiplatformV1FeatureViewFeatureRegistrySourceResponseOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) GoogleCloudAiplatformV1FeatureViewFeatureRegistrySourceResponse {
		return v.FeatureRegistrySource
	}).(GoogleCloudAiplatformV1FeatureViewFeatureRegistrySourceResponseOutput)
}

// Optional. The labels with user-defined metadata to organize your FeatureViews. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information on and examples of labels. No more than 64 user labels can be associated with one FeatureOnlineStore(System labels are excluded)." System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
func (o LookupFeatureViewResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the FeatureView. Format: `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}`
func (o LookupFeatureViewResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
func (o LookupFeatureViewResultOutput) SyncConfig() GoogleCloudAiplatformV1FeatureViewSyncConfigResponseOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) GoogleCloudAiplatformV1FeatureViewSyncConfigResponse {
		return v.SyncConfig
	}).(GoogleCloudAiplatformV1FeatureViewSyncConfigResponseOutput)
}

// Timestamp when this FeatureView was last updated.
func (o LookupFeatureViewResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureViewResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureViewResultOutput{})
}
