// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single Rollout.
 */
export function getRollout(args: GetRolloutArgs, opts?: pulumi.InvokeOptions): Promise<GetRolloutResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:clouddeploy/v1:getRollout", {
        "deliveryPipelineId": args.deliveryPipelineId,
        "location": args.location,
        "project": args.project,
        "releaseId": args.releaseId,
        "rolloutId": args.rolloutId,
    }, opts);
}

export interface GetRolloutArgs {
    deliveryPipelineId: string;
    location: string;
    project?: string;
    releaseId: string;
    rolloutId: string;
}

export interface GetRolloutResult {
    /**
     * User annotations. These attributes can only be set and used by the user, and not by Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
     */
    readonly annotations: {[key: string]: string};
    /**
     * Approval state of the `Rollout`.
     */
    readonly approvalState: string;
    /**
     * Time at which the `Rollout` was approved.
     */
    readonly approveTime: string;
    /**
     * Name of the `ControllerRollout`. Format is `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/a-z{0,62}`.
     */
    readonly controllerRollout: string;
    /**
     * Time at which the `Rollout` was created.
     */
    readonly createTime: string;
    /**
     * Time at which the `Rollout` finished deploying.
     */
    readonly deployEndTime: string;
    /**
     * The reason this rollout failed. This will always be unspecified while the rollout is in progress.
     */
    readonly deployFailureCause: string;
    /**
     * Time at which the `Rollout` started deploying.
     */
    readonly deployStartTime: string;
    /**
     * The resource name of the Cloud Build `Build` object that is used to deploy the Rollout. Format is `projects/{project}/locations/{location}/builds/{build}`.
     */
    readonly deployingBuild: string;
    /**
     * Description of the `Rollout` for user purposes. Max length is 255 characters.
     */
    readonly description: string;
    /**
     * Time at which the `Rollout` was enqueued.
     */
    readonly enqueueTime: string;
    /**
     * This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
     */
    readonly etag: string;
    /**
     * Additional information about the rollout failure, if available.
     */
    readonly failureReason: string;
    /**
     * Labels are attributes that can be set and used by both the user and by Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
     */
    readonly labels: {[key: string]: string};
    /**
     * Metadata contains information about the rollout.
     */
    readonly metadata: outputs.clouddeploy.v1.MetadataResponse;
    /**
     * Optional. Name of the `Rollout`. Format is `projects/{project}/locations/{location}/deliveryPipelines/{deliveryPipeline}/releases/{release}/rollouts/a-z{0,62}`.
     */
    readonly name: string;
    /**
     * The phases that represent the workflows of this `Rollout`.
     */
    readonly phases: outputs.clouddeploy.v1.PhaseResponse[];
    /**
     * Name of the `Rollout` that is rolled back by this `Rollout`. Empty if this `Rollout` wasn't created as a rollback.
     */
    readonly rollbackOfRollout: string;
    /**
     * Names of `Rollouts` that rolled back this `Rollout`.
     */
    readonly rolledBackByRollouts: string[];
    /**
     * Current state of the `Rollout`.
     */
    readonly state: string;
    /**
     * The ID of Target to which this `Rollout` is deploying.
     */
    readonly targetId: string;
    /**
     * Unique identifier of the `Rollout`.
     */
    readonly uid: string;
}
/**
 * Gets details of a single Rollout.
 */
export function getRolloutOutput(args: GetRolloutOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRolloutResult> {
    return pulumi.output(args).apply((a: any) => getRollout(a, opts))
}

export interface GetRolloutOutputArgs {
    deliveryPipelineId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    releaseId: pulumi.Input<string>;
    rolloutId: pulumi.Input<string>;
}
