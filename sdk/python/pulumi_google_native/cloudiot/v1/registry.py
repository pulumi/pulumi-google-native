# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['RegistryArgs', 'Registry']

@pulumi.input_type
class RegistryArgs:
    def __init__(__self__, *,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input['EventNotificationConfigArgs']]]] = None,
                 http_config: Optional[pulumi.Input['HttpConfigArgs']] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input['RegistryLogLevel']] = None,
                 mqtt_config: Optional[pulumi.Input['MqttConfigArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input['StateNotificationConfigArgs']] = None):
        """
        The set of arguments for constructing a Registry resource.
        :param pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]] credentials: The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
        :param pulumi.Input[Sequence[pulumi.Input['EventNotificationConfigArgs']]] event_notification_configs: The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
        :param pulumi.Input['HttpConfigArgs'] http_config: The DeviceService (HTTP) configuration for this device registry.
        :param pulumi.Input[str] id: The identifier of this device registry. For example, `myRegistry`.
        :param pulumi.Input['RegistryLogLevel'] log_level: **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
        :param pulumi.Input['MqttConfigArgs'] mqtt_config: The MQTT configuration for this device registry.
        :param pulumi.Input[str] name: The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
        :param pulumi.Input['StateNotificationConfigArgs'] state_notification_config: The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
        """
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if event_notification_configs is not None:
            pulumi.set(__self__, "event_notification_configs", event_notification_configs)
        if http_config is not None:
            pulumi.set(__self__, "http_config", http_config)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if log_level is not None:
            pulumi.set(__self__, "log_level", log_level)
        if mqtt_config is not None:
            pulumi.set(__self__, "mqtt_config", mqtt_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if state_notification_config is not None:
            pulumi.set(__self__, "state_notification_config", state_notification_config)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]]:
        """
        The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RegistryCredentialArgs']]]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="eventNotificationConfigs")
    def event_notification_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EventNotificationConfigArgs']]]]:
        """
        The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
        """
        return pulumi.get(self, "event_notification_configs")

    @event_notification_configs.setter
    def event_notification_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EventNotificationConfigArgs']]]]):
        pulumi.set(self, "event_notification_configs", value)

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> Optional[pulumi.Input['HttpConfigArgs']]:
        """
        The DeviceService (HTTP) configuration for this device registry.
        """
        return pulumi.get(self, "http_config")

    @http_config.setter
    def http_config(self, value: Optional[pulumi.Input['HttpConfigArgs']]):
        pulumi.set(self, "http_config", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of this device registry. For example, `myRegistry`.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> Optional[pulumi.Input['RegistryLogLevel']]:
        """
        **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
        """
        return pulumi.get(self, "log_level")

    @log_level.setter
    def log_level(self, value: Optional[pulumi.Input['RegistryLogLevel']]):
        pulumi.set(self, "log_level", value)

    @property
    @pulumi.getter(name="mqttConfig")
    def mqtt_config(self) -> Optional[pulumi.Input['MqttConfigArgs']]:
        """
        The MQTT configuration for this device registry.
        """
        return pulumi.get(self, "mqtt_config")

    @mqtt_config.setter
    def mqtt_config(self, value: Optional[pulumi.Input['MqttConfigArgs']]):
        pulumi.set(self, "mqtt_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
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
    @pulumi.getter(name="stateNotificationConfig")
    def state_notification_config(self) -> Optional[pulumi.Input['StateNotificationConfigArgs']]:
        """
        The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
        """
        return pulumi.get(self, "state_notification_config")

    @state_notification_config.setter
    def state_notification_config(self, value: Optional[pulumi.Input['StateNotificationConfigArgs']]):
        pulumi.set(self, "state_notification_config", value)


class Registry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventNotificationConfigArgs']]]]] = None,
                 http_config: Optional[pulumi.Input[pulumi.InputType['HttpConfigArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input['RegistryLogLevel']] = None,
                 mqtt_config: Optional[pulumi.Input[pulumi.InputType['MqttConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input[pulumi.InputType['StateNotificationConfigArgs']]] = None,
                 __props__=None):
        """
        Creates a device registry that contains devices.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]] credentials: The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventNotificationConfigArgs']]]] event_notification_configs: The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
        :param pulumi.Input[pulumi.InputType['HttpConfigArgs']] http_config: The DeviceService (HTTP) configuration for this device registry.
        :param pulumi.Input[str] id: The identifier of this device registry. For example, `myRegistry`.
        :param pulumi.Input['RegistryLogLevel'] log_level: **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
        :param pulumi.Input[pulumi.InputType['MqttConfigArgs']] mqtt_config: The MQTT configuration for this device registry.
        :param pulumi.Input[str] name: The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
        :param pulumi.Input[pulumi.InputType['StateNotificationConfigArgs']] state_notification_config: The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[RegistryArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a device registry that contains devices.

        :param str resource_name: The name of the resource.
        :param RegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegistryCredentialArgs']]]]] = None,
                 event_notification_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EventNotificationConfigArgs']]]]] = None,
                 http_config: Optional[pulumi.Input[pulumi.InputType['HttpConfigArgs']]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 log_level: Optional[pulumi.Input['RegistryLogLevel']] = None,
                 mqtt_config: Optional[pulumi.Input[pulumi.InputType['MqttConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 state_notification_config: Optional[pulumi.Input[pulumi.InputType['StateNotificationConfigArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryArgs.__new__(RegistryArgs)

            __props__.__dict__["credentials"] = credentials
            __props__.__dict__["event_notification_configs"] = event_notification_configs
            __props__.__dict__["http_config"] = http_config
            __props__.__dict__["id"] = id
            __props__.__dict__["location"] = location
            __props__.__dict__["log_level"] = log_level
            __props__.__dict__["mqtt_config"] = mqtt_config
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["state_notification_config"] = state_notification_config
        super(Registry, __self__).__init__(
            'google-native:cloudiot/v1:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Registry':
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegistryArgs.__new__(RegistryArgs)

        __props__.__dict__["credentials"] = None
        __props__.__dict__["event_notification_configs"] = None
        __props__.__dict__["http_config"] = None
        __props__.__dict__["log_level"] = None
        __props__.__dict__["mqtt_config"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state_notification_config"] = None
        return Registry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Sequence['outputs.RegistryCredentialResponse']]:
        """
        The credentials used to verify the device credentials. No more than 10 credentials can be bound to a single registry at a time. The verification process occurs at the time of device creation or update. If this field is empty, no verification is performed. Otherwise, the credentials of a newly created device or added credentials of an updated device should be signed with one of these registry credentials. Note, however, that existing devices will never be affected by modifications to this list of credentials: after a device has been successfully created in a registry, it should be able to connect even if its registry credentials are revoked, deleted, or modified.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="eventNotificationConfigs")
    def event_notification_configs(self) -> pulumi.Output[Sequence['outputs.EventNotificationConfigResponse']]:
        """
        The configuration for notification of telemetry events received from the device. All telemetry events that were successfully published by the device and acknowledged by Cloud IoT Core are guaranteed to be delivered to Cloud Pub/Sub. If multiple configurations match a message, only the first matching configuration is used. If you try to publish a device telemetry event using MQTT without specifying a Cloud Pub/Sub topic for the device's registry, the connection closes automatically. If you try to do so using an HTTP connection, an error is returned. Up to 10 configurations may be provided.
        """
        return pulumi.get(self, "event_notification_configs")

    @property
    @pulumi.getter(name="httpConfig")
    def http_config(self) -> pulumi.Output['outputs.HttpConfigResponse']:
        """
        The DeviceService (HTTP) configuration for this device registry.
        """
        return pulumi.get(self, "http_config")

    @property
    @pulumi.getter(name="logLevel")
    def log_level(self) -> pulumi.Output[str]:
        """
        **Beta Feature** The default logging verbosity for activity from devices in this registry. The verbosity level can be overridden by Device.log_level.
        """
        return pulumi.get(self, "log_level")

    @property
    @pulumi.getter(name="mqttConfig")
    def mqtt_config(self) -> pulumi.Output['outputs.MqttConfigResponse']:
        """
        The MQTT configuration for this device registry.
        """
        return pulumi.get(self, "mqtt_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource path name. For example, `projects/example-project/locations/us-central1/registries/my-registry`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="stateNotificationConfig")
    def state_notification_config(self) -> pulumi.Output['outputs.StateNotificationConfigResponse']:
        """
        The configuration for notification of new states received from the device. State updates are guaranteed to be stored in the state history, but notifications to Cloud Pub/Sub are not guaranteed. For example, if permissions are misconfigured or the specified topic doesn't exist, no notification will be published but the state will still be stored in Cloud IoT Core.
        """
        return pulumi.get(self, "state_notification_config")
