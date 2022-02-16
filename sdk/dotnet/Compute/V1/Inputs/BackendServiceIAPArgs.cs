// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Identity-Aware Proxy
    /// </summary>
    public sealed class BackendServiceIAPArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the serving infrastructure will authenticate and authorize all incoming requests. If true, the oauth2ClientId and oauth2ClientSecret fields must be non-empty.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// OAuth2 client ID to use for the authentication flow.
        /// </summary>
        [Input("oauth2ClientId")]
        public Input<string>? Oauth2ClientId { get; set; }

        /// <summary>
        /// OAuth2 client secret to use for the authentication flow. For security reasons, this value cannot be retrieved via the API. Instead, the SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field. @InputOnly
        /// </summary>
        [Input("oauth2ClientSecret")]
        public Input<string>? Oauth2ClientSecret { get; set; }

        public BackendServiceIAPArgs()
        {
        }
    }
}