// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta
{
    public static class GetWorkstationCluster
    {
        /// <summary>
        /// Returns the requested workstation cluster.
        /// </summary>
        public static Task<GetWorkstationClusterResult> InvokeAsync(GetWorkstationClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetWorkstationClusterResult>("google-native:workstations/v1beta:getWorkstationCluster", args ?? new GetWorkstationClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Returns the requested workstation cluster.
        /// </summary>
        public static Output<GetWorkstationClusterResult> Invoke(GetWorkstationClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetWorkstationClusterResult>("google-native:workstations/v1beta:getWorkstationCluster", args ?? new GetWorkstationClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetWorkstationClusterArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("workstationClusterId", required: true)]
        public string WorkstationClusterId { get; set; } = null!;

        public GetWorkstationClusterArgs()
        {
        }
        public static new GetWorkstationClusterArgs Empty => new GetWorkstationClusterArgs();
    }

    public sealed class GetWorkstationClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("workstationClusterId", required: true)]
        public Input<string> WorkstationClusterId { get; set; } = null!;

        public GetWorkstationClusterInvokeArgs()
        {
        }
        public static new GetWorkstationClusterInvokeArgs Empty => new GetWorkstationClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetWorkstationClusterResult
    {
        /// <summary>
        /// Client-specified annotations.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Annotations;
        /// <summary>
        /// Status conditions describing the current resource state.
        /// </summary>
        public readonly ImmutableArray<Outputs.StatusResponse> Conditions;
        /// <summary>
        /// The private IP address of the control plane for this cluster. Workstation VMs need access to this IP address to work with the service, so make sure that your firewall rules allow egress from the workstation VMs to this address.
        /// </summary>
        public readonly string ControlPlaneIp;
        /// <summary>
        /// Time when this resource was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality. Details can be found in the `conditions` field.
        /// </summary>
        public readonly bool Degraded;
        /// <summary>
        /// Time when this resource was soft-deleted.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// Human-readable name for this resource.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Checksum computed by the server. May be sent on update and delete requests to make sure that the client has an up-to-date value before proceeding.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Full name of this resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. Name of the Compute Engine network in which instances associated with this cluster will be created.
        /// </summary>
        public readonly string Network;
        /// <summary>
        /// Configuration for private cluster.
        /// </summary>
        public readonly Outputs.PrivateClusterConfigResponse PrivateClusterConfig;
        /// <summary>
        /// Indicates whether this resource is currently being updated to match its intended state.
        /// </summary>
        public readonly bool Reconciling;
        /// <summary>
        /// Immutable. Name of the Compute Engine subnetwork in which instances associated with this cluster will be created. Must be part of the subnetwork specified for this cluster.
        /// </summary>
        public readonly string Subnetwork;
        /// <summary>
        /// A system-assigned unique identified for this resource.
        /// </summary>
        public readonly string Uid;
        /// <summary>
        /// Time when this resource was most recently updated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetWorkstationClusterResult(
            ImmutableDictionary<string, string> annotations,

            ImmutableArray<Outputs.StatusResponse> conditions,

            string controlPlaneIp,

            string createTime,

            bool degraded,

            string deleteTime,

            string displayName,

            string etag,

            ImmutableDictionary<string, string> labels,

            string name,

            string network,

            Outputs.PrivateClusterConfigResponse privateClusterConfig,

            bool reconciling,

            string subnetwork,

            string uid,

            string updateTime)
        {
            Annotations = annotations;
            Conditions = conditions;
            ControlPlaneIp = controlPlaneIp;
            CreateTime = createTime;
            Degraded = degraded;
            DeleteTime = deleteTime;
            DisplayName = displayName;
            Etag = etag;
            Labels = labels;
            Name = name;
            Network = network;
            PrivateClusterConfig = privateClusterConfig;
            Reconciling = reconciling;
            Subnetwork = subnetwork;
            Uid = uid;
            UpdateTime = updateTime;
        }
    }
}