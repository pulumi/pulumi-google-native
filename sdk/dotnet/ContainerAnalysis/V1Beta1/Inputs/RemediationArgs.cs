// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.Inputs
{

    /// <summary>
    /// Specifies details on how to handle (and presumably, fix) a vulnerability.
    /// </summary>
    public sealed class RemediationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contains a comprehensive human-readable discussion of the remediation.
        /// </summary>
        [Input("details")]
        public Input<string>? Details { get; set; }

        /// <summary>
        /// The type of remediation that can be applied.
        /// </summary>
        [Input("remediationType")]
        public Input<Pulumi.GoogleNative.ContainerAnalysis.V1Beta1.RemediationRemediationType>? RemediationType { get; set; }

        /// <summary>
        /// Contains the URL where to obtain the remediation.
        /// </summary>
        [Input("remediationUri")]
        public Input<Inputs.RelatedUrlArgs>? RemediationUri { get; set; }

        public RemediationArgs()
        {
        }
        public static new RemediationArgs Empty => new RemediationArgs();
    }
}