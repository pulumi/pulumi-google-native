// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.IdentityToolkit.V2
{
    /// <summary>
    /// Create an Oidc Idp configuration for an Identity Toolkit project.
    /// </summary>
    [GoogleNativeResourceType("google-native:identitytoolkit/v2:OauthIdpConfig")]
    public partial class OauthIdpConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The client id of an OAuth client.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client secret of the OAuth client, to enable OIDC code flow.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// The config's display name set by developers.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// True if allows the user to sign in with the provider.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// For OIDC Idps, the issuer identifier.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id to use for this config.
        /// </summary>
        [Output("oauthIdpConfigId")]
        public Output<string?> OauthIdpConfigId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
        /// </summary>
        [Output("responseType")]
        public Output<Outputs.GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeResponse> ResponseType { get; private set; } = null!;

        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a OauthIdpConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OauthIdpConfig(string name, OauthIdpConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:identitytoolkit/v2:OauthIdpConfig", name, args ?? new OauthIdpConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OauthIdpConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:identitytoolkit/v2:OauthIdpConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "project",
                    "tenantId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing OauthIdpConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OauthIdpConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new OauthIdpConfig(name, id, options);
        }
    }

    public sealed class OauthIdpConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client id of an OAuth client.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client secret of the OAuth client, to enable OIDC code flow.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// The config's display name set by developers.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// True if allows the user to sign in with the provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// For OIDC Idps, the issuer identifier.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// The name of the OAuthIdpConfig resource, for example: 'projects/my-awesome-project/oauthIdpConfigs/oauth-config-id'. Ignored during create requests.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id to use for this config.
        /// </summary>
        [Input("oauthIdpConfigId")]
        public Input<string>? OauthIdpConfigId { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The response type to request for in the OAuth authorization flow. You can set either `id_token` or `code` to true, but not both. Setting both types to be simultaneously true (`{code: true, id_token: true}`) is not yet supported.
        /// </summary>
        [Input("responseType")]
        public Input<Inputs.GoogleCloudIdentitytoolkitAdminV2OAuthResponseTypeArgs>? ResponseType { get; set; }

        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public OauthIdpConfigArgs()
        {
        }
        public static new OauthIdpConfigArgs Empty => new OauthIdpConfigArgs();
    }
}
