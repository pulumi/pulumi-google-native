// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new CertificateIssuanceConfig in a given project and location.
type CertificateIssuanceConfig struct {
	pulumi.CustomResourceState

	// The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
	CertificateAuthorityConfig CertificateAuthorityConfigResponseOutput `pulumi:"certificateAuthorityConfig"`
	// Required. A user-provided name of the certificate config.
	CertificateIssuanceConfigId pulumi.StringOutput `pulumi:"certificateIssuanceConfigId"`
	// The creation timestamp of a CertificateIssuanceConfig.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	Description pulumi.StringOutput `pulumi:"description"`
	// The key algorithm to use when generating the private key.
	KeyAlgorithm pulumi.StringOutput `pulumi:"keyAlgorithm"`
	// Set of labels associated with a CertificateIssuanceConfig.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Workload certificate lifetime requested.
	Lifetime pulumi.StringOutput `pulumi:"lifetime"`
	Location pulumi.StringOutput `pulumi:"location"`
	// A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally and match pattern `projects/*/locations/*/certificateIssuanceConfigs/*`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
	RotationWindowPercentage pulumi.IntOutput `pulumi:"rotationWindowPercentage"`
	// The last update timestamp of a CertificateIssuanceConfig.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCertificateIssuanceConfig registers a new resource with the given unique name, arguments, and options.
func NewCertificateIssuanceConfig(ctx *pulumi.Context,
	name string, args *CertificateIssuanceConfigArgs, opts ...pulumi.ResourceOption) (*CertificateIssuanceConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateAuthorityConfig == nil {
		return nil, errors.New("invalid value for required argument 'CertificateAuthorityConfig'")
	}
	if args.CertificateIssuanceConfigId == nil {
		return nil, errors.New("invalid value for required argument 'CertificateIssuanceConfigId'")
	}
	if args.KeyAlgorithm == nil {
		return nil, errors.New("invalid value for required argument 'KeyAlgorithm'")
	}
	if args.Lifetime == nil {
		return nil, errors.New("invalid value for required argument 'Lifetime'")
	}
	if args.RotationWindowPercentage == nil {
		return nil, errors.New("invalid value for required argument 'RotationWindowPercentage'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"certificateIssuanceConfigId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource CertificateIssuanceConfig
	err := ctx.RegisterResource("google-native:certificatemanager/v1:CertificateIssuanceConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificateIssuanceConfig gets an existing CertificateIssuanceConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificateIssuanceConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateIssuanceConfigState, opts ...pulumi.ResourceOption) (*CertificateIssuanceConfig, error) {
	var resource CertificateIssuanceConfig
	err := ctx.ReadResource("google-native:certificatemanager/v1:CertificateIssuanceConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertificateIssuanceConfig resources.
type certificateIssuanceConfigState struct {
}

type CertificateIssuanceConfigState struct {
}

func (CertificateIssuanceConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateIssuanceConfigState)(nil)).Elem()
}

type certificateIssuanceConfigArgs struct {
	// The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
	CertificateAuthorityConfig CertificateAuthorityConfig `pulumi:"certificateAuthorityConfig"`
	// Required. A user-provided name of the certificate config.
	CertificateIssuanceConfigId string `pulumi:"certificateIssuanceConfigId"`
	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	Description *string `pulumi:"description"`
	// The key algorithm to use when generating the private key.
	KeyAlgorithm CertificateIssuanceConfigKeyAlgorithm `pulumi:"keyAlgorithm"`
	// Set of labels associated with a CertificateIssuanceConfig.
	Labels map[string]string `pulumi:"labels"`
	// Workload certificate lifetime requested.
	Lifetime string  `pulumi:"lifetime"`
	Location *string `pulumi:"location"`
	// A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally and match pattern `projects/*/locations/*/certificateIssuanceConfigs/*`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
	RotationWindowPercentage int `pulumi:"rotationWindowPercentage"`
}

// The set of arguments for constructing a CertificateIssuanceConfig resource.
type CertificateIssuanceConfigArgs struct {
	// The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
	CertificateAuthorityConfig CertificateAuthorityConfigInput
	// Required. A user-provided name of the certificate config.
	CertificateIssuanceConfigId pulumi.StringInput
	// One or more paragraphs of text description of a CertificateIssuanceConfig.
	Description pulumi.StringPtrInput
	// The key algorithm to use when generating the private key.
	KeyAlgorithm CertificateIssuanceConfigKeyAlgorithmInput
	// Set of labels associated with a CertificateIssuanceConfig.
	Labels pulumi.StringMapInput
	// Workload certificate lifetime requested.
	Lifetime pulumi.StringInput
	Location pulumi.StringPtrInput
	// A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally and match pattern `projects/*/locations/*/certificateIssuanceConfigs/*`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
	RotationWindowPercentage pulumi.IntInput
}

func (CertificateIssuanceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateIssuanceConfigArgs)(nil)).Elem()
}

type CertificateIssuanceConfigInput interface {
	pulumi.Input

	ToCertificateIssuanceConfigOutput() CertificateIssuanceConfigOutput
	ToCertificateIssuanceConfigOutputWithContext(ctx context.Context) CertificateIssuanceConfigOutput
}

func (*CertificateIssuanceConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateIssuanceConfig)(nil)).Elem()
}

