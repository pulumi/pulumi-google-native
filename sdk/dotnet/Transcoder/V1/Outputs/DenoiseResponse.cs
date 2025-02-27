// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Transcoder.V1.Outputs
{

    /// <summary>
    /// Denoise preprocessing configuration. **Note:** This configuration is not supported.
    /// </summary>
    [OutputType]
    public sealed class DenoiseResponse
    {
        /// <summary>
        /// Set strength of the denoise. Enter a value between 0 and 1. The higher the value, the smoother the image. 0 is no denoising. The default is 0.
        /// </summary>
        public readonly double Strength;
        /// <summary>
        /// Set the denoiser mode. The default is `standard`. Supported denoiser modes: - `standard` - `grain`
        /// </summary>
        public readonly string Tune;

        [OutputConstructor]
        private DenoiseResponse(
            double strength,

            string tune)
        {
            Strength = strength;
            Tune = tune;
        }
    }
}
