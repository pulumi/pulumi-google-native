// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a new ImportJob within a KeyRing. ImportJob.import_method is required.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type ImportJob struct {
	pulumi.CustomResourceState

	// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen ImportMethod is one with a protection level of HSM.
	Attestation KeyOperationAttestationResponseOutput `pulumi:"attestation"`
	// The time at which this ImportJob was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time this ImportJob expired. Only present if state is EXPIRED.
	ExpireEventTime pulumi.StringOutput `pulumi:"expireEventTime"`
	// The time at which this ImportJob is scheduled for expiration and can no longer be used to import key material.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// The time this ImportJob's key material was generated.
	GenerateTime pulumi.StringOutput `pulumi:"generateTime"`
	// Immutable. The wrapping method to be used for incoming key material.
	ImportMethod pulumi.StringOutput `pulumi:"importMethod"`
	// The resource name for this ImportJob in the format `projects/*/locations/*/keyRings/*/importJobs/*`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
	ProtectionLevel pulumi.StringOutput `pulumi:"protectionLevel"`
	// The public key with which to wrap key material prior to import. Only returned if state is ACTIVE.
	PublicKey WrappingPublicKeyResponseOutput `pulumi:"publicKey"`
	// The current state of the ImportJob, indicating if it can be used.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewImportJob registers a new resource with the given unique name, arguments, and options.
func NewImportJob(ctx *pulumi.Context,
	name string, args *ImportJobArgs, opts ...pulumi.ResourceOption) (*ImportJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImportJobId == nil {
		return nil, errors.New("invalid value for required argument 'ImportJobId'")
	}
	if args.ImportMethod == nil {
		return nil, errors.New("invalid value for required argument 'ImportMethod'")
	}
	if args.KeyRingId == nil {
		return nil, errors.New("invalid value for required argument 'KeyRingId'")
	}
	if args.ProtectionLevel == nil {
		return nil, errors.New("invalid value for required argument 'ProtectionLevel'")
	}
	var resource ImportJob
	err := ctx.RegisterResource("google-native:cloudkms/v1:ImportJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImportJob gets an existing ImportJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImportJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImportJobState, opts ...pulumi.ResourceOption) (*ImportJob, error) {
	var resource ImportJob
	err := ctx.ReadResource("google-native:cloudkms/v1:ImportJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImportJob resources.
type importJobState struct {
}

type ImportJobState struct {
}

func (ImportJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*importJobState)(nil)).Elem()
}

type importJobArgs struct {
	ImportJobId string `pulumi:"importJobId"`
	// Immutable. The wrapping method to be used for incoming key material.
	ImportMethod ImportJobImportMethod `pulumi:"importMethod"`
	KeyRingId    string                `pulumi:"keyRingId"`
	Location     *string               `pulumi:"location"`
	Project      *string               `pulumi:"project"`
	// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
	ProtectionLevel ImportJobProtectionLevel `pulumi:"protectionLevel"`
}

// The set of arguments for constructing a ImportJob resource.
type ImportJobArgs struct {
	ImportJobId pulumi.StringInput
	// Immutable. The wrapping method to be used for incoming key material.
	ImportMethod ImportJobImportMethodInput
	KeyRingId    pulumi.StringInput
	Location     pulumi.StringPtrInput
	Project      pulumi.StringPtrInput
	// Immutable. The protection level of the ImportJob. This must match the protection_level of the version_template on the CryptoKey you attempt to import into.
	ProtectionLevel ImportJobProtectionLevelInput
}

func (ImportJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*importJobArgs)(nil)).Elem()
}

type ImportJobInput interface {
	pulumi.Input

	ToImportJobOutput() ImportJobOutput
	ToImportJobOutputWithContext(ctx context.Context) ImportJobOutput
}

func (*ImportJob) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportJob)(nil)).Elem()
}

func (i *ImportJob) ToImportJobOutput() ImportJobOutput {
	return i.ToImportJobOutputWithContext(context.Background())
}

func (i *ImportJob) ToImportJobOutputWithContext(ctx context.Context) ImportJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportJobOutput)
}

type ImportJobOutput struct{ *pulumi.OutputState }

func (ImportJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportJob)(nil)).Elem()
}

func (o ImportJobOutput) ToImportJobOutput() ImportJobOutput {
	return o
}

func (o ImportJobOutput) ToImportJobOutputWithContext(ctx context.Context) ImportJobOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImportJobInput)(nil)).Elem(), &ImportJob{})
	pulumi.RegisterOutputType(ImportJobOutput{})
}