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

__all__ = ['PrivateCloudArgs', 'PrivateCloud']

@pulumi.input_type
class PrivateCloudArgs:
    def __init__(__self__, *,
                 management_cluster: pulumi.Input['ManagementClusterArgs'],
                 network_config: pulumi.Input['NetworkConfigArgs'],
                 private_cloud_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['PrivateCloudType']] = None):
        """
        The set of arguments for constructing a PrivateCloud resource.
        :param pulumi.Input['ManagementClusterArgs'] management_cluster: Input only. The management cluster for this private cloud. This field is required during creation of the private cloud to provide details for the default cluster. The following fields can't be changed after private cloud creation: `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.
        :param pulumi.Input['NetworkConfigArgs'] network_config: Network configuration of the private cloud.
        :param pulumi.Input[str] private_cloud_id: Required. The user-provided identifier of the private cloud to be created. This identifier must be unique among each `PrivateCloud` within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
        :param pulumi.Input[str] description: User-provided description for this private cloud.
        :param pulumi.Input[str] request_id: Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        :param pulumi.Input['PrivateCloudType'] type: Optional. Type of the private cloud. Defaults to STANDARD.
        """
        pulumi.set(__self__, "management_cluster", management_cluster)
        pulumi.set(__self__, "network_config", network_config)
        pulumi.set(__self__, "private_cloud_id", private_cloud_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Input['ManagementClusterArgs']:
        """
        Input only. The management cluster for this private cloud. This field is required during creation of the private cloud to provide details for the default cluster. The following fields can't be changed after private cloud creation: `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.
        """
        return pulumi.get(self, "management_cluster")

    @management_cluster.setter
    def management_cluster(self, value: pulumi.Input['ManagementClusterArgs']):
        pulumi.set(self, "management_cluster", value)

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Input['NetworkConfigArgs']:
        """
        Network configuration of the private cloud.
        """
        return pulumi.get(self, "network_config")

    @network_config.setter
    def network_config(self, value: pulumi.Input['NetworkConfigArgs']):
        pulumi.set(self, "network_config", value)

    @property
    @pulumi.getter(name="privateCloudId")
    def private_cloud_id(self) -> pulumi.Input[str]:
        """
        Required. The user-provided identifier of the private cloud to be created. This identifier must be unique among each `PrivateCloud` within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
        """
        return pulumi.get(self, "private_cloud_id")

    @private_cloud_id.setter
    def private_cloud_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_cloud_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User-provided description for this private cloud.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input['PrivateCloudType']]:
        """
        Optional. Type of the private cloud. Defaults to STANDARD.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input['PrivateCloudType']]):
        pulumi.set(self, "type", value)


class PrivateCloud(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['ManagementClusterArgs']]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]] = None,
                 private_cloud_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['PrivateCloudType']] = None,
                 __props__=None):
        """
        Creates a new `PrivateCloud` resource in a given project and location. Private clouds of type `STANDARD` and `TIME_LIMITED` are zonal resources, `STRETCHED` private clouds are regional. Creating a private cloud also creates a [management cluster](https://cloud.google.com/vmware-engine/docs/concepts-vmware-components) for that private cloud.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: User-provided description for this private cloud.
        :param pulumi.Input[pulumi.InputType['ManagementClusterArgs']] management_cluster: Input only. The management cluster for this private cloud. This field is required during creation of the private cloud to provide details for the default cluster. The following fields can't be changed after private cloud creation: `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.
        :param pulumi.Input[pulumi.InputType['NetworkConfigArgs']] network_config: Network configuration of the private cloud.
        :param pulumi.Input[str] private_cloud_id: Required. The user-provided identifier of the private cloud to be created. This identifier must be unique among each `PrivateCloud` within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
        :param pulumi.Input[str] request_id: Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        :param pulumi.Input['PrivateCloudType'] type: Optional. Type of the private cloud. Defaults to STANDARD.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PrivateCloudArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new `PrivateCloud` resource in a given project and location. Private clouds of type `STANDARD` and `TIME_LIMITED` are zonal resources, `STRETCHED` private clouds are regional. Creating a private cloud also creates a [management cluster](https://cloud.google.com/vmware-engine/docs/concepts-vmware-components) for that private cloud.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param PrivateCloudArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PrivateCloudArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_cluster: Optional[pulumi.Input[pulumi.InputType['ManagementClusterArgs']]] = None,
                 network_config: Optional[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]] = None,
                 private_cloud_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['PrivateCloudType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PrivateCloudArgs.__new__(PrivateCloudArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["location"] = location
            if management_cluster is None and not opts.urn:
                raise TypeError("Missing required property 'management_cluster'")
            __props__.__dict__["management_cluster"] = management_cluster
            if network_config is None and not opts.urn:
                raise TypeError("Missing required property 'network_config'")
            __props__.__dict__["network_config"] = network_config
            if private_cloud_id is None and not opts.urn:
                raise TypeError("Missing required property 'private_cloud_id'")
            __props__.__dict__["private_cloud_id"] = private_cloud_id
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["type"] = type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["hcx"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["nsx"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["vcenter"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "private_cloud_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(PrivateCloud, __self__).__init__(
            'google-native:vmwareengine/v1:PrivateCloud',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'PrivateCloud':
        """
        Get an existing PrivateCloud resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = PrivateCloudArgs.__new__(PrivateCloudArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["hcx"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["management_cluster"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["network_config"] = None
        __props__.__dict__["nsx"] = None
        __props__.__dict__["private_cloud_id"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["vcenter"] = None
        return PrivateCloud(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Creation time of this resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        Time when the resource was scheduled for deletion.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        User-provided description for this private cloud.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        Time when the resource will be irreversibly deleted.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def hcx(self) -> pulumi.Output['outputs.HcxResponse']:
        """
        HCX appliance.
        """
        return pulumi.get(self, "hcx")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementCluster")
    def management_cluster(self) -> pulumi.Output['outputs.ManagementClusterResponse']:
        """
        Input only. The management cluster for this private cloud. This field is required during creation of the private cloud to provide details for the default cluster. The following fields can't be changed after private cloud creation: `ManagementCluster.clusterId`, `ManagementCluster.nodeTypeId`.
        """
        return pulumi.get(self, "management_cluster")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of this private cloud. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/us-central1-a/privateClouds/my-cloud`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> pulumi.Output['outputs.NetworkConfigResponse']:
        """
        Network configuration of the private cloud.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter
    def nsx(self) -> pulumi.Output['outputs.NsxResponse']:
        """
        NSX appliance.
        """
        return pulumi.get(self, "nsx")

    @property
    @pulumi.getter(name="privateCloudId")
    def private_cloud_id(self) -> pulumi.Output[str]:
        """
        Required. The user-provided identifier of the private cloud to be created. This identifier must be unique among each `PrivateCloud` within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
        """
        return pulumi.get(self, "private_cloud_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        State of the resource. New values may be added to this enum when appropriate.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Optional. Type of the private cloud. Defaults to STANDARD.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        System-generated unique identifier for the resource.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update time of this resource.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def vcenter(self) -> pulumi.Output['outputs.VcenterResponse']:
        """
        Vcenter appliance.
        """
        return pulumi.get(self, "vcenter")
