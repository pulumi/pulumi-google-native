// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new data exchange.
// Auto-naming is currently not supported for this resource.
type DataExchange struct {
	pulumi.CustomResourceState

	// Required. The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Max length: 100 bytes.
	DataExchangeId pulumi.StringOutput `pulumi:"dataExchangeId"`
	// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description pulumi.StringOutput `pulumi:"description"`
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Documentation describing the data exchange.
	Documentation pulumi.StringOutput `pulumi:"documentation"`
	// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the content of the fields are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
	Icon pulumi.StringOutput `pulumi:"icon"`
	// Number of listings contained in the data exchange.
	ListingCount pulumi.IntOutput    `pulumi:"listingCount"`
	Location     pulumi.StringOutput `pulumi:"location"`
	// The resource name of the data exchange. e.g. `projects/myproject/locations/US/dataExchanges/123`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
	PrimaryContact pulumi.StringOutput `pulumi:"primaryContact"`
	Project        pulumi.StringOutput `pulumi:"project"`
}

// NewDataExchange registers a new resource with the given unique name, arguments, and options.
func NewDataExchange(ctx *pulumi.Context,
	name string, args *DataExchangeArgs, opts ...pulumi.ResourceOption) (*DataExchange, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataExchangeId == nil {
		return nil, errors.New("invalid value for required argument 'DataExchangeId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dataExchangeId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource DataExchange
	err := ctx.RegisterResource("google-native:analyticshub/v1:DataExchange", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataExchange gets an existing DataExchange resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataExchange(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataExchangeState, opts ...pulumi.ResourceOption) (*DataExchange, error) {
	var resource DataExchange
	err := ctx.ReadResource("google-native:analyticshub/v1:DataExchange", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataExchange resources.
type dataExchangeState struct {
}

type DataExchangeState struct {
}

func (DataExchangeState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeState)(nil)).Elem()
}

type dataExchangeArgs struct {
	// Required. The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Max length: 100 bytes.
	DataExchangeId string `pulumi:"dataExchangeId"`
	// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description *string `pulumi:"description"`
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName string `pulumi:"displayName"`
	// Optional. Documentation describing the data exchange.
	Documentation *string `pulumi:"documentation"`
	// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the content of the fields are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
	Icon     *string `pulumi:"icon"`
	Location *string `pulumi:"location"`
	// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
	PrimaryContact *string `pulumi:"primaryContact"`
	Project        *string `pulumi:"project"`
}

// The set of arguments for constructing a DataExchange resource.
type DataExchangeArgs struct {
	// Required. The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Max length: 100 bytes.
	DataExchangeId pulumi.StringInput
	// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
	Description pulumi.StringPtrInput
	// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
	DisplayName pulumi.StringInput
	// Optional. Documentation describing the data exchange.
	Documentation pulumi.StringPtrInput
	// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the content of the fields are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
	Icon     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
	PrimaryContact pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
}

func (DataExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataExchangeArgs)(nil)).Elem()
}

type DataExchangeInput interface {
	pulumi.Input

	ToDataExchangeOutput() DataExchangeOutput
	ToDataExchangeOutputWithContext(ctx context.Context) DataExchangeOutput
}

func (*DataExchange) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchange)(nil)).Elem()
}

func (i *DataExchange) ToDataExchangeOutput() DataExchangeOutput {
	return i.ToDataExchangeOutputWithContext(context.Background())
}

func (i *DataExchange) ToDataExchangeOutputWithContext(ctx context.Context) DataExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataExchangeOutput)
}

type DataExchangeOutput struct{ *pulumi.OutputState }

func (DataExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataExchange)(nil)).Elem()
}

func (o DataExchangeOutput) ToDataExchangeOutput() DataExchangeOutput {
	return o
}

func (o DataExchangeOutput) ToDataExchangeOutputWithContext(ctx context.Context) DataExchangeOutput {
	return o
}

// Required. The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces. Max length: 100 bytes.
func (o DataExchangeOutput) DataExchangeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.DataExchangeId }).(pulumi.StringOutput)
}

// Optional. Description of the data exchange. The description must not contain Unicode non-characters as well as C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
func (o DataExchangeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Human-readable display name of the data exchange. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&) and must not start or end with spaces. Default value is an empty string. Max length: 63 bytes.
func (o DataExchangeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Documentation describing the data exchange.
func (o DataExchangeOutput) Documentation() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Documentation }).(pulumi.StringOutput)
}

// Optional. Base64 encoded image representing the data exchange. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the content of the fields are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
func (o DataExchangeOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Icon }).(pulumi.StringOutput)
}

// Number of listings contained in the data exchange.
func (o DataExchangeOutput) ListingCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.IntOutput { return v.ListingCount }).(pulumi.IntOutput)
}

func (o DataExchangeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the data exchange. e.g. `projects/myproject/locations/US/dataExchanges/123`.
func (o DataExchangeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Email or URL of the primary point of contact of the data exchange. Max Length: 1000 bytes.
func (o DataExchangeOutput) PrimaryContact() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.PrimaryContact }).(pulumi.StringOutput)
}

func (o DataExchangeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataExchange) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataExchangeInput)(nil)).Elem(), &DataExchange{})
	pulumi.RegisterOutputType(DataExchangeOutput{})
}