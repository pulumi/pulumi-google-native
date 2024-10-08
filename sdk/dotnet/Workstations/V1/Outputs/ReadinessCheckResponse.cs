// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1.Outputs
{

    /// <summary>
    /// A readiness check to be performed on a workstation.
    /// </summary>
    [OutputType]
    public sealed class ReadinessCheckResponse
    {
        /// <summary>
        /// Optional. Path to which the request should be sent.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Optional. Port to which the request should be sent.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private ReadinessCheckResponse(
            string path,

            int port)
        {
            Path = path;
            Port = port;
        }
    }
}
