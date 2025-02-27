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
    /// Configuration for ConvexStopPolicy.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1StudySpecConvexStopConfigResponse
    {
        /// <summary>
        /// The number of Trial measurements used in autoregressive model for value prediction. A trial won't be considered early stopping if has fewer measurement points.
        /// </summary>
        public readonly string AutoregressiveOrder;
        /// <summary>
        /// The hyper-parameter name used in the tuning job that stands for learning rate. Leave it blank if learning rate is not in a parameter in tuning. The learning_rate is used to estimate the objective value of the ongoing trial.
        /// </summary>
        public readonly string LearningRateParameterName;
        /// <summary>
        /// Steps used in predicting the final objective for early stopped trials. In general, it's set to be the same as the defined steps in training / tuning. When use_steps is false, this field is set to the maximum elapsed seconds.
        /// </summary>
        public readonly string MaxNumSteps;
        /// <summary>
        /// Minimum number of steps for a trial to complete. Trials which do not have a measurement with num_steps &gt; min_num_steps won't be considered for early stopping. It's ok to set it to 0, and a trial can be early stopped at any stage. By default, min_num_steps is set to be one-tenth of the max_num_steps. When use_steps is false, this field is set to the minimum elapsed seconds.
        /// </summary>
        public readonly string MinNumSteps;
        /// <summary>
        /// This bool determines whether or not the rule is applied based on elapsed_secs or steps. If use_seconds==false, the early stopping decision is made according to the predicted objective values according to the target steps. If use_seconds==true, elapsed_secs is used instead of steps. Also, in this case, the parameters max_num_steps and min_num_steps are overloaded to contain max_elapsed_seconds and min_elapsed_seconds.
        /// </summary>
        public readonly bool UseSeconds;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1StudySpecConvexStopConfigResponse(
            string autoregressiveOrder,

            string learningRateParameterName,

            string maxNumSteps,

            string minNumSteps,

            bool useSeconds)
        {
            AutoregressiveOrder = autoregressiveOrder;
            LearningRateParameterName = learningRateParameterName;
            MaxNumSteps = maxNumSteps;
            MinNumSteps = minNumSteps;
            UseSeconds = useSeconds;
        }
    }
}
