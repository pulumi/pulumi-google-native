// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const InstanceConnectMode = {
    /**
     * Not set.
     */
    ConnectModeUnspecified: "CONNECT_MODE_UNSPECIFIED",
    /**
     * Connect via direct peering to the Memorystore for Redis hosted service.
     */
    DirectPeering: "DIRECT_PEERING",
    /**
     * Connect your Memorystore for Redis instance using Private Service Access. Private services access provides an IP address range for multiple Google Cloud services, including Memorystore.
     */
    PrivateServiceAccess: "PRIVATE_SERVICE_ACCESS",
} as const;

/**
 * Optional. The network connect mode of the Redis instance. If not provided, the connect mode defaults to DIRECT_PEERING.
 */
export type InstanceConnectMode = (typeof InstanceConnectMode)[keyof typeof InstanceConnectMode];

export const InstanceReadReplicasMode = {
    /**
     * If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
     */
    ReadReplicasModeUnspecified: "READ_REPLICAS_MODE_UNSPECIFIED",
    /**
     * If disabled, read endpoint will not be provided and the instance cannot scale up or down the number of replicas.
     */
    ReadReplicasDisabled: "READ_REPLICAS_DISABLED",
    /**
     * If enabled, read endpoint will be provided and the instance can scale up and down the number of replicas. Not valid for basic tier.
     */
    ReadReplicasEnabled: "READ_REPLICAS_ENABLED",
} as const;

/**
 * Optional. Read replicas mode for the instance. Defaults to READ_REPLICAS_DISABLED.
 */
export type InstanceReadReplicasMode = (typeof InstanceReadReplicasMode)[keyof typeof InstanceReadReplicasMode];

export const InstanceTier = {
    /**
     * Not set.
     */
    TierUnspecified: "TIER_UNSPECIFIED",
    /**
     * BASIC tier: standalone instance
     */
    Basic: "BASIC",
    /**
     * STANDARD_HA tier: highly available primary/replica instances
     */
    StandardHa: "STANDARD_HA",
} as const;

/**
 * Required. The service tier of the instance.
 */
export type InstanceTier = (typeof InstanceTier)[keyof typeof InstanceTier];

export const InstanceTransitEncryptionMode = {
    /**
     * Not set.
     */
    TransitEncryptionModeUnspecified: "TRANSIT_ENCRYPTION_MODE_UNSPECIFIED",
    /**
     * Client to Server traffic encryption enabled with server authentication.
     */
    ServerAuthentication: "SERVER_AUTHENTICATION",
    /**
     * TLS is disabled for the instance.
     */
    Disabled: "DISABLED",
} as const;

/**
 * Optional. The TLS mode of the Redis instance. If not provided, TLS is disabled for the instance.
 */
export type InstanceTransitEncryptionMode = (typeof InstanceTransitEncryptionMode)[keyof typeof InstanceTransitEncryptionMode];

export const PersistenceConfigPersistenceMode = {
    /**
     * Not set.
     */
    PersistenceModeUnspecified: "PERSISTENCE_MODE_UNSPECIFIED",
    /**
     * Persistence is disabled for the instance, and any existing snapshots are deleted.
     */
    Disabled: "DISABLED",
    /**
     * RDB based Persistence is enabled.
     */
    Rdb: "RDB",
} as const;

/**
 * Optional. Controls whether Persistence features are enabled. If not provided, the existing value will be used.
 */
export type PersistenceConfigPersistenceMode = (typeof PersistenceConfigPersistenceMode)[keyof typeof PersistenceConfigPersistenceMode];

export const PersistenceConfigRdbSnapshotPeriod = {
    /**
     * Not set.
     */
    SnapshotPeriodUnspecified: "SNAPSHOT_PERIOD_UNSPECIFIED",
    /**
     * Snapshot every 1 hour.
     */
    OneHour: "ONE_HOUR",
    /**
     * Snapshot every 6 hours.
     */
    SixHours: "SIX_HOURS",
    /**
     * Snapshot every 12 hours.
     */
    TwelveHours: "TWELVE_HOURS",
    /**
     * Snapshot every 24 horus.
     */
    TwentyFourHours: "TWENTY_FOUR_HOURS",
} as const;

/**
 * Optional. Period between RDB snapshots. Snapshots will be attempted every period starting from the provided snapshot start time. For example, a start time of 01/01/2033 06:45 and SIX_HOURS snapshot period will do nothing until 01/01/2033, and then trigger snapshots every day at 06:45, 12:45, 18:45, and 00:45 the next day, and so on. If not provided, TWENTY_FOUR_HOURS will be used as default.
 */
export type PersistenceConfigRdbSnapshotPeriod = (typeof PersistenceConfigRdbSnapshotPeriod)[keyof typeof PersistenceConfigRdbSnapshotPeriod];

export const WeeklyMaintenanceWindowDay = {
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
 * Required. The day of week that maintenance updates occur.
 */
export type WeeklyMaintenanceWindowDay = (typeof WeeklyMaintenanceWindowDay)[keyof typeof WeeklyMaintenanceWindowDay];