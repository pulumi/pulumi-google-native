// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified Version.
 */
export function getVersion(args: GetVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetVersionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dialogflow/v3:getVersion", {
        "agentId": args.agentId,
        "flowId": args.flowId,
        "location": args.location,
        "project": args.project,
        "versionId": args.versionId,
    }, opts);
}

export interface GetVersionArgs {
    agentId: string;
    flowId: string;
    location: string;
    project?: string;
    versionId: string;
}

export interface GetVersionResult {
    /**
     * Create time of the version.
     */
    readonly createTime: string;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    readonly description: string;
    /**
     * The human-readable name of the version. Limit of 64 characters.
     */
    readonly displayName: string;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
     */
    readonly name: string;
    /**
     * The NLU settings of the flow at version creation.
     */
    readonly nluSettings: outputs.dialogflow.v3.GoogleCloudDialogflowCxV3NluSettingsResponse;
    /**
     * The state of this version. This field is read-only and cannot be set by create and update methods.
     */
    readonly state: string;
}
/**
 * Retrieves the specified Version.
 */
export function getVersionOutput(args: GetVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVersionResult> {
    return pulumi.output(args).apply((a: any) => getVersion(a, opts))
}

export interface GetVersionOutputArgs {
    agentId: pulumi.Input<string>;
    flowId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    versionId: pulumi.Input<string>;
}
