// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new workstation.
type Workstation struct {
	pulumi.CustomResourceState

	// Client-specified annotations.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Time when this resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Time when this resource was soft-deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// Human-readable name for this resource.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Host to which clients can send HTTPS traffic that will be received by the workstation. Authorized traffic will be received to the workstation as HTTP on port 80. To send traffic to a different port, clients may prefix the host with the destination port in the format `{port}-{host}`.
	Host pulumi.StringOutput `pulumi:"host"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Full name of this resource.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates whether this resource is currently being updated to match its intended state.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Current state of the workstation.
	State pulumi.StringOutput `pulumi:"state"`
	// A system-assigned unique identified for this resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Time when this resource was most recently updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// If set, validate the request and preview the review, but do not actually apply it.
	ValidateOnly         pulumi.BoolPtrOutput `pulumi:"validateOnly"`
	WorkstationClusterId pulumi.StringOutput  `pulumi:"workstationClusterId"`
	WorkstationConfigId  pulumi.StringOutput  `pulumi:"workstationConfigId"`
	// Required. ID to use for the workstation.
	WorkstationId pulumi.StringOutput `pulumi:"workstationId"`
}

// NewWorkstation registers a new resource with the given unique name, arguments, and options.
func NewWorkstation(ctx *pulumi.Context,
	name string, args *WorkstationArgs, opts ...pulumi.ResourceOption) (*Workstation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.WorkstationClusterId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationClusterId'")
	}
	if args.WorkstationConfigId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationConfigId'")
	}
	if args.WorkstationId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"workstationClusterId",
		"workstationConfigId",
		"workstationId",
	})
	opts = append(opts, replaceOnChanges)
	var resource Workstation
	err := ctx.RegisterResource("google-native:workstations/v1beta:Workstation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstation gets an existing Workstation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationState, opts ...pulumi.ResourceOption) (*Workstation, error) {
	var resource Workstation
	err := ctx.ReadResource("google-native:workstations/v1beta:Workstation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workstation resources.
type workstationState struct {
}

type WorkstationState struct {
}

func (WorkstationState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationState)(nil)).Elem()
}

type workstationArgs struct {
	// Client-specified annotations.
	Annotations map[string]string `pulumi:"annotations"`
	// Human-readable name for this resource.
	DisplayName *string `pulumi:"displayName"`
	// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Full name of this resource.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// If set, validate the request and preview the review, but do not actually apply it.
	ValidateOnly         *bool  `pulumi:"validateOnly"`
	WorkstationClusterId string `pulumi:"workstationClusterId"`
	WorkstationConfigId  string `pulumi:"workstationConfigId"`
	// Required. ID to use for the workstation.
	WorkstationId string `pulumi:"workstationId"`
}

// The set of arguments for constructing a Workstation resource.
type WorkstationArgs struct {
	// Client-specified annotations.
	Annotations pulumi.StringMapInput
	// Human-readable name for this resource.
	DisplayName pulumi.StringPtrInput
	// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Full name of this resource.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// If set, validate the request and preview the review, but do not actually apply it.
	ValidateOnly         pulumi.BoolPtrInput
	WorkstationClusterId pulumi.StringInput
	WorkstationConfigId  pulumi.StringInput
	// Required. ID to use for the workstation.
	WorkstationId pulumi.StringInput
}

func (WorkstationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationArgs)(nil)).Elem()
}

type WorkstationInput interface {
	pulumi.Input

	ToWorkstationOutput() WorkstationOutput
	ToWorkstationOutputWithContext(ctx context.Context) WorkstationOutput
}

func (*Workstation) ElementType() reflect.Type {
	return reflect.TypeOf((**Workstation)(nil)).Elem()
}

func (i *Workstation) ToWorkstationOutput() WorkstationOutput {
	return i.ToWorkstationOutputWithContext(context.Background())
}

func (i *Workstation) ToWorkstationOutputWithContext(ctx context.Context) WorkstationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationOutput)
}

type WorkstationOutput struct{ *pulumi.OutputState }

func (WorkstationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workstation)(nil)).Elem()
}

func (o WorkstationOutput) ToWorkstationOutput() WorkstationOutput {
	return o
}

func (o WorkstationOutput) ToWorkstationOutputWithContext(ctx context.Context) WorkstationOutput {
	return o
}

// Client-specified annotations.
func (o WorkstationOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Time when this resource was created.
func (o WorkstationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Time when this resource was soft-deleted.
func (o WorkstationOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// Human-readable name for this resource.
func (o WorkstationOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
func (o WorkstationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Host to which clients can send HTTPS traffic that will be received by the workstation. Authorized traffic will be received to the workstation as HTTP on port 80. To send traffic to a different port, clients may prefix the host with the destination port in the format `{port}-{host}`.
func (o WorkstationOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
func (o WorkstationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o WorkstationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Full name of this resource.
func (o WorkstationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkstationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Indicates whether this resource is currently being updated to match its intended state.
func (o WorkstationOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *Workstation) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Current state of the workstation.
func (o WorkstationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A system-assigned unique identified for this resource.
func (o WorkstationOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Time when this resource was most recently updated.
func (o WorkstationOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// If set, validate the request and preview the review, but do not actually apply it.
func (o WorkstationOutput) ValidateOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workstation) pulumi.BoolPtrOutput { return v.ValidateOnly }).(pulumi.BoolPtrOutput)
}

func (o WorkstationOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

// Required. ID to use for the workstation.
func (o WorkstationOutput) WorkstationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workstation) pulumi.StringOutput { return v.WorkstationId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationInput)(nil)).Elem(), &Workstation{})
	pulumi.RegisterOutputType(WorkstationOutput{})
}