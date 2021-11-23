// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a custom class.
 */
export class CustomClass extends pulumi.CustomResource {
    /**
     * Get an existing CustomClass resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomClass {
        return new CustomClass(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:speech/v1:CustomClass';

    /**
     * Returns true if the given object is an instance of CustomClass.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomClass {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomClass.__pulumiType;
    }

    /**
     * If this custom class is a resource, the custom_class_id is the resource id of the CustomClass. Case sensitive.
     */
    public readonly customClassId!: pulumi.Output<string>;
    /**
     * A collection of class items.
     */
    public readonly items!: pulumi.Output<outputs.speech.v1.ClassItemResponse[]>;
    /**
     * The resource name of the custom class.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a CustomClass resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomClassArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.customClassId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'customClassId'");
            }
            inputs["customClassId"] = args ? args.customClassId : undefined;
            inputs["items"] = args ? args.items : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
        } else {
            inputs["customClassId"] = undefined /*out*/;
            inputs["items"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(CustomClass.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomClass resource.
 */
export interface CustomClassArgs {
    /**
     * The ID to use for the custom class, which will become the final component of the custom class' resource name. This value should be 4-63 characters, and valid characters are /a-z-/.
     */
    customClassId: pulumi.Input<string>;
    /**
     * A collection of class items.
     */
    items?: pulumi.Input<pulumi.Input<inputs.speech.v1.ClassItemArgs>[]>;
    location?: pulumi.Input<string>;
    /**
     * The resource name of the custom class.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}