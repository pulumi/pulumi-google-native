// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// Configuration and status of a managed SSL certificate.
    /// </summary>
    [OutputType]
    public sealed class SslCertificateManagedSslCertificateResponse
    {
        /// <summary>
        /// [Output only] Detailed statuses of the domains specified for managed certificate resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string> DomainStatus;
        /// <summary>
        /// The domains for which a managed SSL certificate will be generated. Each Google-managed SSL certificate supports up to the [maximum number of domains per Google-managed SSL certificate](/load-balancing/docs/quotas#ssl_certificates).
        /// </summary>
        public readonly ImmutableArray<string> Domains;
        /// <summary>
        /// [Output only] Status of the managed certificate resource.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private SslCertificateManagedSslCertificateResponse(
            ImmutableDictionary<string, string> domainStatus,

            ImmutableArray<string> domains,

            string status)
        {
            DomainStatus = domainStatus;
            Domains = domains;
            Status = status;
        }
    }
}
