// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a CaPool.
// Auto-naming is currently not supported for this resource.
type CaPool struct {
	pulumi.CustomResourceState

	// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
	IssuancePolicy IssuancePolicyResponseOutput `pulumi:"issuancePolicy"`
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for this CaPool in the format `projects/*/locations/*/caPools/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	PublishingOptions PublishingOptionsResponseOutput `pulumi:"publishingOptions"`
	// Immutable. The Tier of this CaPool.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewCaPool registers a new resource with the given unique name, arguments, and options.
func NewCaPool(ctx *pulumi.Context,
	name string, args *CaPoolArgs, opts ...pulumi.ResourceOption) (*CaPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPoolId == nil {
		return nil, errors.New("invalid value for required argument 'CaPoolId'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	var resource CaPool
	err := ctx.RegisterResource("google-native:privateca/v1:CaPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCaPool gets an existing CaPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCaPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CaPoolState, opts ...pulumi.ResourceOption) (*CaPool, error) {
	var resource CaPool
	err := ctx.ReadResource("google-native:privateca/v1:CaPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CaPool resources.
type caPoolState struct {
}

type CaPoolState struct {
}

func (CaPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolState)(nil)).Elem()
}

type caPoolArgs struct {
	CaPoolId string `pulumi:"caPoolId"`
	// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
	IssuancePolicy *IssuancePolicy `pulumi:"issuancePolicy"`
	// Optional. Labels with user-defined metadata.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	PublishingOptions *PublishingOptions `pulumi:"publishingOptions"`
	RequestId         *string            `pulumi:"requestId"`
	// Immutable. The Tier of this CaPool.
	Tier CaPoolTier `pulumi:"tier"`
}

// The set of arguments for constructing a CaPool resource.
type CaPoolArgs struct {
	CaPoolId pulumi.StringInput
	// Optional. The IssuancePolicy to control how Certificates will be issued from this CaPool.
	IssuancePolicy IssuancePolicyPtrInput
	// Optional. Labels with user-defined metadata.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Optional. The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
	PublishingOptions PublishingOptionsPtrInput
	RequestId         pulumi.StringPtrInput
	// Immutable. The Tier of this CaPool.
	Tier CaPoolTierInput
}

func (CaPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*caPoolArgs)(nil)).Elem()
}

type CaPoolInput interface {
	pulumi.Input

	ToCaPoolOutput() CaPoolOutput
	ToCaPoolOutputWithContext(ctx context.Context) CaPoolOutput
}

func (*CaPool) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPool)(nil)).Elem()
}

func (i *CaPool) ToCaPoolOutput() CaPoolOutput {
	return i.ToCaPoolOutputWithContext(context.Background())
}

func (i *CaPool) ToCaPoolOutputWithContext(ctx context.Context) CaPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CaPoolOutput)
}

type CaPoolOutput struct{ *pulumi.OutputState }

func (CaPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CaPool)(nil)).Elem()
}

func (o CaPoolOutput) ToCaPoolOutput() CaPoolOutput {
	return o
}

func (o CaPoolOutput) ToCaPoolOutputWithContext(ctx context.Context) CaPoolOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CaPoolInput)(nil)).Elem(), &CaPool{})
	pulumi.RegisterOutputType(CaPoolOutput{})
}