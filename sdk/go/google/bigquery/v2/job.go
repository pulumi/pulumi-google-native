// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Starts a new asynchronous job. Requires the Can View project role.
// Auto-naming is currently not supported for this resource.
type Job struct {
	pulumi.CustomResourceState

	// [Required] Describes the job configuration.
	Configuration JobConfigurationResponseOutput `pulumi:"configuration"`
	// A hash of this resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// [Optional] Reference describing the unique-per-user name of the job.
	JobReference JobReferenceResponseOutput `pulumi:"jobReference"`
	// The type of the resource.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A URL that can be used to access this resource again.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Information about the job, including starting time and ending time of the job.
	Statistics JobStatisticsResponseOutput `pulumi:"statistics"`
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	Status JobStatusResponseOutput `pulumi:"status"`
	// Email address of the user who ran the job.
	UserEmail pulumi.StringOutput `pulumi:"userEmail"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		args = &JobArgs{}
	}

	var resource Job
	err := ctx.RegisterResource("google-native:bigquery/v2:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("google-native:bigquery/v2:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// [Required] Describes the job configuration.
	Configuration *JobConfiguration `pulumi:"configuration"`
	// [Optional] Reference describing the unique-per-user name of the job.
	JobReference *JobReference         `pulumi:"jobReference"`
	Project      *string               `pulumi:"project"`
	Source       pulumi.AssetOrArchive `pulumi:"source"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// [Required] Describes the job configuration.
	Configuration JobConfigurationPtrInput
	// [Optional] Reference describing the unique-per-user name of the job.
	JobReference JobReferencePtrInput
	Project      pulumi.StringPtrInput
	Source       pulumi.AssetOrArchiveInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterOutputType(JobOutput{})
}