// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The configuration used for a Replay.
type GoogleCloudPolicysimulatorV1ReplayConfig struct {
	// The logs to use as input for the Replay.
	LogSource *GoogleCloudPolicysimulatorV1ReplayConfigLogSource `pulumi:"logSource"`
	// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
	PolicyOverlay map[string]string `pulumi:"policyOverlay"`
}

// GoogleCloudPolicysimulatorV1ReplayConfigInput is an input type that accepts GoogleCloudPolicysimulatorV1ReplayConfigArgs and GoogleCloudPolicysimulatorV1ReplayConfigOutput values.
// You can construct a concrete instance of `GoogleCloudPolicysimulatorV1ReplayConfigInput` via:
//
//          GoogleCloudPolicysimulatorV1ReplayConfigArgs{...}
type GoogleCloudPolicysimulatorV1ReplayConfigInput interface {
	pulumi.Input

	ToGoogleCloudPolicysimulatorV1ReplayConfigOutput() GoogleCloudPolicysimulatorV1ReplayConfigOutput
	ToGoogleCloudPolicysimulatorV1ReplayConfigOutputWithContext(context.Context) GoogleCloudPolicysimulatorV1ReplayConfigOutput
}

// The configuration used for a Replay.
type GoogleCloudPolicysimulatorV1ReplayConfigArgs struct {
	// The logs to use as input for the Replay.
	LogSource GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrInput `pulumi:"logSource"`
	// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
	PolicyOverlay pulumi.StringMapInput `pulumi:"policyOverlay"`
}

func (GoogleCloudPolicysimulatorV1ReplayConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfig)(nil)).Elem()
}

func (i GoogleCloudPolicysimulatorV1ReplayConfigArgs) ToGoogleCloudPolicysimulatorV1ReplayConfigOutput() GoogleCloudPolicysimulatorV1ReplayConfigOutput {
	return i.ToGoogleCloudPolicysimulatorV1ReplayConfigOutputWithContext(context.Background())
}

func (i GoogleCloudPolicysimulatorV1ReplayConfigArgs) ToGoogleCloudPolicysimulatorV1ReplayConfigOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudPolicysimulatorV1ReplayConfigOutput)
}

// The configuration used for a Replay.
type GoogleCloudPolicysimulatorV1ReplayConfigOutput struct{ *pulumi.OutputState }

func (GoogleCloudPolicysimulatorV1ReplayConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfig)(nil)).Elem()
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigOutput() GoogleCloudPolicysimulatorV1ReplayConfigOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigOutput {
	return o
}

// The logs to use as input for the Replay.
func (o GoogleCloudPolicysimulatorV1ReplayConfigOutput) LogSource() GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayConfig) *GoogleCloudPolicysimulatorV1ReplayConfigLogSource {
		return v.LogSource
	}).(GoogleCloudPolicysimulatorV1ReplayConfigLogSourcePtrOutput)
}

// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
func (o GoogleCloudPolicysimulatorV1ReplayConfigOutput) PolicyOverlay() pulumi.StringMapOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayConfig) map[string]string { return v.PolicyOverlay }).(pulumi.StringMapOutput)
}

// The configuration used for a Replay.
type GoogleCloudPolicysimulatorV1ReplayConfigResponse struct {
	// The logs to use as input for the Replay.
	LogSource string `pulumi:"logSource"`
	// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
	PolicyOverlay map[string]string `pulumi:"policyOverlay"`
}

// The configuration used for a Replay.
type GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput struct{ *pulumi.OutputState }

func (GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfigResponse)(nil)).Elem()
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigResponseOutput() GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput) ToGoogleCloudPolicysimulatorV1ReplayConfigResponseOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput {
	return o
}

// The logs to use as input for the Replay.
func (o GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput) LogSource() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayConfigResponse) string { return v.LogSource }).(pulumi.StringOutput)
}

// A mapping of the resources that you want to simulate policies for and the policies that you want to simulate. Keys are the full resource names for the resources. For example, `//cloudresourcemanager.googleapis.com/projects/my-project`. For examples of full resource names for Google Cloud services, see https://cloud.google.com/iam/help/troubleshooter/full-resource-names. Values are Policy objects representing the policies that you want to simulate. Replays automatically take into account any IAM policies inherited through the resource hierarchy, and any policies set on descendant resources. You do not need to include these policies in the policy overlay.
func (o GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput) PolicyOverlay() pulumi.StringMapOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayConfigResponse) map[string]string { return v.PolicyOverlay }).(pulumi.StringMapOutput)
}

