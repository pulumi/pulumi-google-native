// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AuditLogConfigLogType = {
    /**
     * Default case. Should never be this.
     */
    LogTypeUnspecified: "LOG_TYPE_UNSPECIFIED",
    /**
     * Admin reads. Example: CloudIAM getIamPolicy
     */
    AdminRead: "ADMIN_READ",
    /**
     * Data writes. Example: CloudSQL Users create
     */
    DataWrite: "DATA_WRITE",
    /**
     * Data reads. Example: CloudSQL Users list
     */
    DataRead: "DATA_READ",
} as const;

/**
 * The log type that this config enables.
 */
export type AuditLogConfigLogType = (typeof AuditLogConfigLogType)[keyof typeof AuditLogConfigLogType];

export const GceRegionalPersistentDiskReclaimPolicy = {
    /**
     * Do not use.
     */
    ReclaimPolicyUnspecified: "RECLAIM_POLICY_UNSPECIFIED",
    /**
     * The persistent disk will be deleted with the Workstation.
     */
    Delete: "DELETE",
    /**
     * The persistent disk will be remain after the workstation is deleted, and the administrator must manually delete the disk.
     */
    Retain: "RETAIN",
} as const;

/**
 * What should happen to the disk after the Workstation is deleted. Defaults to DELETE.
 */
export type GceRegionalPersistentDiskReclaimPolicy = (typeof GceRegionalPersistentDiskReclaimPolicy)[keyof typeof GceRegionalPersistentDiskReclaimPolicy];