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
    /// Represents specification of a Study.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1beta1StudySpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The search algorithm specified for the Study.
        /// </summary>
        [Input("algorithm")]
        public Input<Pulumi.GoogleNative.Aiplatform.V1Beta1.GoogleCloudAiplatformV1beta1StudySpecAlgorithm>? Algorithm { get; set; }

        /// <summary>
        /// The automated early stopping spec using convex stopping rule.
        /// </summary>
        [Input("convexAutomatedStoppingSpec")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1StudySpecConvexAutomatedStoppingSpecArgs>? ConvexAutomatedStoppingSpec { get; set; }

        /// <summary>
        /// Deprecated. The automated early stopping using convex stopping rule.
        /// </summary>
        [Input("convexStopConfig")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1StudySpecConvexStopConfigArgs>? ConvexStopConfig { get; set; }

        /// <summary>
        /// The automated early stopping spec using decay curve rule.
        /// </summary>
        [Input("decayCurveStoppingSpec")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1StudySpecDecayCurveAutomatedStoppingSpecArgs>? DecayCurveStoppingSpec { get; set; }

        /// <summary>
        /// Describe which measurement selection type will be used
        /// </summary>
        [Input("measurementSelectionType")]
        public Input<Pulumi.GoogleNative.Aiplatform.V1Beta1.GoogleCloudAiplatformV1beta1StudySpecMeasurementSelectionType>? MeasurementSelectionType { get; set; }

        /// <summary>
        /// The automated early stopping spec using median rule.
        /// </summary>
        [Input("medianAutomatedStoppingSpec")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1StudySpecMedianAutomatedStoppingSpecArgs>? MedianAutomatedStoppingSpec { get; set; }

        [Input("metrics", required: true)]
        private InputList<Inputs.GoogleCloudAiplatformV1beta1StudySpecMetricSpecArgs>? _metrics;

        /// <summary>
        /// Metric specs for the Study.
        /// </summary>
        public InputList<Inputs.GoogleCloudAiplatformV1beta1StudySpecMetricSpecArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.GoogleCloudAiplatformV1beta1StudySpecMetricSpecArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// The observation noise level of the study. Currently only supported by the Vertex AI Vizier service. Not supported by HyperparameterTuningJob or TrainingPipeline.
        /// </summary>
        [Input("observationNoise")]
        public Input<Pulumi.GoogleNative.Aiplatform.V1Beta1.GoogleCloudAiplatformV1beta1StudySpecObservationNoise>? ObservationNoise { get; set; }

        [Input("parameters", required: true)]
        private InputList<Inputs.GoogleCloudAiplatformV1beta1StudySpecParameterSpecArgs>? _parameters;

        /// <summary>
        /// The set of parameters to tune.
        /// </summary>
        public InputList<Inputs.GoogleCloudAiplatformV1beta1StudySpecParameterSpecArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.GoogleCloudAiplatformV1beta1StudySpecParameterSpecArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Conditions for automated stopping of a Study. Enable automated stopping by configuring at least one condition.
        /// </summary>
        [Input("studyStoppingConfig")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1StudySpecStudyStoppingConfigArgs>? StudyStoppingConfig { get; set; }

        /// <summary>
        /// The configuration info/options for transfer learning. Currently supported for Vertex AI Vizier service, not HyperParameterTuningJob
        /// </summary>
        [Input("transferLearningConfig")]
        public Input<Inputs.GoogleCloudAiplatformV1beta1StudySpecTransferLearningConfigArgs>? TransferLearningConfig { get; set; }

        public GoogleCloudAiplatformV1beta1StudySpecArgs()
        {
        }
        public static new GoogleCloudAiplatformV1beta1StudySpecArgs Empty => new GoogleCloudAiplatformV1beta1StudySpecArgs();
    }
}