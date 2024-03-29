// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const CapacityCommitmentEdition = {
    /**
     * Default value, which will be treated as ENTERPRISE.
     */
    EditionUnspecified: "EDITION_UNSPECIFIED",
    /**
     * Standard edition.
     */
    Standard: "STANDARD",
    /**
     * Enterprise edition.
     */
    Enterprise: "ENTERPRISE",
    /**
     * Enterprise plus edition.
     */
    EnterprisePlus: "ENTERPRISE_PLUS",
} as const;

/**
 * Edition of the capacity commitment.
 */
export type CapacityCommitmentEdition = (typeof CapacityCommitmentEdition)[keyof typeof CapacityCommitmentEdition];

export const CapacityCommitmentPlan = {
    /**
     * Invalid plan value. Requests with this value will be rejected with error code `google.rpc.Code.INVALID_ARGUMENT`.
     */
    CommitmentPlanUnspecified: "COMMITMENT_PLAN_UNSPECIFIED",
    /**
     * Flex commitments have committed period of 1 minute after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Flex: "FLEX",
    /**
     * Same as FLEX, should only be used if flat-rate commitments are still available.
     */
    FlexFlatRate: "FLEX_FLAT_RATE",
    /**
     * Trial commitments have a committed period of 182 days after becoming ACTIVE. After that, they are converted to a new commitment based on the `renewal_plan`. Default `renewal_plan` for Trial commitment is Flex so that it can be deleted right after committed period ends.
     */
    Trial: "TRIAL",
    /**
     * Monthly commitments have a committed period of 30 days after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Monthly: "MONTHLY",
    /**
     * Same as MONTHLY, should only be used if flat-rate commitments are still available.
     */
    MonthlyFlatRate: "MONTHLY_FLAT_RATE",
    /**
     * Annual commitments have a committed period of 365 days after becoming ACTIVE. After that they are converted to a new commitment based on the renewal_plan.
     */
    Annual: "ANNUAL",
    /**
     * Same as ANNUAL, should only be used if flat-rate commitments are still available.
     */
    AnnualFlatRate: "ANNUAL_FLAT_RATE",
    /**
     * 3-year commitments have a committed period of 1095(3 * 365) days after becoming ACTIVE. After that they are converted to a new commitment based on the renewal_plan.
     */
    ThreeYear: "THREE_YEAR",
    /**
     * Should only be used for `renewal_plan` and is only meaningful if edition is specified to values other than EDITION_UNSPECIFIED. Otherwise CreateCapacityCommitmentRequest or UpdateCapacityCommitmentRequest will be rejected with error code `google.rpc.Code.INVALID_ARGUMENT`. If the renewal_plan is NONE, capacity commitment will be removed at the end of its commitment period.
     */
    None: "NONE",
} as const;

/**
 * Capacity commitment commitment plan.
 */
export type CapacityCommitmentPlan = (typeof CapacityCommitmentPlan)[keyof typeof CapacityCommitmentPlan];

export const CapacityCommitmentRenewalPlan = {
    /**
     * Invalid plan value. Requests with this value will be rejected with error code `google.rpc.Code.INVALID_ARGUMENT`.
     */
    CommitmentPlanUnspecified: "COMMITMENT_PLAN_UNSPECIFIED",
    /**
     * Flex commitments have committed period of 1 minute after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Flex: "FLEX",
    /**
     * Same as FLEX, should only be used if flat-rate commitments are still available.
     */
    FlexFlatRate: "FLEX_FLAT_RATE",
    /**
     * Trial commitments have a committed period of 182 days after becoming ACTIVE. After that, they are converted to a new commitment based on the `renewal_plan`. Default `renewal_plan` for Trial commitment is Flex so that it can be deleted right after committed period ends.
     */
    Trial: "TRIAL",
    /**
     * Monthly commitments have a committed period of 30 days after becoming ACTIVE. After that, they are not in a committed period anymore and can be removed any time.
     */
    Monthly: "MONTHLY",
    /**
     * Same as MONTHLY, should only be used if flat-rate commitments are still available.
     */
    MonthlyFlatRate: "MONTHLY_FLAT_RATE",
    /**
     * Annual commitments have a committed period of 365 days after becoming ACTIVE. After that they are converted to a new commitment based on the renewal_plan.
     */
    Annual: "ANNUAL",
    /**
     * Same as ANNUAL, should only be used if flat-rate commitments are still available.
     */
    AnnualFlatRate: "ANNUAL_FLAT_RATE",
    /**
     * 3-year commitments have a committed period of 1095(3 * 365) days after becoming ACTIVE. After that they are converted to a new commitment based on the renewal_plan.
     */
    ThreeYear: "THREE_YEAR",
    /**
     * Should only be used for `renewal_plan` and is only meaningful if edition is specified to values other than EDITION_UNSPECIFIED. Otherwise CreateCapacityCommitmentRequest or UpdateCapacityCommitmentRequest will be rejected with error code `google.rpc.Code.INVALID_ARGUMENT`. If the renewal_plan is NONE, capacity commitment will be removed at the end of its commitment period.
     */
    None: "NONE",
} as const;

/**
 * The plan this capacity commitment is converted to after commitment_end_time passes. Once the plan is changed, committed period is extended according to commitment plan. Only applicable for ANNUAL and TRIAL commitments.
 */
export type CapacityCommitmentRenewalPlan = (typeof CapacityCommitmentRenewalPlan)[keyof typeof CapacityCommitmentRenewalPlan];

export const ReservationEdition = {
    /**
     * Default value, which will be treated as ENTERPRISE.
     */
    EditionUnspecified: "EDITION_UNSPECIFIED",
    /**
     * Standard edition.
     */
    Standard: "STANDARD",
    /**
     * Enterprise edition.
     */
    Enterprise: "ENTERPRISE",
    /**
     * Enterprise plus edition.
     */
    EnterprisePlus: "ENTERPRISE_PLUS",
} as const;

/**
 * Edition of the reservation.
 */
export type ReservationEdition = (typeof ReservationEdition)[keyof typeof ReservationEdition];
