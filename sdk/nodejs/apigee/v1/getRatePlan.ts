// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets the details of a rate plan.
 */
export function getRatePlan(args: GetRatePlanArgs, opts?: pulumi.InvokeOptions): Promise<GetRatePlanResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:apigee/v1:getRatePlan", {
        "apiproductId": args.apiproductId,
        "organizationId": args.organizationId,
        "rateplanId": args.rateplanId,
    }, opts);
}

export interface GetRatePlanArgs {
    apiproductId: string;
    organizationId: string;
    rateplanId: string;
}

export interface GetRatePlanResult {
    /**
     * Name of the API product that the rate plan is associated with.
     */
    readonly apiproduct: string;
    /**
     * Frequency at which the customer will be billed.
     */
    readonly billingPeriod: string;
    /**
     * API call volume ranges and the fees charged when the total number of API calls is within a given range. The method used to calculate the final fee depends on the selected pricing model. For example, if the pricing model is `STAIRSTEP` and the ranges are defined as follows: ``` { "start": 1, "end": 100, "fee": 75 }, { "start": 101, "end": 200, "fee": 100 }, } ``` Then the following fees would be charged based on the total number of API calls (assuming the currency selected is `USD`): * 1 call costs $75 * 50 calls cost $75 * 150 calls cost $100 The number of API calls cannot exceed 200.
     */
    readonly consumptionPricingRates: outputs.apigee.v1.GoogleCloudApigeeV1RateRangeResponse[];
    /**
     * Pricing model used for consumption-based charges.
     */
    readonly consumptionPricingType: string;
    /**
     * Time that the rate plan was created in milliseconds since epoch.
     */
    readonly createdAt: string;
    /**
     * Currency to be used for billing. Consists of a three-letter code as defined by the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) standard.
     */
    readonly currencyCode: string;
    /**
     * Description of the rate plan.
     */
    readonly description: string;
    /**
     * Display name of the rate plan.
     */
    readonly displayName: string;
    /**
     * Time when the rate plan will expire in milliseconds since epoch. Set to 0 or `null` to indicate that the rate plan should never expire.
     */
    readonly endTime: string;
    /**
     * Frequency at which the fixed fee is charged.
     */
    readonly fixedFeeFrequency: number;
    /**
     * Fixed amount that is charged at a defined interval and billed in advance of use of the API product. The fee will be prorated for the first billing period.
     */
    readonly fixedRecurringFee: outputs.apigee.v1.GoogleTypeMoneyResponse;
    /**
     * Time the rate plan was last modified in milliseconds since epoch.
     */
    readonly lastModifiedAt: string;
    /**
     * Name of the rate plan.
     */
    readonly name: string;
    /**
     * DEPRECATED: This field is no longer supported and will eventually be removed when Apigee Hybrid 1.5/1.6 is no longer supported. Instead, use the `billingType` field inside `DeveloperMonetizationConfig` resource. Flag that specifies the billing account type, prepaid or postpaid.
     *
     * @deprecated DEPRECATED: This field is no longer supported and will eventually be removed when Apigee Hybrid 1.5/1.6 is no longer supported. Instead, use the `billingType` field inside `DeveloperMonetizationConfig` resource. Flag that specifies the billing account type, prepaid or postpaid.
     */
    readonly paymentFundingModel: string;
    /**
     * Details of the revenue sharing model.
     */
    readonly revenueShareRates: outputs.apigee.v1.GoogleCloudApigeeV1RevenueShareRangeResponse[];
    /**
     * Method used to calculate the revenue that is shared with developers.
     */
    readonly revenueShareType: string;
    /**
     * Initial, one-time fee paid when purchasing the API product.
     */
    readonly setupFee: outputs.apigee.v1.GoogleTypeMoneyResponse;
    /**
     * Time when the rate plan becomes active in milliseconds since epoch.
     */
    readonly startTime: string;
    /**
     * Current state of the rate plan (draft or published).
     */
    readonly state: string;
}

export function getRatePlanOutput(args: GetRatePlanOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRatePlanResult> {
    return pulumi.output(args).apply(a => getRatePlan(a, opts))
}

export interface GetRatePlanOutputArgs {
    apiproductId: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    rateplanId: pulumi.Input<string>;
}