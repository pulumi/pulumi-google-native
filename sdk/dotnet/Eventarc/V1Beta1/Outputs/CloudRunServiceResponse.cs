// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1Beta1.Outputs
{

    /// <summary>
    /// Represents a Cloud Run service destination.
    /// </summary>
    [OutputType]
    public sealed class CloudRunServiceResponse
    {
        /// <summary>
        /// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// The region the Cloud Run service is deployed in.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// The name of the Cloud run service being addressed (see https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services). Only services located in the same project of the trigger object can be addressed.
        /// </summary>
        public readonly string Service;

        [OutputConstructor]
        private CloudRunServiceResponse(
            string path,

            string region,

            string service)
        {
            Path = path;
            Region = region;
            Service = service;
        }
    }
}