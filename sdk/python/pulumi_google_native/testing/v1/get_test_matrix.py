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
    'GetTestMatrixResult',
    'AwaitableGetTestMatrixResult',
    'get_test_matrix',
    'get_test_matrix_output',
]

@pulumi.output_type
class GetTestMatrixResult:
    def __init__(__self__, client_info=None, environment_matrix=None, fail_fast=None, flaky_test_attempts=None, invalid_matrix_details=None, outcome_summary=None, project=None, result_storage=None, state=None, test_executions=None, test_matrix_id=None, test_specification=None, timestamp=None):
        if client_info and not isinstance(client_info, dict):
            raise TypeError("Expected argument 'client_info' to be a dict")
        pulumi.set(__self__, "client_info", client_info)
        if environment_matrix and not isinstance(environment_matrix, dict):
            raise TypeError("Expected argument 'environment_matrix' to be a dict")
        pulumi.set(__self__, "environment_matrix", environment_matrix)
        if fail_fast and not isinstance(fail_fast, bool):
            raise TypeError("Expected argument 'fail_fast' to be a bool")
        pulumi.set(__self__, "fail_fast", fail_fast)
        if flaky_test_attempts and not isinstance(flaky_test_attempts, int):
            raise TypeError("Expected argument 'flaky_test_attempts' to be a int")
        pulumi.set(__self__, "flaky_test_attempts", flaky_test_attempts)
        if invalid_matrix_details and not isinstance(invalid_matrix_details, str):
            raise TypeError("Expected argument 'invalid_matrix_details' to be a str")
        pulumi.set(__self__, "invalid_matrix_details", invalid_matrix_details)
        if outcome_summary and not isinstance(outcome_summary, str):
            raise TypeError("Expected argument 'outcome_summary' to be a str")
        pulumi.set(__self__, "outcome_summary", outcome_summary)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if result_storage and not isinstance(result_storage, dict):
            raise TypeError("Expected argument 'result_storage' to be a dict")
        pulumi.set(__self__, "result_storage", result_storage)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if test_executions and not isinstance(test_executions, list):
            raise TypeError("Expected argument 'test_executions' to be a list")
        pulumi.set(__self__, "test_executions", test_executions)
        if test_matrix_id and not isinstance(test_matrix_id, str):
            raise TypeError("Expected argument 'test_matrix_id' to be a str")
        pulumi.set(__self__, "test_matrix_id", test_matrix_id)
        if test_specification and not isinstance(test_specification, dict):
            raise TypeError("Expected argument 'test_specification' to be a dict")
        pulumi.set(__self__, "test_specification", test_specification)
        if timestamp and not isinstance(timestamp, str):
            raise TypeError("Expected argument 'timestamp' to be a str")
        pulumi.set(__self__, "timestamp", timestamp)

    @property
    @pulumi.getter(name="clientInfo")
    def client_info(self) -> 'outputs.ClientInfoResponse':
        """
        Information about the client which invoked the test.
        """
        return pulumi.get(self, "client_info")

    @property
    @pulumi.getter(name="environmentMatrix")
    def environment_matrix(self) -> 'outputs.EnvironmentMatrixResponse':
        """
        The devices the tests are being executed on.
        """
        return pulumi.get(self, "environment_matrix")

    @property
    @pulumi.getter(name="failFast")
    def fail_fast(self) -> bool:
        """
        If true, only a single attempt at most will be made to run each execution/shard in the matrix. Flaky test attempts are not affected. Normally, 2 or more attempts are made if a potential infrastructure issue is detected. This feature is for latency sensitive workloads. The incidence of execution failures may be significantly greater for fail-fast matrices and support is more limited because of that expectation.
        """
        return pulumi.get(self, "fail_fast")

    @property
    @pulumi.getter(name="flakyTestAttempts")
    def flaky_test_attempts(self) -> int:
        """
        The number of times a TestExecution should be re-attempted if one or more of its test cases fail for any reason. The maximum number of reruns allowed is 10. Default is 0, which implies no reruns.
        """
        return pulumi.get(self, "flaky_test_attempts")

    @property
    @pulumi.getter(name="invalidMatrixDetails")
    def invalid_matrix_details(self) -> str:
        """
        Describes why the matrix is considered invalid. Only useful for matrices in the INVALID state.
        """
        return pulumi.get(self, "invalid_matrix_details")

    @property
    @pulumi.getter(name="outcomeSummary")
    def outcome_summary(self) -> str:
        """
        Output Only. The overall outcome of the test. Only set when the test matrix state is FINISHED.
        """
        return pulumi.get(self, "outcome_summary")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The cloud project that owns the test matrix.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="resultStorage")
    def result_storage(self) -> 'outputs.ResultStorageResponse':
        """
        Where the results for the matrix are written.
        """
        return pulumi.get(self, "result_storage")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Indicates the current progress of the test matrix.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="testExecutions")
    def test_executions(self) -> Sequence['outputs.TestExecutionResponse']:
        """
        The list of test executions that the service creates for this matrix.
        """
        return pulumi.get(self, "test_executions")

    @property
    @pulumi.getter(name="testMatrixId")
    def test_matrix_id(self) -> str:
        """
        Unique id set by the service.
        """
        return pulumi.get(self, "test_matrix_id")

    @property
    @pulumi.getter(name="testSpecification")
    def test_specification(self) -> 'outputs.TestSpecificationResponse':
        """
        How to run the test.
        """
        return pulumi.get(self, "test_specification")

    @property
    @pulumi.getter
    def timestamp(self) -> str:
        """
        The time this test matrix was initially created.
        """
        return pulumi.get(self, "timestamp")


class AwaitableGetTestMatrixResult(GetTestMatrixResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTestMatrixResult(
            client_info=self.client_info,
            environment_matrix=self.environment_matrix,
            fail_fast=self.fail_fast,
            flaky_test_attempts=self.flaky_test_attempts,
            invalid_matrix_details=self.invalid_matrix_details,
            outcome_summary=self.outcome_summary,
            project=self.project,
            result_storage=self.result_storage,
            state=self.state,
            test_executions=self.test_executions,
            test_matrix_id=self.test_matrix_id,
            test_specification=self.test_specification,
            timestamp=self.timestamp)


def get_test_matrix(project: Optional[str] = None,
                    test_matrix_id: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTestMatrixResult:
    """
    Checks the status of a test matrix. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Matrix does not exist
    """
    __args__ = dict()
    __args__['project'] = project
    __args__['testMatrixId'] = test_matrix_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:testing/v1:getTestMatrix', __args__, opts=opts, typ=GetTestMatrixResult).value

    return AwaitableGetTestMatrixResult(
        client_info=__ret__.client_info,
        environment_matrix=__ret__.environment_matrix,
        fail_fast=__ret__.fail_fast,
        flaky_test_attempts=__ret__.flaky_test_attempts,
        invalid_matrix_details=__ret__.invalid_matrix_details,
        outcome_summary=__ret__.outcome_summary,
        project=__ret__.project,
        result_storage=__ret__.result_storage,
        state=__ret__.state,
        test_executions=__ret__.test_executions,
        test_matrix_id=__ret__.test_matrix_id,
        test_specification=__ret__.test_specification,
        timestamp=__ret__.timestamp)


@_utilities.lift_output_func(get_test_matrix)
def get_test_matrix_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                           test_matrix_id: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTestMatrixResult]:
    """
    Checks the status of a test matrix. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Matrix does not exist
    """
    ...