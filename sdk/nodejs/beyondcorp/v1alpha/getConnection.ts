// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Connection.
 */
export function getConnection(args: GetConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:beyondcorp/v1alpha:getConnection", {
        "connectionId": args.connectionId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetConnectionArgs {
    connectionId: string;
    location: string;
    project?: string;
}

export interface GetConnectionResult {
    /**
     * Address of the remote application endpoint for the BeyondCorp Connection.
     */
    readonly applicationEndpoint: outputs.beyondcorp.v1alpha.ApplicationEndpointResponse;
    /**
     * Optional. List of [google.cloud.beyondcorp.v1main.Connector.name] that are authorised to be associated with this Connection.
     */
    readonly connectors: string[];
    /**
     * Timestamp when the resource was created.
     */
    readonly createTime: string;
    /**
     * Optional. An arbitrary user-provided name for the connection. Cannot exceed 64 characters.
     */
    readonly displayName: string;
    /**
     * Optional. Gateway used by the connection.
     */
    readonly gateway: outputs.beyondcorp.v1alpha.GatewayResponse;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    readonly labels: {[key: string]: string};
    /**
     * Unique resource name of the connection. The name is ignored when creating a connection.
     */
    readonly name: string;
    /**
     * The current state of the connection.
     */
    readonly state: string;
    /**
     * The type of network connectivity used by the connection.
     */
    readonly type: string;
    /**
     * A unique identifier for the instance generated by the system.
     */
    readonly uid: string;
    /**
     * Timestamp when the resource was last modified.
     */
    readonly updateTime: string;
}

export function getConnectionOutput(args: GetConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionResult> {
    return pulumi.output(args).apply(a => getConnection(a, opts))
}

export interface GetConnectionOutputArgs {
    connectionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}