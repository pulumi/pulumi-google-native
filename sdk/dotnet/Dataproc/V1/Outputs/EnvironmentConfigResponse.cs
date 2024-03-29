// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Environment configuration for a workload.
    /// </summary>
    [OutputType]
    public sealed class EnvironmentConfigResponse
    {
        /// <summary>
        /// Optional. Execution configuration for a workload.
        /// </summary>
        public readonly Outputs.ExecutionConfigResponse ExecutionConfig;
        /// <summary>
        /// Optional. Peripherals configuration that workload has access to.
        /// </summary>
        public readonly Outputs.PeripheralsConfigResponse PeripheralsConfig;

        [OutputConstructor]
        private EnvironmentConfigResponse(
            Outputs.ExecutionConfigResponse executionConfig,

            Outputs.PeripheralsConfigResponse peripheralsConfig)
        {
            ExecutionConfig = executionConfig;
            PeripheralsConfig = peripheralsConfig;
        }
    }
}
