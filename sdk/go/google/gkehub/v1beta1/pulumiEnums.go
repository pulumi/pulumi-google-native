// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The log type that this config enables.
type AuditLogConfigLogType string

const (
	// Default case. Should never be this.
	AuditLogConfigLogTypeLogTypeUnspecified = AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED")
	// Admin reads. Example: CloudIAM getIamPolicy
	AuditLogConfigLogTypeAdminRead = AuditLogConfigLogType("ADMIN_READ")
	// Data writes. Example: CloudSQL Users create
	AuditLogConfigLogTypeDataWrite = AuditLogConfigLogType("DATA_WRITE")
	// Data reads. Example: CloudSQL Users list
	AuditLogConfigLogTypeDataRead = AuditLogConfigLogType("DATA_READ")
)

func (AuditLogConfigLogType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return pulumi.ToOutput(e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuditLogConfigLogTypeOutput)
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return e.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return AuditLogConfigLogType(e).ToAuditLogConfigLogTypeOutputWithContext(ctx).ToAuditLogConfigLogTypePtrOutputWithContext(ctx)
}

func (e AuditLogConfigLogType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuditLogConfigLogType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuditLogConfigLogType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuditLogConfigLogTypeOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypeOutputWithContext(ctx context.Context) AuditLogConfigLogTypeOutput {
	return o
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o.ToAuditLogConfigLogTypePtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuditLogConfigLogType) *AuditLogConfigLogType {
		return &v
	}).(AuditLogConfigLogTypePtrOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuditLogConfigLogType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuditLogConfigLogTypePtrOutput struct{ *pulumi.OutputState }

func (AuditLogConfigLogTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return o
}

func (o AuditLogConfigLogTypePtrOutput) Elem() AuditLogConfigLogTypeOutput {
	return o.ApplyT(func(v *AuditLogConfigLogType) AuditLogConfigLogType {
		if v != nil {
			return *v
		}
		var ret AuditLogConfigLogType
		return ret
	}).(AuditLogConfigLogTypeOutput)
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuditLogConfigLogTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuditLogConfigLogType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// AuditLogConfigLogTypeInput is an input type that accepts AuditLogConfigLogTypeArgs and AuditLogConfigLogTypeOutput values.
// You can construct a concrete instance of `AuditLogConfigLogTypeInput` via:
//
//          AuditLogConfigLogTypeArgs{...}
type AuditLogConfigLogTypeInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypeOutput() AuditLogConfigLogTypeOutput
	ToAuditLogConfigLogTypeOutputWithContext(context.Context) AuditLogConfigLogTypeOutput
}

var auditLogConfigLogTypePtrType = reflect.TypeOf((**AuditLogConfigLogType)(nil)).Elem()

type AuditLogConfigLogTypePtrInput interface {
	pulumi.Input

	ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput
	ToAuditLogConfigLogTypePtrOutputWithContext(context.Context) AuditLogConfigLogTypePtrOutput
}

type auditLogConfigLogTypePtr string

func AuditLogConfigLogTypePtr(v string) AuditLogConfigLogTypePtrInput {
	return (*auditLogConfigLogTypePtr)(&v)
}

func (*auditLogConfigLogTypePtr) ElementType() reflect.Type {
	return auditLogConfigLogTypePtrType
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutput() AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutput(in).(AuditLogConfigLogTypePtrOutput)
}

func (in *auditLogConfigLogTypePtr) ToAuditLogConfigLogTypePtrOutputWithContext(ctx context.Context) AuditLogConfigLogTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuditLogConfigLogTypePtrOutput)
}

// Optional. The infrastructure type this Membership is running on.
type MembershipInfrastructureType string

const (
	// No type was specified. Some Hub functionality may require a type be specified, and will not support Memberships with this value.
	MembershipInfrastructureTypeInfrastructureTypeUnspecified = MembershipInfrastructureType("INFRASTRUCTURE_TYPE_UNSPECIFIED")
	// Private infrastructure that is owned or operated by customer. This includes GKE distributions such as GKE-OnPrem and GKE-OnBareMetal.
	MembershipInfrastructureTypeOnPrem = MembershipInfrastructureType("ON_PREM")
	// Public cloud infrastructure.
	MembershipInfrastructureTypeMultiCloud = MembershipInfrastructureType("MULTI_CLOUD")
)

func (MembershipInfrastructureType) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipInfrastructureType)(nil)).Elem()
}

func (e MembershipInfrastructureType) ToMembershipInfrastructureTypeOutput() MembershipInfrastructureTypeOutput {
	return pulumi.ToOutput(e).(MembershipInfrastructureTypeOutput)
}

func (e MembershipInfrastructureType) ToMembershipInfrastructureTypeOutputWithContext(ctx context.Context) MembershipInfrastructureTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MembershipInfrastructureTypeOutput)
}

func (e MembershipInfrastructureType) ToMembershipInfrastructureTypePtrOutput() MembershipInfrastructureTypePtrOutput {
	return e.ToMembershipInfrastructureTypePtrOutputWithContext(context.Background())
}

