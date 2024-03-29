// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Specifies an artifact available as a Google Cloud Storage object.
    /// </summary>
    [OutputType]
    public sealed class SoftwareRecipeArtifactGcsResponse
    {
        /// <summary>
        /// Bucket of the Google Cloud Storage object. Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `my-bucket`.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Must be provided if allow_insecure is false. Generation number of the Google Cloud Storage object. `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `1234567`.
        /// </summary>
        public readonly string Generation;
        /// <summary>
        /// Name of the Google Cloud Storage object. As specified [here] (https://cloud.google.com/storage/docs/naming#objectnames) Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `foo/bar`.
        /// </summary>
        public readonly string Object;

        [OutputConstructor]
        private SoftwareRecipeArtifactGcsResponse(
            string bucket,

            string generation,

            string @object)
        {
            Bucket = bucket;
            Generation = generation;
            Object = @object;
        }
    }
}
