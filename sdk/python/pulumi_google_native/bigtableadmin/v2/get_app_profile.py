# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetAppProfileResult',
    'AwaitableGetAppProfileResult',
    'get_app_profile',
    'get_app_profile_output',
]

@pulumi.output_type
class GetAppProfileResult:
    def __init__(__self__, description=None, etag=None, multi_cluster_routing_use_any=None, name=None, priority=None, single_cluster_routing=None, standard_isolation=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if multi_cluster_routing_use_any and not isinstance(multi_cluster_routing_use_any, dict):
            raise TypeError("Expected argument 'multi_cluster_routing_use_any' to be a dict")
        pulumi.set(__self__, "multi_cluster_routing_use_any", multi_cluster_routing_use_any)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if priority and not isinstance(priority, str):
            raise TypeError("Expected argument 'priority' to be a str")
        pulumi.set(__self__, "priority", priority)
        if single_cluster_routing and not isinstance(single_cluster_routing, dict):
            raise TypeError("Expected argument 'single_cluster_routing' to be a dict")
        pulumi.set(__self__, "single_cluster_routing", single_cluster_routing)
        if standard_isolation and not isinstance(standard_isolation, dict):
            raise TypeError("Expected argument 'standard_isolation' to be a dict")
        pulumi.set(__self__, "standard_isolation", standard_isolation)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Long form description of the use case for this AppProfile.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Strongly validated etag for optimistic concurrency control. Preserve the value returned from `GetAppProfile` when calling `UpdateAppProfile` to fail the request if there has been a modification in the mean time. The `update_mask` of the request need not include `etag` for this protection to apply. See [Wikipedia](https://en.wikipedia.org/wiki/HTTP_ETag) and [RFC 7232](https://tools.ietf.org/html/rfc7232#section-2.3) for more details.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="multiClusterRoutingUseAny")
    def multi_cluster_routing_use_any(self) -> 'outputs.MultiClusterRoutingUseAnyResponse':
        """
        Use a multi-cluster routing policy.
        """
        return pulumi.get(self, "multi_cluster_routing_use_any")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The unique name of the app profile. Values are of the form `projects/{project}/instances/{instance}/appProfiles/_a-zA-Z0-9*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> str:
        """
        This field has been deprecated in favor of `standard_isolation.priority`. If you set this field, `standard_isolation.priority` will be set instead. The priority of requests sent using this app profile.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="singleClusterRouting")
    def single_cluster_routing(self) -> 'outputs.SingleClusterRoutingResponse':
        """
        Use a single-cluster routing policy.
        """
        return pulumi.get(self, "single_cluster_routing")

    @property
    @pulumi.getter(name="standardIsolation")
    def standard_isolation(self) -> 'outputs.StandardIsolationResponse':
        """
        The standard options used for isolating this app profile's traffic from other use cases.
        """
        return pulumi.get(self, "standard_isolation")


class AwaitableGetAppProfileResult(GetAppProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAppProfileResult(
            description=self.description,
            etag=self.etag,
            multi_cluster_routing_use_any=self.multi_cluster_routing_use_any,
            name=self.name,
            priority=self.priority,
            single_cluster_routing=self.single_cluster_routing,
            standard_isolation=self.standard_isolation)


def get_app_profile(app_profile_id: Optional[str] = None,
                    instance_id: Optional[str] = None,
                    project: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAppProfileResult:
    """
    Gets information about an app profile.
    """
    __args__ = dict()
    __args__['appProfileId'] = app_profile_id
    __args__['instanceId'] = instance_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:bigtableadmin/v2:getAppProfile', __args__, opts=opts, typ=GetAppProfileResult).value

    return AwaitableGetAppProfileResult(
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        multi_cluster_routing_use_any=pulumi.get(__ret__, 'multi_cluster_routing_use_any'),
        name=pulumi.get(__ret__, 'name'),
        priority=pulumi.get(__ret__, 'priority'),
        single_cluster_routing=pulumi.get(__ret__, 'single_cluster_routing'),
        standard_isolation=pulumi.get(__ret__, 'standard_isolation'))


@_utilities.lift_output_func(get_app_profile)
def get_app_profile_output(app_profile_id: Optional[pulumi.Input[str]] = None,
                           instance_id: Optional[pulumi.Input[str]] = None,
                           project: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAppProfileResult]:
    """
    Gets information about an app profile.
    """
    ...
