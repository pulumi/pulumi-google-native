// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns a specified storage pool. Gets a list of available storage pools by making a list() request.
 */
export function getStoragePool(args: GetStoragePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetStoragePoolResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getStoragePool", {
        "project": args.project,
        "storagePool": args.storagePool,
        "zone": args.zone,
    }, opts);
}

export interface GetStoragePoolArgs {
    project?: string;
    storagePool: string;
    zone: string;
}

export interface GetStoragePoolResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#storagePool for storage pools.
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this storage pool, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a storage pool.
     */
    readonly labelFingerprint: string;
    /**
     * Labels to apply to this storage pool. These can be later modified by the setLabels method.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Provsioned IOPS of the storage pool.
     */
    readonly provisionedIops: string;
    /**
     * Status information for the storage pool resource.
     */
    readonly resourceStatus: outputs.compute.alpha.StoragePoolResourceStatusResponse;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource's resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * Size, in GiB, of the storage pool.
     */
    readonly sizeGb: string;
    /**
     * The status of storage pool creation. - CREATING: Storage pool is provisioning. storagePool. - FAILED: Storage pool creation failed. - READY: Storage pool is ready for use. - DELETING: Storage pool is deleting. 
     */
    readonly state: string;
    /**
     * Type of the storage pool
     */
    readonly type: string;
    /**
     * URL of the zone where the storage pool resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    readonly zone: string;
}
/**
 * Returns a specified storage pool. Gets a list of available storage pools by making a list() request.
 */
export function getStoragePoolOutput(args: GetStoragePoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStoragePoolResult> {
    return pulumi.output(args).apply((a: any) => getStoragePool(a, opts))
}

export interface GetStoragePoolOutputArgs {
    project?: pulumi.Input<string>;
    storagePool: pulumi.Input<string>;
    zone: pulumi.Input<string>;
}