// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Get a single ChannelConnection.
 */
export function getChannelConnection(args: GetChannelConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetChannelConnectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:eventarc/v1:getChannelConnection", {
        "channelConnectionId": args.channelConnectionId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetChannelConnectionArgs {
    channelConnectionId: string;
    location: string;
    project?: string;
}

export interface GetChannelConnectionResult {
    /**
     * Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
     */
    readonly activationToken: string;
    /**
     * The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
     */
    readonly channel: string;
    /**
     * The creation time.
     */
    readonly createTime: string;
    /**
     * The name of the connection.
     */
    readonly name: string;
    /**
     * Server assigned ID of the resource. The server guarantees uniqueness and immutability until deleted.
     */
    readonly uid: string;
    /**
     * The last-modified time.
     */
    readonly updateTime: string;
}

export function getChannelConnectionOutput(args: GetChannelConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetChannelConnectionResult> {
    return pulumi.output(args).apply(a => getChannelConnection(a, opts))
}

export interface GetChannelConnectionOutputArgs {
    channelConnectionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}