// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a target VPN gateway in the specified project and region using the data included in the request.
type TargetVpnGateway struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of URLs to the ForwardingRule resources. ForwardingRules are created using compute.forwardingRules.insert and associated with a VPN gateway.
	ForwardingRules pulumi.StringArrayOutput `pulumi:"forwardingRules"`
	// Type of resource. Always compute#targetVpnGateway for target VPN gateways.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this TargetVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a TargetVpnGateway.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network pulumi.StringOutput `pulumi:"network"`
	// URL of the region where the target VPN gateway resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the VPN gateway, which can be one of the following: CREATING, READY, FAILED, or DELETING.
	Status pulumi.StringOutput `pulumi:"status"`
	// A list of URLs to VpnTunnel resources. VpnTunnels are created using the compute.vpntunnels.insert method and associated with a VPN gateway.
	Tunnels pulumi.StringArrayOutput `pulumi:"tunnels"`
}

// NewTargetVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewTargetVpnGateway(ctx *pulumi.Context,
	name string, args *TargetVpnGatewayArgs, opts ...pulumi.ResourceOption) (*TargetVpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource TargetVpnGateway
	err := ctx.RegisterResource("google-native:compute/alpha:TargetVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetVpnGateway gets an existing TargetVpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetVpnGatewayState, opts ...pulumi.ResourceOption) (*TargetVpnGateway, error) {
	var resource TargetVpnGateway
	err := ctx.ReadResource("google-native:compute/alpha:TargetVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetVpnGateway resources.
type targetVpnGatewayState struct {
}

type TargetVpnGatewayState struct {
}

func (TargetVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetVpnGatewayState)(nil)).Elem()
}

type targetVpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network *string `pulumi:"network"`
	Project *string `pulumi:"project"`
	Region  string  `pulumi:"region"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a TargetVpnGateway resource.
type TargetVpnGatewayArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
	Network pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	Region  pulumi.StringInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (TargetVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetVpnGatewayArgs)(nil)).Elem()
}

type TargetVpnGatewayInput interface {
	pulumi.Input

	ToTargetVpnGatewayOutput() TargetVpnGatewayOutput
	ToTargetVpnGatewayOutputWithContext(ctx context.Context) TargetVpnGatewayOutput
}

func (*TargetVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetVpnGateway)(nil)).Elem()
}

func (i *TargetVpnGateway) ToTargetVpnGatewayOutput() TargetVpnGatewayOutput {
	return i.ToTargetVpnGatewayOutputWithContext(context.Background())
}

func (i *TargetVpnGateway) ToTargetVpnGatewayOutputWithContext(ctx context.Context) TargetVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetVpnGatewayOutput)
}

type TargetVpnGatewayOutput struct{ *pulumi.OutputState }

func (TargetVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetVpnGateway)(nil)).Elem()
}

func (o TargetVpnGatewayOutput) ToTargetVpnGatewayOutput() TargetVpnGatewayOutput {
	return o
}

func (o TargetVpnGatewayOutput) ToTargetVpnGatewayOutputWithContext(ctx context.Context) TargetVpnGatewayOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o TargetVpnGatewayOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o TargetVpnGatewayOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A list of URLs to the ForwardingRule resources. ForwardingRules are created using compute.forwardingRules.insert and associated with a VPN gateway.
func (o TargetVpnGatewayOutput) ForwardingRules() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringArrayOutput { return v.ForwardingRules }).(pulumi.StringArrayOutput)
}

// Type of resource. Always compute#targetVpnGateway for target VPN gateways.
func (o TargetVpnGatewayOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this TargetVpnGateway, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a TargetVpnGateway.
func (o TargetVpnGatewayOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o TargetVpnGatewayOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o TargetVpnGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// URL of the network to which this VPN gateway is attached. Provided by the client when the VPN gateway is created.
func (o TargetVpnGatewayOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// URL of the region where the target VPN gateway resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
func (o TargetVpnGatewayOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o TargetVpnGatewayOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The status of the VPN gateway, which can be one of the following: CREATING, READY, FAILED, or DELETING.
func (o TargetVpnGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// A list of URLs to VpnTunnel resources. VpnTunnels are created using the compute.vpntunnels.insert method and associated with a VPN gateway.
func (o TargetVpnGatewayOutput) Tunnels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TargetVpnGateway) pulumi.StringArrayOutput { return v.Tunnels }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetVpnGatewayInput)(nil)).Elem(), &TargetVpnGateway{})
	pulumi.RegisterOutputType(TargetVpnGatewayOutput{})
}