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

// Creates an alias from a key/certificate pair. The structure of the request is controlled by the `format` query parameter: - `keycertfile` - Separate PEM-encoded key and certificate files are uploaded. Set `Content-Type: multipart/form-data` and include the `keyFile`, `certFile`, and `password` (if keys are encrypted) fields in the request body. If uploading to a truststore, omit `keyFile`. - `pkcs12` - A PKCS12 file is uploaded. Set `Content-Type: multipart/form-data`, provide the file in the `file` field, and include the `password` field if the file is encrypted in the request body. - `selfsignedcert` - A new private key and certificate are generated. Set `Content-Type: application/json` and include CertificateGenerationSpec in the request body.
// Auto-naming is currently not supported for this resource.
type Alias struct {
	pulumi.CustomResourceState

	// Alias for the key/certificate pair. Values must match the regular expression `[\w\s-.]{1,255}`. This must be provided for all formats except `selfsignedcert`; self-signed certs may specify the alias in either this parameter or the JSON body.
	Alias pulumi.StringOutput `pulumi:"alias"`
	// Chain of certificates under this alias.
	CertsInfo     GoogleCloudApigeeV1CertificateResponseOutput `pulumi:"certsInfo"`
	EnvironmentId pulumi.StringOutput                          `pulumi:"environmentId"`
	// Required. Format of the data. Valid values include: `selfsignedcert`, `keycertfile`, or `pkcs12`
	Format pulumi.StringOutput `pulumi:"format"`
	// Flag that specifies whether to ignore expiry validation. If set to `true`, no expiry validation will be performed.
	IgnoreExpiryValidation pulumi.BoolPtrOutput `pulumi:"ignoreExpiryValidation"`
	// Flag that specifies whether to ignore newline validation. If set to `true`, no error is thrown when the file contains a certificate chain with no newline between each certificate. Defaults to `false`.
	IgnoreNewlineValidation pulumi.BoolPtrOutput `pulumi:"ignoreNewlineValidation"`
	KeystoreId              pulumi.StringOutput  `pulumi:"keystoreId"`
	OrganizationId          pulumi.StringOutput  `pulumi:"organizationId"`
	// DEPRECATED: For improved security, specify the password in the request body instead of using the query parameter. To specify the password in the request body, set `Content-type: multipart/form-data` part with name `password`. Password for the private key file, if required.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Type of alias.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOption) (*Alias, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvironmentId == nil {
		return nil, errors.New("invalid value for required argument 'EnvironmentId'")
	}
	if args.Format == nil {
		return nil, errors.New("invalid value for required argument 'Format'")
	}
	if args.KeystoreId == nil {
		return nil, errors.New("invalid value for required argument 'KeystoreId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"environmentId",
		"format",
		"keystoreId",
		"organizationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alias
	err := ctx.RegisterResource("google-native:apigee/v1:Alias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AliasState, opts ...pulumi.ResourceOption) (*Alias, error) {
	var resource Alias
	err := ctx.ReadResource("google-native:apigee/v1:Alias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alias resources.
type aliasState struct {
}

type AliasState struct {
}

func (AliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasState)(nil)).Elem()
}

type aliasArgs struct {
	// Alias for the key/certificate pair. Values must match the regular expression `[\w\s-.]{1,255}`. This must be provided for all formats except `selfsignedcert`; self-signed certs may specify the alias in either this parameter or the JSON body.
	Alias *string `pulumi:"alias"`
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data          *string `pulumi:"data"`
	EnvironmentId string  `pulumi:"environmentId"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions []map[string]interface{} `pulumi:"extensions"`
	// File to upload.
	File pulumi.AssetOrArchive `pulumi:"file"`
	// Required. Format of the data. Valid values include: `selfsignedcert`, `keycertfile`, or `pkcs12`
	Format string `pulumi:"format"`
	// Flag that specifies whether to ignore expiry validation. If set to `true`, no expiry validation will be performed.
	IgnoreExpiryValidation *bool `pulumi:"ignoreExpiryValidation"`
	// Flag that specifies whether to ignore newline validation. If set to `true`, no error is thrown when the file contains a certificate chain with no newline between each certificate. Defaults to `false`.
	IgnoreNewlineValidation *bool  `pulumi:"ignoreNewlineValidation"`
	KeystoreId              string `pulumi:"keystoreId"`
	OrganizationId          string `pulumi:"organizationId"`
	// DEPRECATED: For improved security, specify the password in the request body instead of using the query parameter. To specify the password in the request body, set `Content-type: multipart/form-data` part with name `password`. Password for the private key file, if required.
	Password *string `pulumi:"password"`
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// Alias for the key/certificate pair. Values must match the regular expression `[\w\s-.]{1,255}`. This must be provided for all formats except `selfsignedcert`; self-signed certs may specify the alias in either this parameter or the JSON body.
	Alias pulumi.StringPtrInput
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType pulumi.StringPtrInput
	// The HTTP request/response body as raw binary.
	Data          pulumi.StringPtrInput
	EnvironmentId pulumi.StringInput
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions pulumi.MapArrayInput
	// File to upload.
	File pulumi.AssetOrArchiveInput
	// Required. Format of the data. Valid values include: `selfsignedcert`, `keycertfile`, or `pkcs12`
	Format pulumi.StringInput
	// Flag that specifies whether to ignore expiry validation. If set to `true`, no expiry validation will be performed.
	IgnoreExpiryValidation pulumi.BoolPtrInput
	// Flag that specifies whether to ignore newline validation. If set to `true`, no error is thrown when the file contains a certificate chain with no newline between each certificate. Defaults to `false`.
	IgnoreNewlineValidation pulumi.BoolPtrInput
	KeystoreId              pulumi.StringInput
	OrganizationId          pulumi.StringInput
	// DEPRECATED: For improved security, specify the password in the request body instead of using the query parameter. To specify the password in the request body, set `Content-type: multipart/form-data` part with name `password`. Password for the private key file, if required.
	Password pulumi.StringPtrInput
}

func (AliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aliasArgs)(nil)).Elem()
}

type AliasInput interface {
	pulumi.Input

	ToAliasOutput() AliasOutput
	ToAliasOutputWithContext(ctx context.Context) AliasOutput
}

func (*Alias) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (i *Alias) ToAliasOutput() AliasOutput {
	return i.ToAliasOutputWithContext(context.Background())
}

func (i *Alias) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AliasOutput)
}

type AliasOutput struct{ *pulumi.OutputState }

func (AliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alias)(nil)).Elem()
}

func (o AliasOutput) ToAliasOutput() AliasOutput {
	return o
}

func (o AliasOutput) ToAliasOutputWithContext(ctx context.Context) AliasOutput {
	return o
}

// Alias for the key/certificate pair. Values must match the regular expression `[\w\s-.]{1,255}`. This must be provided for all formats except `selfsignedcert`; self-signed certs may specify the alias in either this parameter or the JSON body.
func (o AliasOutput) Alias() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Alias }).(pulumi.StringOutput)
}

// Chain of certificates under this alias.
func (o AliasOutput) CertsInfo() GoogleCloudApigeeV1CertificateResponseOutput {
	return o.ApplyT(func(v *Alias) GoogleCloudApigeeV1CertificateResponseOutput { return v.CertsInfo }).(GoogleCloudApigeeV1CertificateResponseOutput)
}

func (o AliasOutput) EnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.EnvironmentId }).(pulumi.StringOutput)
}

// Required. Format of the data. Valid values include: `selfsignedcert`, `keycertfile`, or `pkcs12`
func (o AliasOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Format }).(pulumi.StringOutput)
}

// Flag that specifies whether to ignore expiry validation. If set to `true`, no expiry validation will be performed.
func (o AliasOutput) IgnoreExpiryValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Alias) pulumi.BoolPtrOutput { return v.IgnoreExpiryValidation }).(pulumi.BoolPtrOutput)
}

// Flag that specifies whether to ignore newline validation. If set to `true`, no error is thrown when the file contains a certificate chain with no newline between each certificate. Defaults to `false`.
func (o AliasOutput) IgnoreNewlineValidation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Alias) pulumi.BoolPtrOutput { return v.IgnoreNewlineValidation }).(pulumi.BoolPtrOutput)
}

func (o AliasOutput) KeystoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.KeystoreId }).(pulumi.StringOutput)
}

func (o AliasOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

// DEPRECATED: For improved security, specify the password in the request body instead of using the query parameter. To specify the password in the request body, set `Content-type: multipart/form-data` part with name `password`. Password for the private key file, if required.
func (o AliasOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Type of alias.
func (o AliasOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Alias) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AliasInput)(nil)).Elem(), &Alias{})
	pulumi.RegisterOutputType(AliasOutput{})
}
