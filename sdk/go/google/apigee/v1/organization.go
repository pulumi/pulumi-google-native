// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an Apigee organization. See [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
// Auto-naming is currently not supported for this resource.
type Organization struct {
	pulumi.CustomResourceState

	// Addon configurations of the Apigee organization.
	AddonsConfig GoogleCloudApigeeV1AddonsConfigResponseOutput `pulumi:"addonsConfig"`
	// Not used by Apigee.
	Attributes pulumi.StringArrayOutput `pulumi:"attributes"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
	AuthorizedNetwork pulumi.StringOutput `pulumi:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType pulumi.StringOutput `pulumi:"billingType"`
	// Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when [RuntimeType](#RuntimeType) is `CLOUD`.
	CaCertificate pulumi.StringOutput `pulumi:"caCertificate"`
	// Time that the Apigee organization was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Not used by Apigee.
	CustomerName pulumi.StringOutput `pulumi:"customerName"`
	// Description of the Apigee organization.
	Description pulumi.StringOutput `pulumi:"description"`
	// Display name for the Apigee organization. Unused, but reserved for future use.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// List of environments in the Apigee organization.
	Environments pulumi.StringArrayOutput `pulumi:"environments"`
	// Time that the Apigee organization is scheduled for deletion.
	ExpiresAt pulumi.StringOutput `pulumi:"expiresAt"`
	// Time that the Apigee organization was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Name of the Apigee organization.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for the Portals settings.
	PortalDisabled pulumi.BoolOutput `pulumi:"portalDisabled"`
	// Project ID associated with the Apigee organization.
	Project pulumi.StringOutput `pulumi:"project"`
	// Properties defined in the Apigee organization profile.
	Properties GoogleCloudApigeeV1PropertiesResponseOutput `pulumi:"properties"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
	RuntimeDatabaseEncryptionKeyName pulumi.StringOutput `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	RuntimeType pulumi.StringOutput `pulumi:"runtimeType"`
	// State of the organization. Values other than ACTIVE means the resource is not ready to use.
	State pulumi.StringOutput `pulumi:"state"`
	// Not used by Apigee.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOption) (*Organization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.RuntimeType == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeType'")
	}
	var resource Organization
	err := ctx.RegisterResource("google-native:apigee/v1:Organization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationState, opts ...pulumi.ResourceOption) (*Organization, error) {
	var resource Organization
	err := ctx.ReadResource("google-native:apigee/v1:Organization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Organization resources.
type organizationState struct {
}

type OrganizationState struct {
}

func (OrganizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationState)(nil)).Elem()
}

type organizationArgs struct {
	// Addon configurations of the Apigee organization.
	AddonsConfig *GoogleCloudApigeeV1AddonsConfig `pulumi:"addonsConfig"`
	// Not used by Apigee.
	Attributes []string `pulumi:"attributes"`
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
	AuthorizedNetwork *string `pulumi:"authorizedNetwork"`
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType *OrganizationBillingType `pulumi:"billingType"`
	// Not used by Apigee.
	CustomerName *string `pulumi:"customerName"`
	// Description of the Apigee organization.
	Description *string `pulumi:"description"`
	// Display name for the Apigee organization. Unused, but reserved for future use.
	DisplayName *string `pulumi:"displayName"`
	// Required. Name of the GCP project in which to associate the Apigee organization. Pass the information as a query parameter using the following structure in your request: `projects/`
	Parent string `pulumi:"parent"`
	// Configuration for the Portals settings.
	PortalDisabled *bool `pulumi:"portalDisabled"`
	// Properties defined in the Apigee organization profile.
	Properties *GoogleCloudApigeeV1Properties `pulumi:"properties"`
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
	RuntimeDatabaseEncryptionKeyName *string `pulumi:"runtimeDatabaseEncryptionKeyName"`
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	RuntimeType OrganizationRuntimeType `pulumi:"runtimeType"`
	// Not used by Apigee.
	Type *OrganizationType `pulumi:"type"`
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Addon configurations of the Apigee organization.
	AddonsConfig GoogleCloudApigeeV1AddonsConfigPtrInput
	// Not used by Apigee.
	Attributes pulumi.StringArrayInput
	// Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
	AuthorizedNetwork pulumi.StringPtrInput
	// Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
	BillingType OrganizationBillingTypePtrInput
	// Not used by Apigee.
	CustomerName pulumi.StringPtrInput
	// Description of the Apigee organization.
	Description pulumi.StringPtrInput
	// Display name for the Apigee organization. Unused, but reserved for future use.
	DisplayName pulumi.StringPtrInput
	// Required. Name of the GCP project in which to associate the Apigee organization. Pass the information as a query parameter using the following structure in your request: `projects/`
	Parent pulumi.StringInput
	// Configuration for the Portals settings.
	PortalDisabled pulumi.BoolPtrInput
	// Properties defined in the Apigee organization profile.
	Properties GoogleCloudApigeeV1PropertiesPtrInput
	// Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
	RuntimeDatabaseEncryptionKeyName pulumi.StringPtrInput
	// Runtime type of the Apigee organization based on the Apigee subscription purchased.
	RuntimeType OrganizationRuntimeTypeInput
	// Not used by Apigee.
	Type OrganizationTypePtrInput
}

func (OrganizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationArgs)(nil)).Elem()
}

type OrganizationInput interface {
	pulumi.Input

	ToOrganizationOutput() OrganizationOutput
	ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput
}

func (*Organization) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (i *Organization) ToOrganizationOutput() OrganizationOutput {
	return i.ToOrganizationOutputWithContext(context.Background())
}

func (i *Organization) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationOutput)
}

type OrganizationOutput struct{ *pulumi.OutputState }

func (OrganizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Organization)(nil)).Elem()
}

func (o OrganizationOutput) ToOrganizationOutput() OrganizationOutput {
	return o
}

func (o OrganizationOutput) ToOrganizationOutputWithContext(ctx context.Context) OrganizationOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationInput)(nil)).Elem(), &Organization{})
	pulumi.RegisterOutputType(OrganizationOutput{})
}