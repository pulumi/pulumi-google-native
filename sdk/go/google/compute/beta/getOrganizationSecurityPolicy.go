// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List all of the ordered rules present in a single specified policy.
func LookupOrganizationSecurityPolicy(ctx *pulumi.Context, args *LookupOrganizationSecurityPolicyArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationSecurityPolicyResult, error) {
	var rv LookupOrganizationSecurityPolicyResult
	err := ctx.Invoke("google-native:compute/beta:getOrganizationSecurityPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationSecurityPolicyArgs struct {
	SecurityPolicy string `pulumi:"securityPolicy"`
}

type LookupOrganizationSecurityPolicyResult struct {
	AdaptiveProtectionConfig SecurityPolicyAdaptiveProtectionConfigResponse `pulumi:"adaptiveProtectionConfig"`
	AdvancedOptionsConfig    SecurityPolicyAdvancedOptionsConfigResponse    `pulumi:"advancedOptionsConfig"`
	// A list of associations that belong to this policy.
	Associations []SecurityPolicyAssociationResponse `pulumi:"associations"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description string `pulumi:"description"`
	// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	DisplayName string `pulumi:"displayName"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
	Fingerprint string `pulumi:"fingerprint"`
	// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
	Kind string `pulumi:"kind"`
	// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
	LabelFingerprint string `pulumi:"labelFingerprint"`
	// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `pulumi:"name"`
	// The parent of the security policy.
	Parent                 string                                       `pulumi:"parent"`
	RecaptchaOptionsConfig SecurityPolicyRecaptchaOptionsConfigResponse `pulumi:"recaptchaOptionsConfig"`
	// Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
	RuleTupleCount int `pulumi:"ruleTupleCount"`
	// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
	Rules []SecurityPolicyRuleResponse `pulumi:"rules"`
	// Server-defined URL for the resource.
	SelfLink string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId string `pulumi:"selfLinkWithId"`
	// The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
	Type string `pulumi:"type"`
}

func LookupOrganizationSecurityPolicyOutput(ctx *pulumi.Context, args LookupOrganizationSecurityPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationSecurityPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationSecurityPolicyResult, error) {
			args := v.(LookupOrganizationSecurityPolicyArgs)
			r, err := LookupOrganizationSecurityPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupOrganizationSecurityPolicyResultOutput)
}

type LookupOrganizationSecurityPolicyOutputArgs struct {
	SecurityPolicy pulumi.StringInput `pulumi:"securityPolicy"`
}

func (LookupOrganizationSecurityPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationSecurityPolicyArgs)(nil)).Elem()
}

type LookupOrganizationSecurityPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationSecurityPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationSecurityPolicyResult)(nil)).Elem()
}

func (o LookupOrganizationSecurityPolicyResultOutput) ToLookupOrganizationSecurityPolicyResultOutput() LookupOrganizationSecurityPolicyResultOutput {
	return o
}

func (o LookupOrganizationSecurityPolicyResultOutput) ToLookupOrganizationSecurityPolicyResultOutputWithContext(ctx context.Context) LookupOrganizationSecurityPolicyResultOutput {
	return o
}

func (o LookupOrganizationSecurityPolicyResultOutput) AdaptiveProtectionConfig() SecurityPolicyAdaptiveProtectionConfigResponseOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) SecurityPolicyAdaptiveProtectionConfigResponse {
		return v.AdaptiveProtectionConfig
	}).(SecurityPolicyAdaptiveProtectionConfigResponseOutput)
}

func (o LookupOrganizationSecurityPolicyResultOutput) AdvancedOptionsConfig() SecurityPolicyAdvancedOptionsConfigResponseOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) SecurityPolicyAdvancedOptionsConfigResponse {
		return v.AdvancedOptionsConfig
	}).(SecurityPolicyAdvancedOptionsConfigResponseOutput)
}

// A list of associations that belong to this policy.
func (o LookupOrganizationSecurityPolicyResultOutput) Associations() SecurityPolicyAssociationResponseArrayOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) []SecurityPolicyAssociationResponse {
		return v.Associations
	}).(SecurityPolicyAssociationResponseArrayOutput)
}

// Creation timestamp in RFC3339 text format.
func (o LookupOrganizationSecurityPolicyResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o LookupOrganizationSecurityPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupOrganizationSecurityPolicyResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
func (o LookupOrganizationSecurityPolicyResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// [Output only] Type of the resource. Always compute#securityPolicyfor security policies
func (o LookupOrganizationSecurityPolicyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
func (o LookupOrganizationSecurityPolicyResultOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
func (o LookupOrganizationSecurityPolicyResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o LookupOrganizationSecurityPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// The parent of the security policy.
func (o LookupOrganizationSecurityPolicyResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o LookupOrganizationSecurityPolicyResultOutput) RecaptchaOptionsConfig() SecurityPolicyRecaptchaOptionsConfigResponseOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) SecurityPolicyRecaptchaOptionsConfigResponse {
		return v.RecaptchaOptionsConfig
	}).(SecurityPolicyRecaptchaOptionsConfigResponseOutput)
}

// Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
func (o LookupOrganizationSecurityPolicyResultOutput) RuleTupleCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) int { return v.RuleTupleCount }).(pulumi.IntOutput)
}

// A list of rules that belong to this policy. There must always be a default rule (rule with priority 2147483647 and match "*"). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
func (o LookupOrganizationSecurityPolicyResultOutput) Rules() SecurityPolicyRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) []SecurityPolicyRuleResponse { return v.Rules }).(SecurityPolicyRuleResponseArrayOutput)
}

// Server-defined URL for the resource.
func (o LookupOrganizationSecurityPolicyResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o LookupOrganizationSecurityPolicyResultOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// The type indicates the intended use of the security policy. CLOUD_ARMOR - Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache.
func (o LookupOrganizationSecurityPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationSecurityPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationSecurityPolicyResultOutput{})
}