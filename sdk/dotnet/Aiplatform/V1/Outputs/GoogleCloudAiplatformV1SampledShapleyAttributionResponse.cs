// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Outputs
{

    /// <summary>
    /// An attribution method that approximates Shapley values for features that contribute to the label being predicted. A sampling strategy is used to approximate the value rather than considering all subsets of features.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1SampledShapleyAttributionResponse
    {
        /// <summary>
        /// The number of feature permutations to consider when approximating the Shapley values. Valid range of its value is [1, 50], inclusively.
        /// </summary>
        public readonly int PathCount;

        [OutputConstructor]
        private GoogleCloudAiplatformV1SampledShapleyAttributionResponse(int pathCount)
        {
            PathCount = pathCount;
        }
    }
}