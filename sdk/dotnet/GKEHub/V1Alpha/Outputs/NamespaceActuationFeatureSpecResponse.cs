// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1Alpha.Outputs
{

    /// <summary>
    /// An empty spec for actuation feature. This is required since Feature proto requires a spec.
    /// </summary>
    [OutputType]
    public sealed class NamespaceActuationFeatureSpecResponse
    {
        /// <summary>
        /// actuation_mode controls the behavior of the controller
        /// </summary>
        public readonly string ActuationMode;

        [OutputConstructor]
        private NamespaceActuationFeatureSpecResponse(string actuationMode)
        {
            ActuationMode = actuationMode;
        }
    }
}