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
    'GetStoragePoolResult',
    'AwaitableGetStoragePoolResult',
    'get_storage_pool',
    'get_storage_pool_output',
]

@pulumi.output_type
class GetStoragePoolResult:
    def __init__(__self__, creation_timestamp=None, description=None, kind=None, label_fingerprint=None, labels=None, name=None, provisioned_iops=None, resource_status=None, self_link=None, self_link_with_id=None, size_gb=None, state=None, type=None, zone=None):
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioned_iops and not isinstance(provisioned_iops, str):
            raise TypeError("Expected argument 'provisioned_iops' to be a str")
        pulumi.set(__self__, "provisioned_iops", provisioned_iops)
        if resource_status and not isinstance(resource_status, dict):
            raise TypeError("Expected argument 'resource_status' to be a dict")
        pulumi.set(__self__, "resource_status", resource_status)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if self_link_with_id and not isinstance(self_link_with_id, str):
            raise TypeError("Expected argument 'self_link_with_id' to be a str")
        pulumi.set(__self__, "self_link_with_id", self_link_with_id)
        if size_gb and not isinstance(size_gb, str):
            raise TypeError("Expected argument 'size_gb' to be a str")
        pulumi.set(__self__, "size_gb", size_gb)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource. Provide this property when you create the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#storagePool for storage pools.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for the labels being applied to this storage pool, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve a storage pool.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels to apply to this storage pool. These can be later modified by the setLabels method.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisionedIops")
    def provisioned_iops(self) -> str:
        """
        Provsioned IOPS of the storage pool.
        """
        return pulumi.get(self, "provisioned_iops")

    @property
    @pulumi.getter(name="resourceStatus")
    def resource_status(self) -> 'outputs.StoragePoolResourceStatusResponse':
        """
        Status information for the storage pool resource.
        """
        return pulumi.get(self, "resource_status")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined fully-qualified URL for this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="selfLinkWithId")
    def self_link_with_id(self) -> str:
        """
        Server-defined URL for this resource's resource id.
        """
        return pulumi.get(self, "self_link_with_id")

    @property
    @pulumi.getter(name="sizeGb")
    def size_gb(self) -> str:
        """
        Size, in GiB, of the storage pool.
        """
        return pulumi.get(self, "size_gb")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The status of storage pool creation. - CREATING: Storage pool is provisioning. storagePool. - FAILED: Storage pool creation failed. - READY: Storage pool is ready for use. - DELETING: Storage pool is deleting. 
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the storage pool
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        URL of the zone where the storage pool resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "zone")


class AwaitableGetStoragePoolResult(GetStoragePoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetStoragePoolResult(
            creation_timestamp=self.creation_timestamp,
            description=self.description,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            name=self.name,
            provisioned_iops=self.provisioned_iops,
            resource_status=self.resource_status,
            self_link=self.self_link,
            self_link_with_id=self.self_link_with_id,
            size_gb=self.size_gb,
            state=self.state,
            type=self.type,
            zone=self.zone)


def get_storage_pool(project: Optional[str] = None,
                     storage_pool: Optional[str] = None,
                     zone: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetStoragePoolResult:
    """
    Returns a specified storage pool. Gets a list of available storage pools by making a list() request.
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['storagePool'] = storage_pool
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/alpha:getStoragePool', __args__, opts=opts, typ=GetStoragePoolResult).value

    return AwaitableGetStoragePoolResult(
        creation_timestamp=pulumi.get(__ret__, 'creation_timestamp'),
        description=pulumi.get(__ret__, 'description'),
        kind=pulumi.get(__ret__, 'kind'),
        label_fingerprint=pulumi.get(__ret__, 'label_fingerprint'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        provisioned_iops=pulumi.get(__ret__, 'provisioned_iops'),
        resource_status=pulumi.get(__ret__, 'resource_status'),
        self_link=pulumi.get(__ret__, 'self_link'),
        self_link_with_id=pulumi.get(__ret__, 'self_link_with_id'),
        size_gb=pulumi.get(__ret__, 'size_gb'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_storage_pool)
def get_storage_pool_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                            storage_pool: Optional[pulumi.Input[str]] = None,
                            zone: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetStoragePoolResult]:
    """
    Returns a specified storage pool. Gets a list of available storage pools by making a list() request.
    """
    ...