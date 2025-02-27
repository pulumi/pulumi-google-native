// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

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
type ApiConfigIamPolicy struct {
	pulumi.CustomResourceState

	ApiId pulumi.StringOutput `pulumi:"apiId"`
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs ApigatewayAuditConfigResponseArrayOutput `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings ApigatewayBindingResponseArrayOutput `pulumi:"bindings"`
	ConfigId pulumi.StringOutput                  `pulumi:"configId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringOutput `pulumi:"etag"`
	Location pulumi.StringOutput `pulumi:"location"`
	Project  pulumi.StringOutput `pulumi:"project"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewApiConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewApiConfigIamPolicy(ctx *pulumi.Context,
	name string, args *ApiConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*ApiConfigIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"apiId",
		"configId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApiConfigIamPolicy
	err := ctx.RegisterResource("google-native:apigateway/v1beta:ApiConfigIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiConfigIamPolicy gets an existing ApiConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiConfigIamPolicyState, opts ...pulumi.ResourceOption) (*ApiConfigIamPolicy, error) {
	var resource ApiConfigIamPolicy
	err := ctx.ReadResource("google-native:apigateway/v1beta:ApiConfigIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiConfigIamPolicy resources.
type apiConfigIamPolicyState struct {
}

type ApiConfigIamPolicyState struct {
}

func (ApiConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamPolicyState)(nil)).Elem()
}

type apiConfigIamPolicyArgs struct {
	ApiId string `pulumi:"apiId"`
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []ApigatewayAuditConfig `pulumi:"auditConfigs"`
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []ApigatewayBinding `pulumi:"bindings"`
	ConfigId string              `pulumi:"configId"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     *string `pulumi:"etag"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask *string `pulumi:"updateMask"`
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a ApiConfigIamPolicy resource.
type ApiConfigIamPolicyArgs struct {
	ApiId pulumi.StringInput
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs ApigatewayAuditConfigArrayInput
	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings ApigatewayBindingArrayInput
	ConfigId pulumi.StringInput
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// OPTIONAL: A FieldMask specifying which fields of the policy to modify. Only the fields in the mask will be modified. If no mask is provided, the following default mask is used: `paths: "bindings, etag"`
	UpdateMask pulumi.StringPtrInput
	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version pulumi.IntPtrInput
}

func (ApiConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiConfigIamPolicyArgs)(nil)).Elem()
}

type ApiConfigIamPolicyInput interface {
	pulumi.Input

	ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput
	ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput
}

func (*ApiConfigIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamPolicy)(nil)).Elem()
}

func (i *ApiConfigIamPolicy) ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput {
	return i.ToApiConfigIamPolicyOutputWithContext(context.Background())
}

func (i *ApiConfigIamPolicy) ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiConfigIamPolicyOutput)
}

type ApiConfigIamPolicyOutput struct{ *pulumi.OutputState }

func (ApiConfigIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiConfigIamPolicy)(nil)).Elem()
}

func (o ApiConfigIamPolicyOutput) ToApiConfigIamPolicyOutput() ApiConfigIamPolicyOutput {
	return o
}

func (o ApiConfigIamPolicyOutput) ToApiConfigIamPolicyOutputWithContext(ctx context.Context) ApiConfigIamPolicyOutput {
	return o
}

func (o ApiConfigIamPolicyOutput) ApiId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.ApiId }).(pulumi.StringOutput)
}

// Specifies cloud audit logging configuration for this policy.
func (o ApiConfigIamPolicyOutput) AuditConfigs() ApigatewayAuditConfigResponseArrayOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) ApigatewayAuditConfigResponseArrayOutput { return v.AuditConfigs }).(ApigatewayAuditConfigResponseArrayOutput)
}

// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
func (o ApiConfigIamPolicyOutput) Bindings() ApigatewayBindingResponseArrayOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) ApigatewayBindingResponseArrayOutput { return v.Bindings }).(ApigatewayBindingResponseArrayOutput)
}

func (o ApiConfigIamPolicyOutput) ConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.ConfigId }).(pulumi.StringOutput)
}

// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
func (o ApiConfigIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApiConfigIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ApiConfigIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o ApiConfigIamPolicyOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *ApiConfigIamPolicy) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiConfigIamPolicyInput)(nil)).Elem(), &ApiConfigIamPolicy{})
	pulumi.RegisterOutputType(ApiConfigIamPolicyOutput{})
}
