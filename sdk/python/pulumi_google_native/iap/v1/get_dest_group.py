# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetDestGroupResult',
    'AwaitableGetDestGroupResult',
    'get_dest_group',
    'get_dest_group_output',
]

@pulumi.output_type
class GetDestGroupResult:
    def __init__(__self__, cidrs=None, fqdns=None, name=None):
        if cidrs and not isinstance(cidrs, list):
            raise TypeError("Expected argument 'cidrs' to be a list")
        pulumi.set(__self__, "cidrs", cidrs)
        if fqdns and not isinstance(fqdns, list):
            raise TypeError("Expected argument 'fqdns' to be a list")
        pulumi.set(__self__, "fqdns", fqdns)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def cidrs(self) -> Sequence[str]:
        """
        null List of CIDRs that this group applies to.
        """
        return pulumi.get(self, "cidrs")

    @property
    @pulumi.getter
    def fqdns(self) -> Sequence[str]:
        """
        null List of FQDNs that this group applies to.
        """
        return pulumi.get(self, "fqdns")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. Identifier for the TunnelDestGroup. Must be unique within the project.
        """
        return pulumi.get(self, "name")


class AwaitableGetDestGroupResult(GetDestGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDestGroupResult(
            cidrs=self.cidrs,
            fqdns=self.fqdns,
            name=self.name)


def get_dest_group(dest_group_id: Optional[str] = None,
                   location: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDestGroupResult:
    """
    Retrieves an existing TunnelDestGroup.
    """
    __args__ = dict()
    __args__['destGroupId'] = dest_group_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:iap/v1:getDestGroup', __args__, opts=opts, typ=GetDestGroupResult).value

    return AwaitableGetDestGroupResult(
        cidrs=__ret__.cidrs,
        fqdns=__ret__.fqdns,
        name=__ret__.name)


@_utilities.lift_output_func(get_dest_group)
def get_dest_group_output(dest_group_id: Optional[pulumi.Input[str]] = None,
                          location: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDestGroupResult]:
    """
    Retrieves an existing TunnelDestGroup.
    """
    ...