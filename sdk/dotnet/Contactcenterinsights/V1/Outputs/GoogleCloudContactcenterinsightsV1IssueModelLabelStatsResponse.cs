// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Contactcenterinsights.V1.Outputs
{

    /// <summary>
    /// Aggregated statistics about an issue model.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse
    {
        /// <summary>
        /// Number of conversations the issue model has analyzed at this point in time.
        /// </summary>
        public readonly string AnalyzedConversationsCount;
        /// <summary>
        /// Statistics on each issue. Key is the issue's resource name.
        /// </summary>
        public readonly ImmutableDictionary<string, string> IssueStats;
        /// <summary>
        /// Number of analyzed conversations for which no issue was applicable at this point in time.
        /// </summary>
        public readonly string UnclassifiedConversationsCount;

        [OutputConstructor]
        private GoogleCloudContactcenterinsightsV1IssueModelLabelStatsResponse(
            string analyzedConversationsCount,

            ImmutableDictionary<string, string> issueStats,

            string unclassifiedConversationsCount)
        {
            AnalyzedConversationsCount = analyzedConversationsCount;
            IssueStats = issueStats;
            UnclassifiedConversationsCount = unclassifiedConversationsCount;
        }
    }
}