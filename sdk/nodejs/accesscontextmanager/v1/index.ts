// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./accessLevel";
export * from "./accessPolicy";
export * from "./accessPolicyIamPolicy";
export * from "./gcpUserAccessBinding";
export * from "./getAccessLevel";
export * from "./getAccessPolicy";
export * from "./getAccessPolicyIamPolicy";
export * from "./getGcpUserAccessBinding";
export * from "./getServicePerimeter";
export * from "./servicePerimeter";

// Export enums:
export * from "../../types/enums/accesscontextmanager/v1";

// Import resources to register:
import { AccessLevel } from "./accessLevel";
import { AccessPolicy } from "./accessPolicy";
import { AccessPolicyIamPolicy } from "./accessPolicyIamPolicy";
import { GcpUserAccessBinding } from "./gcpUserAccessBinding";
import { ServicePerimeter } from "./servicePerimeter";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:accesscontextmanager/v1:AccessLevel":
                return new AccessLevel(name, <any>undefined, { urn })
            case "google-native:accesscontextmanager/v1:AccessPolicy":
                return new AccessPolicy(name, <any>undefined, { urn })
            case "google-native:accesscontextmanager/v1:AccessPolicyIamPolicy":
                return new AccessPolicyIamPolicy(name, <any>undefined, { urn })
            case "google-native:accesscontextmanager/v1:GcpUserAccessBinding":
                return new GcpUserAccessBinding(name, <any>undefined, { urn })
            case "google-native:accesscontextmanager/v1:ServicePerimeter":
                return new ServicePerimeter(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "accesscontextmanager/v1", _module)