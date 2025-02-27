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
    /// Domain details of the input feature value. Provides numeric information about the feature, such as its range (min, max). If the feature has been pre-processed, for example with z-scoring, then it provides information about how to recover the original feature. For example, if the input feature is an image and it has been pre-processed to obtain 0-mean and stddev = 1 values, then original_mean, and original_stddev refer to the mean and stddev of the original feature (e.g. image tensor) from which input feature (with mean = 0 and stddev = 1) was obtained.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataFeatureValueDomainResponse
    {
        /// <summary>
        /// The maximum permissible value for this feature.
        /// </summary>
        public readonly double MaxValue;
        /// <summary>
        /// The minimum permissible value for this feature.
        /// </summary>
        public readonly double MinValue;
        /// <summary>
        /// If this input feature has been normalized to a mean value of 0, the original_mean specifies the mean value of the domain prior to normalization.
        /// </summary>
        public readonly double OriginalMean;
        /// <summary>
        /// If this input feature has been normalized to a standard deviation of 1.0, the original_stddev specifies the standard deviation of the domain prior to normalization.
        /// </summary>
        public readonly double OriginalStddev;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1ExplanationMetadataInputMetadataFeatureValueDomainResponse(
            double maxValue,

            double minValue,

            double originalMean,

            double originalStddev)
        {
            MaxValue = maxValue;
            MinValue = minValue;
            OriginalMean = originalMean;
            OriginalStddev = originalStddev;
        }
    }
}
