// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Outputs
{

    /// <summary>
    /// Details of a build occurrence.
    /// </summary>
    [OutputType]
    public sealed class BuildOccurrenceResponse
    {
        /// <summary>
        /// In-toto Statement representation as defined in spec. The intoto_statement can contain any type of provenance. The serialized payload of the statement can be stored and signed in the Occurrence's envelope.
        /// </summary>
        public readonly Outputs.InTotoStatementResponse IntotoStatement;
        /// <summary>
        /// The actual provenance for the build.
        /// </summary>
        public readonly Outputs.BuildProvenanceResponse Provenance;
        /// <summary>
        /// Serialized JSON representation of the provenance, used in generating the build signature in the corresponding build note. After verifying the signature, `provenance_bytes` can be unmarshalled and compared to the provenance to confirm that it is unchanged. A base64-encoded string representation of the provenance bytes is used for the signature in order to interoperate with openssl which expects this format for signature verification. The serialized form is captured both to avoid ambiguity in how the provenance is marshalled to json as well to prevent incompatibilities with future changes.
        /// </summary>
        public readonly string ProvenanceBytes;

        [OutputConstructor]
        private BuildOccurrenceResponse(
            Outputs.InTotoStatementResponse intotoStatement,

            Outputs.BuildProvenanceResponse provenance,

            string provenanceBytes)
        {
            IntotoStatement = intotoStatement;
            Provenance = provenance;
            ProvenanceBytes = provenanceBytes;
        }
    }
}