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
    'GetInstanceAttachmentResult',
    'AwaitableGetInstanceAttachmentResult',
    'get_instance_attachment',
    'get_instance_attachment_output',
]

@pulumi.output_type
class GetInstanceAttachmentResult:
    def __init__(__self__, created_at=None, environment=None, name=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        Time the attachment was created in milliseconds since epoch.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        ID of the attached environment.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        ID of the attachment.
        """
        return pulumi.get(self, "name")


class AwaitableGetInstanceAttachmentResult(GetInstanceAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceAttachmentResult(
            created_at=self.created_at,
            environment=self.environment,
            name=self.name)


def get_instance_attachment(attachment_id: Optional[str] = None,
                            instance_id: Optional[str] = None,
                            organization_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceAttachmentResult:
    """
    Gets an attachment. **Note:** Not supported for Apigee hybrid.
    """
    __args__ = dict()
    __args__['attachmentId'] = attachment_id
    __args__['instanceId'] = instance_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:apigee/v1:getInstanceAttachment', __args__, opts=opts, typ=GetInstanceAttachmentResult).value

    return AwaitableGetInstanceAttachmentResult(
        created_at=pulumi.get(__ret__, 'created_at'),
        environment=pulumi.get(__ret__, 'environment'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_instance_attachment)
def get_instance_attachment_output(attachment_id: Optional[pulumi.Input[str]] = None,
                                   instance_id: Optional[pulumi.Input[str]] = None,
                                   organization_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceAttachmentResult]:
    """
    Gets an attachment. **Note:** Not supported for Apigee hybrid.
    """
    ...
