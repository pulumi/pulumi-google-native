# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetConnectionProfileResult',
    'AwaitableGetConnectionProfileResult',
    'get_connection_profile',
    'get_connection_profile_output',
]

@pulumi.output_type
class GetConnectionProfileResult:
    def __init__(__self__, create_time=None, display_name=None, forward_ssh_connectivity=None, gcs_profile=None, labels=None, mysql_profile=None, name=None, no_connectivity=None, oracle_profile=None, private_connectivity=None, static_service_ip_connectivity=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if forward_ssh_connectivity and not isinstance(forward_ssh_connectivity, dict):
            raise TypeError("Expected argument 'forward_ssh_connectivity' to be a dict")
        pulumi.set(__self__, "forward_ssh_connectivity", forward_ssh_connectivity)
        if gcs_profile and not isinstance(gcs_profile, dict):
            raise TypeError("Expected argument 'gcs_profile' to be a dict")
        pulumi.set(__self__, "gcs_profile", gcs_profile)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if mysql_profile and not isinstance(mysql_profile, dict):
            raise TypeError("Expected argument 'mysql_profile' to be a dict")
        pulumi.set(__self__, "mysql_profile", mysql_profile)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if no_connectivity and not isinstance(no_connectivity, dict):
            raise TypeError("Expected argument 'no_connectivity' to be a dict")
        pulumi.set(__self__, "no_connectivity", no_connectivity)
        if oracle_profile and not isinstance(oracle_profile, dict):
            raise TypeError("Expected argument 'oracle_profile' to be a dict")
        pulumi.set(__self__, "oracle_profile", oracle_profile)
        if private_connectivity and not isinstance(private_connectivity, dict):
            raise TypeError("Expected argument 'private_connectivity' to be a dict")
        pulumi.set(__self__, "private_connectivity", private_connectivity)
        if static_service_ip_connectivity and not isinstance(static_service_ip_connectivity, dict):
            raise TypeError("Expected argument 'static_service_ip_connectivity' to be a dict")
        pulumi.set(__self__, "static_service_ip_connectivity", static_service_ip_connectivity)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The create time of the resource.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="forwardSshConnectivity")
    def forward_ssh_connectivity(self) -> 'outputs.ForwardSshTunnelConnectivityResponse':
        """
        Forward SSH tunnel connectivity.
        """
        return pulumi.get(self, "forward_ssh_connectivity")

    @property
    @pulumi.getter(name="gcsProfile")
    def gcs_profile(self) -> 'outputs.GcsProfileResponse':
        """
        Cloud Storage ConnectionProfile configuration.
        """
        return pulumi.get(self, "gcs_profile")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="mysqlProfile")
    def mysql_profile(self) -> 'outputs.MysqlProfileResponse':
        """
        MySQL ConnectionProfile configuration.
        """
        return pulumi.get(self, "mysql_profile")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noConnectivity")
    def no_connectivity(self) -> 'outputs.NoConnectivitySettingsResponse':
        """
        No connectivity option chosen.
        """
        return pulumi.get(self, "no_connectivity")

    @property
    @pulumi.getter(name="oracleProfile")
    def oracle_profile(self) -> 'outputs.OracleProfileResponse':
        """
        Oracle ConnectionProfile configuration.
        """
        return pulumi.get(self, "oracle_profile")

    @property
    @pulumi.getter(name="privateConnectivity")
    def private_connectivity(self) -> 'outputs.PrivateConnectivityResponse':
        """
        Private connectivity.
        """
        return pulumi.get(self, "private_connectivity")

    @property
    @pulumi.getter(name="staticServiceIpConnectivity")
    def static_service_ip_connectivity(self) -> 'outputs.StaticServiceIpConnectivityResponse':
        """
        Static Service IP connectivity.
        """
        return pulumi.get(self, "static_service_ip_connectivity")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The update time of the resource.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetConnectionProfileResult(GetConnectionProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectionProfileResult(
            create_time=self.create_time,
            display_name=self.display_name,
            forward_ssh_connectivity=self.forward_ssh_connectivity,
            gcs_profile=self.gcs_profile,
            labels=self.labels,
            mysql_profile=self.mysql_profile,
            name=self.name,
            no_connectivity=self.no_connectivity,
            oracle_profile=self.oracle_profile,
            private_connectivity=self.private_connectivity,
            static_service_ip_connectivity=self.static_service_ip_connectivity,
            update_time=self.update_time)


def get_connection_profile(connection_profile_id: Optional[str] = None,
                           location: Optional[str] = None,
                           project: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectionProfileResult:
    """
    Use this method to get details about a connection profile.
    """
    __args__ = dict()
    __args__['connectionProfileId'] = connection_profile_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:datastream/v1alpha1:getConnectionProfile', __args__, opts=opts, typ=GetConnectionProfileResult).value

    return AwaitableGetConnectionProfileResult(
        create_time=__ret__.create_time,
        display_name=__ret__.display_name,
        forward_ssh_connectivity=__ret__.forward_ssh_connectivity,
        gcs_profile=__ret__.gcs_profile,
        labels=__ret__.labels,
        mysql_profile=__ret__.mysql_profile,
        name=__ret__.name,
        no_connectivity=__ret__.no_connectivity,
        oracle_profile=__ret__.oracle_profile,
        private_connectivity=__ret__.private_connectivity,
        static_service_ip_connectivity=__ret__.static_service_ip_connectivity,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_connection_profile)
def get_connection_profile_output(connection_profile_id: Optional[pulumi.Input[str]] = None,
                                  location: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectionProfileResult]:
    """
    Use this method to get details about a connection profile.
    """
    ...