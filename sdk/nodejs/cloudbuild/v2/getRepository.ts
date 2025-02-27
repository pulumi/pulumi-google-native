// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single repository.
 */
export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:cloudbuild/v2:getRepository", {
        "connectionId": args.connectionId,
        "location": args.location,
        "project": args.project,
        "repositoryId": args.repositoryId,
    }, opts);
}

export interface GetRepositoryArgs {
    connectionId: string;
    location: string;
    project?: string;
    repositoryId: string;
}

export interface GetRepositoryResult {
    /**
     * Allows clients to store small amounts of arbitrary data.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Server assigned timestamp for when the connection was created.
     */
    readonly createTime: string;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    readonly etag: string;
    /**
     * Immutable. Resource name of the repository, in the format `projects/*&#47;locations/*&#47;connections/*&#47;repositories/*`.
     */
    readonly name: string;
    /**
     * Git Clone HTTPS URI.
     */
    readonly remoteUri: string;
    /**
     * Server assigned timestamp for when the connection was updated.
     */
    readonly updateTime: string;
    /**
     * External ID of the webhook created for the repository.
     */
    readonly webhookId: string;
}
/**
 * Gets details of a single repository.
 */
export function getRepositoryOutput(args: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getRepository(a, opts))
}

export interface GetRepositoryOutputArgs {
    connectionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    repositoryId: pulumi.Input<string>;
}
