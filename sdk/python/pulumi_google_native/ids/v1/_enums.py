# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'EndpointSeverity',
]


class AuditLogConfigLogType(str, Enum):
    """
    The log type that this config enables.
    """
    LOG_TYPE_UNSPECIFIED = "LOG_TYPE_UNSPECIFIED"
    """
    Default case. Should never be this.
    """
    ADMIN_READ = "ADMIN_READ"
    """
    Admin reads. Example: CloudIAM getIamPolicy
    """
    DATA_WRITE = "DATA_WRITE"
    """
    Data writes. Example: CloudSQL Users create
    """
    DATA_READ = "DATA_READ"
    """
    Data reads. Example: CloudSQL Users list
    """


class EndpointSeverity(str, Enum):
    """
    Required. Lowest threat severity that this endpoint will alert on.
    """
    SEVERITY_UNSPECIFIED = "SEVERITY_UNSPECIFIED"
    """
    Not set.
    """
    INFORMATIONAL = "INFORMATIONAL"
    """
    Informational alerts.
    """
    LOW = "LOW"
    """
    Low severity alerts.
    """
    MEDIUM = "MEDIUM"
    """
    Medium severity alerts.
    """
    HIGH = "HIGH"
    """
    High severity alerts.
    """
    CRITICAL = "CRITICAL"
    """
    Critical severity alerts.
    """
