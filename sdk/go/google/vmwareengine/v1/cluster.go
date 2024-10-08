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

// Creates a new cluster in a given private cloud. Creating a new cluster provides additional nodes for use in the parent private cloud and requires sufficient [node quota](https://cloud.google.com/vmware-engine/quotas).
// Auto-naming is currently not supported for this resource.
type Cluster struct {
	pulumi.CustomResourceState

	// Required. The user-provided identifier of the new `Cluster`. This identifier must be unique among clusters within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Creation time of this resource.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	Location   pulumi.StringOutput `pulumi:"location"`
	// True if the cluster is a management cluster; false otherwise. There can only be one management cluster in a private cloud and it has to be the first one.
	Management pulumi.BoolOutput `pulumi:"management"`
	// The resource name of this cluster. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/clusters/my-cluster`
	Name pulumi.StringOutput `pulumi:"name"`
	// The map of cluster node types in this cluster, where the key is canonical identifier of the node type (corresponds to the `NodeType`).
	NodeTypeConfigs NodeTypeConfigResponseMapOutput `pulumi:"nodeTypeConfigs"`
	PrivateCloudId  pulumi.StringOutput             `pulumi:"privateCloudId"`
	Project         pulumi.StringOutput             `pulumi:"project"`
	// Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// State of the resource.
	State pulumi.StringOutput `pulumi:"state"`
	// Optional. Configuration of a stretched cluster. Required for clusters that belong to a STRETCHED private cloud.
	StretchedClusterConfig StretchedClusterConfigResponseOutput `pulumi:"stretchedClusterConfig"`
	// System-generated unique identifier for the resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Last update time of this resource.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.NodeTypeConfigs == nil {
		return nil, errors.New("invalid value for required argument 'NodeTypeConfigs'")
	}
	if args.PrivateCloudId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"clusterId",
		"location",
		"privateCloudId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Cluster
	err := ctx.RegisterResource("google-native:vmwareengine/v1:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("google-native:vmwareengine/v1:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
}

type ClusterState struct {
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Required. The user-provided identifier of the new `Cluster`. This identifier must be unique among clusters within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	ClusterId string  `pulumi:"clusterId"`
	Location  *string `pulumi:"location"`
	// The map of cluster node types in this cluster, where the key is canonical identifier of the node type (corresponds to the `NodeType`).
	NodeTypeConfigs map[string]NodeTypeConfig `pulumi:"nodeTypeConfigs"`
	PrivateCloudId  string                    `pulumi:"privateCloudId"`
	Project         *string                   `pulumi:"project"`
	// Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Optional. Configuration of a stretched cluster. Required for clusters that belong to a STRETCHED private cloud.
	StretchedClusterConfig *StretchedClusterConfig `pulumi:"stretchedClusterConfig"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Required. The user-provided identifier of the new `Cluster`. This identifier must be unique among clusters within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
	ClusterId pulumi.StringInput
	Location  pulumi.StringPtrInput
	// The map of cluster node types in this cluster, where the key is canonical identifier of the node type (corresponds to the `NodeType`).
	NodeTypeConfigs NodeTypeConfigMapInput
	PrivateCloudId  pulumi.StringInput
	Project         pulumi.StringPtrInput
	// Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Optional. Configuration of a stretched cluster. Required for clusters that belong to a STRETCHED private cloud.
	StretchedClusterConfig StretchedClusterConfigPtrInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}

type ClusterInput interface {
	pulumi.Input

	ToClusterOutput() ClusterOutput
	ToClusterOutputWithContext(ctx context.Context) ClusterOutput
}

func (*Cluster) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (i *Cluster) ToClusterOutput() ClusterOutput {
	return i.ToClusterOutputWithContext(context.Background())
}

func (i *Cluster) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterOutput)
}

type ClusterOutput struct{ *pulumi.OutputState }

func (ClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cluster)(nil)).Elem()
}

func (o ClusterOutput) ToClusterOutput() ClusterOutput {
	return o
}

func (o ClusterOutput) ToClusterOutputWithContext(ctx context.Context) ClusterOutput {
	return o
}

// Required. The user-provided identifier of the new `Cluster`. This identifier must be unique among clusters within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
func (o ClusterOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Creation time of this resource.
func (o ClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o ClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// True if the cluster is a management cluster; false otherwise. There can only be one management cluster in a private cloud and it has to be the first one.
func (o ClusterOutput) Management() pulumi.BoolOutput {
	return o.ApplyT(func(v *Cluster) pulumi.BoolOutput { return v.Management }).(pulumi.BoolOutput)
}

// The resource name of this cluster. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/clusters/my-cluster`
func (o ClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The map of cluster node types in this cluster, where the key is canonical identifier of the node type (corresponds to the `NodeType`).
func (o ClusterOutput) NodeTypeConfigs() NodeTypeConfigResponseMapOutput {
	return o.ApplyT(func(v *Cluster) NodeTypeConfigResponseMapOutput { return v.NodeTypeConfigs }).(NodeTypeConfigResponseMapOutput)
}

func (o ClusterOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o ClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o ClusterOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// State of the resource.
func (o ClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Optional. Configuration of a stretched cluster. Required for clusters that belong to a STRETCHED private cloud.
func (o ClusterOutput) StretchedClusterConfig() StretchedClusterConfigResponseOutput {
	return o.ApplyT(func(v *Cluster) StretchedClusterConfigResponseOutput { return v.StretchedClusterConfig }).(StretchedClusterConfigResponseOutput)
}

// System-generated unique identifier for the resource.
func (o ClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Last update time of this resource.
func (o ClusterOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Cluster) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterInput)(nil)).Elem(), &Cluster{})
	pulumi.RegisterOutputType(ClusterOutput{})
}
