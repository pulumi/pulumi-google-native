// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudChannel.V1.Inputs
{

    /// <summary>
    /// Specifies the override to conditionally apply.
    /// </summary>
    public sealed class GoogleCloudChannelV1ConditionalOverrideArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about the applied override's adjustment.
        /// </summary>
        [Input("adjustment", required: true)]
        public Input<Inputs.GoogleCloudChannelV1RepricingAdjustmentArgs> Adjustment { get; set; } = null!;

        /// <summary>
        /// The RebillingBasis to use for the applied override. Shows the relative cost based on your repricing costs.
        /// </summary>
        [Input("rebillingBasis", required: true)]
        public Input<Pulumi.GoogleNative.CloudChannel.V1.GoogleCloudChannelV1ConditionalOverrideRebillingBasis> RebillingBasis { get; set; } = null!;

        /// <summary>
        /// Specifies the condition which, if met, will apply the override.
        /// </summary>
        [Input("repricingCondition", required: true)]
        public Input<Inputs.GoogleCloudChannelV1RepricingConditionArgs> RepricingCondition { get; set; } = null!;

        public GoogleCloudChannelV1ConditionalOverrideArgs()
        {
        }
        public static new GoogleCloudChannelV1ConditionalOverrideArgs Empty => new GoogleCloudChannelV1ConditionalOverrideArgs();
    }
}