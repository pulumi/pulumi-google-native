// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single ContactCenter.
func LookupContactCenter(ctx *pulumi.Context, args *LookupContactCenterArgs, opts ...pulumi.InvokeOption) (*LookupContactCenterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContactCenterResult
	err := ctx.Invoke("google-native:contactcenteraiplatform/v1alpha1:getContactCenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContactCenterArgs struct {
	ContactCenterId string  `pulumi:"contactCenterId"`
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
}

type LookupContactCenterResult struct {
	// Optional. Info about the first admin user, such as given name and family name.
	AdminUser AdminUserResponse `pulumi:"adminUser"`
	// Optional. Whether to enable users to be created in the CCAIP-instance concurrently to having users in Cloud identity
	CcaipManagedUsers bool `pulumi:"ccaipManagedUsers"`
	// [Output only] Create time stamp
	CreateTime string `pulumi:"createTime"`
	// Immutable. At least 2 and max 16 char long, must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
	CustomerDomainPrefix string `pulumi:"customerDomainPrefix"`
	// A user friendly name for the ContactCenter.
	DisplayName string `pulumi:"displayName"`
	// The configuration of this instance, it is currently immutable once created.
	InstanceConfig InstanceConfigResponse `pulumi:"instanceConfig"`
	// Labels as key value pairs
	Labels map[string]string `pulumi:"labels"`
	// name of resource
	Name string `pulumi:"name"`
	// Optional. Params that sets up Google as IdP.
	SamlParams SAMLParamsResponse `pulumi:"samlParams"`
	// The state of this contact center.
	State string `pulumi:"state"`
	// [Output only] Update time stamp
	UpdateTime string `pulumi:"updateTime"`
	// URIs to access the deployed ContactCenters.
	Uris URIsResponse `pulumi:"uris"`
	// Optional. Email address of the first admin user.
	UserEmail string `pulumi:"userEmail"`
}

func LookupContactCenterOutput(ctx *pulumi.Context, args LookupContactCenterOutputArgs, opts ...pulumi.InvokeOption) LookupContactCenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContactCenterResult, error) {
			args := v.(LookupContactCenterArgs)
			r, err := LookupContactCenter(ctx, &args, opts...)
			var s LookupContactCenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContactCenterResultOutput)
}

type LookupContactCenterOutputArgs struct {
	ContactCenterId pulumi.StringInput    `pulumi:"contactCenterId"`
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupContactCenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactCenterArgs)(nil)).Elem()
}

type LookupContactCenterResultOutput struct{ *pulumi.OutputState }

func (LookupContactCenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContactCenterResult)(nil)).Elem()
}

func (o LookupContactCenterResultOutput) ToLookupContactCenterResultOutput() LookupContactCenterResultOutput {
	return o
}

func (o LookupContactCenterResultOutput) ToLookupContactCenterResultOutputWithContext(ctx context.Context) LookupContactCenterResultOutput {
	return o
}

// Optional. Info about the first admin user, such as given name and family name.
func (o LookupContactCenterResultOutput) AdminUser() AdminUserResponseOutput {
	return o.ApplyT(func(v LookupContactCenterResult) AdminUserResponse { return v.AdminUser }).(AdminUserResponseOutput)
}

// Optional. Whether to enable users to be created in the CCAIP-instance concurrently to having users in Cloud identity
func (o LookupContactCenterResultOutput) CcaipManagedUsers() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupContactCenterResult) bool { return v.CcaipManagedUsers }).(pulumi.BoolOutput)
}

// [Output only] Create time stamp
func (o LookupContactCenterResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Immutable. At least 2 and max 16 char long, must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt).
func (o LookupContactCenterResultOutput) CustomerDomainPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.CustomerDomainPrefix }).(pulumi.StringOutput)
}

// A user friendly name for the ContactCenter.
func (o LookupContactCenterResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The configuration of this instance, it is currently immutable once created.
func (o LookupContactCenterResultOutput) InstanceConfig() InstanceConfigResponseOutput {
	return o.ApplyT(func(v LookupContactCenterResult) InstanceConfigResponse { return v.InstanceConfig }).(InstanceConfigResponseOutput)
}

// Labels as key value pairs
func (o LookupContactCenterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupContactCenterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// name of resource
func (o LookupContactCenterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Params that sets up Google as IdP.
func (o LookupContactCenterResultOutput) SamlParams() SAMLParamsResponseOutput {
	return o.ApplyT(func(v LookupContactCenterResult) SAMLParamsResponse { return v.SamlParams }).(SAMLParamsResponseOutput)
}

// The state of this contact center.
func (o LookupContactCenterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.State }).(pulumi.StringOutput)
}

// [Output only] Update time stamp
func (o LookupContactCenterResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// URIs to access the deployed ContactCenters.
func (o LookupContactCenterResultOutput) Uris() URIsResponseOutput {
	return o.ApplyT(func(v LookupContactCenterResult) URIsResponse { return v.Uris }).(URIsResponseOutput)
}

// Optional. Email address of the first admin user.
func (o LookupContactCenterResultOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContactCenterResult) string { return v.UserEmail }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContactCenterResultOutput{})
}