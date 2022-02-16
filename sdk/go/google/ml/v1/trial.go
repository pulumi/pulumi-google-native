// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds a user provided trial to a study.
// Auto-naming is currently not supported for this resource.
type Trial struct {
	pulumi.CustomResourceState

	// The identifier of the client that originally requested this trial.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Time at which the trial's status changed to COMPLETED.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// The final measurement containing the objective value.
	FinalMeasurement GoogleCloudMlV1__MeasurementResponseOutput `pulumi:"finalMeasurement"`
	// A human readable string describing why the trial is infeasible. This should only be set if trial_infeasible is true.
	InfeasibleReason pulumi.StringOutput `pulumi:"infeasibleReason"`
	// A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
	Measurements GoogleCloudMlV1__MeasurementResponseArrayOutput `pulumi:"measurements"`
	// Name of the trial assigned by the service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parameters of the trial.
	Parameters GoogleCloudMlV1_Trial_ParameterResponseArrayOutput `pulumi:"parameters"`
	// Time at which the trial was started.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
	// The detailed state of a trial.
	State pulumi.StringOutput `pulumi:"state"`
	// If true, the parameters in this trial are not attempted again.
	TrialInfeasible pulumi.BoolOutput `pulumi:"trialInfeasible"`
}

// NewTrial registers a new resource with the given unique name, arguments, and options.
func NewTrial(ctx *pulumi.Context,
	name string, args *TrialArgs, opts ...pulumi.ResourceOption) (*Trial, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StudyId == nil {
		return nil, errors.New("invalid value for required argument 'StudyId'")
	}
	var resource Trial
	err := ctx.RegisterResource("google-native:ml/v1:Trial", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrial gets an existing Trial resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrial(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrialState, opts ...pulumi.ResourceOption) (*Trial, error) {
	var resource Trial
	err := ctx.ReadResource("google-native:ml/v1:Trial", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trial resources.
type trialState struct {
}

type TrialState struct {
}

func (TrialState) ElementType() reflect.Type {
	return reflect.TypeOf((*trialState)(nil)).Elem()
}

type trialArgs struct {
	// The final measurement containing the objective value.
	FinalMeasurement *GoogleCloudMlV1__Measurement `pulumi:"finalMeasurement"`
	Location         *string                       `pulumi:"location"`
	// A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
	Measurements []GoogleCloudMlV1__Measurement `pulumi:"measurements"`
	// The parameters of the trial.
	Parameters []GoogleCloudMlV1_Trial_Parameter `pulumi:"parameters"`
	Project    *string                           `pulumi:"project"`
	// The detailed state of a trial.
	State   *TrialStateEnum `pulumi:"state"`
	StudyId string          `pulumi:"studyId"`
}

// The set of arguments for constructing a Trial resource.
type TrialArgs struct {
	// The final measurement containing the objective value.
	FinalMeasurement GoogleCloudMlV1__MeasurementPtrInput
	Location         pulumi.StringPtrInput
	// A list of measurements that are strictly lexicographically ordered by their induced tuples (steps, elapsed_time). These are used for early stopping computations.
	Measurements GoogleCloudMlV1__MeasurementArrayInput
	// The parameters of the trial.
	Parameters GoogleCloudMlV1_Trial_ParameterArrayInput
	Project    pulumi.StringPtrInput
	// The detailed state of a trial.
	State   TrialStateEnumPtrInput
	StudyId pulumi.StringInput
}

func (TrialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trialArgs)(nil)).Elem()
}

type TrialInput interface {
	pulumi.Input

	ToTrialOutput() TrialOutput
	ToTrialOutputWithContext(ctx context.Context) TrialOutput
}

func (*Trial) ElementType() reflect.Type {
	return reflect.TypeOf((**Trial)(nil)).Elem()
}

func (i *Trial) ToTrialOutput() TrialOutput {
	return i.ToTrialOutputWithContext(context.Background())
}

func (i *Trial) ToTrialOutputWithContext(ctx context.Context) TrialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrialOutput)
}

type TrialOutput struct{ *pulumi.OutputState }

func (TrialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trial)(nil)).Elem()
}

func (o TrialOutput) ToTrialOutput() TrialOutput {
	return o
}

func (o TrialOutput) ToTrialOutputWithContext(ctx context.Context) TrialOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrialInput)(nil)).Elem(), &Trial{})
	pulumi.RegisterOutputType(TrialOutput{})
}