// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a NasJob
func LookupNasJob(ctx *pulumi.Context, args *LookupNasJobArgs, opts ...pulumi.InvokeOption) (*LookupNasJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNasJobResult
	err := ctx.Invoke("google-native:aiplatform/v1:getNasJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNasJobArgs struct {
	Location string  `pulumi:"location"`
	NasJobId string  `pulumi:"nasJobId"`
	Project  *string `pulumi:"project"`
}

type LookupNasJobResult struct {
	// Time when the NasJob was created.
	CreateTime string `pulumi:"createTime"`
	// The display name of the NasJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `pulumi:"displayName"`
	// Optional. Enable a separation of Custom model training and restricted image training for tenant project.
	EnableRestrictedImageTraining bool `pulumi:"enableRestrictedImageTraining"`
	// Customer-managed encryption key options for a NasJob. If this is set, then all resources created by the NasJob will be encrypted with the provided encryption key.
	EncryptionSpec GoogleCloudAiplatformV1EncryptionSpecResponse `pulumi:"encryptionSpec"`
	// Time when the NasJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
	EndTime string `pulumi:"endTime"`
	// Only populated when job's state is JOB_STATE_FAILED or JOB_STATE_CANCELLED.
	Error GoogleRpcStatusResponse `pulumi:"error"`
	// The labels with user-defined metadata to organize NasJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the NasJob.
	Name string `pulumi:"name"`
	// Output of the NasJob.
	NasJobOutput GoogleCloudAiplatformV1NasJobOutputResponse `pulumi:"nasJobOutput"`
	// The specification of a NasJob.
	NasJobSpec GoogleCloudAiplatformV1NasJobSpecResponse `pulumi:"nasJobSpec"`
	// Time when the NasJob for the first time entered the `JOB_STATE_RUNNING` state.
	StartTime string `pulumi:"startTime"`
	// The detailed state of the job.
	State string `pulumi:"state"`
	// Time when the NasJob was most recently updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupNasJobOutput(ctx *pulumi.Context, args LookupNasJobOutputArgs, opts ...pulumi.InvokeOption) LookupNasJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNasJobResult, error) {
			args := v.(LookupNasJobArgs)
			r, err := LookupNasJob(ctx, &args, opts...)
			var s LookupNasJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNasJobResultOutput)
}

type LookupNasJobOutputArgs struct {
	Location pulumi.StringInput    `pulumi:"location"`
	NasJobId pulumi.StringInput    `pulumi:"nasJobId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupNasJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNasJobArgs)(nil)).Elem()
}

type LookupNasJobResultOutput struct{ *pulumi.OutputState }

func (LookupNasJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNasJobResult)(nil)).Elem()
}

func (o LookupNasJobResultOutput) ToLookupNasJobResultOutput() LookupNasJobResultOutput {
	return o
}

func (o LookupNasJobResultOutput) ToLookupNasJobResultOutputWithContext(ctx context.Context) LookupNasJobResultOutput {
	return o
}

// Time when the NasJob was created.
func (o LookupNasJobResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The display name of the NasJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
func (o LookupNasJobResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Enable a separation of Custom model training and restricted image training for tenant project.
func (o LookupNasJobResultOutput) EnableRestrictedImageTraining() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNasJobResult) bool { return v.EnableRestrictedImageTraining }).(pulumi.BoolOutput)
}

// Customer-managed encryption key options for a NasJob. If this is set, then all resources created by the NasJob will be encrypted with the provided encryption key.
func (o LookupNasJobResultOutput) EncryptionSpec() GoogleCloudAiplatformV1EncryptionSpecResponseOutput {
	return o.ApplyT(func(v LookupNasJobResult) GoogleCloudAiplatformV1EncryptionSpecResponse { return v.EncryptionSpec }).(GoogleCloudAiplatformV1EncryptionSpecResponseOutput)
}

// Time when the NasJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
func (o LookupNasJobResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Only populated when job's state is JOB_STATE_FAILED or JOB_STATE_CANCELLED.
func (o LookupNasJobResultOutput) Error() GoogleRpcStatusResponseOutput {
	return o.ApplyT(func(v LookupNasJobResult) GoogleRpcStatusResponse { return v.Error }).(GoogleRpcStatusResponseOutput)
}

// The labels with user-defined metadata to organize NasJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
func (o LookupNasJobResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNasJobResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource name of the NasJob.
func (o LookupNasJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.Name }).(pulumi.StringOutput)
}

// Output of the NasJob.
func (o LookupNasJobResultOutput) NasJobOutput() GoogleCloudAiplatformV1NasJobOutputResponseOutput {
	return o.ApplyT(func(v LookupNasJobResult) GoogleCloudAiplatformV1NasJobOutputResponse { return v.NasJobOutput }).(GoogleCloudAiplatformV1NasJobOutputResponseOutput)
}

// The specification of a NasJob.
func (o LookupNasJobResultOutput) NasJobSpec() GoogleCloudAiplatformV1NasJobSpecResponseOutput {
	return o.ApplyT(func(v LookupNasJobResult) GoogleCloudAiplatformV1NasJobSpecResponse { return v.NasJobSpec }).(GoogleCloudAiplatformV1NasJobSpecResponseOutput)
}

// Time when the NasJob for the first time entered the `JOB_STATE_RUNNING` state.
func (o LookupNasJobResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The detailed state of the job.
func (o LookupNasJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.State }).(pulumi.StringOutput)
}

// Time when the NasJob was most recently updated.
func (o LookupNasJobResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasJobResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNasJobResultOutput{})
}