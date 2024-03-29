// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ClusterArgs } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { ExternalAccessRuleArgs } from "./externalAccessRule";
export type ExternalAccessRule = import("./externalAccessRule").ExternalAccessRule;
export const ExternalAccessRule: typeof import("./externalAccessRule").ExternalAccessRule = null as any;
utilities.lazyLoad(exports, ["ExternalAccessRule"], () => require("./externalAccessRule"));

export { ExternalAddressArgs } from "./externalAddress";
export type ExternalAddress = import("./externalAddress").ExternalAddress;
export const ExternalAddress: typeof import("./externalAddress").ExternalAddress = null as any;
utilities.lazyLoad(exports, ["ExternalAddress"], () => require("./externalAddress"));

export { GetClusterArgs, GetClusterResult, GetClusterOutputArgs } from "./getCluster";
export const getCluster: typeof import("./getCluster").getCluster = null as any;
export const getClusterOutput: typeof import("./getCluster").getClusterOutput = null as any;
utilities.lazyLoad(exports, ["getCluster","getClusterOutput"], () => require("./getCluster"));

export { GetExternalAccessRuleArgs, GetExternalAccessRuleResult, GetExternalAccessRuleOutputArgs } from "./getExternalAccessRule";
export const getExternalAccessRule: typeof import("./getExternalAccessRule").getExternalAccessRule = null as any;
export const getExternalAccessRuleOutput: typeof import("./getExternalAccessRule").getExternalAccessRuleOutput = null as any;
utilities.lazyLoad(exports, ["getExternalAccessRule","getExternalAccessRuleOutput"], () => require("./getExternalAccessRule"));

export { GetExternalAddressArgs, GetExternalAddressResult, GetExternalAddressOutputArgs } from "./getExternalAddress";
export const getExternalAddress: typeof import("./getExternalAddress").getExternalAddress = null as any;
export const getExternalAddressOutput: typeof import("./getExternalAddress").getExternalAddressOutput = null as any;
utilities.lazyLoad(exports, ["getExternalAddress","getExternalAddressOutput"], () => require("./getExternalAddress"));

export { GetHcxActivationKeyArgs, GetHcxActivationKeyResult, GetHcxActivationKeyOutputArgs } from "./getHcxActivationKey";
export const getHcxActivationKey: typeof import("./getHcxActivationKey").getHcxActivationKey = null as any;
export const getHcxActivationKeyOutput: typeof import("./getHcxActivationKey").getHcxActivationKeyOutput = null as any;
utilities.lazyLoad(exports, ["getHcxActivationKey","getHcxActivationKeyOutput"], () => require("./getHcxActivationKey"));

export { GetLoggingServerArgs, GetLoggingServerResult, GetLoggingServerOutputArgs } from "./getLoggingServer";
export const getLoggingServer: typeof import("./getLoggingServer").getLoggingServer = null as any;
export const getLoggingServerOutput: typeof import("./getLoggingServer").getLoggingServerOutput = null as any;
utilities.lazyLoad(exports, ["getLoggingServer","getLoggingServerOutput"], () => require("./getLoggingServer"));

export { GetManagementDnsZoneBindingArgs, GetManagementDnsZoneBindingResult, GetManagementDnsZoneBindingOutputArgs } from "./getManagementDnsZoneBinding";
export const getManagementDnsZoneBinding: typeof import("./getManagementDnsZoneBinding").getManagementDnsZoneBinding = null as any;
export const getManagementDnsZoneBindingOutput: typeof import("./getManagementDnsZoneBinding").getManagementDnsZoneBindingOutput = null as any;
utilities.lazyLoad(exports, ["getManagementDnsZoneBinding","getManagementDnsZoneBindingOutput"], () => require("./getManagementDnsZoneBinding"));

export { GetNetworkPeeringArgs, GetNetworkPeeringResult, GetNetworkPeeringOutputArgs } from "./getNetworkPeering";
export const getNetworkPeering: typeof import("./getNetworkPeering").getNetworkPeering = null as any;
export const getNetworkPeeringOutput: typeof import("./getNetworkPeering").getNetworkPeeringOutput = null as any;
utilities.lazyLoad(exports, ["getNetworkPeering","getNetworkPeeringOutput"], () => require("./getNetworkPeering"));

