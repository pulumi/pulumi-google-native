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

__all__ = ['CatalogItemArgs', 'CatalogItem']

@pulumi.input_type
class CatalogItemArgs:
    def __init__(__self__, *,
                 catalog_id: pulumi.Input[str],
                 category_hierarchies: pulumi.Input[Sequence[pulumi.Input['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]],
                 id: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 item_attributes: Optional[pulumi.Input['GoogleCloudRecommendationengineV1beta1FeatureMapArgs']] = None,
                 item_group_id: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 product_metadata: Optional[pulumi.Input['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CatalogItem resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]] category_hierarchies: Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
        :param pulumi.Input[str] id: Catalog item identifier. UTF-8 encoded string with a length limit of 128 bytes. This id must be unique among all catalog items within the same catalog. It should also be used when logging user events in order for the user events to be joined with the Catalog.
        :param pulumi.Input[str] title: Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
        :param pulumi.Input[str] description: Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
        :param pulumi.Input['GoogleCloudRecommendationengineV1beta1FeatureMapArgs'] item_attributes: Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
        :param pulumi.Input[str] item_group_id: Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
        :param pulumi.Input[str] language_code: Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.
        :param pulumi.Input['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs'] product_metadata: Optional. Metadata specific to retail products.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "category_hierarchies", category_hierarchies)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if item_attributes is not None:
            pulumi.set(__self__, "item_attributes", item_attributes)
        if item_group_id is not None:
            pulumi.set(__self__, "item_group_id", item_group_id)
        if language_code is not None:
            warnings.warn("""Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.""", DeprecationWarning)
            pulumi.log.warn("""language_code is deprecated: Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.""")
        if language_code is not None:
            pulumi.set(__self__, "language_code", language_code)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if product_metadata is not None:
            pulumi.set(__self__, "product_metadata", product_metadata)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="categoryHierarchies")
    def category_hierarchies(self) -> pulumi.Input[Sequence[pulumi.Input['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]]:
        """
        Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
        """
        return pulumi.get(self, "category_hierarchies")

    @category_hierarchies.setter
    def category_hierarchies(self, value: pulumi.Input[Sequence[pulumi.Input['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]]):
        pulumi.set(self, "category_hierarchies", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        Catalog item identifier. UTF-8 encoded string with a length limit of 128 bytes. This id must be unique among all catalog items within the same catalog. It should also be used when logging user events in order for the user events to be joined with the Catalog.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="itemAttributes")
    def item_attributes(self) -> Optional[pulumi.Input['GoogleCloudRecommendationengineV1beta1FeatureMapArgs']]:
        """
        Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
        """
        return pulumi.get(self, "item_attributes")

    @item_attributes.setter
    def item_attributes(self, value: Optional[pulumi.Input['GoogleCloudRecommendationengineV1beta1FeatureMapArgs']]):
        pulumi.set(self, "item_attributes", value)

    @property
    @pulumi.getter(name="itemGroupId")
    def item_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
        """
        return pulumi.get(self, "item_group_id")

    @item_group_id.setter
    def item_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "item_group_id", value)

    @property
    @pulumi.getter(name="languageCode")
    @_utilities.deprecated("""Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.""")
    def language_code(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="productMetadata")
    def product_metadata(self) -> Optional[pulumi.Input['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs']]:
        """
        Optional. Metadata specific to retail products.
        """
        return pulumi.get(self, "product_metadata")

    @product_metadata.setter
    def product_metadata(self, value: Optional[pulumi.Input['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs']]):
        pulumi.set(self, "product_metadata", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class CatalogItem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 category_hierarchies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 item_attributes: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1FeatureMapArgs']]] = None,
                 item_group_id: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 product_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a catalog item.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]]] category_hierarchies: Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
        :param pulumi.Input[str] description: Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
        :param pulumi.Input[str] id: Catalog item identifier. UTF-8 encoded string with a length limit of 128 bytes. This id must be unique among all catalog items within the same catalog. It should also be used when logging user events in order for the user events to be joined with the Catalog.
        :param pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1FeatureMapArgs']] item_attributes: Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
        :param pulumi.Input[str] item_group_id: Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
        :param pulumi.Input[str] language_code: Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.
        :param pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs']] product_metadata: Optional. Metadata specific to retail products.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
        :param pulumi.Input[str] title: Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CatalogItemArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a catalog item.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param CatalogItemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CatalogItemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 category_hierarchies: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 item_attributes: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1FeatureMapArgs']]] = None,
                 item_group_id: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 product_metadata: Optional[pulumi.Input[pulumi.InputType['GoogleCloudRecommendationengineV1beta1ProductCatalogItemArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CatalogItemArgs.__new__(CatalogItemArgs)

            if catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_id'")
            __props__.__dict__["catalog_id"] = catalog_id
            if category_hierarchies is None and not opts.urn:
                raise TypeError("Missing required property 'category_hierarchies'")
            __props__.__dict__["category_hierarchies"] = category_hierarchies
            __props__.__dict__["description"] = description
            if id is None and not opts.urn:
                raise TypeError("Missing required property 'id'")
            __props__.__dict__["id"] = id
            __props__.__dict__["item_attributes"] = item_attributes
            __props__.__dict__["item_group_id"] = item_group_id
            __props__.__dict__["language_code"] = language_code
            __props__.__dict__["location"] = location
            __props__.__dict__["product_metadata"] = product_metadata
            __props__.__dict__["project"] = project
            __props__.__dict__["tags"] = tags
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["catalogId", "location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(CatalogItem, __self__).__init__(
            'google-native:recommendationengine/v1beta1:CatalogItem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CatalogItem':
        """
        Get an existing CatalogItem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CatalogItemArgs.__new__(CatalogItemArgs)

        __props__.__dict__["catalog_id"] = None
        __props__.__dict__["category_hierarchies"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["item_attributes"] = None
        __props__.__dict__["item_group_id"] = None
        __props__.__dict__["language_code"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["product_metadata"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["title"] = None
        return CatalogItem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="categoryHierarchies")
    def category_hierarchies(self) -> pulumi.Output[Sequence['outputs.GoogleCloudRecommendationengineV1beta1CatalogItemCategoryHierarchyResponse']]:
        """
        Catalog item categories. This field is repeated for supporting one catalog item belonging to several parallel category hierarchies. For example, if a shoes product belongs to both ["Shoes & Accessories" -> "Shoes"] and ["Sports & Fitness" -> "Athletic Clothing" -> "Shoes"], it could be represented as: "categoryHierarchies": [ { "categories": ["Shoes & Accessories", "Shoes"]}, { "categories": ["Sports & Fitness", "Athletic Clothing", "Shoes"] } ]
        """
        return pulumi.get(self, "category_hierarchies")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Catalog item description. UTF-8 encoded string with a length limit of 5 KiB.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="itemAttributes")
    def item_attributes(self) -> pulumi.Output['outputs.GoogleCloudRecommendationengineV1beta1FeatureMapResponse']:
        """
        Optional. Highly encouraged. Extra catalog item attributes to be included in the recommendation model. For example, for retail products, this could include the store name, vendor, style, color, etc. These are very strong signals for recommendation model, thus we highly recommend providing the item attributes here.
        """
        return pulumi.get(self, "item_attributes")

    @property
    @pulumi.getter(name="itemGroupId")
    def item_group_id(self) -> pulumi.Output[str]:
        """
        Optional. Variant group identifier for prediction results. UTF-8 encoded string with a length limit of 128 bytes. This field must be enabled before it can be used. [Learn more](/recommendations-ai/docs/catalog#item-group-id).
        """
        return pulumi.get(self, "item_group_id")

    @property
    @pulumi.getter(name="languageCode")
    @_utilities.deprecated("""Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.""")
    def language_code(self) -> pulumi.Output[str]:
        """
        Optional. Deprecated. The model automatically detects the text language. Your catalog can include text in different languages, but duplicating catalog items to provide text in multiple languages can result in degraded model performance.
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="productMetadata")
    def product_metadata(self) -> pulumi.Output['outputs.GoogleCloudRecommendationengineV1beta1ProductCatalogItemResponse']:
        """
        Optional. Metadata specific to retail products.
        """
        return pulumi.get(self, "product_metadata")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. Filtering tags associated with the catalog item. Each tag should be a UTF-8 encoded string with a length limit of 1 KiB. This tag can be used for filtering recommendation results by passing the tag as part of the predict request filter.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        Catalog item title. UTF-8 encoded string with a length limit of 1 KiB.
        """
        return pulumi.get(self, "title")

