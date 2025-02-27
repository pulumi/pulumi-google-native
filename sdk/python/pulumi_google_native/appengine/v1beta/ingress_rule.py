# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = ['IngressRuleArgs', 'IngressRule']

@pulumi.input_type
class IngressRuleArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[str],
                 action: Optional[pulumi.Input['IngressRuleAction']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 source_range: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IngressRule resource.
        :param pulumi.Input['IngressRuleAction'] action: The action to take on matched requests.
        :param pulumi.Input[str] description: An optional string description of this rule. This field has a maximum length of 400 characters.
        :param pulumi.Input[int] priority: A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
        :param pulumi.Input[str] source_range: IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
        """
        pulumi.set(__self__, "app_id", app_id)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if source_range is not None:
            pulumi.set(__self__, "source_range", source_range)

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "app_id", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input['IngressRuleAction']]:
        """
        The action to take on matched requests.
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input['IngressRuleAction']]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional string description of this rule. This field has a maximum length of 400 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="sourceRange")
    def source_range(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
        """
        return pulumi.get(self, "source_range")

    @source_range.setter
    def source_range(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_range", value)


class IngressRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['IngressRuleAction']] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 source_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a firewall rule for the application.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input['IngressRuleAction'] action: The action to take on matched requests.
        :param pulumi.Input[str] description: An optional string description of this rule. This field has a maximum length of 400 characters.
        :param pulumi.Input[int] priority: A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
        :param pulumi.Input[str] source_range: IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IngressRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a firewall rule for the application.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param IngressRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IngressRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input['IngressRuleAction']] = None,
                 app_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 source_range: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IngressRuleArgs.__new__(IngressRuleArgs)

            __props__.__dict__["action"] = action
            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["description"] = description
            __props__.__dict__["priority"] = priority
            __props__.__dict__["source_range"] = source_range
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["appId"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(IngressRule, __self__).__init__(
            'google-native:appengine/v1beta:IngressRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'IngressRule':
        """
        Get an existing IngressRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = IngressRuleArgs.__new__(IngressRuleArgs)

        __props__.__dict__["action"] = None
        __props__.__dict__["app_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["priority"] = None
        __props__.__dict__["source_range"] = None
        return IngressRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        The action to take on matched requests.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "app_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        An optional string description of this rule. This field has a maximum length of 400 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        A positive integer between 1, Int32.MaxValue-1 that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.A default rule at priority Int32.MaxValue matches all IPv4 and IPv6 traffic when no previous rule matches. Only the action of this rule can be modified by the user.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="sourceRange")
    def source_range(self) -> pulumi.Output[str]:
        """
        IP address or range, defined using CIDR notation, of requests that this rule applies to. You can use the wildcard character "*" to match all IPs equivalent to "0/0" and "::/0" together. Examples: 192.168.1.1 or 192.168.0.0/16 or 2001:db8::/32 or 2001:0db8:0000:0042:0000:8a2e:0370:7334. Truncation will be silently performed on addresses which are not properly truncated. For example, 1.2.3.4/24 is accepted as the same address as 1.2.3.0/24. Similarly, for IPv6, 2001:db8::1/32 is accepted as the same address as 2001:db8::/32.
        """
        return pulumi.get(self, "source_range")

