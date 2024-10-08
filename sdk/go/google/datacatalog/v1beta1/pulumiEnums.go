// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of the entry. Only used for Entries with types in the EntryType enum.
type EntryType string

const (
	// Default unknown type.
	EntryTypeEntryTypeUnspecified = EntryType("ENTRY_TYPE_UNSPECIFIED")
	// Output only. The type of entry that has a GoogleSQL schema, including logical views.
	EntryTypeTable = EntryType("TABLE")
	// Output only. The type of models. https://cloud.google.com/bigquery-ml/docs/bigqueryml-intro
	EntryTypeModel = EntryType("MODEL")
	// Output only. An entry type which is used for streaming entries. Example: Pub/Sub topic.
	EntryTypeDataStream = EntryType("DATA_STREAM")
	// An entry type which is a set of files or objects. Example: Cloud Storage fileset.
	EntryTypeFileset = EntryType("FILESET")
)

func (EntryType) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryType)(nil)).Elem()
}

func (e EntryType) ToEntryTypeOutput() EntryTypeOutput {
	return pulumi.ToOutput(e).(EntryTypeOutput)
}

func (e EntryType) ToEntryTypeOutputWithContext(ctx context.Context) EntryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EntryTypeOutput)
}

func (e EntryType) ToEntryTypePtrOutput() EntryTypePtrOutput {
	return e.ToEntryTypePtrOutputWithContext(context.Background())
}

func (e EntryType) ToEntryTypePtrOutputWithContext(ctx context.Context) EntryTypePtrOutput {
	return EntryType(e).ToEntryTypeOutputWithContext(ctx).ToEntryTypePtrOutputWithContext(ctx)
}

func (e EntryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EntryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EntryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EntryTypeOutput struct{ *pulumi.OutputState }

func (EntryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntryType)(nil)).Elem()
}

func (o EntryTypeOutput) ToEntryTypeOutput() EntryTypeOutput {
	return o
}

func (o EntryTypeOutput) ToEntryTypeOutputWithContext(ctx context.Context) EntryTypeOutput {
	return o
}

func (o EntryTypeOutput) ToEntryTypePtrOutput() EntryTypePtrOutput {
	return o.ToEntryTypePtrOutputWithContext(context.Background())
}

func (o EntryTypeOutput) ToEntryTypePtrOutputWithContext(ctx context.Context) EntryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntryType) *EntryType {
		return &v
	}).(EntryTypePtrOutput)
}

func (o EntryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EntryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EntryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EntryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EntryTypePtrOutput struct{ *pulumi.OutputState }

func (EntryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntryType)(nil)).Elem()
}

func (o EntryTypePtrOutput) ToEntryTypePtrOutput() EntryTypePtrOutput {
	return o
}

func (o EntryTypePtrOutput) ToEntryTypePtrOutputWithContext(ctx context.Context) EntryTypePtrOutput {
	return o
}

func (o EntryTypePtrOutput) Elem() EntryTypeOutput {
	return o.ApplyT(func(v *EntryType) EntryType {
		if v != nil {
			return *v
		}
		var ret EntryType
		return ret
	}).(EntryTypeOutput)
}

func (o EntryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EntryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EntryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EntryTypeInput is an input type that accepts values of the EntryType enum
// A concrete instance of `EntryTypeInput` can be one of the following:
//
//	EntryTypeEntryTypeUnspecified
//	EntryTypeTable
//	EntryTypeModel
//	EntryTypeDataStream
//	EntryTypeFileset
type EntryTypeInput interface {
	pulumi.Input

	ToEntryTypeOutput() EntryTypeOutput
	ToEntryTypeOutputWithContext(context.Context) EntryTypeOutput
}

var entryTypePtrType = reflect.TypeOf((**EntryType)(nil)).Elem()

type EntryTypePtrInput interface {
	pulumi.Input

	ToEntryTypePtrOutput() EntryTypePtrOutput
	ToEntryTypePtrOutputWithContext(context.Context) EntryTypePtrOutput
}

type entryTypePtr string

func EntryTypePtr(v string) EntryTypePtrInput {
	return (*entryTypePtr)(&v)
}

func (*entryTypePtr) ElementType() reflect.Type {
	return entryTypePtrType
}

func (in *entryTypePtr) ToEntryTypePtrOutput() EntryTypePtrOutput {
	return pulumi.ToOutput(in).(EntryTypePtrOutput)
}

func (in *entryTypePtr) ToEntryTypePtrOutputWithContext(ctx context.Context) EntryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EntryTypePtrOutput)
}

