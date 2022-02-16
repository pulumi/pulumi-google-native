// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new HMAC key for the specified service account.
// Auto-naming is currently not supported for this resource.
type HmacKey struct {
	pulumi.CustomResourceState

	// The ID of the HMAC Key.
	AccessId pulumi.StringOutput `pulumi:"accessId"`
	// HTTP 1.1 Entity tag for the HMAC key.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The kind of item this is. For HMAC Key metadata, this is always storage#hmacKeyMetadata.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Project ID owning the service account to which the key authenticates.
	Project pulumi.StringOutput `pulumi:"project"`
	// The link to this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The email address of the key's associated service account.
	ServiceAccountEmail pulumi.StringOutput `pulumi:"serviceAccountEmail"`
	// The state of the key. Can be one of ACTIVE, INACTIVE, or DELETED.
	State pulumi.StringOutput `pulumi:"state"`
	// The creation time of the HMAC key in RFC 3339 format.
	TimeCreated pulumi.StringOutput `pulumi:"timeCreated"`
	// The last modification time of the HMAC key metadata in RFC 3339 format.
	Updated pulumi.StringOutput `pulumi:"updated"`
}

// NewHmacKey registers a new resource with the given unique name, arguments, and options.
func NewHmacKey(ctx *pulumi.Context,
	name string, args *HmacKeyArgs, opts ...pulumi.ResourceOption) (*HmacKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountEmail == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountEmail'")
	}
	var resource HmacKey
	err := ctx.RegisterResource("google-native:storage/v1:HmacKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHmacKey gets an existing HmacKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHmacKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HmacKeyState, opts ...pulumi.ResourceOption) (*HmacKey, error) {
	var resource HmacKey
	err := ctx.ReadResource("google-native:storage/v1:HmacKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HmacKey resources.
type hmacKeyState struct {
}

type HmacKeyState struct {
}

func (HmacKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*hmacKeyState)(nil)).Elem()
}

type hmacKeyArgs struct {
	Project             *string `pulumi:"project"`
	ServiceAccountEmail string  `pulumi:"serviceAccountEmail"`
	UserProject         *string `pulumi:"userProject"`
}

// The set of arguments for constructing a HmacKey resource.
type HmacKeyArgs struct {
	Project             pulumi.StringPtrInput
	ServiceAccountEmail pulumi.StringInput
	UserProject         pulumi.StringPtrInput
}

func (HmacKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hmacKeyArgs)(nil)).Elem()
}

type HmacKeyInput interface {
	pulumi.Input

	ToHmacKeyOutput() HmacKeyOutput
	ToHmacKeyOutputWithContext(ctx context.Context) HmacKeyOutput
}

func (*HmacKey) ElementType() reflect.Type {
	return reflect.TypeOf((**HmacKey)(nil)).Elem()
}

func (i *HmacKey) ToHmacKeyOutput() HmacKeyOutput {
	return i.ToHmacKeyOutputWithContext(context.Background())
}

func (i *HmacKey) ToHmacKeyOutputWithContext(ctx context.Context) HmacKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HmacKeyOutput)
}

type HmacKeyOutput struct{ *pulumi.OutputState }

func (HmacKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HmacKey)(nil)).Elem()
}

func (o HmacKeyOutput) ToHmacKeyOutput() HmacKeyOutput {
	return o
}

func (o HmacKeyOutput) ToHmacKeyOutputWithContext(ctx context.Context) HmacKeyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HmacKeyInput)(nil)).Elem(), &HmacKey{})
	pulumi.RegisterOutputType(HmacKeyOutput{})
}