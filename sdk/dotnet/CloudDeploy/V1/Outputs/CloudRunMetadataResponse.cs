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
    /// CloudRunMetadata contains information from a Cloud Run deployment.
    /// </summary>
    [OutputType]
    public sealed class CloudRunMetadataResponse
    {
        /// <summary>
        /// The Cloud Run Revision id associated with a `Rollout`.
        /// </summary>
        public readonly string Revision;
        /// <summary>
        /// The name of the Cloud Run Service that is associated with a `Rollout`. Format is projects/{project}/locations/{location}/services/{service}.
        /// </summary>
        public readonly string Service;
        /// <summary>
        /// The Cloud Run Service urls that are associated with a `Rollout`.
        /// </summary>
        public readonly ImmutableArray<string> ServiceUrls;

        [OutputConstructor]
        private CloudRunMetadataResponse(
            string revision,

            string service,

            ImmutableArray<string> serviceUrls)
        {
            Revision = revision;
            Service = service;
            ServiceUrls = serviceUrls;
        }
    }
}