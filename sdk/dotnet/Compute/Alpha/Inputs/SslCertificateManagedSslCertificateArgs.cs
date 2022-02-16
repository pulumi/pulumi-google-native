// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Configuration and status of a managed SSL certificate.
    /// </summary>
    public sealed class SslCertificateManagedSslCertificateArgs : Pulumi.ResourceArgs
    {
        [Input("domains")]
        private InputList<string>? _domains;

        /// <summary>
        /// The domains for which a managed SSL certificate will be generated. Each Google-managed SSL certificate supports up to the [maximum number of domains per Google-managed SSL certificate](/load-balancing/docs/quotas#ssl_certificates).
        /// </summary>
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        public SslCertificateManagedSslCertificateArgs()
        {
        }
    }
}