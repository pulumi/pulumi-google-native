// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CertificateManager.V1
{
    public static class GetCertificateIssuanceConfig
    {
        /// <summary>
        /// Gets details of a single CertificateIssuanceConfig.
        /// </summary>
        public static Task<GetCertificateIssuanceConfigResult> InvokeAsync(GetCertificateIssuanceConfigArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateIssuanceConfigResult>("google-native:certificatemanager/v1:getCertificateIssuanceConfig", args ?? new GetCertificateIssuanceConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single CertificateIssuanceConfig.
        /// </summary>
        public static Output<GetCertificateIssuanceConfigResult> Invoke(GetCertificateIssuanceConfigInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateIssuanceConfigResult>("google-native:certificatemanager/v1:getCertificateIssuanceConfig", args ?? new GetCertificateIssuanceConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateIssuanceConfigArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateIssuanceConfigId", required: true)]
        public string CertificateIssuanceConfigId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateIssuanceConfigArgs()
        {
        }
        public static new GetCertificateIssuanceConfigArgs Empty => new GetCertificateIssuanceConfigArgs();
    }

    public sealed class GetCertificateIssuanceConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("certificateIssuanceConfigId", required: true)]
        public Input<string> CertificateIssuanceConfigId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCertificateIssuanceConfigInvokeArgs()
        {
        }
        public static new GetCertificateIssuanceConfigInvokeArgs Empty => new GetCertificateIssuanceConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateIssuanceConfigResult
    {
        /// <summary>
        /// The CA that issues the workload certificate. It includes the CA address, type, authentication to CA service, etc.
        /// </summary>
        public readonly Outputs.CertificateAuthorityConfigResponse CertificateAuthorityConfig;
        /// <summary>
        /// The creation timestamp of a CertificateIssuanceConfig.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// One or more paragraphs of text description of a CertificateIssuanceConfig.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The key algorithm to use when generating the private key.
        /// </summary>
        public readonly string KeyAlgorithm;
        /// <summary>
        /// Set of labels associated with a CertificateIssuanceConfig.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Workload certificate lifetime requested.
        /// </summary>
        public readonly string Lifetime;
        /// <summary>
        /// A user-defined name of the certificate issuance config. CertificateIssuanceConfig names must be unique globally and match pattern `projects/*/locations/*/certificateIssuanceConfigs/*`.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the percentage of elapsed time of the certificate lifetime to wait before renewing the certificate. Must be a number between 1-99, inclusive.
        /// </summary>
        public readonly int RotationWindowPercentage;
        /// <summary>
        /// The last update timestamp of a CertificateIssuanceConfig.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetCertificateIssuanceConfigResult(
            Outputs.CertificateAuthorityConfigResponse certificateAuthorityConfig,

            string createTime,

            string description,

            string keyAlgorithm,

            ImmutableDictionary<string, string> labels,

            string lifetime,

            string name,

            int rotationWindowPercentage,

            string updateTime)
        {
            CertificateAuthorityConfig = certificateAuthorityConfig;
            CreateTime = createTime;
            Description = description;
            KeyAlgorithm = keyAlgorithm;
            Labels = labels;
            Lifetime = lifetime;
            Name = name;
            RotationWindowPercentage = rotationWindowPercentage;
            UpdateTime = updateTime;
        }
    }
}
