// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V1.Outputs
{

    /// <summary>
    /// ResourceRequirements describes the compute resource requirements.
    /// </summary>
    [OutputType]
    public sealed class ResourceRequirementsResponse
    {
        /// <summary>
        /// (Optional) Only memory and CPU are supported. Limits describes the maximum amount of compute resources allowed. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        /// </summary>
        public readonly ImmutableDictionary<string, string> Limits;
        /// <summary>
        /// (Optional) Only memory and CPU are supported. Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        /// </summary>
        public readonly ImmutableDictionary<string, string> Requests;

        [OutputConstructor]
        private ResourceRequirementsResponse(
            ImmutableDictionary<string, string> limits,

            ImmutableDictionary<string, string> requests)
        {
            Limits = limits;
            Requests = requests;
        }
    }
}