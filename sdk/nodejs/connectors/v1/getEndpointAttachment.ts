// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single EndpointAttachment.
 */
export function getEndpointAttachment(args: GetEndpointAttachmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointAttachmentResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:connectors/v1:getEndpointAttachment", {
        "endpointAttachmentId": args.endpointAttachmentId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEndpointAttachmentArgs {
    endpointAttachmentId: string;
    location: string;
    project?: string;
}

export interface GetEndpointAttachmentResult {
    /**
     * Created time.
     */
    readonly createTime: string;
    /**
     * Optional. Description of the resource.
     */
    readonly description: string;
    /**
     * The Private Service Connect connection endpoint ip
     */
    readonly endpointIp: string;
    /**
     * Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
     */
    readonly labels: {[key: string]: string};
    /**
     * Resource name of the Endpoint Attachment. Format: projects/{project}/locations/{location}/endpointAttachments/{endpoint_attachment}
     */
    readonly name: string;
    /**
     * The path of the service attachment
     */
    readonly serviceAttachment: string;
    /**
     * Updated time.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single EndpointAttachment.
 */
export function getEndpointAttachmentOutput(args: GetEndpointAttachmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointAttachmentResult> {
    return pulumi.output(args).apply((a: any) => getEndpointAttachment(a, opts))
}

export interface GetEndpointAttachmentOutputArgs {
    endpointAttachmentId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}