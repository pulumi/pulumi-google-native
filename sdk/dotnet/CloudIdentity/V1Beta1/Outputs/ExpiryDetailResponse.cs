// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// The `MembershipRole` expiry details.
    /// </summary>
    [OutputType]
    public sealed class ExpiryDetailResponse
    {
        /// <summary>
        /// The time at which the `MembershipRole` will expire.
        /// </summary>
        public readonly string ExpireTime;

        [OutputConstructor]
        private ExpiryDetailResponse(string expireTime)
        {
            ExpireTime = expireTime;
        }
    }
}