// Represents primitive types - string, bool etc.
type GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType string

const (
	// This is the default invalid value for a type.
	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePrimitiveTypeUnspecified = GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("PRIMITIVE_TYPE_UNSPECIFIED")
	// A double precision number.
	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeDouble = GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("DOUBLE")
	// An UTF-8 string.
	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeString = GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("STRING")
	// A boolean value.
	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeBool = GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("BOOL")
	// A timestamp.
	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeTimestamp = GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("TIMESTAMP")
)

func (GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType)(nil)).Elem()
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput {
	return pulumi.ToOutput(e).(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput)
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutputWithContext(ctx context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput)
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return e.ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(context.Background())
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(ctx context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType(e).ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutputWithContext(ctx).ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(ctx)
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput struct{ *pulumi.OutputState }

func (GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType)(nil)).Elem()
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput {
	return o
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutputWithContext(ctx context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput {
	return o
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return o.ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(context.Background())
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(ctx context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) *GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType {
		return &v
	}).(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput)
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput struct{ *pulumi.OutputState }

func (GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType)(nil)).Elem()
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return o
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(ctx context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return o
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput) Elem() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput {
	return o.ApplyT(func(v *GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType {
		if v != nil {
			return *v
		}
		var ret GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType
		return ret
	}).(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput)
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeInput is an input type that accepts values of the GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType enum
// A concrete instance of `GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeInput` can be one of the following:
//
//	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePrimitiveTypeUnspecified
//	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeDouble
//	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeString
//	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeBool
//	GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeTimestamp
type GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeInput interface {
	pulumi.Input

	ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput
	ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutputWithContext(context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput
}

var googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrType = reflect.TypeOf((**GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType)(nil)).Elem()

type GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrInput interface {
	pulumi.Input

	ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput
	ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput
}

type googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtr string

func GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtr(v string) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrInput {
	return (*googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtr)(&v)
}

func (*googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtr) ElementType() reflect.Type {
	return googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrType
}

func (in *googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtr) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput() GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return pulumi.ToOutput(in).(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput)
}

func (in *googleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtr) ToGoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutputWithContext(ctx context.Context) GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput)
}

type TaxonomyActivatedPolicyTypesItem string

const (
	// Unspecified policy type.
	TaxonomyActivatedPolicyTypesItemPolicyTypeUnspecified = TaxonomyActivatedPolicyTypesItem("POLICY_TYPE_UNSPECIFIED")
	// Fine grained access control policy, which enables access control on tagged resources.
	TaxonomyActivatedPolicyTypesItemFineGrainedAccessControl = TaxonomyActivatedPolicyTypesItem("FINE_GRAINED_ACCESS_CONTROL")
)

func (TaxonomyActivatedPolicyTypesItem) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyActivatedPolicyTypesItem)(nil)).Elem()
}

func (e TaxonomyActivatedPolicyTypesItem) ToTaxonomyActivatedPolicyTypesItemOutput() TaxonomyActivatedPolicyTypesItemOutput {
	return pulumi.ToOutput(e).(TaxonomyActivatedPolicyTypesItemOutput)
}

func (e TaxonomyActivatedPolicyTypesItem) ToTaxonomyActivatedPolicyTypesItemOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TaxonomyActivatedPolicyTypesItemOutput)
}

func (e TaxonomyActivatedPolicyTypesItem) ToTaxonomyActivatedPolicyTypesItemPtrOutput() TaxonomyActivatedPolicyTypesItemPtrOutput {
	return e.ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(context.Background())
}

func (e TaxonomyActivatedPolicyTypesItem) ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemPtrOutput {
	return TaxonomyActivatedPolicyTypesItem(e).ToTaxonomyActivatedPolicyTypesItemOutputWithContext(ctx).ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(ctx)
}

func (e TaxonomyActivatedPolicyTypesItem) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaxonomyActivatedPolicyTypesItem) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TaxonomyActivatedPolicyTypesItem) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TaxonomyActivatedPolicyTypesItem) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TaxonomyActivatedPolicyTypesItemOutput struct{ *pulumi.OutputState }

func (TaxonomyActivatedPolicyTypesItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaxonomyActivatedPolicyTypesItem)(nil)).Elem()
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToTaxonomyActivatedPolicyTypesItemOutput() TaxonomyActivatedPolicyTypesItemOutput {
	return o
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToTaxonomyActivatedPolicyTypesItemOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemOutput {
	return o
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToTaxonomyActivatedPolicyTypesItemPtrOutput() TaxonomyActivatedPolicyTypesItemPtrOutput {
	return o.ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(context.Background())
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaxonomyActivatedPolicyTypesItem) *TaxonomyActivatedPolicyTypesItem {
		return &v
	}).(TaxonomyActivatedPolicyTypesItemPtrOutput)
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaxonomyActivatedPolicyTypesItem) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaxonomyActivatedPolicyTypesItemOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TaxonomyActivatedPolicyTypesItem) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TaxonomyActivatedPolicyTypesItemPtrOutput struct{ *pulumi.OutputState }

func (TaxonomyActivatedPolicyTypesItemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaxonomyActivatedPolicyTypesItem)(nil)).Elem()
}

func (o TaxonomyActivatedPolicyTypesItemPtrOutput) ToTaxonomyActivatedPolicyTypesItemPtrOutput() TaxonomyActivatedPolicyTypesItemPtrOutput {
	return o
}

func (o TaxonomyActivatedPolicyTypesItemPtrOutput) ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemPtrOutput {
	return o
}

func (o TaxonomyActivatedPolicyTypesItemPtrOutput) Elem() TaxonomyActivatedPolicyTypesItemOutput {
	return o.ApplyT(func(v *TaxonomyActivatedPolicyTypesItem) TaxonomyActivatedPolicyTypesItem {
		if v != nil {
			return *v
		}
		var ret TaxonomyActivatedPolicyTypesItem
		return ret
	}).(TaxonomyActivatedPolicyTypesItemOutput)
}

func (o TaxonomyActivatedPolicyTypesItemPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TaxonomyActivatedPolicyTypesItemPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TaxonomyActivatedPolicyTypesItem) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TaxonomyActivatedPolicyTypesItemInput is an input type that accepts values of the TaxonomyActivatedPolicyTypesItem enum
// A concrete instance of `TaxonomyActivatedPolicyTypesItemInput` can be one of the following:
//
//	TaxonomyActivatedPolicyTypesItemPolicyTypeUnspecified
//	TaxonomyActivatedPolicyTypesItemFineGrainedAccessControl
type TaxonomyActivatedPolicyTypesItemInput interface {
	pulumi.Input

	ToTaxonomyActivatedPolicyTypesItemOutput() TaxonomyActivatedPolicyTypesItemOutput
	ToTaxonomyActivatedPolicyTypesItemOutputWithContext(context.Context) TaxonomyActivatedPolicyTypesItemOutput
}

var taxonomyActivatedPolicyTypesItemPtrType = reflect.TypeOf((**TaxonomyActivatedPolicyTypesItem)(nil)).Elem()

type TaxonomyActivatedPolicyTypesItemPtrInput interface {
	pulumi.Input

	ToTaxonomyActivatedPolicyTypesItemPtrOutput() TaxonomyActivatedPolicyTypesItemPtrOutput
	ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(context.Context) TaxonomyActivatedPolicyTypesItemPtrOutput
}

type taxonomyActivatedPolicyTypesItemPtr string

func TaxonomyActivatedPolicyTypesItemPtr(v string) TaxonomyActivatedPolicyTypesItemPtrInput {
	return (*taxonomyActivatedPolicyTypesItemPtr)(&v)
}

func (*taxonomyActivatedPolicyTypesItemPtr) ElementType() reflect.Type {
	return taxonomyActivatedPolicyTypesItemPtrType
}

func (in *taxonomyActivatedPolicyTypesItemPtr) ToTaxonomyActivatedPolicyTypesItemPtrOutput() TaxonomyActivatedPolicyTypesItemPtrOutput {
	return pulumi.ToOutput(in).(TaxonomyActivatedPolicyTypesItemPtrOutput)
}

func (in *taxonomyActivatedPolicyTypesItemPtr) ToTaxonomyActivatedPolicyTypesItemPtrOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TaxonomyActivatedPolicyTypesItemPtrOutput)
}

