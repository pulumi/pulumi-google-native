# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'GoogleDatastoreAdminV1IndexedPropertyArgs',
]

@pulumi.input_type
class GoogleDatastoreAdminV1IndexedPropertyArgs:
    def __init__(__self__, *,
                 direction: pulumi.Input['GoogleDatastoreAdminV1IndexedPropertyDirection'],
                 name: pulumi.Input[str]):
        """
        A property of an index.
        :param pulumi.Input['GoogleDatastoreAdminV1IndexedPropertyDirection'] direction: The indexed property's direction. Must not be DIRECTION_UNSPECIFIED.
        :param pulumi.Input[str] name: The property name to index.
        """
        pulumi.set(__self__, "direction", direction)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def direction(self) -> pulumi.Input['GoogleDatastoreAdminV1IndexedPropertyDirection']:
        """
        The indexed property's direction. Must not be DIRECTION_UNSPECIFIED.
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: pulumi.Input['GoogleDatastoreAdminV1IndexedPropertyDirection']):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The property name to index.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

