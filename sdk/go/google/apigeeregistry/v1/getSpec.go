// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// GetApiSpec returns a specified spec.
func LookupSpec(ctx *pulumi.Context, args *LookupSpecArgs, opts ...pulumi.InvokeOption) (*LookupSpecResult, error) {
	var rv LookupSpecResult
	err := ctx.Invoke("google-native:apigeeregistry/v1:getSpec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSpecArgs struct {
	ApiId     string  `pulumi:"apiId"`
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	SpecId    string  `pulumi:"specId"`
	VersionId string  `pulumi:"versionId"`
}

type LookupSpecResult struct {
	// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
	Annotations map[string]string `pulumi:"annotations"`
	// Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
	Contents string `pulumi:"contents"`
	// Creation timestamp; when the spec resource was created.
	CreateTime string `pulumi:"createTime"`
	// A detailed description.
	Description string `pulumi:"description"`
	// A possibly-hierarchical name used to refer to the spec from other specs.
	Filename string `pulumi:"filename"`
	// A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.
	Hash string `pulumi:"hash"`
	// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
	Labels map[string]string `pulumi:"labels"`
	// A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include "application/vnd.apigee.proto", "application/vnd.apigee.openapi", and "application/vnd.apigee.graphql", with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
	MimeType string `pulumi:"mimeType"`
	// Resource name.
	Name string `pulumi:"name"`
	// Revision creation timestamp; when the represented revision was created.
	RevisionCreateTime string `pulumi:"revisionCreateTime"`
	// Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.
	RevisionId string `pulumi:"revisionId"`
	// Last update timestamp: when the represented revision was last modified.
	RevisionUpdateTime string `pulumi:"revisionUpdateTime"`
	// The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.
	SizeBytes int `pulumi:"sizeBytes"`
	// The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
	SourceUri string `pulumi:"sourceUri"`
}

func LookupSpecOutput(ctx *pulumi.Context, args LookupSpecOutputArgs, opts ...pulumi.InvokeOption) LookupSpecResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpecResult, error) {
			args := v.(LookupSpecArgs)
			r, err := LookupSpec(ctx, &args, opts...)
			var s LookupSpecResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSpecResultOutput)
}

type LookupSpecOutputArgs struct {
	ApiId     pulumi.StringInput    `pulumi:"apiId"`
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	SpecId    pulumi.StringInput    `pulumi:"specId"`
	VersionId pulumi.StringInput    `pulumi:"versionId"`
}

func (LookupSpecOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpecArgs)(nil)).Elem()
}

type LookupSpecResultOutput struct{ *pulumi.OutputState }

func (LookupSpecResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpecResult)(nil)).Elem()
}

func (o LookupSpecResultOutput) ToLookupSpecResultOutput() LookupSpecResultOutput {
	return o
}

func (o LookupSpecResultOutput) ToLookupSpecResultOutputWithContext(ctx context.Context) LookupSpecResultOutput {
	return o
}

// Annotations attach non-identifying metadata to resources. Annotation keys and values are less restricted than those of labels, but should be generally used for small values of broad interest. Larger, topic- specific metadata should be stored in Artifacts.
func (o LookupSpecResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSpecResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// Input only. The contents of the spec. Provided by API callers when specs are created or updated. To access the contents of a spec, use GetApiSpecContents.
func (o LookupSpecResultOutput) Contents() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.Contents }).(pulumi.StringOutput)
}

// Creation timestamp; when the spec resource was created.
func (o LookupSpecResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A detailed description.
func (o LookupSpecResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.Description }).(pulumi.StringOutput)
}

// A possibly-hierarchical name used to refer to the spec from other specs.
func (o LookupSpecResultOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.Filename }).(pulumi.StringOutput)
}

// A SHA-256 hash of the spec's contents. If the spec is gzipped, this is the hash of the uncompressed spec.
func (o LookupSpecResultOutput) Hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.Hash }).(pulumi.StringOutput)
}

// Labels attach identifying metadata to resources. Identifying metadata can be used to filter list operations. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one resource (System labels are excluded). See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "apigeeregistry.googleapis.com/" and cannot be changed.
func (o LookupSpecResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSpecResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// A style (format) descriptor for this spec that is specified as a Media Type (https://en.wikipedia.org/wiki/Media_type). Possible values include "application/vnd.apigee.proto", "application/vnd.apigee.openapi", and "application/vnd.apigee.graphql", with possible suffixes representing compression types. These hypothetical names are defined in the vendor tree defined in RFC6838 (https://tools.ietf.org/html/rfc6838) and are not final. Content types can specify compression. Currently only GZip compression is supported (indicated with "+gzip").
func (o LookupSpecResultOutput) MimeType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.MimeType }).(pulumi.StringOutput)
}

// Resource name.
func (o LookupSpecResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.Name }).(pulumi.StringOutput)
}

// Revision creation timestamp; when the represented revision was created.
func (o LookupSpecResultOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Immutable. The revision ID of the spec. A new revision is committed whenever the spec contents are changed. The format is an 8-character hexadecimal string.
func (o LookupSpecResultOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.RevisionId }).(pulumi.StringOutput)
}

// Last update timestamp: when the represented revision was last modified.
func (o LookupSpecResultOutput) RevisionUpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.RevisionUpdateTime }).(pulumi.StringOutput)
}

// The size of the spec file in bytes. If the spec is gzipped, this is the size of the uncompressed spec.
func (o LookupSpecResultOutput) SizeBytes() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSpecResult) int { return v.SizeBytes }).(pulumi.IntOutput)
}

// The original source URI of the spec (if one exists). This is an external location that can be used for reference purposes but which may not be authoritative since this external resource may change after the spec is retrieved.
func (o LookupSpecResultOutput) SourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpecResult) string { return v.SourceUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpecResultOutput{})
}