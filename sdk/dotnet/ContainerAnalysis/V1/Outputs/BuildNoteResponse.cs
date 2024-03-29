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
    /// Note holding the version of the provider's builder and the signature of the provenance message in the build details occurrence.
    /// </summary>
    [OutputType]
    public sealed class BuildNoteResponse
    {
        /// <summary>
        /// Immutable. Version of the builder which produced this build.
        /// </summary>
        public readonly string BuilderVersion;

        [OutputConstructor]
        private BuildNoteResponse(string builderVersion)
        {
            BuilderVersion = builderVersion;
        }
    }
}