func (e MembershipInfrastructureType) ToMembershipInfrastructureTypePtrOutputWithContext(ctx context.Context) MembershipInfrastructureTypePtrOutput {
	return MembershipInfrastructureType(e).ToMembershipInfrastructureTypeOutputWithContext(ctx).ToMembershipInfrastructureTypePtrOutputWithContext(ctx)
}

func (e MembershipInfrastructureType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MembershipInfrastructureType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MembershipInfrastructureType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MembershipInfrastructureType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MembershipInfrastructureTypeOutput struct{ *pulumi.OutputState }

func (MembershipInfrastructureTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MembershipInfrastructureType)(nil)).Elem()
}

func (o MembershipInfrastructureTypeOutput) ToMembershipInfrastructureTypeOutput() MembershipInfrastructureTypeOutput {
	return o
}

func (o MembershipInfrastructureTypeOutput) ToMembershipInfrastructureTypeOutputWithContext(ctx context.Context) MembershipInfrastructureTypeOutput {
	return o
}

func (o MembershipInfrastructureTypeOutput) ToMembershipInfrastructureTypePtrOutput() MembershipInfrastructureTypePtrOutput {
	return o.ToMembershipInfrastructureTypePtrOutputWithContext(context.Background())
}

func (o MembershipInfrastructureTypeOutput) ToMembershipInfrastructureTypePtrOutputWithContext(ctx context.Context) MembershipInfrastructureTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MembershipInfrastructureType) *MembershipInfrastructureType {
		return &v
	}).(MembershipInfrastructureTypePtrOutput)
}

func (o MembershipInfrastructureTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MembershipInfrastructureTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MembershipInfrastructureType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MembershipInfrastructureTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MembershipInfrastructureTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MembershipInfrastructureType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MembershipInfrastructureTypePtrOutput struct{ *pulumi.OutputState }

func (MembershipInfrastructureTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MembershipInfrastructureType)(nil)).Elem()
}

func (o MembershipInfrastructureTypePtrOutput) ToMembershipInfrastructureTypePtrOutput() MembershipInfrastructureTypePtrOutput {
	return o
}

func (o MembershipInfrastructureTypePtrOutput) ToMembershipInfrastructureTypePtrOutputWithContext(ctx context.Context) MembershipInfrastructureTypePtrOutput {
	return o
}

func (o MembershipInfrastructureTypePtrOutput) Elem() MembershipInfrastructureTypeOutput {
	return o.ApplyT(func(v *MembershipInfrastructureType) MembershipInfrastructureType {
		if v != nil {
			return *v
		}
		var ret MembershipInfrastructureType
		return ret
	}).(MembershipInfrastructureTypeOutput)
}

func (o MembershipInfrastructureTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MembershipInfrastructureTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MembershipInfrastructureType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MembershipInfrastructureTypeInput is an input type that accepts MembershipInfrastructureTypeArgs and MembershipInfrastructureTypeOutput values.
// You can construct a concrete instance of `MembershipInfrastructureTypeInput` via:
//
//          MembershipInfrastructureTypeArgs{...}
type MembershipInfrastructureTypeInput interface {
	pulumi.Input

	ToMembershipInfrastructureTypeOutput() MembershipInfrastructureTypeOutput
	ToMembershipInfrastructureTypeOutputWithContext(context.Context) MembershipInfrastructureTypeOutput
}

var membershipInfrastructureTypePtrType = reflect.TypeOf((**MembershipInfrastructureType)(nil)).Elem()

type MembershipInfrastructureTypePtrInput interface {
	pulumi.Input

	ToMembershipInfrastructureTypePtrOutput() MembershipInfrastructureTypePtrOutput
	ToMembershipInfrastructureTypePtrOutputWithContext(context.Context) MembershipInfrastructureTypePtrOutput
}

type membershipInfrastructureTypePtr string

func MembershipInfrastructureTypePtr(v string) MembershipInfrastructureTypePtrInput {
	return (*membershipInfrastructureTypePtr)(&v)
}

func (*membershipInfrastructureTypePtr) ElementType() reflect.Type {
	return membershipInfrastructureTypePtrType
}

func (in *membershipInfrastructureTypePtr) ToMembershipInfrastructureTypePtrOutput() MembershipInfrastructureTypePtrOutput {
	return pulumi.ToOutput(in).(MembershipInfrastructureTypePtrOutput)
}

func (in *membershipInfrastructureTypePtr) ToMembershipInfrastructureTypePtrOutputWithContext(ctx context.Context) MembershipInfrastructureTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MembershipInfrastructureTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypeInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuditLogConfigLogTypePtrInput)(nil)).Elem(), AuditLogConfigLogType("LOG_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipInfrastructureTypeInput)(nil)).Elem(), MembershipInfrastructureType("INFRASTRUCTURE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*MembershipInfrastructureTypePtrInput)(nil)).Elem(), MembershipInfrastructureType("INFRASTRUCTURE_TYPE_UNSPECIFIED"))
	pulumi.RegisterOutputType(AuditLogConfigLogTypeOutput{})
	pulumi.RegisterOutputType(AuditLogConfigLogTypePtrOutput{})
	pulumi.RegisterOutputType(MembershipInfrastructureTypeOutput{})
	pulumi.RegisterOutputType(MembershipInfrastructureTypePtrOutput{})
}