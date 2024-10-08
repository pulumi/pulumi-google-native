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
    /// CloudRunRenderMetadata contains Cloud Run information associated with a `Release` render.
    /// </summary>
    [OutputType]
    public sealed class CloudRunRenderMetadataResponse
    {
        /// <summary>
        /// The name of the Cloud Run Service in the rendered manifest. Format is `projects/{project}/locations/{location}/services/{service}`.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private CloudRunRenderMetadataResponse(string service)
        {
            Service = service;
        }
    }
}
