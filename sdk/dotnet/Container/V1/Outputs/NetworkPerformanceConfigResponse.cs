// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Outputs
{

    /// <summary>
    /// Configuration of all network bandwidth tiers
    /// </summary>
    [OutputType]
    public sealed class NetworkPerformanceConfigResponse
    {
        /// <summary>
        /// Specifies the total network bandwidth tier for the NodePool.
        /// </summary>
        public readonly string TotalEgressBandwidthTier;

        [OutputConstructor]
        private NetworkPerformanceConfigResponse(string totalEgressBandwidthTier)
        {
            TotalEgressBandwidthTier = totalEgressBandwidthTier;
        }
    }
}