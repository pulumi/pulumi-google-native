// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Aiplatform.V1Beta1.Outputs
{

    /// <summary>
    /// A description of resources that to large degree are decided by Vertex AI, and require only a modest additional configuration. Each Model supporting these resources documents its specific guidelines.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudAiplatformV1beta1AutomaticResourcesResponse
    {
        /// <summary>
        /// Immutable. The maximum number of replicas this DeployedModel may be deployed on when the traffic against it increases. If the requested value is too large, the deployment will error, but if deployment succeeds then the ability to scale the model to that many replicas is guaranteed (barring service outages). If traffic against the DeployedModel increases beyond what its replicas at maximum may handle, a portion of the traffic will be dropped. If this value is not provided, a no upper bound for scaling under heavy traffic will be assume, though Vertex AI may be unable to scale beyond certain replica number.
        /// </summary>
        public readonly int MaxReplicaCount;
        /// <summary>
        /// Immutable. The minimum number of replicas this DeployedModel will be always deployed on. If traffic against it increases, it may dynamically be deployed onto more replicas up to max_replica_count, and as traffic decreases, some of these extra replicas may be freed. If the requested value is too large, the deployment will error.
        /// </summary>
        public readonly int MinReplicaCount;

        [OutputConstructor]
        private GoogleCloudAiplatformV1beta1AutomaticResourcesResponse(
            int maxReplicaCount,

            int minReplicaCount)
        {
            MaxReplicaCount = maxReplicaCount;
            MinReplicaCount = minReplicaCount;
        }
    }
}