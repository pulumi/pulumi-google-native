// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Looker.V1.Outputs
{

    /// <summary>
    /// Metadata about users for a Looker instance.
    /// </summary>
    [OutputType]
    public sealed class UserMetadataResponse
    {
        /// <summary>
        /// Optional. The number of additional developer users the instance owner has purchased.
        /// </summary>
        public readonly int AdditionalDeveloperUserCount;
        /// <summary>
        /// Optional. The number of additional standard users the instance owner has purchased.
        /// </summary>
        public readonly int AdditionalStandardUserCount;
        /// <summary>
        /// Optional. The number of additional viewer users the instance owner has purchased.
        /// </summary>
        public readonly int AdditionalViewerUserCount;

        [OutputConstructor]
        private UserMetadataResponse(
            int additionalDeveloperUserCount,

            int additionalStandardUserCount,

            int additionalViewerUserCount)
        {
            AdditionalDeveloperUserCount = additionalDeveloperUserCount;
            AdditionalStandardUserCount = additionalStandardUserCount;
            AdditionalViewerUserCount = additionalViewerUserCount;
        }
    }
}