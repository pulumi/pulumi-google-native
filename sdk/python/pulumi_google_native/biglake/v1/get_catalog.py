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
    'GetCatalogResult',
    'AwaitableGetCatalogResult',
    'get_catalog',
    'get_catalog_output',
]

@pulumi.output_type
class GetCatalogResult:
    def __init__(__self__, create_time=None, delete_time=None, expire_time=None, name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time of the catalog.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        The deletion time of the catalog. Only set after the catalog is deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        The time when this catalog is considered expired. Only set after the catalog is deleted.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name. Format: projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last modification time of the catalog.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetCatalogResult(GetCatalogResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCatalogResult(
            create_time=self.create_time,
            delete_time=self.delete_time,
            expire_time=self.expire_time,
            name=self.name,
            update_time=self.update_time)


def get_catalog(catalog_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCatalogResult:
    """
    Gets the catalog specified by the resource name.
    """
    __args__ = dict()
    __args__['catalogId'] = catalog_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:biglake/v1:getCatalog', __args__, opts=opts, typ=GetCatalogResult).value

    return AwaitableGetCatalogResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        delete_time=pulumi.get(__ret__, 'delete_time'),
        expire_time=pulumi.get(__ret__, 'expire_time'),
        name=pulumi.get(__ret__, 'name'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_catalog)
def get_catalog_output(catalog_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCatalogResult]:
    """
    Gets the catalog specified by the resource name.
    """
    ...
