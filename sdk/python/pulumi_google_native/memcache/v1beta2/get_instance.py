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
    'GetInstanceResult',
    'AwaitableGetInstanceResult',
    'get_instance',
    'get_instance_output',
]

@pulumi.output_type
class GetInstanceResult:
    def __init__(__self__, authorized_network=None, create_time=None, discovery_endpoint=None, display_name=None, instance_messages=None, labels=None, memcache_full_version=None, memcache_nodes=None, memcache_version=None, name=None, node_config=None, node_count=None, parameters=None, state=None, update_available=None, update_time=None, zones=None):
        if authorized_network and not isinstance(authorized_network, str):
            raise TypeError("Expected argument 'authorized_network' to be a str")
        pulumi.set(__self__, "authorized_network", authorized_network)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if discovery_endpoint and not isinstance(discovery_endpoint, str):
            raise TypeError("Expected argument 'discovery_endpoint' to be a str")
        pulumi.set(__self__, "discovery_endpoint", discovery_endpoint)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if instance_messages and not isinstance(instance_messages, list):
            raise TypeError("Expected argument 'instance_messages' to be a list")
        pulumi.set(__self__, "instance_messages", instance_messages)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if memcache_full_version and not isinstance(memcache_full_version, str):
            raise TypeError("Expected argument 'memcache_full_version' to be a str")
        pulumi.set(__self__, "memcache_full_version", memcache_full_version)
        if memcache_nodes and not isinstance(memcache_nodes, list):
            raise TypeError("Expected argument 'memcache_nodes' to be a list")
        pulumi.set(__self__, "memcache_nodes", memcache_nodes)
        if memcache_version and not isinstance(memcache_version, str):
            raise TypeError("Expected argument 'memcache_version' to be a str")
        pulumi.set(__self__, "memcache_version", memcache_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if node_config and not isinstance(node_config, dict):
            raise TypeError("Expected argument 'node_config' to be a dict")
        pulumi.set(__self__, "node_config", node_config)
        if node_count and not isinstance(node_count, int):
            raise TypeError("Expected argument 'node_count' to be a int")
        pulumi.set(__self__, "node_count", node_count)
        if parameters and not isinstance(parameters, dict):
            raise TypeError("Expected argument 'parameters' to be a dict")
        pulumi.set(__self__, "parameters", parameters)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if update_available and not isinstance(update_available, bool):
            raise TypeError("Expected argument 'update_available' to be a bool")
        pulumi.set(__self__, "update_available", update_available)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if zones and not isinstance(zones, list):
            raise TypeError("Expected argument 'zones' to be a list")
        pulumi.set(__self__, "zones", zones)

    @property
    @pulumi.getter(name="authorizedNetwork")
    def authorized_network(self) -> str:
        """
        The full name of the Google Compute Engine [network](https://cloud.google.com/vpc/docs/vpc) to which the instance is connected. If left unspecified, the `default` network will be used.
        """
        return pulumi.get(self, "authorized_network")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="discoveryEndpoint")
    def discovery_endpoint(self) -> str:
        """
        Endpoint for the Discovery API.
        """
        return pulumi.get(self, "discovery_endpoint")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        User provided name for the instance, which is only used for display purposes. Cannot be more than 80 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="instanceMessages")
    def instance_messages(self) -> Sequence['outputs.InstanceMessageResponse']:
        """
        List of messages that describe the current state of the Memcached instance.
        """
        return pulumi.get(self, "instance_messages")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="memcacheFullVersion")
    def memcache_full_version(self) -> str:
        """
        The full version of memcached server running on this instance. System automatically determines the full memcached version for an instance based on the input MemcacheVersion. The full version format will be "memcached-1.5.16".
        """
        return pulumi.get(self, "memcache_full_version")

    @property
    @pulumi.getter(name="memcacheNodes")
    def memcache_nodes(self) -> Sequence['outputs.NodeResponse']:
        """
        List of Memcached nodes. Refer to Node message for more details.
        """
        return pulumi.get(self, "memcache_nodes")

    @property
    @pulumi.getter(name="memcacheVersion")
    def memcache_version(self) -> str:
        """
        The major version of Memcached software. If not provided, latest supported version will be used. Currently the latest supported major version is `MEMCACHE_1_5`. The minor version will be automatically determined by our system based on the latest supported minor version.
        """
        return pulumi.get(self, "memcache_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the resource in this scope including project and location using the form: `projects/{project_id}/locations/{location_id}/instances/{instance_id}` Note: Memcached instances are managed and addressed at the regional level so `location_id` here refers to a Google Cloud region; however, users may choose which zones Memcached nodes should be provisioned in within an instance. Refer to zones field for more details.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nodeConfig")
    def node_config(self) -> 'outputs.NodeConfigResponse':
        """
        Configuration for Memcached nodes.
        """
        return pulumi.get(self, "node_config")

    @property
    @pulumi.getter(name="nodeCount")
    def node_count(self) -> int:
        """
        Number of nodes in the Memcached instance.
        """
        return pulumi.get(self, "node_count")

    @property
    @pulumi.getter
    def parameters(self) -> 'outputs.MemcacheParametersResponse':
        """
        User defined parameters to apply to the memcached process on each node.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of this Memcached instance.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateAvailable")
    def update_available(self) -> bool:
        """
        Returns true if there is an update waiting to be applied
        """
        return pulumi.get(self, "update_available")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time the instance was updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def zones(self) -> Sequence[str]:
        """
        Zones in which Memcached nodes should be provisioned. Memcached nodes will be equally distributed across these zones. If not provided, the service will by default create nodes in all zones in the region for the instance.
        """
        return pulumi.get(self, "zones")


class AwaitableGetInstanceResult(GetInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInstanceResult(
            authorized_network=self.authorized_network,
            create_time=self.create_time,
            discovery_endpoint=self.discovery_endpoint,
            display_name=self.display_name,
            instance_messages=self.instance_messages,
            labels=self.labels,
            memcache_full_version=self.memcache_full_version,
            memcache_nodes=self.memcache_nodes,
            memcache_version=self.memcache_version,
            name=self.name,
            node_config=self.node_config,
            node_count=self.node_count,
            parameters=self.parameters,
            state=self.state,
            update_available=self.update_available,
            update_time=self.update_time,
            zones=self.zones)


def get_instance(instance_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInstanceResult:
    """
    Gets details of a single Instance.
    """
    __args__ = dict()
    __args__['instanceId'] = instance_id
    __args__['location'] = location
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:memcache/v1beta2:getInstance', __args__, opts=opts, typ=GetInstanceResult).value

    return AwaitableGetInstanceResult(
        authorized_network=__ret__.authorized_network,
        create_time=__ret__.create_time,
        discovery_endpoint=__ret__.discovery_endpoint,
        display_name=__ret__.display_name,
        instance_messages=__ret__.instance_messages,
        labels=__ret__.labels,
        memcache_full_version=__ret__.memcache_full_version,
        memcache_nodes=__ret__.memcache_nodes,
        memcache_version=__ret__.memcache_version,
        name=__ret__.name,
        node_config=__ret__.node_config,
        node_count=__ret__.node_count,
        parameters=__ret__.parameters,
        state=__ret__.state,
        update_available=__ret__.update_available,
        update_time=__ret__.update_time,
        zones=__ret__.zones)


@_utilities.lift_output_func(get_instance)
def get_instance_output(instance_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetInstanceResult]:
    """
    Gets details of a single Instance.
    """
    ...