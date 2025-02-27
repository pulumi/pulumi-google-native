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
    /// Configuration for OIDC Auth flow.
    /// </summary>
    [OutputType]
    public sealed class IdentityServiceOidcConfigResponse
    {
        /// <summary>
        /// PEM-encoded CA for OIDC provider.
        /// </summary>
        public readonly string CertificateAuthorityData;
        /// <summary>
        /// ID for OIDC client application.
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// Input only. Unencrypted OIDC client secret will be passed to the GKE Hub CLH.
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// Flag to denote if reverse proxy is used to connect to auth provider. This flag should be set to true when provider is not reachable by Google Cloud Console.
        /// </summary>
        public readonly bool DeployCloudConsoleProxy;
        /// <summary>
        /// Enable access token.
        /// </summary>
        public readonly bool EnableAccessToken;
        /// <summary>
        /// Encrypted OIDC Client secret
        /// </summary>
        public readonly string EncryptedClientSecret;
        /// <summary>
        /// Comma-separated list of key-value pairs.
        /// </summary>
        public readonly string ExtraParams;
        /// <summary>
        /// Prefix to prepend to group name.
        /// </summary>
        public readonly string GroupPrefix;
        /// <summary>
        /// Claim in OIDC ID token that holds group information.
        /// </summary>
        public readonly string GroupsClaim;
        /// <summary>
        /// URI for the OIDC provider. This should point to the level below .well-known/openid-configuration.
        /// </summary>
        public readonly string IssuerUri;
        /// <summary>
        /// Registered redirect uri to redirect users going through OAuth flow using kubectl plugin.
        /// </summary>
        public readonly string KubectlRedirectUri;
        /// <summary>
        /// Comma-separated list of identifiers.
        /// </summary>
        public readonly string Scopes;
        /// <summary>
        /// Claim in OIDC ID token that holds username.
        /// </summary>
        public readonly string UserClaim;
        /// <summary>
        /// Prefix to prepend to user name.
        /// </summary>
        public readonly string UserPrefix;

        [OutputConstructor]
        private IdentityServiceOidcConfigResponse(
            string certificateAuthorityData,

            string clientId,

            string clientSecret,

            bool deployCloudConsoleProxy,

            bool enableAccessToken,

            string encryptedClientSecret,

            string extraParams,

            string groupPrefix,

            string groupsClaim,

            string issuerUri,

            string kubectlRedirectUri,

            string scopes,

            string userClaim,

            string userPrefix)
        {
            CertificateAuthorityData = certificateAuthorityData;
            ClientId = clientId;
            ClientSecret = clientSecret;
            DeployCloudConsoleProxy = deployCloudConsoleProxy;
            EnableAccessToken = enableAccessToken;
            EncryptedClientSecret = encryptedClientSecret;
            ExtraParams = extraParams;
            GroupPrefix = groupPrefix;
            GroupsClaim = groupsClaim;
            IssuerUri = issuerUri;
            KubectlRedirectUri = kubectlRedirectUri;
            Scopes = scopes;
            UserClaim = userClaim;
            UserPrefix = userPrefix;
        }
    }
}
