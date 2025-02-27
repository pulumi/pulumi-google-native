// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyRuleRateLimitOptionsThresholdResponse
    {
        /// <summary>
        /// Number of HTTP(S) requests for calculating the threshold.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// Interval over which the threshold is computed.
        /// </summary>
        public readonly int IntervalSec;

        [OutputConstructor]
        private SecurityPolicyRuleRateLimitOptionsThresholdResponse(
            int count,

            int intervalSec)
        {
            Count = count;
            IntervalSec = intervalSec;
        }
    }
}
