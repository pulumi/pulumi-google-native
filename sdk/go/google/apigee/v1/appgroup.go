// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an AppGroup. Once created, user can register apps under the AppGroup to obtain secret key and password. At creation time, the AppGroup's state is set as `active`.
type Appgroup struct {
	pulumi.CustomResourceState

	// Internal identifier that cannot be edited
	AppGroupId pulumi.StringOutput `pulumi:"appGroupId"`
	// A list of attributes
	Attributes GoogleCloudApigeeV1AttributeResponseArrayOutput `pulumi:"attributes"`
	// channel identifier identifies the owner maintaing this grouping.
	ChannelId pulumi.StringOutput `pulumi:"channelId"`
	// A reference to the associated storefront/marketplace.
	ChannelUri pulumi.StringOutput `pulumi:"channelUri"`
	// Created time as milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// app group name displayed in the UI
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Modified time as milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Immutable. Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._\-$ %.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. the org the app group is created
	Organization   pulumi.StringOutput `pulumi:"organization"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Valid values are `active` or `inactive`. Note that the status of the AppGroup should be updated via UpdateAppGroupRequest by setting the action as `active` or `inactive`.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAppgroup registers a new resource with the given unique name, arguments, and options.
func NewAppgroup(ctx *pulumi.Context,
	name string, args *AppgroupArgs, opts ...pulumi.ResourceOption) (*Appgroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Appgroup
	err := ctx.RegisterResource("google-native:apigee/v1:Appgroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppgroup gets an existing Appgroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppgroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppgroupState, opts ...pulumi.ResourceOption) (*Appgroup, error) {
	var resource Appgroup
	err := ctx.ReadResource("google-native:apigee/v1:Appgroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Appgroup resources.
type appgroupState struct {
}

type AppgroupState struct {
}

func (AppgroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*appgroupState)(nil)).Elem()
}

type appgroupArgs struct {
	// A list of attributes
	Attributes []GoogleCloudApigeeV1Attribute `pulumi:"attributes"`
	// channel identifier identifies the owner maintaing this grouping.
	ChannelId *string `pulumi:"channelId"`
	// A reference to the associated storefront/marketplace.
	ChannelUri *string `pulumi:"channelUri"`
	// app group name displayed in the UI
	DisplayName *string `pulumi:"displayName"`
	// Immutable. Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._\-$ %.
	Name *string `pulumi:"name"`
	// Immutable. the org the app group is created
	Organization   *string `pulumi:"organization"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a Appgroup resource.
type AppgroupArgs struct {
	// A list of attributes
	Attributes GoogleCloudApigeeV1AttributeArrayInput
	// channel identifier identifies the owner maintaing this grouping.
	ChannelId pulumi.StringPtrInput
	// A reference to the associated storefront/marketplace.
	ChannelUri pulumi.StringPtrInput
	// app group name displayed in the UI
	DisplayName pulumi.StringPtrInput
	// Immutable. Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._\-$ %.
	Name pulumi.StringPtrInput
	// Immutable. the org the app group is created
	Organization   pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (AppgroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appgroupArgs)(nil)).Elem()
}

type AppgroupInput interface {
	pulumi.Input

	ToAppgroupOutput() AppgroupOutput
	ToAppgroupOutputWithContext(ctx context.Context) AppgroupOutput
}

func (*Appgroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Appgroup)(nil)).Elem()
}

func (i *Appgroup) ToAppgroupOutput() AppgroupOutput {
	return i.ToAppgroupOutputWithContext(context.Background())
}

func (i *Appgroup) ToAppgroupOutputWithContext(ctx context.Context) AppgroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppgroupOutput)
}

type AppgroupOutput struct{ *pulumi.OutputState }

func (AppgroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Appgroup)(nil)).Elem()
}

func (o AppgroupOutput) ToAppgroupOutput() AppgroupOutput {
	return o
}

func (o AppgroupOutput) ToAppgroupOutputWithContext(ctx context.Context) AppgroupOutput {
	return o
}

// Internal identifier that cannot be edited
func (o AppgroupOutput) AppGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.AppGroupId }).(pulumi.StringOutput)
}

// A list of attributes
func (o AppgroupOutput) Attributes() GoogleCloudApigeeV1AttributeResponseArrayOutput {
	return o.ApplyT(func(v *Appgroup) GoogleCloudApigeeV1AttributeResponseArrayOutput { return v.Attributes }).(GoogleCloudApigeeV1AttributeResponseArrayOutput)
}

// channel identifier identifies the owner maintaing this grouping.
func (o AppgroupOutput) ChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.ChannelId }).(pulumi.StringOutput)
}

// A reference to the associated storefront/marketplace.
func (o AppgroupOutput) ChannelUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.ChannelUri }).(pulumi.StringOutput)
}

// Created time as milliseconds since epoch.
func (o AppgroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// app group name displayed in the UI
func (o AppgroupOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Modified time as milliseconds since epoch.
func (o AppgroupOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.LastModifiedAt }).(pulumi.StringOutput)
}

// Immutable. Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._\-$ %.
func (o AppgroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Immutable. the org the app group is created
func (o AppgroupOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

func (o AppgroupOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// Valid values are `active` or `inactive`. Note that the status of the AppGroup should be updated via UpdateAppGroupRequest by setting the action as `active` or `inactive`.
func (o AppgroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Appgroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppgroupInput)(nil)).Elem(), &Appgroup{})
	pulumi.RegisterOutputType(AppgroupOutput{})
}
