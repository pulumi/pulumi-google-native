// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns the specified node template. Gets a list of available node templates by making a list() request.
func LookupNodeTemplate(ctx *pulumi.Context, args *LookupNodeTemplateArgs, opts ...pulumi.InvokeOption) (*LookupNodeTemplateResult, error) {
	var rv LookupNodeTemplateResult
	err := ctx.Invoke("google-native:compute/v1:getNodeTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNodeTemplateArgs struct {
	NodeTemplate string  `pulumi:"nodeTemplate"`
	Project      *string `pulumi:"project"`
	Region       string  `pulumi:"region"`
}

type LookupNodeTemplateResult struct {
	Accelerators []AcceleratorConfigResponse `pulumi:"accelerators"`
	// CPU overcommit.
	CpuOvercommitType string `pulumi:"cpuOvercommitType"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string              `pulumi:"description"`
	Disks       []LocalDiskResponse `pulumi:"disks"`
	// The type of the resource. Always compute#nodeTemplate for node templates.
	Kind string `pulumi:"kind"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// The node type to use for nodes group that are created from this template.
	NodeType string `pulumi:"nodeType"`
	// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityResponse `pulumi:"nodeTypeFlexibility"`
	// The name of the region where the node template resides, such as us-central1.
	Region string `pulumi:"region"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
	ServerBinding ServerBindingResponse `pulumi:"serverBinding"`
	// The status of the node template. One of the following values: CREATING, READY, and DELETING.
	Status string `pulumi:"status"`
	// An optional, human-readable explanation of the status.
	StatusMessage string `pulumi:"statusMessage"`
}

func LookupNodeTemplateOutput(ctx *pulumi.Context, args LookupNodeTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupNodeTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNodeTemplateResult, error) {
			args := v.(LookupNodeTemplateArgs)
			r, err := LookupNodeTemplate(ctx, &args, opts...)
			return *r, err
		}).(LookupNodeTemplateResultOutput)
}

type LookupNodeTemplateOutputArgs struct {
	NodeTemplate pulumi.StringInput    `pulumi:"nodeTemplate"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	Region       pulumi.StringInput    `pulumi:"region"`
}

func (LookupNodeTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeTemplateArgs)(nil)).Elem()
}

type LookupNodeTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupNodeTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNodeTemplateResult)(nil)).Elem()
}

func (o LookupNodeTemplateResultOutput) ToLookupNodeTemplateResultOutput() LookupNodeTemplateResultOutput {
	return o
}

func (o LookupNodeTemplateResultOutput) ToLookupNodeTemplateResultOutputWithContext(ctx context.Context) LookupNodeTemplateResultOutput {
	return o
}

func (o LookupNodeTemplateResultOutput) Accelerators() AcceleratorConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) []AcceleratorConfigResponse { return v.Accelerators }).(AcceleratorConfigResponseArrayOutput)
}

// CPU overcommit.
func (o LookupNodeTemplateResultOutput) CpuOvercommitType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.CpuOvercommitType }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupNodeTemplateResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupNodeTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNodeTemplateResultOutput) Disks() LocalDiskResponseArrayOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) []LocalDiskResponse { return v.Disks }).(LocalDiskResponseArrayOutput)
}

// The type of the resource. Always compute#nodeTemplate for node templates.
func (o LookupNodeTemplateResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupNodeTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Labels to use for node affinity, which will be used in instance scheduling.
func (o LookupNodeTemplateResultOutput) NodeAffinityLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) map[string]string { return v.NodeAffinityLabels }).(pulumi.StringMapOutput)
}

// The node type to use for nodes group that are created from this template.
func (o LookupNodeTemplateResultOutput) NodeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.NodeType }).(pulumi.StringOutput)
}

// The flexible properties of the desired node type. Node groups that use this node template will create nodes of a type that matches these properties. This field is mutually exclusive with the node_type property; you can only define one or the other, but not both.
func (o LookupNodeTemplateResultOutput) NodeTypeFlexibility() NodeTemplateNodeTypeFlexibilityResponseOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) NodeTemplateNodeTypeFlexibilityResponse { return v.NodeTypeFlexibility }).(NodeTemplateNodeTypeFlexibilityResponseOutput)
}

// The name of the region where the node template resides, such as us-central1.
func (o LookupNodeTemplateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.Region }).(pulumi.StringOutput)
}

// Server-defined URL for the resource.
func (o LookupNodeTemplateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Sets the binding properties for the physical server. Valid values include: - *[Default]* RESTART_NODE_ON_ANY_SERVER: Restarts VMs on any available physical server - RESTART_NODE_ON_MINIMAL_SERVER: Restarts VMs on the same physical server whenever possible See Sole-tenant node options for more information.
func (o LookupNodeTemplateResultOutput) ServerBinding() ServerBindingResponseOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) ServerBindingResponse { return v.ServerBinding }).(ServerBindingResponseOutput)
}

// The status of the node template. One of the following values: CREATING, READY, and DELETING.
func (o LookupNodeTemplateResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.Status }).(pulumi.StringOutput)
}

// An optional, human-readable explanation of the status.
func (o LookupNodeTemplateResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNodeTemplateResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNodeTemplateResultOutput{})
}