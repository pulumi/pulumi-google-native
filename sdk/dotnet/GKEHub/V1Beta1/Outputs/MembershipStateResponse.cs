// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Beta1.Outputs
{

    /// <summary>
    /// State of the Membership resource.
    /// </summary>
    [OutputType]
    public sealed class MembershipStateResponse
    {
        /// <summary>
        /// The current state of the Membership resource.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// This field is never set by the Hub Service.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// This field is never set by the Hub Service.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private MembershipStateResponse(
            string code,

            string description,

            string updateTime)
        {
            Code = code;
            Description = description;
            UpdateTime = updateTime;
        }
    }
}