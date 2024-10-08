// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single FirewallEndpointAssociation.
 */
export function getFirewallEndpointAssociation(args: GetFirewallEndpointAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallEndpointAssociationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:networksecurity/v1beta1:getFirewallEndpointAssociation", {
        "firewallEndpointAssociationId": args.firewallEndpointAssociationId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetFirewallEndpointAssociationArgs {
    firewallEndpointAssociationId: string;
    location: string;
    project?: string;
}

export interface GetFirewallEndpointAssociationResult {
    /**
     * Create time stamp
     */
    readonly createTime: string;
    /**
     * The URL of the FirewallEndpoint that is being associated.
     */
    readonly firewallEndpoint: string;
    /**
     * Optional. Labels as key value pairs
     */
    readonly labels: {[key: string]: string};
    /**
     * name of resource
     */
    readonly name: string;
    /**
     * The URL of the network that is being associated.
     */
    readonly network: string;
    /**
     * Whether reconciling is in progress, recommended per https://google.aip.dev/128.
     */
    readonly reconciling: boolean;
    /**
     * Current state of the association.
     */
    readonly state: string;
    /**
     * Optional. The URL of the TlsInspectionPolicy that is being associated.
     */
    readonly tlsInspectionPolicy: string;
    /**
     * Update time stamp
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single FirewallEndpointAssociation.
 */
export function getFirewallEndpointAssociationOutput(args: GetFirewallEndpointAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallEndpointAssociationResult> {
    return pulumi.output(args).apply((a: any) => getFirewallEndpointAssociation(a, opts))
}

export interface GetFirewallEndpointAssociationOutputArgs {
    firewallEndpointAssociationId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
