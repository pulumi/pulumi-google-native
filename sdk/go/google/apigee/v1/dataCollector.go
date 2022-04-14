// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new data collector.
type DataCollector struct {
	pulumi.CustomResourceState

	// The time at which the data collector was created in milliseconds since the epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// A description of the data collector.
	Description pulumi.StringOutput `pulumi:"description"`
	// The time at which the Data Collector was last updated in milliseconds since the epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// ID of the data collector. Must begin with `dc_`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The type of data this data collector will collect.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataCollector registers a new resource with the given unique name, arguments, and options.
func NewDataCollector(ctx *pulumi.Context,
	name string, args *DataCollectorArgs, opts ...pulumi.ResourceOption) (*DataCollector, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	var resource DataCollector
	err := ctx.RegisterResource("google-native:apigee/v1:DataCollector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCollector gets an existing DataCollector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCollector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCollectorState, opts ...pulumi.ResourceOption) (*DataCollector, error) {
	var resource DataCollector
	err := ctx.ReadResource("google-native:apigee/v1:DataCollector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCollector resources.
type dataCollectorState struct {
}

type DataCollectorState struct {
}

func (DataCollectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectorState)(nil)).Elem()
}

type dataCollectorArgs struct {
	// ID of the data collector. Overrides any ID in the data collector resource. Must be a string beginning with `dc_` that contains only letters, numbers, and underscores.
	DataCollectorId *string `pulumi:"dataCollectorId"`
	// A description of the data collector.
	Description *string `pulumi:"description"`
	// ID of the data collector. Must begin with `dc_`.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Immutable. The type of data this data collector will collect.
	Type *DataCollectorType `pulumi:"type"`
}

// The set of arguments for constructing a DataCollector resource.
type DataCollectorArgs struct {
	// ID of the data collector. Overrides any ID in the data collector resource. Must be a string beginning with `dc_` that contains only letters, numbers, and underscores.
	DataCollectorId pulumi.StringPtrInput
	// A description of the data collector.
	Description pulumi.StringPtrInput
	// ID of the data collector. Must begin with `dc_`.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Immutable. The type of data this data collector will collect.
	Type DataCollectorTypePtrInput
}

func (DataCollectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCollectorArgs)(nil)).Elem()
}

type DataCollectorInput interface {
	pulumi.Input

	ToDataCollectorOutput() DataCollectorOutput
	ToDataCollectorOutputWithContext(ctx context.Context) DataCollectorOutput
}

func (*DataCollector) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollector)(nil)).Elem()
}

func (i *DataCollector) ToDataCollectorOutput() DataCollectorOutput {
	return i.ToDataCollectorOutputWithContext(context.Background())
}

func (i *DataCollector) ToDataCollectorOutputWithContext(ctx context.Context) DataCollectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCollectorOutput)
}

type DataCollectorOutput struct{ *pulumi.OutputState }

func (DataCollectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCollector)(nil)).Elem()
}

func (o DataCollectorOutput) ToDataCollectorOutput() DataCollectorOutput {
	return o
}

func (o DataCollectorOutput) ToDataCollectorOutputWithContext(ctx context.Context) DataCollectorOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCollectorInput)(nil)).Elem(), &DataCollector{})
	pulumi.RegisterOutputType(DataCollectorOutput{})
}