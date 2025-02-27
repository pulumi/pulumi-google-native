// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a CustomJob. A created CustomJob right away will be attempted to be run.
// Auto-naming is currently not supported for this resource.
type CustomJob struct {
	pulumi.CustomResourceState

	// Time when the CustomJob was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The display name of the CustomJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Customer-managed encryption key options for a CustomJob. If this is set, then all resources created by the CustomJob will be encrypted with the provided encryption key.
	EncryptionSpec GoogleCloudAiplatformV1EncryptionSpecResponseOutput `pulumi:"encryptionSpec"`
	// Time when the CustomJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Only populated when job's state is `JOB_STATE_FAILED` or `JOB_STATE_CANCELLED`.
	Error GoogleRpcStatusResponseOutput `pulumi:"error"`
	// Job spec.
	JobSpec GoogleCloudAiplatformV1CustomJobSpecResponseOutput `pulumi:"jobSpec"`
	// The labels with user-defined metadata to organize CustomJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// Resource name of a CustomJob.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Time when the CustomJob for the first time entered the `JOB_STATE_RUNNING` state.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The detailed state of the job.
	State pulumi.StringOutput `pulumi:"state"`
	// Time when the CustomJob was most recently updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// URIs for accessing [interactive shells](https://cloud.google.com/vertex-ai/docs/training/monitor-debug-interactive-shell) (one URI for each training node). Only available if job_spec.enable_web_access is `true`. The keys are names of each node in the training job; for example, `workerpool0-0` for the primary node, `workerpool1-0` for the first node in the second worker pool, and `workerpool1-1` for the second node in the second worker pool. The values are the URIs for each node's interactive shell.
	WebAccessUris pulumi.StringMapOutput `pulumi:"webAccessUris"`
}

// NewCustomJob registers a new resource with the given unique name, arguments, and options.
func NewCustomJob(ctx *pulumi.Context,
	name string, args *CustomJobArgs, opts ...pulumi.ResourceOption) (*CustomJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.JobSpec == nil {
		return nil, errors.New("invalid value for required argument 'JobSpec'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CustomJob
	err := ctx.RegisterResource("google-native:aiplatform/v1:CustomJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomJob gets an existing CustomJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomJobState, opts ...pulumi.ResourceOption) (*CustomJob, error) {
	var resource CustomJob
	err := ctx.ReadResource("google-native:aiplatform/v1:CustomJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomJob resources.
type customJobState struct {
}

type CustomJobState struct {
}

func (CustomJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*customJobState)(nil)).Elem()
}

type customJobArgs struct {
	// The display name of the CustomJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key options for a CustomJob. If this is set, then all resources created by the CustomJob will be encrypted with the provided encryption key.
	EncryptionSpec *GoogleCloudAiplatformV1EncryptionSpec `pulumi:"encryptionSpec"`
	// Job spec.
	JobSpec GoogleCloudAiplatformV1CustomJobSpec `pulumi:"jobSpec"`
	// The labels with user-defined metadata to organize CustomJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
}

// The set of arguments for constructing a CustomJob resource.
type CustomJobArgs struct {
	// The display name of the CustomJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName pulumi.StringInput
	// Customer-managed encryption key options for a CustomJob. If this is set, then all resources created by the CustomJob will be encrypted with the provided encryption key.
	EncryptionSpec GoogleCloudAiplatformV1EncryptionSpecPtrInput
	// Job spec.
	JobSpec GoogleCloudAiplatformV1CustomJobSpecInput
	// The labels with user-defined metadata to organize CustomJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
}

func (CustomJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customJobArgs)(nil)).Elem()
}

type CustomJobInput interface {
	pulumi.Input

	ToCustomJobOutput() CustomJobOutput
	ToCustomJobOutputWithContext(ctx context.Context) CustomJobOutput
}

func (*CustomJob) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomJob)(nil)).Elem()
}

func (i *CustomJob) ToCustomJobOutput() CustomJobOutput {
	return i.ToCustomJobOutputWithContext(context.Background())
}

func (i *CustomJob) ToCustomJobOutputWithContext(ctx context.Context) CustomJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomJobOutput)
}

type CustomJobOutput struct{ *pulumi.OutputState }

func (CustomJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomJob)(nil)).Elem()
}

func (o CustomJobOutput) ToCustomJobOutput() CustomJobOutput {
	return o
}

func (o CustomJobOutput) ToCustomJobOutputWithContext(ctx context.Context) CustomJobOutput {
	return o
}

// Time when the CustomJob was created.
func (o CustomJobOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The display name of the CustomJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
func (o CustomJobOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Customer-managed encryption key options for a CustomJob. If this is set, then all resources created by the CustomJob will be encrypted with the provided encryption key.
func (o CustomJobOutput) EncryptionSpec() GoogleCloudAiplatformV1EncryptionSpecResponseOutput {
	return o.ApplyT(func(v *CustomJob) GoogleCloudAiplatformV1EncryptionSpecResponseOutput { return v.EncryptionSpec }).(GoogleCloudAiplatformV1EncryptionSpecResponseOutput)
}

// Time when the CustomJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
func (o CustomJobOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.EndTime }).(pulumi.StringOutput)
}

// Only populated when job's state is `JOB_STATE_FAILED` or `JOB_STATE_CANCELLED`.
func (o CustomJobOutput) Error() GoogleRpcStatusResponseOutput {
	return o.ApplyT(func(v *CustomJob) GoogleRpcStatusResponseOutput { return v.Error }).(GoogleRpcStatusResponseOutput)
}

// Job spec.
func (o CustomJobOutput) JobSpec() GoogleCloudAiplatformV1CustomJobSpecResponseOutput {
	return o.ApplyT(func(v *CustomJob) GoogleCloudAiplatformV1CustomJobSpecResponseOutput { return v.JobSpec }).(GoogleCloudAiplatformV1CustomJobSpecResponseOutput)
}

// The labels with user-defined metadata to organize CustomJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
func (o CustomJobOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o CustomJobOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Resource name of a CustomJob.
func (o CustomJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CustomJobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Time when the CustomJob for the first time entered the `JOB_STATE_RUNNING` state.
func (o CustomJobOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.StartTime }).(pulumi.StringOutput)
}

// The detailed state of the job.
func (o CustomJobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Time when the CustomJob was most recently updated.
func (o CustomJobOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// URIs for accessing [interactive shells](https://cloud.google.com/vertex-ai/docs/training/monitor-debug-interactive-shell) (one URI for each training node). Only available if job_spec.enable_web_access is `true`. The keys are names of each node in the training job; for example, `workerpool0-0` for the primary node, `workerpool1-0` for the first node in the second worker pool, and `workerpool1-1` for the second node in the second worker pool. The values are the URIs for each node's interactive shell.
func (o CustomJobOutput) WebAccessUris() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomJob) pulumi.StringMapOutput { return v.WebAccessUris }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomJobInput)(nil)).Elem(), &CustomJob{})
	pulumi.RegisterOutputType(CustomJobOutput{})
}
