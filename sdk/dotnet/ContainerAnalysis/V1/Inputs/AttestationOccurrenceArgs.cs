// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// Occurrence that represents a single "attestation". The authenticity of an attestation can be verified using the attached signature. If the verifier trusts the public key of the signer, then verifying the signature is sufficient to establish trust. In this circumstance, the authority to which this attestation is attached is primarily useful for lookup (how to find this attestation if you already know the authority and artifact to be verified) and intent (for which authority this attestation was intended to sign.
    /// </summary>
    public sealed class AttestationOccurrenceArgs : global::Pulumi.ResourceArgs
    {
        [Input("jwts")]
        private InputList<Inputs.JwtArgs>? _jwts;

        /// <summary>
        /// One or more JWTs encoding a self-contained attestation. Each JWT encodes the payload that it verifies within the JWT itself. Verifier implementation SHOULD ignore the `serialized_payload` field when verifying these JWTs. If only JWTs are present on this AttestationOccurrence, then the `serialized_payload` SHOULD be left empty. Each JWT SHOULD encode a claim specific to the `resource_uri` of this Occurrence, but this is not validated by Grafeas metadata API implementations. The JWT itself is opaque to Grafeas.
        /// </summary>
        public InputList<Inputs.JwtArgs> Jwts
        {
            get => _jwts ?? (_jwts = new InputList<Inputs.JwtArgs>());
            set => _jwts = value;
        }

        /// <summary>
        /// The serialized payload that is verified by one or more `signatures`.
        /// </summary>
        [Input("serializedPayload", required: true)]
        public Input<string> SerializedPayload { get; set; } = null!;

        [Input("signatures")]
        private InputList<Inputs.SignatureArgs>? _signatures;

        /// <summary>
        /// One or more signatures over `serialized_payload`. Verifier implementations should consider this attestation message verified if at least one `signature` verifies `serialized_payload`. See `Signature` in common.proto for more details on signature structure and verification.
        /// </summary>
        public InputList<Inputs.SignatureArgs> Signatures
        {
            get => _signatures ?? (_signatures = new InputList<Inputs.SignatureArgs>());
            set => _signatures = value;
        }

        public AttestationOccurrenceArgs()
        {
        }
        public static new AttestationOccurrenceArgs Empty => new AttestationOccurrenceArgs();
    }
}
