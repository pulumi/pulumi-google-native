// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an access level. The long-running operation from this RPC has a successful status after the access level propagates to long-lasting storage. If access levels contain errors, an error response is returned for the first error encountered.
 */
export class AccessLevel extends pulumi.CustomResource {
    /**
     * Get an existing AccessLevel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AccessLevel {
        return new AccessLevel(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:accesscontextmanager/v1:AccessLevel';

    /**
     * Returns true if the given object is an instance of AccessLevel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccessLevel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccessLevel.__pulumiType;
    }

    /**
     * A `BasicLevel` composed of `Conditions`.
     */
    public readonly basic!: pulumi.Output<outputs.accesscontextmanager.v1.BasicLevelResponse>;
    /**
     * A `CustomLevel` written in the Common Expression Language.
     */
    public readonly custom!: pulumi.Output<outputs.accesscontextmanager.v1.CustomLevelResponse>;
    /**
     * Description of the `AccessLevel` and its use. Does not affect behavior.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Resource name for the Access Level. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{access_policy}/accessLevels/{access_level}`. The maximum length of the `access_level` component is 50 characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a AccessLevel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccessLevelArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.accessPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'accessPolicyId'");
            }
            resourceInputs["accessPolicyId"] = args ? args.accessPolicyId : undefined;
            resourceInputs["basic"] = args ? args.basic : undefined;
            resourceInputs["custom"] = args ? args.custom : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
        } else {
            resourceInputs["basic"] = undefined /*out*/;
            resourceInputs["custom"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["title"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccessLevel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a AccessLevel resource.
 */
export interface AccessLevelArgs {
    accessPolicyId: pulumi.Input<string>;
    /**
     * A `BasicLevel` composed of `Conditions`.
     */
    basic?: pulumi.Input<inputs.accesscontextmanager.v1.BasicLevelArgs>;
    /**
     * A `CustomLevel` written in the Common Expression Language.
     */
    custom?: pulumi.Input<inputs.accesscontextmanager.v1.CustomLevelArgs>;
    /**
     * Description of the `AccessLevel` and its use. Does not affect behavior.
     */
    description?: pulumi.Input<string>;
    /**
     * Resource name for the Access Level. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{access_policy}/accessLevels/{access_level}`. The maximum length of the `access_level` component is 50 characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    title?: pulumi.Input<string>;
}