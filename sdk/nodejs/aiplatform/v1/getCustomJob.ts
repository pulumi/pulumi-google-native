// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a CustomJob.
 */
export function getCustomJob(args: GetCustomJobArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:aiplatform/v1:getCustomJob", {
        "customJobId": args.customJobId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetCustomJobArgs {
    customJobId: string;
    location: string;
    project?: string;
}

export interface GetCustomJobResult {
    /**
     * Time when the CustomJob was created.
     */
    readonly createTime: string;
    /**
     * The display name of the CustomJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     */
    readonly displayName: string;
    /**
     * Customer-managed encryption key options for a CustomJob. If this is set, then all resources created by the CustomJob will be encrypted with the provided encryption key.
     */
    readonly encryptionSpec: outputs.aiplatform.v1.GoogleCloudAiplatformV1EncryptionSpecResponse;
    /**
     * Time when the CustomJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
     */
    readonly endTime: string;
    /**
     * Only populated when job's state is `JOB_STATE_FAILED` or `JOB_STATE_CANCELLED`.
     */
    readonly error: outputs.aiplatform.v1.GoogleRpcStatusResponse;
    /**
     * Job spec.
     */
    readonly jobSpec: outputs.aiplatform.v1.GoogleCloudAiplatformV1CustomJobSpecResponse;
    /**
     * The labels with user-defined metadata to organize CustomJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * Resource name of a CustomJob.
     */
    readonly name: string;
    /**
     * Time when the CustomJob for the first time entered the `JOB_STATE_RUNNING` state.
     */
    readonly startTime: string;
    /**
     * The detailed state of the job.
     */
    readonly state: string;
    /**
     * Time when the CustomJob was most recently updated.
     */
    readonly updateTime: string;
    /**
     * URIs for accessing [interactive shells](https://cloud.google.com/vertex-ai/docs/training/monitor-debug-interactive-shell) (one URI for each training node). Only available if job_spec.enable_web_access is `true`. The keys are names of each node in the training job; for example, `workerpool0-0` for the primary node, `workerpool1-0` for the first node in the second worker pool, and `workerpool1-1` for the second node in the second worker pool. The values are the URIs for each node's interactive shell.
     */
    readonly webAccessUris: {[key: string]: string};
}
/**
 * Gets a CustomJob.
 */
export function getCustomJobOutput(args: GetCustomJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomJobResult> {
    return pulumi.output(args).apply((a: any) => getCustomJob(a, opts))
}

export interface GetCustomJobOutputArgs {
    customJobId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}