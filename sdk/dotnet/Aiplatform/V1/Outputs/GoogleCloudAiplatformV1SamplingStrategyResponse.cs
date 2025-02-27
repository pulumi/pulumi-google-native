// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Outputs
{

    /// <summary>
    /// Sampling Strategy for logging, can be for both training and prediction dataset.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1SamplingStrategyResponse
    {
        /// <summary>
        /// Random sample config. Will support more sampling strategies later.
        /// </summary>
        public readonly Outputs.GoogleCloudAiplatformV1SamplingStrategyRandomSampleConfigResponse RandomSampleConfig;

        [OutputConstructor]
        private GoogleCloudAiplatformV1SamplingStrategyResponse(Outputs.GoogleCloudAiplatformV1SamplingStrategyRandomSampleConfigResponse randomSampleConfig)
        {
            RandomSampleConfig = randomSampleConfig;
        }
    }
}
