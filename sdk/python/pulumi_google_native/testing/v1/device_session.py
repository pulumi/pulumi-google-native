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
from ._inputs import *

__all__ = ['DeviceSessionArgs', 'DeviceSession']

@pulumi.input_type
class DeviceSessionArgs:
    def __init__(__self__, *,
                 android_device: pulumi.Input['AndroidDeviceArgs'],
                 expire_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DeviceSession resource.
        :param pulumi.Input['AndroidDeviceArgs'] android_device: The requested device
        :param pulumi.Input[str] expire_time: Optional. If the device is still in use at this time, any connections will be ended and the SessionState will transition from ACTIVE to FINISHED.
        :param pulumi.Input[str] name: Optional. Name of the DeviceSession, e.g. "projects/{project_id}/deviceSessions/{session_id}"
        :param pulumi.Input[str] ttl: Optional. The amount of time that a device will be initially allocated for. This can eventually be extended with the UpdateDeviceSession RPC. Default: 30 minutes.
        """
        pulumi.set(__self__, "android_device", android_device)
        if expire_time is not None:
            pulumi.set(__self__, "expire_time", expire_time)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="androidDevice")
    def android_device(self) -> pulumi.Input['AndroidDeviceArgs']:
        """
        The requested device
        """
        return pulumi.get(self, "android_device")

    @android_device.setter
    def android_device(self, value: pulumi.Input['AndroidDeviceArgs']):
        pulumi.set(self, "android_device", value)

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. If the device is still in use at this time, any connections will be ended and the SessionState will transition from ACTIVE to FINISHED.
        """
        return pulumi.get(self, "expire_time")

    @expire_time.setter
    def expire_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expire_time", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Name of the DeviceSession, e.g. "projects/{project_id}/deviceSessions/{session_id}"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The amount of time that a device will be initially allocated for. This can eventually be extended with the UpdateDeviceSession RPC. Default: 30 minutes.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


class DeviceSession(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 android_device: Optional[pulumi.Input[pulumi.InputType['AndroidDeviceArgs']]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        POST /v1/projects/{project_id}/deviceSessions
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AndroidDeviceArgs']] android_device: The requested device
        :param pulumi.Input[str] expire_time: Optional. If the device is still in use at this time, any connections will be ended and the SessionState will transition from ACTIVE to FINISHED.
        :param pulumi.Input[str] name: Optional. Name of the DeviceSession, e.g. "projects/{project_id}/deviceSessions/{session_id}"
        :param pulumi.Input[str] ttl: Optional. The amount of time that a device will be initially allocated for. This can eventually be extended with the UpdateDeviceSession RPC. Default: 30 minutes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceSessionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        POST /v1/projects/{project_id}/deviceSessions
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param DeviceSessionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceSessionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 android_device: Optional[pulumi.Input[pulumi.InputType['AndroidDeviceArgs']]] = None,
                 expire_time: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceSessionArgs.__new__(DeviceSessionArgs)

            if android_device is None and not opts.urn:
                raise TypeError("Missing required property 'android_device'")
            __props__.__dict__["android_device"] = android_device
            __props__.__dict__["expire_time"] = expire_time
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["active_start_time"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["display_name"] = None
            __props__.__dict__["inactivity_timeout"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["state_histories"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DeviceSession, __self__).__init__(
            'google-native:testing/v1:DeviceSession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DeviceSession':
        """
        Get an existing DeviceSession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeviceSessionArgs.__new__(DeviceSessionArgs)

        __props__.__dict__["active_start_time"] = None
        __props__.__dict__["android_device"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["inactivity_timeout"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["state_histories"] = None
        __props__.__dict__["ttl"] = None
        return DeviceSession(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activeStartTime")
    def active_start_time(self) -> pulumi.Output[str]:
        """
        The timestamp that the session first became ACTIVE.
        """
        return pulumi.get(self, "active_start_time")

    @property
    @pulumi.getter(name="androidDevice")
    def android_device(self) -> pulumi.Output['outputs.AndroidDeviceResponse']:
        """
        The requested device
        """
        return pulumi.get(self, "android_device")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time that the Session was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The title of the DeviceSession to be presented in the UI.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        Optional. If the device is still in use at this time, any connections will be ended and the SessionState will transition from ACTIVE to FINISHED.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="inactivityTimeout")
    def inactivity_timeout(self) -> pulumi.Output[str]:
        """
        The interval of time that this device must be interacted with before it transitions from ACTIVE to TIMEOUT_INACTIVITY.
        """
        return pulumi.get(self, "inactivity_timeout")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Optional. Name of the DeviceSession, e.g. "projects/{project_id}/deviceSessions/{session_id}"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of the DeviceSession.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateHistories")
    def state_histories(self) -> pulumi.Output[Sequence['outputs.SessionStateEventResponse']]:
        """
        The historical state transitions of the session_state message including the current session state.
        """
        return pulumi.get(self, "state_histories")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[str]:
        """
        Optional. The amount of time that a device will be initially allocated for. This can eventually be extended with the UpdateDeviceSession RPC. Default: 30 minutes.
        """
        return pulumi.get(self, "ttl")
