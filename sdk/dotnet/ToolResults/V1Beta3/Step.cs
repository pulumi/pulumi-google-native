// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3
{
    /// <summary>
    /// Creates a Step. The returned Step will have the id set. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - FAILED_PRECONDITION - if the step is too large (more than 10Mib) - NOT_FOUND - if the containing Execution does not exist
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:toolresults/v1beta3:Step")]
    public partial class Step : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the step status was set to complete. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        /// </summary>
        [Output("completionTime")]
        public Output<Outputs.TimestampResponse> CompletionTime { get; private set; } = null!;

        /// <summary>
        /// The time when the step was created. - In response: always set - In create/update request: never set
        /// </summary>
        [Output("creationTime")]
        public Output<Outputs.TimestampResponse> CreationTime { get; private set; } = null!;

        /// <summary>
        /// A description of this tool For example: mvn clean package -D skipTests=true - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// How much the device resource is used to perform the test. This is the device usage used for billing purpose, which is different from the run_duration, for example, infrastructure failure won't be charged for device usage. PRECONDITION_FAILED will be returned if one attempts to set a device_usage on a step which already has this field set. - In response: present if previously set. - In create request: optional - In update request: optional
        /// </summary>
        [Output("deviceUsageDuration")]
        public Output<Outputs.DurationResponse> DeviceUsageDuration { get; private set; } = null!;

        /// <summary>
        /// If the execution containing this step has any dimension_definition set, then this field allows the child to specify the values of the dimensions. The keys must exactly match the dimension_definition of the execution. For example, if the execution has `dimension_definition = ['attempt', 'device']` then a step must define values for those dimensions, eg. `dimension_value = ['attempt': '1', 'device': 'Nexus 6']` If a step does not participate in one dimension of the matrix, the value for that dimension should be empty string. For example, if one of the tests is executed by a runner which does not support retries, the step could have `dimension_value = ['attempt': '', 'device': 'Nexus 6']` If the step does not participate in any dimensions of the matrix, it may leave dimension_value unset. A PRECONDITION_FAILED will be returned if any of the keys do not exist in the dimension_definition of the execution. A PRECONDITION_FAILED will be returned if another step in this execution already has the same name and dimension_value, but differs on other data fields, for example, step field is different. A PRECONDITION_FAILED will be returned if dimension_value is set, and there is a dimension_definition in the execution which is not specified as one of the keys. - In response: present if set by create - In create request: optional - In update request: never set
        /// </summary>
        [Output("dimensionValue")]
        public Output<ImmutableArray<Outputs.StepDimensionValueEntryResponse>> DimensionValue { get; private set; } = null!;

        [Output("executionId")]
        public Output<string> ExecutionId { get; private set; } = null!;

        /// <summary>
        /// Whether any of the outputs of this step are images whose thumbnails can be fetched with ListThumbnails. - In response: always set - In create/update request: never set
        /// </summary>
        [Output("hasImages")]
        public Output<bool> HasImages { get; private set; } = null!;

        [Output("historyId")]
        public Output<string> HistoryId { get; private set; } = null!;

        /// <summary>
        /// Arbitrary user-supplied key/value pairs that are associated with the step. Users are responsible for managing the key namespace such that keys don't accidentally collide. An INVALID_ARGUMENT will be returned if the number of labels exceeds 100 or if the length of any of the keys or values exceeds 100 characters. - In response: always set - In create request: optional - In update request: optional; any new key/value pair will be added to the map, and any new value for an existing key will update that key's value
        /// </summary>
        [Output("labels")]
        public Output<ImmutableArray<Outputs.StepLabelsEntryResponse>> Labels { get; private set; } = null!;

        /// <summary>
        /// Details when multiple steps are run with the same configuration as a group. These details can be used identify which group this step is part of. It also identifies the groups 'primary step' which indexes all the group members. - In response: present if previously set. - In create request: optional, set iff this step was performed more than once. - In update request: optional
        /// </summary>
        [Output("multiStep")]
        public Output<Outputs.MultiStepResponse> MultiStep { get; private set; } = null!;

        /// <summary>
        /// A short human-readable name to display in the UI. Maximum of 100 characters. For example: Clean build A PRECONDITION_FAILED will be returned upon creating a new step if it shares its name and dimension_value with an existing step. If two steps represent a similar action, but have different dimension values, they should share the same name. For instance, if the same set of tests is run on two different platforms, the two steps should have the same name. - In response: always set - In create request: always set - In update request: never set
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Classification of the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        [Output("outcome")]
        public Output<Outputs.OutcomeResponse> Outcome { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A unique request ID for server to detect duplicated requests. For example, a UUID. Optional, but strongly recommended.
        /// </summary>
        [Output("requestId")]
        public Output<string?> RequestId { get; private set; } = null!;

        /// <summary>
        /// How long it took for this step to run. If unset, this is set to the difference between creation_time and completion_time when the step is set to the COMPLETE state. In some cases, it is appropriate to set this value separately: For instance, if a step is created, but the operation it represents is queued for a few minutes before it executes, it would be appropriate not to include the time spent queued in its run_duration. PRECONDITION_FAILED will be returned if one attempts to set a run_duration on a step which already has this field set. - In response: present if previously set; always present on COMPLETE step - In create request: optional - In update request: optional
        /// </summary>
        [Output("runDuration")]
        public Output<Outputs.DurationResponse> RunDuration { get; private set; } = null!;

        /// <summary>
        /// The initial state is IN_PROGRESS. The only legal state transitions are * IN_PROGRESS -&gt; COMPLETE A PRECONDITION_FAILED will be returned if an invalid transition is requested. It is valid to create Step with a state set to COMPLETE. The state can only be set to COMPLETE once. A PRECONDITION_FAILED will be returned if the state is set to COMPLETE multiple times. - In response: always set - In create/update request: optional
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// A unique identifier within a Execution for this Step. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response: always set - In create/update request: never set
        /// </summary>
        [Output("stepId")]
        public Output<string> StepId { get; private set; } = null!;

        /// <summary>
        /// An execution of a test runner.
        /// </summary>
        [Output("testExecutionStep")]
        public Output<Outputs.TestExecutionStepResponse> TestExecutionStep { get; private set; } = null!;

        /// <summary>
        /// An execution of a tool (used for steps we don't explicitly support).
        /// </summary>
        [Output("toolExecutionStep")]
        public Output<Outputs.ToolExecutionStepResponse> ToolExecutionStep { get; private set; } = null!;


        /// <summary>
        /// Create a Step resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Step(string name, StepArgs args, CustomResourceOptions? options = null)
            : base("google-native:toolresults/v1beta3:Step", name, args ?? new StepArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Step(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:toolresults/v1beta3:Step", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "executionId",
                    "historyId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Step resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Step Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Step(name, id, options);
        }
    }

    public sealed class StepArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the step status was set to complete. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        /// </summary>
        [Input("completionTime")]
        public Input<Inputs.TimestampArgs>? CompletionTime { get; set; }

        /// <summary>
        /// The time when the step was created. - In response: always set - In create/update request: never set
        /// </summary>
        [Input("creationTime")]
        public Input<Inputs.TimestampArgs>? CreationTime { get; set; }

        /// <summary>
        /// A description of this tool For example: mvn clean package -D skipTests=true - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// How much the device resource is used to perform the test. This is the device usage used for billing purpose, which is different from the run_duration, for example, infrastructure failure won't be charged for device usage. PRECONDITION_FAILED will be returned if one attempts to set a device_usage on a step which already has this field set. - In response: present if previously set. - In create request: optional - In update request: optional
        /// </summary>
        [Input("deviceUsageDuration")]
        public Input<Inputs.DurationArgs>? DeviceUsageDuration { get; set; }

        [Input("dimensionValue")]
        private InputList<Inputs.StepDimensionValueEntryArgs>? _dimensionValue;

        /// <summary>
        /// If the execution containing this step has any dimension_definition set, then this field allows the child to specify the values of the dimensions. The keys must exactly match the dimension_definition of the execution. For example, if the execution has `dimension_definition = ['attempt', 'device']` then a step must define values for those dimensions, eg. `dimension_value = ['attempt': '1', 'device': 'Nexus 6']` If a step does not participate in one dimension of the matrix, the value for that dimension should be empty string. For example, if one of the tests is executed by a runner which does not support retries, the step could have `dimension_value = ['attempt': '', 'device': 'Nexus 6']` If the step does not participate in any dimensions of the matrix, it may leave dimension_value unset. A PRECONDITION_FAILED will be returned if any of the keys do not exist in the dimension_definition of the execution. A PRECONDITION_FAILED will be returned if another step in this execution already has the same name and dimension_value, but differs on other data fields, for example, step field is different. A PRECONDITION_FAILED will be returned if dimension_value is set, and there is a dimension_definition in the execution which is not specified as one of the keys. - In response: present if set by create - In create request: optional - In update request: never set
        /// </summary>
        public InputList<Inputs.StepDimensionValueEntryArgs> DimensionValue
        {
            get => _dimensionValue ?? (_dimensionValue = new InputList<Inputs.StepDimensionValueEntryArgs>());
            set => _dimensionValue = value;
        }

        [Input("executionId", required: true)]
        public Input<string> ExecutionId { get; set; } = null!;

        /// <summary>
        /// Whether any of the outputs of this step are images whose thumbnails can be fetched with ListThumbnails. - In response: always set - In create/update request: never set
        /// </summary>
        [Input("hasImages")]
        public Input<bool>? HasImages { get; set; }

        [Input("historyId", required: true)]
        public Input<string> HistoryId { get; set; } = null!;

        [Input("labels")]
        private InputList<Inputs.StepLabelsEntryArgs>? _labels;

        /// <summary>
        /// Arbitrary user-supplied key/value pairs that are associated with the step. Users are responsible for managing the key namespace such that keys don't accidentally collide. An INVALID_ARGUMENT will be returned if the number of labels exceeds 100 or if the length of any of the keys or values exceeds 100 characters. - In response: always set - In create request: optional - In update request: optional; any new key/value pair will be added to the map, and any new value for an existing key will update that key's value
        /// </summary>
        public InputList<Inputs.StepLabelsEntryArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.StepLabelsEntryArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Details when multiple steps are run with the same configuration as a group. These details can be used identify which group this step is part of. It also identifies the groups 'primary step' which indexes all the group members. - In response: present if previously set. - In create request: optional, set iff this step was performed more than once. - In update request: optional
        /// </summary>
        [Input("multiStep")]
        public Input<Inputs.MultiStepArgs>? MultiStep { get; set; }

        /// <summary>
        /// A short human-readable name to display in the UI. Maximum of 100 characters. For example: Clean build A PRECONDITION_FAILED will be returned upon creating a new step if it shares its name and dimension_value with an existing step. If two steps represent a similar action, but have different dimension values, they should share the same name. For instance, if the same set of tests is run on two different platforms, the two steps should have the same name. - In response: always set - In create request: always set - In update request: never set
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Classification of the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        /// </summary>
        [Input("outcome")]
        public Input<Inputs.OutcomeArgs>? Outcome { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A unique request ID for server to detect duplicated requests. For example, a UUID. Optional, but strongly recommended.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// How long it took for this step to run. If unset, this is set to the difference between creation_time and completion_time when the step is set to the COMPLETE state. In some cases, it is appropriate to set this value separately: For instance, if a step is created, but the operation it represents is queued for a few minutes before it executes, it would be appropriate not to include the time spent queued in its run_duration. PRECONDITION_FAILED will be returned if one attempts to set a run_duration on a step which already has this field set. - In response: present if previously set; always present on COMPLETE step - In create request: optional - In update request: optional
        /// </summary>
        [Input("runDuration")]
        public Input<Inputs.DurationArgs>? RunDuration { get; set; }

        /// <summary>
        /// The initial state is IN_PROGRESS. The only legal state transitions are * IN_PROGRESS -&gt; COMPLETE A PRECONDITION_FAILED will be returned if an invalid transition is requested. It is valid to create Step with a state set to COMPLETE. The state can only be set to COMPLETE once. A PRECONDITION_FAILED will be returned if the state is set to COMPLETE multiple times. - In response: always set - In create/update request: optional
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.ToolResults.V1Beta3.StepState>? State { get; set; }

        /// <summary>
        /// A unique identifier within a Execution for this Step. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response: always set - In create/update request: never set
        /// </summary>
        [Input("stepId")]
        public Input<string>? StepId { get; set; }

        /// <summary>
        /// An execution of a test runner.
        /// </summary>
        [Input("testExecutionStep")]
        public Input<Inputs.TestExecutionStepArgs>? TestExecutionStep { get; set; }

        /// <summary>
        /// An execution of a tool (used for steps we don't explicitly support).
        /// </summary>
        [Input("toolExecutionStep")]
        public Input<Inputs.ToolExecutionStepArgs>? ToolExecutionStep { get; set; }

        public StepArgs()
        {
        }
        public static new StepArgs Empty => new StepArgs();
    }
}
