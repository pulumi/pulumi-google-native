// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new process.
type Process struct {
	pulumi.CustomResourceState

	// Optional. The attributes of the process. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the process). Up to 100 attributes are allowed.
	Attributes pulumi.MapOutput `pulumi:"attributes"`
	// Optional. A human-readable name you can set to display in a user interface. Must be not longer than 200 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	Location    pulumi.StringOutput `pulumi:"location"`
	// Immutable. The resource name of the lineage process. Format: `projects/{project}/locations/{location}/processes/{process}`. Can be specified or auto-assigned. {process} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The origin of this process and its runs and lineage events.
	Origin  GoogleCloudDatacatalogLineageV1OriginResponseOutput `pulumi:"origin"`
	Project pulumi.StringOutput                                 `pulumi:"project"`
	// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
}

// NewProcess registers a new resource with the given unique name, arguments, and options.
func NewProcess(ctx *pulumi.Context,
	name string, args *ProcessArgs, opts ...pulumi.ResourceOption) (*Process, error) {
	if args == nil {
		args = &ProcessArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Process
	err := ctx.RegisterResource("google-native:datalineage/v1:Process", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProcess gets an existing Process resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProcess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProcessState, opts ...pulumi.ResourceOption) (*Process, error) {
	var resource Process
	err := ctx.ReadResource("google-native:datalineage/v1:Process", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Process resources.
type processState struct {
}

type ProcessState struct {
}

func (ProcessState) ElementType() reflect.Type {
	return reflect.TypeOf((*processState)(nil)).Elem()
}

type processArgs struct {
	// Optional. The attributes of the process. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the process). Up to 100 attributes are allowed.
	Attributes map[string]interface{} `pulumi:"attributes"`
	// Optional. A human-readable name you can set to display in a user interface. Must be not longer than 200 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
	DisplayName *string `pulumi:"displayName"`
	Location    *string `pulumi:"location"`
	// Immutable. The resource name of the lineage process. Format: `projects/{project}/locations/{location}/processes/{process}`. Can be specified or auto-assigned. {process} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
	Name *string `pulumi:"name"`
	// Optional. The origin of this process and its runs and lineage events.
	Origin  *GoogleCloudDatacatalogLineageV1Origin `pulumi:"origin"`
	Project *string                                `pulumi:"project"`
	// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Process resource.
type ProcessArgs struct {
	// Optional. The attributes of the process. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the process). Up to 100 attributes are allowed.
	Attributes pulumi.MapInput
	// Optional. A human-readable name you can set to display in a user interface. Must be not longer than 200 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
	DisplayName pulumi.StringPtrInput
	Location    pulumi.StringPtrInput
	// Immutable. The resource name of the lineage process. Format: `projects/{project}/locations/{location}/processes/{process}`. Can be specified or auto-assigned. {process} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
	Name pulumi.StringPtrInput
	// Optional. The origin of this process and its runs and lineage events.
	Origin  GoogleCloudDatacatalogLineageV1OriginPtrInput
	Project pulumi.StringPtrInput
	// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
	RequestId pulumi.StringPtrInput
}

func (ProcessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*processArgs)(nil)).Elem()
}

type ProcessInput interface {
	pulumi.Input

	ToProcessOutput() ProcessOutput
	ToProcessOutputWithContext(ctx context.Context) ProcessOutput
}

func (*Process) ElementType() reflect.Type {
	return reflect.TypeOf((**Process)(nil)).Elem()
}

func (i *Process) ToProcessOutput() ProcessOutput {
	return i.ToProcessOutputWithContext(context.Background())
}

func (i *Process) ToProcessOutputWithContext(ctx context.Context) ProcessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProcessOutput)
}

type ProcessOutput struct{ *pulumi.OutputState }

func (ProcessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Process)(nil)).Elem()
}

func (o ProcessOutput) ToProcessOutput() ProcessOutput {
	return o
}

func (o ProcessOutput) ToProcessOutputWithContext(ctx context.Context) ProcessOutput {
	return o
}

// Optional. The attributes of the process. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the process). Up to 100 attributes are allowed.
func (o ProcessOutput) Attributes() pulumi.MapOutput {
	return o.ApplyT(func(v *Process) pulumi.MapOutput { return v.Attributes }).(pulumi.MapOutput)
}

// Optional. A human-readable name you can set to display in a user interface. Must be not longer than 200 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
func (o ProcessOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Process) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ProcessOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Process) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The resource name of the lineage process. Format: `projects/{project}/locations/{location}/processes/{process}`. Can be specified or auto-assigned. {process} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
func (o ProcessOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Process) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. The origin of this process and its runs and lineage events.
func (o ProcessOutput) Origin() GoogleCloudDatacatalogLineageV1OriginResponseOutput {
	return o.ApplyT(func(v *Process) GoogleCloudDatacatalogLineageV1OriginResponseOutput { return v.Origin }).(GoogleCloudDatacatalogLineageV1OriginResponseOutput)
}

func (o ProcessOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Process) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A unique identifier for this request. Restricted to 36 ASCII characters. A random UUID is recommended. This request is idempotent only if a `request_id` is provided.
func (o ProcessOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Process) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProcessInput)(nil)).Elem(), &Process{})
	pulumi.RegisterOutputType(ProcessOutput{})
}
