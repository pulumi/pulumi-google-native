// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new EventSubscription in a given project,location and connection.
 */
export class EventSubscription extends pulumi.CustomResource {
    /**
     * Get an existing EventSubscription resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EventSubscription {
        return new EventSubscription(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:connectors/v1:EventSubscription';

    /**
     * Returns true if the given object is an instance of EventSubscription.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSubscription {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSubscription.__pulumiType;
    }

    public readonly connectionId!: pulumi.Output<string>;
    /**
     * Created time.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. The destination to hit when we receive an event
     */
    public readonly destinations!: pulumi.Output<outputs.connectors.v1.EventSubscriptionDestinationResponse>;
    /**
     * Required. Identifier to assign to the Event Subscription. Must be unique within scope of the parent resource.
     */
    public readonly eventSubscriptionId!: pulumi.Output<string>;
    /**
     * Optional. Event type id of the event of current EventSubscription.
     */
    public readonly eventTypeId!: pulumi.Output<string>;
    /**
     * Optional. JMS is the source for the event listener.
     */
    public readonly jms!: pulumi.Output<outputs.connectors.v1.JMSResponse>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name of the EventSubscription. Format: projects/{project}/locations/{location}/connections/{connection}/eventSubscriptions/{event_subscription}
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. Status indicates the status of the event subscription resource
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.connectors.v1.EventSubscriptionStatusResponse>;
    /**
     * Optional. name of the Subscriber for the current EventSubscription.
     */
    public readonly subscriber!: pulumi.Output<string>;
    /**
     * Optional. Link for Subscriber of the current EventSubscription.
     */
    public readonly subscriberLink!: pulumi.Output<string>;
    /**
     * Updated time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a EventSubscription resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSubscriptionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionId'");
            }
            if ((!args || args.eventSubscriptionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'eventSubscriptionId'");
            }
            resourceInputs["connectionId"] = args ? args.connectionId : undefined;
            resourceInputs["destinations"] = args ? args.destinations : undefined;
            resourceInputs["eventSubscriptionId"] = args ? args.eventSubscriptionId : undefined;
            resourceInputs["eventTypeId"] = args ? args.eventTypeId : undefined;
            resourceInputs["jms"] = args ? args.jms : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["subscriber"] = args ? args.subscriber : undefined;
            resourceInputs["subscriberLink"] = args ? args.subscriberLink : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["connectionId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["destinations"] = undefined /*out*/;
            resourceInputs["eventSubscriptionId"] = undefined /*out*/;
            resourceInputs["eventTypeId"] = undefined /*out*/;
            resourceInputs["jms"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["subscriber"] = undefined /*out*/;
            resourceInputs["subscriberLink"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["connectionId", "eventSubscriptionId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EventSubscription.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EventSubscription resource.
 */
export interface EventSubscriptionArgs {
    connectionId: pulumi.Input<string>;
    /**
     * Optional. The destination to hit when we receive an event
     */
    destinations?: pulumi.Input<inputs.connectors.v1.EventSubscriptionDestinationArgs>;
    /**
     * Required. Identifier to assign to the Event Subscription. Must be unique within scope of the parent resource.
     */
    eventSubscriptionId: pulumi.Input<string>;
    /**
     * Optional. Event type id of the event of current EventSubscription.
     */
    eventTypeId?: pulumi.Input<string>;
    /**
     * Optional. JMS is the source for the event listener.
     */
    jms?: pulumi.Input<inputs.connectors.v1.JMSArgs>;
    location?: pulumi.Input<string>;
    /**
     * Resource name of the EventSubscription. Format: projects/{project}/locations/{location}/connections/{connection}/eventSubscriptions/{event_subscription}
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. name of the Subscriber for the current EventSubscription.
     */
    subscriber?: pulumi.Input<string>;
    /**
     * Optional. Link for Subscriber of the current EventSubscription.
     */
    subscriberLink?: pulumi.Input<string>;
}
