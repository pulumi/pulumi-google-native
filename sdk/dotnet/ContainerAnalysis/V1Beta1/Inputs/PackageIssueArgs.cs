// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// This message wraps a location affected by a vulnerability and its associated fix (if one is available).
    /// </summary>
    public sealed class PackageIssueArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the vulnerability.
        /// </summary>
        [Input("affectedLocation", required: true)]
        public Input<Inputs.VulnerabilityLocationArgs> AffectedLocation { get; set; } = null!;

        /// <summary>
        /// The location of the available fix for vulnerability.
        /// </summary>
        [Input("fixedLocation")]
        public Input<Inputs.VulnerabilityLocationArgs>? FixedLocation { get; set; }

        /// <summary>
        /// The type of package (e.g. OS, MAVEN, GO).
        /// </summary>
        [Input("packageType")]
        public Input<string>? PackageType { get; set; }

        public PackageIssueArgs()
        {
        }
    }
}