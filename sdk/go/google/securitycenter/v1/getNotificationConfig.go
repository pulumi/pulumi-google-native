// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a notification config.
func LookupNotificationConfig(ctx *pulumi.Context, args *LookupNotificationConfigArgs, opts ...pulumi.InvokeOption) (*LookupNotificationConfigResult, error) {
	var rv LookupNotificationConfigResult
	err := ctx.Invoke("google-native:securitycenter/v1:getNotificationConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationConfigArgs struct {
	NotificationConfigId string `pulumi:"notificationConfigId"`
	OrganizationId       string `pulumi:"organizationId"`
}

type LookupNotificationConfigResult struct {
	// The description of the notification config (max of 1024 characters).
	Description string `pulumi:"description"`
	// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
	Name string `pulumi:"name"`
	// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
	PubsubTopic string `pulumi:"pubsubTopic"`
	// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
	ServiceAccount string `pulumi:"serviceAccount"`
	// The config for triggering streaming-based notifications.
	StreamingConfig StreamingConfigResponse `pulumi:"streamingConfig"`
}

func LookupNotificationConfigOutput(ctx *pulumi.Context, args LookupNotificationConfigOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationConfigResult, error) {
			args := v.(LookupNotificationConfigArgs)
			r, err := LookupNotificationConfig(ctx, &args, opts...)
			var s LookupNotificationConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotificationConfigResultOutput)
}

type LookupNotificationConfigOutputArgs struct {
	NotificationConfigId pulumi.StringInput `pulumi:"notificationConfigId"`
	OrganizationId       pulumi.StringInput `pulumi:"organizationId"`
}

func (LookupNotificationConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationConfigArgs)(nil)).Elem()
}

type LookupNotificationConfigResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationConfigResult)(nil)).Elem()
}

func (o LookupNotificationConfigResultOutput) ToLookupNotificationConfigResultOutput() LookupNotificationConfigResultOutput {
	return o
}

func (o LookupNotificationConfigResultOutput) ToLookupNotificationConfigResultOutputWithContext(ctx context.Context) LookupNotificationConfigResultOutput {
	return o
}

// The description of the notification config (max of 1024 characters).
func (o LookupNotificationConfigResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationConfigResult) string { return v.Description }).(pulumi.StringOutput)
}

// The relative resource name of this notification config. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/notificationConfigs/notify_public_bucket".
func (o LookupNotificationConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

// The Pub/Sub topic to send notifications to. Its format is "projects/[project_id]/topics/[topic]".
func (o LookupNotificationConfigResultOutput) PubsubTopic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationConfigResult) string { return v.PubsubTopic }).(pulumi.StringOutput)
}

// The service account that needs "pubsub.topics.publish" permission to publish to the Pub/Sub topic.
func (o LookupNotificationConfigResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotificationConfigResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The config for triggering streaming-based notifications.
func (o LookupNotificationConfigResultOutput) StreamingConfig() StreamingConfigResponseOutput {
	return o.ApplyT(func(v LookupNotificationConfigResult) StreamingConfigResponse { return v.StreamingConfig }).(StreamingConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationConfigResultOutput{})
}