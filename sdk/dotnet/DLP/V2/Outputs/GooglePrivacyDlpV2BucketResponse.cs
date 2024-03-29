// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Outputs
{

    /// <summary>
    /// Bucket is represented as a range, along with replacement values.
    /// </summary>
    [OutputType]
    public sealed class GooglePrivacyDlpV2BucketResponse
    {
        /// <summary>
        /// Upper bound of the range, exclusive; type must match min.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2ValueResponse Max;
        /// <summary>
        /// Lower bound of the range, inclusive. Type should be the same as max if used.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2ValueResponse Min;
        /// <summary>
        /// Replacement value for this bucket.
        /// </summary>
        public readonly Outputs.GooglePrivacyDlpV2ValueResponse ReplacementValue;

        [OutputConstructor]
        private GooglePrivacyDlpV2BucketResponse(
            Outputs.GooglePrivacyDlpV2ValueResponse max,

            Outputs.GooglePrivacyDlpV2ValueResponse min,

            Outputs.GooglePrivacyDlpV2ValueResponse replacementValue)
        {
            Max = max;
            Min = min;
            ReplacementValue = replacementValue;
        }
    }
}
