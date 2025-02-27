// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a NotebookRuntimeTemplate.
func LookupNotebookRuntimeTemplate(ctx *pulumi.Context, args *LookupNotebookRuntimeTemplateArgs, opts ...pulumi.InvokeOption) (*LookupNotebookRuntimeTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNotebookRuntimeTemplateResult
	err := ctx.Invoke("google-native:aiplatform/v1beta1:getNotebookRuntimeTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotebookRuntimeTemplateArgs struct {
	Location                  string  `pulumi:"location"`
	NotebookRuntimeTemplateId string  `pulumi:"notebookRuntimeTemplateId"`
	Project                   *string `pulumi:"project"`
}

type LookupNotebookRuntimeTemplateResult struct {
	// Timestamp when this NotebookRuntimeTemplate was created.
	CreateTime string `pulumi:"createTime"`
	// Optional. The specification of persistent disk attached to the runtime as data disk storage.
	DataPersistentDiskSpec GoogleCloudAiplatformV1beta1PersistentDiskSpecResponse `pulumi:"dataPersistentDiskSpec"`
	// The description of the NotebookRuntimeTemplate.
	Description string `pulumi:"description"`
	// The display name of the NotebookRuntimeTemplate. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	DisplayName string `pulumi:"displayName"`
	// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
	Etag string `pulumi:"etag"`
	// EUC configuration of the NotebookRuntimeTemplate.
	EucConfig GoogleCloudAiplatformV1beta1NotebookEucConfigResponse `pulumi:"eucConfig"`
	// The idle shutdown configuration of NotebookRuntimeTemplate. This config will only be set when idle shutdown is enabled.
	IdleShutdownConfig GoogleCloudAiplatformV1beta1NotebookIdleShutdownConfigResponse `pulumi:"idleShutdownConfig"`
	// The default template to use if not specified.
	IsDefault bool `pulumi:"isDefault"`
	// The labels with user-defined metadata to organize the NotebookRuntimeTemplates. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Immutable. The specification of a single machine for the template.
	MachineSpec GoogleCloudAiplatformV1beta1MachineSpecResponse `pulumi:"machineSpec"`
	// The resource name of the NotebookRuntimeTemplate.
	Name string `pulumi:"name"`
	// Optional. Network spec.
	NetworkSpec GoogleCloudAiplatformV1beta1NetworkSpecResponse `pulumi:"networkSpec"`
	// Optional. Immutable. The type of the notebook runtime template.
	NotebookRuntimeType string `pulumi:"notebookRuntimeType"`
	// The service account that the runtime workload runs as. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Timestamp when this NotebookRuntimeTemplate was most recently updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupNotebookRuntimeTemplateOutput(ctx *pulumi.Context, args LookupNotebookRuntimeTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupNotebookRuntimeTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotebookRuntimeTemplateResult, error) {
			args := v.(LookupNotebookRuntimeTemplateArgs)
			r, err := LookupNotebookRuntimeTemplate(ctx, &args, opts...)
			var s LookupNotebookRuntimeTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotebookRuntimeTemplateResultOutput)
}

type LookupNotebookRuntimeTemplateOutputArgs struct {
	Location                  pulumi.StringInput    `pulumi:"location"`
	NotebookRuntimeTemplateId pulumi.StringInput    `pulumi:"notebookRuntimeTemplateId"`
	Project                   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupNotebookRuntimeTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookRuntimeTemplateArgs)(nil)).Elem()
}

type LookupNotebookRuntimeTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupNotebookRuntimeTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookRuntimeTemplateResult)(nil)).Elem()
}

func (o LookupNotebookRuntimeTemplateResultOutput) ToLookupNotebookRuntimeTemplateResultOutput() LookupNotebookRuntimeTemplateResultOutput {
	return o
}

func (o LookupNotebookRuntimeTemplateResultOutput) ToLookupNotebookRuntimeTemplateResultOutputWithContext(ctx context.Context) LookupNotebookRuntimeTemplateResultOutput {
	return o
}

// Timestamp when this NotebookRuntimeTemplate was created.
func (o LookupNotebookRuntimeTemplateResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. The specification of persistent disk attached to the runtime as data disk storage.
func (o LookupNotebookRuntimeTemplateResultOutput) DataPersistentDiskSpec() GoogleCloudAiplatformV1beta1PersistentDiskSpecResponseOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) GoogleCloudAiplatformV1beta1PersistentDiskSpecResponse {
		return v.DataPersistentDiskSpec
	}).(GoogleCloudAiplatformV1beta1PersistentDiskSpecResponseOutput)
}

// The description of the NotebookRuntimeTemplate.
func (o LookupNotebookRuntimeTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// The display name of the NotebookRuntimeTemplate. The name can be up to 128 characters long and can consist of any UTF-8 characters.
func (o LookupNotebookRuntimeTemplateResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
func (o LookupNotebookRuntimeTemplateResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.Etag }).(pulumi.StringOutput)
}

// EUC configuration of the NotebookRuntimeTemplate.
func (o LookupNotebookRuntimeTemplateResultOutput) EucConfig() GoogleCloudAiplatformV1beta1NotebookEucConfigResponseOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) GoogleCloudAiplatformV1beta1NotebookEucConfigResponse {
		return v.EucConfig
	}).(GoogleCloudAiplatformV1beta1NotebookEucConfigResponseOutput)
}

// The idle shutdown configuration of NotebookRuntimeTemplate. This config will only be set when idle shutdown is enabled.
func (o LookupNotebookRuntimeTemplateResultOutput) IdleShutdownConfig() GoogleCloudAiplatformV1beta1NotebookIdleShutdownConfigResponseOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) GoogleCloudAiplatformV1beta1NotebookIdleShutdownConfigResponse {
		return v.IdleShutdownConfig
	}).(GoogleCloudAiplatformV1beta1NotebookIdleShutdownConfigResponseOutput)
}

// The default template to use if not specified.
func (o LookupNotebookRuntimeTemplateResultOutput) IsDefault() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) bool { return v.IsDefault }).(pulumi.BoolOutput)
}

// The labels with user-defined metadata to organize the NotebookRuntimeTemplates. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
func (o LookupNotebookRuntimeTemplateResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Immutable. The specification of a single machine for the template.
func (o LookupNotebookRuntimeTemplateResultOutput) MachineSpec() GoogleCloudAiplatformV1beta1MachineSpecResponseOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) GoogleCloudAiplatformV1beta1MachineSpecResponse {
		return v.MachineSpec
	}).(GoogleCloudAiplatformV1beta1MachineSpecResponseOutput)
}

// The resource name of the NotebookRuntimeTemplate.
func (o LookupNotebookRuntimeTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. Network spec.
func (o LookupNotebookRuntimeTemplateResultOutput) NetworkSpec() GoogleCloudAiplatformV1beta1NetworkSpecResponseOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) GoogleCloudAiplatformV1beta1NetworkSpecResponse {
		return v.NetworkSpec
	}).(GoogleCloudAiplatformV1beta1NetworkSpecResponseOutput)
}

// Optional. Immutable. The type of the notebook runtime template.
func (o LookupNotebookRuntimeTemplateResultOutput) NotebookRuntimeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.NotebookRuntimeType }).(pulumi.StringOutput)
}

// The service account that the runtime workload runs as. You can use any service account within the same project, but you must have the service account user permission to use the instance. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.
func (o LookupNotebookRuntimeTemplateResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Timestamp when this NotebookRuntimeTemplate was most recently updated.
func (o LookupNotebookRuntimeTemplateResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookRuntimeTemplateResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotebookRuntimeTemplateResultOutput{})
}
