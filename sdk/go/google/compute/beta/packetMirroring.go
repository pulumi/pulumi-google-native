// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a PacketMirroring resource in the specified project and region using the data included in the request.
type PacketMirroring struct {
	pulumi.CustomResourceState

	// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
	CollectorIlb PacketMirroringForwardingRuleInfoResponseOutput `pulumi:"collectorIlb"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
	Enable pulumi.StringOutput `pulumi:"enable"`
	// Filter for mirrored traffic. If unspecified, all traffic is mirrored.
	Filter PacketMirroringFilterResponseOutput `pulumi:"filter"`
	// Type of the resource. Always compute#packetMirroring for packet mirrorings.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	MirroredResources PacketMirroringMirroredResourceInfoResponseOutput `pulumi:"mirroredResources"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	Network PacketMirroringNetworkInfoResponseOutput `pulumi:"network"`
	// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// URI of the region where the packetMirroring resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewPacketMirroring registers a new resource with the given unique name, arguments, and options.
func NewPacketMirroring(ctx *pulumi.Context,
	name string, args *PacketMirroringArgs, opts ...pulumi.ResourceOption) (*PacketMirroring, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	var resource PacketMirroring
	err := ctx.RegisterResource("google-native:compute/beta:PacketMirroring", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPacketMirroring gets an existing PacketMirroring resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPacketMirroring(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PacketMirroringState, opts ...pulumi.ResourceOption) (*PacketMirroring, error) {
	var resource PacketMirroring
	err := ctx.ReadResource("google-native:compute/beta:PacketMirroring", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PacketMirroring resources.
type packetMirroringState struct {
}

type PacketMirroringState struct {
}

func (PacketMirroringState) ElementType() reflect.Type {
	return reflect.TypeOf((*packetMirroringState)(nil)).Elem()
}

type packetMirroringArgs struct {
	// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
	CollectorIlb *PacketMirroringForwardingRuleInfo `pulumi:"collectorIlb"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
	Enable *PacketMirroringEnable `pulumi:"enable"`
	// Filter for mirrored traffic. If unspecified, all traffic is mirrored.
	Filter *PacketMirroringFilter `pulumi:"filter"`
	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	MirroredResources *PacketMirroringMirroredResourceInfo `pulumi:"mirroredResources"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	Network *PacketMirroringNetworkInfo `pulumi:"network"`
	// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
	Priority  *int    `pulumi:"priority"`
	Project   *string `pulumi:"project"`
	Region    string  `pulumi:"region"`
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a PacketMirroring resource.
type PacketMirroringArgs struct {
	// The Forwarding Rule resource of type loadBalancingScheme=INTERNAL that will be used as collector for mirrored traffic. The specified forwarding rule must have isMirroringCollector set to true.
	CollectorIlb PacketMirroringForwardingRuleInfoPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Indicates whether or not this packet mirroring takes effect. If set to FALSE, this packet mirroring policy will not be enforced on the network. The default is TRUE.
	Enable PacketMirroringEnablePtrInput
	// Filter for mirrored traffic. If unspecified, all traffic is mirrored.
	Filter PacketMirroringFilterPtrInput
	// PacketMirroring mirroredResourceInfos. MirroredResourceInfo specifies a set of mirrored VM instances, subnetworks and/or tags for which traffic from/to all VM instances will be mirrored.
	MirroredResources PacketMirroringMirroredResourceInfoPtrInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// Specifies the mirrored VPC network. Only packets in this network will be mirrored. All mirrored VMs should have a NIC in the given network. All mirrored subnetworks should belong to the given network.
	Network PacketMirroringNetworkInfoPtrInput
	// The priority of applying this configuration. Priority is used to break ties in cases where there is more than one matching rule. In the case of two rules that apply for a given Instance, the one with the lowest-numbered priority value wins. Default value is 1000. Valid range is 0 through 65535.
	Priority  pulumi.IntPtrInput
	Project   pulumi.StringPtrInput
	Region    pulumi.StringInput
	RequestId pulumi.StringPtrInput
}

func (PacketMirroringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packetMirroringArgs)(nil)).Elem()
}

type PacketMirroringInput interface {
	pulumi.Input

	ToPacketMirroringOutput() PacketMirroringOutput
	ToPacketMirroringOutputWithContext(ctx context.Context) PacketMirroringOutput
}

func (*PacketMirroring) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketMirroring)(nil)).Elem()
}

func (i *PacketMirroring) ToPacketMirroringOutput() PacketMirroringOutput {
	return i.ToPacketMirroringOutputWithContext(context.Background())
}

func (i *PacketMirroring) ToPacketMirroringOutputWithContext(ctx context.Context) PacketMirroringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PacketMirroringOutput)
}

type PacketMirroringOutput struct{ *pulumi.OutputState }

func (PacketMirroringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PacketMirroring)(nil)).Elem()
}

func (o PacketMirroringOutput) ToPacketMirroringOutput() PacketMirroringOutput {
	return o
}

func (o PacketMirroringOutput) ToPacketMirroringOutputWithContext(ctx context.Context) PacketMirroringOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PacketMirroringInput)(nil)).Elem(), &PacketMirroring{})
	pulumi.RegisterOutputType(PacketMirroringOutput{})
}