// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a ServiceAccountKey.
// Auto-naming is currently not supported for this resource.
type Key struct {
	pulumi.CustomResourceState

	// The key status.
	Disabled pulumi.BoolOutput `pulumi:"disabled"`
	// Specifies the algorithm (and possibly key size) for the key.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// The key origin.
	KeyOrigin pulumi.StringOutput `pulumi:"keyOrigin"`
	// The key type.
	KeyType pulumi.StringOutput `pulumi:"keyType"`
	// The resource name of the service account key in the following format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The private key data. Only provided in `CreateServiceAccountKey` responses. Make sure to keep the private key data secure because it allows for the assertion of the service account identity. When base64 decoded, the private key data can be used to authenticate with Google API client libraries and with gcloud auth activate-service-account.
	PrivateKeyData pulumi.StringOutput `pulumi:"privateKeyData"`
	// The output format for the private key. Only provided in `CreateServiceAccountKey` responses, not in `GetServiceAccountKey` or `ListServiceAccountKey` responses. Google never exposes system-managed private keys, and never retains user-managed private keys.
	PrivateKeyType pulumi.StringOutput `pulumi:"privateKeyType"`
	// The public key data. Only provided in `GetServiceAccountKey` responses.
	PublicKeyData pulumi.StringOutput `pulumi:"publicKeyData"`
	// The key can be used after this timestamp.
	ValidAfterTime pulumi.StringOutput `pulumi:"validAfterTime"`
	// The key can be used before this timestamp. For system-managed key pairs, this timestamp is the end time for the private key signing operation. The public key could still be used for verification for a few hours after this time.
	ValidBeforeTime pulumi.StringOutput `pulumi:"validBeforeTime"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceAccountId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAccountId'")
	}
	var resource Key
	err := ctx.RegisterResource("google-native:iam/v1:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("google-native:iam/v1:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// Which type of key and algorithm to use for the key. The default is currently a 2K RSA key. However this may change in the future.
	KeyAlgorithm *KeyKeyAlgorithm `pulumi:"keyAlgorithm"`
	// The output format of the private key. The default value is `TYPE_GOOGLE_CREDENTIALS_FILE`, which is the Google Credentials File format.
	PrivateKeyType   *KeyPrivateKeyType `pulumi:"privateKeyType"`
	Project          *string            `pulumi:"project"`
	ServiceAccountId string             `pulumi:"serviceAccountId"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// Which type of key and algorithm to use for the key. The default is currently a 2K RSA key. However this may change in the future.
	KeyAlgorithm KeyKeyAlgorithmPtrInput
	// The output format of the private key. The default value is `TYPE_GOOGLE_CREDENTIALS_FILE`, which is the Google Credentials File format.
	PrivateKeyType   KeyPrivateKeyTypePtrInput
	Project          pulumi.StringPtrInput
	ServiceAccountId pulumi.StringInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Key)(nil)).Elem()
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterOutputType(KeyOutput{})
}