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
    /// JobSpec describes how the job will look.
    /// </summary>
    [OutputType]
    public sealed class JobSpecResponse
    {
        /// <summary>
        /// Optional. Describes the execution that will be created when running a job.
        /// </summary>
        public readonly Outputs.ExecutionTemplateSpecResponse Template;

        [OutputConstructor]
        private JobSpecResponse(Outputs.ExecutionTemplateSpecResponse template)
        {
            Template = template;
        }
    }
}