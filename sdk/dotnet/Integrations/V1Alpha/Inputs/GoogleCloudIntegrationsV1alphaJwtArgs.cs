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
    /// Represents JSON web token(JWT), which is a compact, URL-safe means of representing claims to be transferred between two parties, enabling the claims to be digitally signed or integrity protected.
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaJwtArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The token calculated by the header, payload and signature.
        /// </summary>
        [Input("jwt")]
        public Input<string>? Jwt { get; set; }

        /// <summary>
        /// Identifies which algorithm is used to generate the signature.
        /// </summary>
        [Input("jwtHeader")]
        public Input<string>? JwtHeader { get; set; }

        /// <summary>
        /// Contains a set of claims. The JWT specification defines seven Registered Claim Names which are the standard fields commonly included in tokens. Custom claims are usually also included, depending on the purpose of the token.
        /// </summary>
        [Input("jwtPayload")]
        public Input<string>? JwtPayload { get; set; }

        /// <summary>
        /// User's pre-shared secret to sign the token.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public GoogleCloudIntegrationsV1alphaJwtArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaJwtArgs Empty => new GoogleCloudIntegrationsV1alphaJwtArgs();
    }
}