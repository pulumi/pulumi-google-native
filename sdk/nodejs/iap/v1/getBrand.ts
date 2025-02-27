// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Retrieves the OAuth brand of the project.
 */
export function getBrand(args: GetBrandArgs, opts?: pulumi.InvokeOptions): Promise<GetBrandResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:iap/v1:getBrand", {
        "brandId": args.brandId,
        "project": args.project,
    }, opts);
}

export interface GetBrandArgs {
    brandId: string;
    project?: string;
}

export interface GetBrandResult {
    /**
     * Application name displayed on OAuth consent screen.
     */
    readonly applicationTitle: string;
    /**
     * Identifier of the brand. NOTE: GCP project number achieves the same brand identification purpose as only one brand per project can be created.
     */
    readonly name: string;
    /**
     * Whether the brand is only intended for usage inside the G Suite organization only.
     */
    readonly orgInternalOnly: boolean;
    /**
     * Support email displayed on the OAuth consent screen.
     */
    readonly supportEmail: string;
}
/**
 * Retrieves the OAuth brand of the project.
 */
export function getBrandOutput(args: GetBrandOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBrandResult> {
    return pulumi.output(args).apply((a: any) => getBrand(a, opts))
}

export interface GetBrandOutputArgs {
    brandId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
