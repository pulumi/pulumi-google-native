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
    'GetAuthConfigResult',
    'AwaitableGetAuthConfigResult',
    'get_auth_config',
    'get_auth_config_output',
]

@pulumi.output_type
class GetAuthConfigResult:
    def __init__(__self__, certificate_id=None, create_time=None, creator_email=None, credential_type=None, decrypted_credential=None, description=None, display_name=None, encrypted_credential=None, expiry_notification_duration=None, last_modifier_email=None, name=None, override_valid_time=None, reason=None, state=None, update_time=None, valid_time=None, visibility=None):
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if creator_email and not isinstance(creator_email, str):
            raise TypeError("Expected argument 'creator_email' to be a str")
        pulumi.set(__self__, "creator_email", creator_email)
        if credential_type and not isinstance(credential_type, str):
            raise TypeError("Expected argument 'credential_type' to be a str")
        pulumi.set(__self__, "credential_type", credential_type)
        if decrypted_credential and not isinstance(decrypted_credential, dict):
            raise TypeError("Expected argument 'decrypted_credential' to be a dict")
        pulumi.set(__self__, "decrypted_credential", decrypted_credential)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if encrypted_credential and not isinstance(encrypted_credential, str):
            raise TypeError("Expected argument 'encrypted_credential' to be a str")
        pulumi.set(__self__, "encrypted_credential", encrypted_credential)
        if expiry_notification_duration and not isinstance(expiry_notification_duration, list):
            raise TypeError("Expected argument 'expiry_notification_duration' to be a list")
        pulumi.set(__self__, "expiry_notification_duration", expiry_notification_duration)
        if last_modifier_email and not isinstance(last_modifier_email, str):
            raise TypeError("Expected argument 'last_modifier_email' to be a str")
        pulumi.set(__self__, "last_modifier_email", last_modifier_email)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if override_valid_time and not isinstance(override_valid_time, str):
            raise TypeError("Expected argument 'override_valid_time' to be a str")
        pulumi.set(__self__, "override_valid_time", override_valid_time)
        if reason and not isinstance(reason, str):
            raise TypeError("Expected argument 'reason' to be a str")
        pulumi.set(__self__, "reason", reason)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if valid_time and not isinstance(valid_time, str):
            raise TypeError("Expected argument 'valid_time' to be a str")
        pulumi.set(__self__, "valid_time", valid_time)
        if visibility and not isinstance(visibility, str):
            raise TypeError("Expected argument 'visibility' to be a str")
        pulumi.set(__self__, "visibility", visibility)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        """
        Certificate id for client certificate
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp when the auth config is created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="creatorEmail")
    def creator_email(self) -> str:
        """
        The creator's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        """
        return pulumi.get(self, "creator_email")

    @property
    @pulumi.getter(name="credentialType")
    def credential_type(self) -> str:
        """
        Credential type of the encrypted credential.
        """
        return pulumi.get(self, "credential_type")

    @property
    @pulumi.getter(name="decryptedCredential")
    def decrypted_credential(self) -> 'outputs.GoogleCloudIntegrationsV1alphaCredentialResponse':
        """
        Raw auth credentials.
        """
        return pulumi.get(self, "decrypted_credential")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the auth config.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The name of the auth config.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="encryptedCredential")
    def encrypted_credential(self) -> str:
        """
        Auth credential encrypted by Cloud KMS. Can be decrypted as Credential with proper KMS key.
        """
        return pulumi.get(self, "encrypted_credential")

    @property
    @pulumi.getter(name="expiryNotificationDuration")
    def expiry_notification_duration(self) -> Sequence[str]:
        """
        User can define the time to receive notification after which the auth config becomes invalid. Support up to 30 days. Support granularity in hours.
        """
        return pulumi.get(self, "expiry_notification_duration")

    @property
    @pulumi.getter(name="lastModifierEmail")
    def last_modifier_email(self) -> str:
        """
        The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
        """
        return pulumi.get(self, "last_modifier_email")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the SFDC instance projects/{project}/locations/{location}/authConfigs/{authConfig}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overrideValidTime")
    def override_valid_time(self) -> str:
        """
        User provided expiry time to override. For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
        """
        return pulumi.get(self, "override_valid_time")

    @property
    @pulumi.getter
    def reason(self) -> str:
        """
        The reason / details of the current status.
        """
        return pulumi.get(self, "reason")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The status of the auth config.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp when the auth config is modified.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="validTime")
    def valid_time(self) -> str:
        """
        The time until the auth config is valid. Empty or max value is considered the auth config won't expire.
        """
        return pulumi.get(self, "valid_time")

    @property
    @pulumi.getter
    def visibility(self) -> str:
        """
        The visibility of the auth config.
        """
        return pulumi.get(self, "visibility")


class AwaitableGetAuthConfigResult(GetAuthConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthConfigResult(
            certificate_id=self.certificate_id,
            create_time=self.create_time,
            creator_email=self.creator_email,
            credential_type=self.credential_type,
            decrypted_credential=self.decrypted_credential,
            description=self.description,
            display_name=self.display_name,
            encrypted_credential=self.encrypted_credential,
            expiry_notification_duration=self.expiry_notification_duration,
            last_modifier_email=self.last_modifier_email,
            name=self.name,
            override_valid_time=self.override_valid_time,
            reason=self.reason,
            state=self.state,
            update_time=self.update_time,
            valid_time=self.valid_time,
            visibility=self.visibility)


def get_auth_config(auth_config_id: Optional[str] = None,
                    location: Optional[str] = None,
                    product_id: Optional[str] = None,
                    project: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthConfigResult:
    """
    Gets a complete auth config. If the auth config doesn't exist, Code.NOT_FOUND exception will be thrown. Returns the decrypted auth config.
    """
    __args__ = dict()
    __args__['authConfigId'] = auth_config_id
    __args__['location'] = location
    __args__['productId'] = product_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:integrations/v1alpha:getAuthConfig', __args__, opts=opts, typ=GetAuthConfigResult).value

    return AwaitableGetAuthConfigResult(
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        create_time=pulumi.get(__ret__, 'create_time'),
        creator_email=pulumi.get(__ret__, 'creator_email'),
        credential_type=pulumi.get(__ret__, 'credential_type'),
        decrypted_credential=pulumi.get(__ret__, 'decrypted_credential'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        encrypted_credential=pulumi.get(__ret__, 'encrypted_credential'),
        expiry_notification_duration=pulumi.get(__ret__, 'expiry_notification_duration'),
        last_modifier_email=pulumi.get(__ret__, 'last_modifier_email'),
        name=pulumi.get(__ret__, 'name'),
        override_valid_time=pulumi.get(__ret__, 'override_valid_time'),
        reason=pulumi.get(__ret__, 'reason'),
        state=pulumi.get(__ret__, 'state'),
        update_time=pulumi.get(__ret__, 'update_time'),
        valid_time=pulumi.get(__ret__, 'valid_time'),
        visibility=pulumi.get(__ret__, 'visibility'))


@_utilities.lift_output_func(get_auth_config)
def get_auth_config_output(auth_config_id: Optional[pulumi.Input[str]] = None,
                           location: Optional[pulumi.Input[str]] = None,
                           product_id: Optional[pulumi.Input[str]] = None,
                           project: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAuthConfigResult]:
    """
    Gets a complete auth config. If the auth config doesn't exist, Code.NOT_FOUND exception will be thrown. Returns the decrypted auth config.
    """
    ...