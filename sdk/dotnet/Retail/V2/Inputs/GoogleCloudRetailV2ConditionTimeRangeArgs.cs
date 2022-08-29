// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2.Inputs
{

    /// <summary>
    /// Used for time-dependent conditions. Example: Want to have rule applied for week long sale.
    /// </summary>
    public sealed class GoogleCloudRetailV2ConditionTimeRangeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End of time range. Range is inclusive.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        /// <summary>
        /// Start of time range. Range is inclusive.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public GoogleCloudRetailV2ConditionTimeRangeArgs()
        {
        }
        public static new GoogleCloudRetailV2ConditionTimeRangeArgs Empty => new GoogleCloudRetailV2ConditionTimeRangeArgs();
    }
}