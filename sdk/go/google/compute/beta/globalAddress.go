// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an address resource in the specified project by using the data included in the request.
type GlobalAddress struct {
	pulumi.CustomResourceState

	// The static IP address represented by this resource.
	Address pulumi.StringOutput `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType pulumi.StringOutput `pulumi:"addressType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion pulumi.StringOutput `pulumi:"ipVersion"`
	// Type of the resource. Always compute#address for addresses.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for the labels being applied to this Address, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an Address.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network pulumi.StringOutput `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier pulumi.StringOutput `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.IntOutput `pulumi:"prefixLength"`
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *IPsec-encrypted Cloud Interconnect* configuration. These addresses are regional resources. Not currently available publicly. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose pulumi.StringOutput `pulumi:"purpose"`
	// The URL of the region where a regional address resides. For regional addresses, you must specify the region as a path parameter in the HTTP request URL. *This field is not applicable to global addresses.*
	Region pulumi.StringOutput `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The status of the address, which can be one of RESERVING, RESERVED, or IN_USE. An address that is RESERVING is currently in the process of being reserved. A RESERVED address is currently reserved and available to use. An IN_USE address is currently being used by another resource and is not available.
	Status pulumi.StringOutput `pulumi:"status"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The URLs of the resources that are using this address.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewGlobalAddress registers a new resource with the given unique name, arguments, and options.
func NewGlobalAddress(ctx *pulumi.Context,
	name string, args *GlobalAddressArgs, opts ...pulumi.ResourceOption) (*GlobalAddress, error) {
	if args == nil {
		args = &GlobalAddressArgs{}
	}

	var resource GlobalAddress
	err := ctx.RegisterResource("google-native:compute/beta:GlobalAddress", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalAddress gets an existing GlobalAddress resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalAddressState, opts ...pulumi.ResourceOption) (*GlobalAddress, error) {
	var resource GlobalAddress
	err := ctx.ReadResource("google-native:compute/beta:GlobalAddress", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalAddress resources.
type globalAddressState struct {
}

type GlobalAddressState struct {
}

func (GlobalAddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalAddressState)(nil)).Elem()
}

type globalAddressArgs struct {
	// The static IP address represented by this resource.
	Address *string `pulumi:"address"`
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType *GlobalAddressAddressType `pulumi:"addressType"`
	// An optional description of this resource. Provide this field when you create the resource.
	Description *string `pulumi:"description"`
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion *GlobalAddressIpVersion `pulumi:"ipVersion"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name *string `pulumi:"name"`
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network *string `pulumi:"network"`
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier *GlobalAddressNetworkTier `pulumi:"networkTier"`
	// The prefix length if the resource represents an IP range.
	PrefixLength *int    `pulumi:"prefixLength"`
	Project      *string `pulumi:"project"`
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *IPsec-encrypted Cloud Interconnect* configuration. These addresses are regional resources. Not currently available publicly. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose *GlobalAddressPurpose `pulumi:"purpose"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork *string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a GlobalAddress resource.
type GlobalAddressArgs struct {
	// The static IP address represented by this resource.
	Address pulumi.StringPtrInput
	// The type of address to reserve, either INTERNAL or EXTERNAL. If unspecified, defaults to EXTERNAL.
	AddressType GlobalAddressAddressTypePtrInput
	// An optional description of this resource. Provide this field when you create the resource.
	Description pulumi.StringPtrInput
	// The IP version that will be used by this address. Valid options are IPV4 or IPV6. This can only be specified for a global address.
	IpVersion GlobalAddressIpVersionPtrInput
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels pulumi.StringMapInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`. The first character must be a lowercase letter, and all following characters (except for the last character) must be a dash, lowercase letter, or digit. The last character must be a lowercase letter or digit.
	Name pulumi.StringPtrInput
	// The URL of the network in which to reserve the address. This field can only be used with INTERNAL type with the VPC_PEERING purpose.
	Network pulumi.StringPtrInput
	// This signifies the networking tier used for configuring this address and can only take the following values: PREMIUM or STANDARD. Internal IP addresses are always Premium Tier; global external IP addresses are always Premium Tier; regional external IP addresses can be either Standard or Premium Tier. If this field is not specified, it is assumed to be PREMIUM.
	NetworkTier GlobalAddressNetworkTierPtrInput
	// The prefix length if the resource represents an IP range.
	PrefixLength pulumi.IntPtrInput
	Project      pulumi.StringPtrInput
	// The purpose of this resource, which can be one of the following values: - GCE_ENDPOINT for addresses that are used by VM instances, alias IP ranges, load balancers, and similar resources. - DNS_RESOLVER for a DNS resolver address in a subnetwork for a Cloud DNS inbound forwarder IP addresses (regional internal IP address in a subnet of a VPC network) - VPC_PEERING for global internal IP addresses used for private services access allocated ranges. - NAT_AUTO for the regional external IP addresses used by Cloud NAT when allocating addresses using automatic NAT IP address allocation. - IPSEC_INTERCONNECT for addresses created from a private IP range that are reserved for a VLAN attachment in an *IPsec-encrypted Cloud Interconnect* configuration. These addresses are regional resources. Not currently available publicly. - `SHARED_LOADBALANCER_VIP` for an internal IP address that is assigned to multiple internal forwarding rules. - `PRIVATE_SERVICE_CONNECT` for a private network address that is used to configure Private Service Connect. Only global internal addresses can use this purpose.
	Purpose GlobalAddressPurposePtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// The URL of the subnetwork in which to reserve the address. If an IP address is specified, it must be within the subnetwork's IP range. This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose.
	Subnetwork pulumi.StringPtrInput
}

func (GlobalAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalAddressArgs)(nil)).Elem()
}

type GlobalAddressInput interface {
	pulumi.Input

	ToGlobalAddressOutput() GlobalAddressOutput
	ToGlobalAddressOutputWithContext(ctx context.Context) GlobalAddressOutput
}

func (*GlobalAddress) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalAddress)(nil)).Elem()
}

func (i *GlobalAddress) ToGlobalAddressOutput() GlobalAddressOutput {
	return i.ToGlobalAddressOutputWithContext(context.Background())
}

func (i *GlobalAddress) ToGlobalAddressOutputWithContext(ctx context.Context) GlobalAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalAddressOutput)
}

type GlobalAddressOutput struct{ *pulumi.OutputState }

func (GlobalAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalAddress)(nil)).Elem()
}

func (o GlobalAddressOutput) ToGlobalAddressOutput() GlobalAddressOutput {
	return o
}

func (o GlobalAddressOutput) ToGlobalAddressOutputWithContext(ctx context.Context) GlobalAddressOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalAddressInput)(nil)).Elem(), &GlobalAddress{})
	pulumi.RegisterOutputType(GlobalAddressOutput{})
}