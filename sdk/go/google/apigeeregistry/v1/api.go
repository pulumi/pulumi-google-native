// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateApi creates a specified API.
type Api struct {
	pulumi.CustomResourceState

	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g. "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
	Availability pulumi.StringOutput `pulumi:"availability"`
	// Creation timestamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A detailed description.
	Description pulumi.StringOutput `pulumi:"description"`
	// Human-meaningful name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The recommended deployment of the API. Format: apis/{api}/deployments/{deployment}
	RecommendedDeployment pulumi.StringOutput `pulumi:"recommendedDeployment"`
	// The recommended version of the API. Format: apis/{api}/versions/{version}
	RecommendedVersion pulumi.StringOutput `pulumi:"recommendedVersion"`
	// Last update timestamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewApi registers a new resource with the given unique name, arguments, and options.
func NewApi(ctx *pulumi.Context,
	name string, args *ApiArgs, opts ...pulumi.ResourceOption) (*Api, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	var resource Api
	err := ctx.RegisterResource("google-native:apigeeregistry/v1:Api", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApi gets an existing Api resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApi(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiState, opts ...pulumi.ResourceOption) (*Api, error) {
	var resource Api
	err := ctx.ReadResource("google-native:apigeeregistry/v1:Api", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Api resources.
type apiState struct {
}

type ApiState struct {
}

func (ApiState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiState)(nil)).Elem()
}

type apiArgs struct {
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations map[string]string `pulumi:"annotations"`
	// Required. The ID to use for the api, which will become the final component of the api's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.
	ApiId string `pulumi:"apiId"`
	// A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g. "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
	Availability *string `pulumi:"availability"`
	// A detailed description.
	Description *string `pulumi:"description"`
	// Human-meaningful name.
	DisplayName *string `pulumi:"displayName"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Resource name.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The recommended deployment of the API. Format: apis/{api}/deployments/{deployment}
	RecommendedDeployment *string `pulumi:"recommendedDeployment"`
	// The recommended version of the API. Format: apis/{api}/versions/{version}
	RecommendedVersion *string `pulumi:"recommendedVersion"`
}

// The set of arguments for constructing a Api resource.
type ApiArgs struct {
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations pulumi.StringMapInput
	// Required. The ID to use for the api, which will become the final component of the api's resource name. This value should be 4-63 characters, and valid characters are /a-z-/. Following AIP-162, IDs must not have the form of a UUID.
	ApiId pulumi.StringInput
	// A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g. "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
	Availability pulumi.StringPtrInput
	// A detailed description.
	Description pulumi.StringPtrInput
	// Human-meaningful name.
	DisplayName pulumi.StringPtrInput
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Resource name.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The recommended deployment of the API. Format: apis/{api}/deployments/{deployment}
	RecommendedDeployment pulumi.StringPtrInput
	// The recommended version of the API. Format: apis/{api}/versions/{version}
	RecommendedVersion pulumi.StringPtrInput
}

func (ApiArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiArgs)(nil)).Elem()
}

type ApiInput interface {
	pulumi.Input

	ToApiOutput() ApiOutput
	ToApiOutputWithContext(ctx context.Context) ApiOutput
}

func (*Api) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (i *Api) ToApiOutput() ApiOutput {
	return i.ToApiOutputWithContext(context.Background())
}

func (i *Api) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiOutput)
}

type ApiOutput struct{ *pulumi.OutputState }

func (ApiOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Api)(nil)).Elem()
}

func (o ApiOutput) ToApiOutput() ApiOutput {
	return o
}

func (o ApiOutput) ToApiOutputWithContext(ctx context.Context) ApiOutput {
	return o
}

// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
func (o ApiOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Api) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// A user-definable description of the availability of this service. Format: free-form, but we expect single words that describe availability, e.g. "NONE", "TESTING", "PREVIEW", "GENERAL", "DEPRECATED", "SHUTDOWN".
func (o ApiOutput) Availability() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Availability }).(pulumi.StringOutput)
}

// Creation timestamp.
func (o ApiOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A detailed description.
func (o ApiOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Human-meaningful name.
func (o ApiOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
func (o ApiOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Api) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource name.
func (o ApiOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The recommended deployment of the API. Format: apis/{api}/deployments/{deployment}
func (o ApiOutput) RecommendedDeployment() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.RecommendedDeployment }).(pulumi.StringOutput)
}

// The recommended version of the API. Format: apis/{api}/versions/{version}
func (o ApiOutput) RecommendedVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.RecommendedVersion }).(pulumi.StringOutput)
}

// Last update timestamp.
func (o ApiOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Api) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiInput)(nil)).Elem(), &Api{})
	pulumi.RegisterOutputType(ApiOutput{})
}