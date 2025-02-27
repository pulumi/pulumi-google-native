// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a Version in the specified Flow. This method is a [long-running operation](https://cloud.google.com/dialogflow/cx/docs/how/long-running-operation). The returned `Operation` type has the following method-specific fields: - `metadata`: CreateVersionOperationMetadata - `response`: Version
 */
export class Version extends pulumi.CustomResource {
    /**
     * Get an existing Version resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Version {
        return new Version(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dialogflow/v3beta1:Version';

    /**
     * Returns true if the given object is an instance of Version.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Version {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Version.__pulumiType;
    }

    public readonly agentId!: pulumi.Output<string>;
    /**
     * Create time of the version.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The human-readable name of the version. Limit of 64 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    public readonly flowId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The NLU settings of the flow at version creation.
     */
    public /*out*/ readonly nluSettings!: pulumi.Output<outputs.dialogflow.v3beta1.GoogleCloudDialogflowCxV3beta1NluSettingsResponse>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The state of this version. This field is read-only and cannot be set by create and update methods.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;

    /**
     * Create a Version resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.agentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'agentId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.flowId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flowId'");
            }
            resourceInputs["agentId"] = args ? args.agentId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["flowId"] = args ? args.flowId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["nluSettings"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["agentId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["flowId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["nluSettings"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["agentId", "flowId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Version.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Version resource.
 */
export interface VersionArgs {
    agentId: pulumi.Input<string>;
    /**
     * The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
     */
    description?: pulumi.Input<string>;
    /**
     * The human-readable name of the version. Limit of 64 characters.
     */
    displayName: pulumi.Input<string>;
    flowId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
