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

export const CustomDomainState = {
    /**
     * Unspecified state.
     */
    CustomDomainStateUnspecified: "CUSTOM_DOMAIN_STATE_UNSPECIFIED",
    /**
     * DNS record is not created.
     */
    Unverified: "UNVERIFIED",
    /**
     * DNS record is created.
     */
    Verified: "VERIFIED",
    /**
     * Calling SLM to update.
     */
    Modifying: "MODIFYING",
    /**
     * ManagedCertificate is ready.
     */
    Available: "AVAILABLE",
    /**
     * ManagedCertificate is not ready.
     */
    Unavailable: "UNAVAILABLE",
    /**
     * Status is not known.
     */
    Unknown: "UNKNOWN",
} as const;

/**
 * Domain state.
 */
export type CustomDomainState = (typeof CustomDomainState)[keyof typeof CustomDomainState];

export const InstancePlatformEdition = {
    /**
     * Platform edition is unspecified.
     */
    PlatformEditionUnspecified: "PLATFORM_EDITION_UNSPECIFIED",
    /**
     * Trial.
     */
    LookerCoreTrial: "LOOKER_CORE_TRIAL",
    /**
     * Standard.
     */
    LookerCoreStandard: "LOOKER_CORE_STANDARD",
    /**
     * Subscription Standard.
     */
    LookerCoreStandardAnnual: "LOOKER_CORE_STANDARD_ANNUAL",
    /**
     * Subscription Enterprise.
     */
    LookerCoreEnterpriseAnnual: "LOOKER_CORE_ENTERPRISE_ANNUAL",
    /**
     * Subscription Embed.
     */
    LookerCoreEmbedAnnual: "LOOKER_CORE_EMBED_ANNUAL",
} as const;

/**
 * Platform edition.
 */
export type InstancePlatformEdition = (typeof InstancePlatformEdition)[keyof typeof InstancePlatformEdition];

export const MaintenanceWindowDayOfWeek = {
    /**
     * The day of the week is unspecified.
     */
    DayOfWeekUnspecified: "DAY_OF_WEEK_UNSPECIFIED",
    /**
     * Monday
     */
    Monday: "MONDAY",
    /**
     * Tuesday
     */
    Tuesday: "TUESDAY",
    /**
     * Wednesday
     */
    Wednesday: "WEDNESDAY",
    /**
     * Thursday
     */
    Thursday: "THURSDAY",
    /**
     * Friday
     */
    Friday: "FRIDAY",
    /**
     * Saturday
     */
    Saturday: "SATURDAY",
    /**
     * Sunday
     */
    Sunday: "SUNDAY",
} as const;

/**
 * Required. Day of the week for this MaintenanceWindow (in UTC).
 */
export type MaintenanceWindowDayOfWeek = (typeof MaintenanceWindowDayOfWeek)[keyof typeof MaintenanceWindowDayOfWeek];
