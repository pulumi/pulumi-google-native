// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates and returns a new ReferenceImage resource. The `bounding_poly` field is optional. If `bounding_poly` is not specified, the system will try to detect regions of interest in the image that are compatible with the product_category on the parent product. If it is specified, detection is ALWAYS skipped. The system converts polygons into non-rotated rectangles. Note that the pipeline will resize the image if the image resolution is too large to process (above 50MP). Possible errors: * Returns INVALID_ARGUMENT if the image_uri is missing or longer than 4096 characters. * Returns INVALID_ARGUMENT if the product does not exist. * Returns INVALID_ARGUMENT if bounding_poly is not provided, and nothing compatible with the parent product's product_category is detected. * Returns INVALID_ARGUMENT if bounding_poly contains more than 10 polygons.
type ReferenceImage struct {
	pulumi.CustomResourceState

	// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
	BoundingPolys BoundingPolyResponseArrayOutput `pulumi:"boundingPolys"`
	// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewReferenceImage registers a new resource with the given unique name, arguments, and options.
func NewReferenceImage(ctx *pulumi.Context,
	name string, args *ReferenceImageArgs, opts ...pulumi.ResourceOption) (*ReferenceImage, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProductId == nil {
		return nil, errors.New("invalid value for required argument 'ProductId'")
	}
	if args.Uri == nil {
		return nil, errors.New("invalid value for required argument 'Uri'")
	}
	var resource ReferenceImage
	err := ctx.RegisterResource("google-native:vision/v1:ReferenceImage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReferenceImage gets an existing ReferenceImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReferenceImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReferenceImageState, opts ...pulumi.ResourceOption) (*ReferenceImage, error) {
	var resource ReferenceImage
	err := ctx.ReadResource("google-native:vision/v1:ReferenceImage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReferenceImage resources.
type referenceImageState struct {
}

type ReferenceImageState struct {
}

func (ReferenceImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceImageState)(nil)).Elem()
}

type referenceImageArgs struct {
	// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
	BoundingPolys []BoundingPoly `pulumi:"boundingPolys"`
	Location      *string        `pulumi:"location"`
	// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
	Name      *string `pulumi:"name"`
	ProductId string  `pulumi:"productId"`
	Project   *string `pulumi:"project"`
	// A user-supplied resource id for the ReferenceImage to be added. If set, the server will attempt to use this value as the resource id. If it is already in use, an error is returned with code ALREADY_EXISTS. Must be at most 128 characters long. It cannot contain the character `/`.
	ReferenceImageId *string `pulumi:"referenceImageId"`
	// The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
	Uri string `pulumi:"uri"`
}

// The set of arguments for constructing a ReferenceImage resource.
type ReferenceImageArgs struct {
	// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
	BoundingPolys BoundingPolyArrayInput
	Location      pulumi.StringPtrInput
	// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
	Name      pulumi.StringPtrInput
	ProductId pulumi.StringInput
	Project   pulumi.StringPtrInput
	// A user-supplied resource id for the ReferenceImage to be added. If set, the server will attempt to use this value as the resource id. If it is already in use, an error is returned with code ALREADY_EXISTS. Must be at most 128 characters long. It cannot contain the character `/`.
	ReferenceImageId pulumi.StringPtrInput
	// The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
	Uri pulumi.StringInput
}

func (ReferenceImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*referenceImageArgs)(nil)).Elem()
}

type ReferenceImageInput interface {
	pulumi.Input

	ToReferenceImageOutput() ReferenceImageOutput
	ToReferenceImageOutputWithContext(ctx context.Context) ReferenceImageOutput
}

func (*ReferenceImage) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceImage)(nil)).Elem()
}

func (i *ReferenceImage) ToReferenceImageOutput() ReferenceImageOutput {
	return i.ToReferenceImageOutputWithContext(context.Background())
}

func (i *ReferenceImage) ToReferenceImageOutputWithContext(ctx context.Context) ReferenceImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceImageOutput)
}

type ReferenceImageOutput struct{ *pulumi.OutputState }

func (ReferenceImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReferenceImage)(nil)).Elem()
}

func (o ReferenceImageOutput) ToReferenceImageOutput() ReferenceImageOutput {
	return o
}

func (o ReferenceImageOutput) ToReferenceImageOutputWithContext(ctx context.Context) ReferenceImageOutput {
	return o
}

// Optional. Bounding polygons around the areas of interest in the reference image. If this field is empty, the system will try to detect regions of interest. At most 10 bounding polygons will be used. The provided shape is converted into a non-rotated rectangle. Once converted, the small edge of the rectangle must be greater than or equal to 300 pixels. The aspect ratio must be 1:4 or less (i.e. 1:3 is ok; 1:5 is not).
func (o ReferenceImageOutput) BoundingPolys() BoundingPolyResponseArrayOutput {
	return o.ApplyT(func(v *ReferenceImage) BoundingPolyResponseArrayOutput { return v.BoundingPolys }).(BoundingPolyResponseArrayOutput)
}

// The resource name of the reference image. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID/referenceImages/IMAGE_ID`. This field is ignored when creating a reference image.
func (o ReferenceImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Google Cloud Storage URI of the reference image. The URI must start with `gs://`.
func (o ReferenceImageOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v *ReferenceImage) pulumi.StringOutput { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReferenceImageInput)(nil)).Elem(), &ReferenceImage{})
	pulumi.RegisterOutputType(ReferenceImageOutput{})
}