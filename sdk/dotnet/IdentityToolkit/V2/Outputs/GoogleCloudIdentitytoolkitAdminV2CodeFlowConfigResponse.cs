// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IdentityToolkit.V2.Outputs
{

    /// <summary>
    /// Additional config for Apple for code flow.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIdentitytoolkitAdminV2CodeFlowConfigResponse
    {
        /// <summary>
        /// Key ID for the private key.
        /// </summary>
        public readonly string KeyId;
        /// <summary>
        /// Private key used for signing the client secret JWT.
        /// </summary>
        public readonly string PrivateKey;
        /// <summary>
        /// Apple Developer Team ID.
        /// </summary>
        public readonly string TeamId;

        [OutputConstructor]
        private GoogleCloudIdentitytoolkitAdminV2CodeFlowConfigResponse(
            string keyId,

            string privateKey,

            string teamId)
        {
            KeyId = keyId;
            PrivateKey = privateKey;
            TeamId = teamId;
        }
    }
}