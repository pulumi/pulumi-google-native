// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Returns information about a specific job. Job information is available for a six month period after creation. Requires that you're the person who ran the job, or have the Is Owner project role.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("google-native:bigquery/v2:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	JobId    string  `pulumi:"jobId"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupJobResult struct {
	// [Required] Describes the job configuration.
	Configuration JobConfigurationResponse `pulumi:"configuration"`
	// A hash of this resource.
	Etag string `pulumi:"etag"`
	// [Optional] Reference describing the unique-per-user name of the job.
	JobReference JobReferenceResponse `pulumi:"jobReference"`
	// The type of the resource.
	Kind string `pulumi:"kind"`
	// A URL that can be used to access this resource again.
	SelfLink string `pulumi:"selfLink"`
	// Information about the job, including starting time and ending time of the job.
	Statistics JobStatisticsResponse `pulumi:"statistics"`
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	Status JobStatusResponse `pulumi:"status"`
	// Email address of the user who ran the job.
	UserEmail string `pulumi:"userEmail"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	JobId    pulumi.StringInput    `pulumi:"jobId"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

// [Required] Describes the job configuration.
func (o LookupJobResultOutput) Configuration() JobConfigurationResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobConfigurationResponse { return v.Configuration }).(JobConfigurationResponseOutput)
}

// A hash of this resource.
func (o LookupJobResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Etag }).(pulumi.StringOutput)
}

// [Optional] Reference describing the unique-per-user name of the job.
func (o LookupJobResultOutput) JobReference() JobReferenceResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobReferenceResponse { return v.JobReference }).(JobReferenceResponseOutput)
}

// The type of the resource.
func (o LookupJobResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A URL that can be used to access this resource again.
func (o LookupJobResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Information about the job, including starting time and ending time of the job.
func (o LookupJobResultOutput) Statistics() JobStatisticsResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobStatisticsResponse { return v.Statistics }).(JobStatisticsResponseOutput)
}

// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
func (o LookupJobResultOutput) Status() JobStatusResponseOutput {
	return o.ApplyT(func(v LookupJobResult) JobStatusResponse { return v.Status }).(JobStatusResponseOutput)
}

// Email address of the user who ran the job.
func (o LookupJobResultOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobResult) string { return v.UserEmail }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}