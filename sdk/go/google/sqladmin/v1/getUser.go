// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a resource containing information about a user.
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	var rv LookupUserResult
	err := ctx.Invoke("google-native:sqladmin/v1:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserArgs struct {
	Host     *string `pulumi:"host"`
	Instance string  `pulumi:"instance"`
	Name     string  `pulumi:"name"`
	Project  *string `pulumi:"project"`
}

type LookupUserResult struct {
	// Dual password status for the user.
	DualPasswordType string `pulumi:"dualPasswordType"`
	// This field is deprecated and will be removed from a future version of the API.
	//
	// Deprecated: This field is deprecated and will be removed from a future version of the API.
	Etag string `pulumi:"etag"`
	// Optional. The host from which the user can connect. For `insert` operations, host defaults to an empty string. For `update` operations, host is specified as part of the request URL. The host name cannot be updated after insertion. For a MySQL instance, it's required; for a PostgreSQL or SQL Server instance, it's optional.
	Host string `pulumi:"host"`
	// The name of the Cloud SQL instance. This does not include the project ID. Can be omitted for `update` because it is already specified on the URL.
	Instance string `pulumi:"instance"`
	// This is always `sql#user`.
	Kind string `pulumi:"kind"`
	// The name of the user in the Cloud SQL instance. Can be omitted for `update` because it is already specified in the URL.
	Name string `pulumi:"name"`
	// The password for the user.
	Password string `pulumi:"password"`
	// User level password validation policy.
	PasswordPolicy UserPasswordValidationPolicyResponse `pulumi:"passwordPolicy"`
	// The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable. Can be omitted for `update` because it is already specified on the URL.
	Project              string                       `pulumi:"project"`
	SqlserverUserDetails SqlServerUserDetailsResponse `pulumi:"sqlserverUserDetails"`
	// The user type. It determines the method to authenticate the user during login. The default is the database's built-in user type.
	Type string `pulumi:"type"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

type LookupUserOutputArgs struct {
	Host     pulumi.StringPtrInput `pulumi:"host"`
	Instance pulumi.StringInput    `pulumi:"instance"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// Dual password status for the user.
func (o LookupUserResultOutput) DualPasswordType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.DualPasswordType }).(pulumi.StringOutput)
}

// This field is deprecated and will be removed from a future version of the API.
//
// Deprecated: This field is deprecated and will be removed from a future version of the API.
func (o LookupUserResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Optional. The host from which the user can connect. For `insert` operations, host defaults to an empty string. For `update` operations, host is specified as part of the request URL. The host name cannot be updated after insertion. For a MySQL instance, it's required; for a PostgreSQL or SQL Server instance, it's optional.
func (o LookupUserResultOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Host }).(pulumi.StringOutput)
}

// The name of the Cloud SQL instance. This does not include the project ID. Can be omitted for `update` because it is already specified on the URL.
func (o LookupUserResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Instance }).(pulumi.StringOutput)
}

// This is always `sql#user`.
func (o LookupUserResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the user in the Cloud SQL instance. Can be omitted for `update` because it is already specified in the URL.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// The password for the user.
func (o LookupUserResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Password }).(pulumi.StringOutput)
}

// User level password validation policy.
func (o LookupUserResultOutput) PasswordPolicy() UserPasswordValidationPolicyResponseOutput {
	return o.ApplyT(func(v LookupUserResult) UserPasswordValidationPolicyResponse { return v.PasswordPolicy }).(UserPasswordValidationPolicyResponseOutput)
}

// The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable. Can be omitted for `update` because it is already specified on the URL.
func (o LookupUserResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupUserResultOutput) SqlserverUserDetails() SqlServerUserDetailsResponseOutput {
	return o.ApplyT(func(v LookupUserResult) SqlServerUserDetailsResponse { return v.SqlserverUserDetails }).(SqlServerUserDetailsResponseOutput)
}

// The user type. It determines the method to authenticate the user during login. The default is the database's built-in user type.
func (o LookupUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}