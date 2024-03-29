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
from ._enums import *
from ._inputs import *

__all__ = ['TenantArgs', 'Tenant']

@pulumi.input_type
class TenantArgs:
    def __init__(__self__, *,
                 allow_password_signup: Optional[pulumi.Input[bool]] = None,
                 autodelete_anonymous_users: Optional[pulumi.Input[bool]] = None,
                 client: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs']] = None,
                 disable_auth: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email_privacy_config: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs']] = None,
                 enable_anonymous_user: Optional[pulumi.Input[bool]] = None,
                 enable_email_link_signin: Optional[pulumi.Input[bool]] = None,
                 inheritance: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs']] = None,
                 mfa_config: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs']] = None,
                 monitoring: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs']] = None,
                 password_policy_config: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recaptcha_config: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs']] = None,
                 sms_region_config: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs']] = None,
                 test_phone_numbers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Tenant resource.
        :param pulumi.Input[bool] allow_password_signup: Whether to allow email/password user authentication.
        :param pulumi.Input[bool] autodelete_anonymous_users: Whether anonymous users will be auto-deleted after a period of 30 days.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs'] client: Options related to how clients making requests on behalf of a project should be configured.
        :param pulumi.Input[bool] disable_auth: Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
        :param pulumi.Input[str] display_name: Display name of the tenant.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs'] email_privacy_config: Configuration for settings related to email privacy and public visibility.
        :param pulumi.Input[bool] enable_anonymous_user: Whether to enable anonymous user authentication.
        :param pulumi.Input[bool] enable_email_link_signin: Whether to enable email link user authentication.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs'] inheritance: Specify the settings that the tenant could inherit.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs'] mfa_config: The tenant-level configuration of MFA options.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs'] monitoring: Configuration related to monitoring project activity.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs'] password_policy_config: The tenant-level password policy config
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs'] recaptcha_config: The tenant-level reCAPTCHA config.
        :param pulumi.Input['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs'] sms_region_config: Configures which regions are enabled for SMS verification code sending.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] test_phone_numbers: A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
        """
        if allow_password_signup is not None:
            pulumi.set(__self__, "allow_password_signup", allow_password_signup)
        if autodelete_anonymous_users is not None:
            pulumi.set(__self__, "autodelete_anonymous_users", autodelete_anonymous_users)
        if client is not None:
            pulumi.set(__self__, "client", client)
        if disable_auth is not None:
            pulumi.set(__self__, "disable_auth", disable_auth)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if email_privacy_config is not None:
            pulumi.set(__self__, "email_privacy_config", email_privacy_config)
        if enable_anonymous_user is not None:
            pulumi.set(__self__, "enable_anonymous_user", enable_anonymous_user)
        if enable_email_link_signin is not None:
            pulumi.set(__self__, "enable_email_link_signin", enable_email_link_signin)
        if inheritance is not None:
            pulumi.set(__self__, "inheritance", inheritance)
        if mfa_config is not None:
            pulumi.set(__self__, "mfa_config", mfa_config)
        if monitoring is not None:
            pulumi.set(__self__, "monitoring", monitoring)
        if password_policy_config is not None:
            pulumi.set(__self__, "password_policy_config", password_policy_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if recaptcha_config is not None:
            pulumi.set(__self__, "recaptcha_config", recaptcha_config)
        if sms_region_config is not None:
            pulumi.set(__self__, "sms_region_config", sms_region_config)
        if test_phone_numbers is not None:
            pulumi.set(__self__, "test_phone_numbers", test_phone_numbers)

    @property
    @pulumi.getter(name="allowPasswordSignup")
    def allow_password_signup(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow email/password user authentication.
        """
        return pulumi.get(self, "allow_password_signup")

    @allow_password_signup.setter
    def allow_password_signup(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_password_signup", value)

    @property
    @pulumi.getter(name="autodeleteAnonymousUsers")
    def autodelete_anonymous_users(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether anonymous users will be auto-deleted after a period of 30 days.
        """
        return pulumi.get(self, "autodelete_anonymous_users")

    @autodelete_anonymous_users.setter
    def autodelete_anonymous_users(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "autodelete_anonymous_users", value)

    @property
    @pulumi.getter
    def client(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs']]:
        """
        Options related to how clients making requests on behalf of a project should be configured.
        """
        return pulumi.get(self, "client")

    @client.setter
    def client(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs']]):
        pulumi.set(self, "client", value)

    @property
    @pulumi.getter(name="disableAuth")
    def disable_auth(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
        """
        return pulumi.get(self, "disable_auth")

    @disable_auth.setter
    def disable_auth(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_auth", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of the tenant.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="emailPrivacyConfig")
    def email_privacy_config(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs']]:
        """
        Configuration for settings related to email privacy and public visibility.
        """
        return pulumi.get(self, "email_privacy_config")

    @email_privacy_config.setter
    def email_privacy_config(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs']]):
        pulumi.set(self, "email_privacy_config", value)

    @property
    @pulumi.getter(name="enableAnonymousUser")
    def enable_anonymous_user(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable anonymous user authentication.
        """
        return pulumi.get(self, "enable_anonymous_user")

    @enable_anonymous_user.setter
    def enable_anonymous_user(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_anonymous_user", value)

    @property
    @pulumi.getter(name="enableEmailLinkSignin")
    def enable_email_link_signin(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable email link user authentication.
        """
        return pulumi.get(self, "enable_email_link_signin")

    @enable_email_link_signin.setter
    def enable_email_link_signin(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_email_link_signin", value)

    @property
    @pulumi.getter
    def inheritance(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs']]:
        """
        Specify the settings that the tenant could inherit.
        """
        return pulumi.get(self, "inheritance")

    @inheritance.setter
    def inheritance(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs']]):
        pulumi.set(self, "inheritance", value)

    @property
    @pulumi.getter(name="mfaConfig")
    def mfa_config(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs']]:
        """
        The tenant-level configuration of MFA options.
        """
        return pulumi.get(self, "mfa_config")

    @mfa_config.setter
    def mfa_config(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs']]):
        pulumi.set(self, "mfa_config", value)

    @property
    @pulumi.getter
    def monitoring(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs']]:
        """
        Configuration related to monitoring project activity.
        """
        return pulumi.get(self, "monitoring")

    @monitoring.setter
    def monitoring(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs']]):
        pulumi.set(self, "monitoring", value)

    @property
    @pulumi.getter(name="passwordPolicyConfig")
    def password_policy_config(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs']]:
        """
        The tenant-level password policy config
        """
        return pulumi.get(self, "password_policy_config")

    @password_policy_config.setter
    def password_policy_config(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs']]):
        pulumi.set(self, "password_policy_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="recaptchaConfig")
    def recaptcha_config(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs']]:
        """
        The tenant-level reCAPTCHA config.
        """
        return pulumi.get(self, "recaptcha_config")

    @recaptcha_config.setter
    def recaptcha_config(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs']]):
        pulumi.set(self, "recaptcha_config", value)

    @property
    @pulumi.getter(name="smsRegionConfig")
    def sms_region_config(self) -> Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs']]:
        """
        Configures which regions are enabled for SMS verification code sending.
        """
        return pulumi.get(self, "sms_region_config")

    @sms_region_config.setter
    def sms_region_config(self, value: Optional[pulumi.Input['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs']]):
        pulumi.set(self, "sms_region_config", value)

    @property
    @pulumi.getter(name="testPhoneNumbers")
    def test_phone_numbers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
        """
        return pulumi.get(self, "test_phone_numbers")

    @test_phone_numbers.setter
    def test_phone_numbers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "test_phone_numbers", value)


class Tenant(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_password_signup: Optional[pulumi.Input[bool]] = None,
                 autodelete_anonymous_users: Optional[pulumi.Input[bool]] = None,
                 client: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs']]] = None,
                 disable_auth: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email_privacy_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs']]] = None,
                 enable_anonymous_user: Optional[pulumi.Input[bool]] = None,
                 enable_email_link_signin: Optional[pulumi.Input[bool]] = None,
                 inheritance: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs']]] = None,
                 mfa_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs']]] = None,
                 monitoring: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs']]] = None,
                 password_policy_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recaptcha_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs']]] = None,
                 sms_region_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs']]] = None,
                 test_phone_numbers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a tenant. Requires write permission on the Agent project.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_password_signup: Whether to allow email/password user authentication.
        :param pulumi.Input[bool] autodelete_anonymous_users: Whether anonymous users will be auto-deleted after a period of 30 days.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs']] client: Options related to how clients making requests on behalf of a project should be configured.
        :param pulumi.Input[bool] disable_auth: Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
        :param pulumi.Input[str] display_name: Display name of the tenant.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs']] email_privacy_config: Configuration for settings related to email privacy and public visibility.
        :param pulumi.Input[bool] enable_anonymous_user: Whether to enable anonymous user authentication.
        :param pulumi.Input[bool] enable_email_link_signin: Whether to enable email link user authentication.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs']] inheritance: Specify the settings that the tenant could inherit.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs']] mfa_config: The tenant-level configuration of MFA options.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs']] monitoring: Configuration related to monitoring project activity.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs']] password_policy_config: The tenant-level password policy config
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs']] recaptcha_config: The tenant-level reCAPTCHA config.
        :param pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs']] sms_region_config: Configures which regions are enabled for SMS verification code sending.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] test_phone_numbers: A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TenantArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a tenant. Requires write permission on the Agent project.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param TenantArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TenantArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_password_signup: Optional[pulumi.Input[bool]] = None,
                 autodelete_anonymous_users: Optional[pulumi.Input[bool]] = None,
                 client: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigArgs']]] = None,
                 disable_auth: Optional[pulumi.Input[bool]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 email_privacy_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigArgs']]] = None,
                 enable_anonymous_user: Optional[pulumi.Input[bool]] = None,
                 enable_email_link_signin: Optional[pulumi.Input[bool]] = None,
                 inheritance: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2InheritanceArgs']]] = None,
                 mfa_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigArgs']]] = None,
                 monitoring: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2MonitoringConfigArgs']]] = None,
                 password_policy_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 recaptcha_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigArgs']]] = None,
                 sms_region_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigArgs']]] = None,
                 test_phone_numbers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TenantArgs.__new__(TenantArgs)

            __props__.__dict__["allow_password_signup"] = allow_password_signup
            __props__.__dict__["autodelete_anonymous_users"] = autodelete_anonymous_users
            __props__.__dict__["client"] = client
            __props__.__dict__["disable_auth"] = disable_auth
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["email_privacy_config"] = email_privacy_config
            __props__.__dict__["enable_anonymous_user"] = enable_anonymous_user
            __props__.__dict__["enable_email_link_signin"] = enable_email_link_signin
            __props__.__dict__["inheritance"] = inheritance
            __props__.__dict__["mfa_config"] = mfa_config
            __props__.__dict__["monitoring"] = monitoring
            __props__.__dict__["password_policy_config"] = password_policy_config
            __props__.__dict__["project"] = project
            __props__.__dict__["recaptcha_config"] = recaptcha_config
            __props__.__dict__["sms_region_config"] = sms_region_config
            __props__.__dict__["test_phone_numbers"] = test_phone_numbers
            __props__.__dict__["hash_config"] = None
            __props__.__dict__["name"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Tenant, __self__).__init__(
            'google-native:identitytoolkit/v2:Tenant',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Tenant':
        """
        Get an existing Tenant resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TenantArgs.__new__(TenantArgs)

        __props__.__dict__["allow_password_signup"] = None
        __props__.__dict__["autodelete_anonymous_users"] = None
        __props__.__dict__["client"] = None
        __props__.__dict__["disable_auth"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["email_privacy_config"] = None
        __props__.__dict__["enable_anonymous_user"] = None
        __props__.__dict__["enable_email_link_signin"] = None
        __props__.__dict__["hash_config"] = None
        __props__.__dict__["inheritance"] = None
        __props__.__dict__["mfa_config"] = None
        __props__.__dict__["monitoring"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["password_policy_config"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["recaptcha_config"] = None
        __props__.__dict__["sms_region_config"] = None
        __props__.__dict__["test_phone_numbers"] = None
        return Tenant(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowPasswordSignup")
    def allow_password_signup(self) -> pulumi.Output[bool]:
        """
        Whether to allow email/password user authentication.
        """
        return pulumi.get(self, "allow_password_signup")

    @property
    @pulumi.getter(name="autodeleteAnonymousUsers")
    def autodelete_anonymous_users(self) -> pulumi.Output[bool]:
        """
        Whether anonymous users will be auto-deleted after a period of 30 days.
        """
        return pulumi.get(self, "autodelete_anonymous_users")

    @property
    @pulumi.getter
    def client(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponse']:
        """
        Options related to how clients making requests on behalf of a project should be configured.
        """
        return pulumi.get(self, "client")

    @property
    @pulumi.getter(name="disableAuth")
    def disable_auth(self) -> pulumi.Output[bool]:
        """
        Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
        """
        return pulumi.get(self, "disable_auth")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Display name of the tenant.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="emailPrivacyConfig")
    def email_privacy_config(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2EmailPrivacyConfigResponse']:
        """
        Configuration for settings related to email privacy and public visibility.
        """
        return pulumi.get(self, "email_privacy_config")

    @property
    @pulumi.getter(name="enableAnonymousUser")
    def enable_anonymous_user(self) -> pulumi.Output[bool]:
        """
        Whether to enable anonymous user authentication.
        """
        return pulumi.get(self, "enable_anonymous_user")

    @property
    @pulumi.getter(name="enableEmailLinkSignin")
    def enable_email_link_signin(self) -> pulumi.Output[bool]:
        """
        Whether to enable email link user authentication.
        """
        return pulumi.get(self, "enable_email_link_signin")

    @property
    @pulumi.getter(name="hashConfig")
    def hash_config(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2HashConfigResponse']:
        """
        Hash config information of a tenant for display on Pantheon. This can only be displayed on Pantheon to avoid the sensitive information to get accidentally leaked. Only returned in GetTenant response to restrict reading of this information. Requires firebaseauth.configs.getHashConfig permission on the agent project for returning this field.
        """
        return pulumi.get(self, "hash_config")

    @property
    @pulumi.getter
    def inheritance(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2InheritanceResponse']:
        """
        Specify the settings that the tenant could inherit.
        """
        return pulumi.get(self, "inheritance")

    @property
    @pulumi.getter(name="mfaConfig")
    def mfa_config(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponse']:
        """
        The tenant-level configuration of MFA options.
        """
        return pulumi.get(self, "mfa_config")

    @property
    @pulumi.getter
    def monitoring(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponse']:
        """
        Configuration related to monitoring project activity.
        """
        return pulumi.get(self, "monitoring")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of a tenant. For example: "projects/{project-id}/tenants/{tenant-id}"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="passwordPolicyConfig")
    def password_policy_config(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2PasswordPolicyConfigResponse']:
        """
        The tenant-level password policy config
        """
        return pulumi.get(self, "password_policy_config")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="recaptchaConfig")
    def recaptcha_config(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2RecaptchaConfigResponse']:
        """
        The tenant-level reCAPTCHA config.
        """
        return pulumi.get(self, "recaptcha_config")

    @property
    @pulumi.getter(name="smsRegionConfig")
    def sms_region_config(self) -> pulumi.Output['outputs.GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponse']:
        """
        Configures which regions are enabled for SMS verification code sending.
        """
        return pulumi.get(self, "sms_region_config")

    @property
    @pulumi.getter(name="testPhoneNumbers")
    def test_phone_numbers(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
        """
        return pulumi.get(self, "test_phone_numbers")

