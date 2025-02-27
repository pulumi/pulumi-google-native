// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single connection.
 */
export function getConnection(args: GetConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:cloudbuild/v2:getConnection", {
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
     * Allows clients to store small amounts of arbitrary data.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Server assigned timestamp for when the connection was created.
     */
    readonly createTime: string;
    /**
     * If disabled is set to true, functionality is disabled for this connection. Repository based API methods and webhooks processing for repositories in this connection will be disabled.
     */
    readonly disabled: boolean;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    readonly etag: string;
    /**
     * Configuration for connections to github.com.
     */
    readonly githubConfig: outputs.cloudbuild.v2.GitHubConfigResponse;
    /**
     * Configuration for connections to an instance of GitHub Enterprise.
     */
    readonly githubEnterpriseConfig: outputs.cloudbuild.v2.GoogleDevtoolsCloudbuildV2GitHubEnterpriseConfigResponse;
    /**
     * Configuration for connections to gitlab.com or an instance of GitLab Enterprise.
     */
    readonly gitlabConfig: outputs.cloudbuild.v2.GoogleDevtoolsCloudbuildV2GitLabConfigResponse;
    /**
     * Installation state of the Connection.
     */
    readonly installationState: outputs.cloudbuild.v2.InstallationStateResponse;
    /**
     * Immutable. The resource name of the connection, in the format `projects/{project}/locations/{location}/connections/{connection_id}`.
     */
    readonly name: string;
    /**
     * Set to true when the connection is being set up or updated in the background.
     */
    readonly reconciling: boolean;
    /**
     * Server assigned timestamp for when the connection was updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single connection.
 */
export function getConnectionOutput(args: GetConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionResult> {
    return pulumi.output(args).apply((a: any) => getConnection(a, opts))
}

export interface GetConnectionOutputArgs {
    connectionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
