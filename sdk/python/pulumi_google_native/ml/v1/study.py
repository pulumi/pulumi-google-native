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

__all__ = ['StudyArgs', 'Study']

@pulumi.input_type
class StudyArgs:
    def __init__(__self__, *,
                 study_config: pulumi.Input['GoogleCloudMlV1__StudyConfigArgs'],
                 study_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Study resource.
        :param pulumi.Input['GoogleCloudMlV1__StudyConfigArgs'] study_config: Configuration of the study.
        :param pulumi.Input[str] study_id: Required. The ID to use for the study, which will become the final component of the study's resource name.
        """
        pulumi.set(__self__, "study_config", study_config)
        pulumi.set(__self__, "study_id", study_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="studyConfig")
    def study_config(self) -> pulumi.Input['GoogleCloudMlV1__StudyConfigArgs']:
        """
        Configuration of the study.
        """
        return pulumi.get(self, "study_config")

    @study_config.setter
    def study_config(self, value: pulumi.Input['GoogleCloudMlV1__StudyConfigArgs']):
        pulumi.set(self, "study_config", value)

    @property
    @pulumi.getter(name="studyId")
    def study_id(self) -> pulumi.Input[str]:
        """
        Required. The ID to use for the study, which will become the final component of the study's resource name.
        """
        return pulumi.get(self, "study_id")

    @study_id.setter
    def study_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "study_id", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Study(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 study_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__StudyConfigArgs']]] = None,
                 study_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a study.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudMlV1__StudyConfigArgs']] study_config: Configuration of the study.
        :param pulumi.Input[str] study_id: Required. The ID to use for the study, which will become the final component of the study's resource name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StudyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a study.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param StudyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StudyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 study_config: Optional[pulumi.Input[pulumi.InputType['GoogleCloudMlV1__StudyConfigArgs']]] = None,
                 study_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = StudyArgs.__new__(StudyArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            if study_config is None and not opts.urn:
                raise TypeError("Missing required property 'study_config'")
            __props__.__dict__["study_config"] = study_config
            if study_id is None and not opts.urn:
                raise TypeError("Missing required property 'study_id'")
            __props__.__dict__["study_id"] = study_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["inactive_reason"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
        super(Study, __self__).__init__(
            'google-native:ml/v1:Study',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Study':
        """
        Get an existing Study resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = StudyArgs.__new__(StudyArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["inactive_reason"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["study_config"] = None
        return Study(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time at which the study was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="inactiveReason")
    def inactive_reason(self) -> pulumi.Output[str]:
        """
        A human readable reason why the Study is inactive. This should be empty if a study is ACTIVE or COMPLETED.
        """
        return pulumi.get(self, "inactive_reason")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of a study.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The detailed state of a study.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="studyConfig")
    def study_config(self) -> pulumi.Output['outputs.GoogleCloudMlV1__StudyConfigResponse']:
        """
        Configuration of the study.
        """
        return pulumi.get(self, "study_config")
