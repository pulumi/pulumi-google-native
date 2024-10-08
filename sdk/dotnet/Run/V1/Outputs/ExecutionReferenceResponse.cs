// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Outputs
{

    /// <summary>
    /// Reference to an Execution. Use /Executions.GetExecution with the given name to get full execution including the latest status.
    /// </summary>
    [OutputType]
    public sealed class ExecutionReferenceResponse
    {
        /// <summary>
        /// Optional. Completion timestamp of the execution.
        /// </summary>
        public readonly string CompletionTimestamp;
        /// <summary>
        /// Optional. Creation timestamp of the execution.
        /// </summary>
        public readonly string CreationTimestamp;
        /// <summary>
        /// Optional. Name of the execution.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private ExecutionReferenceResponse(
            string completionTimestamp,

            string creationTimestamp,

            string name)
        {
            CompletionTimestamp = completionTimestamp;
            CreationTimestamp = creationTimestamp;
            Name = name;
        }
    }
}
