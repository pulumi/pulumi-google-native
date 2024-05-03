// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Outputs
{

    /// <summary>
    /// Contains the paths to the artifacts, relative to the URI, for a phase.
    /// </summary>
    [OutputType]
    public sealed class PhaseArtifactResponse
    {
        /// <summary>
        /// File path of the directory of rendered job manifests relative to the URI. This is only set if it is applicable.
        /// </summary>
        public readonly string JobManifestsPath;
        /// <summary>
        /// File path of the rendered manifest relative to the URI.
        /// </summary>
        public readonly string ManifestPath;
        /// <summary>
        /// File path of the resolved Skaffold configuration relative to the URI.
        /// </summary>
        public readonly string SkaffoldConfigPath;

        [OutputConstructor]
        private PhaseArtifactResponse(
            string jobManifestsPath,

            string manifestPath,

            string skaffoldConfigPath)
        {
            JobManifestsPath = jobManifestsPath;
            ManifestPath = manifestPath;
            SkaffoldConfigPath = skaffoldConfigPath;
        }
    }
}