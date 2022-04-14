// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets an alias.
 */
export function getAlias(args: GetAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetAliasResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:apigee/v1:getAlias", {
        "aliasId": args.aliasId,
        "environmentId": args.environmentId,
        "keystoreId": args.keystoreId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetAliasArgs {
    aliasId: string;
    environmentId: string;
    keystoreId: string;
    organizationId: string;
}

export interface GetAliasResult {
    /**
     * Resource ID for this alias. Values must match the regular expression `[^/]{1,255}`.
     */
    readonly alias: string;
    /**
     * Chain of certificates under this alias.
     */
    readonly certsInfo: outputs.apigee.v1.GoogleCloudApigeeV1CertificateResponse;
    /**
     * Type of alias.
     */
    readonly type: string;
}

export function getAliasOutput(args: GetAliasOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAliasResult> {
    return pulumi.output(args).apply(a => getAlias(a, opts))
}

export interface GetAliasOutputArgs {
    aliasId: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    keystoreId: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}