# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from .. import _utilities
import typing

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_google_native.dataflow.v1b3 as __v1b3
    v1b3 = __v1b3
else:
    v1b3 = _utilities.lazy_import('pulumi_google_native.dataflow.v1b3')

