// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single TrustConfig.
func LookupTrustConfig(ctx *pulumi.Context, args *LookupTrustConfigArgs, opts ...pulumi.InvokeOption) (*LookupTrustConfigResult, error) {
	var rv LookupTrustConfigResult
	err := ctx.Invoke("google-native:certificatemanager/v1:getTrustConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrustConfigArgs struct {
	Location      string  `pulumi:"location"`
	Project       *string `pulumi:"project"`
	TrustConfigId string  `pulumi:"trustConfigId"`
}

type LookupTrustConfigResult struct {
	// The creation timestamp of a TrustConfig.
	CreateTime string `pulumi:"createTime"`
	// One or more paragraphs of text description of a TrustConfig.
	Description string `pulumi:"description"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// Set of labels associated with a TrustConfig.
	Labels map[string]string `pulumi:"labels"`
	// A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*/locations/*/trustConfigs/*`.
	Name string `pulumi:"name"`
	// Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
	TrustStores []TrustStoreResponse `pulumi:"trustStores"`
	// The last update timestamp of a TrustConfig.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupTrustConfigOutput(ctx *pulumi.Context, args LookupTrustConfigOutputArgs, opts ...pulumi.InvokeOption) LookupTrustConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrustConfigResult, error) {
			args := v.(LookupTrustConfigArgs)
			r, err := LookupTrustConfig(ctx, &args, opts...)
			var s LookupTrustConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrustConfigResultOutput)
}

type LookupTrustConfigOutputArgs struct {
	Location      pulumi.StringInput    `pulumi:"location"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	TrustConfigId pulumi.StringInput    `pulumi:"trustConfigId"`
}

func (LookupTrustConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustConfigArgs)(nil)).Elem()
}

type LookupTrustConfigResultOutput struct{ *pulumi.OutputState }

func (LookupTrustConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrustConfigResult)(nil)).Elem()
}

func (o LookupTrustConfigResultOutput) ToLookupTrustConfigResultOutput() LookupTrustConfigResultOutput {
	return o
}

func (o LookupTrustConfigResultOutput) ToLookupTrustConfigResultOutputWithContext(ctx context.Context) LookupTrustConfigResultOutput {
	return o
}

// The creation timestamp of a TrustConfig.
func (o LookupTrustConfigResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// One or more paragraphs of text description of a TrustConfig.
func (o LookupTrustConfigResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) string { return v.Description }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o LookupTrustConfigResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Set of labels associated with a TrustConfig.
func (o LookupTrustConfigResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A user-defined name of the trust config. TrustConfig names must be unique globally and match pattern `projects/*/locations/*/trustConfigs/*`.
func (o LookupTrustConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// Set of trust stores to perform validation against. This field is supported when TrustConfig is configured with Load Balancers, currently not supported for SPIFFE certificate validation. Only one TrustStore specified is currently allowed.
func (o LookupTrustConfigResultOutput) TrustStores() TrustStoreResponseArrayOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) []TrustStoreResponse { return v.TrustStores }).(TrustStoreResponseArrayOutput)
}

// The last update timestamp of a TrustConfig.
func (o LookupTrustConfigResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTrustConfigResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrustConfigResultOutput{})
}