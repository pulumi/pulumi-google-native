// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Runtime in a given project and location.
// Auto-naming is currently not supported for this resource.
type Runtime struct {
	pulumi.CustomResourceState

	// The config settings for accessing runtime.
	AccessConfig RuntimeAccessConfigResponseOutput `pulumi:"accessConfig"`
	// Runtime creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Runtime health_state.
	HealthState pulumi.StringOutput `pulumi:"healthState"`
	// Contains Runtime daemon metrics such as Service status and JupyterLab stats.
	Metrics RuntimeMetricsResponseOutput `pulumi:"metrics"`
	// The resource name of the runtime. Format: `projects/{project}/locations/{location}/runtimes/{runtimeId}`
	Name pulumi.StringOutput `pulumi:"name"`
	// The config settings for software inside the runtime.
	SoftwareConfig RuntimeSoftwareConfigResponseOutput `pulumi:"softwareConfig"`
	// Runtime state.
	State pulumi.StringOutput `pulumi:"state"`
	// Runtime update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	VirtualMachine VirtualMachineResponseOutput `pulumi:"virtualMachine"`
}

// NewRuntime registers a new resource with the given unique name, arguments, and options.
func NewRuntime(ctx *pulumi.Context,
	name string, args *RuntimeArgs, opts ...pulumi.ResourceOption) (*Runtime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuntimeId == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeId'")
	}
	var resource Runtime
	err := ctx.RegisterResource("google-native:notebooks/v1:Runtime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntime gets an existing Runtime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeState, opts ...pulumi.ResourceOption) (*Runtime, error) {
	var resource Runtime
	err := ctx.ReadResource("google-native:notebooks/v1:Runtime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Runtime resources.
type runtimeState struct {
}

type RuntimeState struct {
}

func (RuntimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeState)(nil)).Elem()
}

type runtimeArgs struct {
	// The config settings for accessing runtime.
	AccessConfig *RuntimeAccessConfig `pulumi:"accessConfig"`
	Location     *string              `pulumi:"location"`
	Project      *string              `pulumi:"project"`
	// Required. User-defined unique ID of this Runtime.
	RuntimeId string `pulumi:"runtimeId"`
	// The config settings for software inside the runtime.
	SoftwareConfig *RuntimeSoftwareConfig `pulumi:"softwareConfig"`
	// Use a Compute Engine VM image to start the managed notebook instance.
	VirtualMachine *VirtualMachine `pulumi:"virtualMachine"`
}

// The set of arguments for constructing a Runtime resource.
type RuntimeArgs struct {
	// The config settings for accessing runtime.
	AccessConfig RuntimeAccessConfigPtrInput
	Location     pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	// Required. User-defined unique ID of this Runtime.
	RuntimeId pulumi.StringInput
	// The config settings for software inside the runtime.
	SoftwareConfig RuntimeSoftwareConfigPtrInput
	// Use a Compute Engine VM image to start the managed notebook instance.
	VirtualMachine VirtualMachinePtrInput
}

func (RuntimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeArgs)(nil)).Elem()
}

type RuntimeInput interface {
	pulumi.Input

	ToRuntimeOutput() RuntimeOutput
	ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput
}

func (*Runtime) ElementType() reflect.Type {
	return reflect.TypeOf((**Runtime)(nil)).Elem()
}

func (i *Runtime) ToRuntimeOutput() RuntimeOutput {
	return i.ToRuntimeOutputWithContext(context.Background())
}

func (i *Runtime) ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeOutput)
}

type RuntimeOutput struct{ *pulumi.OutputState }

func (RuntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Runtime)(nil)).Elem()
}

func (o RuntimeOutput) ToRuntimeOutput() RuntimeOutput {
	return o
}

func (o RuntimeOutput) ToRuntimeOutputWithContext(ctx context.Context) RuntimeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeInput)(nil)).Elem(), &Runtime{})
	pulumi.RegisterOutputType(RuntimeOutput{})
}