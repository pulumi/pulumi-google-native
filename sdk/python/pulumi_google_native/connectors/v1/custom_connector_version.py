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

__all__ = ['CustomConnectorVersionArgs', 'CustomConnectorVersion']

@pulumi.input_type
class CustomConnectorVersionArgs:
    def __init__(__self__, *,
                 auth_config: pulumi.Input['AuthConfigArgs'],
                 custom_connector_id: pulumi.Input[str],
                 custom_connector_version_id: pulumi.Input[str],
                 destination_config: pulumi.Input['DestinationConfigArgs'],
                 type: pulumi.Input['CustomConnectorVersionType'],
                 enable_backend_destination_config: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 spec_location: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomConnectorVersion resource.
        :param pulumi.Input['AuthConfigArgs'] auth_config: Configuration for establishing the authentication to the connector destination.
        :param pulumi.Input[str] custom_connector_version_id: Required. Identifier to assign to the CreateCustomConnectorVersion. Must be unique within scope of the parent resource.
        :param pulumi.Input['DestinationConfigArgs'] destination_config: Configuration of the customConnector's destination.
        :param pulumi.Input['CustomConnectorVersionType'] type: Type of the customConnector.
        :param pulumi.Input[bool] enable_backend_destination_config: Optional. Whether to enable backend destination config. This is the backend server that the connector connects to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input[str] spec_location: Optional. Location of the custom connector spec.
        """
        pulumi.set(__self__, "auth_config", auth_config)
        pulumi.set(__self__, "custom_connector_id", custom_connector_id)
        pulumi.set(__self__, "custom_connector_version_id", custom_connector_version_id)
        pulumi.set(__self__, "destination_config", destination_config)
        pulumi.set(__self__, "type", type)
        if enable_backend_destination_config is not None:
            pulumi.set(__self__, "enable_backend_destination_config", enable_backend_destination_config)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if spec_location is not None:
            pulumi.set(__self__, "spec_location", spec_location)

    @property
    @pulumi.getter(name="authConfig")
    def auth_config(self) -> pulumi.Input['AuthConfigArgs']:
        """
        Configuration for establishing the authentication to the connector destination.
        """
        return pulumi.get(self, "auth_config")

    @auth_config.setter
    def auth_config(self, value: pulumi.Input['AuthConfigArgs']):
        pulumi.set(self, "auth_config", value)

    @property
    @pulumi.getter(name="customConnectorId")
    def custom_connector_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "custom_connector_id")

    @custom_connector_id.setter
    def custom_connector_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "custom_connector_id", value)

    @property
    @pulumi.getter(name="customConnectorVersionId")
    def custom_connector_version_id(self) -> pulumi.Input[str]:
        """
        Required. Identifier to assign to the CreateCustomConnectorVersion. Must be unique within scope of the parent resource.
        """
        return pulumi.get(self, "custom_connector_version_id")

    @custom_connector_version_id.setter
    def custom_connector_version_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "custom_connector_version_id", value)

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> pulumi.Input['DestinationConfigArgs']:
        """
        Configuration of the customConnector's destination.
        """
        return pulumi.get(self, "destination_config")

    @destination_config.setter
    def destination_config(self, value: pulumi.Input['DestinationConfigArgs']):
        pulumi.set(self, "destination_config", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['CustomConnectorVersionType']:
        """
        Type of the customConnector.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['CustomConnectorVersionType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="enableBackendDestinationConfig")
    def enable_backend_destination_config(self) -> Optional[pulumi.Input[bool]]:
        """
        Optional. Whether to enable backend destination config. This is the backend server that the connector connects to.
        """
        return pulumi.get(self, "enable_backend_destination_config")

    @enable_backend_destination_config.setter
    def enable_backend_destination_config(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_backend_destination_config", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="specLocation")
    def spec_location(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Location of the custom connector spec.
        """
        return pulumi.get(self, "spec_location")

    @spec_location.setter
    def spec_location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec_location", value)


class CustomConnectorVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_config: Optional[pulumi.Input[pulumi.InputType['AuthConfigArgs']]] = None,
                 custom_connector_id: Optional[pulumi.Input[str]] = None,
                 custom_connector_version_id: Optional[pulumi.Input[str]] = None,
                 destination_config: Optional[pulumi.Input[pulumi.InputType['DestinationConfigArgs']]] = None,
                 enable_backend_destination_config: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 spec_location: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['CustomConnectorVersionType']] = None,
                 __props__=None):
        """
        Creates a new CustomConnectorVersion in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AuthConfigArgs']] auth_config: Configuration for establishing the authentication to the connector destination.
        :param pulumi.Input[str] custom_connector_version_id: Required. Identifier to assign to the CreateCustomConnectorVersion. Must be unique within scope of the parent resource.
        :param pulumi.Input[pulumi.InputType['DestinationConfigArgs']] destination_config: Configuration of the customConnector's destination.
        :param pulumi.Input[bool] enable_backend_destination_config: Optional. Whether to enable backend destination config. This is the backend server that the connector connects to.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        :param pulumi.Input[str] spec_location: Optional. Location of the custom connector spec.
        :param pulumi.Input['CustomConnectorVersionType'] type: Type of the customConnector.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomConnectorVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new CustomConnectorVersion in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param CustomConnectorVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomConnectorVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_config: Optional[pulumi.Input[pulumi.InputType['AuthConfigArgs']]] = None,
                 custom_connector_id: Optional[pulumi.Input[str]] = None,
                 custom_connector_version_id: Optional[pulumi.Input[str]] = None,
                 destination_config: Optional[pulumi.Input[pulumi.InputType['DestinationConfigArgs']]] = None,
                 enable_backend_destination_config: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 spec_location: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['CustomConnectorVersionType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomConnectorVersionArgs.__new__(CustomConnectorVersionArgs)

            if auth_config is None and not opts.urn:
                raise TypeError("Missing required property 'auth_config'")
            __props__.__dict__["auth_config"] = auth_config
            if custom_connector_id is None and not opts.urn:
                raise TypeError("Missing required property 'custom_connector_id'")
            __props__.__dict__["custom_connector_id"] = custom_connector_id
            if custom_connector_version_id is None and not opts.urn:
                raise TypeError("Missing required property 'custom_connector_version_id'")
            __props__.__dict__["custom_connector_version_id"] = custom_connector_version_id
            if destination_config is None and not opts.urn:
                raise TypeError("Missing required property 'destination_config'")
            __props__.__dict__["destination_config"] = destination_config
            __props__.__dict__["enable_backend_destination_config"] = enable_backend_destination_config
            __props__.__dict__["labels"] = labels
            __props__.__dict__["project"] = project
            __props__.__dict__["spec_location"] = spec_location
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["custom_connector_id", "custom_connector_version_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(CustomConnectorVersion, __self__).__init__(
            'google-native:connectors/v1:CustomConnectorVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CustomConnectorVersion':
        """
        Get an existing CustomConnectorVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CustomConnectorVersionArgs.__new__(CustomConnectorVersionArgs)

        __props__.__dict__["auth_config"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["custom_connector_id"] = None
        __props__.__dict__["custom_connector_version_id"] = None
        __props__.__dict__["destination_config"] = None
        __props__.__dict__["enable_backend_destination_config"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["spec_location"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["update_time"] = None
        return CustomConnectorVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authConfig")
    def auth_config(self) -> pulumi.Output['outputs.AuthConfigResponse']:
        """
        Configuration for establishing the authentication to the connector destination.
        """
        return pulumi.get(self, "auth_config")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Created time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="customConnectorId")
    def custom_connector_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "custom_connector_id")

    @property
    @pulumi.getter(name="customConnectorVersionId")
    def custom_connector_version_id(self) -> pulumi.Output[str]:
        """
        Required. Identifier to assign to the CreateCustomConnectorVersion. Must be unique within scope of the parent resource.
        """
        return pulumi.get(self, "custom_connector_version_id")

    @property
    @pulumi.getter(name="destinationConfig")
    def destination_config(self) -> pulumi.Output['outputs.DestinationConfigResponse']:
        """
        Configuration of the customConnector's destination.
        """
        return pulumi.get(self, "destination_config")

    @property
    @pulumi.getter(name="enableBackendDestinationConfig")
    def enable_backend_destination_config(self) -> pulumi.Output[bool]:
        """
        Optional. Whether to enable backend destination config. This is the backend server that the connector connects to.
        """
        return pulumi.get(self, "enable_backend_destination_config")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Identifier. Resource name of the Version. Format: projects/{project}/locations/{location}/customConnectors/{custom_connector}/customConnectorVersions/{custom_connector_version}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="specLocation")
    def spec_location(self) -> pulumi.Output[str]:
        """
        Optional. Location of the custom connector spec.
        """
        return pulumi.get(self, "spec_location")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the customConnector.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Updated time.
        """
        return pulumi.get(self, "update_time")
