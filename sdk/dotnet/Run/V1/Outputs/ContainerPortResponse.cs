// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Outputs
{

    /// <summary>
    /// ContainerPort represents a network port in a single container.
    /// </summary>
    [OutputType]
    public sealed class ContainerPortResponse
    {
        /// <summary>
        /// Port number the container listens on. If present, this must be a valid port number, 0 &lt; x &lt; 65536. If not present, it will default to port 8080. For more information, see https://cloud.google.com/run/docs/container-contract#port
        /// </summary>
        public readonly int ContainerPort;
        /// <summary>
        /// If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Protocol for port. Must be "TCP". Defaults to "TCP".
        /// </summary>
        public readonly string Protocol;

        [OutputConstructor]
        private ContainerPortResponse(
            int containerPort,

            string name,

            string protocol)
        {
            ContainerPort = containerPort;
            Name = name;
            Protocol = protocol;
        }
    }
}
