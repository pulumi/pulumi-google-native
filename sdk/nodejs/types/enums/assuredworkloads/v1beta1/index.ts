// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType = {
    /**
     * Unknown resource type.
     */
    ResourceTypeUnspecified: "RESOURCE_TYPE_UNSPECIFIED",
    /**
     * Deprecated. Existing workloads will continue to support this, but new CreateWorkloadRequests should not specify this as an input value.
     */
    ConsumerProject: "CONSUMER_PROJECT",
    /**
     * Consumer Folder.
     */
    ConsumerFolder: "CONSUMER_FOLDER",
    /**
     * Consumer project containing encryption keys.
     */
    EncryptionKeysProject: "ENCRYPTION_KEYS_PROJECT",
    /**
     * Keyring resource that hosts encryption keys.
     */
    Keyring: "KEYRING",
} as const;

/**
 * Indicates the type of resource. This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT)
 */
export type GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType = (typeof GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType)[keyof typeof GoogleCloudAssuredworkloadsV1beta1WorkloadResourceSettingsResourceType];

export const WorkloadComplianceRegime = {
    /**
     * Unknown compliance regime.
     */
    ComplianceRegimeUnspecified: "COMPLIANCE_REGIME_UNSPECIFIED",
    /**
     * Information protection as per DoD IL4 requirements.
     */
    Il4: "IL4",
    /**
     * Criminal Justice Information Services (CJIS) Security policies.
     */
    Cjis: "CJIS",
    /**
     * FedRAMP High data protection controls
     */
    FedrampHigh: "FEDRAMP_HIGH",
    /**
     * FedRAMP Moderate data protection controls
     */
    FedrampModerate: "FEDRAMP_MODERATE",
    /**
     * Assured Workloads For US Regions data protection controls
     */
    UsRegionalAccess: "US_REGIONAL_ACCESS",
    /**
     * Health Insurance Portability and Accountability Act controls
     */
    Hipaa: "HIPAA",
    /**
     * Health Information Trust Alliance controls
     */
    Hitrust: "HITRUST",
    /**
     * Assured Workloads For EU Regions and Support controls
     */
    EuRegionsAndSupport: "EU_REGIONS_AND_SUPPORT",
    /**
     * Assured Workloads For Canada Regions and Support controls
     */
    CaRegionsAndSupport: "CA_REGIONS_AND_SUPPORT",
    /**
     * International Traffic in Arms Regulations
     */
    Itar: "ITAR",
    /**
     * Assured Workloads for Australia Regions and Support controls Available for public preview consumption. Don't create production workloads.
     */
    AuRegionsAndUsSupport: "AU_REGIONS_AND_US_SUPPORT",
    /**
     * Assured Workloads for Partners;
     */
    AssuredWorkloadsForPartners: "ASSURED_WORKLOADS_FOR_PARTNERS",
} as const;

/**
 * Required. Immutable. Compliance Regime associated with this workload.
 */
export type WorkloadComplianceRegime = (typeof WorkloadComplianceRegime)[keyof typeof WorkloadComplianceRegime];

export const WorkloadPartner = {
    /**
     * Unknown partner regime/controls.
     */
    PartnerUnspecified: "PARTNER_UNSPECIFIED",
    /**
     * S3NS regime/controls.
     */
    LocalControlsByS3ns: "LOCAL_CONTROLS_BY_S3NS",
} as const;

/**
 * Optional. Compliance Regime associated with this workload.
 */
export type WorkloadPartner = (typeof WorkloadPartner)[keyof typeof WorkloadPartner];