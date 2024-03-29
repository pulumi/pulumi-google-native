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
    /// Noise sigma by features. Noise sigma represents the standard deviation of the gaussian kernel that will be used to add noise to interpolated inputs prior to computing gradients.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1FeatureNoiseSigmaResponse
    {
        /// <summary>
        /// Noise sigma per feature. No noise is added to features that are not set.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1FeatureNoiseSigmaNoiseSigmaForFeatureResponse> NoiseSigma;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1FeatureNoiseSigmaResponse(ImmutableArray<Outputs.GoogleCloudAiplatformV1beta1FeatureNoiseSigmaNoiseSigmaForFeatureResponse> noiseSigma)
        {
            NoiseSigma = noiseSigma;
        }
    }
}
