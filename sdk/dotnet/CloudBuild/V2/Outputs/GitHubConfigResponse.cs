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
    /// Configuration for connections to github.com.
    /// </summary>
    [OutputType]
    public sealed class GitHubConfigResponse
    {
        /// <summary>
        /// GitHub App installation id.
        /// </summary>
        public readonly string AppInstallationId;
        /// <summary>
        /// OAuth credential of the account that authorized the Cloud Build GitHub App. It is recommended to use a robot account instead of a human user account. The OAuth token must be tied to the Cloud Build GitHub App.
        /// </summary>
        public readonly Outputs.OAuthCredentialResponse AuthorizerCredential;

        [OutputConstructor]
        private GitHubConfigResponse(
            string appInstallationId,

            Outputs.OAuthCredentialResponse authorizerCredential)
        {
            AppInstallationId = appInstallationId;
            AuthorizerCredential = authorizerCredential;
        }
    }
}
