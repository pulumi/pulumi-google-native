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
    /// Configuration for Binary Authorization.
    /// </summary>
    public sealed class BinaryAuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Binary Authorization.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Mode of operation for binauthz policy evaluation. Currently the only options are equivalent to enable/disable. If unspecified, defaults to DISABLED.
        /// </summary>
        [Input("evaluationMode")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.BinaryAuthorizationEvaluationMode>? EvaluationMode { get; set; }

        public BinaryAuthorizationArgs()
        {
        }
    }
}