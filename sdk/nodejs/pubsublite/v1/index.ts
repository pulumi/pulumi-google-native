// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getReservation";
export * from "./getSubscription";
export * from "./getTopic";
export * from "./reservation";
export * from "./subscription";
export * from "./topic";

// Export enums:
export * from "../../types/enums/pubsublite/v1";

// Import resources to register:
import { Reservation } from "./reservation";
import { Subscription } from "./subscription";
import { Topic } from "./topic";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:pubsublite/v1:Reservation":
                return new Reservation(name, <any>undefined, { urn })
            case "google-native:pubsublite/v1:Subscription":
                return new Subscription(name, <any>undefined, { urn })
            case "google-native:pubsublite/v1:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "pubsublite/v1", _module)