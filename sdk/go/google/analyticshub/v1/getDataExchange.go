// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the details of a data exchange.
func LookupDataExchange(ctx *pulumi.Context, args *LookupDataExchangeArgs, opts ...pulumi.InvokeOption) (*LookupDataExchangeResult, error) {
	var rv LookupDataExchangeResult
	err := ctx.Invoke("google-native:analyticshub/v1:getDataExchange", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataExchangeArgs struct {
	DataExchangeId string  `pulumi:"dataExchangeId"`
	Location       string  `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

type LookupDataExchangeResult struct {
	// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description string `pulumi:"description"`
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName string `pulumi:"displayName"`
	// Optional. Documentation describing the data exchange.
	Documentation string `pulumi:"documentation"`
	// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the content of the fields are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
	Icon string `pulumi:"icon"`
	// Number of listings contained in the data exchange.
	ListingCount int `pulumi:"listingCount"`
	// The resource name of the data exchange. e.g. `projects/myproject/locations/US/dataExchanges/123`.
	Name string `pulumi:"name"`
	// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
	PrimaryContact string `pulumi:"primaryContact"`
}

func LookupDataExchangeOutput(ctx *pulumi.Context, args LookupDataExchangeOutputArgs, opts ...pulumi.InvokeOption) LookupDataExchangeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataExchangeResult, error) {
			args := v.(LookupDataExchangeArgs)
			r, err := LookupDataExchange(ctx, &args, opts...)
			var s LookupDataExchangeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataExchangeResultOutput)
}

type LookupDataExchangeOutputArgs struct {
	DataExchangeId pulumi.StringInput    `pulumi:"dataExchangeId"`
	Location       pulumi.StringInput    `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDataExchangeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataExchangeArgs)(nil)).Elem()
}

type LookupDataExchangeResultOutput struct{ *pulumi.OutputState }

func (LookupDataExchangeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataExchangeResult)(nil)).Elem()
}

func (o LookupDataExchangeResultOutput) ToLookupDataExchangeResultOutput() LookupDataExchangeResultOutput {
	return o
}

func (o LookupDataExchangeResultOutput) ToLookupDataExchangeResultOutputWithContext(ctx context.Context) LookupDataExchangeResultOutput {
	return o
}

// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
func (o LookupDataExchangeResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) string { return v.Description }).(pulumi.StringOutput)
}

// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
func (o LookupDataExchangeResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Documentation describing the data exchange.
func (o LookupDataExchangeResultOutput) Documentation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) string { return v.Documentation }).(pulumi.StringOutput)
}

// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the content of the fields are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
func (o LookupDataExchangeResultOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) string { return v.Icon }).(pulumi.StringOutput)
}

// Number of listings contained in the data exchange.
func (o LookupDataExchangeResultOutput) ListingCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) int { return v.ListingCount }).(pulumi.IntOutput)
}

// The resource name of the data exchange. e.g. `projects/myproject/locations/US/dataExchanges/123`.
func (o LookupDataExchangeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
func (o LookupDataExchangeResultOutput) PrimaryContact() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataExchangeResult) string { return v.PrimaryContact }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataExchangeResultOutput{})
}