// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a domain mapping on the specified site.
 */
export function getDomain(args: GetDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:firebasehosting/v1beta1:getDomain", {
        "domainId": args.domainId,
        "project": args.project,
        "siteId": args.siteId,
    }, opts);
}

export interface GetDomainArgs {
    domainId: string;
    project?: string;
    siteId: string;
}

export interface GetDomainResult {
    /**
     * The domain name of the association.
     */
    readonly domainName: string;
    /**
     * If set, the domain should redirect with the provided parameters.
     */
    readonly domainRedirect: outputs.firebasehosting.v1beta1.DomainRedirectResponse;
    /**
     * Information about the provisioning of certificates and the health of the DNS resolution for the domain.
     */
    readonly provisioning: outputs.firebasehosting.v1beta1.DomainProvisioningResponse;
    /**
     * The site name of the association.
     */
    readonly site: string;
    /**
     * Additional status of the domain association.
     */
    readonly status: string;
    /**
     * The time at which the domain was last updated.
     */
    readonly updateTime: string;
}

export function getDomainOutput(args: GetDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainResult> {
    return pulumi.output(args).apply(a => getDomain(a, opts))
}

export interface GetDomainOutputArgs {
    domainId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    siteId: pulumi.Input<string>;
}