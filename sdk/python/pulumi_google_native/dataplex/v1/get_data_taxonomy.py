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
    'GetDataTaxonomyResult',
    'AwaitableGetDataTaxonomyResult',
    'get_data_taxonomy',
    'get_data_taxonomy_output',
]

@pulumi.output_type
class GetDataTaxonomyResult:
    def __init__(__self__, attribute_count=None, create_time=None, description=None, display_name=None, etag=None, labels=None, name=None, uid=None, update_time=None):
        if attribute_count and not isinstance(attribute_count, int):
            raise TypeError("Expected argument 'attribute_count' to be a int")
        pulumi.set(__self__, "attribute_count", attribute_count)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="attributeCount")
    def attribute_count(self) -> int:
        """
        The number of attributes in the DataTaxonomy.
        """
        return pulumi.get(self, "attribute_count")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the DataTaxonomy was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the DataTaxonomy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. User-defined labels for the DataTaxonomy.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The relative resource name of the DataTaxonomy, of the form: projects/{project_number}/locations/{location_id}/dataTaxonomies/{data_taxonomy_id}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        System generated globally unique ID for the dataTaxonomy. This ID will be different if the DataTaxonomy is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the DataTaxonomy was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDataTaxonomyResult(GetDataTaxonomyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataTaxonomyResult(
            attribute_count=self.attribute_count,
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            etag=self.etag,
            labels=self.labels,
            name=self.name,
            uid=self.uid,
            update_time=self.update_time)


def get_data_taxonomy(data_taxonomy_id: Optional[str] = None,
                      location: Optional[str] = None,
                      project: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataTaxonomyResult:
    """
    Retrieves a DataTaxonomy resource.
    """
    __args__ = dict()
    __args__['dataTaxonomyId'] = data_taxonomy_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataplex/v1:getDataTaxonomy', __args__, opts=opts, typ=GetDataTaxonomyResult).value

    return AwaitableGetDataTaxonomyResult(
        attribute_count=pulumi.get(__ret__, 'attribute_count'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        etag=pulumi.get(__ret__, 'etag'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_data_taxonomy)
def get_data_taxonomy_output(data_taxonomy_id: Optional[pulumi.Input[str]] = None,
                             location: Optional[pulumi.Input[str]] = None,
                             project: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataTaxonomyResult]:
    """
    Retrieves a DataTaxonomy resource.
    """
    ...