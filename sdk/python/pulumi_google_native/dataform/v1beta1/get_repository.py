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
    'GetRepositoryResult',
    'AwaitableGetRepositoryResult',
    'get_repository',
    'get_repository_output',
]

@pulumi.output_type
class GetRepositoryResult:
    def __init__(__self__, display_name=None, git_remote_settings=None, labels=None, name=None, npmrc_environment_variables_secret_version=None, service_account=None, set_authenticated_user_admin=None, workspace_compilation_overrides=None):
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if git_remote_settings and not isinstance(git_remote_settings, dict):
            raise TypeError("Expected argument 'git_remote_settings' to be a dict")
        pulumi.set(__self__, "git_remote_settings", git_remote_settings)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if npmrc_environment_variables_secret_version and not isinstance(npmrc_environment_variables_secret_version, str):
            raise TypeError("Expected argument 'npmrc_environment_variables_secret_version' to be a str")
        pulumi.set(__self__, "npmrc_environment_variables_secret_version", npmrc_environment_variables_secret_version)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if set_authenticated_user_admin and not isinstance(set_authenticated_user_admin, bool):
            raise TypeError("Expected argument 'set_authenticated_user_admin' to be a bool")
        pulumi.set(__self__, "set_authenticated_user_admin", set_authenticated_user_admin)
        if workspace_compilation_overrides and not isinstance(workspace_compilation_overrides, dict):
            raise TypeError("Expected argument 'workspace_compilation_overrides' to be a dict")
        pulumi.set(__self__, "workspace_compilation_overrides", workspace_compilation_overrides)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. The repository's user-friendly name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="gitRemoteSettings")
    def git_remote_settings(self) -> 'outputs.GitRemoteSettingsResponse':
        """
        Optional. If set, configures this repository to be linked to a Git remote.
        """
        return pulumi.get(self, "git_remote_settings")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. Repository user labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The repository's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="npmrcEnvironmentVariablesSecretVersion")
    def npmrc_environment_variables_secret_version(self) -> str:
        """
        Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*/secrets/*/versions/*`. The file itself must be in a JSON format.
        """
        return pulumi.get(self, "npmrc_environment_variables_secret_version")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        Optional. The service account to run workflow invocations under.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="setAuthenticatedUserAdmin")
    def set_authenticated_user_admin(self) -> bool:
        """
        Optional. Input only. If set to true, the authenticated user will be granted the roles/dataform.admin role on the created repository. To modify access to the created repository later apply setIamPolicy from https://cloud.google.com/dataform/reference/rest#rest-resource:-v1beta1.projects.locations.repositories
        """
        return pulumi.get(self, "set_authenticated_user_admin")

    @property
    @pulumi.getter(name="workspaceCompilationOverrides")
    def workspace_compilation_overrides(self) -> 'outputs.WorkspaceCompilationOverridesResponse':
        """
        Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
        """
        return pulumi.get(self, "workspace_compilation_overrides")


class AwaitableGetRepositoryResult(GetRepositoryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRepositoryResult(
            display_name=self.display_name,
            git_remote_settings=self.git_remote_settings,
            labels=self.labels,
            name=self.name,
            npmrc_environment_variables_secret_version=self.npmrc_environment_variables_secret_version,
            service_account=self.service_account,
            set_authenticated_user_admin=self.set_authenticated_user_admin,
            workspace_compilation_overrides=self.workspace_compilation_overrides)


def get_repository(location: Optional[str] = None,
                   project: Optional[str] = None,
                   repository_id: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRepositoryResult:
    """
    Fetches a single Repository.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['repositoryId'] = repository_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataform/v1beta1:getRepository', __args__, opts=opts, typ=GetRepositoryResult).value

    return AwaitableGetRepositoryResult(
        display_name=pulumi.get(__ret__, 'display_name'),
        git_remote_settings=pulumi.get(__ret__, 'git_remote_settings'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        npmrc_environment_variables_secret_version=pulumi.get(__ret__, 'npmrc_environment_variables_secret_version'),
        service_account=pulumi.get(__ret__, 'service_account'),
        set_authenticated_user_admin=pulumi.get(__ret__, 'set_authenticated_user_admin'),
        workspace_compilation_overrides=pulumi.get(__ret__, 'workspace_compilation_overrides'))


@_utilities.lift_output_func(get_repository)
def get_repository_output(location: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          repository_id: Optional[pulumi.Input[str]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRepositoryResult]:
    """
    Fetches a single Repository.
    """
    ...
