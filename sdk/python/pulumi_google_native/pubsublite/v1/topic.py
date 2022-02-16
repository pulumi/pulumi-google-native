# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TopicArgs', 'Topic']

@pulumi.input_type
class TopicArgs:
    def __init__(__self__, *,
                 topic_id: pulumi.Input[str],
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_config: Optional[pulumi.Input['PartitionConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation_config: Optional[pulumi.Input['ReservationConfigArgs']] = None,
                 retention_config: Optional[pulumi.Input['RetentionConfigArgs']] = None):
        """
        The set of arguments for constructing a Topic resource.
        :param pulumi.Input[str] name: The name of the topic. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
        :param pulumi.Input['PartitionConfigArgs'] partition_config: The settings for this topic's partitions.
        :param pulumi.Input['ReservationConfigArgs'] reservation_config: The settings for this topic's Reservation usage.
        :param pulumi.Input['RetentionConfigArgs'] retention_config: The settings for this topic's message retention.
        """
        pulumi.set(__self__, "topic_id", topic_id)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if partition_config is not None:
            pulumi.set(__self__, "partition_config", partition_config)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if reservation_config is not None:
            pulumi.set(__self__, "reservation_config", reservation_config)
        if retention_config is not None:
            pulumi.set(__self__, "retention_config", retention_config)

    @property
    @pulumi.getter(name="topicId")
    def topic_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "topic_id")

    @topic_id.setter
    def topic_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_id", value)

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
        The name of the topic. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="partitionConfig")
    def partition_config(self) -> Optional[pulumi.Input['PartitionConfigArgs']]:
        """
        The settings for this topic's partitions.
        """
        return pulumi.get(self, "partition_config")

    @partition_config.setter
    def partition_config(self, value: Optional[pulumi.Input['PartitionConfigArgs']]):
        pulumi.set(self, "partition_config", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="reservationConfig")
    def reservation_config(self) -> Optional[pulumi.Input['ReservationConfigArgs']]:
        """
        The settings for this topic's Reservation usage.
        """
        return pulumi.get(self, "reservation_config")

    @reservation_config.setter
    def reservation_config(self, value: Optional[pulumi.Input['ReservationConfigArgs']]):
        pulumi.set(self, "reservation_config", value)

    @property
    @pulumi.getter(name="retentionConfig")
    def retention_config(self) -> Optional[pulumi.Input['RetentionConfigArgs']]:
        """
        The settings for this topic's message retention.
        """
        return pulumi.get(self, "retention_config")

    @retention_config.setter
    def retention_config(self, value: Optional[pulumi.Input['RetentionConfigArgs']]):
        pulumi.set(self, "retention_config", value)


class Topic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_config: Optional[pulumi.Input[pulumi.InputType['PartitionConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation_config: Optional[pulumi.Input[pulumi.InputType['ReservationConfigArgs']]] = None,
                 retention_config: Optional[pulumi.Input[pulumi.InputType['RetentionConfigArgs']]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new topic.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the topic. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
        :param pulumi.Input[pulumi.InputType['PartitionConfigArgs']] partition_config: The settings for this topic's partitions.
        :param pulumi.Input[pulumi.InputType['ReservationConfigArgs']] reservation_config: The settings for this topic's Reservation usage.
        :param pulumi.Input[pulumi.InputType['RetentionConfigArgs']] retention_config: The settings for this topic's message retention.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new topic.

        :param str resource_name: The name of the resource.
        :param TopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 partition_config: Optional[pulumi.Input[pulumi.InputType['PartitionConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation_config: Optional[pulumi.Input[pulumi.InputType['ReservationConfigArgs']]] = None,
                 retention_config: Optional[pulumi.Input[pulumi.InputType['RetentionConfigArgs']]] = None,
                 topic_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = TopicArgs.__new__(TopicArgs)

            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            __props__.__dict__["partition_config"] = partition_config
            __props__.__dict__["project"] = project
            __props__.__dict__["reservation_config"] = reservation_config
            __props__.__dict__["retention_config"] = retention_config
            if topic_id is None and not opts.urn:
                raise TypeError("Missing required property 'topic_id'")
            __props__.__dict__["topic_id"] = topic_id
        super(Topic, __self__).__init__(
            'google-native:pubsublite/v1:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TopicArgs.__new__(TopicArgs)

        __props__.__dict__["name"] = None
        __props__.__dict__["partition_config"] = None
        __props__.__dict__["reservation_config"] = None
        __props__.__dict__["retention_config"] = None
        return Topic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the topic. Structured like: projects/{project_number}/locations/{location}/topics/{topic_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="partitionConfig")
    def partition_config(self) -> pulumi.Output['outputs.PartitionConfigResponse']:
        """
        The settings for this topic's partitions.
        """
        return pulumi.get(self, "partition_config")

    @property
    @pulumi.getter(name="reservationConfig")
    def reservation_config(self) -> pulumi.Output['outputs.ReservationConfigResponse']:
        """
        The settings for this topic's Reservation usage.
        """
        return pulumi.get(self, "reservation_config")

    @property
    @pulumi.getter(name="retentionConfig")
    def retention_config(self) -> pulumi.Output['outputs.RetentionConfigResponse']:
        """
        The settings for this topic's message retention.
        """
        return pulumi.get(self, "retention_config")
