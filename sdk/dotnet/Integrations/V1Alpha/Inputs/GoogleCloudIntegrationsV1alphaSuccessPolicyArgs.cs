// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// Policy that dictates the behavior for the task after it completes successfully.
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaSuccessPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// State to which the execution snapshot status will be set if the task succeeds.
        /// </summary>
        [Input("finalState")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.GoogleCloudIntegrationsV1alphaSuccessPolicyFinalState>? FinalState { get; set; }

        public GoogleCloudIntegrationsV1alphaSuccessPolicyArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaSuccessPolicyArgs Empty => new GoogleCloudIntegrationsV1alphaSuccessPolicyArgs();
    }
}