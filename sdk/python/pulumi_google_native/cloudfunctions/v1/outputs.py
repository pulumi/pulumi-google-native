# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *

__all__ = [
    'AuditConfigResponse',
    'AuditLogConfigResponse',
    'BindingResponse',
    'EventTriggerResponse',
    'ExprResponse',
    'FailurePolicyResponse',
    'HttpsTriggerResponse',
    'RetryResponse',
    'SecretEnvVarResponse',
    'SecretVersionResponse',
    'SecretVolumeResponse',
    'SourceRepositoryResponse',
]

@pulumi.output_type
class AuditConfigResponse(dict):
    """
    Specifies the audit configuration for a service. The configuration determines which permission types are logged, and what identities, if any, are exempted from logging. An AuditConfig must have one or more AuditLogConfigs. If there are AuditConfigs for both `allServices` and a specific service, the union of the two AuditConfigs is used for that service: the log_types specified in each AuditConfig are enabled, and the exempted_members in each AuditLogConfig are exempted. Example Policy with multiple AuditConfigs: { "audit_configs": [ { "service": "allServices", "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type": "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com", "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type": "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ logging. It also exempts `jose@example.com` from DATA_READ logging, and `aliya@example.com` from DATA_WRITE logging.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "auditLogConfigs":
            suggest = "audit_log_configs"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuditConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 audit_log_configs: Sequence['outputs.AuditLogConfigResponse'],
                 service: str):
        """
        Specifies the audit configuration for a service. The configuration determines which permission types are logged, and what identities, if any, are exempted from logging. An AuditConfig must have one or more AuditLogConfigs. If there are AuditConfigs for both `allServices` and a specific service, the union of the two AuditConfigs is used for that service: the log_types specified in each AuditConfig are enabled, and the exempted_members in each AuditLogConfig are exempted. Example Policy with multiple AuditConfigs: { "audit_configs": [ { "service": "allServices", "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type": "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com", "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type": "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ logging. It also exempts `jose@example.com` from DATA_READ logging, and `aliya@example.com` from DATA_WRITE logging.
        :param Sequence['AuditLogConfigResponse'] audit_log_configs: The configuration for logging of each type of permission.
        :param str service: Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        pulumi.set(__self__, "audit_log_configs", audit_log_configs)
        pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="auditLogConfigs")
    def audit_log_configs(self) -> Sequence['outputs.AuditLogConfigResponse']:
        """
        The configuration for logging of each type of permission.
        """
        return pulumi.get(self, "audit_log_configs")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        """
        return pulumi.get(self, "service")


@pulumi.output_type
class AuditLogConfigResponse(dict):
    """
    Provides the configuration for logging a type of permissions. Example: { "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from DATA_READ logging.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "exemptedMembers":
            suggest = "exempted_members"
        elif key == "logType":
            suggest = "log_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuditLogConfigResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuditLogConfigResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuditLogConfigResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 exempted_members: Sequence[str],
                 log_type: str):
        """
        Provides the configuration for logging a type of permissions. Example: { "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from DATA_READ logging.
        :param Sequence[str] exempted_members: Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        :param str log_type: The log type that this config enables.
        """
        pulumi.set(__self__, "exempted_members", exempted_members)
        pulumi.set(__self__, "log_type", log_type)

    @property
    @pulumi.getter(name="exemptedMembers")
    def exempted_members(self) -> Sequence[str]:
        """
        Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        """
        return pulumi.get(self, "exempted_members")

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> str:
        """
        The log type that this config enables.
        """
        return pulumi.get(self, "log_type")


@pulumi.output_type
class BindingResponse(dict):
    """
    Associates `members`, or principals, with a `role`.
    """
    def __init__(__self__, *,
                 condition: 'outputs.ExprResponse',
                 members: Sequence[str],
                 role: str):
        """
        Associates `members`, or principals, with a `role`.
        :param 'ExprResponse' condition: The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        :param Sequence[str] members: Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        :param str role: Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        pulumi.set(__self__, "condition", condition)
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def condition(self) -> 'outputs.ExprResponse':
        """
        The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter
    def members(self) -> Sequence[str]:
        """
        Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. 
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
        """
        return pulumi.get(self, "role")


@pulumi.output_type
class EventTriggerResponse(dict):
    """
    Describes EventTrigger, used to request events be sent from another service.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "eventType":
            suggest = "event_type"
        elif key == "failurePolicy":
            suggest = "failure_policy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in EventTriggerResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        EventTriggerResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        EventTriggerResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 event_type: str,
                 failure_policy: 'outputs.FailurePolicyResponse',
                 resource: str,
                 service: str):
        """
        Describes EventTrigger, used to request events be sent from another service.
        :param str event_type: The type of event to observe. For example: `providers/cloud.storage/eventTypes/object.change` and `providers/cloud.pubsub/eventTypes/topic.publish`. Event types match pattern `providers/*/eventTypes/*.*`. The pattern contains: 1. namespace: For example, `cloud.storage` and `google.firebase.analytics`. 2. resource type: The type of resource on which event occurs. For example, the Google Cloud Storage API includes the type `object`. 3. action: The action that generates the event. For example, action for a Google Cloud Storage Object is 'change'. These parts are lower case.
        :param 'FailurePolicyResponse' failure_policy: Specifies policy for failed executions.
        :param str resource: The resource(s) from which to observe events, for example, `projects/_/buckets/myBucket`. Not all syntactically correct values are accepted by all services. For example: 1. The authorization model must support it. Google Cloud Functions only allows EventTriggers to be deployed that observe resources in the same project as the `CloudFunction`. 2. The resource type must match the pattern expected for an `event_type`. For example, an `EventTrigger` that has an `event_type` of "google.pubsub.topic.publish" should have a resource that matches Google Cloud Pub/Sub topics. Additionally, some services may support short names when creating an `EventTrigger`. These will always be returned in the normalized "long" format. See each *service's* documentation for supported formats.
        :param str service: The hostname of the service that should be observed. If no string is provided, the default service implementing the API will be used. For example, `storage.googleapis.com` is the default for all event types in the `google.storage` namespace.
        """
        pulumi.set(__self__, "event_type", event_type)
        pulumi.set(__self__, "failure_policy", failure_policy)
        pulumi.set(__self__, "resource", resource)
        pulumi.set(__self__, "service", service)

    @property
    @pulumi.getter(name="eventType")
    def event_type(self) -> str:
        """
        The type of event to observe. For example: `providers/cloud.storage/eventTypes/object.change` and `providers/cloud.pubsub/eventTypes/topic.publish`. Event types match pattern `providers/*/eventTypes/*.*`. The pattern contains: 1. namespace: For example, `cloud.storage` and `google.firebase.analytics`. 2. resource type: The type of resource on which event occurs. For example, the Google Cloud Storage API includes the type `object`. 3. action: The action that generates the event. For example, action for a Google Cloud Storage Object is 'change'. These parts are lower case.
        """
        return pulumi.get(self, "event_type")

    @property
    @pulumi.getter(name="failurePolicy")
    def failure_policy(self) -> 'outputs.FailurePolicyResponse':
        """
        Specifies policy for failed executions.
        """
        return pulumi.get(self, "failure_policy")

    @property
    @pulumi.getter
    def resource(self) -> str:
        """
        The resource(s) from which to observe events, for example, `projects/_/buckets/myBucket`. Not all syntactically correct values are accepted by all services. For example: 1. The authorization model must support it. Google Cloud Functions only allows EventTriggers to be deployed that observe resources in the same project as the `CloudFunction`. 2. The resource type must match the pattern expected for an `event_type`. For example, an `EventTrigger` that has an `event_type` of "google.pubsub.topic.publish" should have a resource that matches Google Cloud Pub/Sub topics. Additionally, some services may support short names when creating an `EventTrigger`. These will always be returned in the normalized "long" format. See each *service's* documentation for supported formats.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter
    def service(self) -> str:
        """
        The hostname of the service that should be observed. If no string is provided, the default service implementing the API will be used. For example, `storage.googleapis.com` is the default for all event types in the `google.storage` namespace.
        """
        return pulumi.get(self, "service")


@pulumi.output_type
class ExprResponse(dict):
    """
    Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
    """
    def __init__(__self__, *,
                 description: str,
                 expression: str,
                 location: str,
                 title: str):
        """
        Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
        :param str description: Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        :param str expression: Textual representation of an expression in Common Expression Language syntax.
        :param str location: Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        :param str title: Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def expression(self) -> str:
        """
        Textual representation of an expression in Common Expression Language syntax.
        """
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def title(self) -> str:
        """
        Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
        """
        return pulumi.get(self, "title")


@pulumi.output_type
class FailurePolicyResponse(dict):
    """
    Describes the policy in case of function's execution failure. If empty, then defaults to ignoring failures (i.e. not retrying them).
    """
    def __init__(__self__, *,
                 retry: 'outputs.RetryResponse'):
        """
        Describes the policy in case of function's execution failure. If empty, then defaults to ignoring failures (i.e. not retrying them).
        :param 'RetryResponse' retry: If specified, then the function will be retried in case of a failure.
        """
        pulumi.set(__self__, "retry", retry)

    @property
    @pulumi.getter
    def retry(self) -> 'outputs.RetryResponse':
        """
        If specified, then the function will be retried in case of a failure.
        """
        return pulumi.get(self, "retry")


@pulumi.output_type
class HttpsTriggerResponse(dict):
    """
    Describes HttpsTrigger, could be used to connect web hooks to function.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "securityLevel":
            suggest = "security_level"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in HttpsTriggerResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        HttpsTriggerResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        HttpsTriggerResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 security_level: str,
                 url: str):
        """
        Describes HttpsTrigger, could be used to connect web hooks to function.
        :param str security_level: The security level for the function.
        :param str url: The deployed url for the function.
        """
        pulumi.set(__self__, "security_level", security_level)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="securityLevel")
    def security_level(self) -> str:
        """
        The security level for the function.
        """
        return pulumi.get(self, "security_level")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The deployed url for the function.
        """
        return pulumi.get(self, "url")


@pulumi.output_type
class RetryResponse(dict):
    """
    Describes the retry policy in case of function's execution failure. A function execution will be retried on any failure. A failed execution will be retried up to 7 days with an exponential backoff (capped at 10 seconds). Retried execution is charged as any other execution.
    """
    def __init__(__self__):
        """
        Describes the retry policy in case of function's execution failure. A function execution will be retried on any failure. A failed execution will be retried up to 7 days with an exponential backoff (capped at 10 seconds). Retried execution is charged as any other execution.
        """
        pass


@pulumi.output_type
class SecretEnvVarResponse(dict):
    """
    Configuration for a secret environment variable. It has the information necessary to fetch the secret value from secret manager and expose it as an environment variable.
    """
    def __init__(__self__, *,
                 key: str,
                 project: str,
                 secret: str,
                 version: str):
        """
        Configuration for a secret environment variable. It has the information necessary to fetch the secret value from secret manager and expose it as an environment variable.
        :param str key: Name of the environment variable.
        :param str project: Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
        :param str secret: Name of the secret in secret manager (not the full resource name).
        :param str version: Version of the secret (version number or the string 'latest'). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new instances start.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "secret", secret)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        Name of the environment variable.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def secret(self) -> str:
        """
        Name of the secret in secret manager (not the full resource name).
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of the secret (version number or the string 'latest'). It is recommended to use a numeric version for secret environment variables as any updates to the secret value is not reflected until new instances start.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SecretVersionResponse(dict):
    """
    Configuration for a single version.
    """
    def __init__(__self__, *,
                 path: str,
                 version: str):
        """
        Configuration for a single version.
        :param str path: Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mount_path as '/etc/secrets' and path as `/secret_foo` would mount the secret value file at `/etc/secrets/secret_foo`.
        :param str version: Version of the secret (version number or the string 'latest'). It is preferrable to use `latest` version with secret volumes as secret value changes are reflected immediately.
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        Relative path of the file under the mount path where the secret value for this version will be fetched and made available. For example, setting the mount_path as '/etc/secrets' and path as `/secret_foo` would mount the secret value file at `/etc/secrets/secret_foo`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Version of the secret (version number or the string 'latest'). It is preferrable to use `latest` version with secret volumes as secret value changes are reflected immediately.
        """
        return pulumi.get(self, "version")


@pulumi.output_type
class SecretVolumeResponse(dict):
    """
    Configuration for a secret volume. It has the information necessary to fetch the secret value from secret manager and make it available as files mounted at the requested paths within the application container. Secret value is not a part of the configuration. Every filesystem read operation performs a lookup in secret manager to retrieve the secret value.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "mountPath":
            suggest = "mount_path"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretVolumeResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretVolumeResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretVolumeResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 mount_path: str,
                 project: str,
                 secret: str,
                 versions: Sequence['outputs.SecretVersionResponse']):
        """
        Configuration for a secret volume. It has the information necessary to fetch the secret value from secret manager and make it available as files mounted at the requested paths within the application container. Secret value is not a part of the configuration. Every filesystem read operation performs a lookup in secret manager to retrieve the secret value.
        :param str mount_path: The path within the container to mount the secret volume. For example, setting the mount_path as `/etc/secrets` would mount the secret value files under the `/etc/secrets` directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: /etc/secrets Restricted mount paths: /cloudsql, /dev/log, /pod, /proc, /var/log
        :param str project: Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
        :param str secret: Name of the secret in secret manager (not the full resource name).
        :param Sequence['SecretVersionResponse'] versions: List of secret versions to mount for this secret. If empty, the `latest` version of the secret will be made available in a file named after the secret under the mount point.
        """
        pulumi.set(__self__, "mount_path", mount_path)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "secret", secret)
        pulumi.set(__self__, "versions", versions)

    @property
    @pulumi.getter(name="mountPath")
    def mount_path(self) -> str:
        """
        The path within the container to mount the secret volume. For example, setting the mount_path as `/etc/secrets` would mount the secret value files under the `/etc/secrets` directory. This directory will also be completely shadowed and unavailable to mount any other secrets. Recommended mount paths: /etc/secrets Restricted mount paths: /cloudsql, /dev/log, /pod, /proc, /var/log
        """
        return pulumi.get(self, "mount_path")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        Project identifier (preferrably project number but can also be the project ID) of the project that contains the secret. If not set, it will be populated with the function's project assuming that the secret exists in the same project as of the function.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def secret(self) -> str:
        """
        Name of the secret in secret manager (not the full resource name).
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.SecretVersionResponse']:
        """
        List of secret versions to mount for this secret. If empty, the `latest` version of the secret will be made available in a file named after the secret under the mount point.
        """
        return pulumi.get(self, "versions")


@pulumi.output_type
class SourceRepositoryResponse(dict):
    """
    Describes SourceRepository, used to represent parameters related to source repository where a function is hosted.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "deployedUrl":
            suggest = "deployed_url"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SourceRepositoryResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SourceRepositoryResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SourceRepositoryResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 deployed_url: str,
                 url: str):
        """
        Describes SourceRepository, used to represent parameters related to source repository where a function is hosted.
        :param str deployed_url: The URL pointing to the hosted repository where the function were defined at the time of deployment. It always points to a specific commit in the format described above.
        :param str url: The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats: To refer to a specific commit: `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*` To refer to a moveable alias (branch): `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*` In particular, to refer to HEAD use `master` moveable alias. To refer to a specific fixed alias (tag): `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*` You may omit `paths/*` if you want to use the main directory.
        """
        pulumi.set(__self__, "deployed_url", deployed_url)
        pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="deployedUrl")
    def deployed_url(self) -> str:
        """
        The URL pointing to the hosted repository where the function were defined at the time of deployment. It always points to a specific commit in the format described above.
        """
        return pulumi.get(self, "deployed_url")

    @property
    @pulumi.getter
    def url(self) -> str:
        """
        The URL pointing to the hosted repository where the function is defined. There are supported Cloud Source Repository URLs in the following formats: To refer to a specific commit: `https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*` To refer to a moveable alias (branch): `https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*` In particular, to refer to HEAD use `master` moveable alias. To refer to a specific fixed alias (tag): `https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*` You may omit `paths/*` if you want to use the main directory.
        """
        return pulumi.get(self, "url")

