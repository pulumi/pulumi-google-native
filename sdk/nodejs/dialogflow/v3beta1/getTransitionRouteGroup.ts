// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves the specified TransitionRouteGroup.
 */
export function getTransitionRouteGroup(args: GetTransitionRouteGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitionRouteGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:dialogflow/v3beta1:getTransitionRouteGroup", {
        "agentId": args.agentId,
        "flowId": args.flowId,
        "languageCode": args.languageCode,
        "location": args.location,
        "project": args.project,
        "transitionRouteGroupId": args.transitionRouteGroupId,
    }, opts);
}

export interface GetTransitionRouteGroupArgs {
    agentId: string;
    flowId: string;
    languageCode?: string;
    location: string;
    project?: string;
    transitionRouteGroupId: string;
}

export interface GetTransitionRouteGroupResult {
    /**
     * The human-readable name of the transition route group, unique within the flow. The display name can be no longer than 30 characters.
     */
    readonly displayName: string;
    /**
     * The unique identifier of the transition route group. TransitionRouteGroups.CreateTransitionRouteGroup populates the name automatically. Format: `projects//locations//agents//flows//transitionRouteGroups/`.
     */
    readonly name: string;
    /**
     * Transition routes associated with the TransitionRouteGroup.
     */
    readonly transitionRoutes: outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1TransitionRouteResponse[];
}

export function getTransitionRouteGroupOutput(args: GetTransitionRouteGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitionRouteGroupResult> {
    return pulumi.output(args).apply(a => getTransitionRouteGroup(a, opts))
}

export interface GetTransitionRouteGroupOutputArgs {
    agentId: pulumi.Input<string>;
    flowId: pulumi.Input<string>;
    languageCode?: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    transitionRouteGroupId: pulumi.Input<string>;
}