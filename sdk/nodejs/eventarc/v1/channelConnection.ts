// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Create a new ChannelConnection in a particular project and location.
 */
export class ChannelConnection extends pulumi.CustomResource {
    /**
     * Get an existing ChannelConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ChannelConnection {
        return new ChannelConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:eventarc/v1:ChannelConnection';

    /**
     * Returns true if the given object is an instance of ChannelConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ChannelConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ChannelConnection.__pulumiType;
    }

    /**
     * Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
     */
    public readonly activationToken!: pulumi.Output<string>;
    /**
     * The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
     */
    public readonly channel!: pulumi.Output<string>;
    /**
     * Required. The user-provided ID to be assigned to the channel connection.
     */
    public readonly channelConnectionId!: pulumi.Output<string>;
    /**
     * The creation time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the connection.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Server assigned ID of the resource. The server guarantees uniqueness and immutability until deleted.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The last-modified time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ChannelConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ChannelConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.channel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channel'");
            }
            if ((!args || args.channelConnectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'channelConnectionId'");
            }
            resourceInputs["activationToken"] = args ? args.activationToken : undefined;
            resourceInputs["channel"] = args ? args.channel : undefined;
            resourceInputs["channelConnectionId"] = args ? args.channelConnectionId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["activationToken"] = undefined /*out*/;
            resourceInputs["channel"] = undefined /*out*/;
            resourceInputs["channelConnectionId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["channelConnectionId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ChannelConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ChannelConnection resource.
 */
export interface ChannelConnectionArgs {
    /**
     * Input only. Activation token for the channel. The token will be used during the creation of ChannelConnection to bind the channel with the provider project. This field will not be stored in the provider resource.
     */
    activationToken?: pulumi.Input<string>;
    /**
     * The name of the connected subscriber Channel. This is a weak reference to avoid cross project and cross accounts references. This must be in `projects/{project}/location/{location}/channels/{channel_id}` format.
     */
    channel: pulumi.Input<string>;
    /**
     * Required. The user-provided ID to be assigned to the channel connection.
     */
    channelConnectionId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The name of the connection.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}