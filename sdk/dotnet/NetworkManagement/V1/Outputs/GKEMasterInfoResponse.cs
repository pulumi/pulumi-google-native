// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1.Outputs
{

    /// <summary>
    /// For display only. Metadata associated with a Google Kubernetes Engine (GKE) cluster master.
    /// </summary>
    [OutputType]
    public sealed class GKEMasterInfoResponse
    {
        /// <summary>
        /// URI of a GKE cluster network.
        /// </summary>
        public readonly string ClusterNetworkUri;
        /// <summary>
        /// URI of a GKE cluster.
        /// </summary>
        public readonly string ClusterUri;
        /// <summary>
        /// External IP address of a GKE cluster master.
        /// </summary>
        public readonly string ExternalIp;
        /// <summary>
        /// Internal IP address of a GKE cluster master.
        /// </summary>
        public readonly string InternalIp;

        [OutputConstructor]
        private GKEMasterInfoResponse(
            string clusterNetworkUri,

            string clusterUri,

            string externalIp,

            string internalIp)
        {
            ClusterNetworkUri = clusterNetworkUri;
            ClusterUri = clusterUri;
            ExternalIp = externalIp;
            InternalIp = internalIp;
        }
    }
}
