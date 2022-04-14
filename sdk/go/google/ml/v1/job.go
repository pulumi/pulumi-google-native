// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a training or a batch prediction job.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Job struct {
	pulumi.CustomResourceState

	// When the job was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// When the job processing was completed.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The details of a failure or a cancellation.
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The user-specified id of the job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// It's only effect when the job is in QUEUED state. If it's positive, it indicates the job's position in the job scheduler. It's 0 when the job is already scheduled.
	JobPosition pulumi.StringOutput `pulumi:"jobPosition"`
	// Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Input parameters to create a prediction job.
	PredictionInput GoogleCloudMlV1__PredictionInputResponseOutput `pulumi:"predictionInput"`
	// The current prediction job result.
	PredictionOutput GoogleCloudMlV1__PredictionOutputResponseOutput `pulumi:"predictionOutput"`
	// When the job processing was started.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The detailed state of a job.
	State pulumi.StringOutput `pulumi:"state"`
	// Input parameters to create a training job.
	TrainingInput GoogleCloudMlV1__TrainingInputResponseOutput `pulumi:"trainingInput"`
	// The current training job result.
	TrainingOutput GoogleCloudMlV1__TrainingOutputResponseOutput `pulumi:"trainingOutput"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	var resource Job
	err := ctx.RegisterResource("google-native:ml/v1:Job", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:ml/v1:Job", name, id, state, &resource, opts...)
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
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
	Etag *string `pulumi:"etag"`
	// The user-specified id of the job.
	JobId string `pulumi:"jobId"`
	// Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
	Labels map[string]string `pulumi:"labels"`
	// Input parameters to create a prediction job.
	PredictionInput *GoogleCloudMlV1__PredictionInput `pulumi:"predictionInput"`
	// The current prediction job result.
	PredictionOutput *GoogleCloudMlV1__PredictionOutput `pulumi:"predictionOutput"`
	Project          *string                            `pulumi:"project"`
	// Input parameters to create a training job.
	TrainingInput *GoogleCloudMlV1__TrainingInput `pulumi:"trainingInput"`
	// The current training job result.
	TrainingOutput *GoogleCloudMlV1__TrainingOutput `pulumi:"trainingOutput"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a job from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform job updates in order to avoid race conditions: An `etag` is returned in the response to `GetJob`, and systems are expected to put that etag in the request to `UpdateJob` to ensure that their change will be applied to the same version of the job.
	Etag pulumi.StringPtrInput
	// The user-specified id of the job.
	JobId pulumi.StringInput
	// Optional. One or more labels that you can add, to organize your jobs. Each label is a key-value pair, where both the key and the value are arbitrary strings that you supply. For more information, see the documentation on using labels.
	Labels pulumi.StringMapInput
	// Input parameters to create a prediction job.
	PredictionInput GoogleCloudMlV1__PredictionInputPtrInput
	// The current prediction job result.
	PredictionOutput GoogleCloudMlV1__PredictionOutputPtrInput
	Project          pulumi.StringPtrInput
	// Input parameters to create a training job.
	TrainingInput GoogleCloudMlV1__TrainingInputPtrInput
	// The current training job result.
	TrainingOutput GoogleCloudMlV1__TrainingOutputPtrInput
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