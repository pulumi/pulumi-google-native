// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified firewall rule.
func LookupIngressRule(ctx *pulumi.Context, args *LookupIngressRuleArgs, opts ...pulumi.InvokeOption) (*LookupIngressRuleResult, error) {
	var rv LookupIngressRuleResult
	err := ctx.Invoke("google-native:appengine/v1beta:getIngressRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIngressRuleArgs struct {
	AppId         string `pulumi:"appId"`
	IngressRuleId string `pulumi:"ingressRuleId"`
}

type LookupIngressRuleResult struct {
	// The action to take on matched requests.
	Action string `pulumi:"action"`
	// An optional string description of this rule. This field has a maximum length of 400 characters.
	Description string `pulumi:"description"`
	// A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
	Priority int `pulumi:"priority"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
	SourceRange string `pulumi:"sourceRange"`
}

func LookupIngressRuleOutput(ctx *pulumi.Context, args LookupIngressRuleOutputArgs, opts ...pulumi.InvokeOption) LookupIngressRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIngressRuleResult, error) {
			args := v.(LookupIngressRuleArgs)
			r, err := LookupIngressRule(ctx, &args, opts...)
			var s LookupIngressRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIngressRuleResultOutput)
}

type LookupIngressRuleOutputArgs struct {
	AppId         pulumi.StringInput `pulumi:"appId"`
	IngressRuleId pulumi.StringInput `pulumi:"ingressRuleId"`
}

func (LookupIngressRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIngressRuleArgs)(nil)).Elem()
}

type LookupIngressRuleResultOutput struct{ *pulumi.OutputState }

func (LookupIngressRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIngressRuleResult)(nil)).Elem()
}

func (o LookupIngressRuleResultOutput) ToLookupIngressRuleResultOutput() LookupIngressRuleResultOutput {
	return o
}

func (o LookupIngressRuleResultOutput) ToLookupIngressRuleResultOutputWithContext(ctx context.Context) LookupIngressRuleResultOutput {
	return o
}

// The action to take on matched requests.
func (o LookupIngressRuleResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIngressRuleResult) string { return v.Action }).(pulumi.StringOutput)
}

// An optional string description of this rule. This field has a maximum length of 400 characters.
func (o LookupIngressRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIngressRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

// A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
func (o LookupIngressRuleResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIngressRuleResult) int { return v.Priority }).(pulumi.IntOutput)
}

// IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
func (o LookupIngressRuleResultOutput) SourceRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIngressRuleResult) string { return v.SourceRange }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIngressRuleResultOutput{})
}