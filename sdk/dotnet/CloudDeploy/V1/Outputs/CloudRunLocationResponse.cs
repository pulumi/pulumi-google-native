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
    /// Information specifying where to deploy a Cloud Run Service.
    /// </summary>
    [OutputType]
    public sealed class CloudRunLocationResponse
    {
        /// <summary>
        /// The location for the Cloud Run Service. Format must be `projects/{project}/locations/{location}`.
        /// </summary>
        public readonly string Location;

        [OutputConstructor]
        private CloudRunLocationResponse(string location)
        {
            Location = location;
        }
    }
}
