// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// GPUDriverInstallationConfig specifies the version of GPU driver to be auto installed.
    /// </summary>
    public sealed class GPUDriverInstallationConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode for how the GPU driver is installed.
        /// </summary>
        [Input("gpuDriverVersion")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.GPUDriverInstallationConfigGpuDriverVersion>? GpuDriverVersion { get; set; }

        public GPUDriverInstallationConfigArgs()
        {
        }
        public static new GPUDriverInstallationConfigArgs Empty => new GPUDriverInstallationConfigArgs();
    }
}