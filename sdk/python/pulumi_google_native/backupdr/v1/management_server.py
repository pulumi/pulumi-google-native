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

__all__ = ['ManagementServerArgs', 'ManagementServer']

@pulumi.input_type
class ManagementServerArgs:
    def __init__(__self__, *,
                 management_server_id: pulumi.Input[str],
                 networks: pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]],
                 type: pulumi.Input['ManagementServerType'],
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ManagementServer resource.
        :param pulumi.Input[str] management_server_id: Required. The name of the management server to create. The name must be unique for the specified project and location.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]] networks: VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
        :param pulumi.Input['ManagementServerType'] type: The type of the ManagementServer resource.
        :param pulumi.Input[str] description: Optional. The description of the ManagementServer instance (2048 characters or less).
        :param pulumi.Input[str] etag: Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
        :param pulumi.Input[str] request_id: Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        pulumi.set(__self__, "management_server_id", management_server_id)
        pulumi.set(__self__, "networks", networks)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="managementServerId")
    def management_server_id(self) -> pulumi.Input[str]:
        """
        Required. The name of the management server to create. The name must be unique for the specified project and location.
        """
        return pulumi.get(self, "management_server_id")

    @management_server_id.setter
    def management_server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "management_server_id", value)

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]]:
        """
        VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
        """
        return pulumi.get(self, "networks")

    @networks.setter
    def networks(self, value: pulumi.Input[Sequence[pulumi.Input['NetworkConfigArgs']]]):
        pulumi.set(self, "networks", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['ManagementServerType']:
        """
        The type of the ManagementServer resource.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['ManagementServerType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The description of the ManagementServer instance (2048 characters or less).
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
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
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class ManagementServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_server_id: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['ManagementServerType']] = None,
                 __props__=None):
        """
        Creates a new ManagementServer in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. The description of the ManagementServer instance (2048 characters or less).
        :param pulumi.Input[str] etag: Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
        :param pulumi.Input[str] management_server_id: Required. The name of the management server to create. The name must be unique for the specified project and location.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]]] networks: VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
        :param pulumi.Input[str] request_id: Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        :param pulumi.Input['ManagementServerType'] type: The type of the ManagementServer resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ManagementServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new ManagementServer in a given project and location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ManagementServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ManagementServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 management_server_id: Optional[pulumi.Input[str]] = None,
                 networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkConfigArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input['ManagementServerType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ManagementServerArgs.__new__(ManagementServerArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            if management_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'management_server_id'")
            __props__.__dict__["management_server_id"] = management_server_id
            if networks is None and not opts.urn:
                raise TypeError("Missing required property 'networks'")
            __props__.__dict__["networks"] = networks
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["management_uri"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["oauth2_client_id"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["workforce_identity_based_management_uri"] = None
            __props__.__dict__["workforce_identity_based_oauth2_client_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "management_server_id", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ManagementServer, __self__).__init__(
            'google-native:backupdr/v1:ManagementServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ManagementServer':
        """
        Get an existing ManagementServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ManagementServerArgs.__new__(ManagementServerArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["management_server_id"] = None
        __props__.__dict__["management_uri"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["networks"] = None
        __props__.__dict__["oauth2_client_id"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["workforce_identity_based_management_uri"] = None
        __props__.__dict__["workforce_identity_based_oauth2_client_id"] = None
        return ManagementServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. The description of the ManagementServer instance (2048 characters or less).
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Optional. Server specified ETag for the ManagementServer resource to prevent simultaneous updates from overwiting each other.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels to represent user provided metadata. Labels currently defined: 1. migrate_from_go= If set to true, the MS is created in migration ready mode.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managementServerId")
    def management_server_id(self) -> pulumi.Output[str]:
        """
        Required. The name of the management server to create. The name must be unique for the specified project and location.
        """
        return pulumi.get(self, "management_server_id")

    @property
    @pulumi.getter(name="managementUri")
    def management_uri(self) -> pulumi.Output['outputs.ManagementURIResponse']:
        """
        The hostname or ip address of the exposed AGM endpoints, used by clients to connect to AGM/RD graphical user interface and APIs.
        """
        return pulumi.get(self, "management_uri")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def networks(self) -> pulumi.Output[Sequence['outputs.NetworkConfigResponse']]:
        """
        VPC networks to which the ManagementServer instance is connected. For this version, only a single network is supported.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="oauth2ClientId")
    def oauth2_client_id(self) -> pulumi.Output[str]:
        """
        The OAuth 2.0 client id is required to make API calls to the BackupDR instance API of this ManagementServer. This is the value that should be provided in the ‘aud’ field of the OIDC ID Token (see openid specification https://openid.net/specs/openid-connect-core-1_0.html#IDToken).
        """
        return pulumi.get(self, "oauth2_client_id")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The ManagementServer state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the ManagementServer resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the instance was updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="workforceIdentityBasedManagementUri")
    def workforce_identity_based_management_uri(self) -> pulumi.Output['outputs.WorkforceIdentityBasedManagementURIResponse']:
        """
        The hostnames of the exposed AGM endpoints for both types of user i.e. 1p and 3p, used to connect AGM/RM UI.
        """
        return pulumi.get(self, "workforce_identity_based_management_uri")

    @property
    @pulumi.getter(name="workforceIdentityBasedOauth2ClientId")
    def workforce_identity_based_oauth2_client_id(self) -> pulumi.Output['outputs.WorkforceIdentityBasedOAuth2ClientIDResponse']:
        """
        The OAuth client IDs for both types of user i.e. 1p and 3p.
        """
        return pulumi.get(self, "workforce_identity_based_oauth2_client_id")
