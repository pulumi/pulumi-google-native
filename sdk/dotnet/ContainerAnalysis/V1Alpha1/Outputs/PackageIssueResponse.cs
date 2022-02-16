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
    /// This message wraps a location affected by a vulnerability and its associated fix (if one is available).
    /// </summary>
    [OutputType]
    public sealed class PackageIssueResponse
    {
        /// <summary>
        /// The location of the vulnerability.
        /// </summary>
        public readonly Outputs.VulnerabilityLocationResponse AffectedLocation;
        /// <summary>
        /// The distro or language system assigned severity for this vulnerability when that is available and note provider assigned severity when distro or language system has not yet assigned a severity for this vulnerability.
        /// </summary>
        public readonly string EffectiveSeverity;
        /// <summary>
        /// The location of the available fix for vulnerability.
        /// </summary>
        public readonly Outputs.VulnerabilityLocationResponse FixedLocation;
        /// <summary>
        /// The type of package (e.g. OS, MAVEN, GO).
        /// </summary>
        public readonly string PackageType;
        public readonly string SeverityName;

        [OutputConstructor]
        private PackageIssueResponse(
            Outputs.VulnerabilityLocationResponse affectedLocation,

            string effectiveSeverity,

            Outputs.VulnerabilityLocationResponse fixedLocation,

            string packageType,

            string severityName)
        {
            AffectedLocation = affectedLocation;
            EffectiveSeverity = effectiveSeverity;
            FixedLocation = fixedLocation;
            PackageType = packageType;
            SeverityName = severityName;
        }
    }
}