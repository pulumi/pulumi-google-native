// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.NetworkManagement.V1Beta1.Outputs
{

    /// <summary>
    /// A simulated forwarding path is composed of multiple steps. Each step has a well-defined state and an associated configuration.
    /// </summary>
    [OutputType]
    public sealed class StepResponse
    {
        /// <summary>
        /// Display information of the final state "abort" and reason.
        /// </summary>
        public readonly Outputs.AbortInfoResponse Abort;
        /// <summary>
        /// Display information of an App Engine service version.
        /// </summary>
        public readonly Outputs.AppEngineVersionInfoResponse AppEngineVersion;
        /// <summary>
        /// This is a step that leads to the final state Drop.
        /// </summary>
        public readonly bool CausesDrop;
        /// <summary>
        /// Display information of a Cloud function.
        /// </summary>
        public readonly Outputs.CloudFunctionInfoResponse CloudFunction;
        /// <summary>
        /// Display information of a Cloud SQL instance.
        /// </summary>
        public readonly Outputs.CloudSQLInstanceInfoResponse CloudSqlInstance;
        /// <summary>
        /// Display information of the final state "deliver" and reason.
        /// </summary>
        public readonly Outputs.DeliverInfoResponse Deliver;
        /// <summary>
        /// A description of the step. Usually this is a summary of the state.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Display information of the final state "drop" and reason.
        /// </summary>
        public readonly Outputs.DropInfoResponse Drop;
        /// <summary>
        /// Display information of the source and destination under analysis. The endpoint information in an intermediate state may differ with the initial input, as it might be modified by state like NAT, or Connection Proxy.
        /// </summary>
        public readonly Outputs.EndpointInfoResponse Endpoint;
        /// <summary>
        /// Display information of a Compute Engine firewall rule.
        /// </summary>
        public readonly Outputs.FirewallInfoResponse Firewall;
        /// <summary>
        /// Display information of the final state "forward" and reason.
        /// </summary>
        public readonly Outputs.ForwardInfoResponse Forward;
        /// <summary>
        /// Display information of a Compute Engine forwarding rule.
        /// </summary>
        public readonly Outputs.ForwardingRuleInfoResponse ForwardingRule;
        /// <summary>
        /// Display information of a Google Kubernetes Engine cluster master.
        /// </summary>
        public readonly Outputs.GKEMasterInfoResponse GkeMaster;
        /// <summary>
        /// Display information of a Compute Engine instance.
        /// </summary>
        public readonly Outputs.InstanceInfoResponse Instance;
        /// <summary>
        /// Display information of the load balancers.
        /// </summary>
        public readonly Outputs.LoadBalancerInfoResponse LoadBalancer;
        /// <summary>
        /// Display information of a Google Cloud network.
        /// </summary>
        public readonly Outputs.NetworkInfoResponse Network;
        /// <summary>
        /// Project ID that contains the configuration this step is validating.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// Display information of a Compute Engine route.
        /// </summary>
        public readonly Outputs.RouteInfoResponse Route;
        /// <summary>
        /// Each step is in one of the pre-defined states.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Display information of a VPC connector.
        /// </summary>
        public readonly Outputs.VpcConnectorInfoResponse VpcConnector;
        /// <summary>
        /// Display information of a Compute Engine VPN gateway.
        /// </summary>
        public readonly Outputs.VpnGatewayInfoResponse VpnGateway;
        /// <summary>
        /// Display information of a Compute Engine VPN tunnel.
        /// </summary>
        public readonly Outputs.VpnTunnelInfoResponse VpnTunnel;

        [OutputConstructor]
        private StepResponse(
            Outputs.AbortInfoResponse abort,

            Outputs.AppEngineVersionInfoResponse appEngineVersion,

            bool causesDrop,

            Outputs.CloudFunctionInfoResponse cloudFunction,

            Outputs.CloudSQLInstanceInfoResponse cloudSqlInstance,

            Outputs.DeliverInfoResponse deliver,

            string description,

            Outputs.DropInfoResponse drop,

            Outputs.EndpointInfoResponse endpoint,

            Outputs.FirewallInfoResponse firewall,

            Outputs.ForwardInfoResponse forward,

            Outputs.ForwardingRuleInfoResponse forwardingRule,

            Outputs.GKEMasterInfoResponse gkeMaster,

            Outputs.InstanceInfoResponse instance,

            Outputs.LoadBalancerInfoResponse loadBalancer,

            Outputs.NetworkInfoResponse network,

            string project,

            Outputs.RouteInfoResponse route,

            string state,

            Outputs.VpcConnectorInfoResponse vpcConnector,

            Outputs.VpnGatewayInfoResponse vpnGateway,

            Outputs.VpnTunnelInfoResponse vpnTunnel)
        {
            Abort = abort;
            AppEngineVersion = appEngineVersion;
            CausesDrop = causesDrop;
            CloudFunction = cloudFunction;
            CloudSqlInstance = cloudSqlInstance;
            Deliver = deliver;
            Description = description;
            Drop = drop;
            Endpoint = endpoint;
            Firewall = firewall;
            Forward = forward;
            ForwardingRule = forwardingRule;
            GkeMaster = gkeMaster;
            Instance = instance;
            LoadBalancer = loadBalancer;
            Network = network;
            Project = project;
            Route = route;
            State = state;
            VpcConnector = vpcConnector;
            VpnGateway = vpnGateway;
            VpnTunnel = vpnTunnel;
        }
    }
}