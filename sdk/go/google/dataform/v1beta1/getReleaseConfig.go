// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches a single ReleaseConfig.
func LookupReleaseConfig(ctx *pulumi.Context, args *LookupReleaseConfigArgs, opts ...pulumi.InvokeOption) (*LookupReleaseConfigResult, error) {
	var rv LookupReleaseConfigResult
	err := ctx.Invoke("google-native:dataform/v1beta1:getReleaseConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReleaseConfigArgs struct {
	Location        string  `pulumi:"location"`
	Project         *string `pulumi:"project"`
	ReleaseConfigId string  `pulumi:"releaseConfigId"`
	RepositoryId    string  `pulumi:"repositoryId"`
}

type LookupReleaseConfigResult struct {
	// Optional. If set, fields of `code_compilation_config` override the default compilation settings that are specified in dataform.json.
	CodeCompilationConfig CodeCompilationConfigResponse `pulumi:"codeCompilationConfig"`
	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule string `pulumi:"cronSchedule"`
	// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository. Examples: - a commit SHA: `12ade345` - a tag: `tag1` - a branch name: `branch1`
	GitCommitish string `pulumi:"gitCommitish"`
	// The release config's name.
	Name string `pulumi:"name"`
	// Records of the 10 most recent scheduled release attempts. Updated whenever automatic creation of a compilation result is triggered by cron_schedule.
	RecentScheduledReleaseRecords []ScheduledReleaseRecordResponse `pulumi:"recentScheduledReleaseRecords"`
	// Optional. The name of the currently released compilation result for this release config. This value is updated when a compilation result is created from this release config, or when this resource is updated by API call (perhaps to roll back to an earlier release). The compilation result must have been created using this release config. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
	ReleaseCompilationResult string `pulumi:"releaseCompilationResult"`
	// Optional. Specifies the time zone to be used when interpreting cron_schedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone string `pulumi:"timeZone"`
}

func LookupReleaseConfigOutput(ctx *pulumi.Context, args LookupReleaseConfigOutputArgs, opts ...pulumi.InvokeOption) LookupReleaseConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReleaseConfigResult, error) {
			args := v.(LookupReleaseConfigArgs)
			r, err := LookupReleaseConfig(ctx, &args, opts...)
			var s LookupReleaseConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReleaseConfigResultOutput)
}

type LookupReleaseConfigOutputArgs struct {
	Location        pulumi.StringInput    `pulumi:"location"`
	Project         pulumi.StringPtrInput `pulumi:"project"`
	ReleaseConfigId pulumi.StringInput    `pulumi:"releaseConfigId"`
	RepositoryId    pulumi.StringInput    `pulumi:"repositoryId"`
}

func (LookupReleaseConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseConfigArgs)(nil)).Elem()
}

type LookupReleaseConfigResultOutput struct{ *pulumi.OutputState }

func (LookupReleaseConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReleaseConfigResult)(nil)).Elem()
}

func (o LookupReleaseConfigResultOutput) ToLookupReleaseConfigResultOutput() LookupReleaseConfigResultOutput {
	return o
}

func (o LookupReleaseConfigResultOutput) ToLookupReleaseConfigResultOutputWithContext(ctx context.Context) LookupReleaseConfigResultOutput {
	return o
}

// Optional. If set, fields of `code_compilation_config` override the default compilation settings that are specified in dataform.json.
func (o LookupReleaseConfigResultOutput) CodeCompilationConfig() CodeCompilationConfigResponseOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) CodeCompilationConfigResponse { return v.CodeCompilationConfig }).(CodeCompilationConfigResponseOutput)
}

// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
func (o LookupReleaseConfigResultOutput) CronSchedule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) string { return v.CronSchedule }).(pulumi.StringOutput)
}

// Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository. Examples: - a commit SHA: `12ade345` - a tag: `tag1` - a branch name: `branch1`
func (o LookupReleaseConfigResultOutput) GitCommitish() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) string { return v.GitCommitish }).(pulumi.StringOutput)
}

// The release config's name.
func (o LookupReleaseConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// Records of the 10 most recent scheduled release attempts. Updated whenever automatic creation of a compilation result is triggered by cron_schedule.
func (o LookupReleaseConfigResultOutput) RecentScheduledReleaseRecords() ScheduledReleaseRecordResponseArrayOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) []ScheduledReleaseRecordResponse {
		return v.RecentScheduledReleaseRecords
	}).(ScheduledReleaseRecordResponseArrayOutput)
}

// Optional. The name of the currently released compilation result for this release config. This value is updated when a compilation result is created from this release config, or when this resource is updated by API call (perhaps to roll back to an earlier release). The compilation result must have been created using this release config. Must be in the format `projects/*/locations/*/repositories/*/compilationResults/*`.
func (o LookupReleaseConfigResultOutput) ReleaseCompilationResult() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) string { return v.ReleaseCompilationResult }).(pulumi.StringOutput)
}

// Optional. Specifies the time zone to be used when interpreting cron_schedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
func (o LookupReleaseConfigResultOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupReleaseConfigResult) string { return v.TimeZone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReleaseConfigResultOutput{})
}