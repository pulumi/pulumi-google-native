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
    'GetFeatureViewResult',
    'AwaitableGetFeatureViewResult',
    'get_feature_view',
    'get_feature_view_output',
]

@pulumi.output_type
class GetFeatureViewResult:
    def __init__(__self__, big_query_source=None, create_time=None, etag=None, feature_registry_source=None, labels=None, name=None, sync_config=None, update_time=None, vector_search_config=None):
        if big_query_source and not isinstance(big_query_source, dict):
            raise TypeError("Expected argument 'big_query_source' to be a dict")
        pulumi.set(__self__, "big_query_source", big_query_source)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if feature_registry_source and not isinstance(feature_registry_source, dict):
            raise TypeError("Expected argument 'feature_registry_source' to be a dict")
        pulumi.set(__self__, "feature_registry_source", feature_registry_source)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sync_config and not isinstance(sync_config, dict):
            raise TypeError("Expected argument 'sync_config' to be a dict")
        pulumi.set(__self__, "sync_config", sync_config)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if vector_search_config and not isinstance(vector_search_config, dict):
            raise TypeError("Expected argument 'vector_search_config' to be a dict")
        pulumi.set(__self__, "vector_search_config", vector_search_config)

    @property
    @pulumi.getter(name="bigQuerySource")
    def big_query_source(self) -> 'outputs.GoogleCloudAiplatformV1beta1FeatureViewBigQuerySourceResponse':
        """
        Optional. Configures how data is supposed to be extracted from a BigQuery source to be loaded onto the FeatureOnlineStore.
        """
        return pulumi.get(self, "big_query_source")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Timestamp when this FeatureView was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Optional. Used to perform consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="featureRegistrySource")
    def feature_registry_source(self) -> 'outputs.GoogleCloudAiplatformV1beta1FeatureViewFeatureRegistrySourceResponse':
        """
        Optional. Configures the features from a Feature Registry source that need to be loaded onto the FeatureOnlineStore.
        """
        return pulumi.get(self, "feature_registry_source")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. The labels with user-defined metadata to organize your FeatureViews. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information on and examples of labels. No more than 64 user labels can be associated with one FeatureOnlineStore(System labels are excluded)." System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the FeatureView. Format: `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="syncConfig")
    def sync_config(self) -> 'outputs.GoogleCloudAiplatformV1beta1FeatureViewSyncConfigResponse':
        """
        Configures when data is to be synced/updated for this FeatureView. At the end of the sync the latest featureValues for each entityId of this FeatureView are made ready for online serving.
        """
        return pulumi.get(self, "sync_config")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Timestamp when this FeatureView was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vectorSearchConfig")
    def vector_search_config(self) -> 'outputs.GoogleCloudAiplatformV1beta1FeatureViewVectorSearchConfigResponse':
        """
        Optional. Configuration for vector search. It contains the required configurations to create an index from source data, so that approximate nearest neighbor (a.k.a ANN) algorithms search can be performed during online serving.
        """
        return pulumi.get(self, "vector_search_config")


class AwaitableGetFeatureViewResult(GetFeatureViewResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeatureViewResult(
            big_query_source=self.big_query_source,
            create_time=self.create_time,
            etag=self.etag,
            feature_registry_source=self.feature_registry_source,
            labels=self.labels,
            name=self.name,
            sync_config=self.sync_config,
            update_time=self.update_time,
            vector_search_config=self.vector_search_config)


def get_feature_view(feature_online_store_id: Optional[str] = None,
                     feature_view_id: Optional[str] = None,
                     location: Optional[str] = None,
                     project: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeatureViewResult:
    """
    Gets details of a single FeatureView.
    """
    __args__ = dict()
    __args__['featureOnlineStoreId'] = feature_online_store_id
    __args__['featureViewId'] = feature_view_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:aiplatform/v1beta1:getFeatureView', __args__, opts=opts, typ=GetFeatureViewResult).value

    return AwaitableGetFeatureViewResult(
        big_query_source=pulumi.get(__ret__, 'big_query_source'),
        create_time=pulumi.get(__ret__, 'create_time'),
        etag=pulumi.get(__ret__, 'etag'),
        feature_registry_source=pulumi.get(__ret__, 'feature_registry_source'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        sync_config=pulumi.get(__ret__, 'sync_config'),
        update_time=pulumi.get(__ret__, 'update_time'),
        vector_search_config=pulumi.get(__ret__, 'vector_search_config'))


@_utilities.lift_output_func(get_feature_view)
def get_feature_view_output(feature_online_store_id: Optional[pulumi.Input[str]] = None,
                            feature_view_id: Optional[pulumi.Input[str]] = None,
                            location: Optional[pulumi.Input[str]] = None,
                            project: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeatureViewResult]:
    """
    Gets details of a single FeatureView.
    """
    ...