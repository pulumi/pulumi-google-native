// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    /// <summary>
    /// Represents output related to a built-in algorithm Job.
    /// </summary>
    public sealed class GoogleCloudMlV1__BuiltInAlgorithmOutputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Framework on which the built-in algorithm was trained.
        /// </summary>
        [Input("framework")]
        public Input<string>? Framework { get; set; }

        /// <summary>
        /// The Cloud Storage path to the `model/` directory where the training job saves the trained model. Only set for successful jobs that don't use hyperparameter tuning.
        /// </summary>
        [Input("modelPath")]
        public Input<string>? ModelPath { get; set; }

        /// <summary>
        /// Python version on which the built-in algorithm was trained.
        /// </summary>
        [Input("pythonVersion")]
        public Input<string>? PythonVersion { get; set; }

        /// <summary>
        /// AI Platform runtime version on which the built-in algorithm was trained.
        /// </summary>
        [Input("runtimeVersion")]
        public Input<string>? RuntimeVersion { get; set; }

        public GoogleCloudMlV1__BuiltInAlgorithmOutputArgs()
        {
        }
    }
}