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

// Creates a resident SecurityHealthAnalyticsCustomModule at the scope of the given CRM parent, and also creates inherited SecurityHealthAnalyticsCustomModules for all CRM descendants of the given parent. These modules are enabled by default.
// Auto-naming is currently not supported for this resource.
type OrganizationSecurityHealthAnalyticsSettingCustomModule struct {
	pulumi.CustomResourceState

	// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
	AncestorModule pulumi.StringOutput `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig GoogleCloudSecuritycenterV1CustomConfigResponseOutput `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The enablement state of the custom module.
	EnablementState pulumi.StringOutput `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The time at which the custom module was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrganizationSecurityHealthAnalyticsSettingCustomModule registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSecurityHealthAnalyticsSettingCustomModule(ctx *pulumi.Context,
	name string, args *OrganizationSecurityHealthAnalyticsSettingCustomModuleArgs, opts ...pulumi.ResourceOption) (*OrganizationSecurityHealthAnalyticsSettingCustomModule, error) {
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
	var resource OrganizationSecurityHealthAnalyticsSettingCustomModule
	err := ctx.RegisterResource("google-native:securitycenter/v1:OrganizationSecurityHealthAnalyticsSettingCustomModule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSecurityHealthAnalyticsSettingCustomModule gets an existing OrganizationSecurityHealthAnalyticsSettingCustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSecurityHealthAnalyticsSettingCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSecurityHealthAnalyticsSettingCustomModuleState, opts ...pulumi.ResourceOption) (*OrganizationSecurityHealthAnalyticsSettingCustomModule, error) {
	var resource OrganizationSecurityHealthAnalyticsSettingCustomModule
	err := ctx.ReadResource("google-native:securitycenter/v1:OrganizationSecurityHealthAnalyticsSettingCustomModule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSecurityHealthAnalyticsSettingCustomModule resources.
type organizationSecurityHealthAnalyticsSettingCustomModuleState struct {
}

type OrganizationSecurityHealthAnalyticsSettingCustomModuleState struct {
}

func (OrganizationSecurityHealthAnalyticsSettingCustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityHealthAnalyticsSettingCustomModuleState)(nil)).Elem()
}

type organizationSecurityHealthAnalyticsSettingCustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig *GoogleCloudSecuritycenterV1CustomConfig `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName *string `pulumi:"displayName"`
	// The enablement state of the custom module.
	EnablementState *OrganizationSecurityHealthAnalyticsSettingCustomModuleEnablementState `pulumi:"enablementState"`
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a OrganizationSecurityHealthAnalyticsSettingCustomModule resource.
type OrganizationSecurityHealthAnalyticsSettingCustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig GoogleCloudSecuritycenterV1CustomConfigPtrInput
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringPtrInput
	// The enablement state of the custom module.
	EnablementState OrganizationSecurityHealthAnalyticsSettingCustomModuleEnablementStatePtrInput
	// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (OrganizationSecurityHealthAnalyticsSettingCustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSecurityHealthAnalyticsSettingCustomModuleArgs)(nil)).Elem()
}

type OrganizationSecurityHealthAnalyticsSettingCustomModuleInput interface {
	pulumi.Input

	ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutput() OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput
	ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutputWithContext(ctx context.Context) OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput
}

func (*OrganizationSecurityHealthAnalyticsSettingCustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSecurityHealthAnalyticsSettingCustomModule)(nil)).Elem()
}

func (i *OrganizationSecurityHealthAnalyticsSettingCustomModule) ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutput() OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput {
	return i.ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutputWithContext(context.Background())
}

func (i *OrganizationSecurityHealthAnalyticsSettingCustomModule) ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutputWithContext(ctx context.Context) OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput)
}

type OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput struct{ *pulumi.OutputState }

func (OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationSecurityHealthAnalyticsSettingCustomModule)(nil)).Elem()
}

func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutput() OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput {
	return o
}

func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) ToOrganizationSecurityHealthAnalyticsSettingCustomModuleOutputWithContext(ctx context.Context) OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput {
	return o
}

// If empty, indicates that the custom module was created in the organization, folder, or project in which you are viewing the custom module. Otherwise, `ancestor_module` specifies the organization or folder from which the custom module is inherited.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) AncestorModule() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput {
		return v.AncestorModule
	}).(pulumi.StringOutput)
}

// The user specified custom configuration for the module.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) CustomConfig() GoogleCloudSecuritycenterV1CustomConfigResponseOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) GoogleCloudSecuritycenterV1CustomConfigResponseOutput {
		return v.CustomConfig
	}).(GoogleCloudSecuritycenterV1CustomConfigResponseOutput)
}

// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a lowercase letter, and contain alphanumeric characters or underscores only.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput {
		return v.DisplayName
	}).(pulumi.StringOutput)
}

// The enablement state of the custom module.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) EnablementState() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput {
		return v.EnablementState
	}).(pulumi.StringOutput)
}

// The editor that last updated the custom module.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput {
		return v.LastEditor
	}).(pulumi.StringOutput)
}

// Immutable. The resource name of the custom module. Its format is "organizations/{organization}/securityHealthAnalyticsSettings/customModules/{customModule}", or "folders/{folder}/securityHealthAnalyticsSettings/customModules/{customModule}", or "projects/{project}/securityHealthAnalyticsSettings/customModules/{customModule}" The id {customModule} is server-generated and is not user settable. It will be a numeric id containing 1-20 digits.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput {
		return v.OrganizationId
	}).(pulumi.StringOutput)
}

// The time at which the custom module was last updated.
func (o OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationSecurityHealthAnalyticsSettingCustomModule) pulumi.StringOutput {
		return v.UpdateTime
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationSecurityHealthAnalyticsSettingCustomModuleInput)(nil)).Elem(), &OrganizationSecurityHealthAnalyticsSettingCustomModule{})
	pulumi.RegisterOutputType(OrganizationSecurityHealthAnalyticsSettingCustomModuleOutput{})
}
