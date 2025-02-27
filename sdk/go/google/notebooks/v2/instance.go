// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Instance in a given project and location.
// Auto-naming is currently not supported for this resource.
type Instance struct {
	pulumi.CustomResourceState

	// Instance creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Email address of entity that sent original CreateInstance request.
	Creator pulumi.StringOutput `pulumi:"creator"`
	// Optional. If true, the notebook instance will not register with the proxy.
	DisableProxyAccess pulumi.BoolOutput `pulumi:"disableProxyAccess"`
	// Optional. Compute Engine setup for the notebook. Uses notebook-defined fields.
	GceSetup GceSetupResponseOutput `pulumi:"gceSetup"`
	// Additional information about instance health. Example: healthInfo": { "docker_proxy_agent_status": "1", "docker_status": "1", "jupyterlab_api_status": "-1", "jupyterlab_status": "-1", "updated": "2020-10-18 09:40:03.573409" }
	HealthInfo pulumi.StringMapOutput `pulumi:"healthInfo"`
	// Instance health_state.
	HealthState pulumi.StringOutput `pulumi:"healthState"`
	// Required. User-defined unique ID of this instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Optional. Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
	InstanceOwners pulumi.StringArrayOutput `pulumi:"instanceOwners"`
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The name of this notebook instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The proxy endpoint that is used to access the Jupyter notebook.
	ProxyUri pulumi.StringOutput `pulumi:"proxyUri"`
	// Optional. Idempotent request UUID.
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// The state of this instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Instance update time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The upgrade history of this instance.
	UpgradeHistory UpgradeHistoryEntryResponseArrayOutput `pulumi:"upgradeHistory"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("google-native:notebooks/v2:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("google-native:notebooks/v2:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Optional. If true, the notebook instance will not register with the proxy.
	DisableProxyAccess *bool `pulumi:"disableProxyAccess"`
	// Optional. Compute Engine setup for the notebook. Uses notebook-defined fields.
	GceSetup *GceSetup `pulumi:"gceSetup"`
	// Required. User-defined unique ID of this instance.
	InstanceId string `pulumi:"instanceId"`
	// Optional. Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
	InstanceOwners []string `pulumi:"instanceOwners"`
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// Optional. Idempotent request UUID.
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Optional. If true, the notebook instance will not register with the proxy.
	DisableProxyAccess pulumi.BoolPtrInput
	// Optional. Compute Engine setup for the notebook. Uses notebook-defined fields.
	GceSetup GceSetupPtrInput
	// Required. User-defined unique ID of this instance.
	InstanceId pulumi.StringInput
	// Optional. Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
	InstanceOwners pulumi.StringArrayInput
	// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Optional. Idempotent request UUID.
	RequestId pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Instance creation time.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Email address of entity that sent original CreateInstance request.
func (o InstanceOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Creator }).(pulumi.StringOutput)
}

// Optional. If true, the notebook instance will not register with the proxy.
func (o InstanceOutput) DisableProxyAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.DisableProxyAccess }).(pulumi.BoolOutput)
}

// Optional. Compute Engine setup for the notebook. Uses notebook-defined fields.
func (o InstanceOutput) GceSetup() GceSetupResponseOutput {
	return o.ApplyT(func(v *Instance) GceSetupResponseOutput { return v.GceSetup }).(GceSetupResponseOutput)
}

// Additional information about instance health. Example: healthInfo": { "docker_proxy_agent_status": "1", "docker_status": "1", "jupyterlab_api_status": "-1", "jupyterlab_status": "-1", "updated": "2020-10-18 09:40:03.573409" }
func (o InstanceOutput) HealthInfo() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.HealthInfo }).(pulumi.StringMapOutput)
}

// Instance health_state.
func (o InstanceOutput) HealthState() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.HealthState }).(pulumi.StringOutput)
}

// Required. User-defined unique ID of this instance.
func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Optional. Input only. The owner of this instance after creation. Format: `alias@example.com` Currently supports one owner only. If not specified, all of the service account users of your VM instance's service account can use the instance.
func (o InstanceOutput) InstanceOwners() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.InstanceOwners }).(pulumi.StringArrayOutput)
}

// Optional. Labels to apply to this instance. These can be later modified by the UpdateInstance method.
func (o InstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of this notebook instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o InstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The proxy endpoint that is used to access the Jupyter notebook.
func (o InstanceOutput) ProxyUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ProxyUri }).(pulumi.StringOutput)
}

// Optional. Idempotent request UUID.
func (o InstanceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// The state of this instance.
func (o InstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Instance update time.
func (o InstanceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The upgrade history of this instance.
func (o InstanceOutput) UpgradeHistory() UpgradeHistoryEntryResponseArrayOutput {
	return o.ApplyT(func(v *Instance) UpgradeHistoryEntryResponseArrayOutput { return v.UpgradeHistory }).(UpgradeHistoryEntryResponseArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
