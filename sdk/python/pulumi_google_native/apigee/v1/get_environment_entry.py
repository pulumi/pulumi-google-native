# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetEnvironmentEntryResult',
    'AwaitableGetEnvironmentEntryResult',
    'get_environment_entry',
    'get_environment_entry_output',
]

@pulumi.output_type
class GetEnvironmentEntryResult:
    def __init__(__self__, name=None, value=None):
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource URI that can be used to identify the scope of the key value map entries.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Data or payload that is being retrieved and associated with the unique key.
        """
        return pulumi.get(self, "value")


class AwaitableGetEnvironmentEntryResult(GetEnvironmentEntryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEnvironmentEntryResult(
            name=self.name,
            value=self.value)


def get_environment_entry(entry_id: Optional[str] = None,
                          environment_id: Optional[str] = None,
                          keyvaluemap_id: Optional[str] = None,
                          organization_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEnvironmentEntryResult:
    """
    Get the Key value entry value for org, env or apis scoped Key value map.
    """
    __args__ = dict()
    __args__['entryId'] = entry_id
    __args__['environmentId'] = environment_id
    __args__['keyvaluemapId'] = keyvaluemap_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getEnvironmentEntry', __args__, opts=opts, typ=GetEnvironmentEntryResult).value

    return AwaitableGetEnvironmentEntryResult(
        name=__ret__.name,
        value=__ret__.value)


@_utilities.lift_output_func(get_environment_entry)
def get_environment_entry_output(entry_id: Optional[pulumi.Input[str]] = None,
                                 environment_id: Optional[pulumi.Input[str]] = None,
                                 keyvaluemap_id: Optional[pulumi.Input[str]] = None,
                                 organization_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEnvironmentEntryResult]:
    """
    Get the Key value entry value for org, env or apis scoped Key value map.
    """
    ...