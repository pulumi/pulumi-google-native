// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches the representation of an existing Response Policy Rule.
func LookupResponsePolicyRule(ctx *pulumi.Context, args *LookupResponsePolicyRuleArgs, opts ...pulumi.InvokeOption) (*LookupResponsePolicyRuleResult, error) {
	var rv LookupResponsePolicyRuleResult
	err := ctx.Invoke("google-native:dns/v1beta2:getResponsePolicyRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResponsePolicyRuleArgs struct {
	ClientOperationId  *string `pulumi:"clientOperationId"`
	Project            *string `pulumi:"project"`
	ResponsePolicy     string  `pulumi:"responsePolicy"`
	ResponsePolicyRule string  `pulumi:"responsePolicyRule"`
}

type LookupResponsePolicyRuleResult struct {
	// Answer this query with a behavior rather than DNS data.
	Behavior string `pulumi:"behavior"`
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	DnsName string `pulumi:"dnsName"`
	Kind    string `pulumi:"kind"`
	// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
	LocalData ResponsePolicyRuleLocalDataResponse `pulumi:"localData"`
	// An identifier for this rule. Must be unique with the ResponsePolicy.
	RuleName string `pulumi:"ruleName"`
}

func LookupResponsePolicyRuleOutput(ctx *pulumi.Context, args LookupResponsePolicyRuleOutputArgs, opts ...pulumi.InvokeOption) LookupResponsePolicyRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResponsePolicyRuleResult, error) {
			args := v.(LookupResponsePolicyRuleArgs)
			r, err := LookupResponsePolicyRule(ctx, &args, opts...)
			return *r, err
		}).(LookupResponsePolicyRuleResultOutput)
}

type LookupResponsePolicyRuleOutputArgs struct {
	ClientOperationId  pulumi.StringPtrInput `pulumi:"clientOperationId"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
	ResponsePolicy     pulumi.StringInput    `pulumi:"responsePolicy"`
	ResponsePolicyRule pulumi.StringInput    `pulumi:"responsePolicyRule"`
}

func (LookupResponsePolicyRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponsePolicyRuleArgs)(nil)).Elem()
}

type LookupResponsePolicyRuleResultOutput struct{ *pulumi.OutputState }

func (LookupResponsePolicyRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponsePolicyRuleResult)(nil)).Elem()
}

func (o LookupResponsePolicyRuleResultOutput) ToLookupResponsePolicyRuleResultOutput() LookupResponsePolicyRuleResultOutput {
	return o
}

func (o LookupResponsePolicyRuleResultOutput) ToLookupResponsePolicyRuleResultOutputWithContext(ctx context.Context) LookupResponsePolicyRuleResultOutput {
	return o
}

// Answer this query with a behavior rather than DNS data.
func (o LookupResponsePolicyRuleResultOutput) Behavior() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponsePolicyRuleResult) string { return v.Behavior }).(pulumi.StringOutput)
}

// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
func (o LookupResponsePolicyRuleResultOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponsePolicyRuleResult) string { return v.DnsName }).(pulumi.StringOutput)
}

func (o LookupResponsePolicyRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponsePolicyRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Answer this query directly with DNS data. These ResourceRecordSets override any other DNS behavior for the matched name; in particular they override private zones, the public internet, and GCP internal DNS. No SOA nor NS types are allowed.
func (o LookupResponsePolicyRuleResultOutput) LocalData() ResponsePolicyRuleLocalDataResponseOutput {
	return o.ApplyT(func(v LookupResponsePolicyRuleResult) ResponsePolicyRuleLocalDataResponse { return v.LocalData }).(ResponsePolicyRuleLocalDataResponseOutput)
}

// An identifier for this rule. Must be unique with the ResponsePolicy.
func (o LookupResponsePolicyRuleResultOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResponsePolicyRuleResult) string { return v.RuleName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResponsePolicyRuleResultOutput{})
}