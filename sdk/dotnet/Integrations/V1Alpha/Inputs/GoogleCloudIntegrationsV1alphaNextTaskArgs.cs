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
    public sealed class GoogleCloudIntegrationsV1alphaNextTaskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Standard filter expression for this task to become an eligible next task.
        /// </summary>
        [Input("condition")]
        public Input<string>? Condition { get; set; }

        /// <summary>
        /// User-provided description intended to give additional business context about the task.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User-provided label that is attached to this edge in the UI.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// ID of the next task.
        /// </summary>
        [Input("taskConfigId")]
        public Input<string>? TaskConfigId { get; set; }

        /// <summary>
        /// Task number of the next task.
        /// </summary>
        [Input("taskId")]
        public Input<string>? TaskId { get; set; }

        public GoogleCloudIntegrationsV1alphaNextTaskArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaNextTaskArgs Empty => new GoogleCloudIntegrationsV1alphaNextTaskArgs();
    }
}
