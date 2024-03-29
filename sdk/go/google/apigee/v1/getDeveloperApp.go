// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the details for a developer app.
func LookupDeveloperApp(ctx *pulumi.Context, args *LookupDeveloperAppArgs, opts ...pulumi.InvokeOption) (*LookupDeveloperAppResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDeveloperAppResult
	err := ctx.Invoke("google-native:apigee/v1:getDeveloperApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeveloperAppArgs struct {
	AppId          string  `pulumi:"appId"`
	DeveloperId    string  `pulumi:"developerId"`
	Entity         *string `pulumi:"entity"`
	OrganizationId string  `pulumi:"organizationId"`
	Query          *string `pulumi:"query"`
}

type LookupDeveloperAppResult struct {
	// List of API products associated with the developer app.
	ApiProducts []string `pulumi:"apiProducts"`
	// Developer app family.
	AppFamily string `pulumi:"appFamily"`
	// ID of the developer app.
	AppId string `pulumi:"appId"`
	// List of attributes for the developer app.
	Attributes []GoogleCloudApigeeV1AttributeResponse `pulumi:"attributes"`
	// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
	CallbackUrl string `pulumi:"callbackUrl"`
	// Time the developer app was created in milliseconds since epoch.
	CreatedAt string `pulumi:"createdAt"`
	// Set of credentials for the developer app consisting of the consumer key/secret pairs associated with the API products.
	Credentials []GoogleCloudApigeeV1CredentialResponse `pulumi:"credentials"`
	// ID of the developer.
	DeveloperId string `pulumi:"developerId"`
	// Expiration time, in milliseconds, for the consumer key that is generated for the developer app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
	KeyExpiresIn string `pulumi:"keyExpiresIn"`
	// Time the developer app was modified in milliseconds since epoch.
	LastModifiedAt string `pulumi:"lastModifiedAt"`
	// Name of the developer app.
	Name string `pulumi:"name"`
	// Scopes to apply to the developer app. The specified scopes must already exist for the API product that you associate with the developer app.
	Scopes []string `pulumi:"scopes"`
	// Status of the credential. Valid values include `approved` or `revoked`.
	Status string `pulumi:"status"`
}

func LookupDeveloperAppOutput(ctx *pulumi.Context, args LookupDeveloperAppOutputArgs, opts ...pulumi.InvokeOption) LookupDeveloperAppResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeveloperAppResult, error) {
			args := v.(LookupDeveloperAppArgs)
			r, err := LookupDeveloperApp(ctx, &args, opts...)
			var s LookupDeveloperAppResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDeveloperAppResultOutput)
}

type LookupDeveloperAppOutputArgs struct {
	AppId          pulumi.StringInput    `pulumi:"appId"`
	DeveloperId    pulumi.StringInput    `pulumi:"developerId"`
	Entity         pulumi.StringPtrInput `pulumi:"entity"`
	OrganizationId pulumi.StringInput    `pulumi:"organizationId"`
	Query          pulumi.StringPtrInput `pulumi:"query"`
}

func (LookupDeveloperAppOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeveloperAppArgs)(nil)).Elem()
}

type LookupDeveloperAppResultOutput struct{ *pulumi.OutputState }

func (LookupDeveloperAppResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeveloperAppResult)(nil)).Elem()
}

func (o LookupDeveloperAppResultOutput) ToLookupDeveloperAppResultOutput() LookupDeveloperAppResultOutput {
	return o
}

func (o LookupDeveloperAppResultOutput) ToLookupDeveloperAppResultOutputWithContext(ctx context.Context) LookupDeveloperAppResultOutput {
	return o
}

// List of API products associated with the developer app.
func (o LookupDeveloperAppResultOutput) ApiProducts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) []string { return v.ApiProducts }).(pulumi.StringArrayOutput)
}

// Developer app family.
func (o LookupDeveloperAppResultOutput) AppFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.AppFamily }).(pulumi.StringOutput)
}

// ID of the developer app.
func (o LookupDeveloperAppResultOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.AppId }).(pulumi.StringOutput)
}

// List of attributes for the developer app.
func (o LookupDeveloperAppResultOutput) Attributes() GoogleCloudApigeeV1AttributeResponseArrayOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) []GoogleCloudApigeeV1AttributeResponse { return v.Attributes }).(GoogleCloudApigeeV1AttributeResponseArrayOutput)
}

// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to developer apps.
func (o LookupDeveloperAppResultOutput) CallbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.CallbackUrl }).(pulumi.StringOutput)
}

// Time the developer app was created in milliseconds since epoch.
func (o LookupDeveloperAppResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Set of credentials for the developer app consisting of the consumer key/secret pairs associated with the API products.
func (o LookupDeveloperAppResultOutput) Credentials() GoogleCloudApigeeV1CredentialResponseArrayOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) []GoogleCloudApigeeV1CredentialResponse { return v.Credentials }).(GoogleCloudApigeeV1CredentialResponseArrayOutput)
}

// ID of the developer.
func (o LookupDeveloperAppResultOutput) DeveloperId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.DeveloperId }).(pulumi.StringOutput)
}

// Expiration time, in milliseconds, for the consumer key that is generated for the developer app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
func (o LookupDeveloperAppResultOutput) KeyExpiresIn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.KeyExpiresIn }).(pulumi.StringOutput)
}

// Time the developer app was modified in milliseconds since epoch.
func (o LookupDeveloperAppResultOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

// Name of the developer app.
func (o LookupDeveloperAppResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.Name }).(pulumi.StringOutput)
}

// Scopes to apply to the developer app. The specified scopes must already exist for the API product that you associate with the developer app.
func (o LookupDeveloperAppResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Status of the credential. Valid values include `approved` or `revoked`.
func (o LookupDeveloperAppResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDeveloperAppResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeveloperAppResultOutput{})
}
