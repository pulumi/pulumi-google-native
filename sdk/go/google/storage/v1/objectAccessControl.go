// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new ACL entry on the specified object.
// Auto-naming is currently not supported for this resource.
type ObjectAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity, if any.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity, if any.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The ID for the entity, if any.
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// HTTP 1.1 Entity tag for the access-control entry.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The content generation of the object, if applied to an object.
	Generation pulumi.StringOutput `pulumi:"generation"`
	// The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the object, if applied to an object.
	Object pulumi.StringOutput `pulumi:"object"`
	// The project team associated with the entity, if any.
	ProjectTeam ObjectAccessControlProjectTeamResponseOutput `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role pulumi.StringOutput `pulumi:"role"`
	// The link to this access-control entry.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewObjectAccessControl registers a new resource with the given unique name, arguments, and options.
func NewObjectAccessControl(ctx *pulumi.Context,
	name string, args *ObjectAccessControlArgs, opts ...pulumi.ResourceOption) (*ObjectAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Object == nil {
		return nil, errors.New("invalid value for required argument 'Object'")
	}
	var resource ObjectAccessControl
	err := ctx.RegisterResource("google-native:storage/v1:ObjectAccessControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectAccessControl gets an existing ObjectAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectAccessControlState, opts ...pulumi.ResourceOption) (*ObjectAccessControl, error) {
	var resource ObjectAccessControl
	err := ctx.ReadResource("google-native:storage/v1:ObjectAccessControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectAccessControl resources.
type objectAccessControlState struct {
}

type ObjectAccessControlState struct {
}

func (ObjectAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAccessControlState)(nil)).Elem()
}

type objectAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The domain associated with the entity, if any.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity, if any.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity *string `pulumi:"entity"`
	// The ID for the entity, if any.
	EntityId *string `pulumi:"entityId"`
	// HTTP 1.1 Entity tag for the access-control entry.
	Etag *string `pulumi:"etag"`
	// The content generation of the object, if applied to an object.
	Generation *string `pulumi:"generation"`
	// The ID of the access-control entry.
	Id *string `pulumi:"id"`
	// The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
	Kind *string `pulumi:"kind"`
	// The name of the object, if applied to an object.
	Object string `pulumi:"object"`
	// The project team associated with the entity, if any.
	ProjectTeam *ObjectAccessControlProjectTeam `pulumi:"projectTeam"`
	// The access permission for the entity.
	Role *string `pulumi:"role"`
	// The link to this access-control entry.
	SelfLink *string `pulumi:"selfLink"`
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject *string `pulumi:"userProject"`
}

// The set of arguments for constructing a ObjectAccessControl resource.
type ObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The domain associated with the entity, if any.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity, if any.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms:
	// - user-userId
	// - user-email
	// - group-groupId
	// - group-email
	// - domain-domain
	// - project-team-projectId
	// - allUsers
	// - allAuthenticatedUsers Examples:
	// - The user liz@example.com would be user-liz@example.com.
	// - The group example@googlegroups.com would be group-example@googlegroups.com.
	// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringPtrInput
	// The ID for the entity, if any.
	EntityId pulumi.StringPtrInput
	// HTTP 1.1 Entity tag for the access-control entry.
	Etag pulumi.StringPtrInput
	// The content generation of the object, if applied to an object.
	Generation pulumi.StringPtrInput
	// The ID of the access-control entry.
	Id pulumi.StringPtrInput
	// The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
	Kind pulumi.StringPtrInput
	// The name of the object, if applied to an object.
	Object pulumi.StringInput
	// The project team associated with the entity, if any.
	ProjectTeam ObjectAccessControlProjectTeamPtrInput
	// The access permission for the entity.
	Role pulumi.StringPtrInput
	// The link to this access-control entry.
	SelfLink pulumi.StringPtrInput
	// The project to be billed for this request. Required for Requester Pays buckets.
	UserProject pulumi.StringPtrInput
}

func (ObjectAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectAccessControlArgs)(nil)).Elem()
}

type ObjectAccessControlInput interface {
	pulumi.Input

	ToObjectAccessControlOutput() ObjectAccessControlOutput
	ToObjectAccessControlOutputWithContext(ctx context.Context) ObjectAccessControlOutput
}

func (*ObjectAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAccessControl)(nil)).Elem()
}

func (i *ObjectAccessControl) ToObjectAccessControlOutput() ObjectAccessControlOutput {
	return i.ToObjectAccessControlOutputWithContext(context.Background())
}

func (i *ObjectAccessControl) ToObjectAccessControlOutputWithContext(ctx context.Context) ObjectAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ObjectAccessControlOutput)
}

type ObjectAccessControlOutput struct{ *pulumi.OutputState }

func (ObjectAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ObjectAccessControl)(nil)).Elem()
}

func (o ObjectAccessControlOutput) ToObjectAccessControlOutput() ObjectAccessControlOutput {
	return o
}

func (o ObjectAccessControlOutput) ToObjectAccessControlOutputWithContext(ctx context.Context) ObjectAccessControlOutput {
	return o
}

// The name of the bucket.
func (o ObjectAccessControlOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The domain associated with the entity, if any.
func (o ObjectAccessControlOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The email address associated with the entity, if any.
func (o ObjectAccessControlOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The entity holding the permission, in one of the following forms:
// - user-userId
// - user-email
// - group-groupId
// - group-email
// - domain-domain
// - project-team-projectId
// - allUsers
// - allAuthenticatedUsers Examples:
// - The user liz@example.com would be user-liz@example.com.
// - The group example@googlegroups.com would be group-example@googlegroups.com.
// - To refer to all members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
func (o ObjectAccessControlOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// The ID for the entity, if any.
func (o ObjectAccessControlOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// HTTP 1.1 Entity tag for the access-control entry.
func (o ObjectAccessControlOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The content generation of the object, if applied to an object.
func (o ObjectAccessControlOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Generation }).(pulumi.StringOutput)
}

// The kind of item this is. For object access control entries, this is always storage#objectAccessControl.
func (o ObjectAccessControlOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The name of the object, if applied to an object.
func (o ObjectAccessControlOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Object }).(pulumi.StringOutput)
}

// The project team associated with the entity, if any.
func (o ObjectAccessControlOutput) ProjectTeam() ObjectAccessControlProjectTeamResponseOutput {
	return o.ApplyT(func(v *ObjectAccessControl) ObjectAccessControlProjectTeamResponseOutput { return v.ProjectTeam }).(ObjectAccessControlProjectTeamResponseOutput)
}

// The access permission for the entity.
func (o ObjectAccessControlOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

// The link to this access-control entry.
func (o ObjectAccessControlOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ObjectAccessControl) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ObjectAccessControlInput)(nil)).Elem(), &ObjectAccessControl{})
	pulumi.RegisterOutputType(ObjectAccessControlOutput{})
}