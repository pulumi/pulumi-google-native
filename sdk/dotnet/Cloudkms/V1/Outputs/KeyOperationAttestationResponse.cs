// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Cloudkms.V1.Outputs
{

    /// <summary>
    /// Contains an HSM-generated attestation about a key operation. For more information, see [Verifying attestations] (https://cloud.google.com/kms/docs/attest-key).
    /// </summary>
    [OutputType]
    public sealed class KeyOperationAttestationResponse
    {
        /// <summary>
        /// The certificate chains needed to validate the attestation
        /// </summary>
        public readonly Outputs.CertificateChainsResponse CertChains;
        /// <summary>
        /// The attestation data provided by the HSM when the key operation was performed.
        /// </summary>
        public readonly string Content;
        /// <summary>
        /// The format of the attestation data.
        /// </summary>
        public readonly string Format;

        [OutputConstructor]
        private KeyOperationAttestationResponse(
            Outputs.CertificateChainsResponse certChains,

            string content,

            string format)
        {
            CertChains = certChains;
            Content = content;
            Format = format;
        }
    }
}