export { GetNetworkPolicyArgs, GetNetworkPolicyResult, GetNetworkPolicyOutputArgs } from "./getNetworkPolicy";
export const getNetworkPolicy: typeof import("./getNetworkPolicy").getNetworkPolicy = null as any;
export const getNetworkPolicyOutput: typeof import("./getNetworkPolicy").getNetworkPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getNetworkPolicy","getNetworkPolicyOutput"], () => require("./getNetworkPolicy"));

export { GetPrivateCloudArgs, GetPrivateCloudResult, GetPrivateCloudOutputArgs } from "./getPrivateCloud";
export const getPrivateCloud: typeof import("./getPrivateCloud").getPrivateCloud = null as any;
export const getPrivateCloudOutput: typeof import("./getPrivateCloud").getPrivateCloudOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateCloud","getPrivateCloudOutput"], () => require("./getPrivateCloud"));

export { GetPrivateCloudClusterIamPolicyArgs, GetPrivateCloudClusterIamPolicyResult, GetPrivateCloudClusterIamPolicyOutputArgs } from "./getPrivateCloudClusterIamPolicy";
export const getPrivateCloudClusterIamPolicy: typeof import("./getPrivateCloudClusterIamPolicy").getPrivateCloudClusterIamPolicy = null as any;
export const getPrivateCloudClusterIamPolicyOutput: typeof import("./getPrivateCloudClusterIamPolicy").getPrivateCloudClusterIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateCloudClusterIamPolicy","getPrivateCloudClusterIamPolicyOutput"], () => require("./getPrivateCloudClusterIamPolicy"));

export { GetPrivateCloudHcxActivationKeyIamPolicyArgs, GetPrivateCloudHcxActivationKeyIamPolicyResult, GetPrivateCloudHcxActivationKeyIamPolicyOutputArgs } from "./getPrivateCloudHcxActivationKeyIamPolicy";
export const getPrivateCloudHcxActivationKeyIamPolicy: typeof import("./getPrivateCloudHcxActivationKeyIamPolicy").getPrivateCloudHcxActivationKeyIamPolicy = null as any;
export const getPrivateCloudHcxActivationKeyIamPolicyOutput: typeof import("./getPrivateCloudHcxActivationKeyIamPolicy").getPrivateCloudHcxActivationKeyIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateCloudHcxActivationKeyIamPolicy","getPrivateCloudHcxActivationKeyIamPolicyOutput"], () => require("./getPrivateCloudHcxActivationKeyIamPolicy"));

export { GetPrivateCloudIamPolicyArgs, GetPrivateCloudIamPolicyResult, GetPrivateCloudIamPolicyOutputArgs } from "./getPrivateCloudIamPolicy";
export const getPrivateCloudIamPolicy: typeof import("./getPrivateCloudIamPolicy").getPrivateCloudIamPolicy = null as any;
export const getPrivateCloudIamPolicyOutput: typeof import("./getPrivateCloudIamPolicy").getPrivateCloudIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateCloudIamPolicy","getPrivateCloudIamPolicyOutput"], () => require("./getPrivateCloudIamPolicy"));

export { GetPrivateConnectionArgs, GetPrivateConnectionResult, GetPrivateConnectionOutputArgs } from "./getPrivateConnection";
export const getPrivateConnection: typeof import("./getPrivateConnection").getPrivateConnection = null as any;
export const getPrivateConnectionOutput: typeof import("./getPrivateConnection").getPrivateConnectionOutput = null as any;
utilities.lazyLoad(exports, ["getPrivateConnection","getPrivateConnectionOutput"], () => require("./getPrivateConnection"));

export { GetVmwareEngineNetworkArgs, GetVmwareEngineNetworkResult, GetVmwareEngineNetworkOutputArgs } from "./getVmwareEngineNetwork";
export const getVmwareEngineNetwork: typeof import("./getVmwareEngineNetwork").getVmwareEngineNetwork = null as any;
export const getVmwareEngineNetworkOutput: typeof import("./getVmwareEngineNetwork").getVmwareEngineNetworkOutput = null as any;
utilities.lazyLoad(exports, ["getVmwareEngineNetwork","getVmwareEngineNetworkOutput"], () => require("./getVmwareEngineNetwork"));

export { HcxActivationKeyArgs } from "./hcxActivationKey";
export type HcxActivationKey = import("./hcxActivationKey").HcxActivationKey;
export const HcxActivationKey: typeof import("./hcxActivationKey").HcxActivationKey = null as any;
utilities.lazyLoad(exports, ["HcxActivationKey"], () => require("./hcxActivationKey"));

