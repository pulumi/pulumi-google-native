// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Outputs
{

    /// <summary>
    /// Represents configuration for the VMware VCenter for the user cluster.
    /// </summary>
    [OutputType]
    public sealed class VmwareVCenterConfigResponse
    {
        /// <summary>
        /// The vCenter IP address.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// Contains the vCenter CA certificate public key for SSL verification.
        /// </summary>
        public readonly string CaCertData;
        /// <summary>
        /// The name of the vCenter cluster for the user cluster.
        /// </summary>
        public readonly string Cluster;
        /// <summary>
        /// The name of the vCenter datacenter for the user cluster.
        /// </summary>
        public readonly string Datacenter;
        /// <summary>
        /// The name of the vCenter datastore for the user cluster.
        /// </summary>
        public readonly string Datastore;
        /// <summary>
        /// The name of the vCenter folder for the user cluster.
        /// </summary>
        public readonly string Folder;
        /// <summary>
        /// The name of the vCenter resource pool for the user cluster.
        /// </summary>
        public readonly string ResourcePool;

        [OutputConstructor]
        private VmwareVCenterConfigResponse(
            string address,

            string caCertData,

            string cluster,

            string datacenter,

            string datastore,

            string folder,

            string resourcePool)
        {
            Address = address;
            CaCertData = caCertData;
            Cluster = cluster;
            Datacenter = datacenter;
            Datastore = datastore;
            Folder = folder;
            ResourcePool = resourcePool;
        }
    }
}