// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BillingAccountSinkOutputVersionFormat = {
    /**
     * An unspecified format version that will default to V2.
     */
    VersionFormatUnspecified: "VERSION_FORMAT_UNSPECIFIED",
    /**
     * LogEntry version 2 format.
     */
    V2: "V2",
    /**
     * LogEntry version 1 format.
     */
    V1: "V1",
} as const;

/**
 * Deprecated. This field is unused.
 */
export type BillingAccountSinkOutputVersionFormat = (typeof BillingAccountSinkOutputVersionFormat)[keyof typeof BillingAccountSinkOutputVersionFormat];