// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha
{
    /// <summary>
    /// Creates a SslCertificate resource in the specified project and region using the data included in the request
    /// </summary>
    [GoogleNativeResourceType("google-native:compute/alpha:RegionSslCertificate")]
    public partial class RegionSslCertificate : Pulumi.CustomResource
    {
        /// <summary>
        /// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        /// </summary>
        [Output("certificate")]
        public Output<string> Certificate { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Expire time of the certificate. RFC3339
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Type of the resource. Always compute#sslCertificate for SSL certificates.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Configuration and status of a managed SSL certificate.
        /// </summary>
        [Output("managed")]
        public Output<Outputs.SslCertificateManagedSslCertificateResponse> Managed { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        /// </summary>
        [Output("privateKey")]
        public Output<string> PrivateKey { get; private set; } = null!;

        /// <summary>
        /// URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// [Output only] Server-defined URL for the resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL for this resource with the resource id.
        /// </summary>
        [Output("selfLinkWithId")]
        public Output<string> SelfLinkWithId { get; private set; } = null!;

        /// <summary>
        /// Configuration and status of a self-managed SSL certificate.
        /// </summary>
        [Output("selfManaged")]
        public Output<Outputs.SslCertificateSelfManagedSslCertificateResponse> SelfManaged { get; private set; } = null!;

        /// <summary>
        /// Domains associated with the certificate via Subject Alternative Name.
        /// </summary>
        [Output("subjectAlternativeNames")]
        public Output<ImmutableArray<string>> SubjectAlternativeNames { get; private set; } = null!;

        /// <summary>
        /// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RegionSslCertificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionSslCertificate(string name, RegionSslCertificateArgs args, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:RegionSslCertificate", name, args ?? new RegionSslCertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionSslCertificate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:compute/alpha:RegionSslCertificate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegionSslCertificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionSslCertificate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegionSslCertificate(name, id, options);
        }
    }

    public sealed class RegionSslCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
        /// </summary>
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Configuration and status of a managed SSL certificate.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.SslCertificateManagedSslCertificateArgs>? Managed { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
        /// </summary>
        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// Configuration and status of a self-managed SSL certificate.
        /// </summary>
        [Input("selfManaged")]
        public Input<Inputs.SslCertificateSelfManagedSslCertificateArgs>? SelfManaged { get; set; }

        /// <summary>
        /// (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.RegionSslCertificateType>? Type { get; set; }

        public RegionSslCertificateArgs()
        {
        }
    }
}