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
    /// Note holding the version of the provider's builder and the signature of the provenance message in the build details occurrence.
    /// </summary>
    public sealed class BuildArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. Version of the builder which produced this build.
        /// </summary>
        [Input("builderVersion", required: true)]
        public Input<string> BuilderVersion { get; set; } = null!;

        /// <summary>
        /// Signature of the build in occurrences pointing to this build note containing build details.
        /// </summary>
        [Input("signature")]
        public Input<Inputs.BuildSignatureArgs>? Signature { get; set; }

        public BuildArgs()
        {
        }
        public static new BuildArgs Empty => new BuildArgs();
    }
}
