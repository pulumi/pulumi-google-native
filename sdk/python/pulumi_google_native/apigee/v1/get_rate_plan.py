# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetRatePlanResult',
    'AwaitableGetRatePlanResult',
    'get_rate_plan',
    'get_rate_plan_output',
]

@pulumi.output_type
class GetRatePlanResult:
    def __init__(__self__, apiproduct=None, billing_period=None, consumption_pricing_rates=None, consumption_pricing_type=None, created_at=None, currency_code=None, description=None, display_name=None, end_time=None, fixed_fee_frequency=None, fixed_recurring_fee=None, last_modified_at=None, name=None, revenue_share_rates=None, revenue_share_type=None, setup_fee=None, start_time=None, state=None):
        if apiproduct and not isinstance(apiproduct, str):
            raise TypeError("Expected argument 'apiproduct' to be a str")
        pulumi.set(__self__, "apiproduct", apiproduct)
        if billing_period and not isinstance(billing_period, str):
            raise TypeError("Expected argument 'billing_period' to be a str")
        pulumi.set(__self__, "billing_period", billing_period)
        if consumption_pricing_rates and not isinstance(consumption_pricing_rates, list):
            raise TypeError("Expected argument 'consumption_pricing_rates' to be a list")
        pulumi.set(__self__, "consumption_pricing_rates", consumption_pricing_rates)
        if consumption_pricing_type and not isinstance(consumption_pricing_type, str):
            raise TypeError("Expected argument 'consumption_pricing_type' to be a str")
        pulumi.set(__self__, "consumption_pricing_type", consumption_pricing_type)
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if currency_code and not isinstance(currency_code, str):
            raise TypeError("Expected argument 'currency_code' to be a str")
        pulumi.set(__self__, "currency_code", currency_code)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if fixed_fee_frequency and not isinstance(fixed_fee_frequency, int):
            raise TypeError("Expected argument 'fixed_fee_frequency' to be a int")
        pulumi.set(__self__, "fixed_fee_frequency", fixed_fee_frequency)
        if fixed_recurring_fee and not isinstance(fixed_recurring_fee, dict):
            raise TypeError("Expected argument 'fixed_recurring_fee' to be a dict")
        pulumi.set(__self__, "fixed_recurring_fee", fixed_recurring_fee)
        if last_modified_at and not isinstance(last_modified_at, str):
            raise TypeError("Expected argument 'last_modified_at' to be a str")
        pulumi.set(__self__, "last_modified_at", last_modified_at)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if revenue_share_rates and not isinstance(revenue_share_rates, list):
            raise TypeError("Expected argument 'revenue_share_rates' to be a list")
        pulumi.set(__self__, "revenue_share_rates", revenue_share_rates)
        if revenue_share_type and not isinstance(revenue_share_type, str):
            raise TypeError("Expected argument 'revenue_share_type' to be a str")
        pulumi.set(__self__, "revenue_share_type", revenue_share_type)
        if setup_fee and not isinstance(setup_fee, dict):
            raise TypeError("Expected argument 'setup_fee' to be a dict")
        pulumi.set(__self__, "setup_fee", setup_fee)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def apiproduct(self) -> str:
        """
        Name of the API product that the rate plan is associated with.
        """
        return pulumi.get(self, "apiproduct")

    @property
    @pulumi.getter(name="billingPeriod")
    def billing_period(self) -> str:
        """
        Frequency at which the customer will be billed.
        """
        return pulumi.get(self, "billing_period")

    @property
    @pulumi.getter(name="consumptionPricingRates")
    def consumption_pricing_rates(self) -> Sequence['outputs.GoogleCloudApigeeV1RateRangeResponse']:
        """
        API call volume ranges and the fees charged when the total number of API calls is within a given range. The method used to calculate the final fee depends on the selected pricing model. For example, if the pricing model is `STAIRSTEP` and the ranges are defined as follows: ``` { "start": 1, "end": 100, "fee": 75 }, { "start": 101, "end": 200, "fee": 100 }, } ``` Then the following fees would be charged based on the total number of API calls (assuming the currency selected is `USD`): * 1 call costs $75 * 50 calls cost $75 * 150 calls cost $100 The number of API calls cannot exceed 200.
        """
        return pulumi.get(self, "consumption_pricing_rates")

    @property
    @pulumi.getter(name="consumptionPricingType")
    def consumption_pricing_type(self) -> str:
        """
        Pricing model used for consumption-based charges.
        """
        return pulumi.get(self, "consumption_pricing_type")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Time that the rate plan was created in milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="currencyCode")
    def currency_code(self) -> str:
        """
        Currency to be used for billing. Consists of a three-letter code as defined by the [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) standard.
        """
        return pulumi.get(self, "currency_code")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the rate plan.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name of the rate plan.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Time when the rate plan will expire in milliseconds since epoch. Set to 0 or `null` to indicate that the rate plan should never expire.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="fixedFeeFrequency")
    def fixed_fee_frequency(self) -> int:
        """
        Frequency at which the fixed fee is charged.
        """
        return pulumi.get(self, "fixed_fee_frequency")

    @property
    @pulumi.getter(name="fixedRecurringFee")
    def fixed_recurring_fee(self) -> 'outputs.GoogleTypeMoneyResponse':
        """
        Fixed amount that is charged at a defined interval and billed in advance of use of the API product. The fee will be prorated for the first billing period.
        """
        return pulumi.get(self, "fixed_recurring_fee")

    @property
    @pulumi.getter(name="lastModifiedAt")
    def last_modified_at(self) -> str:
        """
        Time the rate plan was last modified in milliseconds since epoch.
        """
        return pulumi.get(self, "last_modified_at")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the rate plan.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="revenueShareRates")
    def revenue_share_rates(self) -> Sequence['outputs.GoogleCloudApigeeV1RevenueShareRangeResponse']:
        """
        Details of the revenue sharing model.
        """
        return pulumi.get(self, "revenue_share_rates")

    @property
    @pulumi.getter(name="revenueShareType")
    def revenue_share_type(self) -> str:
        """
        Method used to calculate the revenue that is shared with developers.
        """
        return pulumi.get(self, "revenue_share_type")

    @property
    @pulumi.getter(name="setupFee")
    def setup_fee(self) -> 'outputs.GoogleTypeMoneyResponse':
        """
        Initial, one-time fee paid when purchasing the API product.
        """
        return pulumi.get(self, "setup_fee")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Time when the rate plan becomes active in milliseconds since epoch.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the rate plan (draft or published).
        """
        return pulumi.get(self, "state")


class AwaitableGetRatePlanResult(GetRatePlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRatePlanResult(
            apiproduct=self.apiproduct,
            billing_period=self.billing_period,
            consumption_pricing_rates=self.consumption_pricing_rates,
            consumption_pricing_type=self.consumption_pricing_type,
            created_at=self.created_at,
            currency_code=self.currency_code,
            description=self.description,
            display_name=self.display_name,
            end_time=self.end_time,
            fixed_fee_frequency=self.fixed_fee_frequency,
            fixed_recurring_fee=self.fixed_recurring_fee,
            last_modified_at=self.last_modified_at,
            name=self.name,
            revenue_share_rates=self.revenue_share_rates,
            revenue_share_type=self.revenue_share_type,
            setup_fee=self.setup_fee,
            start_time=self.start_time,
            state=self.state)


def get_rate_plan(apiproduct_id: Optional[str] = None,
                  organization_id: Optional[str] = None,
                  rateplan_id: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRatePlanResult:
    """
    Gets the details of a rate plan.
    """
    __args__ = dict()
    __args__['apiproductId'] = apiproduct_id
    __args__['organizationId'] = organization_id
    __args__['rateplanId'] = rateplan_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getRatePlan', __args__, opts=opts, typ=GetRatePlanResult).value

    return AwaitableGetRatePlanResult(
        apiproduct=__ret__.apiproduct,
        billing_period=__ret__.billing_period,
        consumption_pricing_rates=__ret__.consumption_pricing_rates,
        consumption_pricing_type=__ret__.consumption_pricing_type,
        created_at=__ret__.created_at,
        currency_code=__ret__.currency_code,
        description=__ret__.description,
        display_name=__ret__.display_name,
        end_time=__ret__.end_time,
        fixed_fee_frequency=__ret__.fixed_fee_frequency,
        fixed_recurring_fee=__ret__.fixed_recurring_fee,
        last_modified_at=__ret__.last_modified_at,
        name=__ret__.name,
        revenue_share_rates=__ret__.revenue_share_rates,
        revenue_share_type=__ret__.revenue_share_type,
        setup_fee=__ret__.setup_fee,
        start_time=__ret__.start_time,
        state=__ret__.state)


@_utilities.lift_output_func(get_rate_plan)
def get_rate_plan_output(apiproduct_id: Optional[pulumi.Input[str]] = None,
                         organization_id: Optional[pulumi.Input[str]] = None,
                         rateplan_id: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRatePlanResult]:
    """
    Gets the details of a rate plan.
    """
    ...