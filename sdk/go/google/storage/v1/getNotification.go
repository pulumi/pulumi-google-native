// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// View a notification configuration.
func LookupNotification(ctx *pulumi.Context, args *LookupNotificationArgs, opts ...pulumi.InvokeOption) (*LookupNotificationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNotificationResult
	err := ctx.Invoke("google-native:storage/v1:getNotification", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationArgs struct {
	Bucket       string  `pulumi:"bucket"`
	Notification string  `pulumi:"notification"`
	UserProject  *string `pulumi:"userProject"`
}

type LookupNotificationResult struct {
	// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// HTTP 1.1 Entity tag for this subscription notification.
	Etag string `pulumi:"etag"`
	// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
	EventTypes []string `pulumi:"eventTypes"`
	// The kind of item this is. For notifications, this is always storage#notification.
	Kind string `pulumi:"kind"`
	// If present, only apply this notification configuration to object names that begin with this prefix.
	ObjectNamePrefix string `pulumi:"objectNamePrefix"`
	// The desired content of the Payload.
	PayloadFormat string `pulumi:"payloadFormat"`
	// The canonical URL of this notification.
	SelfLink string `pulumi:"selfLink"`
	// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
	Topic string `pulumi:"topic"`
}

func LookupNotificationOutput(ctx *pulumi.Context, args LookupNotificationOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationResult, error) {
			args := v.(LookupNotificationArgs)
			r, err := LookupNotification(ctx, &args, opts...)
			var s LookupNotificationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationResultOutput)
}

type LookupNotificationOutputArgs struct {
	Bucket       pulumi.StringInput    `pulumi:"bucket"`
	Notification pulumi.StringInput    `pulumi:"notification"`
	UserProject  pulumi.StringPtrInput `pulumi:"userProject"`
}

func (LookupNotificationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationArgs)(nil)).Elem()
}

type LookupNotificationResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationResult)(nil)).Elem()
}

func (o LookupNotificationResultOutput) ToLookupNotificationResultOutput() LookupNotificationResultOutput {
	return o
}

func (o LookupNotificationResultOutput) ToLookupNotificationResultOutputWithContext(ctx context.Context) LookupNotificationResultOutput {
	return o
}

// An optional list of additional attributes to attach to each Cloud PubSub message published for this notification subscription.
func (o LookupNotificationResultOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotificationResult) map[string]string { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// HTTP 1.1 Entity tag for this subscription notification.
func (o LookupNotificationResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.Etag }).(pulumi.StringOutput)
}

// If present, only send notifications about listed event types. If empty, sent notifications for all event types.
func (o LookupNotificationResultOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNotificationResult) []string { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// The kind of item this is. For notifications, this is always storage#notification.
func (o LookupNotificationResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.Kind }).(pulumi.StringOutput)
}

// If present, only apply this notification configuration to object names that begin with this prefix.
func (o LookupNotificationResultOutput) ObjectNamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.ObjectNamePrefix }).(pulumi.StringOutput)
}

// The desired content of the Payload.
func (o LookupNotificationResultOutput) PayloadFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.PayloadFormat }).(pulumi.StringOutput)
}

// The canonical URL of this notification.
func (o LookupNotificationResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// The Cloud PubSub topic to which this subscription publishes. Formatted as: '//pubsub.googleapis.com/projects/{project-identifier}/topics/{my-topic}'
func (o LookupNotificationResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationResult) string { return v.Topic }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationResultOutput{})
}
