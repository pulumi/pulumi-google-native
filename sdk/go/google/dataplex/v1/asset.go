// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an asset resource.
// Auto-naming is currently not supported for this resource.
type Asset struct {
	pulumi.CustomResourceState

	// The time when the asset was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Description of the asset.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec GoogleCloudDataplexV1AssetDiscoverySpecResponseOutput `pulumi:"discoverySpec"`
	// Status of the discovery feature applied to data referenced by this asset.
	DiscoveryStatus GoogleCloudDataplexV1AssetDiscoveryStatusResponseOutput `pulumi:"discoveryStatus"`
	// Optional. User friendly display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. User defined labels for the asset.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The relative resource name of the asset, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/assets/{asset_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specification of the resource that is referenced by this asset.
	ResourceSpec GoogleCloudDataplexV1AssetResourceSpecResponseOutput `pulumi:"resourceSpec"`
	// Status of the resource referenced by this asset.
	ResourceStatus GoogleCloudDataplexV1AssetResourceStatusResponseOutput `pulumi:"resourceStatus"`
	// Status of the security policy applied to resource referenced by this asset.
	SecurityStatus GoogleCloudDataplexV1AssetSecurityStatusResponseOutput `pulumi:"securityStatus"`
	// Current state of the asset.
	State pulumi.StringOutput `pulumi:"state"`
	// System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the asset was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAsset registers a new resource with the given unique name, arguments, and options.
func NewAsset(ctx *pulumi.Context,
	name string, args *AssetArgs, opts ...pulumi.ResourceOption) (*Asset, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AssetId == nil {
		return nil, errors.New("invalid value for required argument 'AssetId'")
	}
	if args.LakeId == nil {
		return nil, errors.New("invalid value for required argument 'LakeId'")
	}
	if args.ResourceSpec == nil {
		return nil, errors.New("invalid value for required argument 'ResourceSpec'")
	}
	var resource Asset
	err := ctx.RegisterResource("google-native:dataplex/v1:Asset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAsset gets an existing Asset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAsset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetState, opts ...pulumi.ResourceOption) (*Asset, error) {
	var resource Asset
	err := ctx.ReadResource("google-native:dataplex/v1:Asset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Asset resources.
type assetState struct {
}

type AssetState struct {
}

func (AssetState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetState)(nil)).Elem()
}

type assetArgs struct {
	// Required. Asset identifier. This ID will be used to generate names such as table names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the zone.
	AssetId string `pulumi:"assetId"`
	// Optional. Description of the asset.
	Description *string `pulumi:"description"`
	// Optional. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec *GoogleCloudDataplexV1AssetDiscoverySpec `pulumi:"discoverySpec"`
	// Optional. User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// Optional. User defined labels for the asset.
	Labels   map[string]string `pulumi:"labels"`
	LakeId   string            `pulumi:"lakeId"`
	Location *string           `pulumi:"location"`
	Project  *string           `pulumi:"project"`
	// Specification of the resource that is referenced by this asset.
	ResourceSpec GoogleCloudDataplexV1AssetResourceSpec `pulumi:"resourceSpec"`
	// Optional. Only validate the request, but do not perform mutations. The default is false.
	ValidateOnly *string `pulumi:"validateOnly"`
	Zone         *string `pulumi:"zone"`
}

// The set of arguments for constructing a Asset resource.
type AssetArgs struct {
	// Required. Asset identifier. This ID will be used to generate names such as table names when publishing metadata to Hive Metastore and BigQuery. * Must contain only lowercase letters, numbers and hyphens. * Must start with a letter. * Must end with a number or a letter. * Must be between 1-63 characters. * Must be unique within the zone.
	AssetId pulumi.StringInput
	// Optional. Description of the asset.
	Description pulumi.StringPtrInput
	// Optional. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
	DiscoverySpec GoogleCloudDataplexV1AssetDiscoverySpecPtrInput
	// Optional. User friendly display name.
	DisplayName pulumi.StringPtrInput
	// Optional. User defined labels for the asset.
	Labels   pulumi.StringMapInput
	LakeId   pulumi.StringInput
	Location pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Specification of the resource that is referenced by this asset.
	ResourceSpec GoogleCloudDataplexV1AssetResourceSpecInput
	// Optional. Only validate the request, but do not perform mutations. The default is false.
	ValidateOnly pulumi.StringPtrInput
	Zone         pulumi.StringPtrInput
}

func (AssetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetArgs)(nil)).Elem()
}

