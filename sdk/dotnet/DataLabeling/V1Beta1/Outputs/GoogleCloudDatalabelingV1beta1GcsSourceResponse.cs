// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataLabeling.V1Beta1.Outputs
{

    /// <summary>
    /// Source of the Cloud Storage file to be imported.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatalabelingV1beta1GcsSourceResponse
    {
        /// <summary>
        /// The input URI of source file. This must be a Cloud Storage path (`gs://...`).
        /// </summary>
        public readonly string InputUri;
        /// <summary>
        /// The format of the source file. Only "text/csv" is supported.
        /// </summary>
        public readonly string MimeType;

        [OutputConstructor]
        private GoogleCloudDatalabelingV1beta1GcsSourceResponse(
            string inputUri,

            string mimeType)
        {
            InputUri = inputUri;
            MimeType = mimeType;
        }
    }
}