// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Outputs
{

    /// <summary>
    /// A status to accept. Either a status code class like "2xx", or an integer status code like "200".
    /// </summary>
    [OutputType]
    public sealed class ResponseStatusCodeResponse
    {
        /// <summary>
        /// A class of status codes to accept.
        /// </summary>
        public readonly string StatusClass;
        /// <summary>
        /// A status code to accept.
        /// </summary>
        public readonly int StatusValue;

        [OutputConstructor]
        private ResponseStatusCodeResponse(
            string statusClass,

            int statusValue)
        {
            StatusClass = statusClass;
            StatusValue = statusValue;
        }
    }
}