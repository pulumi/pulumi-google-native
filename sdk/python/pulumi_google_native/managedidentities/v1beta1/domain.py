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

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 domain_name: pulumi.Input[str],
                 locations: pulumi.Input[Sequence[pulumi.Input[str]]],
                 reserved_ip_range: pulumi.Input[str],
                 admin: Optional[pulumi.Input[str]] = None,
                 audit_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[str] domain_name: Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] reserved_ip_range: The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        :param pulumi.Input[str] admin: Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
        :param pulumi.Input[bool] audit_logs_enabled: Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels that can contain user-provided metadata.
        """
        pulumi.set(__self__, "domain_name", domain_name)
        pulumi.set(__self__, "locations", locations)
        pulumi.set(__self__, "reserved_ip_range", reserved_ip_range)
        if admin is not None:
            pulumi.set(__self__, "admin", admin)
        if audit_logs_enabled is not None:
            pulumi.set(__self__, "audit_logs_enabled", audit_logs_enabled)
        if authorized_networks is not None:
            pulumi.set(__self__, "authorized_networks", authorized_networks)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Input[str]:
        """
        Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
        """
        return pulumi.get(self, "domain_name")

    @domain_name.setter
    def domain_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain_name", value)

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @locations.setter
    def locations(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "locations", value)

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> pulumi.Input[str]:
        """
        The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        """
        return pulumi.get(self, "reserved_ip_range")

    @reserved_ip_range.setter
    def reserved_ip_range(self, value: pulumi.Input[str]):
        pulumi.set(self, "reserved_ip_range", value)

    @property
    @pulumi.getter
    def admin(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
        """
        return pulumi.get(self, "admin")

    @admin.setter
    def admin(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin", value)

    @property
    @pulumi.getter(name="auditLogsEnabled")
    def audit_logs_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        """
        return pulumi.get(self, "audit_logs_enabled")

    @audit_logs_enabled.setter
    def audit_logs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "audit_logs_enabled", value)

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @authorized_networks.setter
    def authorized_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "authorized_networks", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Resource labels that can contain user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[str]] = None,
                 audit_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a Microsoft AD domain.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin: Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
        :param pulumi.Input[bool] audit_logs_enabled: Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] authorized_networks: Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
        :param pulumi.Input[str] domain_name: Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Resource labels that can contain user-provided metadata.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] locations: Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        :param pulumi.Input[str] reserved_ip_range: The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DomainArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Microsoft AD domain.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin: Optional[pulumi.Input[str]] = None,
                 audit_logs_enabled: Optional[pulumi.Input[bool]] = None,
                 authorized_networks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reserved_ip_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["admin"] = admin
            __props__.__dict__["audit_logs_enabled"] = audit_logs_enabled
            __props__.__dict__["authorized_networks"] = authorized_networks
            if domain_name is None and not opts.urn:
                raise TypeError("Missing required property 'domain_name'")
            __props__.__dict__["domain_name"] = domain_name
            __props__.__dict__["labels"] = labels
            if locations is None and not opts.urn:
                raise TypeError("Missing required property 'locations'")
            __props__.__dict__["locations"] = locations
            __props__.__dict__["project"] = project
            if reserved_ip_range is None and not opts.urn:
                raise TypeError("Missing required property 'reserved_ip_range'")
            __props__.__dict__["reserved_ip_range"] = reserved_ip_range
            __props__.__dict__["create_time"] = None
            __props__.__dict__["fqdn"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["status_message"] = None
            __props__.__dict__["trusts"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["domainName", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Domain, __self__).__init__(
            'google-native:managedidentities/v1beta1:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DomainArgs.__new__(DomainArgs)

        __props__.__dict__["admin"] = None
        __props__.__dict__["audit_logs_enabled"] = None
        __props__.__dict__["authorized_networks"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["domain_name"] = None
        __props__.__dict__["fqdn"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["locations"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["reserved_ip_range"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["status_message"] = None
        __props__.__dict__["trusts"] = None
        __props__.__dict__["update_time"] = None
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def admin(self) -> pulumi.Output[str]:
        """
        Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
        """
        return pulumi.get(self, "admin")

    @property
    @pulumi.getter(name="auditLogsEnabled")
    def audit_logs_enabled(self) -> pulumi.Output[bool]:
        """
        Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
        """
        return pulumi.get(self, "audit_logs_enabled")

    @property
    @pulumi.getter(name="authorizedNetworks")
    def authorized_networks(self) -> pulumi.Output[Sequence[str]]:
        """
        Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
        """
        return pulumi.get(self, "authorized_networks")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time the instance was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        Required. A domain name, e.g. mydomain.myorg.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric. * Must be unique within the customer project.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter
    def fqdn(self) -> pulumi.Output[str]:
        """
        The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Resource labels that can contain user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def locations(self) -> pulumi.Output[Sequence[str]]:
        """
        Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
        """
        return pulumi.get(self, "locations")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="reservedIpRange")
    def reserved_ip_range(self) -> pulumi.Output[str]:
        """
        The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
        """
        return pulumi.get(self, "reserved_ip_range")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of this domain.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="statusMessage")
    def status_message(self) -> pulumi.Output[str]:
        """
        Additional information about the current status of this domain, if available.
        """
        return pulumi.get(self, "status_message")

    @property
    @pulumi.getter
    def trusts(self) -> pulumi.Output[Sequence['outputs.TrustResponse']]:
        """
        The current trusts associated with the domain.
        """
        return pulumi.get(self, "trusts")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update time.
        """
        return pulumi.get(self, "update_time")