func (i *CertificateIssuanceConfig) ToCertificateIssuanceConfigOutput() CertificateIssuanceConfigOutput {
	return i.ToCertificateIssuanceConfigOutputWithContext(context.Background())
}

func (i *CertificateIssuanceConfig) ToCertificateIssuanceConfigOutputWithContext(ctx context.Context) CertificateIssuanceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateIssuanceConfigOutput)
}

type CertificateIssuanceConfigOutput struct{ *pulumi.OutputState }

func (CertificateIssuanceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificateIssuanceConfig)(nil)).Elem()
}

func (o CertificateIssuanceConfigOutput) ToCertificateIssuanceConfigOutput() CertificateIssuanceConfigOutput {
	return o
}

func (o CertificateIssuanceConfigOutput) ToCertificateIssuanceConfigOutputWithContext(ctx context.Context) CertificateIssuanceConfigOutput {
	return o
}

// The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
func (o CertificateIssuanceConfigOutput) CertificateAuthorityConfig() CertificateAuthorityConfigResponseOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) CertificateAuthorityConfigResponseOutput {
		return v.CertificateAuthorityConfig
	}).(CertificateAuthorityConfigResponseOutput)
}

// Required. A user-provided name of the certificate config.
func (o CertificateIssuanceConfigOutput) CertificateIssuanceConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.CertificateIssuanceConfigId }).(pulumi.StringOutput)
}

// The creation timestamp of a CertificateIssuanceConfig.
func (o CertificateIssuanceConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// One or more paragraphs of text description of a CertificateIssuanceConfig.
func (o CertificateIssuanceConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The key algorithm to use when generating the private key.
func (o CertificateIssuanceConfigOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

// Set of labels associated with a CertificateIssuanceConfig.
func (o CertificateIssuanceConfigOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Workload certificate lifetime requested.
func (o CertificateIssuanceConfigOutput) Lifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.Lifetime }).(pulumi.StringOutput)
}

func (o CertificateIssuanceConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally and match pattern `projects/*/locations/*/certificateIssuanceConfigs/*`.
func (o CertificateIssuanceConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CertificateIssuanceConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
func (o CertificateIssuanceConfigOutput) RotationWindowPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.IntOutput { return v.RotationWindowPercentage }).(pulumi.IntOutput)
}

// The last update timestamp of a CertificateIssuanceConfig.
func (o CertificateIssuanceConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CertificateIssuanceConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertificateIssuanceConfigInput)(nil)).Elem(), &CertificateIssuanceConfig{})
	pulumi.RegisterOutputType(CertificateIssuanceConfigOutput{})
}