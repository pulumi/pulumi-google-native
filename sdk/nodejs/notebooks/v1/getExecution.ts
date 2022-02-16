// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets details of executions
 */
export function getExecution(args: GetExecutionArgs, opts?: pulumi.InvokeOptions): Promise<GetExecutionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:notebooks/v1:getExecution", {
        "executionId": args.executionId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetExecutionArgs {
    executionId: string;
    location: string;
    project?: string;
}

export interface GetExecutionResult {
    /**
     * Time the Execution was instantiated.
     */
    readonly createTime: string;
    /**
     * A brief description of this execution.
     */
    readonly description: string;
    /**
     * Name used for UI purposes. Name can only contain alphanumeric characters and underscores '_'.
     */
    readonly displayName: string;
    /**
     * execute metadata including name, hardware spec, region, labels, etc.
     */
    readonly executionTemplate: outputs.notebooks.v1.ExecutionTemplateResponse;
    /**
     * The URI of the external job used to execute the notebook.
     */
    readonly jobUri: string;
    /**
     * The resource name of the execute. Format: `projects/{project_id}/locations/{location}/executions/{execution_id}`
     */
    readonly name: string;
    /**
     * Output notebook file generated by this execution
     */
    readonly outputNotebookFile: string;
    /**
     * State of the underlying AI Platform job.
     */
    readonly state: string;
    /**
     * Time the Execution was last updated.
     */
    readonly updateTime: string;
}

export function getExecutionOutput(args: GetExecutionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExecutionResult> {
    return pulumi.output(args).apply(a => getExecution(a, opts))
}

export interface GetExecutionOutputArgs {
    executionId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}