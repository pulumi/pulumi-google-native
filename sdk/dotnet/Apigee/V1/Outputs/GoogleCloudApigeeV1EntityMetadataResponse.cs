// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1.Outputs
{

    /// <summary>
    /// Metadata common to many entities in this API.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudApigeeV1EntityMetadataResponse
    {
        /// <summary>
        /// Time at which the API proxy was created, in milliseconds since epoch.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Time at which the API proxy was most recently modified, in milliseconds since epoch.
        /// </summary>
        public readonly string LastModifiedAt;
        /// <summary>
        /// The type of entity described
        /// </summary>
        public readonly string SubType;

        [OutputConstructor]
        private GoogleCloudApigeeV1EntityMetadataResponse(
            string createdAt,

            string lastModifiedAt,

            string subType)
        {
            CreatedAt = createdAt;
            LastModifiedAt = lastModifiedAt;
            SubType = subType;
        }
    }
}
