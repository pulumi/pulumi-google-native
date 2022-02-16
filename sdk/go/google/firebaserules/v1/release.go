// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Create a `Release`. Release names should reflect the developer's deployment practices. For example, the release name may include the environment name, application name, application version, or any other name meaningful to the developer. Once a `Release` refers to a `Ruleset`, the rules can be enforced by Firebase Rules-enabled services. More than one `Release` may be 'live' concurrently. Consider the following three `Release` names for `projects/foo` and the `Ruleset` to which they refer. Release Name -> Ruleset Name * projects/foo/releases/prod -> projects/foo/rulesets/uuid123 * projects/foo/releases/prod/beta -> projects/foo/rulesets/uuid123 * projects/foo/releases/prod/v23 -> projects/foo/rulesets/uuid456 The relationships reflect a `Ruleset` rollout in progress. The `prod` and `prod/beta` releases refer to the same `Ruleset`. However, `prod/v23` refers to a new `Ruleset`. The `Ruleset` reference for a `Release` may be updated using the UpdateRelease method.
type Release struct {
	pulumi.CustomResourceState

	// Time the release was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Format: `projects/{project_id}/releases/{release_id}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName pulumi.StringOutput `pulumi:"rulesetName"`
	// Time the release was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RulesetName == nil {
		return nil, errors.New("invalid value for required argument 'RulesetName'")
	}
	var resource Release
	err := ctx.RegisterResource("google-native:firebaserules/v1:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("google-native:firebaserules/v1:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
}

type ReleaseState struct {
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// Format: `projects/{project_id}/releases/{release_id}`
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName string `pulumi:"rulesetName"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// Format: `projects/{project_id}/releases/{release_id}`
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Name of the `Ruleset` referred to by this `Release`. The `Ruleset` must exist the `Release` to be created.
	RulesetName pulumi.StringInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterOutputType(ReleaseOutput{})
}