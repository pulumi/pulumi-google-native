// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-native/sdk/go/google"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "google-native:healthcare/v1beta1:Annotation":
		r = &Annotation{}
	case "google-native:healthcare/v1beta1:AnnotationStore":
		r = &AnnotationStore{}
	case "google-native:healthcare/v1beta1:AttributeDefinition":
		r = &AttributeDefinition{}
	case "google-native:healthcare/v1beta1:Consent":
		r = &Consent{}
	case "google-native:healthcare/v1beta1:ConsentArtifact":
		r = &ConsentArtifact{}
	case "google-native:healthcare/v1beta1:ConsentStore":
		r = &ConsentStore{}
	case "google-native:healthcare/v1beta1:Dataset":
		r = &Dataset{}
	case "google-native:healthcare/v1beta1:DatasetAnnotationStoreIamPolicy":
		r = &DatasetAnnotationStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetConsentStoreIamPolicy":
		r = &DatasetConsentStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetDicomStoreIamPolicy":
		r = &DatasetDicomStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetFhirStoreIamPolicy":
		r = &DatasetFhirStoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetHl7V2StoreIamPolicy":
		r = &DatasetHl7V2StoreIamPolicy{}
	case "google-native:healthcare/v1beta1:DatasetIamPolicy":
		r = &DatasetIamPolicy{}
	case "google-native:healthcare/v1beta1:DicomStore":
		r = &DicomStore{}
	case "google-native:healthcare/v1beta1:FhirStore":
		r = &FhirStore{}
	case "google-native:healthcare/v1beta1:Hl7V2Store":
		r = &Hl7V2Store{}
	case "google-native:healthcare/v1beta1:Message":
		r = &Message{}
	case "google-native:healthcare/v1beta1:UserDataMapping":
		r = &UserDataMapping{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := google.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"healthcare/v1beta1",
		&module{version},
	)
}