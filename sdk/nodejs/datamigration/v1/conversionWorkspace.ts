// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new conversion workspace in a given project and location.
 */
export class ConversionWorkspace extends pulumi.CustomResource {
    /**
     * Get an existing ConversionWorkspace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConversionWorkspace {
        return new ConversionWorkspace(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datamigration/v1:ConversionWorkspace';

    /**
     * Returns true if the given object is an instance of ConversionWorkspace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConversionWorkspace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConversionWorkspace.__pulumiType;
    }

    /**
     * Required. The ID of the conversion workspace to create.
     */
    public readonly conversionWorkspaceId!: pulumi.Output<string>;
    /**
     * The timestamp when the workspace resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The destination engine details.
     */
    public readonly destination!: pulumi.Output<outputs.datamigration.v1.DatabaseEngineInfoResponse>;
    /**
     * The display name for the workspace
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * A generic list of settings for the workspace. The settings are database pair dependant and can indicate default behavior for the mapping rules engine or turn on or off specific features. Such examples can be: convert_foreign_key_to_interleave=true, skip_triggers=false, ignore_non_table_synonyms=true
     */
    public readonly globalSettings!: pulumi.Output<{[key: string]: string}>;
    /**
     * Whether the workspace has uncommitted changes (changes which were made after the workspace was committed)
     */
    public /*out*/ readonly hasUncommittedChanges!: pulumi.Output<boolean>;
    /**
     * The latest commit id
     */
    public /*out*/ readonly latestCommitId!: pulumi.Output<string>;
    /**
     * The timestamp when the workspace was committed.
     */
    public /*out*/ readonly latestCommitTime!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Full name of the workspace resource, in the form of: projects/{project}/locations/{location}/conversionWorkspaces/{conversion_workspace}.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The source engine details.
     */
    public readonly source!: pulumi.Output<outputs.datamigration.v1.DatabaseEngineInfoResponse>;
    /**
     * The timestamp when the workspace resource was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ConversionWorkspace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConversionWorkspaceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.conversionWorkspaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'conversionWorkspaceId'");
            }
            if ((!args || args.destination === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destination'");
            }
            if ((!args || args.source === undefined) && !opts.urn) {
                throw new Error("Missing required property 'source'");
            }
            resourceInputs["conversionWorkspaceId"] = args ? args.conversionWorkspaceId : undefined;
            resourceInputs["destination"] = args ? args.destination : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["globalSettings"] = args ? args.globalSettings : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["hasUncommittedChanges"] = undefined /*out*/;
            resourceInputs["latestCommitId"] = undefined /*out*/;
            resourceInputs["latestCommitTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["conversionWorkspaceId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["destination"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["globalSettings"] = undefined /*out*/;
            resourceInputs["hasUncommittedChanges"] = undefined /*out*/;
            resourceInputs["latestCommitId"] = undefined /*out*/;
            resourceInputs["latestCommitTime"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["source"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["conversionWorkspaceId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ConversionWorkspace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConversionWorkspace resource.
 */
export interface ConversionWorkspaceArgs {
    /**
     * Required. The ID of the conversion workspace to create.
     */
    conversionWorkspaceId: pulumi.Input<string>;
    /**
     * The destination engine details.
     */
    destination: pulumi.Input<inputs.datamigration.v1.DatabaseEngineInfoArgs>;
    /**
     * The display name for the workspace
     */
    displayName?: pulumi.Input<string>;
    /**
     * A generic list of settings for the workspace. The settings are database pair dependant and can indicate default behavior for the mapping rules engine or turn on or off specific features. Such examples can be: convert_foreign_key_to_interleave=true, skip_triggers=false, ignore_non_table_synonyms=true
     */
    globalSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Full name of the workspace resource, in the form of: projects/{project}/locations/{location}/conversionWorkspaces/{conversion_workspace}.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    requestId?: pulumi.Input<string>;
    /**
     * The source engine details.
     */
    source: pulumi.Input<inputs.datamigration.v1.DatabaseEngineInfoArgs>;
}