type AssetInput interface {
	pulumi.Input

	ToAssetOutput() AssetOutput
	ToAssetOutputWithContext(ctx context.Context) AssetOutput
}

func (*Asset) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (i *Asset) ToAssetOutput() AssetOutput {
	return i.ToAssetOutputWithContext(context.Background())
}

func (i *Asset) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetOutput)
}

type AssetOutput struct{ *pulumi.OutputState }

func (AssetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Asset)(nil)).Elem()
}

func (o AssetOutput) ToAssetOutput() AssetOutput {
	return o
}

func (o AssetOutput) ToAssetOutputWithContext(ctx context.Context) AssetOutput {
	return o
}

// The time when the asset was created.
func (o AssetOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Description of the asset.
func (o AssetOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. Specification of the discovery feature applied to data referenced by this asset. When this spec is left unset, the asset will use the spec set on the parent zone.
func (o AssetOutput) DiscoverySpec() GoogleCloudDataplexV1AssetDiscoverySpecResponseOutput {
	return o.ApplyT(func(v *Asset) GoogleCloudDataplexV1AssetDiscoverySpecResponseOutput { return v.DiscoverySpec }).(GoogleCloudDataplexV1AssetDiscoverySpecResponseOutput)
}

// Status of the discovery feature applied to data referenced by this asset.
func (o AssetOutput) DiscoveryStatus() GoogleCloudDataplexV1AssetDiscoveryStatusResponseOutput {
	return o.ApplyT(func(v *Asset) GoogleCloudDataplexV1AssetDiscoveryStatusResponseOutput { return v.DiscoveryStatus }).(GoogleCloudDataplexV1AssetDiscoveryStatusResponseOutput)
}

// Optional. User friendly display name.
func (o AssetOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. User defined labels for the asset.
func (o AssetOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The relative resource name of the asset, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/zones/{zone_id}/assets/{asset_id}.
func (o AssetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specification of the resource that is referenced by this asset.
func (o AssetOutput) ResourceSpec() GoogleCloudDataplexV1AssetResourceSpecResponseOutput {
	return o.ApplyT(func(v *Asset) GoogleCloudDataplexV1AssetResourceSpecResponseOutput { return v.ResourceSpec }).(GoogleCloudDataplexV1AssetResourceSpecResponseOutput)
}

// Status of the resource referenced by this asset.
func (o AssetOutput) ResourceStatus() GoogleCloudDataplexV1AssetResourceStatusResponseOutput {
	return o.ApplyT(func(v *Asset) GoogleCloudDataplexV1AssetResourceStatusResponseOutput { return v.ResourceStatus }).(GoogleCloudDataplexV1AssetResourceStatusResponseOutput)
}

// Status of the security policy applied to resource referenced by this asset.
func (o AssetOutput) SecurityStatus() GoogleCloudDataplexV1AssetSecurityStatusResponseOutput {
	return o.ApplyT(func(v *Asset) GoogleCloudDataplexV1AssetSecurityStatusResponseOutput { return v.SecurityStatus }).(GoogleCloudDataplexV1AssetSecurityStatusResponseOutput)
}

// Current state of the asset.
func (o AssetOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// System generated globally unique ID for the asset. This ID will be different if the asset is deleted and re-created with the same name.
func (o AssetOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the asset was last updated.
func (o AssetOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Asset) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssetInput)(nil)).Elem(), &Asset{})
	pulumi.RegisterOutputType(AssetOutput{})
}