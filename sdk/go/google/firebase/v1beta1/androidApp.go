// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Requests the creation of a new AndroidApp in the specified FirebaseProject. The result of this call is an `Operation` which can be used to track the provisioning process. The `Operation` is automatically deleted after completion, so there is no need to call `DeleteOperation`.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type AndroidApp struct {
	pulumi.CustomResourceState

	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId pulumi.StringOutput `pulumi:"apiKeyId"`
	// Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The user-assigned display name for the `AndroidApp`.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName pulumi.StringOutput `pulumi:"packageName"`
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewAndroidApp registers a new resource with the given unique name, arguments, and options.
func NewAndroidApp(ctx *pulumi.Context,
	name string, args *AndroidAppArgs, opts ...pulumi.ResourceOption) (*AndroidApp, error) {
	if args == nil {
		args = &AndroidAppArgs{}
	}

	var resource AndroidApp
	err := ctx.RegisterResource("google-native:firebase/v1beta1:AndroidApp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAndroidApp gets an existing AndroidApp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAndroidApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AndroidAppState, opts ...pulumi.ResourceOption) (*AndroidApp, error) {
	var resource AndroidApp
	err := ctx.ReadResource("google-native:firebase/v1beta1:AndroidApp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AndroidApp resources.
type androidAppState struct {
}

type AndroidAppState struct {
}

func (AndroidAppState) ElementType() reflect.Type {
	return reflect.TypeOf((*androidAppState)(nil)).Elem()
}

type androidAppArgs struct {
	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId *string `pulumi:"apiKeyId"`
	// Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId *string `pulumi:"appId"`
	// The user-assigned display name for the `AndroidApp`.
	DisplayName *string `pulumi:"displayName"`
	// The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
	Name *string `pulumi:"name"`
	// Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName *string `pulumi:"packageName"`
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a AndroidApp resource.
type AndroidAppArgs struct {
	// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
	ApiKeyId pulumi.StringPtrInput
	// Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
	AppId pulumi.StringPtrInput
	// The user-assigned display name for the `AndroidApp`.
	DisplayName pulumi.StringPtrInput
	// The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
	Name pulumi.StringPtrInput
	// Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
	PackageName pulumi.StringPtrInput
	// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.
	Project pulumi.StringPtrInput
}

func (AndroidAppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*androidAppArgs)(nil)).Elem()
}

type AndroidAppInput interface {
	pulumi.Input

	ToAndroidAppOutput() AndroidAppOutput
	ToAndroidAppOutputWithContext(ctx context.Context) AndroidAppOutput
}

func (*AndroidApp) ElementType() reflect.Type {
	return reflect.TypeOf((**AndroidApp)(nil)).Elem()
}

func (i *AndroidApp) ToAndroidAppOutput() AndroidAppOutput {
	return i.ToAndroidAppOutputWithContext(context.Background())
}

func (i *AndroidApp) ToAndroidAppOutputWithContext(ctx context.Context) AndroidAppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AndroidAppOutput)
}

type AndroidAppOutput struct{ *pulumi.OutputState }

func (AndroidAppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AndroidApp)(nil)).Elem()
}

func (o AndroidAppOutput) ToAndroidAppOutput() AndroidAppOutput {
	return o
}

func (o AndroidAppOutput) ToAndroidAppOutputWithContext(ctx context.Context) AndroidAppOutput {
	return o
}

// The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.
func (o AndroidAppOutput) ApiKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidApp) pulumi.StringOutput { return v.ApiKeyId }).(pulumi.StringOutput)
}

// Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
func (o AndroidAppOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidApp) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The user-assigned display name for the `AndroidApp`.
func (o AndroidAppOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidApp) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
func (o AndroidAppOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidApp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
func (o AndroidAppOutput) PackageName() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidApp) pulumi.StringOutput { return v.PackageName }).(pulumi.StringOutput)
}

// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.
func (o AndroidAppOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AndroidApp) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AndroidAppInput)(nil)).Elem(), &AndroidApp{})
	pulumi.RegisterOutputType(AndroidAppOutput{})
}