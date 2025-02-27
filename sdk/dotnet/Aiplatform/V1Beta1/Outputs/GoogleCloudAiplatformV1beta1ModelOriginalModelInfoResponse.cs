// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// Contains information about the original Model if this Model is a copy.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1ModelOriginalModelInfoResponse
    {
        /// <summary>
        /// The resource name of the Model this Model is a copy of, including the revision. Format: `projects/{project}/locations/{location}/models/{model_id}@{version_id}`
        /// </summary>
        public readonly string Model;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1ModelOriginalModelInfoResponse(string model)
        {
            Model = model;
        }
    }
}
