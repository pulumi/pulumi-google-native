// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Specifies the mapping between the response code that will be returned along with the custom error content and the response code returned by the backend service.
    /// </summary>
    [OutputType]
    public sealed class CustomErrorResponsePolicyCustomErrorResponseRuleResponse
    {
        /// <summary>
        /// Valid values include: - A number between 400 and 599: For example 401 or 503, in which case the load balancer applies the policy if the error code exactly matches this value. - 5xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 500 to 599. - 4xx: Load Balancer will apply the policy if the backend service responds with any response code in the range of 400 to 499. Values must be unique within matchResponseCodes and across all errorResponseRules of CustomErrorResponsePolicy.
        /// </summary>
        public readonly ImmutableArray<string> MatchResponseCodes;
        /// <summary>
        /// The HTTP status code returned with the response containing the custom error content. If overrideResponseCode is not supplied, the same response code returned by the original backend bucket or backend service is returned to the client.
        /// </summary>
        public readonly int OverrideResponseCode;
        /// <summary>
        /// The full path to a file within backendBucket . For example: /errors/defaultError.html path must start with a leading slash. path cannot have trailing slashes. If the file is not available in backendBucket or the load balancer cannot reach the BackendBucket, a simple Not Found Error is returned to the client. The value must be from 1 to 1024 characters
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private CustomErrorResponsePolicyCustomErrorResponseRuleResponse(
            ImmutableArray<string> matchResponseCodes,

            int overrideResponseCode,

            string path)
        {
            MatchResponseCodes = matchResponseCodes;
            OverrideResponseCode = overrideResponseCode;
            Path = path;
        }
    }
}