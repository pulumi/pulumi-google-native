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

// Sets the access control policy on the specified resource. Replaces any existing policy. Can return `NOT_FOUND`, `INVALID_ARGUMENT`, and `PERMISSION_DENIED` errors.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type InstanceBackupIamPolicy struct {
	pulumi.CustomResourceState

	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	BackupId     pulumi.StringOutput            `pulumi:"backupId"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingResponseArrayOutput `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag       pulumi.StringOutput `pulumi:"etag"`
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	Location   pulumi.StringOutput `pulumi:"location"`
	Project    pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewInstanceBackupIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewInstanceBackupIamPolicy(ctx *pulumi.Context,
	name string, args *InstanceBackupIamPolicyArgs, opts ...pulumi.ResourceOption) (*InstanceBackupIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"backupId",
		"instanceId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstanceBackupIamPolicy
	err := ctx.RegisterResource("google-native:looker/v1:InstanceBackupIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceBackupIamPolicy gets an existing InstanceBackupIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceBackupIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceBackupIamPolicyState, opts ...pulumi.ResourceOption) (*InstanceBackupIamPolicy, error) {
	var resource InstanceBackupIamPolicy
	err := ctx.ReadResource("google-native:looker/v1:InstanceBackupIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceBackupIamPolicy resources.
type instanceBackupIamPolicyState struct {
}

type InstanceBackupIamPolicyState struct {
}

func (InstanceBackupIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceBackupIamPolicyState)(nil)).Elem()
}

type instanceBackupIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `pulumi:"auditConfigs"`
	BackupId     string        `pulumi:"backupId"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `pulumi:"bindings"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag       *string `pulumi:"etag"`
	InstanceId string  `pulumi:"instanceId"`
	Location   *string `pulumi:"location"`
	Project    *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a InstanceBackupIamPolicy resource.
type InstanceBackupIamPolicyArgs struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs AuditConfigArrayInput
	BackupId     pulumi.StringInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings BindingArrayInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag       pulumi.StringPtrInput
	InstanceId pulumi.StringInput
	Location   pulumi.StringPtrInput
	Project    pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (InstanceBackupIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceBackupIamPolicyArgs)(nil)).Elem()
}

type InstanceBackupIamPolicyInput interface {
	pulumi.Input

	ToInstanceBackupIamPolicyOutput() InstanceBackupIamPolicyOutput
	ToInstanceBackupIamPolicyOutputWithContext(ctx context.Context) InstanceBackupIamPolicyOutput
}

func (*InstanceBackupIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupIamPolicy)(nil)).Elem()
}

func (i *InstanceBackupIamPolicy) ToInstanceBackupIamPolicyOutput() InstanceBackupIamPolicyOutput {
	return i.ToInstanceBackupIamPolicyOutputWithContext(context.Background())
}

func (i *InstanceBackupIamPolicy) ToInstanceBackupIamPolicyOutputWithContext(ctx context.Context) InstanceBackupIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceBackupIamPolicyOutput)
}

type InstanceBackupIamPolicyOutput struct{ *pulumi.OutputState }

func (InstanceBackupIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceBackupIamPolicy)(nil)).Elem()
}

func (o InstanceBackupIamPolicyOutput) ToInstanceBackupIamPolicyOutput() InstanceBackupIamPolicyOutput {
	return o
}

func (o InstanceBackupIamPolicyOutput) ToInstanceBackupIamPolicyOutputWithContext(ctx context.Context) InstanceBackupIamPolicyOutput {
	return o
}

// Specifies cloud audit logging configuration for this policy.
func (o InstanceBackupIamPolicyOutput) AuditConfigs() AuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) AuditConfigResponseArrayOutput { return v.AuditConfigs }).(AuditConfigResponseArrayOutput)
}

func (o InstanceBackupIamPolicyOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o InstanceBackupIamPolicyOutput) Bindings() BindingResponseArrayOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) BindingResponseArrayOutput { return v.Bindings }).(BindingResponseArrayOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o InstanceBackupIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o InstanceBackupIamPolicyOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o InstanceBackupIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o InstanceBackupIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o InstanceBackupIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceBackupIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceBackupIamPolicyInput)(nil)).Elem(), &InstanceBackupIamPolicy{})
	pulumi.RegisterOutputType(InstanceBackupIamPolicyOutput{})
}