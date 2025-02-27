// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets information about a Job.
 */
export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:run/v2:getJob", {
        "jobId": args.jobId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetJobArgs {
    jobId: string;
    location: string;
    project?: string;
}

export interface GetJobResult {
    /**
     * Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected on new resources. All system annotations in v1 now have a corresponding field in v2 Job. This field follows Kubernetes annotations' namespacing, limits, and rules.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Settings for the Binary Authorization feature.
     */
    readonly binaryAuthorization: outputs.run.v2.GoogleCloudRunV2BinaryAuthorizationResponse;
    /**
     * Arbitrary identifier for the API client.
     */
    readonly client: string;
    /**
     * Arbitrary version identifier for the API client.
     */
    readonly clientVersion: string;
    /**
     * The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
     */
    readonly conditions: outputs.run.v2.GoogleCloudRunV2ConditionResponse[];
    /**
     * The creation time.
     */
    readonly createTime: string;
    /**
     * Email address of the authenticated creator.
     */
    readonly creator: string;
    /**
     * The deletion time.
     */
    readonly deleteTime: string;
    /**
     * A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
     */
    readonly etag: string;
    /**
     * Number of executions created for this job.
     */
    readonly executionCount: number;
    /**
     * For a deleted resource, the time after which it will be permamently deleted.
     */
    readonly expireTime: string;
    /**
     * A number that monotonically increases every time the user modifies the desired state.
     */
    readonly generation: string;
    /**
     * Unstructured key value map that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels. Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Job.
     */
    readonly labels: {[key: string]: string};
    /**
     * Email address of the last authenticated modifier.
     */
    readonly lastModifier: string;
    /**
     * Name of the last created execution.
     */
    readonly latestCreatedExecution: outputs.run.v2.GoogleCloudRunV2ExecutionReferenceResponse;
    /**
     * The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
     */
    readonly launchStage: string;
    /**
     * The fully qualified name of this Job. Format: projects/{project}/locations/{location}/jobs/{job}
     */
    readonly name: string;
    /**
     * The generation of this Job. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
     */
    readonly observedGeneration: string;
    /**
     * Returns true if the Job is currently being acted upon by the system to bring it into the desired state. When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, `observed_generation` and `latest_succeeded_execution`, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `observed_generation` and `generation`, `latest_succeeded_execution` and `latest_created_execution`. If reconciliation failed, `observed_generation` and `latest_succeeded_execution` will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in `terminal_condition` and `conditions`.
     */
    readonly reconciling: boolean;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * The template used to create executions for this Job.
     */
    readonly template: outputs.run.v2.GoogleCloudRunV2ExecutionTemplateResponse;
    /**
     * The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state.
     */
    readonly terminalCondition: outputs.run.v2.GoogleCloudRunV2ConditionResponse;
    /**
     * Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
     */
    readonly uid: string;
    /**
     * The last-modified time.
     */
    readonly updateTime: string;
}
/**
 * Gets information about a Job.
 */
export function getJobOutput(args: GetJobOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobResult> {
    return pulumi.output(args).apply((a: any) => getJob(a, opts))
}

export interface GetJobOutputArgs {
    jobId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
