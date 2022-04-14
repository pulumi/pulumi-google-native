// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets a managed service. Authentication is required unless the service is public.
 */
export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:servicemanagement/v1:getService", {
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetServiceArgs {
    serviceName: string;
}

export interface GetServiceResult {
    /**
     * ID of the project that produces and owns this service.
     */
    readonly producerProjectId: string;
    /**
     * The name of the service. See the [overview](https://cloud.google.com/service-management/overview) for naming requirements.
     */
    readonly serviceName: string;
}

export function getServiceOutput(args: GetServiceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceResult> {
    return pulumi.output(args).apply(a => getService(a, opts))
}

export interface GetServiceOutputArgs {
    serviceName: pulumi.Input<string>;
}