// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT)
type GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType string

const (
	// Unknown resource type.
	GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeResourceTypeUnspecified = GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("RESOURCE_TYPE_UNSPECIFIED")
	// Deprecated. Existing workloads will continue to support this, but new CreateWorkloadRequests should not specify this as an input value.
	GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeConsumerProject = GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("CONSUMER_PROJECT")
	// Consumer Folder.
	GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeConsumerFolder = GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("CONSUMER_FOLDER")
	// Consumer project containing encryption keys.
	GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeEncryptionKeysProject = GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("ENCRYPTION_KEYS_PROJECT")
	// Keyring resource that hosts encryption keys.
	GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeKeyring = GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("KEYRING")
)

func (GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType)(nil)).Elem()
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput {
	return pulumi.ToOutput(e).(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput)
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutputWithContext(ctx context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput)
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return e.ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(context.Background())
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(ctx context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType(e).ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutputWithContext(ctx).ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(ctx)
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput struct{ *pulumi.OutputState }

func (GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType)(nil)).Elem()
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput {
	return o
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutputWithContext(ctx context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput {
	return o
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return o.ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(context.Background())
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(ctx context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) *GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType {
		return &v
	}).(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput)
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput struct{ *pulumi.OutputState }

func (GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType)(nil)).Elem()
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return o
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(ctx context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return o
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput) Elem() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput {
	return o.ApplyT(func(v *GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType {
		if v != nil {
			return *v
		}
		var ret GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType
		return ret
	}).(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput)
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeInput is an input type that accepts GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeArgs and GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput values.
// You can construct a concrete instance of `GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeInput` via:
//
//	GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeArgs{...}
type GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeInput interface {
	pulumi.Input

	ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput
	ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutputWithContext(context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput
}

var googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrType = reflect.TypeOf((**GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType)(nil)).Elem()

type GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrInput interface {
	pulumi.Input

	ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput
	ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput
}

type googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtr string

func GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtr(v string) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrInput {
	return (*googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtr)(&v)
}

func (*googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtr) ElementType() reflect.Type {
	return googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrType
}

func (in *googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtr) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput() GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return pulumi.ToOutput(in).(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput)
}

func (in *googleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtr) ToGoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutputWithContext(ctx context.Context) GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput)
}

// Required. Immutable. Compliance Regime associated with this workload.
type WorkloadComplianceRegime string

const (
	// Unknown compliance regime.
	WorkloadComplianceRegimeComplianceRegimeUnspecified = WorkloadComplianceRegime("COMPLIANCE_REGIME_UNSPECIFIED")
	// Information protection as per DoD IL4 requirements.
	WorkloadComplianceRegimeIl4 = WorkloadComplianceRegime("IL4")
	// Criminal Justice Information Services (CJIS) Security policies.
	WorkloadComplianceRegimeCjis = WorkloadComplianceRegime("CJIS")
	// FedRAMP High data protection controls
	WorkloadComplianceRegimeFedrampHigh = WorkloadComplianceRegime("FEDRAMP_HIGH")
	// FedRAMP Moderate data protection controls
	WorkloadComplianceRegimeFedrampModerate = WorkloadComplianceRegime("FEDRAMP_MODERATE")
	// Assured Workloads For US Regions data protection controls
	WorkloadComplianceRegimeUsRegionalAccess = WorkloadComplianceRegime("US_REGIONAL_ACCESS")
	// Health Insurance Portability and Accountability Act controls
	WorkloadComplianceRegimeHipaa = WorkloadComplianceRegime("HIPAA")
	// Health Information Trust Alliance controls
	WorkloadComplianceRegimeHitrust = WorkloadComplianceRegime("HITRUST")
	// Assured Workloads For EU Regions and Support controls
	WorkloadComplianceRegimeEuRegionsAndSupport = WorkloadComplianceRegime("EU_REGIONS_AND_SUPPORT")
	// Assured Workloads For Canada Regions and Support controls
	WorkloadComplianceRegimeCaRegionsAndSupport = WorkloadComplianceRegime("CA_REGIONS_AND_SUPPORT")
	// International Traffic in Arms Regulations
	WorkloadComplianceRegimeItar = WorkloadComplianceRegime("ITAR")
	// Assured Workloads for Australia Regions and Support controls Available for public preview consumption. Don't create production workloads.
	WorkloadComplianceRegimeAuRegionsAndUsSupport = WorkloadComplianceRegime("AU_REGIONS_AND_US_SUPPORT")
	// Assured Workloads for Partners;
	WorkloadComplianceRegimeAssuredWorkloadsForPartners = WorkloadComplianceRegime("ASSURED_WORKLOADS_FOR_PARTNERS")
	// Assured Workloads for Israel
	WorkloadComplianceRegimeIsrRegions = WorkloadComplianceRegime("ISR_REGIONS")
)

func (WorkloadComplianceRegime) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadComplianceRegime)(nil)).Elem()
}

