// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a authorized orgs desc based on the resource name.
func LookupAuthorizedOrgsDesc(ctx *pulumi.Context, args *LookupAuthorizedOrgsDescArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizedOrgsDescResult, error) {
	var rv LookupAuthorizedOrgsDescResult
	err := ctx.Invoke("google-native:accesscontextmanager/v1:getAuthorizedOrgsDesc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAuthorizedOrgsDescArgs struct {
	AccessPolicyId       string `pulumi:"accessPolicyId"`
	AuthorizedOrgsDescId string `pulumi:"authorizedOrgsDescId"`
}

type LookupAuthorizedOrgsDescResult struct {
	// The asset type of this authorized orgs desc. e.g. device, credential strength.
	AssetType string `pulumi:"assetType"`
	// Authorization direction of this authorization relationship. i.e. Whether to allow specified orgs to evaluate this org's traffic, or allow specified orgs' traffic to be evaluated by this org. Orgs specified as `AUTHORIZATION_DIRECTION_TO` in this AuthorizedOrgsDesc[com.google.identity.accesscontextmanager.v1.AuthorizedOrgsDesc] must also specify this org as the `AUTHORIZATION_DIRECTION_FROM` in their own AuthorizedOrgsDesc in order for this relationship to take effect. Orgs specified as `AUTHORIZATION_DIRECTION_FROM` in this AuthorizedOrgsDesc[com.google.identity.accesscontextmanager.v1.AuthorizedOrgsDesc] must also specify this org as the `AUTHORIZATION_DIRECTION_TO` in their own AuthorizedOrgsDesc in order for this relationship to take effect.
	AuthorizationDirection string `pulumi:"authorizationDirection"`
	// The authorization type of this authorized orgs desc. e.g.authorization, troubleshooting or logging.
	AuthorizationType string `pulumi:"authorizationType"`
	// Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "accessPolicies/122256/authorizedOrgs/b3-BhcX_Ud5N"
	Name string `pulumi:"name"`
	// The list of organization ids in this AuthorizedOrgsDesc.
	Orgs []string `pulumi:"orgs"`
}

func LookupAuthorizedOrgsDescOutput(ctx *pulumi.Context, args LookupAuthorizedOrgsDescOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizedOrgsDescResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizedOrgsDescResult, error) {
			args := v.(LookupAuthorizedOrgsDescArgs)
			r, err := LookupAuthorizedOrgsDesc(ctx, &args, opts...)
			var s LookupAuthorizedOrgsDescResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizedOrgsDescResultOutput)
}

type LookupAuthorizedOrgsDescOutputArgs struct {
	AccessPolicyId       pulumi.StringInput `pulumi:"accessPolicyId"`
	AuthorizedOrgsDescId pulumi.StringInput `pulumi:"authorizedOrgsDescId"`
}

func (LookupAuthorizedOrgsDescOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizedOrgsDescArgs)(nil)).Elem()
}

type LookupAuthorizedOrgsDescResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizedOrgsDescResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizedOrgsDescResult)(nil)).Elem()
}

func (o LookupAuthorizedOrgsDescResultOutput) ToLookupAuthorizedOrgsDescResultOutput() LookupAuthorizedOrgsDescResultOutput {
	return o
}

func (o LookupAuthorizedOrgsDescResultOutput) ToLookupAuthorizedOrgsDescResultOutputWithContext(ctx context.Context) LookupAuthorizedOrgsDescResultOutput {
	return o
}

// The asset type of this authorized orgs desc. e.g. device, credential strength.
func (o LookupAuthorizedOrgsDescResultOutput) AssetType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.AssetType }).(pulumi.StringOutput)
}

// Authorization direction of this authorization relationship. i.e. Whether to allow specified orgs to evaluate this org's traffic, or allow specified orgs' traffic to be evaluated by this org. Orgs specified as `AUTHORIZATION_DIRECTION_TO` in this AuthorizedOrgsDesc[com.google.identity.accesscontextmanager.v1.AuthorizedOrgsDesc] must also specify this org as the `AUTHORIZATION_DIRECTION_FROM` in their own AuthorizedOrgsDesc in order for this relationship to take effect. Orgs specified as `AUTHORIZATION_DIRECTION_FROM` in this AuthorizedOrgsDesc[com.google.identity.accesscontextmanager.v1.AuthorizedOrgsDesc] must also specify this org as the `AUTHORIZATION_DIRECTION_TO` in their own AuthorizedOrgsDesc in order for this relationship to take effect.
func (o LookupAuthorizedOrgsDescResultOutput) AuthorizationDirection() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.AuthorizationDirection }).(pulumi.StringOutput)
}

// The authorization type of this authorized orgs desc. e.g.authorization, troubleshooting or logging.
func (o LookupAuthorizedOrgsDescResultOutput) AuthorizationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.AuthorizationType }).(pulumi.StringOutput)
}

// Assigned by the server during creation. The last segment has an arbitrary length and has only URI unreserved characters (as defined by [RFC 3986 Section 2.3](https://tools.ietf.org/html/rfc3986#section-2.3)). Should not be specified by the client during creation. Example: "accessPolicies/122256/authorizedOrgs/b3-BhcX_Ud5N"
func (o LookupAuthorizedOrgsDescResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of organization ids in this AuthorizedOrgsDesc.
func (o LookupAuthorizedOrgsDescResultOutput) Orgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizedOrgsDescResult) []string { return v.Orgs }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizedOrgsDescResultOutput{})
}