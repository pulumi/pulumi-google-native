// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta4

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a resource containing information about a database inside a Cloud SQL instance.
func LookupDatabase(ctx *pulumi.Context, args *LookupDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupDatabaseResult, error) {
	var rv LookupDatabaseResult
	err := ctx.Invoke("google-native:sqladmin/v1beta4:getDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDatabaseArgs struct {
	Database string  `pulumi:"database"`
	Instance string  `pulumi:"instance"`
	Project  *string `pulumi:"project"`
}

type LookupDatabaseResult struct {
	// The Cloud SQL charset value.
	Charset string `pulumi:"charset"`
	// The Cloud SQL collation value.
	Collation string `pulumi:"collation"`
	// The name of the Cloud SQL instance. This does not include the project ID.
	Instance string `pulumi:"instance"`
	// This is always `sql#database`.
	Kind string `pulumi:"kind"`
	// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
	Name string `pulumi:"name"`
	// The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
	Project string `pulumi:"project"`
	// The URI of this resource.
	SelfLink                 string                           `pulumi:"selfLink"`
	SqlserverDatabaseDetails SqlServerDatabaseDetailsResponse `pulumi:"sqlserverDatabaseDetails"`
}

func LookupDatabaseOutput(ctx *pulumi.Context, args LookupDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatabaseResult, error) {
			args := v.(LookupDatabaseArgs)
			r, err := LookupDatabase(ctx, &args, opts...)
			return *r, err
		}).(LookupDatabaseResultOutput)
}

type LookupDatabaseOutputArgs struct {
	Database pulumi.StringInput    `pulumi:"database"`
	Instance pulumi.StringInput    `pulumi:"instance"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseArgs)(nil)).Elem()
}

type LookupDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatabaseResult)(nil)).Elem()
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutput() LookupDatabaseResultOutput {
	return o
}

func (o LookupDatabaseResultOutput) ToLookupDatabaseResultOutputWithContext(ctx context.Context) LookupDatabaseResultOutput {
	return o
}

// The Cloud SQL charset value.
func (o LookupDatabaseResultOutput) Charset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Charset }).(pulumi.StringOutput)
}

// The Cloud SQL collation value.
func (o LookupDatabaseResultOutput) Collation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Collation }).(pulumi.StringOutput)
}

// The name of the Cloud SQL instance. This does not include the project ID.
func (o LookupDatabaseResultOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Instance }).(pulumi.StringOutput)
}

// This is always `sql#database`.
func (o LookupDatabaseResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Kind }).(pulumi.StringOutput)
}

// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
func (o LookupDatabaseResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Name }).(pulumi.StringOutput)
}

// The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
func (o LookupDatabaseResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.Project }).(pulumi.StringOutput)
}

// The URI of this resource.
func (o LookupDatabaseResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatabaseResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupDatabaseResultOutput) SqlserverDatabaseDetails() SqlServerDatabaseDetailsResponseOutput {
	return o.ApplyT(func(v LookupDatabaseResult) SqlServerDatabaseDetailsResponse { return v.SqlserverDatabaseDetails }).(SqlServerDatabaseDetailsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatabaseResultOutput{})
}