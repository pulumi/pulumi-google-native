// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// IntegerFieldInfo defines output for any integer type field.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataProfileResultProfileFieldProfileInfoIntegerFieldInfoResponse
    {
        /// <summary>
        /// The average of non-null values of integer field in the sampled data. Return NaN, if the field has a NaN. Optional if zero non-null rows.
        /// </summary>
        public readonly double Average;
        /// <summary>
        /// The maximum value of an integer field in the sampled data. Return NaN, if the field has a NaN. Optional if zero non-null rows.
        /// </summary>
        public readonly string Max;
        /// <summary>
        /// The minimum value of an integer field in the sampled data. Return NaN, if the field has a NaN. Optional if zero non-null rows.
        /// </summary>
        public readonly string Min;
        /// <summary>
        /// A quartile divide the number of data points into four parts, or quarters, of more-or-less equal size. Three main quartiles used are: The first quartile (Q1) splits off the lowest 25% of data from the highest 75%. It is also known as the lower or 25th empirical quartile, as 25% of the data is below this point. The second quartile (Q2) is the median of a data set. So, 50% of the data lies below this point. The third quartile (Q3) splits off the highest 25% of data from the lowest 75%. It is known as the upper or 75th empirical quartile, as 75% of the data lies below this point. So, here the quartiles is provided as an ordered list of quartile values, occurring in order Q1, median, Q3.
        /// </summary>
        public readonly ImmutableArray<string> Quartiles;
        /// <summary>
        /// The standard deviation of non-null of integer field in the sampled data. Return NaN, if the field has a NaN. Optional if zero non-null rows.
        /// </summary>
        public readonly double StandardDeviation;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataProfileResultProfileFieldProfileInfoIntegerFieldInfoResponse(
            double average,

            string max,

            string min,

            ImmutableArray<string> quartiles,

            double standardDeviation)
        {
            Average = average;
            Max = max;
            Min = min;
            Quartiles = quartiles;
            StandardDeviation = standardDeviation;
        }
    }
}