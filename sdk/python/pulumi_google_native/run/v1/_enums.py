# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'AuditLogConfigLogType',
    'DomainMappingSpecCertificateMode',
    'ResourceRecordType',
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


class DomainMappingSpecCertificateMode(str, Enum):
    """
    The mode of the certificate.
    """
    CERTIFICATE_MODE_UNSPECIFIED = "CERTIFICATE_MODE_UNSPECIFIED"
    NONE = "NONE"
    """
    Do not provision an HTTPS certificate.
    """
    AUTOMATIC = "AUTOMATIC"
    """
    Automatically provisions an HTTPS certificate via GoogleCA.
    """


class ResourceRecordType(str, Enum):
    """
    Resource record type. Example: `AAAA`.
    """
    RECORD_TYPE_UNSPECIFIED = "RECORD_TYPE_UNSPECIFIED"
    """
    An unknown resource record.
    """
    A = "A"
    """
    An A resource record. Data is an IPv4 address.
    """
    AAAA = "AAAA"
    """
    An AAAA resource record. Data is an IPv6 address.
    """
    CNAME = "CNAME"
    """
    A CNAME resource record. Data is a domain name to be aliased.
    """
