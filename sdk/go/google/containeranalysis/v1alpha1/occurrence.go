// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new `Occurrence`. Use this method to create `Occurrences` for a resource.
// Auto-naming is currently not supported for this resource.
type Occurrence struct {
	pulumi.CustomResourceState

	// Describes an attestation of an artifact.
	Attestation AttestationResponseOutput `pulumi:"attestation"`
	// Build details for a verifiable build.
	BuildDetails BuildDetailsResponseOutput `pulumi:"buildDetails"`
	// Describes whether or not a resource passes compliance checks.
	Compliance ComplianceOccurrenceResponseOutput `pulumi:"compliance"`
	// The time this `Occurrence` was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Describes the deployment of an artifact on a runtime.
	Deployment DeploymentResponseOutput `pulumi:"deployment"`
	// Describes how this resource derives from the basis in the associated note.
	DerivedImage DerivedResponseOutput `pulumi:"derivedImage"`
	// Describes the initial scan status for this resource.
	Discovered DiscoveredResponseOutput `pulumi:"discovered"`
	// This represents a DSSE attestation occurrence
	DsseAttestation DSSEAttestationOccurrenceResponseOutput `pulumi:"dsseAttestation"`
	// https://github.com/secure-systems-lab/dsse
	Envelope EnvelopeResponseOutput `pulumi:"envelope"`
	// Describes the installation of a package on the linked resource.
	Installation InstallationResponseOutput `pulumi:"installation"`
	// This explicitly denotes which of the `Occurrence` details are specified. This field can be used as a filter in list requests.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The name of the `Occurrence` in the form "projects/{project_id}/occurrences/{OCCURRENCE_ID}"
	Name pulumi.StringOutput `pulumi:"name"`
	// An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
	NoteName pulumi.StringOutput `pulumi:"noteName"`
	// A description of actions that can be taken to remedy the `Note`
	Remediation pulumi.StringOutput `pulumi:"remediation"`
	//  The resource for which the `Occurrence` applies.
	Resource ResourceResponseOutput `pulumi:"resource"`
	// The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
	ResourceUrl pulumi.StringOutput `pulumi:"resourceUrl"`
	// Describes a specific software bill of materials document.
	Sbom DocumentOccurrenceResponseOutput `pulumi:"sbom"`
	// Describes a specific SPDX File.
	SpdxFile FileOccurrenceResponseOutput `pulumi:"spdxFile"`
	// Describes a specific SPDX Package.
	SpdxPackage PackageInfoOccurrenceResponseOutput `pulumi:"spdxPackage"`
	// Describes a specific relationship between SPDX elements.
	SpdxRelationship RelationshipOccurrenceResponseOutput `pulumi:"spdxRelationship"`
	// The time this `Occurrence` was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Describes an upgrade.
	Upgrade UpgradeOccurrenceResponseOutput `pulumi:"upgrade"`
	// Details of a security vulnerability note.
	VulnerabilityDetails VulnerabilityDetailsResponseOutput `pulumi:"vulnerabilityDetails"`
}