func (e WorkloadComplianceRegime) ToWorkloadComplianceRegimeOutput() WorkloadComplianceRegimeOutput {
	return pulumi.ToOutput(e).(WorkloadComplianceRegimeOutput)
}

func (e WorkloadComplianceRegime) ToWorkloadComplianceRegimeOutputWithContext(ctx context.Context) WorkloadComplianceRegimeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkloadComplianceRegimeOutput)
}

func (e WorkloadComplianceRegime) ToWorkloadComplianceRegimePtrOutput() WorkloadComplianceRegimePtrOutput {
	return e.ToWorkloadComplianceRegimePtrOutputWithContext(context.Background())
}

func (e WorkloadComplianceRegime) ToWorkloadComplianceRegimePtrOutputWithContext(ctx context.Context) WorkloadComplianceRegimePtrOutput {
	return WorkloadComplianceRegime(e).ToWorkloadComplianceRegimeOutputWithContext(ctx).ToWorkloadComplianceRegimePtrOutputWithContext(ctx)
}

func (e WorkloadComplianceRegime) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadComplianceRegime) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadComplianceRegime) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkloadComplianceRegime) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkloadComplianceRegimeOutput struct{ *pulumi.OutputState }

func (WorkloadComplianceRegimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadComplianceRegime)(nil)).Elem()
}

func (o WorkloadComplianceRegimeOutput) ToWorkloadComplianceRegimeOutput() WorkloadComplianceRegimeOutput {
	return o
}

func (o WorkloadComplianceRegimeOutput) ToWorkloadComplianceRegimeOutputWithContext(ctx context.Context) WorkloadComplianceRegimeOutput {
	return o
}

func (o WorkloadComplianceRegimeOutput) ToWorkloadComplianceRegimePtrOutput() WorkloadComplianceRegimePtrOutput {
	return o.ToWorkloadComplianceRegimePtrOutputWithContext(context.Background())
}

func (o WorkloadComplianceRegimeOutput) ToWorkloadComplianceRegimePtrOutputWithContext(ctx context.Context) WorkloadComplianceRegimePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadComplianceRegime) *WorkloadComplianceRegime {
		return &v
	}).(WorkloadComplianceRegimePtrOutput)
}

func (o WorkloadComplianceRegimeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkloadComplianceRegimeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadComplianceRegime) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkloadComplianceRegimeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadComplianceRegimeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadComplianceRegime) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkloadComplianceRegimePtrOutput struct{ *pulumi.OutputState }

func (WorkloadComplianceRegimePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadComplianceRegime)(nil)).Elem()
}

func (o WorkloadComplianceRegimePtrOutput) ToWorkloadComplianceRegimePtrOutput() WorkloadComplianceRegimePtrOutput {
	return o
}

func (o WorkloadComplianceRegimePtrOutput) ToWorkloadComplianceRegimePtrOutputWithContext(ctx context.Context) WorkloadComplianceRegimePtrOutput {
	return o
}

func (o WorkloadComplianceRegimePtrOutput) Elem() WorkloadComplianceRegimeOutput {
	return o.ApplyT(func(v *WorkloadComplianceRegime) WorkloadComplianceRegime {
		if v != nil {
			return *v
		}
		var ret WorkloadComplianceRegime
		return ret
	}).(WorkloadComplianceRegimeOutput)
}

func (o WorkloadComplianceRegimePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadComplianceRegimePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkloadComplianceRegime) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkloadComplianceRegimeInput is an input type that accepts WorkloadComplianceRegimeArgs and WorkloadComplianceRegimeOutput values.
// You can construct a concrete instance of `WorkloadComplianceRegimeInput` via:
//
//	WorkloadComplianceRegimeArgs{...}
type WorkloadComplianceRegimeInput interface {
	pulumi.Input

	ToWorkloadComplianceRegimeOutput() WorkloadComplianceRegimeOutput
	ToWorkloadComplianceRegimeOutputWithContext(context.Context) WorkloadComplianceRegimeOutput
}

var workloadComplianceRegimePtrType = reflect.TypeOf((**WorkloadComplianceRegime)(nil)).Elem()

type WorkloadComplianceRegimePtrInput interface {
	pulumi.Input

	ToWorkloadComplianceRegimePtrOutput() WorkloadComplianceRegimePtrOutput
	ToWorkloadComplianceRegimePtrOutputWithContext(context.Context) WorkloadComplianceRegimePtrOutput
}

type workloadComplianceRegimePtr string

func WorkloadComplianceRegimePtr(v string) WorkloadComplianceRegimePtrInput {
	return (*workloadComplianceRegimePtr)(&v)
}

func (*workloadComplianceRegimePtr) ElementType() reflect.Type {
	return workloadComplianceRegimePtrType
}

func (in *workloadComplianceRegimePtr) ToWorkloadComplianceRegimePtrOutput() WorkloadComplianceRegimePtrOutput {
	return pulumi.ToOutput(in).(WorkloadComplianceRegimePtrOutput)
}

func (in *workloadComplianceRegimePtr) ToWorkloadComplianceRegimePtrOutputWithContext(ctx context.Context) WorkloadComplianceRegimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkloadComplianceRegimePtrOutput)
}

// Optional. Partner regime associated with this workload.
type WorkloadPartner string

const (
	WorkloadPartnerPartnerUnspecified = WorkloadPartner("PARTNER_UNSPECIFIED")
	// Enum representing S3NS partner.
	WorkloadPartnerLocalControlsByS3ns = WorkloadPartner("LOCAL_CONTROLS_BY_S3NS")
	// Enum representing T_SYSTEM partner.
	WorkloadPartnerSovereignControlsByTSystems = WorkloadPartner("SOVEREIGN_CONTROLS_BY_T_SYSTEMS")
)

func (WorkloadPartner) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadPartner)(nil)).Elem()
}

func (e WorkloadPartner) ToWorkloadPartnerOutput() WorkloadPartnerOutput {
	return pulumi.ToOutput(e).(WorkloadPartnerOutput)
}

func (e WorkloadPartner) ToWorkloadPartnerOutputWithContext(ctx context.Context) WorkloadPartnerOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WorkloadPartnerOutput)
}

func (e WorkloadPartner) ToWorkloadPartnerPtrOutput() WorkloadPartnerPtrOutput {
	return e.ToWorkloadPartnerPtrOutputWithContext(context.Background())
}

func (e WorkloadPartner) ToWorkloadPartnerPtrOutputWithContext(ctx context.Context) WorkloadPartnerPtrOutput {
	return WorkloadPartner(e).ToWorkloadPartnerOutputWithContext(ctx).ToWorkloadPartnerPtrOutputWithContext(ctx)
}

func (e WorkloadPartner) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadPartner) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WorkloadPartner) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WorkloadPartner) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WorkloadPartnerOutput struct{ *pulumi.OutputState }

func (WorkloadPartnerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkloadPartner)(nil)).Elem()
}

func (o WorkloadPartnerOutput) ToWorkloadPartnerOutput() WorkloadPartnerOutput {
	return o
}

func (o WorkloadPartnerOutput) ToWorkloadPartnerOutputWithContext(ctx context.Context) WorkloadPartnerOutput {
	return o
}

