# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'GceClusterConfigPrivateIpv6GoogleAccess',
    'GkeNodePoolTargetRolesItem',
    'InstanceGroupConfigPreemptibility',
    'MetricMetricSource',
    'ReservationAffinityConsumeReservationType',
    'SoftwareConfigOptionalComponentsItem',
]


class GceClusterConfigPrivateIpv6GoogleAccess(str, Enum):
    """
    Optional. The type of IPv6 access for a cluster.
    """
    PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED = "PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED"
    """
    If unspecified, Compute Engine default behavior will apply, which is the same as INHERIT_FROM_SUBNETWORK.
    """
    INHERIT_FROM_SUBNETWORK = "INHERIT_FROM_SUBNETWORK"
    """
    Private access to and from Google Services configuration inherited from the subnetwork configuration. This is the default Compute Engine behavior.
    """
    OUTBOUND = "OUTBOUND"
    """
    Enables outbound private IPv6 access to Google Services from the Dataproc cluster.
    """
    BIDIRECTIONAL = "BIDIRECTIONAL"
    """
    Enables bidirectional private IPv6 access between Google Services and the Dataproc cluster.
    """


class GkeNodePoolTargetRolesItem(str, Enum):
    ROLE_UNSPECIFIED = "ROLE_UNSPECIFIED"
    """
    Role is unspecified.
    """
    DEFAULT = "DEFAULT"
    """
    Any roles that are not directly assigned to a NodePool run on the default role's NodePool.
    """
    CONTROLLER = "CONTROLLER"
    """
    Run controllers and webhooks.
    """
    SPARK_DRIVER = "SPARK_DRIVER"
    """
    Run spark driver.
    """
    SPARK_EXECUTOR = "SPARK_EXECUTOR"
    """
    Run spark executors.
    """


class InstanceGroupConfigPreemptibility(str, Enum):
    """
    Optional. Specifies the preemptibility of the instance group.The default value for master and worker groups is NON_PREEMPTIBLE. This default cannot be changed.The default value for secondary instances is PREEMPTIBLE.
    """
    PREEMPTIBILITY_UNSPECIFIED = "PREEMPTIBILITY_UNSPECIFIED"
    """
    Preemptibility is unspecified, the system will choose the appropriate setting for each instance group.
    """
    NON_PREEMPTIBLE = "NON_PREEMPTIBLE"
    """
    Instances are non-preemptible.This option is allowed for all instance groups and is the only valid value for Master and Worker instance groups.
    """
    PREEMPTIBLE = "PREEMPTIBLE"
    """
    Instances are preemptible.This option is allowed only for secondary worker groups.
    """


class MetricMetricSource(str, Enum):
    """
    Required. MetricSource that should be enabled
    """
    METRIC_SOURCE_UNSPECIFIED = "METRIC_SOURCE_UNSPECIFIED"
    """
    Required unspecified metric source
    """
    MONITORING_AGENT_DEFAULTS = "MONITORING_AGENT_DEFAULTS"
    """
    all default monitoring agent metrics that are published with prefix "agent.googleapis.com" when we enable a monitoring agent in Compute Engine
    """
    HDFS = "HDFS"
    """
    Hdfs metric source
    """
    SPARK = "SPARK"
    """
    Spark metric source
    """
    YARN = "YARN"
    """
    Yarn metric source
    """
    SPARK_HISTORY_SERVER = "SPARK_HISTORY_SERVER"
    """
    Spark history server metric source
    """
    HIVESERVER2 = "HIVESERVER2"
    """
    hiveserver2 metric source
    """


class ReservationAffinityConsumeReservationType(str, Enum):
    """
    Optional. Type of reservation to consume
    """
    TYPE_UNSPECIFIED = "TYPE_UNSPECIFIED"
    NO_RESERVATION = "NO_RESERVATION"
    """
    Do not consume from any allocated capacity.
    """
    ANY_RESERVATION = "ANY_RESERVATION"
    """
    Consume any reservation available.
    """
    SPECIFIC_RESERVATION = "SPECIFIC_RESERVATION"
    """
    Must consume from a specific reservation. Must specify key value fields for specifying the reservations.
    """


class SoftwareConfigOptionalComponentsItem(str, Enum):
    COMPONENT_UNSPECIFIED = "COMPONENT_UNSPECIFIED"
    """
    Unspecified component. Specifying this will cause Cluster creation to fail.
    """
    ANACONDA = "ANACONDA"
    """
    The Anaconda python distribution. The Anaconda component is not supported in the Dataproc 2.0 image. The 2.0 image is pre-installed with Miniconda.
    """
    DOCKER = "DOCKER"
    """
    Docker
    """
    DRUID = "DRUID"
    """
    The Druid query engine. (alpha)
    """
    FLINK = "FLINK"
    """
    Flink
    """
    HBASE = "HBASE"
    """
    HBase. (beta)
    """
    HIVE_WEBHCAT = "HIVE_WEBHCAT"
    """
    The Hive Web HCatalog (the REST service for accessing HCatalog).
    """
    JUPYTER = "JUPYTER"
    """
    The Jupyter Notebook.
    """
    PRESTO = "PRESTO"
    """
    The Presto query engine.
    """
    RANGER = "RANGER"
    """
    The Ranger service.
    """
    SOLR = "SOLR"
    """
    The Solr service.
    """
    ZEPPELIN = "ZEPPELIN"
    """
    The Zeppelin notebook.
    """
    ZOOKEEPER = "ZOOKEEPER"
    """
    The Zookeeper service.
    """