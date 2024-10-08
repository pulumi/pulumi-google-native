// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates an AppGroup. Once created, user can register apps under the AppGroup to obtain secret key and password. At creation time, the AppGroup's state is set as `active`.
 */
export class Appgroup extends pulumi.CustomResource {
    /**
     * Get an existing Appgroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Appgroup {
        return new Appgroup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Appgroup';

    /**
     * Returns true if the given object is an instance of Appgroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Appgroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Appgroup.__pulumiType;
    }

    /**
     * Internal identifier that cannot be edited
     */
    public /*out*/ readonly appGroupId!: pulumi.Output<string>;
    /**
     * A list of attributes
     */
    public readonly attributes!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1AttributeResponse[]>;
    /**
     * channel identifier identifies the owner maintaing this grouping.
     */
    public readonly channelId!: pulumi.Output<string>;
    /**
     * A reference to the associated storefront/marketplace.
     */
    public readonly channelUri!: pulumi.Output<string>;
    /**
     * Created time as milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * app group name displayed in the UI
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Modified time as milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Immutable. Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._\-$ %.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Immutable. the org the app group is created
     */
    public readonly organization!: pulumi.Output<string>;
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * Valid values are `active` or `inactive`. Note that the status of the AppGroup should be updated via UpdateAppGroupRequest by setting the action as `active` or `inactive`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Appgroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AppgroupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["channelId"] = args ? args.channelId : undefined;
            resourceInputs["channelUri"] = args ? args.channelUri : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["appGroupId"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["appGroupId"] = undefined /*out*/;
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["channelId"] = undefined /*out*/;
            resourceInputs["channelUri"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organization"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["organizationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Appgroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Appgroup resource.
 */
export interface AppgroupArgs {
    /**
     * A list of attributes
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1AttributeArgs>[]>;
    /**
     * channel identifier identifies the owner maintaing this grouping.
     */
    channelId?: pulumi.Input<string>;
    /**
     * A reference to the associated storefront/marketplace.
     */
    channelUri?: pulumi.Input<string>;
    /**
     * app group name displayed in the UI
     */
    displayName?: pulumi.Input<string>;
    /**
     * Immutable. Name of the AppGroup. Characters you can use in the name are restricted to: A-Z0-9._\-$ %.
     */
    name?: pulumi.Input<string>;
    /**
     * Immutable. the org the app group is created
     */
    organization?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}
