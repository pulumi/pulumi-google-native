// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1.Inputs
{

    /// <summary>
    /// Requests are randomly selected.
    /// </summary>
    public sealed class GoogleCloudAiplatformV1SamplingStrategyRandomSampleConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sample rate (0, 1]
        /// </summary>
        [Input("sampleRate")]
        public Input<double>? SampleRate { get; set; }

        public GoogleCloudAiplatformV1SamplingStrategyRandomSampleConfigArgs()
        {
        }
        public static new GoogleCloudAiplatformV1SamplingStrategyRandomSampleConfigArgs Empty => new GoogleCloudAiplatformV1SamplingStrategyRandomSampleConfigArgs();
    }
}