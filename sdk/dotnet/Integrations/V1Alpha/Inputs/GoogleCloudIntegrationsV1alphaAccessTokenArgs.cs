// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// The access token represents the authorization of a specific application to access specific parts of a user’s data.
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaAccessTokenArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access token encapsulating the security identity of a process or thread.
        /// </summary>
        [Input("accessToken")]
        public Input<string>? AccessToken { get; set; }

        /// <summary>
        /// The approximate time until the access token retrieved is valid.
        /// </summary>
        [Input("accessTokenExpireTime", required: true)]
        public Input<string> AccessTokenExpireTime { get; set; } = null!;

        /// <summary>
        /// If the access token will expire, use the refresh token to obtain another access token.
        /// </summary>
        [Input("refreshToken")]
        public Input<string>? RefreshToken { get; set; }

        /// <summary>
        /// The approximate time until the refresh token retrieved is valid.
        /// </summary>
        [Input("refreshTokenExpireTime")]
        public Input<string>? RefreshTokenExpireTime { get; set; }

        /// <summary>
        /// Only support "bearer" token in v1 as bearer token is the predominant type used with OAuth 2.0.
        /// </summary>
        [Input("tokenType")]
        public Input<string>? TokenType { get; set; }

        public GoogleCloudIntegrationsV1alphaAccessTokenArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaAccessTokenArgs Empty => new GoogleCloudIntegrationsV1alphaAccessTokenArgs();
    }
}