// Summary statistics about the replayed log entries.
type GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse struct {
	// The number of replayed log entries with a difference between baseline and simulated policies.
	DifferenceCount int `pulumi:"differenceCount"`
	// The number of log entries that could not be replayed.
	ErrorCount int `pulumi:"errorCount"`
	// The total number of log entries replayed.
	LogCount int `pulumi:"logCount"`
	// The date of the newest log entry replayed.
	NewestDate GoogleTypeDateResponse `pulumi:"newestDate"`
	// The date of the oldest log entry replayed.
	OldestDate GoogleTypeDateResponse `pulumi:"oldestDate"`
	// The number of replayed log entries with no difference between baseline and simulated policies.
	UnchangedCount int `pulumi:"unchangedCount"`
}

// Summary statistics about the replayed log entries.
type GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput struct{ *pulumi.OutputState }

func (GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse)(nil)).Elem()
}

func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) ToGoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput() GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput {
	return o
}

func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) ToGoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutputWithContext(ctx context.Context) GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput {
	return o
}

// The number of replayed log entries with a difference between baseline and simulated policies.
func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) DifferenceCount() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse) int { return v.DifferenceCount }).(pulumi.IntOutput)
}

// The number of log entries that could not be replayed.
func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) ErrorCount() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse) int { return v.ErrorCount }).(pulumi.IntOutput)
}

// The total number of log entries replayed.
func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) LogCount() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse) int { return v.LogCount }).(pulumi.IntOutput)
}

// The date of the newest log entry replayed.
func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) NewestDate() GoogleTypeDateResponseOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse) GoogleTypeDateResponse {
		return v.NewestDate
	}).(GoogleTypeDateResponseOutput)
}

// The date of the oldest log entry replayed.
func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) OldestDate() GoogleTypeDateResponseOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse) GoogleTypeDateResponse {
		return v.OldestDate
	}).(GoogleTypeDateResponseOutput)
}

// The number of replayed log entries with no difference between baseline and simulated policies.
func (o GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput) UnchangedCount() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse) int { return v.UnchangedCount }).(pulumi.IntOutput)
}

// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
type GoogleTypeDateResponse struct {
	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
	Day int `pulumi:"day"`
	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
	Month int `pulumi:"month"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	Year int `pulumi:"year"`
}

// Represents a whole or partial calendar date, such as a birthday. The time of day and time zone are either specified elsewhere or are insignificant. The date is relative to the Gregorian Calendar. This can represent one of the following: * A full date, with non-zero year, month, and day values. * A month and day, with a zero year (for example, an anniversary). * A year on its own, with a zero month and a zero day. * A year and month, with a zero day (for example, a credit card expiration date). Related types: * google.type.TimeOfDay * google.type.DateTime * google.protobuf.Timestamp
type GoogleTypeDateResponseOutput struct{ *pulumi.OutputState }

func (GoogleTypeDateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleTypeDateResponse)(nil)).Elem()
}

func (o GoogleTypeDateResponseOutput) ToGoogleTypeDateResponseOutput() GoogleTypeDateResponseOutput {
	return o
}

func (o GoogleTypeDateResponseOutput) ToGoogleTypeDateResponseOutputWithContext(ctx context.Context) GoogleTypeDateResponseOutput {
	return o
}

// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0 to specify a year by itself or a year and month where the day isn't significant.
func (o GoogleTypeDateResponseOutput) Day() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleTypeDateResponse) int { return v.Day }).(pulumi.IntOutput)
}

// Month of a year. Must be from 1 to 12, or 0 to specify a year without a month and day.
func (o GoogleTypeDateResponseOutput) Month() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleTypeDateResponse) int { return v.Month }).(pulumi.IntOutput)
}

// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
func (o GoogleTypeDateResponseOutput) Year() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleTypeDateResponse) int { return v.Year }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudPolicysimulatorV1ReplayConfigInput)(nil)).Elem(), GoogleCloudPolicysimulatorV1ReplayConfigArgs{})
	pulumi.RegisterOutputType(GoogleCloudPolicysimulatorV1ReplayConfigOutput{})
	pulumi.RegisterOutputType(GoogleCloudPolicysimulatorV1ReplayConfigResponseOutput{})
	pulumi.RegisterOutputType(GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponseOutput{})
	pulumi.RegisterOutputType(GoogleTypeDateResponseOutput{})
}