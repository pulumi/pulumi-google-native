// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Inputs
{

    /// <summary>
    /// Extracts an archive of the type specified in the specified directory.
    /// </summary>
    public sealed class SoftwareRecipeStepExtractArchiveArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the relevant artifact in the recipe.
        /// </summary>
        [Input("artifactId", required: true)]
        public Input<string> ArtifactId { get; set; } = null!;

        /// <summary>
        /// Directory to extract archive to. Defaults to `/` on Linux or `C:\` on Windows.
        /// </summary>
        [Input("destination")]
        public Input<string>? Destination { get; set; }

        /// <summary>
        /// The type of the archive to extract.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.GoogleNative.OSConfig.V1Beta.SoftwareRecipeStepExtractArchiveType> Type { get; set; } = null!;

        public SoftwareRecipeStepExtractArchiveArgs()
        {
        }
    }
}