// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a job trigger. See https://cloud.google.com/dlp/docs/creating-job-triggers to learn more.
func LookupOrganizationJobTrigger(ctx *pulumi.Context, args *LookupOrganizationJobTriggerArgs, opts ...pulumi.InvokeOption) (*LookupOrganizationJobTriggerResult, error) {
	var rv LookupOrganizationJobTriggerResult
	err := ctx.Invoke("google-native:dlp/v2:getOrganizationJobTrigger", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOrganizationJobTriggerArgs struct {
	JobTriggerId   string `pulumi:"jobTriggerId"`
	Location       string `pulumi:"location"`
	OrganizationId string `pulumi:"organizationId"`
}

type LookupOrganizationJobTriggerResult struct {
	// The creation timestamp of a triggeredJob.
	CreateTime string `pulumi:"createTime"`
	// User provided description (max 256 chars)
	Description string `pulumi:"description"`
	// Display name (max 100 chars)
	DisplayName string `pulumi:"displayName"`
	// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
	Errors []GooglePrivacyDlpV2ErrorResponse `pulumi:"errors"`
	// For inspect jobs, a snapshot of the configuration.
	InspectJob GooglePrivacyDlpV2InspectJobConfigResponse `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime string `pulumi:"lastRunTime"`
	// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
	Name string `pulumi:"name"`
	// A status for this trigger.
	Status string `pulumi:"status"`
	// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
	Triggers []GooglePrivacyDlpV2TriggerResponse `pulumi:"triggers"`
	// The last update timestamp of a triggeredJob.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupOrganizationJobTriggerOutput(ctx *pulumi.Context, args LookupOrganizationJobTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupOrganizationJobTriggerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrganizationJobTriggerResult, error) {
			args := v.(LookupOrganizationJobTriggerArgs)
			r, err := LookupOrganizationJobTrigger(ctx, &args, opts...)
			var s LookupOrganizationJobTriggerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrganizationJobTriggerResultOutput)
}

type LookupOrganizationJobTriggerOutputArgs struct {
	JobTriggerId   pulumi.StringInput `pulumi:"jobTriggerId"`
	Location       pulumi.StringInput `pulumi:"location"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupOrganizationJobTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationJobTriggerArgs)(nil)).Elem()
}

type LookupOrganizationJobTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupOrganizationJobTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrganizationJobTriggerResult)(nil)).Elem()
}

func (o LookupOrganizationJobTriggerResultOutput) ToLookupOrganizationJobTriggerResultOutput() LookupOrganizationJobTriggerResultOutput {
	return o
}

func (o LookupOrganizationJobTriggerResultOutput) ToLookupOrganizationJobTriggerResultOutputWithContext(ctx context.Context) LookupOrganizationJobTriggerResultOutput {
	return o
}

// The creation timestamp of a triggeredJob.
func (o LookupOrganizationJobTriggerResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// User provided description (max 256 chars)
func (o LookupOrganizationJobTriggerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.Description }).(pulumi.StringOutput)
}

// Display name (max 100 chars)
func (o LookupOrganizationJobTriggerResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A stream of errors encountered when the trigger was activated. Repeated errors may result in the JobTrigger automatically being paused. Will return the last 100 errors. Whenever the JobTrigger is modified this list will be cleared.
func (o LookupOrganizationJobTriggerResultOutput) Errors() GooglePrivacyDlpV2ErrorResponseArrayOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) []GooglePrivacyDlpV2ErrorResponse { return v.Errors }).(GooglePrivacyDlpV2ErrorResponseArrayOutput)
}

// For inspect jobs, a snapshot of the configuration.
func (o LookupOrganizationJobTriggerResultOutput) InspectJob() GooglePrivacyDlpV2InspectJobConfigResponseOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) GooglePrivacyDlpV2InspectJobConfigResponse {
		return v.InspectJob
	}).(GooglePrivacyDlpV2InspectJobConfigResponseOutput)
}

// The timestamp of the last time this trigger executed.
func (o LookupOrganizationJobTriggerResultOutput) LastRunTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.LastRunTime }).(pulumi.StringOutput)
}

// Unique resource name for the triggeredJob, assigned by the service when the triggeredJob is created, for example `projects/dlp-test-project/jobTriggers/53234423`.
func (o LookupOrganizationJobTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

// A status for this trigger.
func (o LookupOrganizationJobTriggerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.Status }).(pulumi.StringOutput)
}

// A list of triggers which will be OR'ed together. Only one in the list needs to trigger for a job to be started. The list may contain only a single Schedule trigger and must have at least one object.
func (o LookupOrganizationJobTriggerResultOutput) Triggers() GooglePrivacyDlpV2TriggerResponseArrayOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) []GooglePrivacyDlpV2TriggerResponse { return v.Triggers }).(GooglePrivacyDlpV2TriggerResponseArrayOutput)
}

// The last update timestamp of a triggeredJob.
func (o LookupOrganizationJobTriggerResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrganizationJobTriggerResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrganizationJobTriggerResultOutput{})
}