func (o WorkloadPartnerOutput) ToWorkloadPartnerPtrOutput() WorkloadPartnerPtrOutput {
	return o.ToWorkloadPartnerPtrOutputWithContext(context.Background())
}

func (o WorkloadPartnerOutput) ToWorkloadPartnerPtrOutputWithContext(ctx context.Context) WorkloadPartnerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WorkloadPartner) *WorkloadPartner {
		return &v
	}).(WorkloadPartnerPtrOutput)
}

func (o WorkloadPartnerOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WorkloadPartnerOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadPartner) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WorkloadPartnerOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadPartnerOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WorkloadPartner) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WorkloadPartnerPtrOutput struct{ *pulumi.OutputState }

func (WorkloadPartnerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkloadPartner)(nil)).Elem()
}

func (o WorkloadPartnerPtrOutput) ToWorkloadPartnerPtrOutput() WorkloadPartnerPtrOutput {
	return o
}

func (o WorkloadPartnerPtrOutput) ToWorkloadPartnerPtrOutputWithContext(ctx context.Context) WorkloadPartnerPtrOutput {
	return o
}

func (o WorkloadPartnerPtrOutput) Elem() WorkloadPartnerOutput {
	return o.ApplyT(func(v *WorkloadPartner) WorkloadPartner {
		if v != nil {
			return *v
		}
		var ret WorkloadPartner
		return ret
	}).(WorkloadPartnerOutput)
}

func (o WorkloadPartnerPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WorkloadPartnerPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WorkloadPartner) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// WorkloadPartnerInput is an input type that accepts WorkloadPartnerArgs and WorkloadPartnerOutput values.
// You can construct a concrete instance of `WorkloadPartnerInput` via:
//
//	WorkloadPartnerArgs{...}
type WorkloadPartnerInput interface {
	pulumi.Input

	ToWorkloadPartnerOutput() WorkloadPartnerOutput
	ToWorkloadPartnerOutputWithContext(context.Context) WorkloadPartnerOutput
}

var workloadPartnerPtrType = reflect.TypeOf((**WorkloadPartner)(nil)).Elem()

type WorkloadPartnerPtrInput interface {
	pulumi.Input

	ToWorkloadPartnerPtrOutput() WorkloadPartnerPtrOutput
	ToWorkloadPartnerPtrOutputWithContext(context.Context) WorkloadPartnerPtrOutput
}

type workloadPartnerPtr string

func WorkloadPartnerPtr(v string) WorkloadPartnerPtrInput {
	return (*workloadPartnerPtr)(&v)
}

func (*workloadPartnerPtr) ElementType() reflect.Type {
	return workloadPartnerPtrType
}

func (in *workloadPartnerPtr) ToWorkloadPartnerPtrOutput() WorkloadPartnerPtrOutput {
	return pulumi.ToOutput(in).(WorkloadPartnerPtrOutput)
}

func (in *workloadPartnerPtr) ToWorkloadPartnerPtrOutputWithContext(ctx context.Context) WorkloadPartnerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WorkloadPartnerPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeInput)(nil)).Elem(), GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("RESOURCE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrInput)(nil)).Elem(), GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType("RESOURCE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadComplianceRegimeInput)(nil)).Elem(), WorkloadComplianceRegime("COMPLIANCE_REGIME_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadComplianceRegimePtrInput)(nil)).Elem(), WorkloadComplianceRegime("COMPLIANCE_REGIME_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadPartnerInput)(nil)).Elem(), WorkloadPartner("PARTNER_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*WorkloadPartnerPtrInput)(nil)).Elem(), WorkloadPartner("PARTNER_UNSPECIFIED"))
	pulumi.RegisterOutputType(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypeOutput{})
	pulumi.RegisterOutputType(GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceTypePtrOutput{})
	pulumi.RegisterOutputType(WorkloadComplianceRegimeOutput{})
	pulumi.RegisterOutputType(WorkloadComplianceRegimePtrOutput{})
	pulumi.RegisterOutputType(WorkloadPartnerOutput{})
	pulumi.RegisterOutputType(WorkloadPartnerPtrOutput{})
}