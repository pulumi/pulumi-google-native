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
    'GetDiscoveryConfigResult',
    'AwaitableGetDiscoveryConfigResult',
    'get_discovery_config',
    'get_discovery_config_output',
]

@pulumi.output_type
class GetDiscoveryConfigResult:
    def __init__(__self__, actions=None, create_time=None, display_name=None, errors=None, inspect_templates=None, last_run_time=None, name=None, org_config=None, status=None, targets=None, update_time=None):
        if actions and not isinstance(actions, list):
            raise TypeError("Expected argument 'actions' to be a list")
        pulumi.set(__self__, "actions", actions)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if errors and not isinstance(errors, list):
            raise TypeError("Expected argument 'errors' to be a list")
        pulumi.set(__self__, "errors", errors)
        if inspect_templates and not isinstance(inspect_templates, list):
            raise TypeError("Expected argument 'inspect_templates' to be a list")
        pulumi.set(__self__, "inspect_templates", inspect_templates)
        if last_run_time and not isinstance(last_run_time, str):
            raise TypeError("Expected argument 'last_run_time' to be a str")
        pulumi.set(__self__, "last_run_time", last_run_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org_config and not isinstance(org_config, dict):
            raise TypeError("Expected argument 'org_config' to be a dict")
        pulumi.set(__self__, "org_config", org_config)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if targets and not isinstance(targets, list):
            raise TypeError("Expected argument 'targets' to be a list")
        pulumi.set(__self__, "targets", targets)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def actions(self) -> Sequence['outputs.GooglePrivacyDlpV2DataProfileActionResponse']:
        """
        Actions to execute at the completion of scanning.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation timestamp of a DiscoveryConfig.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Display name (max 100 chars)
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def errors(self) -> Sequence['outputs.GooglePrivacyDlpV2ErrorResponse']:
        """
        A stream of errors encountered when the config was activated. Repeated errors may result in the config automatically being paused. Output only field. Will return the last 100 errors. Whenever the config is modified this list will be cleared.
        """
        return pulumi.get(self, "errors")

    @property
    @pulumi.getter(name="inspectTemplates")
    def inspect_templates(self) -> Sequence[str]:
        """
        Detection logic for profile generation. Not all template features are used by Discovery. FindingLimits, include_quote and exclude_info_types have no impact on Discovery. Multiple templates may be provided if there is data in multiple regions. At most one template must be specified per-region (including "global"). Each region is scanned using the applicable template. If no region-specific template is specified, but a "global" template is specified, it will be copied to that region and used instead. If no global or region-specific template is provided for a region with data, that region's data will not be scanned. For more information, see https://cloud.google.com/dlp/docs/data-profiles#data-residency.
        """
        return pulumi.get(self, "inspect_templates")

    @property
    @pulumi.getter(name="lastRunTime")
    def last_run_time(self) -> str:
        """
        The timestamp of the last time this config was executed.
        """
        return pulumi.get(self, "last_run_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique resource name for the DiscoveryConfig, assigned by the service when the DiscoveryConfig is created, for example `projects/dlp-test-project/locations/global/discoveryConfigs/53234423`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgConfig")
    def org_config(self) -> 'outputs.GooglePrivacyDlpV2OrgConfigResponse':
        """
        Only set when the parent is an org.
        """
        return pulumi.get(self, "org_config")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        A status for this configuration.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def targets(self) -> Sequence['outputs.GooglePrivacyDlpV2DiscoveryTargetResponse']:
        """
        Target to match against for determining what to scan and how frequently.
        """
        return pulumi.get(self, "targets")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of a DiscoveryConfig.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDiscoveryConfigResult(GetDiscoveryConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDiscoveryConfigResult(
            actions=self.actions,
            create_time=self.create_time,
            display_name=self.display_name,
            errors=self.errors,
            inspect_templates=self.inspect_templates,
            last_run_time=self.last_run_time,
            name=self.name,
            org_config=self.org_config,
            status=self.status,
            targets=self.targets,
            update_time=self.update_time)


def get_discovery_config(discovery_config_id: Optional[str] = None,
                         location: Optional[str] = None,
                         project: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDiscoveryConfigResult:
    """
    Gets a discovery configuration.
    """
    __args__ = dict()
    __args__['discoveryConfigId'] = discovery_config_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dlp/v2:getDiscoveryConfig', __args__, opts=opts, typ=GetDiscoveryConfigResult).value

    return AwaitableGetDiscoveryConfigResult(
        actions=pulumi.get(__ret__, 'actions'),
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        errors=pulumi.get(__ret__, 'errors'),
        inspect_templates=pulumi.get(__ret__, 'inspect_templates'),
        last_run_time=pulumi.get(__ret__, 'last_run_time'),
        name=pulumi.get(__ret__, 'name'),
        org_config=pulumi.get(__ret__, 'org_config'),
        status=pulumi.get(__ret__, 'status'),
        targets=pulumi.get(__ret__, 'targets'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_discovery_config)
def get_discovery_config_output(discovery_config_id: Optional[pulumi.Input[str]] = None,
                                location: Optional[pulumi.Input[str]] = None,
                                project: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDiscoveryConfigResult]:
    """
    Gets a discovery configuration.
    """
    ...