// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// The task that is next in line to be executed, if the condition specified evaluated to true.
    /// </summary>
    public sealed class EnterpriseCrmEventbusProtoNextTaskArgs : global::Pulumi.ResourceArgs
    {
        [Input("combinedConditions")]
        private InputList<Inputs.EnterpriseCrmEventbusProtoCombinedConditionArgs>? _combinedConditions;

        /// <summary>
        /// Combined condition for this task to become an eligible next task. Each of these combined_conditions are joined with logical OR. DEPRECATED: use `condition`
        /// </summary>
        [Obsolete(@"Combined condition for this task to become an eligible next task. Each of these combined_conditions are joined with logical OR. DEPRECATED: use `condition`")]
        public InputList<Inputs.EnterpriseCrmEventbusProtoCombinedConditionArgs> CombinedConditions
        {
            get => _combinedConditions ?? (_combinedConditions = new InputList<Inputs.EnterpriseCrmEventbusProtoCombinedConditionArgs>());
            set => _combinedConditions = value;
        }

        /// <summary>
        /// Standard filter expression for this task to become an eligible next task.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// User-provided description intended to give more business context about the next task edge or condition.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-provided label that is attached to this edge in the UI.
        /// </summary>
        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// ID of the next task.
        /// </summary>
        [Input("taskConfigId")]
        public Input<string>? TaskConfigId { get; set; }

        /// <summary>
        /// Task number of the next task.
        /// </summary>
        [Input("taskNumber")]
        public Input<string>? TaskNumber { get; set; }

        public EnterpriseCrmEventbusProtoNextTaskArgs()
        {
        }
        public static new EnterpriseCrmEventbusProtoNextTaskArgs Empty => new EnterpriseCrmEventbusProtoNextTaskArgs();
    }
}
