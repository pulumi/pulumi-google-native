// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified firewall policy.
 */
export function getFirewallpolicy(args: GetFirewallpolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallpolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:recaptchaenterprise/v1:getFirewallpolicy", {
        "firewallpolicyId": args.firewallpolicyId,
        "project": args.project,
    }, opts);
}

export interface GetFirewallpolicyArgs {
    firewallpolicyId: string;
    project?: string;
}

export interface GetFirewallpolicyResult {
    /**
     * The actions that the caller should take regarding user access. There should be at most one terminal action. A terminal action is any action that forces a response, such as AllowAction, BlockAction or SubstituteAction. Zero or more non-terminal actions such as SetHeader might be specified. A single policy can contain up to 16 actions.
     */
    readonly actions: outputs.recaptchaenterprise.v1.GoogleCloudRecaptchaenterpriseV1FirewallActionResponse[];
    /**
     * A CEL (Common Expression Language) conditional expression that specifies if this policy applies to an incoming user request. If this condition evaluates to true and the requested path matched the path pattern, the associated actions should be executed by the caller. The condition string is checked for CEL syntax correctness on creation. For more information, see the [CEL spec](https://github.com/google/cel-spec) and its [language definition](https://github.com/google/cel-spec/blob/master/doc/langdef.md). A condition has a max length of 500 characters.
     */
    readonly condition: string;
    /**
     * A description of what this policy aims to achieve, for convenience purposes. The description can at most include 256 UTF-8 characters.
     */
    readonly description: string;
    /**
     * The resource name for the FirewallPolicy in the format "projects/{project}/firewallpolicies/{firewallpolicy}".
     */
    readonly name: string;
    /**
     * The path for which this policy applies, specified as a glob pattern. For more information on glob, see the [manual page](https://man7.org/linux/man-pages/man7/glob.7.html). A path has a max length of 200 characters.
     */
    readonly path: string;
}
/**
 * Returns the specified firewall policy.
 */
export function getFirewallpolicyOutput(args: GetFirewallpolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFirewallpolicyResult> {
    return pulumi.output(args).apply((a: any) => getFirewallpolicy(a, opts))
}

export interface GetFirewallpolicyOutputArgs {
    firewallpolicyId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}