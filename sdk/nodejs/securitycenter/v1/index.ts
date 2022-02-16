// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getMuteConfig";
export * from "./getNotificationConfig";
export * from "./getOrganizationMuteConfig";
export * from "./getOrganizationSourceIamPolicy";
export * from "./getSource";
export * from "./muteConfig";
export * from "./notificationConfig";
export * from "./organizationMuteConfig";
export * from "./organizationSourceIamPolicy";
export * from "./source";

// Export enums:
export * from "../../types/enums/securitycenter/v1";

// Import resources to register:
import { MuteConfig } from "./muteConfig";
import { NotificationConfig } from "./notificationConfig";
import { OrganizationMuteConfig } from "./organizationMuteConfig";
import { OrganizationSourceIamPolicy } from "./organizationSourceIamPolicy";
import { Source } from "./source";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:securitycenter/v1:MuteConfig":
                return new MuteConfig(name, <any>undefined, { urn })
            case "google-native:securitycenter/v1:NotificationConfig":
                return new NotificationConfig(name, <any>undefined, { urn })
            case "google-native:securitycenter/v1:OrganizationMuteConfig":
                return new OrganizationMuteConfig(name, <any>undefined, { urn })
            case "google-native:securitycenter/v1:OrganizationSourceIamPolicy":
                return new OrganizationSourceIamPolicy(name, <any>undefined, { urn })
            case "google-native:securitycenter/v1:Source":
                return new Source(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "securitycenter/v1", _module)