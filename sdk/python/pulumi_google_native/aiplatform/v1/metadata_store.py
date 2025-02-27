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
from ._inputs import *

__all__ = ['MetadataStoreArgs', 'MetadataStore']

@pulumi.input_type
class MetadataStoreArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_spec: Optional[pulumi.Input['GoogleCloudAiplatformV1EncryptionSpecArgs']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata_store_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MetadataStore resource.
        :param pulumi.Input[str] description: Description of the MetadataStore.
        :param pulumi.Input['GoogleCloudAiplatformV1EncryptionSpecArgs'] encryption_spec: Customer-managed encryption key spec for a Metadata Store. If set, this Metadata Store and all sub-resources of this Metadata Store are secured using this key.
        :param pulumi.Input[str] metadata_store_id: The {metadatastore} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}` If not provided, the MetadataStore's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all MetadataStores in the parent Location. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting MetadataStore.)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if encryption_spec is not None:
            pulumi.set(__self__, "encryption_spec", encryption_spec)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if metadata_store_id is not None:
            pulumi.set(__self__, "metadata_store_id", metadata_store_id)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the MetadataStore.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="encryptionSpec")
    def encryption_spec(self) -> Optional[pulumi.Input['GoogleCloudAiplatformV1EncryptionSpecArgs']]:
        """
        Customer-managed encryption key spec for a Metadata Store. If set, this Metadata Store and all sub-resources of this Metadata Store are secured using this key.
        """
        return pulumi.get(self, "encryption_spec")

    @encryption_spec.setter
    def encryption_spec(self, value: Optional[pulumi.Input['GoogleCloudAiplatformV1EncryptionSpecArgs']]):
        pulumi.set(self, "encryption_spec", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="metadataStoreId")
    def metadata_store_id(self) -> Optional[pulumi.Input[str]]:
        """
        The {metadatastore} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}` If not provided, the MetadataStore's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all MetadataStores in the parent Location. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting MetadataStore.)
        """
        return pulumi.get(self, "metadata_store_id")

    @metadata_store_id.setter
    def metadata_store_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_store_id", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class MetadataStore(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudAiplatformV1EncryptionSpecArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata_store_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Initializes a MetadataStore, including allocation of resources.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of the MetadataStore.
        :param pulumi.Input[pulumi.InputType['GoogleCloudAiplatformV1EncryptionSpecArgs']] encryption_spec: Customer-managed encryption key spec for a Metadata Store. If set, this Metadata Store and all sub-resources of this Metadata Store are secured using this key.
        :param pulumi.Input[str] metadata_store_id: The {metadatastore} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}` If not provided, the MetadataStore's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all MetadataStores in the parent Location. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting MetadataStore.)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[MetadataStoreArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Initializes a MetadataStore, including allocation of resources.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param MetadataStoreArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MetadataStoreArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 encryption_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudAiplatformV1EncryptionSpecArgs']]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata_store_id: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MetadataStoreArgs.__new__(MetadataStoreArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["encryption_spec"] = encryption_spec
            __props__.__dict__["location"] = location
            __props__.__dict__["metadata_store_id"] = metadata_store_id
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(MetadataStore, __self__).__init__(
            'google-native:aiplatform/v1:MetadataStore',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'MetadataStore':
        """
        Get an existing MetadataStore resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = MetadataStoreArgs.__new__(MetadataStoreArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["encryption_spec"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["metadata_store_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return MetadataStore(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Timestamp when this MetadataStore was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description of the MetadataStore.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="encryptionSpec")
    def encryption_spec(self) -> pulumi.Output['outputs.GoogleCloudAiplatformV1EncryptionSpecResponse']:
        """
        Customer-managed encryption key spec for a Metadata Store. If set, this Metadata Store and all sub-resources of this Metadata Store are secured using this key.
        """
        return pulumi.get(self, "encryption_spec")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="metadataStoreId")
    def metadata_store_id(self) -> pulumi.Output[Optional[str]]:
        """
        The {metadatastore} portion of the resource name with the format: `projects/{project}/locations/{location}/metadataStores/{metadatastore}` If not provided, the MetadataStore's ID will be a UUID generated by the service. Must be 4-128 characters in length. Valid characters are `/a-z-/`. Must be unique across all MetadataStores in the parent Location. (Otherwise the request will fail with ALREADY_EXISTS, or PERMISSION_DENIED if the caller can't view the preexisting MetadataStore.)
        """
        return pulumi.get(self, "metadata_store_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the MetadataStore instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output['outputs.GoogleCloudAiplatformV1MetadataStoreMetadataStoreStateResponse']:
        """
        State information of the MetadataStore.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Timestamp when this MetadataStore was last updated.
        """
        return pulumi.get(self, "update_time")

