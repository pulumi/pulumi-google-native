import pulumi
from pulumi_google_native.container import v1 as container
import pulumi_kubernetes as k8s

config = pulumi.Config("google-native")
PROJECT_ID = config.require("project")
LOCATION = config.require("location")

server_config = container.get_server_config_output(location=LOCATION)

cluster = container.Cluster("cluster", container.ClusterArgs(
    initial_cluster_version=server_config.valid_master_versions[0],
    node_pools=[container.NodePoolArgs(initial_node_count=1, name="initial")],
))

k8sprov = k8s.Provider("k8sprov", k8s.ProviderArgs(
    kubeconfig=cluster.get_kubeconfig()))
namespace = k8s.core.v1.Namespace(
    "ns", opts=pulumi.ResourceOptions(provider=k8sprov))

pulumi.export("ns", namespace.metadata.name)
