// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Alpha1.Outputs
{

    /// <summary>
    /// Assessment provides all information that is related to a single vulnerability for this product.
    /// </summary>
    [OutputType]
    public sealed class AssessmentResponse
    {
        /// <summary>
        /// Holds the MITRE standard Common Vulnerabilities and Exposures (CVE) tracking number for the vulnerability.
        /// </summary>
        public readonly string Cve;
        /// <summary>
        /// Contains information about the impact of this vulnerability, this will change with time.
        /// </summary>
        public readonly ImmutableArray<string> Impacts;
        /// <summary>
        /// Justification provides the justification when the state of the assessment if NOT_AFFECTED.
        /// </summary>
        public readonly Outputs.JustificationResponse Justification;
        /// <summary>
        /// A detailed description of this Vex.
        /// </summary>
        public readonly string LongDescription;
        /// <summary>
        /// Holds a list of references associated with this vulnerability item and assessment. These uris have additional information about the vulnerability and the assessment itself. E.g. Link to a document which details how this assessment concluded the state of this vulnerability.
        /// </summary>
        public readonly ImmutableArray<Outputs.URIResponse> RelatedUris;
        /// <summary>
        /// Specifies details on how to handle (and presumably, fix) a vulnerability.
        /// </summary>
        public readonly ImmutableArray<Outputs.RemediationResponse> Remediations;
        /// <summary>
        /// A one sentence description of this Vex.
        /// </summary>
        public readonly string ShortDescription;
        /// <summary>
        /// Provides the state of this Vulnerability assessment.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private AssessmentResponse(
            string cve,

            ImmutableArray<string> impacts,

            Outputs.JustificationResponse justification,

            string longDescription,

            ImmutableArray<Outputs.URIResponse> relatedUris,

            ImmutableArray<Outputs.RemediationResponse> remediations,

            string shortDescription,

            string state)
        {
            Cve = cve;
            Impacts = impacts;
            Justification = justification;
            LongDescription = longDescription;
            RelatedUris = relatedUris;
            Remediations = remediations;
            ShortDescription = shortDescription;
            State = state;
        }
    }
}