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

// Creates a new Automation in a given project and location.
// Auto-naming is currently not supported for this resource.
type Automation struct {
	pulumi.CustomResourceState

	// Optional. User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. Annotations must meet the following constraints: * Annotations are key/value pairs. * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (`/`). * The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between. * The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots(`.`), not longer than 253 characters in total, followed by a slash (`/`). See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set for more details.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Required. ID of the `Automation`.
	AutomationId pulumi.StringOutput `pulumi:"automationId"`
	// Time at which the automation was created.
	CreateTime         pulumi.StringOutput `pulumi:"createTime"`
	DeliveryPipelineId pulumi.StringOutput `pulumi:"deliveryPipelineId"`
	// Optional. Description of the `Automation`. Max length is 255 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. The weak etag of the `Automation` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 63 characters.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Name of the `Automation`. Format is `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/automations/{automation}`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// List of Automation rules associated with the Automation resource. Must have at least one rule and limited to 250 rules per Delivery Pipeline. Note: the order of the rules here is not the same as the order of execution.
	Rules AutomationRuleResponseArrayOutput `pulumi:"rules"`
	// Selected resources to which the automation will be applied.
	Selector AutomationResourceSelectorResponseOutput `pulumi:"selector"`
	// Email address of the user-managed IAM service account that creates Cloud Deploy release and rollout resources.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Optional. When Suspended, automation is deactivated from execution.
	Suspended pulumi.BoolOutput `pulumi:"suspended"`
	// Unique identifier of the `Automation`.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Time at which the automation was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAutomation registers a new resource with the given unique name, arguments, and options.
func NewAutomation(ctx *pulumi.Context,
	name string, args *AutomationArgs, opts ...pulumi.ResourceOption) (*Automation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationId == nil {
		return nil, errors.New("invalid value for required argument 'AutomationId'")
	}
	if args.DeliveryPipelineId == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryPipelineId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.Selector == nil {
		return nil, errors.New("invalid value for required argument 'Selector'")
	}
	if args.ServiceAccount == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccount'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"automationId",
		"deliveryPipelineId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Automation
	err := ctx.RegisterResource("google-native:clouddeploy/v1:Automation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutomation gets an existing Automation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutomation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutomationState, opts ...pulumi.ResourceOption) (*Automation, error) {
	var resource Automation
	err := ctx.ReadResource("google-native:clouddeploy/v1:Automation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Automation resources.
type automationState struct {
}

type AutomationState struct {
}

func (AutomationState) ElementType() reflect.Type {
	return reflect.TypeOf((*automationState)(nil)).Elem()
}

type automationArgs struct {
	// Optional. User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. Annotations must meet the following constraints: * Annotations are key/value pairs. * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (`/`). * The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between. * The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots(`.`), not longer than 253 characters in total, followed by a slash (`/`). See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set for more details.
	Annotations map[string]string `pulumi:"annotations"`
	// Required. ID of the `Automation`.
	AutomationId       string `pulumi:"automationId"`
	DeliveryPipelineId string `pulumi:"deliveryPipelineId"`
	// Optional. Description of the `Automation`. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// Optional. The weak etag of the `Automation` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 63 characters.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// List of Automation rules associated with the Automation resource. Must have at least one rule and limited to 250 rules per Delivery Pipeline. Note: the order of the rules here is not the same as the order of execution.
	Rules []AutomationRule `pulumi:"rules"`
	// Selected resources to which the automation will be applied.
	Selector AutomationResourceSelector `pulumi:"selector"`
	// Email address of the user-managed IAM service account that creates Cloud Deploy release and rollout resources.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Optional. When Suspended, automation is deactivated from execution.
	Suspended *bool `pulumi:"suspended"`
}

// The set of arguments for constructing a Automation resource.
type AutomationArgs struct {
	// Optional. User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. Annotations must meet the following constraints: * Annotations are key/value pairs. * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (`/`). * The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between. * The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots(`.`), not longer than 253 characters in total, followed by a slash (`/`). See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set for more details.
	Annotations pulumi.StringMapInput
	// Required. ID of the `Automation`.
	AutomationId       pulumi.StringInput
	DeliveryPipelineId pulumi.StringInput
	// Optional. Description of the `Automation`. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// Optional. The weak etag of the `Automation` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 63 characters.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// List of Automation rules associated with the Automation resource. Must have at least one rule and limited to 250 rules per Delivery Pipeline. Note: the order of the rules here is not the same as the order of execution.
	Rules AutomationRuleArrayInput
	// Selected resources to which the automation will be applied.
	Selector AutomationResourceSelectorInput
	// Email address of the user-managed IAM service account that creates Cloud Deploy release and rollout resources.
	ServiceAccount pulumi.StringInput
	// Optional. When Suspended, automation is deactivated from execution.
	Suspended pulumi.BoolPtrInput
}

func (AutomationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*automationArgs)(nil)).Elem()
}

type AutomationInput interface {
	pulumi.Input

	ToAutomationOutput() AutomationOutput
	ToAutomationOutputWithContext(ctx context.Context) AutomationOutput
}

func (*Automation) ElementType() reflect.Type {
	return reflect.TypeOf((**Automation)(nil)).Elem()
}

func (i *Automation) ToAutomationOutput() AutomationOutput {
	return i.ToAutomationOutputWithContext(context.Background())
}

func (i *Automation) ToAutomationOutputWithContext(ctx context.Context) AutomationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationOutput)
}

