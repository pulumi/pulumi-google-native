// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single ServiceClass.
 */
export function getServiceClass(args: GetServiceClassArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceClassResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:networkconnectivity/v1:getServiceClass", {
        "location": args.location,
        "project": args.project,
        "serviceClassId": args.serviceClassId,
    }, opts);
}

export interface GetServiceClassArgs {
    location: string;
    project?: string;
    serviceClassId: string;
}

export interface GetServiceClassResult {
    /**
     * Time when the ServiceClass was created.
     */
    readonly createTime: string;
    /**
     * A description of this resource.
     */
    readonly description: string;
    /**
     * User-defined labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * Immutable. The name of a ServiceClass resource. Format: projects/{project}/locations/{location}/serviceClasses/{service_class} See: https://google.aip.dev/122#fields-representing-resource-names
     */
    readonly name: string;
    /**
     * The generated service class name. Use this name to refer to the Service class in Service Connection Maps and Service Connection Policies.
     */
    readonly serviceClass: string;
    /**
     * URIs of all Service Connection Maps using this service class.
     */
    readonly serviceConnectionMaps: string[];
    /**
     * Time when the ServiceClass was updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single ServiceClass.
 */
export function getServiceClassOutput(args: GetServiceClassOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceClassResult> {
    return pulumi.output(args).apply((a: any) => getServiceClass(a, opts))
}

export interface GetServiceClassOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    serviceClassId: pulumi.Input<string>;
}