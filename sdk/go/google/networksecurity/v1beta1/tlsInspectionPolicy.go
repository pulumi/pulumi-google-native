// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new TlsInspectionPolicy in a given project and location.
type TlsInspectionPolicy struct {
	pulumi.CustomResourceState

	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool pulumi.StringOutput `pulumi:"caPool"`
	// The timestamp when the resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	CustomTlsFeatures pulumi.StringArrayOutput `pulumi:"customTlsFeatures"`
	// Optional. Free-text description of the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trust_config. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trust_config will be accepted. This defaults to FALSE (use public CAs in addition to trust_config) for backwards compatibility, but trusting public root CAs is *not recommended* unless the traffic in question is outbound to public web servers. When possible, prefer setting this to "false" and explicitly specifying trusted CAs and certificates in a TrustConfig. Note that Secure Web Proxy does not yet honor this field.
	ExcludePublicCaSet pulumi.BoolOutput   `pulumi:"excludePublicCaSet"`
	Location           pulumi.StringOutput `pulumi:"location"`
	// Optional. Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	MinTlsVersion pulumi.StringOutput `pulumi:"minTlsVersion"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	TlsFeatureProfile pulumi.StringOutput `pulumi:"tlsFeatureProfile"`
	// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
	TlsInspectionPolicyId pulumi.StringOutput `pulumi:"tlsInspectionPolicyId"`
	// Optional. A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Note that Secure Web Proxy does not yet honor this field.
	TrustConfig pulumi.StringOutput `pulumi:"trustConfig"`
	// The timestamp when the resource was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewTlsInspectionPolicy registers a new resource with the given unique name, arguments, and options.
func NewTlsInspectionPolicy(ctx *pulumi.Context,
	name string, args *TlsInspectionPolicyArgs, opts ...pulumi.ResourceOption) (*TlsInspectionPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CaPool == nil {
		return nil, errors.New("invalid value for required argument 'CaPool'")
	}
	if args.TlsInspectionPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'TlsInspectionPolicyId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"tlsInspectionPolicyId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TlsInspectionPolicy
	err := ctx.RegisterResource("google-native:networksecurity/v1beta1:TlsInspectionPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTlsInspectionPolicy gets an existing TlsInspectionPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTlsInspectionPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TlsInspectionPolicyState, opts ...pulumi.ResourceOption) (*TlsInspectionPolicy, error) {
	var resource TlsInspectionPolicy
	err := ctx.ReadResource("google-native:networksecurity/v1beta1:TlsInspectionPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TlsInspectionPolicy resources.
type tlsInspectionPolicyState struct {
}

type TlsInspectionPolicyState struct {
}

func (TlsInspectionPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsInspectionPolicyState)(nil)).Elem()
}

type tlsInspectionPolicyArgs struct {
	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool string `pulumi:"caPool"`
	// Optional. List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	CustomTlsFeatures []string `pulumi:"customTlsFeatures"`
	// Optional. Free-text description of the resource.
	Description *string `pulumi:"description"`
	// Optional. If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trust_config. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trust_config will be accepted. This defaults to FALSE (use public CAs in addition to trust_config) for backwards compatibility, but trusting public root CAs is *not recommended* unless the traffic in question is outbound to public web servers. When possible, prefer setting this to "false" and explicitly specifying trusted CAs and certificates in a TrustConfig. Note that Secure Web Proxy does not yet honor this field.
	ExcludePublicCaSet *bool   `pulumi:"excludePublicCaSet"`
	Location           *string `pulumi:"location"`
	// Optional. Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	MinTlsVersion *TlsInspectionPolicyMinTlsVersion `pulumi:"minTlsVersion"`
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	TlsFeatureProfile *TlsInspectionPolicyTlsFeatureProfile `pulumi:"tlsFeatureProfile"`
	// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
	TlsInspectionPolicyId string `pulumi:"tlsInspectionPolicyId"`
	// Optional. A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Note that Secure Web Proxy does not yet honor this field.
	TrustConfig *string `pulumi:"trustConfig"`
}

// The set of arguments for constructing a TlsInspectionPolicy resource.
type TlsInspectionPolicyArgs struct {
	// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
	CaPool pulumi.StringInput
	// Optional. List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
	CustomTlsFeatures pulumi.StringArrayInput
	// Optional. Free-text description of the resource.
	Description pulumi.StringPtrInput
	// Optional. If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trust_config. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trust_config will be accepted. This defaults to FALSE (use public CAs in addition to trust_config) for backwards compatibility, but trusting public root CAs is *not recommended* unless the traffic in question is outbound to public web servers. When possible, prefer setting this to "false" and explicitly specifying trusted CAs and certificates in a TrustConfig. Note that Secure Web Proxy does not yet honor this field.
	ExcludePublicCaSet pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	// Optional. Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	MinTlsVersion TlsInspectionPolicyMinTlsVersionPtrInput
	// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
	TlsFeatureProfile TlsInspectionPolicyTlsFeatureProfilePtrInput
	// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
	TlsInspectionPolicyId pulumi.StringInput
	// Optional. A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Note that Secure Web Proxy does not yet honor this field.
	TrustConfig pulumi.StringPtrInput
}

func (TlsInspectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tlsInspectionPolicyArgs)(nil)).Elem()
}

