// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { ConfigArgs } from "./config";
export type Config = import("./config").Config;
export const Config: typeof import("./config").Config = null as any;
utilities.lazyLoad(exports, ["Config"], () => require("./config"));

export { ConfigIamBindingArgs } from "./configIamBinding";
export type ConfigIamBinding = import("./configIamBinding").ConfigIamBinding;
export const ConfigIamBinding: typeof import("./configIamBinding").ConfigIamBinding = null as any;
utilities.lazyLoad(exports, ["ConfigIamBinding"], () => require("./configIamBinding"));

export { ConfigIamMemberArgs } from "./configIamMember";
export type ConfigIamMember = import("./configIamMember").ConfigIamMember;
export const ConfigIamMember: typeof import("./configIamMember").ConfigIamMember = null as any;
utilities.lazyLoad(exports, ["ConfigIamMember"], () => require("./configIamMember"));

export { ConfigIamPolicyArgs } from "./configIamPolicy";
export type ConfigIamPolicy = import("./configIamPolicy").ConfigIamPolicy;
export const ConfigIamPolicy: typeof import("./configIamPolicy").ConfigIamPolicy = null as any;
utilities.lazyLoad(exports, ["ConfigIamPolicy"], () => require("./configIamPolicy"));

export { GetConfigArgs, GetConfigResult, GetConfigOutputArgs } from "./getConfig";
export const getConfig: typeof import("./getConfig").getConfig = null as any;
export const getConfigOutput: typeof import("./getConfig").getConfigOutput = null as any;
utilities.lazyLoad(exports, ["getConfig","getConfigOutput"], () => require("./getConfig"));

export { GetConfigIamPolicyArgs, GetConfigIamPolicyResult, GetConfigIamPolicyOutputArgs } from "./getConfigIamPolicy";
export const getConfigIamPolicy: typeof import("./getConfigIamPolicy").getConfigIamPolicy = null as any;
export const getConfigIamPolicyOutput: typeof import("./getConfigIamPolicy").getConfigIamPolicyOutput = null as any;
utilities.lazyLoad(exports, ["getConfigIamPolicy","getConfigIamPolicyOutput"], () => require("./getConfigIamPolicy"));

export { GetVariableArgs, GetVariableResult, GetVariableOutputArgs } from "./getVariable";
export const getVariable: typeof import("./getVariable").getVariable = null as any;
export const getVariableOutput: typeof import("./getVariable").getVariableOutput = null as any;
utilities.lazyLoad(exports, ["getVariable","getVariableOutput"], () => require("./getVariable"));

export { GetWaiterArgs, GetWaiterResult, GetWaiterOutputArgs } from "./getWaiter";
export const getWaiter: typeof import("./getWaiter").getWaiter = null as any;
export const getWaiterOutput: typeof import("./getWaiter").getWaiterOutput = null as any;
utilities.lazyLoad(exports, ["getWaiter","getWaiterOutput"], () => require("./getWaiter"));

export { VariableArgs } from "./variable";
export type Variable = import("./variable").Variable;
export const Variable: typeof import("./variable").Variable = null as any;
utilities.lazyLoad(exports, ["Variable"], () => require("./variable"));

export { WaiterArgs } from "./waiter";
export type Waiter = import("./waiter").Waiter;
export const Waiter: typeof import("./waiter").Waiter = null as any;
utilities.lazyLoad(exports, ["Waiter"], () => require("./waiter"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:runtimeconfig/v1beta1:Config":
                return new Config(name, <any>undefined, { urn })
            case "google-native:runtimeconfig/v1beta1:ConfigIamBinding":
                return new ConfigIamBinding(name, <any>undefined, { urn })
            case "google-native:runtimeconfig/v1beta1:ConfigIamMember":
                return new ConfigIamMember(name, <any>undefined, { urn })
            case "google-native:runtimeconfig/v1beta1:ConfigIamPolicy":
                return new ConfigIamPolicy(name, <any>undefined, { urn })
            case "google-native:runtimeconfig/v1beta1:Variable":
                return new Variable(name, <any>undefined, { urn })
            case "google-native:runtimeconfig/v1beta1:Waiter":
                return new Waiter(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "runtimeconfig/v1beta1", _module)
