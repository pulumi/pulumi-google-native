// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new Certificate in a given Project, Location from a particular CertificateAuthority.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Certificate struct {
	pulumi.CustomResourceState

	// A structured description of the issued X.509 certificate.
	CertificateDescription CertificateDescriptionResponseOutput `pulumi:"certificateDescription"`
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config CertificateConfigResponseOutput `pulumi:"config"`
	// The time at which this Certificate was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime pulumi.StringOutput `pulumi:"lifetime"`
	// The resource path for this Certificate in the format `projects/*/locations/*/certificateAuthorities/*/certificates/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The pem-encoded, signed X.509 certificate.
	PemCertificate pulumi.StringOutput `pulumi:"pemCertificate"`
	// The chain that may be used to verify the X.509 certificate. Expected to be in issuer-to-root order according to RFC 5246.
	PemCertificateChain pulumi.StringArrayOutput `pulumi:"pemCertificateChain"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr pulumi.StringOutput `pulumi:"pemCsr"`
	// Details regarding the revocation of this Certificate. This Certificate is considered revoked if and only if this field is present.
	RevocationDetails RevocationDetailsResponseOutput `pulumi:"revocationDetails"`
	// The time at which this Certificate was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthorityId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityId'")
	}
	if args.Lifetime == nil {
		return nil, errors.New("invalid value for required argument 'Lifetime'")
	}
	var resource Certificate
	err := ctx.RegisterResource("google-native:privateca/v1beta1:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("google-native:privateca/v1beta1:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
}

type CertificateState struct {
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	CertificateAuthorityId string `pulumi:"certificateAuthorityId"`
	// Optional. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`. This field is required when using a CertificateAuthority in the Enterprise CertificateAuthority.Tier, but is optional and its value is ignored otherwise.
	CertificateId *string `pulumi:"certificateId"`
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config *CertificateConfig `pulumi:"config"`
	// Optional. Labels with user-defined metadata.
	Labels map[string]string `pulumi:"labels"`
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime string  `pulumi:"lifetime"`
	Location *string `pulumi:"location"`
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr  *string `pulumi:"pemCsr"`
	Project *string `pulumi:"project"`
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	CertificateAuthorityId pulumi.StringInput
	// Optional. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`. This field is required when using a CertificateAuthority in the Enterprise CertificateAuthority.Tier, but is optional and its value is ignored otherwise.
	CertificateId pulumi.StringPtrInput
	// Immutable. A description of the certificate and key that does not require X.509 or ASN.1.
	Config CertificateConfigPtrInput
	// Optional. Labels with user-defined metadata.
	Labels pulumi.StringMapInput
	// Immutable. The desired lifetime of a certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate. Note that the lifetime may be truncated if it would extend past the life of any certificate authority in the issuing chain.
	Lifetime pulumi.StringInput
	Location pulumi.StringPtrInput
	// Immutable. A pem-encoded X.509 certificate signing request (CSR).
	PemCsr  pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. An ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and t he request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}

type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(ctx context.Context) CertificateOutput
}

func (*Certificate) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (i *Certificate) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i *Certificate) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateInput)(nil)).Elem(), &Certificate{})
	pulumi.RegisterOutputType(CertificateOutput{})
}