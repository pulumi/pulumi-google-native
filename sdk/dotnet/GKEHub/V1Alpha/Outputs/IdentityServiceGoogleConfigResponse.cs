// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// Configuration for the Google Plugin Auth flow.
    /// </summary>
    [OutputType]
    public sealed class IdentityServiceGoogleConfigResponse
    {
        /// <summary>
        /// Disable automatic configuration of Google Plugin on supported platforms.
        /// </summary>
        public readonly bool Disable;

        [OutputConstructor]
        private IdentityServiceGoogleConfigResponse(bool disable)
        {
            Disable = disable;
        }
    }
}