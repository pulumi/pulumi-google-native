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

__all__ = [
    'GetTaxonomyResult',
    'AwaitableGetTaxonomyResult',
    'get_taxonomy',
    'get_taxonomy_output',
]

@pulumi.output_type
class GetTaxonomyResult:
    def __init__(__self__, activated_policy_types=None, description=None, display_name=None, name=None, policy_tag_count=None, service=None, taxonomy_timestamps=None):
        if activated_policy_types and not isinstance(activated_policy_types, list):
            raise TypeError("Expected argument 'activated_policy_types' to be a list")
        pulumi.set(__self__, "activated_policy_types", activated_policy_types)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_tag_count and not isinstance(policy_tag_count, int):
            raise TypeError("Expected argument 'policy_tag_count' to be a int")
        pulumi.set(__self__, "policy_tag_count", policy_tag_count)
        if service and not isinstance(service, dict):
            raise TypeError("Expected argument 'service' to be a dict")
        pulumi.set(__self__, "service", service)
        if taxonomy_timestamps and not isinstance(taxonomy_timestamps, dict):
            raise TypeError("Expected argument 'taxonomy_timestamps' to be a dict")
        pulumi.set(__self__, "taxonomy_timestamps", taxonomy_timestamps)

    @property
    @pulumi.getter(name="activatedPolicyTypes")
    def activated_policy_types(self) -> Sequence[str]:
        """
        Optional. A list of policy types that are activated for this taxonomy. If not set, defaults to an empty list.
        """
        return pulumi.get(self, "activated_policy_types")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of this taxonomy. If not set, defaults to empty. The description must contain only Unicode characters, tabs, newlines, carriage returns, and page breaks, and be at most 2000 bytes long when encoded in UTF-8.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        User-defined name of this taxonomy. The name can't start or end with spaces, must contain only Unicode letters, numbers, underscores, dashes, and spaces, and be at most 200 bytes long when encoded in UTF-8. The taxonomy display name must be unique within an organization.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of this taxonomy in URL format. Note: Policy tag manager generates unique taxonomy IDs.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyTagCount")
    def policy_tag_count(self) -> int:
        """
        Number of policy tags in this taxonomy.
        """
        return pulumi.get(self, "policy_tag_count")

    @property
    @pulumi.getter
    def service(self) -> 'outputs.GoogleCloudDatacatalogV1TaxonomyServiceResponse':
        """
        Identity of the service which owns the Taxonomy. This field is only populated when the taxonomy is created by a Google Cloud service. Currently only 'DATAPLEX' is supported.
        """
        return pulumi.get(self, "service")

    @property
    @pulumi.getter(name="taxonomyTimestamps")
    def taxonomy_timestamps(self) -> 'outputs.GoogleCloudDatacatalogV1SystemTimestampsResponse':
        """
        Creation and modification timestamps of this taxonomy.
        """
        return pulumi.get(self, "taxonomy_timestamps")


class AwaitableGetTaxonomyResult(GetTaxonomyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTaxonomyResult(
            activated_policy_types=self.activated_policy_types,
            description=self.description,
            display_name=self.display_name,
            name=self.name,
            policy_tag_count=self.policy_tag_count,
            service=self.service,
            taxonomy_timestamps=self.taxonomy_timestamps)


def get_taxonomy(location: Optional[str] = None,
                 project: Optional[str] = None,
                 taxonomy_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTaxonomyResult:
    """
    Gets a taxonomy.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['taxonomyId'] = taxonomy_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datacatalog/v1:getTaxonomy', __args__, opts=opts, typ=GetTaxonomyResult).value

    return AwaitableGetTaxonomyResult(
        activated_policy_types=pulumi.get(__ret__, 'activated_policy_types'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        name=pulumi.get(__ret__, 'name'),
        policy_tag_count=pulumi.get(__ret__, 'policy_tag_count'),
        service=pulumi.get(__ret__, 'service'),
        taxonomy_timestamps=pulumi.get(__ret__, 'taxonomy_timestamps'))


@_utilities.lift_output_func(get_taxonomy)
def get_taxonomy_output(location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        taxonomy_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTaxonomyResult]:
    """
    Gets a taxonomy.
    """
    ...
