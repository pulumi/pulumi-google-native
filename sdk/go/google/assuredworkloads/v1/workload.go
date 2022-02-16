// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates Assured Workload.
type Workload struct {
	pulumi.CustomResourceState

	// Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// Immutable. Compliance Regime associated with this workload.
	ComplianceRegime pulumi.StringOutput `pulumi:"complianceRegime"`
	// Immutable. The Workload creation timestamp.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
	EnableSovereignControls pulumi.BoolOutput `pulumi:"enableSovereignControls"`
	// Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update & Delete operations.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Represents the KAJ enrollment state of the given workload.
	KajEnrollmentState pulumi.StringOutput `pulumi:"kajEnrollmentState"`
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponseOutput `pulumi:"kmsSettings"`
	// Optional. Labels applied to the workload.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
	Name pulumi.StringOutput `pulumi:"name"`
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
	ProvisionedResourcesParent pulumi.StringOutput `pulumi:"provisionedResourcesParent"`
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponseArrayOutput `pulumi:"resourceSettings"`
	// The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
	Resources GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponseArrayOutput `pulumi:"resources"`
	// Represents the SAA enrollment response of the given workload. SAA enrollment response is queried during GetWorkload call. In failure cases, user friendly error message is shown in SAA details page.
	SaaEnrollmentResponse GoogleCloudAssuredworkloadsV1WorkloadSaaEnrollmentResponseResponseOutput `pulumi:"saaEnrollmentResponse"`
}

// NewWorkload registers a new resource with the given unique name, arguments, and options.
func NewWorkload(ctx *pulumi.Context,
	name string, args *WorkloadArgs, opts ...pulumi.ResourceOption) (*Workload, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComplianceRegime == nil {
		return nil, errors.New("invalid value for required argument 'ComplianceRegime'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource Workload
	err := ctx.RegisterResource("google-native:assuredworkloads/v1:Workload", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkload gets an existing Workload resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkload(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadState, opts ...pulumi.ResourceOption) (*Workload, error) {
	var resource Workload
	err := ctx.ReadResource("google-native:assuredworkloads/v1:Workload", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workload resources.
type workloadState struct {
}

type WorkloadState struct {
}

func (WorkloadState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadState)(nil)).Elem()
}

type workloadArgs struct {
	// Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccount *string `pulumi:"billingAccount"`
	// Immutable. Compliance Regime associated with this workload.
	ComplianceRegime WorkloadComplianceRegime `pulumi:"complianceRegime"`
	// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName string `pulumi:"displayName"`
	// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
	EnableSovereignControls *bool `pulumi:"enableSovereignControls"`
	// Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update & Delete operations.
	Etag       *string `pulumi:"etag"`
	ExternalId *string `pulumi:"externalId"`
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings *GoogleCloudAssuredworkloadsV1WorkloadKMSSettings `pulumi:"kmsSettings"`
	// Optional. Labels applied to the workload.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
	ProvisionedResourcesParent *string `pulumi:"provisionedResourcesParent"`
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings []GoogleCloudAssuredworkloadsV1WorkloadResourceSettings `pulumi:"resourceSettings"`
}

// The set of arguments for constructing a Workload resource.
type WorkloadArgs struct {
	// Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
	BillingAccount pulumi.StringPtrInput
	// Immutable. Compliance Regime associated with this workload.
	ComplianceRegime WorkloadComplianceRegimeInput
	// The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
	DisplayName pulumi.StringInput
	// Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
	EnableSovereignControls pulumi.BoolPtrInput
	// Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update & Delete operations.
	Etag       pulumi.StringPtrInput
	ExternalId pulumi.StringPtrInput
	// Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
	KmsSettings GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsPtrInput
	// Optional. Labels applied to the workload.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
	ProvisionedResourcesParent pulumi.StringPtrInput
	// Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
	ResourceSettings GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsArrayInput
}

func (WorkloadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadArgs)(nil)).Elem()
}

type WorkloadInput interface {
	pulumi.Input

	ToWorkloadOutput() WorkloadOutput
	ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput
}

func (*Workload) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (i *Workload) ToWorkloadOutput() WorkloadOutput {
	return i.ToWorkloadOutputWithContext(context.Background())
}

func (i *Workload) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkloadOutput)
}

type WorkloadOutput struct{ *pulumi.OutputState }

func (WorkloadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workload)(nil)).Elem()
}

func (o WorkloadOutput) ToWorkloadOutput() WorkloadOutput {
	return o
}

func (o WorkloadOutput) ToWorkloadOutputWithContext(ctx context.Context) WorkloadOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadInput)(nil)).Elem(), &Workload{})
	pulumi.RegisterOutputType(WorkloadOutput{})
}