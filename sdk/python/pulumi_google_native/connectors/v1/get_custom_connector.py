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
    'GetCustomConnectorResult',
    'AwaitableGetCustomConnectorResult',
    'get_custom_connector',
    'get_custom_connector_output',
]

@pulumi.output_type
class GetCustomConnectorResult:
    def __init__(__self__, create_time=None, custom_connector_type=None, description=None, display_name=None, labels=None, launch_stage=None, logo=None, name=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if custom_connector_type and not isinstance(custom_connector_type, str):
            raise TypeError("Expected argument 'custom_connector_type' to be a str")
        pulumi.set(__self__, "custom_connector_type", custom_connector_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if launch_stage and not isinstance(launch_stage, str):
            raise TypeError("Expected argument 'launch_stage' to be a str")
        pulumi.set(__self__, "launch_stage", launch_stage)
        if logo and not isinstance(logo, str):
            raise TypeError("Expected argument 'logo' to be a str")
        pulumi.set(__self__, "logo", logo)
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
        Created time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="customConnectorType")
    def custom_connector_type(self) -> str:
        """
        Type of the custom connector.
        """
        return pulumi.get(self, "custom_connector_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. Display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="launchStage")
    def launch_stage(self) -> str:
        """
        Launch stage.
        """
        return pulumi.get(self, "launch_stage")

    @property
    @pulumi.getter
    def logo(self) -> str:
        """
        Optional. Logo of the resource.
        """
        return pulumi.get(self, "logo")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Identifier. Resource name of the CustomConnector. Format: projects/{project}/locations/{location}/customConnectors/{connector}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Updated time.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetCustomConnectorResult(GetCustomConnectorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCustomConnectorResult(
            create_time=self.create_time,
            custom_connector_type=self.custom_connector_type,
            description=self.description,
            display_name=self.display_name,
            labels=self.labels,
            launch_stage=self.launch_stage,
            logo=self.logo,
            name=self.name,
            update_time=self.update_time)


def get_custom_connector(custom_connector_id: Optional[str] = None,
                         project: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCustomConnectorResult:
    """
    Gets details of a single CustomConnector.
    """
    __args__ = dict()
    __args__['customConnectorId'] = custom_connector_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:connectors/v1:getCustomConnector', __args__, opts=opts, typ=GetCustomConnectorResult).value

    return AwaitableGetCustomConnectorResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        custom_connector_type=pulumi.get(__ret__, 'custom_connector_type'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        labels=pulumi.get(__ret__, 'labels'),
        launch_stage=pulumi.get(__ret__, 'launch_stage'),
        logo=pulumi.get(__ret__, 'logo'),
        name=pulumi.get(__ret__, 'name'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_custom_connector)
def get_custom_connector_output(custom_connector_id: Optional[pulumi.Input[str]] = None,
                                project: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCustomConnectorResult]:
    """
    Gets details of a single CustomConnector.
    """
    ...