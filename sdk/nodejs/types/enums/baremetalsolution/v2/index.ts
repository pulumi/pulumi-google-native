// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const AllowedClientMountPermissions = {
    /**
     * Permissions were not specified.
     */
    MountPermissionsUnspecified: "MOUNT_PERMISSIONS_UNSPECIFIED",
    /**
     * NFS share can be mount with read-only permissions.
     */
    Read: "READ",
    /**
     * NFS share can be mount with read-write permissions.
     */
    ReadWrite: "READ_WRITE",
} as const;

/**
 * Mount permissions.
 */
export type AllowedClientMountPermissions = (typeof AllowedClientMountPermissions)[keyof typeof AllowedClientMountPermissions];

export const InstanceConfigNetworkConfig = {
    /**
     * The unspecified network configuration.
     */
    NetworkconfigUnspecified: "NETWORKCONFIG_UNSPECIFIED",
    /**
     * Instance part of single client network and single private network.
     */
    SingleVlan: "SINGLE_VLAN",
    /**
     * Instance part of multiple (or single) client networks and private networks.
     */
    MultiVlan: "MULTI_VLAN",
} as const;

/**
 * The type of network configuration on the instance.
 */
export type InstanceConfigNetworkConfig = (typeof InstanceConfigNetworkConfig)[keyof typeof InstanceConfigNetworkConfig];

export const LogicalNetworkInterfaceNetworkType = {
    /**
     * Unspecified value.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Client network, a network peered to a Google Cloud VPC.
     */
    Client: "CLIENT",
    /**
     * Private network, a network local to the Bare Metal Solution environment.
     */
    Private: "PRIVATE",
} as const;

/**
 * Type of network.
 */
export type LogicalNetworkInterfaceNetworkType = (typeof LogicalNetworkInterfaceNetworkType)[keyof typeof LogicalNetworkInterfaceNetworkType];

export const NetworkConfigBandwidth = {
    /**
     * Unspecified value.
     */
    BandwidthUnspecified: "BANDWIDTH_UNSPECIFIED",
    /**
     * 1 Gbps.
     */
    Bw1Gbps: "BW_1_GBPS",
    /**
     * 2 Gbps.
     */
    Bw2Gbps: "BW_2_GBPS",
    /**
     * 5 Gbps.
     */
    Bw5Gbps: "BW_5_GBPS",
    /**
     * 10 Gbps.
     */
    Bw10Gbps: "BW_10_GBPS",
} as const;

/**
 * Interconnect bandwidth. Set only when type is CLIENT.
 */
export type NetworkConfigBandwidth = (typeof NetworkConfigBandwidth)[keyof typeof NetworkConfigBandwidth];

export const NetworkConfigServiceCidr = {
    /**
     * Unspecified value.
     */
    ServiceCidrUnspecified: "SERVICE_CIDR_UNSPECIFIED",
    /**
     * Services are disabled for the given network.
     */
    Disabled: "DISABLED",
    /**
     * Use the highest /26 block of the network to host services.
     */
    High26: "HIGH_26",
    /**
     * Use the highest /27 block of the network to host services.
     */
    High27: "HIGH_27",
    /**
     * Use the highest /28 block of the network to host services.
     */
    High28: "HIGH_28",
} as const;

/**
 * Service CIDR, if any.
 */
export type NetworkConfigServiceCidr = (typeof NetworkConfigServiceCidr)[keyof typeof NetworkConfigServiceCidr];

export const NetworkConfigType = {
    /**
     * Unspecified value.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Client network, that is a network peered to a GCP VPC.
     */
    Client: "CLIENT",
    /**
     * Private network, that is a network local to the BMS POD.
     */
    Private: "PRIVATE",
} as const;

/**
 * The type of this network, either Client or Private.
 */
export type NetworkConfigType = (typeof NetworkConfigType)[keyof typeof NetworkConfigType];

export const NfsExportPermissions = {
    /**
     * Unspecified value.
     */
    PermissionsUnspecified: "PERMISSIONS_UNSPECIFIED",
    /**
     * Read-only permission.
     */
    ReadOnly: "READ_ONLY",
    /**
     * Read-write permission.
     */
    ReadWrite: "READ_WRITE",
} as const;

/**
 * Export permissions.
 */
export type NfsExportPermissions = (typeof NfsExportPermissions)[keyof typeof NfsExportPermissions];

export const NfsShareStorageType = {
    /**
     * The storage type for this volume is unknown.
     */
    StorageTypeUnspecified: "STORAGE_TYPE_UNSPECIFIED",
    /**
     * The storage type for this volume is SSD.
     */
    Ssd: "SSD",
    /**
     * This storage type for this volume is HDD.
     */
    Hdd: "HDD",
} as const;

/**
 * Immutable. The storage type of the underlying volume.
 */
export type NfsShareStorageType = (typeof NfsShareStorageType)[keyof typeof NfsShareStorageType];

export const VolumeConfigPerformanceTier = {
    /**
     * Value is not specified.
     */
    VolumePerformanceTierUnspecified: "VOLUME_PERFORMANCE_TIER_UNSPECIFIED",
    /**
     * Regular volumes, shared aggregates.
     */
    VolumePerformanceTierShared: "VOLUME_PERFORMANCE_TIER_SHARED",
    /**
     * Assigned aggregates.
     */
    VolumePerformanceTierAssigned: "VOLUME_PERFORMANCE_TIER_ASSIGNED",
    /**
     * High throughput aggregates.
     */
    VolumePerformanceTierHt: "VOLUME_PERFORMANCE_TIER_HT",
} as const;

/**
 * Performance tier of the Volume. Default is SHARED.
 */
export type VolumeConfigPerformanceTier = (typeof VolumeConfigPerformanceTier)[keyof typeof VolumeConfigPerformanceTier];

export const VolumeConfigProtocol = {
    /**
     * Unspecified value.
     */
    ProtocolUnspecified: "PROTOCOL_UNSPECIFIED",
    /**
     * Fibre channel.
     */
    ProtocolFc: "PROTOCOL_FC",
    /**
     * Network file system.
     */
    ProtocolNfs: "PROTOCOL_NFS",
} as const;

/**
 * Volume protocol.
 */
export type VolumeConfigProtocol = (typeof VolumeConfigProtocol)[keyof typeof VolumeConfigProtocol];

export const VolumeConfigType = {
    /**
     * The unspecified type.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * This Volume is on flash.
     */
    Flash: "FLASH",
    /**
     * This Volume is on disk.
     */
    Disk: "DISK",
} as const;

/**
 * The type of this Volume.
 */
export type VolumeConfigType = (typeof VolumeConfigType)[keyof typeof VolumeConfigType];
