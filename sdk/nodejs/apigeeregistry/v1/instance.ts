// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Provisions instance resources for the Registry.
 * Auto-naming is currently not supported for this resource.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigeeregistry/v1:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Config of the Instance.
     */
    public readonly config!: pulumi.Output<outputs.apigeeregistry.v1.ConfigResponse>;
    /**
     * Creation timestamp.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Required. Identifier to assign to the Instance. Must be unique within scope of the parent resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Format: `projects/*&#47;locations/*&#47;instance`. Currently only locations/global is supported.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The current state of the Instance.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Extra information of Instance.State if the state is `FAILED`.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * Last update timestamp.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["config"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Config of the Instance.
     */
    config: pulumi.Input<inputs.apigeeregistry.v1.ConfigArgs>;
    /**
     * Required. Identifier to assign to the Instance. Must be unique within scope of the parent resource.
     */
    instanceId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Format: `projects/*&#47;locations/*&#47;instance`. Currently only locations/global is supported.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}