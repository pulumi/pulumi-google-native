// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	iam "github.com/pulumi/pulumi-google-native/sdk/go/google/iam/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Sets the IAM policy for a taxonomy or a policy tag.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type TaxonomyPolicyTagIamMember struct {
	pulumi.CustomResourceState

	// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
	Condition iam.ConditionPtrOutput `pulumi:"condition"`
	// The etag of the resource's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`.
	Member pulumi.StringOutput `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it is not provided, a default will be supplied.
	Project pulumi.StringOutput `pulumi:"project"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewTaxonomyPolicyTagIamMember registers a new resource with the given unique name, arguments, and options.
func NewTaxonomyPolicyTagIamMember(ctx *pulumi.Context,
	name string, args *TaxonomyPolicyTagIamMemberArgs, opts ...pulumi.ResourceOption) (*TaxonomyPolicyTagIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource TaxonomyPolicyTagIamMember
	err := ctx.RegisterResource("google-native:datacatalog/v1beta1:TaxonomyPolicyTagIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTaxonomyPolicyTagIamMember gets an existing TaxonomyPolicyTagIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTaxonomyPolicyTagIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaxonomyPolicyTagIamMemberState, opts ...pulumi.ResourceOption) (*TaxonomyPolicyTagIamMember, error) {
	var resource TaxonomyPolicyTagIamMember
	err := ctx.ReadResource("google-native:datacatalog/v1beta1:TaxonomyPolicyTagIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TaxonomyPolicyTagIamMember resources.
type taxonomyPolicyTagIamMemberState struct {
}

type TaxonomyPolicyTagIamMemberState struct {
}

func (TaxonomyPolicyTagIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyPolicyTagIamMemberState)(nil)).Elem()
}

type taxonomyPolicyTagIamMemberArgs struct {
	// An IAM Condition for a given binding.
	Condition *iam.Condition `pulumi:"condition"`
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member string `pulumi:"member"`
	// The name of the resource to manage IAM policies for.
	Name string `pulumi:"name"`
	// The role that should be applied.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a TaxonomyPolicyTagIamMember resource.
type TaxonomyPolicyTagIamMemberArgs struct {
	// An IAM Condition for a given binding.
	Condition iam.ConditionPtrInput
	// Identity that will be granted the privilege in role. The entry can have one of the following values:
	//
	//  * user:{emailid}: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	//  * serviceAccount:{emailid}: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	//  * group:{emailid}: An email address that represents a Google group. For example, admins@example.com.
	//  * domain:{domain}: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Member pulumi.StringInput
	// The name of the resource to manage IAM policies for.
	Name pulumi.StringInput
	// The role that should be applied.
	Role pulumi.StringInput
}

func (TaxonomyPolicyTagIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taxonomyPolicyTagIamMemberArgs)(nil)).Elem()
}

type TaxonomyPolicyTagIamMemberInput interface {
	pulumi.Input

	ToTaxonomyPolicyTagIamMemberOutput() TaxonomyPolicyTagIamMemberOutput
	ToTaxonomyPolicyTagIamMemberOutputWithContext(ctx context.Context) TaxonomyPolicyTagIamMemberOutput
}

func (*TaxonomyPolicyTagIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyPolicyTagIamMember)(nil)).Elem()
}

func (i *TaxonomyPolicyTagIamMember) ToTaxonomyPolicyTagIamMemberOutput() TaxonomyPolicyTagIamMemberOutput {
	return i.ToTaxonomyPolicyTagIamMemberOutputWithContext(context.Background())
}

func (i *TaxonomyPolicyTagIamMember) ToTaxonomyPolicyTagIamMemberOutputWithContext(ctx context.Context) TaxonomyPolicyTagIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyPolicyTagIamMemberOutput)
}

type TaxonomyPolicyTagIamMemberOutput struct{ *pulumi.OutputState }

func (TaxonomyPolicyTagIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyPolicyTagIamMember)(nil)).Elem()
}

func (o TaxonomyPolicyTagIamMemberOutput) ToTaxonomyPolicyTagIamMemberOutput() TaxonomyPolicyTagIamMemberOutput {
	return o
}

func (o TaxonomyPolicyTagIamMemberOutput) ToTaxonomyPolicyTagIamMemberOutputWithContext(ctx context.Context) TaxonomyPolicyTagIamMemberOutput {
	return o
}

// An IAM Condition for a given binding. See https://cloud.google.com/iam/docs/conditions-overview for additional details.
func (o TaxonomyPolicyTagIamMemberOutput) Condition() iam.ConditionPtrOutput {
	return o.ApplyT(func(v *TaxonomyPolicyTagIamMember) iam.ConditionPtrOutput { return v.Condition }).(iam.ConditionPtrOutput)
}

// The etag of the resource's IAM policy.
func (o TaxonomyPolicyTagIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyPolicyTagIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`.
func (o TaxonomyPolicyTagIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyPolicyTagIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

// The name of the resource to manage IAM policies for.
func (o TaxonomyPolicyTagIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyPolicyTagIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it is not provided, a default will be supplied.
func (o TaxonomyPolicyTagIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyPolicyTagIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o TaxonomyPolicyTagIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *TaxonomyPolicyTagIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyPolicyTagIamMemberInput)(nil)).Elem(), &TaxonomyPolicyTagIamMember{})
	pulumi.RegisterOutputType(TaxonomyPolicyTagIamMemberOutput{})
}