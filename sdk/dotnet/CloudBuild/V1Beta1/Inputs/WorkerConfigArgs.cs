// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudBuild.V1Beta1.Inputs
{

    /// <summary>
    /// Defines the configuration to be used for creating workers in the pool.
    /// </summary>
    public sealed class WorkerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of the disk attached to the worker, in GB. See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). Specify a value of up to 1000. If `0` is specified, Cloud Build will use a standard disk size.
        /// </summary>
        [Input("diskSizeGb")]
        public Input<string>? DiskSizeGb { get; set; }

        /// <summary>
        /// Machine type of a worker, such as `n1-standard-1`. See [Worker pool config file](https://cloud.google.com/cloud-build/docs/custom-workers/worker-pool-config-file). If left blank, Cloud Build will use `n1-standard-1`.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        /// <summary>
        /// If true, workers are created without any public address, which prevents network egress to public IPs.
        /// </summary>
        [Input("noExternalIp")]
        public Input<bool>? NoExternalIp { get; set; }

        public WorkerConfigArgs()
        {
        }
        public static new WorkerConfigArgs Empty => new WorkerConfigArgs();
    }
}
