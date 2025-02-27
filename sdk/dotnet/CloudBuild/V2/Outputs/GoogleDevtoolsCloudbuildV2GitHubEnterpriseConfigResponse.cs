// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V2.Outputs
{

    /// <summary>
    /// Configuration for connections to an instance of GitHub Enterprise.
    /// </summary>
    [OutputType]
    public sealed class GoogleDevtoolsCloudbuildV2GitHubEnterpriseConfigResponse
    {
        /// <summary>
        /// API Key used for authentication of webhook events.
        /// </summary>
        public readonly string ApiKey;
        /// <summary>
        /// Id of the GitHub App created from the manifest.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// ID of the installation of the GitHub App.
        /// </summary>
        public readonly string AppInstallationId;
        /// <summary>
        /// The URL-friendly name of the GitHub App.
        /// </summary>
        public readonly string AppSlug;
        /// <summary>
        /// The URI of the GitHub Enterprise host this connection is for.
        /// </summary>
        public readonly string HostUri;
        /// <summary>
        /// SecretManager resource containing the private key of the GitHub App, formatted as `projects/*/secrets/*/versions/*`.
        /// </summary>
        public readonly string PrivateKeySecretVersion;
        /// <summary>
        /// GitHub Enterprise version installed at the host_uri.
        /// </summary>
        public readonly string ServerVersion;
        /// <summary>
        /// Configuration for using Service Directory to privately connect to a GitHub Enterprise server. This should only be set if the GitHub Enterprise server is hosted on-premises and not reachable by public internet. If this field is left empty, calls to the GitHub Enterprise server will be made over the public internet.
        /// </summary>
        public readonly Outputs.GoogleDevtoolsCloudbuildV2ServiceDirectoryConfigResponse ServiceDirectoryConfig;
        /// <summary>
        /// SSL certificate to use for requests to GitHub Enterprise.
        /// </summary>
        public readonly string SslCa;
        /// <summary>
        /// SecretManager resource containing the webhook secret of the GitHub App, formatted as `projects/*/secrets/*/versions/*`.
        /// </summary>
        public readonly string WebhookSecretSecretVersion;

        [OutputConstructor]
        private GoogleDevtoolsCloudbuildV2GitHubEnterpriseConfigResponse(
            string apiKey,

            string appId,

            string appInstallationId,

            string appSlug,

            string hostUri,

            string privateKeySecretVersion,

            string serverVersion,

            Outputs.GoogleDevtoolsCloudbuildV2ServiceDirectoryConfigResponse serviceDirectoryConfig,

            string sslCa,

            string webhookSecretSecretVersion)
        {
            ApiKey = apiKey;
            AppId = appId;
            AppInstallationId = appInstallationId;
            AppSlug = appSlug;
            HostUri = hostUri;
            PrivateKeySecretVersion = privateKeySecretVersion;
            ServerVersion = serverVersion;
            ServiceDirectoryConfig = serviceDirectoryConfig;
            SslCa = sslCa;
            WebhookSecretSecretVersion = webhookSecretSecretVersion;
        }
    }
}
