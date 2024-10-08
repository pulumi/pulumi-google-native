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
    /// Represents the configuration for a replica in a cluster.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudMlV1__ReplicaConfigResponse
    {
        /// <summary>
        /// Represents the type and number of accelerators used by the replica. [Learn about restrictions on accelerator configurations for training.](/ai-platform/training/docs/using-gpus#compute-engine-machine-types-with-gpu)
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__AcceleratorConfigResponse AcceleratorConfig;
        /// <summary>
        /// Arguments to the entrypoint command. The following rules apply for container_command and container_args: - If you do not supply command or args: The defaults defined in the Docker image are used. - If you supply a command but no args: The default EntryPoint and the default Cmd defined in the Docker image are ignored. Your command is run without any arguments. - If you supply only args: The default Entrypoint defined in the Docker image is run with the args that you supplied. - If you supply a command and args: The default Entrypoint and the default Cmd defined in the Docker image are ignored. Your command is run with your args. It cannot be set if custom container image is not provided. Note that this field and [TrainingInput.args] are mutually exclusive, i.e., both cannot be set at the same time.
        /// </summary>
        public readonly ImmutableArray<string> ContainerArgs;
        /// <summary>
        /// The command with which the replica's custom container is run. If provided, it will override default ENTRYPOINT of the docker image. If not provided, the docker image's ENTRYPOINT is used. It cannot be set if custom container image is not provided. Note that this field and [TrainingInput.args] are mutually exclusive, i.e., both cannot be set at the same time.
        /// </summary>
        public readonly ImmutableArray<string> ContainerCommand;
        /// <summary>
        /// Represents the configuration of disk options.
        /// </summary>
        public readonly Outputs.GoogleCloudMlV1__DiskConfigResponse DiskConfig;
        /// <summary>
        /// The Docker image to run on the replica. This image must be in Container Registry. Learn more about [configuring custom containers](/ai-platform/training/docs/distributed-training-containers).
        /// </summary>
        public readonly string ImageUri;
        /// <summary>
        /// The AI Platform runtime version that includes a TensorFlow version matching the one used in the custom container. This field is required if the replica is a TPU worker that uses a custom container. Otherwise, do not specify this field. This must be a [runtime version that currently supports training with TPUs](/ml-engine/docs/tensorflow/runtime-version-list#tpu-support). Note that the version of TensorFlow included in a runtime version may differ from the numbering of the runtime version itself, because it may have a different [patch version](https://www.tensorflow.org/guide/version_compat#semantic_versioning_20). In this field, you must specify the runtime version (TensorFlow minor version). For example, if your custom container runs TensorFlow `1.x.y`, specify `1.x`.
        /// </summary>
        public readonly string TpuTfVersion;

        [OutputConstructor]
        private GoogleCloudMlV1__ReplicaConfigResponse(
            Outputs.GoogleCloudMlV1__AcceleratorConfigResponse acceleratorConfig,

            ImmutableArray<string> containerArgs,

            ImmutableArray<string> containerCommand,

            Outputs.GoogleCloudMlV1__DiskConfigResponse diskConfig,

            string imageUri,

            string tpuTfVersion)
        {
            AcceleratorConfig = acceleratorConfig;
            ContainerArgs = containerArgs;
            ContainerCommand = containerCommand;
            DiskConfig = diskConfig;
            ImageUri = imageUri;
            TpuTfVersion = tpuTfVersion;
        }
    }
}
