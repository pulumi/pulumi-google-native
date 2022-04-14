// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Environment.
// Auto-naming is currently not supported for this resource.
type Environment struct {
	pulumi.CustomResourceState

	// Use a container image to start the notebook instance.
	ContainerImage ContainerImageResponseOutput `pulumi:"containerImage"`
	// The time at which this environment was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A brief description of this environment.
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Name of this environment. Format: `projects/{project_id}/locations/{location}/environments/{environment_id}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	PostStartupScript pulumi.StringOutput `pulumi:"postStartupScript"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage VmImageResponseOutput `pulumi:"vmImage"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	var resource Environment
	err := ctx.RegisterResource("google-native:notebooks/v1:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("google-native:notebooks/v1:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
}

type EnvironmentState struct {
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// Use a container image to start the notebook instance.
	ContainerImage *ContainerImage `pulumi:"containerImage"`
	// A brief description of this environment.
	Description *string `pulumi:"description"`
	// Display name of this environment for the UI.
	DisplayName *string `pulumi:"displayName"`
	// Required. User-defined unique ID of this environment. The `environment_id` must be 1 to 63 characters long and contain only lowercase letters, numeric characters, and dashes. The first character must be a lowercase letter and the last character cannot be a dash.
	EnvironmentId string  `pulumi:"environmentId"`
	Location      *string `pulumi:"location"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	PostStartupScript *string `pulumi:"postStartupScript"`
	Project           *string `pulumi:"project"`
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage *VmImage `pulumi:"vmImage"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// Use a container image to start the notebook instance.
	ContainerImage ContainerImagePtrInput
	// A brief description of this environment.
	Description pulumi.StringPtrInput
	// Display name of this environment for the UI.
	DisplayName pulumi.StringPtrInput
	// Required. User-defined unique ID of this environment. The `environment_id` must be 1 to 63 characters long and contain only lowercase letters, numeric characters, and dashes. The first character must be a lowercase letter and the last character cannot be a dash.
	EnvironmentId pulumi.StringInput
	Location      pulumi.StringPtrInput
	// Path to a Bash script that automatically runs after a notebook instance fully boots up. The path must be a URL or Cloud Storage path. Example: `"gs://path-to-file/file-name"`
	PostStartupScript pulumi.StringPtrInput
	Project           pulumi.StringPtrInput
	// Use a Compute Engine VM image to start the notebook instance.
	VmImage VmImagePtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
}