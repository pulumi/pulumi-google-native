// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

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
	case "google-native:analyticshub/v1:DataExchange":
		r = &DataExchange{}
	case "google-native:analyticshub/v1:DataExchangeIamBinding":
		r = &DataExchangeIamBinding{}
	case "google-native:analyticshub/v1:DataExchangeIamMember":
		r = &DataExchangeIamMember{}
	case "google-native:analyticshub/v1:DataExchangeIamPolicy":
		r = &DataExchangeIamPolicy{}
	case "google-native:analyticshub/v1:DataExchangeListingIamBinding":
		r = &DataExchangeListingIamBinding{}
	case "google-native:analyticshub/v1:DataExchangeListingIamMember":
		r = &DataExchangeListingIamMember{}
	case "google-native:analyticshub/v1:DataExchangeListingIamPolicy":
		r = &DataExchangeListingIamPolicy{}
	case "google-native:analyticshub/v1:Listing":
		r = &Listing{}
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
		"analyticshub/v1",
		&module{version},
	)
}