type TlsInspectionPolicyInput interface {
	pulumi.Input

	ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput
	ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput
}

func (*TlsInspectionPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsInspectionPolicy)(nil)).Elem()
}

func (i *TlsInspectionPolicy) ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput {
	return i.ToTlsInspectionPolicyOutputWithContext(context.Background())
}

func (i *TlsInspectionPolicy) ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsInspectionPolicyOutput)
}

type TlsInspectionPolicyOutput struct{ *pulumi.OutputState }

func (TlsInspectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsInspectionPolicy)(nil)).Elem()
}

func (o TlsInspectionPolicyOutput) ToTlsInspectionPolicyOutput() TlsInspectionPolicyOutput {
	return o
}

func (o TlsInspectionPolicyOutput) ToTlsInspectionPolicyOutputWithContext(ctx context.Context) TlsInspectionPolicyOutput {
	return o
}

// A CA pool resource used to issue interception certificates. The CA pool string has a relative resource path following the form "projects/{project}/locations/{location}/caPools/{ca_pool}".
func (o TlsInspectionPolicyOutput) CaPool() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.CaPool }).(pulumi.StringOutput)
}

// The timestamp when the resource was created.
func (o TlsInspectionPolicyOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. List of custom TLS cipher suites selected. This field is valid only if the selected tls_feature_profile is CUSTOM. The compute.SslPoliciesService.ListAvailableFeatures method returns the set of features that can be specified in this list. Note that Secure Web Proxy does not yet honor this field.
func (o TlsInspectionPolicyOutput) CustomTlsFeatures() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringArrayOutput { return v.CustomTlsFeatures }).(pulumi.StringArrayOutput)
}

// Optional. Free-text description of the resource.
func (o TlsInspectionPolicyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. If FALSE (the default), use our default set of public CAs in addition to any CAs specified in trust_config. These public CAs are currently based on the Mozilla Root Program and are subject to change over time. If TRUE, do not accept our default set of public CAs. Only CAs specified in trust_config will be accepted. This defaults to FALSE (use public CAs in addition to trust_config) for backwards compatibility, but trusting public root CAs is *not recommended* unless the traffic in question is outbound to public web servers. When possible, prefer setting this to "false" and explicitly specifying trusted CAs and certificates in a TrustConfig. Note that Secure Web Proxy does not yet honor this field.
func (o TlsInspectionPolicyOutput) ExcludePublicCaSet() pulumi.BoolOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.BoolOutput { return v.ExcludePublicCaSet }).(pulumi.BoolOutput)
}

func (o TlsInspectionPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Optional. Minimum TLS version that the firewall should use when negotiating connections with both clients and servers. If this is not set, then the default value is to allow the broadest set of clients and servers (TLS 1.0 or higher). Setting this to more restrictive values may improve security, but may also prevent the firewall from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
func (o TlsInspectionPolicyOutput) MinTlsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.MinTlsVersion }).(pulumi.StringOutput)
}

// Name of the resource. Name is of the form projects/{project}/locations/{location}/tlsInspectionPolicies/{tls_inspection_policy} tls_inspection_policy should match the pattern:(^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o TlsInspectionPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TlsInspectionPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. The selected Profile. If this is not set, then the default value is to allow the broadest set of clients and servers ("PROFILE_COMPATIBLE"). Setting this to more restrictive values may improve security, but may also prevent the TLS inspection proxy from connecting to some clients or servers. Note that Secure Web Proxy does not yet honor this field.
func (o TlsInspectionPolicyOutput) TlsFeatureProfile() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.TlsFeatureProfile }).(pulumi.StringOutput)
}

// Required. Short name of the TlsInspectionPolicy resource to be created. This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. "tls_inspection_policy1".
func (o TlsInspectionPolicyOutput) TlsInspectionPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.TlsInspectionPolicyId }).(pulumi.StringOutput)
}

// Optional. A TrustConfig resource used when making a connection to the TLS server. This is a relative resource path following the form "projects/{project}/locations/{location}/trustConfigs/{trust_config}". This is necessary to intercept TLS connections to servers with certificates signed by a private CA or self-signed certificates. Note that Secure Web Proxy does not yet honor this field.
func (o TlsInspectionPolicyOutput) TrustConfig() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.TrustConfig }).(pulumi.StringOutput)
}

// The timestamp when the resource was updated.
func (o TlsInspectionPolicyOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TlsInspectionPolicy) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TlsInspectionPolicyInput)(nil)).Elem(), &TlsInspectionPolicy{})
	pulumi.RegisterOutputType(TlsInspectionPolicyOutput{})
}
