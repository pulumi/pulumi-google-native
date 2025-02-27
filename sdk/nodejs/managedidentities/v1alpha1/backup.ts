// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a Backup for a domain.
 * Auto-naming is currently not supported for this resource.
 */
export class Backup extends pulumi.CustomResource {
    /**
     * Get an existing Backup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Backup {
        return new Backup(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:managedidentities/v1alpha1:Backup';

    /**
     * Returns true if the given object is an instance of Backup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Backup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Backup.__pulumiType;
    }

    /**
     * Required. Backup Id, unique name to identify the backups with the following restrictions: * Must be lowercase letters, numbers, and hyphens * Must start with a letter. * Must contain between 1-63 characters. * Must end with a number or a letter. * Must be unique within the domain.
     */
    public readonly backupId!: pulumi.Output<string>;
    /**
     * The time the backups was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. A short description of the backup.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly domainId!: pulumi.Output<string>;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The unique name of the Backup in the form of projects/{project_id}/locations/global/domains/{domain_name}/backups/{name}
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The current state of the backup.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current status of this backup, if available.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Indicates whether it’s an on-demand backup or scheduled.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Last update time.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Backup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackupArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.backupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backupId'");
            }
            if ((!args || args.domainId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainId'");
            }
            resourceInputs["backupId"] = args ? args.backupId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["backupId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["domainId"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["statusMessage"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["backupId", "domainId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Backup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Backup resource.
 */
export interface BackupArgs {
    /**
     * Required. Backup Id, unique name to identify the backups with the following restrictions: * Must be lowercase letters, numbers, and hyphens * Must start with a letter. * Must contain between 1-63 characters. * Must end with a number or a letter. * Must be unique within the domain.
     */
    backupId: pulumi.Input<string>;
    /**
     * Optional. A short description of the backup.
     */
    description?: pulumi.Input<string>;
    domainId: pulumi.Input<string>;
    /**
     * Optional. Resource labels to represent user provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    project?: pulumi.Input<string>;
}
