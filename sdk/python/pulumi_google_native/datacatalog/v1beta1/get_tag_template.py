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
    'GetTagTemplateResult',
    'AwaitableGetTagTemplateResult',
    'get_tag_template',
    'get_tag_template_output',
]

@pulumi.output_type
class GetTagTemplateResult:
    def __init__(__self__, display_name=None, fields=None, name=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name for this template. Defaults to an empty string.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def fields(self) -> Mapping[str, 'outputs.GoogleCloudDatacatalogV1beta1TagTemplateFieldResponse']:
        """
        Map of tag template field IDs to the settings for the field. This map is an exhaustive list of the allowed fields. This map must contain at least one field and at most 500 fields. The keys to this map are tag template field IDs. Field IDs can contain letters (both uppercase and lowercase), numbers (0-9) and underscores (_). Field IDs must be at least 1 character long and at most 64 characters long. Field IDs must start with a letter or underscore.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the tag template in URL format. Example: * projects/{project_id}/locations/{location}/tagTemplates/{tag_template_id} Note that this TagTemplate and its child resources may not actually be stored in the location in this name.
        """
        return pulumi.get(self, "name")


class AwaitableGetTagTemplateResult(GetTagTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTagTemplateResult(
            display_name=self.display_name,
            fields=self.fields,
            name=self.name)


def get_tag_template(location: Optional[str] = None,
                     project: Optional[str] = None,
                     tag_template_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTagTemplateResult:
    """
    Gets a tag template.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['tagTemplateId'] = tag_template_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datacatalog/v1beta1:getTagTemplate', __args__, opts=opts, typ=GetTagTemplateResult).value

    return AwaitableGetTagTemplateResult(
        display_name=pulumi.get(__ret__, 'display_name'),
        fields=pulumi.get(__ret__, 'fields'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_tag_template)
def get_tag_template_output(location: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            tag_template_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTagTemplateResult]:
    """
    Gets a tag template.
    """
    ...
