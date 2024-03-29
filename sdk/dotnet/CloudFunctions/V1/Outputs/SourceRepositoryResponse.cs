// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V1.Outputs
{

    /// <summary>
    /// Describes SourceRepository, used to represent parameters related to source repository where a function is hosted.
    /// </summary>
    [OutputType]
    public sealed class SourceRepositoryResponse
    {
        /// <summary>
        /// The URL pointing to the hosted repository where the function were defined at the time of deployment. It always points to a specific commit in the format described above.
        /// </summary>
        public readonly string DeployedUrl;
        /// <summary>
        /// The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats: To refer to a specific commit: `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*` To refer to a moveable alias (branch): `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*` In particular, to refer to HEAD use `master` moveable alias. To refer to a specific fixed alias (tag): `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*` You may omit `paths/*` if you want to use the main directory.
        /// </summary>
        public readonly string Url;

        [OutputConstructor]
        private SourceRepositoryResponse(
            string deployedUrl,

            string url)
        {
            DeployedUrl = deployedUrl;
            Url = url;
        }
    }
}
