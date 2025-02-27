// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// Buckets values based on fixed size ranges. The Bucketing transformation can provide all of this functionality, but requires more configuration. This message is provided as a convenience to the user for simple bucketing strategies. The transformed value will be a hyphenated string of {lower_bound}-{upper_bound}. For example, if lower_bound = 10 and upper_bound = 20, all values that are within this bucket will be replaced with "10-20". This can be used on data of type: double, long. If the bound Value type differs from the type of data being transformed, we will first attempt converting the type of the data to be transformed to match the type of the bound before comparing. See https://cloud.google.com/dlp/docs/concepts-bucketing to learn more.
    /// </summary>
    public sealed class GooglePrivacyDlpV2FixedSizeBucketingConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of each bucket (except for minimum and maximum buckets). So if `lower_bound` = 10, `upper_bound` = 89, and `bucket_size` = 10, then the following buckets would be used: -10, 10-20, 20-30, 30-40, 40-50, 50-60, 60-70, 70-80, 80-89, 89+. Precision up to 2 decimals works.
        /// </summary>
        [Input("bucketSize", required: true)]
        public Input<double> BucketSize { get; set; } = null!;

        /// <summary>
        /// Lower bound value of buckets. All values less than `lower_bound` are grouped together into a single bucket; for example if `lower_bound` = 10, then all values less than 10 are replaced with the value "-10".
        /// </summary>
        [Input("lowerBound", required: true)]
        public Input<Inputs.GooglePrivacyDlpV2ValueArgs> LowerBound { get; set; } = null!;

        /// <summary>
        /// Upper bound value of buckets. All values greater than upper_bound are grouped together into a single bucket; for example if `upper_bound` = 89, then all values greater than 89 are replaced with the value "89+".
        /// </summary>
        [Input("upperBound", required: true)]
        public Input<Inputs.GooglePrivacyDlpV2ValueArgs> UpperBound { get; set; } = null!;

        public GooglePrivacyDlpV2FixedSizeBucketingConfigArgs()
        {
        }
        public static new GooglePrivacyDlpV2FixedSizeBucketingConfigArgs Empty => new GooglePrivacyDlpV2FixedSizeBucketingConfigArgs();
    }
}
