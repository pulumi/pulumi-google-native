// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataflow.V1b3.Outputs
{

    /// <summary>
    /// Configuration options for sampling elements.
    /// </summary>
    [OutputType]
    public sealed class DataSamplingConfigResponse
    {
        /// <summary>
        /// List of given sampling behaviors to enable. For example, specifying behaviors = [ALWAYS_ON] samples in-flight elements but does not sample exceptions. Can be used to specify multiple behaviors like, behaviors = [ALWAYS_ON, EXCEPTIONS] for specifying periodic sampling and exception sampling. If DISABLED is in the list, then sampling will be disabled and ignore the other given behaviors. Ordering does not matter.
        /// </summary>
        public readonly ImmutableArray<string> Behaviors;

        [OutputConstructor]
        private DataSamplingConfigResponse(ImmutableArray<string> behaviors)
        {
            Behaviors = behaviors;
        }
    }
}