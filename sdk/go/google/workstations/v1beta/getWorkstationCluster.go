// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the requested WorkstationCluster.
func LookupWorkstationCluster(ctx *pulumi.Context, args *LookupWorkstationClusterArgs, opts ...pulumi.InvokeOption) (*LookupWorkstationClusterResult, error) {
	var rv LookupWorkstationClusterResult
	err := ctx.Invoke("google-native:workstations/v1beta:getWorkstationCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkstationClusterArgs struct {
	Location             string  `pulumi:"location"`
	Project              *string `pulumi:"project"`
	WorkstationClusterId string  `pulumi:"workstationClusterId"`
}

type LookupWorkstationClusterResult struct {
	// Client-specified annotations.
	Annotations map[string]string `pulumi:"annotations"`
	// Status conditions describing the current resource state.
	Conditions []StatusResponse `pulumi:"conditions"`
	// Time when this resource was created.
	CreateTime string `pulumi:"createTime"`
	// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality. Details can be found in the `conditions` field.
	Degraded bool `pulumi:"degraded"`
	// Time when this resource was soft-deleted.
	DeleteTime string `pulumi:"deleteTime"`
	// Human-readable name for this resource.
	DisplayName string `pulumi:"displayName"`
	// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag string `pulumi:"etag"`
	// Full name of this resource.
	Name string `pulumi:"name"`
	// Name of the Compute Engine network in which instances associated with this cluster will be created.
	Network string `pulumi:"network"`
	// Configuration for private cluster.
	PrivateClusterConfig PrivateClusterConfigResponse `pulumi:"privateClusterConfig"`
	// Indicates whether this resource is currently being updated to match its intended state.
	Reconciling bool `pulumi:"reconciling"`
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created. Must be part of the subnetwork specified for this cluster.
	Subnetwork string `pulumi:"subnetwork"`
	// A system-assigned unique identified for this resource.
	Uid string `pulumi:"uid"`
	// Time when this resource was most recently updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupWorkstationClusterOutput(ctx *pulumi.Context, args LookupWorkstationClusterOutputArgs, opts ...pulumi.InvokeOption) LookupWorkstationClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkstationClusterResult, error) {
			args := v.(LookupWorkstationClusterArgs)
			r, err := LookupWorkstationCluster(ctx, &args, opts...)
			var s LookupWorkstationClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkstationClusterResultOutput)
}

type LookupWorkstationClusterOutputArgs struct {
	Location             pulumi.StringInput    `pulumi:"location"`
	Project              pulumi.StringPtrInput `pulumi:"project"`
	WorkstationClusterId pulumi.StringInput    `pulumi:"workstationClusterId"`
}

func (LookupWorkstationClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkstationClusterArgs)(nil)).Elem()
}

type LookupWorkstationClusterResultOutput struct{ *pulumi.OutputState }

func (LookupWorkstationClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkstationClusterResult)(nil)).Elem()
}

func (o LookupWorkstationClusterResultOutput) ToLookupWorkstationClusterResultOutput() LookupWorkstationClusterResultOutput {
	return o
}

func (o LookupWorkstationClusterResultOutput) ToLookupWorkstationClusterResultOutputWithContext(ctx context.Context) LookupWorkstationClusterResultOutput {
	return o
}

// Client-specified annotations.
func (o LookupWorkstationClusterResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Status conditions describing the current resource state.
func (o LookupWorkstationClusterResultOutput) Conditions() StatusResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) []StatusResponse { return v.Conditions }).(StatusResponseArrayOutput)
}

// Time when this resource was created.
func (o LookupWorkstationClusterResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality. Details can be found in the `conditions` field.
func (o LookupWorkstationClusterResultOutput) Degraded() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) bool { return v.Degraded }).(pulumi.BoolOutput)
}

// Time when this resource was soft-deleted.
func (o LookupWorkstationClusterResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// Human-readable name for this resource.
func (o LookupWorkstationClusterResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
func (o LookupWorkstationClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Full name of this resource.
func (o LookupWorkstationClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Name of the Compute Engine network in which instances associated with this cluster will be created.
func (o LookupWorkstationClusterResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.Network }).(pulumi.StringOutput)
}

// Configuration for private cluster.
func (o LookupWorkstationClusterResultOutput) PrivateClusterConfig() PrivateClusterConfigResponseOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) PrivateClusterConfigResponse { return v.PrivateClusterConfig }).(PrivateClusterConfigResponseOutput)
}

// Indicates whether this resource is currently being updated to match its intended state.
func (o LookupWorkstationClusterResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created. Must be part of the subnetwork specified for this cluster.
func (o LookupWorkstationClusterResultOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.Subnetwork }).(pulumi.StringOutput)
}

// A system-assigned unique identified for this resource.
func (o LookupWorkstationClusterResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.Uid }).(pulumi.StringOutput)
}

// Time when this resource was most recently updated.
func (o LookupWorkstationClusterResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkstationClusterResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkstationClusterResultOutput{})
}