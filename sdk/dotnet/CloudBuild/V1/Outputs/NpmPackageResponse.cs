// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Outputs
{

    /// <summary>
    /// Npm package to upload to Artifact Registry upon successful completion of all build steps.
    /// </summary>
    [OutputType]
    public sealed class NpmPackageResponse
    {
        /// <summary>
        /// Path to the package.json. e.g. workspace/path/to/package
        /// </summary>
        public readonly string PackagePath;
        /// <summary>
        /// Artifact Registry repository, in the form "https://$REGION-npm.pkg.dev/$PROJECT/$REPOSITORY" Npm package in the workspace specified by path will be zipped and uploaded to Artifact Registry with this location as a prefix.
        /// </summary>
        public readonly string Repository;

        [OutputConstructor]
        private NpmPackageResponse(
            string packagePath,

            string repository)
        {
            PackagePath = packagePath;
            Repository = repository;
        }
    }
}