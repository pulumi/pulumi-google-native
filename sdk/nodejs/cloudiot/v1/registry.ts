// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a device registry that contains devices.
 */
export class Registry extends pulumi.CustomResource {
    /**
     * Get an existing Registry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Registry {
        return new Registry(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudiot/v1:Registry';

    /**
     * Returns true if the given object is an instance of Registry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registry.__pulumiType;
    }

    /**
     * The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
     */
    public readonly credentials!: pulumi.Output<outputs.cloudiot.v1.RegistryCredentialResponse[]>;
    /**
     * The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
     */
    public readonly eventNotificationConfigs!: pulumi.Output<outputs.cloudiot.v1.EventNotificationConfigResponse[]>;
    /**
     * The DeviceService (HTTP) configuration for this device registry.
     */
    public readonly httpConfig!: pulumi.Output<outputs.cloudiot.v1.HttpConfigResponse>;
    /**
     * **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
     */
    public readonly logLevel!: pulumi.Output<string>;
    /**
     * The MQTT configuration for this device registry.
     */
    public readonly mqttConfig!: pulumi.Output<outputs.cloudiot.v1.MqttConfigResponse>;
    /**
     * The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
     */
    public readonly stateNotificationConfig!: pulumi.Output<outputs.cloudiot.v1.StateNotificationConfigResponse>;

    /**
     * Create a Registry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegistryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["eventNotificationConfigs"] = args ? args.eventNotificationConfigs : undefined;
            resourceInputs["httpConfig"] = args ? args.httpConfig : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["logLevel"] = args ? args.logLevel : undefined;
            resourceInputs["mqttConfig"] = args ? args.mqttConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["stateNotificationConfig"] = args ? args.stateNotificationConfig : undefined;
        } else {
            resourceInputs["credentials"] = undefined /*out*/;
            resourceInputs["eventNotificationConfigs"] = undefined /*out*/;
            resourceInputs["httpConfig"] = undefined /*out*/;
            resourceInputs["logLevel"] = undefined /*out*/;
            resourceInputs["mqttConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["stateNotificationConfig"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Registry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Registry resource.
 */
export interface RegistryArgs {
    /**
     * The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.cloudiot.v1.RegistryCredentialArgs>[]>;
    /**
     * The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
     */
    eventNotificationConfigs?: pulumi.Input<pulumi.Input<inputs.cloudiot.v1.EventNotificationConfigArgs>[]>;
    /**
     * The DeviceService (HTTP) configuration for this device registry.
     */
    httpConfig?: pulumi.Input<inputs.cloudiot.v1.HttpConfigArgs>;
    /**
     * The identifier of this device registry. For example, `myRegistry`.
     */
    id?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
     */
    logLevel?: pulumi.Input<enums.cloudiot.v1.RegistryLogLevel>;
    /**
     * The MQTT configuration for this device registry.
     */
    mqttConfig?: pulumi.Input<inputs.cloudiot.v1.MqttConfigArgs>;
    /**
     * The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
     */
    stateNotificationConfig?: pulumi.Input<inputs.cloudiot.v1.StateNotificationConfigArgs>;
}