// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Retrieves an SSH public key.
 */
export function getSshPublicKey(args: GetSshPublicKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetSshPublicKeyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:oslogin/v1beta:getSshPublicKey", {
        "sshPublicKeyId": args.sshPublicKeyId,
        "userId": args.userId,
    }, opts);
}

export interface GetSshPublicKeyArgs {
    sshPublicKeyId: string;
    userId: string;
}

export interface GetSshPublicKeyResult {
    /**
     * An expiration time in microseconds since epoch.
     */
    readonly expirationTimeUsec: string;
    /**
     * The SHA-256 fingerprint of the SSH public key.
     */
    readonly fingerprint: string;
    /**
     * Public key text in SSH format, defined by RFC4253 section 6.6.
     */
    readonly key: string;
    /**
     * The canonical resource name.
     */
    readonly name: string;
}

export function getSshPublicKeyOutput(args: GetSshPublicKeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSshPublicKeyResult> {
    return pulumi.output(args).apply(a => getSshPublicKey(a, opts))
}

export interface GetSshPublicKeyOutputArgs {
    sshPublicKeyId: pulumi.Input<string>;
    userId: pulumi.Input<string>;
}