// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new preference set in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class PreferenceSet extends pulumi.CustomResource {
    /**
     * Get an existing PreferenceSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PreferenceSet {
        return new PreferenceSet(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:migrationcenter/v1:PreferenceSet';

    /**
     * Returns true if the given object is an instance of PreferenceSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PreferenceSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PreferenceSet.__pulumiType;
    }

    /**
     * The timestamp when the preference set was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of the preference set.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User-friendly display name. Maximum length is 63 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the preference set.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
     */
    public readonly preferenceSetId!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The timestamp when the preference set was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * A set of preferences that applies to all virtual machines in the context.
     */
    public readonly virtualMachinePreferences!: pulumi.Output<outputs.migrationcenter.v1.VirtualMachinePreferencesResponse>;

    /**
     * Create a PreferenceSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PreferenceSetArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.preferenceSetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'preferenceSetId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["preferenceSetId"] = args ? args.preferenceSetId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["virtualMachinePreferences"] = args ? args.virtualMachinePreferences : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["preferenceSetId"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["virtualMachinePreferences"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "preferenceSetId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(PreferenceSet.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a PreferenceSet resource.
 */
export interface PreferenceSetArgs {
    /**
     * A description of the preference set.
     */
    description?: pulumi.Input<string>;
    /**
     * User-friendly display name. Maximum length is 63 characters.
     */
    displayName?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Required. User specified ID for the preference set. It will become the last component of the preference set name. The ID must be unique within the project, must conform with RFC-1034, is restricted to lower-cased letters, and has a maximum length of 63 characters. The ID must match the regular expression `[a-z]([a-z0-9-]{0,61}[a-z0-9])?`.
     */
    preferenceSetId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * A set of preferences that applies to all virtual machines in the context.
     */
    virtualMachinePreferences?: pulumi.Input<inputs.migrationcenter.v1.VirtualMachinePreferencesArgs>;
}
