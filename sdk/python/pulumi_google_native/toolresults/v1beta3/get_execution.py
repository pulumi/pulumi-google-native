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
    'GetExecutionResult',
    'AwaitableGetExecutionResult',
    'get_execution',
    'get_execution_output',
]

@pulumi.output_type
class GetExecutionResult:
    def __init__(__self__, completion_time=None, creation_time=None, dimension_definitions=None, execution_id=None, outcome=None, specification=None, state=None, test_execution_matrix_id=None):
        if completion_time and not isinstance(completion_time, dict):
            raise TypeError("Expected argument 'completion_time' to be a dict")
        pulumi.set(__self__, "completion_time", completion_time)
        if creation_time and not isinstance(creation_time, dict):
            raise TypeError("Expected argument 'creation_time' to be a dict")
        pulumi.set(__self__, "creation_time", creation_time)
        if dimension_definitions and not isinstance(dimension_definitions, list):
            raise TypeError("Expected argument 'dimension_definitions' to be a list")
        pulumi.set(__self__, "dimension_definitions", dimension_definitions)
        if execution_id and not isinstance(execution_id, str):
            raise TypeError("Expected argument 'execution_id' to be a str")
        pulumi.set(__self__, "execution_id", execution_id)
        if outcome and not isinstance(outcome, dict):
            raise TypeError("Expected argument 'outcome' to be a dict")
        pulumi.set(__self__, "outcome", outcome)
        if specification and not isinstance(specification, dict):
            raise TypeError("Expected argument 'specification' to be a dict")
        pulumi.set(__self__, "specification", specification)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if test_execution_matrix_id and not isinstance(test_execution_matrix_id, str):
            raise TypeError("Expected argument 'test_execution_matrix_id' to be a str")
        pulumi.set(__self__, "test_execution_matrix_id", test_execution_matrix_id)

    @property
    @pulumi.getter(name="completionTime")
    def completion_time(self) -> 'outputs.TimestampResponse':
        """
        The time when the Execution status transitioned to COMPLETE. This value will be set automatically when state transitions to COMPLETE. - In response: set if the execution state is COMPLETE. - In create/update request: never set
        """
        return pulumi.get(self, "completion_time")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> 'outputs.TimestampResponse':
        """
        The time when the Execution was created. This value will be set automatically when CreateExecution is called. - In response: always set - In create/update request: never set
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="dimensionDefinitions")
    def dimension_definitions(self) -> Sequence['outputs.MatrixDimensionDefinitionResponse']:
        """
        The dimensions along which different steps in this execution may vary. This must remain fixed over the life of the execution. Returns INVALID_ARGUMENT if this field is set in an update request. Returns INVALID_ARGUMENT if the same name occurs in more than one dimension_definition. Returns INVALID_ARGUMENT if the size of the list is over 100. - In response: present if set by create - In create request: optional - In update request: never set
        """
        return pulumi.get(self, "dimension_definitions")

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> str:
        """
        A unique identifier within a History for this Execution. Returns INVALID_ARGUMENT if this field is set or overwritten by the caller. - In response always set - In create/update request: never set
        """
        return pulumi.get(self, "execution_id")

    @property
    @pulumi.getter
    def outcome(self) -> 'outputs.OutcomeResponse':
        """
        Classify the result, for example into SUCCESS or FAILURE - In response: present if set by create/update request - In create/update request: optional
        """
        return pulumi.get(self, "outcome")

    @property
    @pulumi.getter
    def specification(self) -> 'outputs.SpecificationResponse':
        """
        Lightweight information about execution request. - In response: present if set by create - In create: optional - In update: optional
        """
        return pulumi.get(self, "specification")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The initial state is IN_PROGRESS. The only legal state transitions is from IN_PROGRESS to COMPLETE. A PRECONDITION_FAILED will be returned if an invalid transition is requested. The state can only be set to COMPLETE once. A FAILED_PRECONDITION will be returned if the state is set to COMPLETE multiple times. If the state is set to COMPLETE, all the in-progress steps within the execution will be set as COMPLETE. If the outcome of the step is not set, the outcome will be set to INCONCLUSIVE. - In response always set - In create/update request: optional
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="testExecutionMatrixId")
    def test_execution_matrix_id(self) -> str:
        """
        TestExecution Matrix ID that the TestExecutionService uses. - In response: present if set by create - In create: optional - In update: never set
        """
        return pulumi.get(self, "test_execution_matrix_id")


class AwaitableGetExecutionResult(GetExecutionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExecutionResult(
            completion_time=self.completion_time,
            creation_time=self.creation_time,
            dimension_definitions=self.dimension_definitions,
            execution_id=self.execution_id,
            outcome=self.outcome,
            specification=self.specification,
            state=self.state,
            test_execution_matrix_id=self.test_execution_matrix_id)


def get_execution(execution_id: Optional[str] = None,
                  history_id: Optional[str] = None,
                  project: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExecutionResult:
    """
    Gets an Execution. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Execution does not exist
    """
    __args__ = dict()
    __args__['executionId'] = execution_id
    __args__['historyId'] = history_id
    __args__['project'] = project
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:toolresults/v1beta3:getExecution', __args__, opts=opts, typ=GetExecutionResult).value

    return AwaitableGetExecutionResult(
        completion_time=__ret__.completion_time,
        creation_time=__ret__.creation_time,
        dimension_definitions=__ret__.dimension_definitions,
        execution_id=__ret__.execution_id,
        outcome=__ret__.outcome,
        specification=__ret__.specification,
        state=__ret__.state,
        test_execution_matrix_id=__ret__.test_execution_matrix_id)


@_utilities.lift_output_func(get_execution)
def get_execution_output(execution_id: Optional[pulumi.Input[str]] = None,
                         history_id: Optional[pulumi.Input[str]] = None,
                         project: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetExecutionResult]:
    """
    Gets an Execution. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to write to project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Execution does not exist
    """
    ...