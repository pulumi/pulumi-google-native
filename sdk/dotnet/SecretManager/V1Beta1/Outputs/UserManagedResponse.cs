// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SecretManager.V1Beta1.Outputs
{

    /// <summary>
    /// A replication policy that replicates the Secret payload into the locations specified in Secret.replication.user_managed.replicas
    /// </summary>
    [OutputType]
    public sealed class UserManagedResponse
    {
        /// <summary>
        /// The list of Replicas for this Secret. Cannot be empty.
        /// </summary>
        public readonly ImmutableArray<Outputs.ReplicaResponse> Replicas;

        [OutputConstructor]
        private UserManagedResponse(ImmutableArray<Outputs.ReplicaResponse> replicas)
        {
            Replicas = replicas;
        }
    }
}