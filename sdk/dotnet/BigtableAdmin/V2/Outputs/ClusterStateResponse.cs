// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// The state of a table's data in a particular cluster.
    /// </summary>
    [OutputType]
    public sealed class ClusterStateResponse
    {
        /// <summary>
        /// The encryption information for the table in this cluster. If the encryption key protecting this resource is customer managed, then its version can be rotated in Cloud Key Management Service (Cloud KMS). The primary version of the key and its status will be reflected here when changes propagate from Cloud KMS.
        /// </summary>
        public readonly ImmutableArray<Outputs.EncryptionInfoResponse> EncryptionInfo;
        /// <summary>
        /// The state of replication for the table in this cluster.
        /// </summary>
        public readonly string ReplicationState;

        [OutputConstructor]
        private ClusterStateResponse(
            ImmutableArray<Outputs.EncryptionInfoResponse> encryptionInfo,

            string replicationState)
        {
            EncryptionInfo = encryptionInfo;
            ReplicationState = replicationState;
        }
    }
}
