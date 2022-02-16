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

__all__ = ['MigrationJobArgs', 'MigrationJob']

@pulumi.input_type
class MigrationJobArgs:
    def __init__(__self__, *,
                 destination: pulumi.Input[str],
                 migration_job_id: pulumi.Input[str],
                 source: pulumi.Input[str],
                 type: pulumi.Input['MigrationJobType'],
                 destination_database: Optional[pulumi.Input['DatabaseTypeArgs']] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 dump_path: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reverse_ssh_connectivity: Optional[pulumi.Input['ReverseSshConnectivityArgs']] = None,
                 source_database: Optional[pulumi.Input['DatabaseTypeArgs']] = None,
                 state: Optional[pulumi.Input['MigrationJobState']] = None,
                 static_ip_connectivity: Optional[pulumi.Input['StaticIpConnectivityArgs']] = None,
                 vpc_peering_connectivity: Optional[pulumi.Input['VpcPeeringConnectivityArgs']] = None):
        """
        The set of arguments for constructing a MigrationJob resource.
        :param pulumi.Input[str] destination: The resource name (URI) of the destination connection profile.
        :param pulumi.Input[str] source: The resource name (URI) of the source connection profile.
        :param pulumi.Input['MigrationJobType'] type: The migration job type.
        :param pulumi.Input['DatabaseTypeArgs'] destination_database: The database engine type and provider of the destination.
        :param pulumi.Input[str] display_name: The migration job display name.
        :param pulumi.Input[str] dump_path: The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        :param pulumi.Input[str] name: The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
        :param pulumi.Input['ReverseSshConnectivityArgs'] reverse_ssh_connectivity: The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        :param pulumi.Input['DatabaseTypeArgs'] source_database: The database engine type and provider of the source.
        :param pulumi.Input['MigrationJobState'] state: The current migration job state.
        :param pulumi.Input['StaticIpConnectivityArgs'] static_ip_connectivity: static ip connectivity data (default, no additional details needed).
        :param pulumi.Input['VpcPeeringConnectivityArgs'] vpc_peering_connectivity: The details of the VPC network that the source database is located in.
        """
        pulumi.set(__self__, "destination", destination)
        pulumi.set(__self__, "migration_job_id", migration_job_id)
        pulumi.set(__self__, "source", source)
        pulumi.set(__self__, "type", type)
        if destination_database is not None:
            pulumi.set(__self__, "destination_database", destination_database)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if dump_path is not None:
            pulumi.set(__self__, "dump_path", dump_path)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if reverse_ssh_connectivity is not None:
            pulumi.set(__self__, "reverse_ssh_connectivity", reverse_ssh_connectivity)
        if source_database is not None:
            pulumi.set(__self__, "source_database", source_database)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if static_ip_connectivity is not None:
            pulumi.set(__self__, "static_ip_connectivity", static_ip_connectivity)
        if vpc_peering_connectivity is not None:
            pulumi.set(__self__, "vpc_peering_connectivity", vpc_peering_connectivity)

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Input[str]:
        """
        The resource name (URI) of the destination connection profile.
        """
        return pulumi.get(self, "destination")

    @destination.setter
    def destination(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination", value)

    @property
    @pulumi.getter(name="migrationJobId")
    def migration_job_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "migration_job_id")

    @migration_job_id.setter
    def migration_job_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "migration_job_id", value)

    @property
    @pulumi.getter
    def source(self) -> pulumi.Input[str]:
        """
        The resource name (URI) of the source connection profile.
        """
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: pulumi.Input[str]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['MigrationJobType']:
        """
        The migration job type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['MigrationJobType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="destinationDatabase")
    def destination_database(self) -> Optional[pulumi.Input['DatabaseTypeArgs']]:
        """
        The database engine type and provider of the destination.
        """
        return pulumi.get(self, "destination_database")

    @destination_database.setter
    def destination_database(self, value: Optional[pulumi.Input['DatabaseTypeArgs']]):
        pulumi.set(self, "destination_database", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The migration job display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="dumpPath")
    def dump_path(self) -> Optional[pulumi.Input[str]]:
        """
        The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
        """
        return pulumi.get(self, "dump_path")

    @dump_path.setter
    def dump_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dump_path", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

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
        The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
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
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="reverseSshConnectivity")
    def reverse_ssh_connectivity(self) -> Optional[pulumi.Input['ReverseSshConnectivityArgs']]:
        """
        The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        """
        return pulumi.get(self, "reverse_ssh_connectivity")

    @reverse_ssh_connectivity.setter
    def reverse_ssh_connectivity(self, value: Optional[pulumi.Input['ReverseSshConnectivityArgs']]):
        pulumi.set(self, "reverse_ssh_connectivity", value)

    @property
    @pulumi.getter(name="sourceDatabase")
    def source_database(self) -> Optional[pulumi.Input['DatabaseTypeArgs']]:
        """
        The database engine type and provider of the source.
        """
        return pulumi.get(self, "source_database")

    @source_database.setter
    def source_database(self, value: Optional[pulumi.Input['DatabaseTypeArgs']]):
        pulumi.set(self, "source_database", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input['MigrationJobState']]:
        """
        The current migration job state.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input['MigrationJobState']]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter(name="staticIpConnectivity")
    def static_ip_connectivity(self) -> Optional[pulumi.Input['StaticIpConnectivityArgs']]:
        """
        static ip connectivity data (default, no additional details needed).
        """
        return pulumi.get(self, "static_ip_connectivity")

    @static_ip_connectivity.setter
    def static_ip_connectivity(self, value: Optional[pulumi.Input['StaticIpConnectivityArgs']]):
        pulumi.set(self, "static_ip_connectivity", value)

    @property
    @pulumi.getter(name="vpcPeeringConnectivity")
    def vpc_peering_connectivity(self) -> Optional[pulumi.Input['VpcPeeringConnectivityArgs']]:
        """
        The details of the VPC network that the source database is located in.
        """
        return pulumi.get(self, "vpc_peering_connectivity")

    @vpc_peering_connectivity.setter
    def vpc_peering_connectivity(self, value: Optional[pulumi.Input['VpcPeeringConnectivityArgs']]):
        pulumi.set(self, "vpc_peering_connectivity", value)


