# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetContextResult',
    'AwaitableGetContextResult',
    'get_context',
    'get_context_output',
]

@pulumi.output_type
class GetContextResult:
    def __init__(__self__, create_time=None, description=None, display_name=None, etag=None, labels=None, metadata=None, name=None, parent_contexts=None, schema_title=None, schema_version=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_contexts and not isinstance(parent_contexts, list):
            raise TypeError("Expected argument 'parent_contexts' to be a list")
        pulumi.set(__self__, "parent_contexts", parent_contexts)
        if schema_title and not isinstance(schema_title, str):
            raise TypeError("Expected argument 'schema_title' to be a str")
        pulumi.set(__self__, "schema_title", schema_title)
        if schema_version and not isinstance(schema_version, str):
            raise TypeError("Expected argument 'schema_version' to be a str")
        pulumi.set(__self__, "schema_version", schema_version)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when this Context was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the Context
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        User provided display name of the Context. May be up to 128 Unicode characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        An eTag used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels with user-defined metadata to organize your Contexts. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. No more than 64 user labels can be associated with one Context (System labels are excluded).
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, Any]:
        """
        Properties of the Context. Top level metadata keys' heading and trailing spaces will be trimmed. The size of this field should not exceed 200KB.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The resource name of the Context.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentContexts")
    def parent_contexts(self) -> Sequence[str]:
        """
        A list of resource names of Contexts that are parents of this Context. A Context may have at most 10 parent_contexts.
        """
        return pulumi.get(self, "parent_contexts")

    @property
    @pulumi.getter(name="schemaTitle")
    def schema_title(self) -> str:
        """
        The title of the schema describing the metadata. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
        """
        return pulumi.get(self, "schema_title")

    @property
    @pulumi.getter(name="schemaVersion")
    def schema_version(self) -> str:
        """
        The version of the schema in schema_name to use. Schema title and version is expected to be registered in earlier Create Schema calls. And both are used together as unique identifiers to identify schemas within the local metadata store.
        """
        return pulumi.get(self, "schema_version")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Timestamp when this Context was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetContextResult(GetContextResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetContextResult(
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            etag=self.etag,
            labels=self.labels,
            metadata=self.metadata,
            name=self.name,
            parent_contexts=self.parent_contexts,
            schema_title=self.schema_title,
            schema_version=self.schema_version,
            update_time=self.update_time)


def get_context(context_id: Optional[str] = None,
                location: Optional[str] = None,
                metadata_store_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetContextResult:
    """
    Retrieves a specific Context.
    """
    __args__ = dict()
    __args__['contextId'] = context_id
    __args__['location'] = location
    __args__['metadataStoreId'] = metadata_store_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:aiplatform/v1:getContext', __args__, opts=opts, typ=GetContextResult).value

    return AwaitableGetContextResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        etag=pulumi.get(__ret__, 'etag'),
        labels=pulumi.get(__ret__, 'labels'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        parent_contexts=pulumi.get(__ret__, 'parent_contexts'),
        schema_title=pulumi.get(__ret__, 'schema_title'),
        schema_version=pulumi.get(__ret__, 'schema_version'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_context)
def get_context_output(context_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       metadata_store_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetContextResult]:
    """
    Retrieves a specific Context.
    """
    ...
