// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single GatewaySecurityPolicyRule.
func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRuleResult
	err := ctx.Invoke("google-native:networksecurity/v1beta1:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleArgs struct {
	GatewaySecurityPolicyId string  `pulumi:"gatewaySecurityPolicyId"`
	Location                string  `pulumi:"location"`
	Project                 *string `pulumi:"project"`
	RuleId                  string  `pulumi:"ruleId"`
}

type LookupRuleResult struct {
	// Optional. CEL expression for matching on L7/application level criteria.
	ApplicationMatcher string `pulumi:"applicationMatcher"`
	// Profile which tells what the primitive action should be.
	BasicProfile string `pulumi:"basicProfile"`
	// Time when the rule was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. Free-text description of the resource.
	Description string `pulumi:"description"`
	// Whether the rule is enforced.
	Enabled bool `pulumi:"enabled"`
	// Immutable. Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule} rule should match the pattern: (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
	Name string `pulumi:"name"`
	// Priority of the rule. Lower number corresponds to higher precedence.
	Priority int `pulumi:"priority"`
	// CEL expression for matching on session criteria.
	SessionMatcher string `pulumi:"sessionMatcher"`
	// Optional. Flag to enable TLS inspection of traffic matching on , can only be true if the parent GatewaySecurityPolicy references a TLSInspectionConfig.
	TlsInspectionEnabled bool `pulumi:"tlsInspectionEnabled"`
	// Time when the rule was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupRuleOutput(ctx *pulumi.Context, args LookupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleResult, error) {
			args := v.(LookupRuleArgs)
			r, err := LookupRule(ctx, &args, opts...)
			var s LookupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleResultOutput)
}

type LookupRuleOutputArgs struct {
	GatewaySecurityPolicyId pulumi.StringInput    `pulumi:"gatewaySecurityPolicyId"`
	Location                pulumi.StringInput    `pulumi:"location"`
	Project                 pulumi.StringPtrInput `pulumi:"project"`
	RuleId                  pulumi.StringInput    `pulumi:"ruleId"`
}

func (LookupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleArgs)(nil)).Elem()
}

type LookupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleResult)(nil)).Elem()
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutput() LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutputWithContext(ctx context.Context) LookupRuleResultOutput {
	return o
}

// Optional. CEL expression for matching on L7/application level criteria.
func (o LookupRuleResultOutput) ApplicationMatcher() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.ApplicationMatcher }).(pulumi.StringOutput)
}

// Profile which tells what the primitive action should be.
func (o LookupRuleResultOutput) BasicProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.BasicProfile }).(pulumi.StringOutput)
}

// Time when the rule was created.
func (o LookupRuleResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Free-text description of the resource.
func (o LookupRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// Whether the rule is enforced.
func (o LookupRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// Immutable. Name of the resource. ame is the full resource name so projects/{project}/locations/{location}/gatewaySecurityPolicies/{gateway_security_policy}/rules/{rule} rule should match the pattern: (^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$).
func (o LookupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Priority of the rule. Lower number corresponds to higher precedence.
func (o LookupRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

// CEL expression for matching on session criteria.
func (o LookupRuleResultOutput) SessionMatcher() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.SessionMatcher }).(pulumi.StringOutput)
}

// Optional. Flag to enable TLS inspection of traffic matching on , can only be true if the parent GatewaySecurityPolicy references a TLSInspectionConfig.
func (o LookupRuleResultOutput) TlsInspectionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRuleResult) bool { return v.TlsInspectionEnabled }).(pulumi.BoolOutput)
}

// Time when the rule was updated.
func (o LookupRuleResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleResultOutput{})
}