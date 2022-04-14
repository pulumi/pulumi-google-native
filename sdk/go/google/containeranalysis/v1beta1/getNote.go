// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified note.
func LookupNote(ctx *pulumi.Context, args *LookupNoteArgs, opts ...pulumi.InvokeOption) (*LookupNoteResult, error) {
	var rv LookupNoteResult
	err := ctx.Invoke("google-native:containeranalysis/v1beta1:getNote", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNoteArgs struct {
	NoteId  string  `pulumi:"noteId"`
	Project *string `pulumi:"project"`
}

type LookupNoteResult struct {
	// A note describing an attestation role.
	AttestationAuthority AuthorityResponse `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage BasisResponse `pulumi:"baseImage"`
	// A note describing build provenance for a verifiable build.
	Build BuildResponse `pulumi:"build"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime string `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployable DeployableResponse `pulumi:"deployable"`
	// A note describing the initial analysis of a resource.
	Discovery DiscoveryResponse `pulumi:"discovery"`
	// Time of expiration for this note. Empty if note does not expire.
	ExpirationTime string `pulumi:"expirationTime"`
	// A note describing an in-toto link.
	Intoto InTotoResponse `pulumi:"intoto"`
	// The type of analysis. This field can be used as a filter in list requests.
	Kind string `pulumi:"kind"`
	// A detailed description of this note.
	LongDescription string `pulumi:"longDescription"`
	// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
	Name string `pulumi:"name"`
	// A note describing a package hosted by various package managers.
	Package PackageResponse `pulumi:"package"`
	// Other notes related to this note.
	RelatedNoteNames []string `pulumi:"relatedNoteNames"`
	// URLs associated with this note.
	RelatedUrl []RelatedUrlResponse `pulumi:"relatedUrl"`
	// A note describing a software bill of materials.
	Sbom DocumentNoteResponse `pulumi:"sbom"`
	// A one sentence description of this note.
	ShortDescription string `pulumi:"shortDescription"`
	// A note describing an SPDX File.
	SpdxFile FileNoteResponse `pulumi:"spdxFile"`
	// A note describing an SPDX Package.
	SpdxPackage PackageInfoNoteResponse `pulumi:"spdxPackage"`
	// A note describing an SPDX File.
	SpdxRelationship RelationshipNoteResponse `pulumi:"spdxRelationship"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime string `pulumi:"updateTime"`
	// A note describing a package vulnerability.
	Vulnerability VulnerabilityResponse `pulumi:"vulnerability"`
}

func LookupNoteOutput(ctx *pulumi.Context, args LookupNoteOutputArgs, opts ...pulumi.InvokeOption) LookupNoteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNoteResult, error) {
			args := v.(LookupNoteArgs)
			r, err := LookupNote(ctx, &args, opts...)
			return *r, err
		}).(LookupNoteResultOutput)
}

type LookupNoteOutputArgs struct {
	NoteId  pulumi.StringInput    `pulumi:"noteId"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupNoteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNoteArgs)(nil)).Elem()
}

type LookupNoteResultOutput struct{ *pulumi.OutputState }

func (LookupNoteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNoteResult)(nil)).Elem()
}

func (o LookupNoteResultOutput) ToLookupNoteResultOutput() LookupNoteResultOutput {
	return o
}

func (o LookupNoteResultOutput) ToLookupNoteResultOutputWithContext(ctx context.Context) LookupNoteResultOutput {
	return o
}

// A note describing an attestation role.
func (o LookupNoteResultOutput) AttestationAuthority() AuthorityResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) AuthorityResponse { return v.AttestationAuthority }).(AuthorityResponseOutput)
}

// A note describing a base image.
func (o LookupNoteResultOutput) BaseImage() BasisResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) BasisResponse { return v.BaseImage }).(BasisResponseOutput)
}

// A note describing build provenance for a verifiable build.
func (o LookupNoteResultOutput) Build() BuildResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) BuildResponse { return v.Build }).(BuildResponseOutput)
}

// The time this note was created. This field can be used as a filter in list requests.
func (o LookupNoteResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A note describing something that can be deployed.
func (o LookupNoteResultOutput) Deployable() DeployableResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) DeployableResponse { return v.Deployable }).(DeployableResponseOutput)
}

// A note describing the initial analysis of a resource.
func (o LookupNoteResultOutput) Discovery() DiscoveryResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) DiscoveryResponse { return v.Discovery }).(DiscoveryResponseOutput)
}

// Time of expiration for this note. Empty if note does not expire.
func (o LookupNoteResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

// A note describing an in-toto link.
func (o LookupNoteResultOutput) Intoto() InTotoResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) InTotoResponse { return v.Intoto }).(InTotoResponseOutput)
}

// The type of analysis. This field can be used as a filter in list requests.
func (o LookupNoteResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A detailed description of this note.
func (o LookupNoteResultOutput) LongDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.LongDescription }).(pulumi.StringOutput)
}

// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
func (o LookupNoteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.Name }).(pulumi.StringOutput)
}

// A note describing a package hosted by various package managers.
func (o LookupNoteResultOutput) Package() PackageResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) PackageResponse { return v.Package }).(PackageResponseOutput)
}

// Other notes related to this note.
func (o LookupNoteResultOutput) RelatedNoteNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNoteResult) []string { return v.RelatedNoteNames }).(pulumi.StringArrayOutput)
}

// URLs associated with this note.
func (o LookupNoteResultOutput) RelatedUrl() RelatedUrlResponseArrayOutput {
	return o.ApplyT(func(v LookupNoteResult) []RelatedUrlResponse { return v.RelatedUrl }).(RelatedUrlResponseArrayOutput)
}

// A note describing a software bill of materials.
func (o LookupNoteResultOutput) Sbom() DocumentNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) DocumentNoteResponse { return v.Sbom }).(DocumentNoteResponseOutput)
}

// A one sentence description of this note.
func (o LookupNoteResultOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.ShortDescription }).(pulumi.StringOutput)
}

// A note describing an SPDX File.
func (o LookupNoteResultOutput) SpdxFile() FileNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) FileNoteResponse { return v.SpdxFile }).(FileNoteResponseOutput)
}

// A note describing an SPDX Package.
func (o LookupNoteResultOutput) SpdxPackage() PackageInfoNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) PackageInfoNoteResponse { return v.SpdxPackage }).(PackageInfoNoteResponseOutput)
}

// A note describing an SPDX File.
func (o LookupNoteResultOutput) SpdxRelationship() RelationshipNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) RelationshipNoteResponse { return v.SpdxRelationship }).(RelationshipNoteResponseOutput)
}

// The time this note was last updated. This field can be used as a filter in list requests.
func (o LookupNoteResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// A note describing a package vulnerability.
func (o LookupNoteResultOutput) Vulnerability() VulnerabilityResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) VulnerabilityResponse { return v.Vulnerability }).(VulnerabilityResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNoteResultOutput{})
}