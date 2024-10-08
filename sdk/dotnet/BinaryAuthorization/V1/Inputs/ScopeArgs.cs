// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BinaryAuthorization.V1.Inputs
{

    /// <summary>
    /// A scope specifier for `CheckSet` objects.
    /// </summary>
    public sealed class ScopeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Matches all Kubernetes service accounts in the provided namespace, unless a more specific `kubernetes_service_account` scope already matched.
        /// </summary>
        [Input("kubernetesNamespace")]
        public Input<string>? KubernetesNamespace { get; set; }

        /// <summary>
        /// Optional. Matches a single Kubernetes service account, e.g. `my-namespace:my-service-account`. `kubernetes_service_account` scope is always more specific than `kubernetes_namespace` scope for the same namespace.
        /// </summary>
        [Input("kubernetesServiceAccount")]
        public Input<string>? KubernetesServiceAccount { get; set; }

        public ScopeArgs()
        {
        }
        public static new ScopeArgs Empty => new ScopeArgs();
    }
}
