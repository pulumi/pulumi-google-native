// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves a `NetworkPolicy` resource by its resource name.
 */
export function getNetworkPolicy(args: GetNetworkPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:vmwareengine/v1:getNetworkPolicy", {
        "location": args.location,
        "networkPolicyId": args.networkPolicyId,
        "project": args.project,
    }, opts);
}

export interface GetNetworkPolicyArgs {
    location: string;
    networkPolicyId: string;
    project?: string;
}

export interface GetNetworkPolicyResult {
    /**
     * Creation time of this resource.
     */
    readonly createTime: string;
    /**
     * Optional. User-provided description for this network policy.
     */
    readonly description: string;
    /**
     * IP address range in CIDR notation used to create internet access and external IP access. An RFC 1918 CIDR block, with a "/26" prefix, is required. The range cannot overlap with any prefixes either in the consumer VPC network or in use by the private clouds attached to that VPC network.
     */
    readonly edgeServicesCidr: string;
    /**
     * Network service that allows External IP addresses to be assigned to VMware workloads. This service can only be enabled when `internet_access` is also enabled.
     */
    readonly externalIp: outputs.vmwareengine.v1.NetworkServiceResponse;
    /**
     * Network service that allows VMware workloads to access the internet.
     */
    readonly internetAccess: outputs.vmwareengine.v1.NetworkServiceResponse;
    /**
     * The resource name of this network policy. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-central1/networkPolicies/my-network-policy`
     */
    readonly name: string;
    /**
     * System-generated unique identifier for the resource.
     */
    readonly uid: string;
    /**
     * Last update time of this resource.
     */
    readonly updateTime: string;
    /**
     * Optional. The relative resource name of the VMware Engine network. Specify the name in the following form: `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}` where `{project}` can either be a project number or a project ID.
     */
    readonly vmwareEngineNetwork: string;
    /**
     * The canonical name of the VMware Engine network in the form: `projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}`
     */
    readonly vmwareEngineNetworkCanonical: string;
}
/**
 * Retrieves a `NetworkPolicy` resource by its resource name.
 */
export function getNetworkPolicyOutput(args: GetNetworkPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkPolicyResult> {
    return pulumi.output(args).apply((a: any) => getNetworkPolicy(a, opts))
}

export interface GetNetworkPolicyOutputArgs {
    location: pulumi.Input<string>;
    networkPolicyId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}