// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudIdentity.V1Beta1.Outputs
{

    /// <summary>
    /// Resource representing the Endpoint Verification-specific attributes of a Device. https://cloud.google.com/endpoint-verification/docs/overview
    /// </summary>
    [OutputType]
    public sealed class EndpointVerificationSpecificAttributesResponse
    {
        /// <summary>
        /// Details of certificates.
        /// </summary>
        public readonly ImmutableArray<Outputs.CertificateAttributesResponse> CertificateAttributes;

        [OutputConstructor]
        private EndpointVerificationSpecificAttributesResponse(ImmutableArray<Outputs.CertificateAttributesResponse> certificateAttributes)
        {
            CertificateAttributes = certificateAttributes;
        }
    }
}
