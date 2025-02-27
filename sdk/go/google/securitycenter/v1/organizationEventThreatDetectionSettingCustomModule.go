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

// Creates an Event Threat Detection custom module.
// Auto-naming is currently not supported for this resource.
type OrganizationEventThreatDetectionSettingCustomModule struct {
	pulumi.CustomResourceState

	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its config value is inherited from the ancestor module.
	Config pulumi.MapOutput `pulumi:"config"`
	// The description for the module.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy.
	EnablementState pulumi.StringOutput `pulumi:"enablementState"`
	// The editor the module was last updated by.
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// Immutable. The resource name of the Event Threat Detection custom module. Its format is: * "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}". * "folders/{folder}/eventThreatDetectionSettings/customModules/{module}". * "projects/{project}/eventThreatDetectionSettings/customModules/{module}".
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the module was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrganizationEventThreatDetectionSettingCustomModule registers a new resource with the given unique name, arguments, and options.
func NewOrganizationEventThreatDetectionSettingCustomModule(ctx *pulumi.Context,
	name string, args *OrganizationEventThreatDetectionSettingCustomModuleArgs, opts ...pulumi.ResourceOption) (*OrganizationEventThreatDetectionSettingCustomModule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrganizationEventThreatDetectionSettingCustomModule
	err := ctx.RegisterResource("google-native:securitycenter/v1:OrganizationEventThreatDetectionSettingCustomModule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationEventThreatDetectionSettingCustomModule gets an existing OrganizationEventThreatDetectionSettingCustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationEventThreatDetectionSettingCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationEventThreatDetectionSettingCustomModuleState, opts ...pulumi.ResourceOption) (*OrganizationEventThreatDetectionSettingCustomModule, error) {
	var resource OrganizationEventThreatDetectionSettingCustomModule
	err := ctx.ReadResource("google-native:securitycenter/v1:OrganizationEventThreatDetectionSettingCustomModule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationEventThreatDetectionSettingCustomModule resources.
type organizationEventThreatDetectionSettingCustomModuleState struct {
}

type OrganizationEventThreatDetectionSettingCustomModuleState struct {
}

func (OrganizationEventThreatDetectionSettingCustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEventThreatDetectionSettingCustomModuleState)(nil)).Elem()
}

type organizationEventThreatDetectionSettingCustomModuleArgs struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its config value is inherited from the ancestor module.
	Config map[string]interface{} `pulumi:"config"`
	// The description for the module.
	Description *string `pulumi:"description"`
	// The human readable name to be displayed for the module.
	DisplayName *string `pulumi:"displayName"`
	// The state of enablement for the module at the given level of the hierarchy.
	EnablementState *OrganizationEventThreatDetectionSettingCustomModuleEnablementState `pulumi:"enablementState"`
	// Immutable. The resource name of the Event Threat Detection custom module. Its format is: * "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}". * "folders/{folder}/eventThreatDetectionSettings/customModules/{module}". * "projects/{project}/eventThreatDetectionSettings/customModules/{module}".
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a OrganizationEventThreatDetectionSettingCustomModule resource.
type OrganizationEventThreatDetectionSettingCustomModuleArgs struct {
	// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its config value is inherited from the ancestor module.
	Config pulumi.MapInput
	// The description for the module.
	Description pulumi.StringPtrInput
	// The human readable name to be displayed for the module.
	DisplayName pulumi.StringPtrInput
	// The state of enablement for the module at the given level of the hierarchy.
	EnablementState OrganizationEventThreatDetectionSettingCustomModuleEnablementStatePtrInput
	// Immutable. The resource name of the Event Threat Detection custom module. Its format is: * "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}". * "folders/{folder}/eventThreatDetectionSettings/customModules/{module}". * "projects/{project}/eventThreatDetectionSettings/customModules/{module}".
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Type for the module. e.g. CONFIGURABLE_BAD_IP.
	Type pulumi.StringPtrInput
}

func (OrganizationEventThreatDetectionSettingCustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationEventThreatDetectionSettingCustomModuleArgs)(nil)).Elem()
}

type OrganizationEventThreatDetectionSettingCustomModuleInput interface {
	pulumi.Input

	ToOrganizationEventThreatDetectionSettingCustomModuleOutput() OrganizationEventThreatDetectionSettingCustomModuleOutput
	ToOrganizationEventThreatDetectionSettingCustomModuleOutputWithContext(ctx context.Context) OrganizationEventThreatDetectionSettingCustomModuleOutput
}

func (*OrganizationEventThreatDetectionSettingCustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationEventThreatDetectionSettingCustomModule)(nil)).Elem()
}

func (i *OrganizationEventThreatDetectionSettingCustomModule) ToOrganizationEventThreatDetectionSettingCustomModuleOutput() OrganizationEventThreatDetectionSettingCustomModuleOutput {
	return i.ToOrganizationEventThreatDetectionSettingCustomModuleOutputWithContext(context.Background())
}

func (i *OrganizationEventThreatDetectionSettingCustomModule) ToOrganizationEventThreatDetectionSettingCustomModuleOutputWithContext(ctx context.Context) OrganizationEventThreatDetectionSettingCustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationEventThreatDetectionSettingCustomModuleOutput)
}

type OrganizationEventThreatDetectionSettingCustomModuleOutput struct{ *pulumi.OutputState }

func (OrganizationEventThreatDetectionSettingCustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationEventThreatDetectionSettingCustomModule)(nil)).Elem()
}

func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) ToOrganizationEventThreatDetectionSettingCustomModuleOutput() OrganizationEventThreatDetectionSettingCustomModuleOutput {
	return o
}

func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) ToOrganizationEventThreatDetectionSettingCustomModuleOutputWithContext(ctx context.Context) OrganizationEventThreatDetectionSettingCustomModuleOutput {
	return o
}

// Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its config value is inherited from the ancestor module.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) Config() pulumi.MapOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.MapOutput { return v.Config }).(pulumi.MapOutput)
}

// The description for the module.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The human readable name to be displayed for the module.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The state of enablement for the module at the given level of the hierarchy.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) EnablementState() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput {
		return v.EnablementState
	}).(pulumi.StringOutput)
}

// The editor the module was last updated by.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput { return v.LastEditor }).(pulumi.StringOutput)
}

// Immutable. The resource name of the Event Threat Detection custom module. Its format is: * "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}". * "folders/{folder}/eventThreatDetectionSettings/customModules/{module}". * "projects/{project}/eventThreatDetectionSettings/customModules/{module}".
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput {
		return v.OrganizationId
	}).(pulumi.StringOutput)
}

// Type for the module. e.g. CONFIGURABLE_BAD_IP.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time the module was last updated.
func (o OrganizationEventThreatDetectionSettingCustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationEventThreatDetectionSettingCustomModule) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationEventThreatDetectionSettingCustomModuleInput)(nil)).Elem(), &OrganizationEventThreatDetectionSettingCustomModule{})
	pulumi.RegisterOutputType(OrganizationEventThreatDetectionSettingCustomModuleOutput{})
}
