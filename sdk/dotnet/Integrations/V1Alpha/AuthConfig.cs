// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha
{
    /// <summary>
    /// Creates an auth config record. Fetch corresponding credentials for specific auth types, e.g. access token for OAuth 2.0, JWT token for JWT. Encrypt the auth config with Cloud KMS and store the encrypted credentials in Spanner. Returns the encrypted auth config.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:integrations/v1alpha:AuthConfig")]
    public partial class AuthConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Certificate id for client certificate
        /// </summary>
        [Output("certificateId")]
        public Output<string> CertificateId { get; private set; } = null!;

        /// <summary>
        /// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
        /// </summary>
        [Output("clientCertificateEncryptedPrivateKey")]
        public Output<string?> ClientCertificateEncryptedPrivateKey { get; private set; } = null!;

        /// <summary>
        /// 'passphrase' should be left unset if private key is not encrypted. Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
        /// </summary>
        [Output("clientCertificatePassphrase")]
        public Output<string?> ClientCertificatePassphrase { get; private set; } = null!;

        /// <summary>
        /// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
        /// </summary>
        [Output("clientCertificateSslCertificate")]
        public Output<string?> ClientCertificateSslCertificate { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the auth config is created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Output("creatorEmail")]
        public Output<string> CreatorEmail { get; private set; } = null!;

        /// <summary>
        /// Credential type of the encrypted credential.
        /// </summary>
        [Output("credentialType")]
        public Output<string> CredentialType { get; private set; } = null!;

        /// <summary>
        /// Raw auth credentials.
        /// </summary>
        [Output("decryptedCredential")]
        public Output<Outputs.GoogleCloudIntegrationsV1alphaCredentialResponse> DecryptedCredential { get; private set; } = null!;

        /// <summary>
        /// A description of the auth config.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the auth config.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
        /// </summary>
        [Output("encryptedCredential")]
        public Output<string> EncryptedCredential { get; private set; } = null!;

        /// <summary>
        /// User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
        /// </summary>
        [Output("expiryNotificationDuration")]
        public Output<ImmutableArray<string>> ExpiryNotificationDuration { get; private set; } = null!;

        /// <summary>
        /// The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Output("lastModifierEmail")]
        public Output<string> LastModifierEmail { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
        /// </summary>
        [Output("overrideValidTime")]
        public Output<string> OverrideValidTime { get; private set; } = null!;

        [Output("productId")]
        public Output<string> ProductId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The reason / details of the current status.
        /// </summary>
        [Output("reason")]
        public Output<string> Reason { get; private set; } = null!;

        /// <summary>
        /// The status of the auth config.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the auth config is modified.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
        /// </summary>
        [Output("validTime")]
        public Output<string> ValidTime { get; private set; } = null!;

        /// <summary>
        /// The visibility of the auth config.
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;


        /// <summary>
        /// Create a AuthConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthConfig(string name, AuthConfigArgs args, CustomResourceOptions? options = null)
            : base("google-native:integrations/v1alpha:AuthConfig", name, args ?? new AuthConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:integrations/v1alpha:AuthConfig", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "productId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AuthConfig(name, id, options);
        }
    }

    public sealed class AuthConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Certificate id for client certificate
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        /// <summary>
        /// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
        /// </summary>
        [Input("clientCertificateEncryptedPrivateKey")]
        public Input<string>? ClientCertificateEncryptedPrivateKey { get; set; }

        /// <summary>
        /// 'passphrase' should be left unset if private key is not encrypted. Note that 'passphrase' is not the password for web server, but an extra layer of security to protected private key.
        /// </summary>
        [Input("clientCertificatePassphrase")]
        public Input<string>? ClientCertificatePassphrase { get; set; }

        /// <summary>
        /// The ssl certificate encoded in PEM format. This string must include the begin header and end footer lines. For example, -----BEGIN CERTIFICATE----- MIICTTCCAbagAwIBAgIJAPT0tSKNxan/MA0GCSqGSIb3DQEBCwUAMCoxFzAVBgNV BAoTDkdvb2dsZSBURVNUSU5HMQ8wDQYDVQQDEwZ0ZXN0Q0EwHhcNMTUwMTAxMDAw MDAwWhcNMjUwMTAxMDAwMDAwWjAuMRcwFQYDVQQKEw5Hb29nbGUgVEVTVElORzET MBEGA1UEAwwKam9lQGJhbmFuYTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA vDYFgMgxi5W488d9J7UpCInl0NXmZQpJDEHE4hvkaRlH7pnC71H0DLt0/3zATRP1 JzY2+eqBmbGl4/sgZKYv8UrLnNyQNUTsNx1iZAfPUflf5FwgVsai8BM0pUciq1NB xD429VFcrGZNucvFLh72RuRFIKH8WUpiK/iZNFkWhZ0CAwEAAaN3MHUwDgYDVR0P AQH/BAQDAgWgMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAMBgNVHRMB Af8EAjAAMBkGA1UdDgQSBBCVgnFBCWgL/iwCqnGrhTPQMBsGA1UdIwQUMBKAEKey Um2o4k2WiEVA0ldQvNYwDQYJKoZIhvcNAQELBQADgYEAYK986R4E3L1v+Q6esBtW JrUwA9UmJRSQr0N5w3o9XzarU37/bkjOP0Fw0k/A6Vv1n3vlciYfBFaBIam1qRHr 5dMsYf4CZS6w50r7hyzqyrwDoyNxkLnd2PdcHT/sym1QmflsjEs7pejtnohO6N2H wQW6M0H7Zt8claGRla4fKkg= -----END CERTIFICATE-----
        /// </summary>
        [Input("clientCertificateSslCertificate")]
        public Input<string>? ClientCertificateSslCertificate { get; set; }

        /// <summary>
        /// The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Input("creatorEmail")]
        public Input<string>? CreatorEmail { get; set; }

        /// <summary>
        /// Credential type of the encrypted credential.
        /// </summary>
        [Input("credentialType")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.AuthConfigCredentialType>? CredentialType { get; set; }

        /// <summary>
        /// Raw auth credentials.
        /// </summary>
        [Input("decryptedCredential")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaCredentialArgs>? DecryptedCredential { get; set; }

        /// <summary>
        /// A description of the auth config.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the auth config.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
        /// </summary>
        [Input("encryptedCredential")]
        public Input<string>? EncryptedCredential { get; set; }

        [Input("expiryNotificationDuration")]
        private InputList<string>? _expiryNotificationDuration;

        /// <summary>
        /// User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
        /// </summary>
        public InputList<string> ExpiryNotificationDuration
        {
            get => _expiryNotificationDuration ?? (_expiryNotificationDuration = new InputList<string>());
            set => _expiryNotificationDuration = value;
        }

        /// <summary>
        /// The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        /// </summary>
        [Input("lastModifierEmail")]
        public Input<string>? LastModifierEmail { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
        /// </summary>
        [Input("overrideValidTime")]
        public Input<string>? OverrideValidTime { get; set; }

        [Input("productId", required: true)]
        public Input<string> ProductId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The reason / details of the current status.
        /// </summary>
        [Input("reason")]
        public Input<string>? Reason { get; set; }

        /// <summary>
        /// The status of the auth config.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.AuthConfigState>? State { get; set; }

        /// <summary>
        /// The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
        /// </summary>
        [Input("validTime")]
        public Input<string>? ValidTime { get; set; }

        /// <summary>
        /// The visibility of the auth config.
        /// </summary>
        [Input("visibility")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.AuthConfigVisibility>? Visibility { get; set; }

        public AuthConfigArgs()
        {
        }
        public static new AuthConfigArgs Empty => new AuthConfigArgs();
    }
}