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
    'GetSecurityProfileResult',
    'AwaitableGetSecurityProfileResult',
    'get_security_profile',
    'get_security_profile_output',
]

@pulumi.output_type
class GetSecurityProfileResult:
    def __init__(__self__, create_time=None, description=None, etag=None, labels=None, name=None, threat_prevention_profile=None, type=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if threat_prevention_profile and not isinstance(threat_prevention_profile, dict):
            raise TypeError("Expected argument 'threat_prevention_profile' to be a dict")
        pulumi.set(__self__, "threat_prevention_profile", threat_prevention_profile)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Resource creation timestamp.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. An optional description of the profile. Max length 512 characters.
        """
        return pulumi.get(self, "description")

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
        Optional. Labels as key value pairs.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. Identifier. Name of the SecurityProfile resource. It matches pattern `projects|organizations/*/locations/{location}/securityProfiles/{security_profile}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="threatPreventionProfile")
    def threat_prevention_profile(self) -> 'outputs.ThreatPreventionProfileResponse':
        """
        The threat prevention configuration for the SecurityProfile.
        """
        return pulumi.get(self, "threat_prevention_profile")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Immutable. The single ProfileType that the SecurityProfile resource configures.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Last resource update timestamp.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetSecurityProfileResult(GetSecurityProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityProfileResult(
            create_time=self.create_time,
            description=self.description,
            etag=self.etag,
            labels=self.labels,
            name=self.name,
            threat_prevention_profile=self.threat_prevention_profile,
            type=self.type,
            update_time=self.update_time)


def get_security_profile(location: Optional[str] = None,
                         organization_id: Optional[str] = None,
                         security_profile_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityProfileResult:
    """
    Gets details of a single SecurityProfile.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['organizationId'] = organization_id
    __args__['securityProfileId'] = security_profile_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:networksecurity/v1beta1:getSecurityProfile', __args__, opts=opts, typ=GetSecurityProfileResult).value

    return AwaitableGetSecurityProfileResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        threat_prevention_profile=pulumi.get(__ret__, 'threat_prevention_profile'),
        type=pulumi.get(__ret__, 'type'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_security_profile)
def get_security_profile_output(location: Optional[pulumi.Input[str]] = None,
                                organization_id: Optional[pulumi.Input[str]] = None,
                                security_profile_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSecurityProfileResult]:
    """
    Gets details of a single SecurityProfile.
    """
    ...