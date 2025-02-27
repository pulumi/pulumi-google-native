// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets details of a single TargetProject. NOTE: TargetProject is a global resource; hence the only supported value for location is `global`.
 */
export function getTargetProject(args: GetTargetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetProjectResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:vmmigration/v1alpha1:getTargetProject", {
        "location": args.location,
        "project": args.project,
        "targetProjectId": args.targetProjectId,
    }, opts);
}

export interface GetTargetProjectArgs {
    location: string;
    project?: string;
    targetProjectId: string;
}

export interface GetTargetProjectResult {
    /**
     * The time this target project resource was created (not related to when the Compute Engine project it points to was created).
     */
    readonly createTime: string;
    /**
     * The target project's description.
     */
    readonly description: string;
    /**
     * The name of the target project.
     */
    readonly name: string;
    /**
     * The target project ID (number) or project name.
     */
    readonly project: string;
    /**
     * The last time the target project resource was updated.
     */
    readonly updateTime: string;
}
/**
 * Gets details of a single TargetProject. NOTE: TargetProject is a global resource; hence the only supported value for location is `global`.
 */
export function getTargetProjectOutput(args: GetTargetProjectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTargetProjectResult> {
    return pulumi.output(args).apply((a: any) => getTargetProject(a, opts))
}

export interface GetTargetProjectOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    targetProjectId: pulumi.Input<string>;
}
