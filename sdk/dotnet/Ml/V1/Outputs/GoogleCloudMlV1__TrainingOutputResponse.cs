// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Outputs
{

    /// <summary>
    /// Represents results of a training job. Output only.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudMlV1__TrainingOutputResponse
    {
        /// <summary>
        /// Details related to built-in algorithms jobs. Only set for built-in algorithms jobs.
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__BuiltInAlgorithmOutputResponse BuiltInAlgorithmOutput;
        /// <summary>
        /// The number of hyperparameter tuning trials that completed successfully. Only set for hyperparameter tuning jobs.
        /// </summary>
        public readonly string CompletedTrialCount;
        /// <summary>
        /// The amount of ML units consumed by the job.
        /// </summary>
        public readonly double ConsumedMLUnits;
        /// <summary>
        /// The TensorFlow summary tag name used for optimizing hyperparameter tuning trials. See [`HyperparameterSpec.hyperparameterMetricTag`](#HyperparameterSpec.FIELDS.hyperparameter_metric_tag) for more information. Only set for hyperparameter tuning jobs.
        /// </summary>
        public readonly string HyperparameterMetricTag;
        /// <summary>
        /// Whether this job is a built-in Algorithm job.
        /// </summary>
        public readonly bool IsBuiltInAlgorithmJob;
        /// <summary>
        /// Whether this job is a hyperparameter tuning job.
        /// </summary>
        public readonly bool IsHyperparameterTuningJob;
        /// <summary>
        /// Results for individual Hyperparameter trials. Only set for hyperparameter tuning jobs.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudMlV1__HyperparameterOutputResponse> Trials;
        /// <summary>
        /// URIs for accessing [interactive shells](https://cloud.google.com/ai-platform/training/docs/monitor-debug-interactive-shell) (one URI for each training node). Only available if training_input.enable_web_access is `true`. The keys are names of each node in the training job; for example, `master-replica-0` for the master node, `worker-replica-0` for the first worker, and `ps-replica-0` for the first parameter server. The values are the URIs for each node's interactive shell.
        /// </summary>
        public readonly ImmutableDictionary<string, string> WebAccessUris;

        [OutputConstructor]
        private GoogleCloudMlV1__TrainingOutputResponse(
            Outputs.GoogleCloudMlV1__BuiltInAlgorithmOutputResponse builtInAlgorithmOutput,

            string completedTrialCount,

            double consumedMLUnits,

            string hyperparameterMetricTag,

            bool isBuiltInAlgorithmJob,

            bool isHyperparameterTuningJob,

            ImmutableArray<Outputs.GoogleCloudMlV1__HyperparameterOutputResponse> trials,

            ImmutableDictionary<string, string> webAccessUris)
        {
            BuiltInAlgorithmOutput = builtInAlgorithmOutput;
            CompletedTrialCount = completedTrialCount;
            ConsumedMLUnits = consumedMLUnits;
            HyperparameterMetricTag = hyperparameterMetricTag;
            IsBuiltInAlgorithmJob = isBuiltInAlgorithmJob;
            IsHyperparameterTuningJob = isHyperparameterTuningJob;
            Trials = trials;
            WebAccessUris = webAccessUris;
        }
    }
}