// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V1.Outputs
{

    /// <summary>
    /// Describes the policy in case of function's execution failure. If empty, then defaults to ignoring failures (i.e. not retrying them).
    /// </summary>
    [OutputType]
    public sealed class FailurePolicyResponse
    {
        /// <summary>
        /// If specified, then the function will be retried in case of a failure.
        /// </summary>
        public readonly Outputs.RetryResponse Retry;

        [OutputConstructor]
        private FailurePolicyResponse(Outputs.RetryResponse retry)
        {
            Retry = retry;
        }
    }
}