class MigrationJob(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 destination_database: Optional[pulumi.Input[pulumi.InputType['DatabaseTypeArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 dump_path: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 migration_job_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reverse_ssh_connectivity: Optional[pulumi.Input[pulumi.InputType['ReverseSshConnectivityArgs']]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_database: Optional[pulumi.Input[pulumi.InputType['DatabaseTypeArgs']]] = None,
                 state: Optional[pulumi.Input['MigrationJobState']] = None,
                 static_ip_connectivity: Optional[pulumi.Input[pulumi.InputType['StaticIpConnectivityArgs']]] = None,
                 type: Optional[pulumi.Input['MigrationJobType']] = None,
                 vpc_peering_connectivity: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectivityArgs']]] = None,
                 __props__=None):
        """
        Creates a new migration job in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination: The resource name (URI) of the destination connection profile.
        :param pulumi.Input[pulumi.InputType['DatabaseTypeArgs']] destination_database: The database engine type and provider of the destination.
        :param pulumi.Input[str] display_name: The migration job display name.
        :param pulumi.Input[str] dump_path: The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        :param pulumi.Input[str] name: The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
        :param pulumi.Input[pulumi.InputType['ReverseSshConnectivityArgs']] reverse_ssh_connectivity: The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        :param pulumi.Input[str] source: The resource name (URI) of the source connection profile.
        :param pulumi.Input[pulumi.InputType['DatabaseTypeArgs']] source_database: The database engine type and provider of the source.
        :param pulumi.Input['MigrationJobState'] state: The current migration job state.
        :param pulumi.Input[pulumi.InputType['StaticIpConnectivityArgs']] static_ip_connectivity: static ip connectivity data (default, no additional details needed).
        :param pulumi.Input['MigrationJobType'] type: The migration job type.
        :param pulumi.Input[pulumi.InputType['VpcPeeringConnectivityArgs']] vpc_peering_connectivity: The details of the VPC network that the source database is located in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MigrationJobArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new migration job in a given project and location.

        :param str resource_name: The name of the resource.
        :param MigrationJobArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MigrationJobArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination: Optional[pulumi.Input[str]] = None,
                 destination_database: Optional[pulumi.Input[pulumi.InputType['DatabaseTypeArgs']]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 dump_path: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 migration_job_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 reverse_ssh_connectivity: Optional[pulumi.Input[pulumi.InputType['ReverseSshConnectivityArgs']]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 source_database: Optional[pulumi.Input[pulumi.InputType['DatabaseTypeArgs']]] = None,
                 state: Optional[pulumi.Input['MigrationJobState']] = None,
                 static_ip_connectivity: Optional[pulumi.Input[pulumi.InputType['StaticIpConnectivityArgs']]] = None,
                 type: Optional[pulumi.Input['MigrationJobType']] = None,
                 vpc_peering_connectivity: Optional[pulumi.Input[pulumi.InputType['VpcPeeringConnectivityArgs']]] = None,
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
            __props__ = MigrationJobArgs.__new__(MigrationJobArgs)

            if destination is None and not opts.urn:
                raise TypeError("Missing required property 'destination'")
            __props__.__dict__["destination"] = destination
            __props__.__dict__["destination_database"] = destination_database
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["dump_path"] = dump_path
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            if migration_job_id is None and not opts.urn:
                raise TypeError("Missing required property 'migration_job_id'")
            __props__.__dict__["migration_job_id"] = migration_job_id
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["reverse_ssh_connectivity"] = reverse_ssh_connectivity
            if source is None and not opts.urn:
                raise TypeError("Missing required property 'source'")
            __props__.__dict__["source"] = source
            __props__.__dict__["source_database"] = source_database
            __props__.__dict__["state"] = state
            __props__.__dict__["static_ip_connectivity"] = static_ip_connectivity
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["vpc_peering_connectivity"] = vpc_peering_connectivity
            __props__.__dict__["create_time"] = None
            __props__.__dict__["duration"] = None
            __props__.__dict__["end_time"] = None
            __props__.__dict__["error"] = None
            __props__.__dict__["phase"] = None
            __props__.__dict__["update_time"] = None
        super(MigrationJob, __self__).__init__(
            'google-native:datamigration/v1:MigrationJob',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MigrationJob':
        """
        Get an existing MigrationJob resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MigrationJobArgs.__new__(MigrationJobArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["destination"] = None
        __props__.__dict__["destination_database"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["dump_path"] = None
        __props__.__dict__["duration"] = None
        __props__.__dict__["end_time"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["phase"] = None
        __props__.__dict__["reverse_ssh_connectivity"] = None
        __props__.__dict__["source"] = None
        __props__.__dict__["source_database"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["static_ip_connectivity"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["vpc_peering_connectivity"] = None
        return MigrationJob(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the migration job resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def destination(self) -> pulumi.Output[str]:
        """
        The resource name (URI) of the destination connection profile.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="destinationDatabase")
    def destination_database(self) -> pulumi.Output['outputs.DatabaseTypeResponse']:
        """
        The database engine type and provider of the destination.
        """
        return pulumi.get(self, "destination_database")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The migration job display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="dumpPath")
    def dump_path(self) -> pulumi.Output[str]:
        """
        The path to the dump file in Google Cloud Storage, in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
        """
        return pulumi.get(self, "dump_path")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[str]:
        """
        The duration of the migration job (in seconds). A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        If the migration job is completed, the time when it was completed.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        The error details in case of state FAILED.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The resource labels for migration job to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name (URI) of this migration job resource, in the form of: projects/{project}/locations/{location}/migrationJobs/{migrationJob}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def phase(self) -> pulumi.Output[str]:
        """
        The current migration job phase.
        """
        return pulumi.get(self, "phase")

    @property
    @pulumi.getter(name="reverseSshConnectivity")
    def reverse_ssh_connectivity(self) -> pulumi.Output['outputs.ReverseSshConnectivityResponse']:
        """
        The details needed to communicate to the source over Reverse SSH tunnel connectivity.
        """
        return pulumi.get(self, "reverse_ssh_connectivity")

    @property
    @pulumi.getter
    def source(self) -> pulumi.Output[str]:
        """
        The resource name (URI) of the source connection profile.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="sourceDatabase")
    def source_database(self) -> pulumi.Output['outputs.DatabaseTypeResponse']:
        """
        The database engine type and provider of the source.
        """
        return pulumi.get(self, "source_database")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current migration job state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="staticIpConnectivity")
    def static_ip_connectivity(self) -> pulumi.Output['outputs.StaticIpConnectivityResponse']:
        """
        static ip connectivity data (default, no additional details needed).
        """
        return pulumi.get(self, "static_ip_connectivity")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The migration job type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp when the migration job resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpcPeeringConnectivity")
    def vpc_peering_connectivity(self) -> pulumi.Output['outputs.VpcPeeringConnectivityResponse']:
        """
        The details of the VPC network that the source database is located in.
        """
        return pulumi.get(self, "vpc_peering_connectivity")