// TaxonomyActivatedPolicyTypesItemArrayInput is an input type that accepts TaxonomyActivatedPolicyTypesItemArray and TaxonomyActivatedPolicyTypesItemArrayOutput values.
// You can construct a concrete instance of `TaxonomyActivatedPolicyTypesItemArrayInput` via:
//
//	TaxonomyActivatedPolicyTypesItemArray{ TaxonomyActivatedPolicyTypesItemArgs{...} }
type TaxonomyActivatedPolicyTypesItemArrayInput interface {
	pulumi.Input

	ToTaxonomyActivatedPolicyTypesItemArrayOutput() TaxonomyActivatedPolicyTypesItemArrayOutput
	ToTaxonomyActivatedPolicyTypesItemArrayOutputWithContext(context.Context) TaxonomyActivatedPolicyTypesItemArrayOutput
}

type TaxonomyActivatedPolicyTypesItemArray []TaxonomyActivatedPolicyTypesItem

func (TaxonomyActivatedPolicyTypesItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaxonomyActivatedPolicyTypesItem)(nil)).Elem()
}

func (i TaxonomyActivatedPolicyTypesItemArray) ToTaxonomyActivatedPolicyTypesItemArrayOutput() TaxonomyActivatedPolicyTypesItemArrayOutput {
	return i.ToTaxonomyActivatedPolicyTypesItemArrayOutputWithContext(context.Background())
}

func (i TaxonomyActivatedPolicyTypesItemArray) ToTaxonomyActivatedPolicyTypesItemArrayOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaxonomyActivatedPolicyTypesItemArrayOutput)
}

type TaxonomyActivatedPolicyTypesItemArrayOutput struct{ *pulumi.OutputState }

func (TaxonomyActivatedPolicyTypesItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TaxonomyActivatedPolicyTypesItem)(nil)).Elem()
}

func (o TaxonomyActivatedPolicyTypesItemArrayOutput) ToTaxonomyActivatedPolicyTypesItemArrayOutput() TaxonomyActivatedPolicyTypesItemArrayOutput {
	return o
}

func (o TaxonomyActivatedPolicyTypesItemArrayOutput) ToTaxonomyActivatedPolicyTypesItemArrayOutputWithContext(ctx context.Context) TaxonomyActivatedPolicyTypesItemArrayOutput {
	return o
}

func (o TaxonomyActivatedPolicyTypesItemArrayOutput) Index(i pulumi.IntInput) TaxonomyActivatedPolicyTypesItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TaxonomyActivatedPolicyTypesItem {
		return vs[0].([]TaxonomyActivatedPolicyTypesItem)[vs[1].(int)]
	}).(TaxonomyActivatedPolicyTypesItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EntryTypeInput)(nil)).Elem(), EntryType("ENTRY_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*EntryTypePtrInput)(nil)).Elem(), EntryType("ENTRY_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeInput)(nil)).Elem(), GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("PRIMITIVE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrInput)(nil)).Elem(), GoogleCloudDatacatalogV1beta1FieldTypePrimitiveType("PRIMITIVE_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyActivatedPolicyTypesItemInput)(nil)).Elem(), TaxonomyActivatedPolicyTypesItem("POLICY_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyActivatedPolicyTypesItemPtrInput)(nil)).Elem(), TaxonomyActivatedPolicyTypesItem("POLICY_TYPE_UNSPECIFIED"))
	pulumi.RegisterInputType(reflect.TypeOf((*TaxonomyActivatedPolicyTypesItemArrayInput)(nil)).Elem(), TaxonomyActivatedPolicyTypesItemArray{})
	pulumi.RegisterOutputType(EntryTypeOutput{})
	pulumi.RegisterOutputType(EntryTypePtrOutput{})
	pulumi.RegisterOutputType(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypeOutput{})
	pulumi.RegisterOutputType(GoogleCloudDatacatalogV1beta1FieldTypePrimitiveTypePtrOutput{})
	pulumi.RegisterOutputType(TaxonomyActivatedPolicyTypesItemOutput{})
	pulumi.RegisterOutputType(TaxonomyActivatedPolicyTypesItemPtrOutput{})
	pulumi.RegisterOutputType(TaxonomyActivatedPolicyTypesItemArrayOutput{})
}
