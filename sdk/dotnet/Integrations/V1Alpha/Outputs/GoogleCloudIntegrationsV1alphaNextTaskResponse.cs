// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// The task that is next in line to be executed, if the condition specified evaluated to true.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIntegrationsV1alphaNextTaskResponse
    {
        /// <summary>
        /// Standard filter expression for this task to become an eligible next task.
        /// </summary>
        public readonly string Condition;
        /// <summary>
        /// User-provided description intended to give additional business context about the task.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// User-provided label that is attached to this edge in the UI.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// ID of the next task.
        /// </summary>
        public readonly string TaskConfigId;
        /// <summary>
        /// Task number of the next task.
        /// </summary>
        public readonly string TaskId;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaNextTaskResponse(
            string condition,

            string description,

            string displayName,

            string taskConfigId,

            string taskId)
        {
            Condition = condition;
            Description = description;
            DisplayName = displayName;
            TaskConfigId = taskConfigId;
            TaskId = taskId;
        }
    }
}