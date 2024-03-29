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
    'GetHyperparameterTuningJobResult',
    'AwaitableGetHyperparameterTuningJobResult',
    'get_hyperparameter_tuning_job',
    'get_hyperparameter_tuning_job_output',
]

@pulumi.output_type
class GetHyperparameterTuningJobResult:
    def __init__(__self__, create_time=None, display_name=None, encryption_spec=None, end_time=None, error=None, labels=None, max_failed_trial_count=None, max_trial_count=None, name=None, parallel_trial_count=None, start_time=None, state=None, study_spec=None, trial_job_spec=None, trials=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if encryption_spec and not isinstance(encryption_spec, dict):
            raise TypeError("Expected argument 'encryption_spec' to be a dict")
        pulumi.set(__self__, "encryption_spec", encryption_spec)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if error and not isinstance(error, dict):
            raise TypeError("Expected argument 'error' to be a dict")
        pulumi.set(__self__, "error", error)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if max_failed_trial_count and not isinstance(max_failed_trial_count, int):
            raise TypeError("Expected argument 'max_failed_trial_count' to be a int")
        pulumi.set(__self__, "max_failed_trial_count", max_failed_trial_count)
        if max_trial_count and not isinstance(max_trial_count, int):
            raise TypeError("Expected argument 'max_trial_count' to be a int")
        pulumi.set(__self__, "max_trial_count", max_trial_count)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parallel_trial_count and not isinstance(parallel_trial_count, int):
            raise TypeError("Expected argument 'parallel_trial_count' to be a int")
        pulumi.set(__self__, "parallel_trial_count", parallel_trial_count)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if study_spec and not isinstance(study_spec, dict):
            raise TypeError("Expected argument 'study_spec' to be a dict")
        pulumi.set(__self__, "study_spec", study_spec)
        if trial_job_spec and not isinstance(trial_job_spec, dict):
            raise TypeError("Expected argument 'trial_job_spec' to be a dict")
        pulumi.set(__self__, "trial_job_spec", trial_job_spec)
        if trials and not isinstance(trials, list):
            raise TypeError("Expected argument 'trials' to be a list")
        pulumi.set(__self__, "trials", trials)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the HyperparameterTuningJob was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of the HyperparameterTuningJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="encryptionSpec")
    def encryption_spec(self) -> 'outputs.GoogleCloudAiplatformV1EncryptionSpecResponse':
        """
        Customer-managed encryption key options for a HyperparameterTuningJob. If this is set, then all resources created by the HyperparameterTuningJob will be encrypted with the provided encryption key.
        """
        return pulumi.get(self, "encryption_spec")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        Time when the HyperparameterTuningJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def error(self) -> 'outputs.GoogleRpcStatusResponse':
        """
        Only populated when job's state is JOB_STATE_FAILED or JOB_STATE_CANCELLED.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels with user-defined metadata to organize HyperparameterTuningJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="maxFailedTrialCount")
    def max_failed_trial_count(self) -> int:
        """
        The number of failed Trials that need to be seen before failing the HyperparameterTuningJob. If set to 0, Vertex AI decides how many Trials must fail before the whole job fails.
        """
        return pulumi.get(self, "max_failed_trial_count")

    @property
    @pulumi.getter(name="maxTrialCount")
    def max_trial_count(self) -> int:
        """
        The desired total number of Trials.
        """
        return pulumi.get(self, "max_trial_count")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the HyperparameterTuningJob.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parallelTrialCount")
    def parallel_trial_count(self) -> int:
        """
        The desired number of Trials to run in parallel.
        """
        return pulumi.get(self, "parallel_trial_count")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Time when the HyperparameterTuningJob for the first time entered the `JOB_STATE_RUNNING` state.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The detailed state of the job.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="studySpec")
    def study_spec(self) -> 'outputs.GoogleCloudAiplatformV1StudySpecResponse':
        """
        Study configuration of the HyperparameterTuningJob.
        """
        return pulumi.get(self, "study_spec")

    @property
    @pulumi.getter(name="trialJobSpec")
    def trial_job_spec(self) -> 'outputs.GoogleCloudAiplatformV1CustomJobSpecResponse':
        """
        The spec of a trial job. The same spec applies to the CustomJobs created in all the trials.
        """
        return pulumi.get(self, "trial_job_spec")

    @property
    @pulumi.getter
    def trials(self) -> Sequence['outputs.GoogleCloudAiplatformV1TrialResponse']:
        """
        Trials of the HyperparameterTuningJob.
        """
        return pulumi.get(self, "trials")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time when the HyperparameterTuningJob was most recently updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetHyperparameterTuningJobResult(GetHyperparameterTuningJobResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHyperparameterTuningJobResult(
            create_time=self.create_time,
            display_name=self.display_name,
            encryption_spec=self.encryption_spec,
            end_time=self.end_time,
            error=self.error,
            labels=self.labels,
            max_failed_trial_count=self.max_failed_trial_count,
            max_trial_count=self.max_trial_count,
            name=self.name,
            parallel_trial_count=self.parallel_trial_count,
            start_time=self.start_time,
            state=self.state,
            study_spec=self.study_spec,
            trial_job_spec=self.trial_job_spec,
            trials=self.trials,
            update_time=self.update_time)


def get_hyperparameter_tuning_job(hyperparameter_tuning_job_id: Optional[str] = None,
                                  location: Optional[str] = None,
                                  project: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHyperparameterTuningJobResult:
    """
    Gets a HyperparameterTuningJob
    """
    __args__ = dict()
    __args__['hyperparameterTuningJobId'] = hyperparameter_tuning_job_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:aiplatform/v1:getHyperparameterTuningJob', __args__, opts=opts, typ=GetHyperparameterTuningJobResult).value

    return AwaitableGetHyperparameterTuningJobResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        encryption_spec=pulumi.get(__ret__, 'encryption_spec'),
        end_time=pulumi.get(__ret__, 'end_time'),
        error=pulumi.get(__ret__, 'error'),
        labels=pulumi.get(__ret__, 'labels'),
        max_failed_trial_count=pulumi.get(__ret__, 'max_failed_trial_count'),
        max_trial_count=pulumi.get(__ret__, 'max_trial_count'),
        name=pulumi.get(__ret__, 'name'),
        parallel_trial_count=pulumi.get(__ret__, 'parallel_trial_count'),
        start_time=pulumi.get(__ret__, 'start_time'),
        state=pulumi.get(__ret__, 'state'),
        study_spec=pulumi.get(__ret__, 'study_spec'),
        trial_job_spec=pulumi.get(__ret__, 'trial_job_spec'),
        trials=pulumi.get(__ret__, 'trials'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_hyperparameter_tuning_job)
def get_hyperparameter_tuning_job_output(hyperparameter_tuning_job_id: Optional[pulumi.Input[str]] = None,
                                         location: Optional[pulumi.Input[str]] = None,
                                         project: Optional[pulumi.Input[Optional[str]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHyperparameterTuningJobResult]:
    """
    Gets a HyperparameterTuningJob
    """
    ...
