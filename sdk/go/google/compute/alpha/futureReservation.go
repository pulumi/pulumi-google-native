// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Future Reservation.
type FutureReservation struct {
	pulumi.CustomResourceState

	// The creation timestamp for this future reservation in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the future reservation.
	Description pulumi.StringOutput `pulumi:"description"`
	// Type of the resource. Always compute#futureReservation for future reservations.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
	NamePrefix pulumi.StringOutput `pulumi:"namePrefix"`
	// Server-defined fully-qualified URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// List of Projects/Folders to share with.
	ShareSettings ShareSettingsResponseOutput `pulumi:"shareSettings"`
	// Future Reservation configuration to indicate instance properties and total count.
	SpecificSkuProperties FutureReservationSpecificSKUPropertiesResponseOutput `pulumi:"specificSkuProperties"`
	// [Output only] Status of the Future Reservation
	Status FutureReservationStatusResponseOutput `pulumi:"status"`
	// Time window for this Future Reservation.
	TimeWindow FutureReservationTimeWindowResponseOutput `pulumi:"timeWindow"`
	// URL of the Zone where this future reservation resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewFutureReservation registers a new resource with the given unique name, arguments, and options.
func NewFutureReservation(ctx *pulumi.Context,
	name string, args *FutureReservationArgs, opts ...pulumi.ResourceOption) (*FutureReservation, error) {
	if args == nil {
		args = &FutureReservationArgs{}
	}

	var resource FutureReservation
	err := ctx.RegisterResource("google-native:compute/alpha:FutureReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFutureReservation gets an existing FutureReservation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFutureReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FutureReservationState, opts ...pulumi.ResourceOption) (*FutureReservation, error) {
	var resource FutureReservation
	err := ctx.ReadResource("google-native:compute/alpha:FutureReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FutureReservation resources.
type futureReservationState struct {
}

type FutureReservationState struct {
}

func (FutureReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*futureReservationState)(nil)).Elem()
}

type futureReservationArgs struct {
	// An optional description of this resource. Provide this property when you create the future reservation.
	Description *string `pulumi:"description"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
	NamePrefix *string `pulumi:"namePrefix"`
	Project    *string `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// List of Projects/Folders to share with.
	ShareSettings *ShareSettings `pulumi:"shareSettings"`
	// Future Reservation configuration to indicate instance properties and total count.
	SpecificSkuProperties *FutureReservationSpecificSKUProperties `pulumi:"specificSkuProperties"`
	// Time window for this Future Reservation.
	TimeWindow *FutureReservationTimeWindow `pulumi:"timeWindow"`
	Zone       *string                      `pulumi:"zone"`
}

// The set of arguments for constructing a FutureReservation resource.
type FutureReservationArgs struct {
	// An optional description of this resource. Provide this property when you create the future reservation.
	Description pulumi.StringPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
	NamePrefix pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// List of Projects/Folders to share with.
	ShareSettings ShareSettingsPtrInput
	// Future Reservation configuration to indicate instance properties and total count.
	SpecificSkuProperties FutureReservationSpecificSKUPropertiesPtrInput
	// Time window for this Future Reservation.
	TimeWindow FutureReservationTimeWindowPtrInput
	Zone       pulumi.StringPtrInput
}

func (FutureReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*futureReservationArgs)(nil)).Elem()
}

type FutureReservationInput interface {
	pulumi.Input

	ToFutureReservationOutput() FutureReservationOutput
	ToFutureReservationOutputWithContext(ctx context.Context) FutureReservationOutput
}

func (*FutureReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**FutureReservation)(nil)).Elem()
}

func (i *FutureReservation) ToFutureReservationOutput() FutureReservationOutput {
	return i.ToFutureReservationOutputWithContext(context.Background())
}

func (i *FutureReservation) ToFutureReservationOutputWithContext(ctx context.Context) FutureReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FutureReservationOutput)
}

type FutureReservationOutput struct{ *pulumi.OutputState }

func (FutureReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FutureReservation)(nil)).Elem()
}

func (o FutureReservationOutput) ToFutureReservationOutput() FutureReservationOutput {
	return o
}

func (o FutureReservationOutput) ToFutureReservationOutputWithContext(ctx context.Context) FutureReservationOutput {
	return o
}

// The creation timestamp for this future reservation in RFC3339 text format.
func (o FutureReservationOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the future reservation.
func (o FutureReservationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Type of the resource. Always compute#futureReservation for future reservations.
func (o FutureReservationOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o FutureReservationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name prefix for the reservations to be created at the time of delivery. The name prefix must comply with RFC1035. Maximum allowed length for name prefix is 20. Automatically created reservations name format will be -date-####.
func (o FutureReservationOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.NamePrefix }).(pulumi.StringOutput)
}

// Server-defined fully-qualified URL for this resource.
func (o FutureReservationOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o FutureReservationOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// List of Projects/Folders to share with.
func (o FutureReservationOutput) ShareSettings() ShareSettingsResponseOutput {
	return o.ApplyT(func(v *FutureReservation) ShareSettingsResponseOutput { return v.ShareSettings }).(ShareSettingsResponseOutput)
}

// Future Reservation configuration to indicate instance properties and total count.
func (o FutureReservationOutput) SpecificSkuProperties() FutureReservationSpecificSKUPropertiesResponseOutput {
	return o.ApplyT(func(v *FutureReservation) FutureReservationSpecificSKUPropertiesResponseOutput {
		return v.SpecificSkuProperties
	}).(FutureReservationSpecificSKUPropertiesResponseOutput)
}

// [Output only] Status of the Future Reservation
func (o FutureReservationOutput) Status() FutureReservationStatusResponseOutput {
	return o.ApplyT(func(v *FutureReservation) FutureReservationStatusResponseOutput { return v.Status }).(FutureReservationStatusResponseOutput)
}

// Time window for this Future Reservation.
func (o FutureReservationOutput) TimeWindow() FutureReservationTimeWindowResponseOutput {
	return o.ApplyT(func(v *FutureReservation) FutureReservationTimeWindowResponseOutput { return v.TimeWindow }).(FutureReservationTimeWindowResponseOutput)
}

// URL of the Zone where this future reservation resides.
func (o FutureReservationOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *FutureReservation) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FutureReservationInput)(nil)).Elem(), &FutureReservation{})
	pulumi.RegisterOutputType(FutureReservationOutput{})
}