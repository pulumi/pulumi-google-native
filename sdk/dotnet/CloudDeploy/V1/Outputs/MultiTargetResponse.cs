// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// Information specifying a multiTarget.
    /// </summary>
    [OutputType]
    public sealed class MultiTargetResponse
    {
        /// <summary>
        /// The target_ids of this multiTarget.
        /// </summary>
        public readonly ImmutableArray<string> TargetIds;

        [OutputConstructor]
        private MultiTargetResponse(ImmutableArray<string> targetIds)
        {
            TargetIds = targetIds;
        }
    }
}