export { LoggingServerArgs } from "./loggingServer";
export type LoggingServer = import("./loggingServer").LoggingServer;
export const LoggingServer: typeof import("./loggingServer").LoggingServer = null as any;
utilities.lazyLoad(exports, ["LoggingServer"], () => require("./loggingServer"));

export { ManagementDnsZoneBindingArgs } from "./managementDnsZoneBinding";
export type ManagementDnsZoneBinding = import("./managementDnsZoneBinding").ManagementDnsZoneBinding;
export const ManagementDnsZoneBinding: typeof import("./managementDnsZoneBinding").ManagementDnsZoneBinding = null as any;
utilities.lazyLoad(exports, ["ManagementDnsZoneBinding"], () => require("./managementDnsZoneBinding"));

export { NetworkPeeringArgs } from "./networkPeering";
export type NetworkPeering = import("./networkPeering").NetworkPeering;
export const NetworkPeering: typeof import("./networkPeering").NetworkPeering = null as any;
utilities.lazyLoad(exports, ["NetworkPeering"], () => require("./networkPeering"));

export { NetworkPolicyArgs } from "./networkPolicy";
export type NetworkPolicy = import("./networkPolicy").NetworkPolicy;
export const NetworkPolicy: typeof import("./networkPolicy").NetworkPolicy = null as any;
utilities.lazyLoad(exports, ["NetworkPolicy"], () => require("./networkPolicy"));

export { PrivateCloudArgs } from "./privateCloud";
export type PrivateCloud = import("./privateCloud").PrivateCloud;
export const PrivateCloud: typeof import("./privateCloud").PrivateCloud = null as any;
utilities.lazyLoad(exports, ["PrivateCloud"], () => require("./privateCloud"));

export { PrivateCloudClusterIamBindingArgs } from "./privateCloudClusterIamBinding";
export type PrivateCloudClusterIamBinding = import("./privateCloudClusterIamBinding").PrivateCloudClusterIamBinding;
export const PrivateCloudClusterIamBinding: typeof import("./privateCloudClusterIamBinding").PrivateCloudClusterIamBinding = null as any;
utilities.lazyLoad(exports, ["PrivateCloudClusterIamBinding"], () => require("./privateCloudClusterIamBinding"));

export { PrivateCloudClusterIamMemberArgs } from "./privateCloudClusterIamMember";
export type PrivateCloudClusterIamMember = import("./privateCloudClusterIamMember").PrivateCloudClusterIamMember;
export const PrivateCloudClusterIamMember: typeof import("./privateCloudClusterIamMember").PrivateCloudClusterIamMember = null as any;
utilities.lazyLoad(exports, ["PrivateCloudClusterIamMember"], () => require("./privateCloudClusterIamMember"));

export { PrivateCloudClusterIamPolicyArgs } from "./privateCloudClusterIamPolicy";
export type PrivateCloudClusterIamPolicy = import("./privateCloudClusterIamPolicy").PrivateCloudClusterIamPolicy;
export const PrivateCloudClusterIamPolicy: typeof import("./privateCloudClusterIamPolicy").PrivateCloudClusterIamPolicy = null as any;
utilities.lazyLoad(exports, ["PrivateCloudClusterIamPolicy"], () => require("./privateCloudClusterIamPolicy"));

export { PrivateCloudHcxActivationKeyIamBindingArgs } from "./privateCloudHcxActivationKeyIamBinding";
export type PrivateCloudHcxActivationKeyIamBinding = import("./privateCloudHcxActivationKeyIamBinding").PrivateCloudHcxActivationKeyIamBinding;
export const PrivateCloudHcxActivationKeyIamBinding: typeof import("./privateCloudHcxActivationKeyIamBinding").PrivateCloudHcxActivationKeyIamBinding = null as any;
utilities.lazyLoad(exports, ["PrivateCloudHcxActivationKeyIamBinding"], () => require("./privateCloudHcxActivationKeyIamBinding"));

export { PrivateCloudHcxActivationKeyIamMemberArgs } from "./privateCloudHcxActivationKeyIamMember";
export type PrivateCloudHcxActivationKeyIamMember = import("./privateCloudHcxActivationKeyIamMember").PrivateCloudHcxActivationKeyIamMember;
export const PrivateCloudHcxActivationKeyIamMember: typeof import("./privateCloudHcxActivationKeyIamMember").PrivateCloudHcxActivationKeyIamMember = null as any;
utilities.lazyLoad(exports, ["PrivateCloudHcxActivationKeyIamMember"], () => require("./privateCloudHcxActivationKeyIamMember"));

