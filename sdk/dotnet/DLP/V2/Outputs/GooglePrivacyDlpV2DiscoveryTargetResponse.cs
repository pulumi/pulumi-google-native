// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// Target used to match against for Discovery.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2DiscoveryTargetResponse
    {
        /// <summary>
        /// BigQuery target for Discovery. The first target to match a table will be the one applied.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2BigQueryDiscoveryTargetResponse BigQueryTarget;

        [OutputConstructor]
        private GooglePrivacyDlpV2DiscoveryTargetResponse(Outputs.GooglePrivacyDlpV2BigQueryDiscoveryTargetResponse bigQueryTarget)
        {
            BigQueryTarget = bigQueryTarget;
        }
    }
}