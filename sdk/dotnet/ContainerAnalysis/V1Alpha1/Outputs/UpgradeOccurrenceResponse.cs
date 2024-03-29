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
    /// An Upgrade Occurrence represents that a specific resource_url could install a specific upgrade. This presence is supplied via local sources (i.e. it is present in the mirror and the running system has noticed its availability).
    /// </summary>
    [OutputType]
    public sealed class UpgradeOccurrenceResponse
    {
        /// <summary>
        /// Metadata about the upgrade for available for the specific operating system for the resource_url. This allows efficient filtering, as well as making it easier to use the occurrence.
        /// </summary>
        public readonly Outputs.UpgradeDistributionResponse Distribution;
        /// <summary>
        /// Required - The package this Upgrade is for.
        /// </summary>
        public readonly string Package;
        /// <summary>
        /// Required - The version of the package in a machine + human readable form.
        /// </summary>
        public readonly Outputs.VersionResponse ParsedVersion;

        [OutputConstructor]
        private UpgradeOccurrenceResponse(
            Outputs.UpgradeDistributionResponse distribution,

            string package,

            Outputs.VersionResponse parsedVersion)
        {
            Distribution = distribution;
            Package = package;
            ParsedVersion = parsedVersion;
        }
    }
}
