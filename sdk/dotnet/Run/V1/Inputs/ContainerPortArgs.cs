// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Inputs
{

    /// <summary>
    /// ContainerPort represents a network port in a single container.
    /// </summary>
    public sealed class ContainerPortArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port number the container listens on. If present, this must be a valid port number, 0 &lt; x &lt; 65536. If not present, it will default to port 8080. For more information, see https://cloud.google.com/run/docs/container-contract#port
        /// </summary>
        [Input("containerPort")]
        public Input<int>? ContainerPort { get; set; }

        /// <summary>
        /// If specified, used to specify which protocol to use. Allowed values are "http1" and "h2c".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol for port. Must be "TCP". Defaults to "TCP".
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        public ContainerPortArgs()
        {
        }
        public static new ContainerPortArgs Empty => new ContainerPortArgs();
    }
}
