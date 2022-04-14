# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetWaiterResult',
    'AwaitableGetWaiterResult',
    'get_waiter',
    'get_waiter_output',
]

@pulumi.output_type
class GetWaiterResult:
    def __init__(__self__, create_time=None, done=None, error=None, failure=None, name=None, success=None, timeout=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if done and not isinstance(done, bool):
            raise TypeError("Expected argument 'done' to be a bool")
        pulumi.set(__self__, "done", done)
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if failure and not isinstance(failure, dict):
            raise TypeError("Expected argument 'failure' to be a dict")
        pulumi.set(__self__, "failure", failure)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if success and not isinstance(success, dict):
            raise TypeError("Expected argument 'success' to be a dict")
        pulumi.set(__self__, "success", success)
        if timeout and not isinstance(timeout, str):
            raise TypeError("Expected argument 'timeout' to be a str")
        pulumi.set(__self__, "timeout", timeout)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The instant at which this Waiter resource was created. Adding the value of `timeout` to this instant yields the timeout deadline for the waiter.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def done(self) -> bool:
        """
        If the value is `false`, it means the waiter is still waiting for one of its conditions to be met. If true, the waiter has finished. If the waiter finished due to a timeout or failure, `error` will be set.
        """
        return pulumi.get(self, "done")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.StatusResponse':
        """
        If the waiter ended due to a failure or timeout, this value will be set.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def failure(self) -> 'outputs.EndConditionResponse':
        """
        [Optional] The failure condition of this waiter. If this condition is met, `done` will be set to `true` and the `error` code will be set to `ABORTED`. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated. This value is optional; if no failure condition is set, the only failure scenario will be a timeout.
        """
        return pulumi.get(self, "failure")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Waiter resource, in the format: projects/[PROJECT_ID]/configs/[CONFIG_NAME]/waiters/[WAITER_NAME] The `[PROJECT_ID]` must be a valid Google Cloud project ID, the `[CONFIG_NAME]` must be a valid RuntimeConfig resource, the `[WAITER_NAME]` must match RFC 1035 segment specification, and the length of `[WAITER_NAME]` must be less than 64 bytes. After you create a Waiter resource, you cannot change the resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def success(self) -> 'outputs.EndConditionResponse':
        """
        [Required] The success condition. If this condition is met, `done` will be set to `true` and the `error` value will remain unset. The failure condition takes precedence over the success condition. If both conditions are met, a failure will be indicated.
        """
        return pulumi.get(self, "success")

    @property
    @pulumi.getter
    def timeout(self) -> str:
        """
        [Required] Specifies the timeout of the waiter in seconds, beginning from the instant that `waiters().create` method is called. If this time elapses before the success or failure conditions are met, the waiter fails and sets the `error` code to `DEADLINE_EXCEEDED`.
        """
        return pulumi.get(self, "timeout")


class AwaitableGetWaiterResult(GetWaiterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWaiterResult(
            create_time=self.create_time,
            done=self.done,
            error=self.error,
            failure=self.failure,
            name=self.name,
            success=self.success,
            timeout=self.timeout)


def get_waiter(config_id: Optional[str] = None,
               project: Optional[str] = None,
               waiter_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWaiterResult:
    """
    Gets information about a single waiter.
    """
    __args__ = dict()
    __args__['configId'] = config_id
    __args__['project'] = project
    __args__['waiterId'] = waiter_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:runtimeconfig/v1beta1:getWaiter', __args__, opts=opts, typ=GetWaiterResult).value

    return AwaitableGetWaiterResult(
        create_time=__ret__.create_time,
        done=__ret__.done,
        error=__ret__.error,
        failure=__ret__.failure,
        name=__ret__.name,
        success=__ret__.success,
        timeout=__ret__.timeout)


@_utilities.lift_output_func(get_waiter)
def get_waiter_output(config_id: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      waiter_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWaiterResult]:
    """
    Gets information about a single waiter.
    """
    ...