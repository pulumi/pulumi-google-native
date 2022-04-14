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
    'GetMetricDescriptorResult',
    'AwaitableGetMetricDescriptorResult',
    'get_metric_descriptor',
    'get_metric_descriptor_output',
]

@pulumi.output_type
class GetMetricDescriptorResult:
    def __init__(__self__, description=None, display_name=None, labels=None, launch_stage=None, metadata=None, metric_kind=None, monitored_resource_types=None, name=None, type=None, unit=None, value_type=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if labels and not isinstance(labels, list):
            raise TypeError("Expected argument 'labels' to be a list")
        pulumi.set(__self__, "labels", labels)
        if launch_stage and not isinstance(launch_stage, str):
            raise TypeError("Expected argument 'launch_stage' to be a str")
        pulumi.set(__self__, "launch_stage", launch_stage)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if metric_kind and not isinstance(metric_kind, str):
            raise TypeError("Expected argument 'metric_kind' to be a str")
        pulumi.set(__self__, "metric_kind", metric_kind)
        if monitored_resource_types and not isinstance(monitored_resource_types, list):
            raise TypeError("Expected argument 'monitored_resource_types' to be a list")
        pulumi.set(__self__, "monitored_resource_types", monitored_resource_types)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if unit and not isinstance(unit, str):
            raise TypeError("Expected argument 'unit' to be a str")
        pulumi.set(__self__, "unit", unit)
        if value_type and not isinstance(value_type, str):
            raise TypeError("Expected argument 'value_type' to be a str")
        pulumi.set(__self__, "value_type", value_type)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A detailed description of the metric, which can be used in documentation.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count". This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> Sequence['outputs.LabelDescriptorResponse']:
        """
        The set of labels that can be used to describe a specific instance of this metric type. For example, the appengine.googleapis.com/http/server/response_latencies metric type has a label for the HTTP response code, response_code, so you can look at latencies for successful responses or just for responses that failed.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="launchStage")
    def launch_stage(self) -> str:
        """
        Optional. The launch stage of the metric definition.
        """
        return pulumi.get(self, "launch_stage")

    @property
    @pulumi.getter
    def metadata(self) -> 'outputs.MetricDescriptorMetadataResponse':
        """
        Optional. Metadata which can be used to guide usage of the metric.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metricKind")
    def metric_kind(self) -> str:
        """
        Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metric_kind and value_type might not be supported.
        """
        return pulumi.get(self, "metric_kind")

    @property
    @pulumi.getter(name="monitoredResourceTypes")
    def monitored_resource_types(self) -> Sequence[str]:
        """
        Read-only. If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here.
        """
        return pulumi.get(self, "monitored_resource_types")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the metric descriptor.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The metric type, including its DNS name prefix. The type is not URL-encoded. All user-defined metric types have the DNS name custom.googleapis.com or external.googleapis.com. Metric types should use a natural hierarchical grouping. For example: "custom.googleapis.com/invoice/paid/amount" "external.googleapis.com/prometheus/up" "appengine.googleapis.com/http/server/response_latencies" 
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def unit(self) -> str:
        """
        The units in which the metric value is reported. It is only applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of the stored metric values.Different systems might scale the values to be more easily displayed (so a value of 0.02kBy might be displayed as 20By, and a value of 3523kBy might be displayed as 3.5MBy). However, if the unit is kBy, then the value of the metric is always in thousands of bytes, no matter how it might be displayed.If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as 12005.Alternatively, if you want a custom metric to record data in a more granular way, you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).The supported units are a subset of The Unified Code for Units of Measure (https://unitsofmeasure.org/ucum.html) standard:Basic units (UNIT) bit bit By byte s second min minute h hour d day 1 dimensionlessPrefixes (PREFIX) k kilo (10^3) M mega (10^6) G giga (10^9) T tera (10^12) P peta (10^15) E exa (10^18) Z zetta (10^21) Y yotta (10^24) m milli (10^-3) u micro (10^-6) n nano (10^-9) p pico (10^-12) f femto (10^-15) a atto (10^-18) z zepto (10^-21) y yocto (10^-24) Ki kibi (2^10) Mi mebi (2^20) Gi gibi (2^30) Ti tebi (2^40) Pi pebi (2^50)GrammarThe grammar also includes these connectors: / division or ratio (as an infix operator). For examples, kBy/{email} or MiBy/10ms (although you should almost never have /s in a metric unit; rates should always be computed at query time from the underlying cumulative or delta value). . multiplication or composition (as an infix operator). For examples, GBy.d or k{watt}.h.The grammar for a unit is as follows: Expression = Component { "." Component } { "/" Component } ; Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ] | Annotation | "1" ; Annotation = "{" NAME "}" ; Notes: Annotation is just a comment if it follows a UNIT. If the annotation is used alone, then the unit is equivalent to 1. For examples, {request}/s == 1/s, By{transmitted}/s == By/s. NAME is a sequence of non-blank printable ASCII characters not containing { or }. 1 represents a unitary dimensionless unit (https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such as in 1/s. It is typically used when none of the basic units are appropriate. For example, "new users per day" can be represented as 1/d or {new-users}/d (and a metric value 5 would mean "5 new users). Alternatively, "thousands of page views per day" would be represented as 1000/d or k1/d or k{page_views}/d (and a metric value of 5.3 would mean "5300 page views per day"). % represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0..100, and a metric value 3 means "3 percent"). 10^2.% indicates a metric contains a ratio, typically in the range 0..1, that will be multiplied by 100 and displayed as a percentage (so a metric value 0.03 means "3 percent").
        """
        return pulumi.get(self, "unit")

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> str:
        """
        Whether the measurement is an integer, a floating-point number, etc. Some combinations of metric_kind and value_type might not be supported.
        """
        return pulumi.get(self, "value_type")


class AwaitableGetMetricDescriptorResult(GetMetricDescriptorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMetricDescriptorResult(
            description=self.description,
            display_name=self.display_name,
            labels=self.labels,
            launch_stage=self.launch_stage,
            metadata=self.metadata,
            metric_kind=self.metric_kind,
            monitored_resource_types=self.monitored_resource_types,
            name=self.name,
            type=self.type,
            unit=self.unit,
            value_type=self.value_type)


def get_metric_descriptor(metric_descriptor_id: Optional[str] = None,
                          project: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMetricDescriptorResult:
    """
    Gets a single metric descriptor.
    """
    __args__ = dict()
    __args__['metricDescriptorId'] = metric_descriptor_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:monitoring/v3:getMetricDescriptor', __args__, opts=opts, typ=GetMetricDescriptorResult).value

    return AwaitableGetMetricDescriptorResult(
        description=__ret__.description,
        display_name=__ret__.display_name,
        labels=__ret__.labels,
        launch_stage=__ret__.launch_stage,
        metadata=__ret__.metadata,
        metric_kind=__ret__.metric_kind,
        monitored_resource_types=__ret__.monitored_resource_types,
        name=__ret__.name,
        type=__ret__.type,
        unit=__ret__.unit,
        value_type=__ret__.value_type)


@_utilities.lift_output_func(get_metric_descriptor)
def get_metric_descriptor_output(metric_descriptor_id: Optional[pulumi.Input[str]] = None,
                                 project: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetMetricDescriptorResult]:
    """
    Gets a single metric descriptor.
    """
    ...