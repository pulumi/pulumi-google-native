// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Recommendationengine.V1Beta1.Inputs
{

    /// <summary>
    /// A list of string features.
    /// </summary>
    public sealed class GoogleCloudRecommendationengineV1beta1FeatureMapStringListArgs : global::Pulumi.ResourceArgs
    {
        [Input("value")]
        private InputList<string>? _value;

        /// <summary>
        /// String feature value with a length limit of 128 bytes.
        /// </summary>
        public InputList<string> Value
        {
            get => _value ?? (_value = new InputList<string>());
            set => _value = value;
        }

        public GoogleCloudRecommendationengineV1beta1FeatureMapStringListArgs()
        {
        }
        public static new GoogleCloudRecommendationengineV1beta1FeatureMapStringListArgs Empty => new GoogleCloudRecommendationengineV1beta1FeatureMapStringListArgs();
    }
}