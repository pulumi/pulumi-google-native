// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1.Inputs
{

    /// <summary>
    /// The configuration of a trigger that creates a build whenever an event from Repo API is received.
    /// </summary>
    public sealed class RepositoryEventConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Filter to match changes in pull requests.
        /// </summary>
        [Input("pullRequest")]
        public Input<Inputs.PullRequestFilterArgs>? PullRequest { get; set; }

        /// <summary>
        /// Filter to match changes in refs like branches, tags.
        /// </summary>
        [Input("push")]
        public Input<Inputs.PushFilterArgs>? Push { get; set; }

        /// <summary>
        /// The resource name of the Repo API resource.
        /// </summary>
        [Input("repository")]
        public Input<string>? Repository { get; set; }

        public RepositoryEventConfigArgs()
        {
        }
        public static new RepositoryEventConfigArgs Empty => new RepositoryEventConfigArgs();
    }
}
