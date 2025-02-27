// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns an IAM policy for the specified object.
 */
export function getObjectIamPolicy(args: GetObjectIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetObjectIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:storage/v1:getObjectIamPolicy", {
        "bucket": args.bucket,
        "generation": args.generation,
        "object": args.object,
        "userProject": args.userProject,
    }, opts);
}

export interface GetObjectIamPolicyArgs {
    bucket: string;
    generation?: string;
    object: string;
    userProject?: string;
}

export interface GetObjectIamPolicyResult {
    /**
     * An association between a role, which comes with a set of permissions, and members who may assume that role.
     */
    readonly bindings: outputs.storage.v1.ObjectIamPolicyBindingsItemResponse[];
    /**
     * HTTP 1.1  Entity tag for the policy.
     */
    readonly etag: string;
    /**
     * The kind of item this is. For policies, this is always storage#policy. This field is ignored on input.
     */
    readonly kind: string;
    /**
     * The ID of the resource to which this policy belongs. Will be of the form projects/_/buckets/bucket for buckets, projects/_/buckets/bucket/objects/object for objects, and projects/_/buckets/bucket/managedFolders/managedFolder. A specific generation may be specified by appending #generationNumber to the end of the object name, e.g. projects/_/buckets/my-bucket/objects/data.txt#17. The current generation can be denoted with #0. This field is ignored on input.
     */
    readonly resourceId: string;
    /**
     * The IAM policy format version.
     */
    readonly version: number;
}
/**
 * Returns an IAM policy for the specified object.
 */
export function getObjectIamPolicyOutput(args: GetObjectIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetObjectIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getObjectIamPolicy(a, opts))
}

export interface GetObjectIamPolicyOutputArgs {
    bucket: pulumi.Input<string>;
    generation?: pulumi.Input<string>;
    object: pulumi.Input<string>;
    userProject?: pulumi.Input<string>;
}