export { PrivateCloudHcxActivationKeyIamPolicyArgs } from "./privateCloudHcxActivationKeyIamPolicy";
export type PrivateCloudHcxActivationKeyIamPolicy = import("./privateCloudHcxActivationKeyIamPolicy").PrivateCloudHcxActivationKeyIamPolicy;
export const PrivateCloudHcxActivationKeyIamPolicy: typeof import("./privateCloudHcxActivationKeyIamPolicy").PrivateCloudHcxActivationKeyIamPolicy = null as any;
utilities.lazyLoad(exports, ["PrivateCloudHcxActivationKeyIamPolicy"], () => require("./privateCloudHcxActivationKeyIamPolicy"));

export { PrivateCloudIamBindingArgs } from "./privateCloudIamBinding";
export type PrivateCloudIamBinding = import("./privateCloudIamBinding").PrivateCloudIamBinding;
export const PrivateCloudIamBinding: typeof import("./privateCloudIamBinding").PrivateCloudIamBinding = null as any;
utilities.lazyLoad(exports, ["PrivateCloudIamBinding"], () => require("./privateCloudIamBinding"));

export { PrivateCloudIamMemberArgs } from "./privateCloudIamMember";
export type PrivateCloudIamMember = import("./privateCloudIamMember").PrivateCloudIamMember;
export const PrivateCloudIamMember: typeof import("./privateCloudIamMember").PrivateCloudIamMember = null as any;
utilities.lazyLoad(exports, ["PrivateCloudIamMember"], () => require("./privateCloudIamMember"));

export { PrivateCloudIamPolicyArgs } from "./privateCloudIamPolicy";
export type PrivateCloudIamPolicy = import("./privateCloudIamPolicy").PrivateCloudIamPolicy;
export const PrivateCloudIamPolicy: typeof import("./privateCloudIamPolicy").PrivateCloudIamPolicy = null as any;
utilities.lazyLoad(exports, ["PrivateCloudIamPolicy"], () => require("./privateCloudIamPolicy"));

export { PrivateConnectionArgs } from "./privateConnection";
export type PrivateConnection = import("./privateConnection").PrivateConnection;
export const PrivateConnection: typeof import("./privateConnection").PrivateConnection = null as any;
utilities.lazyLoad(exports, ["PrivateConnection"], () => require("./privateConnection"));

export { VmwareEngineNetworkArgs } from "./vmwareEngineNetwork";
export type VmwareEngineNetwork = import("./vmwareEngineNetwork").VmwareEngineNetwork;
export const VmwareEngineNetwork: typeof import("./vmwareEngineNetwork").VmwareEngineNetwork = null as any;
utilities.lazyLoad(exports, ["VmwareEngineNetwork"], () => require("./vmwareEngineNetwork"));


// Export enums:
export * from "../../types/enums/vmwareengine/v1";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:vmwareengine/v1:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:ExternalAccessRule":
                return new ExternalAccessRule(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:ExternalAddress":
                return new ExternalAddress(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:HcxActivationKey":
                return new HcxActivationKey(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:LoggingServer":
                return new LoggingServer(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:ManagementDnsZoneBinding":
                return new ManagementDnsZoneBinding(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:NetworkPeering":
                return new NetworkPeering(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:NetworkPolicy":
                return new NetworkPolicy(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloud":
                return new PrivateCloud(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudClusterIamBinding":
                return new PrivateCloudClusterIamBinding(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudClusterIamMember":
                return new PrivateCloudClusterIamMember(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudClusterIamPolicy":
                return new PrivateCloudClusterIamPolicy(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudHcxActivationKeyIamBinding":
                return new PrivateCloudHcxActivationKeyIamBinding(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudHcxActivationKeyIamMember":
                return new PrivateCloudHcxActivationKeyIamMember(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudHcxActivationKeyIamPolicy":
                return new PrivateCloudHcxActivationKeyIamPolicy(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudIamBinding":
                return new PrivateCloudIamBinding(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudIamMember":
                return new PrivateCloudIamMember(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateCloudIamPolicy":
                return new PrivateCloudIamPolicy(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:PrivateConnection":
                return new PrivateConnection(name, <any>undefined, { urn })
            case "google-native:vmwareengine/v1:VmwareEngineNetwork":
                return new VmwareEngineNetwork(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "vmwareengine/v1", _module)