// NewOccurrence registers a new resource with the given unique name, arguments, and options.
func NewOccurrence(ctx *pulumi.Context,
	name string, args *OccurrenceArgs, opts ...pulumi.ResourceOption) (*Occurrence, error) {
	if args == nil {
		args = &OccurrenceArgs{}
	}

	var resource Occurrence
	err := ctx.RegisterResource("google-native:containeranalysis/v1alpha1:Occurrence", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOccurrence gets an existing Occurrence resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOccurrence(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OccurrenceState, opts ...pulumi.ResourceOption) (*Occurrence, error) {
	var resource Occurrence
	err := ctx.ReadResource("google-native:containeranalysis/v1alpha1:Occurrence", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Occurrence resources.
type occurrenceState struct {
}

type OccurrenceState struct {
}

func (OccurrenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*occurrenceState)(nil)).Elem()
}

type occurrenceArgs struct {
	// Describes an attestation of an artifact.
	Attestation *Attestation `pulumi:"attestation"`
	// Build details for a verifiable build.
	BuildDetails *BuildDetails `pulumi:"buildDetails"`
	// Describes whether or not a resource passes compliance checks.
	Compliance *ComplianceOccurrence `pulumi:"compliance"`
	// Describes the deployment of an artifact on a runtime.
	Deployment *Deployment `pulumi:"deployment"`
	// Describes how this resource derives from the basis in the associated note.
	DerivedImage *Derived `pulumi:"derivedImage"`
	// Describes the initial scan status for this resource.
	Discovered *Discovered `pulumi:"discovered"`
	// This represents a DSSE attestation occurrence
	DsseAttestation *DSSEAttestationOccurrence `pulumi:"dsseAttestation"`
	// https://github.com/secure-systems-lab/dsse
	Envelope *Envelope `pulumi:"envelope"`
	// Describes the installation of a package on the linked resource.
	Installation *Installation `pulumi:"installation"`
	// An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
	NoteName *string `pulumi:"noteName"`
	Project  *string `pulumi:"project"`
	// A description of actions that can be taken to remedy the `Note`
	Remediation *string `pulumi:"remediation"`
	//  The resource for which the `Occurrence` applies.
	Resource *Resource `pulumi:"resource"`
	// The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
	ResourceUrl *string `pulumi:"resourceUrl"`
	// Describes a specific software bill of materials document.
	Sbom *DocumentOccurrence `pulumi:"sbom"`
	// Describes a specific SPDX File.
	SpdxFile *FileOccurrence `pulumi:"spdxFile"`
	// Describes a specific SPDX Package.
	SpdxPackage *PackageInfoOccurrence `pulumi:"spdxPackage"`
	// Describes a specific relationship between SPDX elements.
	SpdxRelationship *RelationshipOccurrence `pulumi:"spdxRelationship"`
	// Describes an upgrade.
	Upgrade *UpgradeOccurrence `pulumi:"upgrade"`
	// Details of a security vulnerability note.
	VulnerabilityDetails *VulnerabilityDetails `pulumi:"vulnerabilityDetails"`
}

// The set of arguments for constructing a Occurrence resource.
type OccurrenceArgs struct {
	// Describes an attestation of an artifact.
	Attestation AttestationPtrInput
	// Build details for a verifiable build.
	BuildDetails BuildDetailsPtrInput
	// Describes whether or not a resource passes compliance checks.
	Compliance ComplianceOccurrencePtrInput
	// Describes the deployment of an artifact on a runtime.
	Deployment DeploymentPtrInput
	// Describes how this resource derives from the basis in the associated note.
	DerivedImage DerivedPtrInput
	// Describes the initial scan status for this resource.
	Discovered DiscoveredPtrInput
	// This represents a DSSE attestation occurrence
	DsseAttestation DSSEAttestationOccurrencePtrInput
	// https://github.com/secure-systems-lab/dsse
	Envelope EnvelopePtrInput
	// Describes the installation of a package on the linked resource.
	Installation InstallationPtrInput
	// An analysis note associated with this image, in the form "providers/{provider_id}/notes/{NOTE_ID}" This field can be used as a filter in list requests.
	NoteName pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// A description of actions that can be taken to remedy the `Note`
	Remediation pulumi.StringPtrInput
	//  The resource for which the `Occurrence` applies.
	Resource ResourcePtrInput
	// The unique URL of the image or the container for which the `Occurrence` applies. For example, https://gcr.io/project/image@sha256:foo This field can be used as a filter in list requests.
	ResourceUrl pulumi.StringPtrInput
	// Describes a specific software bill of materials document.
	Sbom DocumentOccurrencePtrInput
	// Describes a specific SPDX File.
	SpdxFile FileOccurrencePtrInput
	// Describes a specific SPDX Package.
	SpdxPackage PackageInfoOccurrencePtrInput
	// Describes a specific relationship between SPDX elements.
	SpdxRelationship RelationshipOccurrencePtrInput
	// Describes an upgrade.
	Upgrade UpgradeOccurrencePtrInput
	// Details of a security vulnerability note.
	VulnerabilityDetails VulnerabilityDetailsPtrInput
}

func (OccurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*occurrenceArgs)(nil)).Elem()
}

type OccurrenceInput interface {
	pulumi.Input

	ToOccurrenceOutput() OccurrenceOutput
	ToOccurrenceOutputWithContext(ctx context.Context) OccurrenceOutput
}

func (*Occurrence) ElementType() reflect.Type {
	return reflect.TypeOf((**Occurrence)(nil)).Elem()
}

func (i *Occurrence) ToOccurrenceOutput() OccurrenceOutput {
	return i.ToOccurrenceOutputWithContext(context.Background())
}

func (i *Occurrence) ToOccurrenceOutputWithContext(ctx context.Context) OccurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OccurrenceOutput)
}

type OccurrenceOutput struct{ *pulumi.OutputState }

func (OccurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Occurrence)(nil)).Elem()
}

func (o OccurrenceOutput) ToOccurrenceOutput() OccurrenceOutput {
	return o
}

func (o OccurrenceOutput) ToOccurrenceOutputWithContext(ctx context.Context) OccurrenceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OccurrenceInput)(nil)).Elem(), &Occurrence{})
	pulumi.RegisterOutputType(OccurrenceOutput{})
}