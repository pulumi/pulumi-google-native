// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets details of a single MigratingVm.
 */
export function getMigratingVm(args: GetMigratingVmArgs, opts?: pulumi.InvokeOptions): Promise<GetMigratingVmResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:vmmigration/v1:getMigratingVm", {
        "location": args.location,
        "migratingVmId": args.migratingVmId,
        "project": args.project,
        "sourceId": args.sourceId,
        "view": args.view,
    }, opts);
}

export interface GetMigratingVmArgs {
    location: string;
    migratingVmId: string;
    project?: string;
    sourceId: string;
    view?: string;
}

export interface GetMigratingVmResult {
    /**
     * Details of the VM from an AWS source.
     */
    readonly awsSourceVmDetails: outputs.vmmigration.v1.AwsSourceVmDetailsResponse;
    /**
     * Details of the VM from an Azure source.
     */
    readonly azureSourceVmDetails: outputs.vmmigration.v1.AzureSourceVmDetailsResponse;
    /**
     * Details of the target Persistent Disks in Compute Engine.
     */
    readonly computeEngineDisksTargetDefaults: outputs.vmmigration.v1.ComputeEngineDisksTargetDefaultsResponse;
    /**
     * Details of the target VM in Compute Engine.
     */
    readonly computeEngineTargetDefaults: outputs.vmmigration.v1.ComputeEngineTargetDefaultsResponse;
    /**
     * The time the migrating VM was created (this refers to this resource and not to the time it was installed in the source).
     */
    readonly createTime: string;
    /**
     * Details of the current running replication cycle.
     */
    readonly currentSyncInfo: outputs.vmmigration.v1.ReplicationCycleResponse;
    /**
     * Provides details of future CutoverJobs of a MigratingVm. Set to empty when cutover forecast is unavailable.
     */
    readonly cutoverForecast: outputs.vmmigration.v1.CutoverForecastResponse;
    /**
     * The description attached to the migrating VM by the user.
     */
    readonly description: string;
    /**
     * The display name attached to the MigratingVm by the user.
     */
    readonly displayName: string;
    /**
     * Provides details on the state of the Migrating VM in case of an error in replication.
     */
    readonly error: outputs.vmmigration.v1.StatusResponse;
    /**
     * The group this migrating vm is included in, if any. The group is represented by the full path of the appropriate Group resource.
     */
    readonly group: string;
    /**
     * The labels of the migrating VM.
     */
    readonly labels: {[key: string]: string};
    /**
     * Details of the last replication cycle. This will be updated whenever a replication cycle is finished and is not to be confused with last_sync which is only updated on successful replication cycles.
     */
    readonly lastReplicationCycle: outputs.vmmigration.v1.ReplicationCycleResponse;
    /**
     * The most updated snapshot created time in the source that finished replication.
     */
    readonly lastSync: outputs.vmmigration.v1.ReplicationSyncResponse;
    /**
     * The identifier of the MigratingVm.
     */
    readonly name: string;
    /**
     * The replication schedule policy.
     */
    readonly policy: outputs.vmmigration.v1.SchedulePolicyResponse;
    /**
     * The recent clone jobs performed on the migrating VM. This field holds the vm's last completed clone job and the vm's running clone job, if one exists. Note: To have this field populated you need to explicitly request it via the "view" parameter of the Get/List request.
     */
    readonly recentCloneJobs: outputs.vmmigration.v1.CloneJobResponse[];
    /**
     * The recent cutover jobs performed on the migrating VM. This field holds the vm's last completed cutover job and the vm's running cutover job, if one exists. Note: To have this field populated you need to explicitly request it via the "view" parameter of the Get/List request.
     */
    readonly recentCutoverJobs: outputs.vmmigration.v1.CutoverJobResponse[];
    /**
     * The unique ID of the VM in the source. The VM's name in vSphere can be changed, so this is not the VM's name but rather its moRef id. This id is of the form vm-.
     */
    readonly sourceVmId: string;
    /**
     * State of the MigratingVm.
     */
    readonly state: string;
    /**
     * The last time the migrating VM state was updated.
     */
    readonly stateTime: string;
    /**
     * The last time the migrating VM resource was updated.
     */
    readonly updateTime: string;
    /**
     * Details of the VM from a Vmware source.
     */
    readonly vmwareSourceVmDetails: outputs.vmmigration.v1.VmwareSourceVmDetailsResponse;
}
/**
 * Gets details of a single MigratingVm.
 */
export function getMigratingVmOutput(args: GetMigratingVmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMigratingVmResult> {
    return pulumi.output(args).apply((a: any) => getMigratingVm(a, opts))
}

export interface GetMigratingVmOutputArgs {
    location: pulumi.Input<string>;
    migratingVmId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    sourceId: pulumi.Input<string>;
    view?: pulumi.Input<string>;
}
