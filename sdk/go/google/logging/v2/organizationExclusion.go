// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new exclusion in the _Default sink in a specified parent resource. Only log entries belonging to that resource can be excluded. You can have up to 10 exclusions in a resource.
type OrganizationExclusion struct {
	pulumi.CustomResourceState

	// The creation timestamp of the exclusion.This field may not be present for older exclusions.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A description of this exclusion.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
	Filter pulumi.StringOutput `pulumi:"filter"`
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name           pulumi.StringOutput `pulumi:"name"`
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	// The last update timestamp of the exclusion.This field may not be present for older exclusions.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewOrganizationExclusion registers a new resource with the given unique name, arguments, and options.
func NewOrganizationExclusion(ctx *pulumi.Context,
	name string, args *OrganizationExclusionArgs, opts ...pulumi.ResourceOption) (*OrganizationExclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource OrganizationExclusion
	err := ctx.RegisterResource("google-native:logging/v2:OrganizationExclusion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationExclusion gets an existing OrganizationExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationExclusionState, opts ...pulumi.ResourceOption) (*OrganizationExclusion, error) {
	var resource OrganizationExclusion
	err := ctx.ReadResource("google-native:logging/v2:OrganizationExclusion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationExclusion resources.
type organizationExclusionState struct {
}

type OrganizationExclusionState struct {
}

func (OrganizationExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationExclusionState)(nil)).Elem()
}

type organizationExclusionArgs struct {
	// Optional. A description of this exclusion.
	Description *string `pulumi:"description"`
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled *bool `pulumi:"disabled"`
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
	Filter string `pulumi:"filter"`
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
}

// The set of arguments for constructing a OrganizationExclusion resource.
type OrganizationExclusionArgs struct {
	// Optional. A description of this exclusion.
	Description pulumi.StringPtrInput
	// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
	Disabled pulumi.BoolPtrInput
	// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
	Filter pulumi.StringInput
	// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
}

func (OrganizationExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationExclusionArgs)(nil)).Elem()
}

type OrganizationExclusionInput interface {
	pulumi.Input

	ToOrganizationExclusionOutput() OrganizationExclusionOutput
	ToOrganizationExclusionOutputWithContext(ctx context.Context) OrganizationExclusionOutput
}

func (*OrganizationExclusion) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationExclusion)(nil)).Elem()
}

func (i *OrganizationExclusion) ToOrganizationExclusionOutput() OrganizationExclusionOutput {
	return i.ToOrganizationExclusionOutputWithContext(context.Background())
}

func (i *OrganizationExclusion) ToOrganizationExclusionOutputWithContext(ctx context.Context) OrganizationExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationExclusionOutput)
}

type OrganizationExclusionOutput struct{ *pulumi.OutputState }

func (OrganizationExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationExclusion)(nil)).Elem()
}

func (o OrganizationExclusionOutput) ToOrganizationExclusionOutput() OrganizationExclusionOutput {
	return o
}

func (o OrganizationExclusionOutput) ToOrganizationExclusionOutputWithContext(ctx context.Context) OrganizationExclusionOutput {
	return o
}

// The creation timestamp of the exclusion.This field may not be present for older exclusions.
func (o OrganizationExclusionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A description of this exclusion.
func (o OrganizationExclusionOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. If set to True, then this exclusion is disabled and it does not exclude any log entries. You can update an exclusion to change the value of this field.
func (o OrganizationExclusionOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.BoolOutput { return v.Disabled }).(pulumi.BoolOutput)
}

// An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-queries) that matches the log entries to be excluded. By using the sample function (https://cloud.google.com/logging/docs/view/advanced-queries#sample), you can exclude less than 100% of the matching log entries.For example, the following query matches 99% of low-severity log entries from Google Cloud Storage buckets:resource.type=gcs_bucket severity<ERROR sample(insertId, 0.99)
func (o OrganizationExclusionOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// A client-assigned identifier, such as "load-balancer-exclusion". Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
func (o OrganizationExclusionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o OrganizationExclusionOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// The last update timestamp of the exclusion.This field may not be present for older exclusions.
func (o OrganizationExclusionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationExclusion) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationExclusionInput)(nil)).Elem(), &OrganizationExclusion{})
	pulumi.RegisterOutputType(OrganizationExclusionOutput{})
}