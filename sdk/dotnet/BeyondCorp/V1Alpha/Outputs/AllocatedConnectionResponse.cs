// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1Alpha.Outputs
{

    /// <summary>
    /// Allocated connection of the AppGateway.
    /// </summary>
    [OutputType]
    public sealed class AllocatedConnectionResponse
    {
        /// <summary>
        /// The ingress port of an allocated connection
        /// </summary>
        public readonly int IngressPort;
        /// <summary>
        /// The PSC uri of an allocated connection
        /// </summary>
        public readonly string PscUri;

        [OutputConstructor]
        private AllocatedConnectionResponse(
            int ingressPort,

            string pscUri)
        {
            IngressPort = ingressPort;
            PscUri = pscUri;
        }
    }
}