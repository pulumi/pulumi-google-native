// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V1.Inputs
{

    /// <summary>
    /// Describes the policy in case of function's execution failure. If empty, then defaults to ignoring failures (i.e. not retrying them).
    /// </summary>
    public sealed class FailurePolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If specified, then the function will be retried in case of a failure.
        /// </summary>
        [Input("retry")]
        public Input<Inputs.RetryArgs>? Retry { get; set; }

        public FailurePolicyArgs()
        {
        }
    }
}