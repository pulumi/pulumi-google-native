// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a Conversation.
 */
export function getConversation(args: GetConversationArgs, opts?: pulumi.InvokeOptions): Promise<GetConversationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:discoveryengine/v1alpha:getConversation", {
        "collectionId": args.collectionId,
        "conversationId": args.conversationId,
        "dataStoreId": args.dataStoreId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetConversationArgs {
    collectionId: string;
    conversationId: string;
    dataStoreId: string;
    location: string;
    project?: string;
}

export interface GetConversationResult {
    /**
     * The time the conversation finished.
     */
    readonly endTime: string;
    /**
     * Conversation messages.
     */
    readonly messages: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaConversationMessageResponse[];
    /**
     * Immutable. Fully qualified name `project/*&#47;locations/global/collections/{collection}/dataStore/*&#47;conversations/*`
     */
    readonly name: string;
    /**
     * The time the conversation started.
     */
    readonly startTime: string;
    /**
     * The state of the Conversation.
     */
    readonly state: string;
    /**
     * A unique identifier for tracking users.
     */
    readonly userPseudoId: string;
}
/**
 * Gets a Conversation.
 */
export function getConversationOutput(args: GetConversationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConversationResult> {
    return pulumi.output(args).apply((a: any) => getConversation(a, opts))
}

export interface GetConversationOutputArgs {
    collectionId: pulumi.Input<string>;
    conversationId: pulumi.Input<string>;
    dataStoreId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}