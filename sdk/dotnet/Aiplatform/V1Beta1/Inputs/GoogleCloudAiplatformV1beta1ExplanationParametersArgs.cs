// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Inputs
{

    /// <summary>
    /// Parameters to configure explaining for Model's predictions.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1beta1ExplanationParametersArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Example-based explanations that returns the nearest neighbors from the provided dataset.
        /// </summary>
        [Input("examples")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1ExamplesArgs>? Examples { get; set; }

        /// <summary>
        /// An attribution method that computes Aumann-Shapley values taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1703.01365
        /// </summary>
        [Input("integratedGradientsAttribution")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1IntegratedGradientsAttributionArgs>? IntegratedGradientsAttribution { get; set; }

        [Input("outputIndices")]
        private InputList<object>? _outputIndices;

        /// <summary>
        /// If populated, only returns attributions that have output_index contained in output_indices. It must be an ndarray of integers, with the same shape of the output it's explaining. If not populated, returns attributions for top_k indices of outputs. If neither top_k nor output_indices is populated, returns the argmax index of the outputs. Only applicable to Models that predict multiple outputs (e,g, multi-class Models that predict multiple classes).
        /// </summary>
        public InputList<object> OutputIndices
        {
            get => _outputIndices ?? (_outputIndices = new InputList<object>());
            set => _outputIndices = value;
        }

        /// <summary>
        /// An attribution method that approximates Shapley values for features that contribute to the label being predicted. A sampling strategy is used to approximate the value rather than considering all subsets of features. Refer to this paper for model details: https://arxiv.org/abs/1306.4265.
        /// </summary>
        [Input("sampledShapleyAttribution")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1SampledShapleyAttributionArgs>? SampledShapleyAttribution { get; set; }

        /// <summary>
        /// If populated, returns attributions for top K indices of outputs (defaults to 1). Only applies to Models that predicts more than one outputs (e,g, multi-class Models). When set to -1, returns explanations for all outputs.
        /// </summary>
        [Input("topK")]
        public Input<int>? TopK { get; set; }

        /// <summary>
        /// An attribution method that redistributes Integrated Gradients attribution to segmented regions, taking advantage of the model's fully differentiable structure. Refer to this paper for more details: https://arxiv.org/abs/1906.02825 XRAI currently performs better on natural images, like a picture of a house or an animal. If the images are taken in artificial environments, like a lab or manufacturing line, or from diagnostic equipment, like x-rays or quality-control cameras, use Integrated Gradients instead.
        /// </summary>
        [Input("xraiAttribution")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1XraiAttributionArgs>? XraiAttribution { get; set; }

        public GoogleCloudAiplatformV1beta1ExplanationParametersArgs()
        {
        }
        public static new GoogleCloudAiplatformV1beta1ExplanationParametersArgs Empty => new GoogleCloudAiplatformV1beta1ExplanationParametersArgs();
    }
}