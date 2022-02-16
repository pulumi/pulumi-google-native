// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a session entity type.
 */
export class SessionEntityType extends pulumi.CustomResource {
    /**
     * Get an existing SessionEntityType resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SessionEntityType {
        return new SessionEntityType(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3beta1:SessionEntityType';

    /**
     * Returns true if the given object is an instance of SessionEntityType.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SessionEntityType {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SessionEntityType.__pulumiType;
    }

    /**
     * The collection of entities to override or supplement the custom entity type.
     */
    public readonly entities!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EntityTypeEntityResponse[]>;
    /**
     * Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    public readonly entityOverrideMode!: pulumi.Output<string>;
    /**
     * The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a SessionEntityType resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SessionEntityTypeArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.entities === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entities'");
            }
            if ((!args || args.entityOverrideMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityOverrideMode'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.sessionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sessionId'");
            }
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["entities"] = args ? args.entities : undefined;
            resourceInputs["entityOverrideMode"] = args ? args.entityOverrideMode : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["sessionId"] = args ? args.sessionId : undefined;
        } else {
            resourceInputs["entities"] = undefined /*out*/;
            resourceInputs["entityOverrideMode"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SessionEntityType.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SessionEntityType resource.
 */
export interface SessionEntityTypeArgs {
    agentId: pulumi.Input<string>;
    /**
     * The collection of entities to override or supplement the custom entity type.
     */
    entities: pulumi.Input<pulumi.Input<inputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1EntityTypeEntityArgs>[]>;
    /**
     * Indicates whether the additional data should override or supplement the custom entity type definition.
     */
    entityOverrideMode: pulumi.Input<enums.dialogflow.v3beta1.SessionEntityTypeEntityOverrideMode>;
    environmentId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The unique identifier of the session entity type. Format: `projects//locations//agents//sessions//entityTypes/` or `projects//locations//agents//environments//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sessionId: pulumi.Input<string>;
}