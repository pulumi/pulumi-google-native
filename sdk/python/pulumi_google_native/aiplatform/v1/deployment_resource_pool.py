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

__all__ = ['DeploymentResourcePoolArgs', 'DeploymentResourcePool']

@pulumi.input_type
class DeploymentResourcePoolArgs:
    def __init__(__self__, *,
                 dedicated_resources: pulumi.Input['GoogleCloudAiplatformV1DedicatedResourcesArgs'],
                 deployment_resource_pool_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DeploymentResourcePool resource.
        :param pulumi.Input['GoogleCloudAiplatformV1DedicatedResourcesArgs'] dedicated_resources: The underlying DedicatedResources that the DeploymentResourcePool uses.
        :param pulumi.Input[str] deployment_resource_pool_id: The ID to use for the DeploymentResourcePool, which will become the final component of the DeploymentResourcePool's resource name. The maximum length is 63 characters, and valid characters are `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
        :param pulumi.Input[str] name: Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
        """
        pulumi.set(__self__, "dedicated_resources", dedicated_resources)
        pulumi.set(__self__, "deployment_resource_pool_id", deployment_resource_pool_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="dedicatedResources")
    def dedicated_resources(self) -> pulumi.Input['GoogleCloudAiplatformV1DedicatedResourcesArgs']:
        """
        The underlying DedicatedResources that the DeploymentResourcePool uses.
        """
        return pulumi.get(self, "dedicated_resources")

    @dedicated_resources.setter
    def dedicated_resources(self, value: pulumi.Input['GoogleCloudAiplatformV1DedicatedResourcesArgs']):
        pulumi.set(self, "dedicated_resources", value)

    @property
    @pulumi.getter(name="deploymentResourcePoolId")
    def deployment_resource_pool_id(self) -> pulumi.Input[str]:
        """
        The ID to use for the DeploymentResourcePool, which will become the final component of the DeploymentResourcePool's resource name. The maximum length is 63 characters, and valid characters are `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
        """
        return pulumi.get(self, "deployment_resource_pool_id")

    @deployment_resource_pool_id.setter
    def deployment_resource_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "deployment_resource_pool_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
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


class DeploymentResourcePool(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_resources: Optional[pulumi.Input[pulumi.InputType['GoogleCloudAiplatformV1DedicatedResourcesArgs']]] = None,
                 deployment_resource_pool_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DeploymentResourcePool.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudAiplatformV1DedicatedResourcesArgs']] dedicated_resources: The underlying DedicatedResources that the DeploymentResourcePool uses.
        :param pulumi.Input[str] deployment_resource_pool_id: The ID to use for the DeploymentResourcePool, which will become the final component of the DeploymentResourcePool's resource name. The maximum length is 63 characters, and valid characters are `/^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$/`.
        :param pulumi.Input[str] name: Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeploymentResourcePoolArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DeploymentResourcePool.

        :param str resource_name: The name of the resource.
        :param DeploymentResourcePoolArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeploymentResourcePoolArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dedicated_resources: Optional[pulumi.Input[pulumi.InputType['GoogleCloudAiplatformV1DedicatedResourcesArgs']]] = None,
                 deployment_resource_pool_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeploymentResourcePoolArgs.__new__(DeploymentResourcePoolArgs)

            if dedicated_resources is None and not opts.urn:
                raise TypeError("Missing required property 'dedicated_resources'")
            __props__.__dict__["dedicated_resources"] = dedicated_resources
            if deployment_resource_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'deployment_resource_pool_id'")
            __props__.__dict__["deployment_resource_pool_id"] = deployment_resource_pool_id
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DeploymentResourcePool, __self__).__init__(
            'google-native:aiplatform/v1:DeploymentResourcePool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DeploymentResourcePool':
        """
        Get an existing DeploymentResourcePool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DeploymentResourcePoolArgs.__new__(DeploymentResourcePoolArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["dedicated_resources"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        return DeploymentResourcePool(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Timestamp when this DeploymentResourcePool was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dedicatedResources")
    def dedicated_resources(self) -> pulumi.Output['outputs.GoogleCloudAiplatformV1DedicatedResourcesResponse']:
        """
        The underlying DedicatedResources that the DeploymentResourcePool uses.
        """
        return pulumi.get(self, "dedicated_resources")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The resource name of the DeploymentResourcePool. Format: `projects/{project}/locations/{location}/deploymentResourcePools/{deployment_resource_pool}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")