type AutomationOutput struct{ *pulumi.OutputState }

func (AutomationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Automation)(nil)).Elem()
}

func (o AutomationOutput) ToAutomationOutput() AutomationOutput {
	return o
}

func (o AutomationOutput) ToAutomationOutputWithContext(ctx context.Context) AutomationOutput {
	return o
}

// Optional. User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. Annotations must meet the following constraints: * Annotations are key/value pairs. * Valid annotation keys have two segments: an optional prefix and name, separated by a slash (`/`). * The name segment is required and must be 63 characters or less, beginning and ending with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between. * The prefix is optional. If specified, the prefix must be a DNS subdomain: a series of DNS labels separated by dots(`.`), not longer than 253 characters in total, followed by a slash (`/`). See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/#syntax-and-character-set for more details.
func (o AutomationOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Required. ID of the `Automation`.
func (o AutomationOutput) AutomationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.AutomationId }).(pulumi.StringOutput)
}

// Time at which the automation was created.
func (o AutomationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o AutomationOutput) DeliveryPipelineId() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.DeliveryPipelineId }).(pulumi.StringOutput)
}

// Optional. Description of the `Automation`. Max length is 255 characters.
func (o AutomationOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. The weak etag of the `Automation` resource. This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
func (o AutomationOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 63 characters.
func (o AutomationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o AutomationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of the `Automation`. Format is `projects/{project}/locations/{location}/deliveryPipelines/{delivery_pipeline}/automations/{automation}`.
func (o AutomationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AutomationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
func (o AutomationOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// List of Automation rules associated with the Automation resource. Must have at least one rule and limited to 250 rules per Delivery Pipeline. Note: the order of the rules here is not the same as the order of execution.
func (o AutomationOutput) Rules() AutomationRuleResponseArrayOutput {
	return o.ApplyT(func(v *Automation) AutomationRuleResponseArrayOutput { return v.Rules }).(AutomationRuleResponseArrayOutput)
}

// Selected resources to which the automation will be applied.
func (o AutomationOutput) Selector() AutomationResourceSelectorResponseOutput {
	return o.ApplyT(func(v *Automation) AutomationResourceSelectorResponseOutput { return v.Selector }).(AutomationResourceSelectorResponseOutput)
}

// Email address of the user-managed IAM service account that creates Cloud Deploy release and rollout resources.
func (o AutomationOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Optional. When Suspended, automation is deactivated from execution.
func (o AutomationOutput) Suspended() pulumi.BoolOutput {
	return o.ApplyT(func(v *Automation) pulumi.BoolOutput { return v.Suspended }).(pulumi.BoolOutput)
}

// Unique identifier of the `Automation`.
func (o AutomationOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Time at which the automation was updated.
func (o AutomationOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Automation) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationInput)(nil)).Elem(), &Automation{})
	pulumi.RegisterOutputType(AutomationOutput{})
}