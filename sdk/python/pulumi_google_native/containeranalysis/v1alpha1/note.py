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
from ._inputs import *

__all__ = ['NoteArgs', 'Note']

@pulumi.input_type
class NoteArgs:
    def __init__(__self__, *,
                 attestation_authority: Optional[pulumi.Input['AttestationAuthorityArgs']] = None,
                 base_image: Optional[pulumi.Input['BasisArgs']] = None,
                 build_type: Optional[pulumi.Input['BuildTypeArgs']] = None,
                 compliance: Optional[pulumi.Input['ComplianceNoteArgs']] = None,
                 deployable: Optional[pulumi.Input['DeployableArgs']] = None,
                 discovery: Optional[pulumi.Input['DiscoveryArgs']] = None,
                 dsse_attestation: Optional[pulumi.Input['DSSEAttestationNoteArgs']] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 long_description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 note_id: Optional[pulumi.Input[str]] = None,
                 package: Optional[pulumi.Input['PackageArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 related_url: Optional[pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]]] = None,
                 sbom: Optional[pulumi.Input['DocumentNoteArgs']] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 spdx_file: Optional[pulumi.Input['FileNoteArgs']] = None,
                 spdx_package: Optional[pulumi.Input['PackageInfoNoteArgs']] = None,
                 spdx_relationship: Optional[pulumi.Input['RelationshipNoteArgs']] = None,
                 upgrade: Optional[pulumi.Input['UpgradeNoteArgs']] = None,
                 vulnerability_type: Optional[pulumi.Input['VulnerabilityTypeArgs']] = None):
        """
        The set of arguments for constructing a Note resource.
        :param pulumi.Input['AttestationAuthorityArgs'] attestation_authority: A note describing an attestation role.
        :param pulumi.Input['BasisArgs'] base_image: A note describing a base image.
        :param pulumi.Input['BuildTypeArgs'] build_type: Build provenance type for a verifiable build.
        :param pulumi.Input['ComplianceNoteArgs'] compliance: A note describing a compliance check.
        :param pulumi.Input['DeployableArgs'] deployable: A note describing something that can be deployed.
        :param pulumi.Input['DiscoveryArgs'] discovery: A note describing a provider/analysis type.
        :param pulumi.Input['DSSEAttestationNoteArgs'] dsse_attestation: A note describing a dsse attestation note.
        :param pulumi.Input[str] expiration_time: Time of expiration for this note, null if note does not expire.
        :param pulumi.Input[str] long_description: A detailed description of this `Note`.
        :param pulumi.Input[str] name: The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        :param pulumi.Input[str] note_id: The ID to use for this note.
        :param pulumi.Input['PackageArgs'] package: A note describing a package hosted by various package managers.
        :param pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]] related_url: URLs associated with this note
        :param pulumi.Input['DocumentNoteArgs'] sbom: A note describing a software bill of materials.
        :param pulumi.Input[str] short_description: A one sentence description of this `Note`.
        :param pulumi.Input['FileNoteArgs'] spdx_file: A note describing an SPDX File.
        :param pulumi.Input['PackageInfoNoteArgs'] spdx_package: A note describing an SPDX Package.
        :param pulumi.Input['RelationshipNoteArgs'] spdx_relationship: A note describing a relationship between SPDX elements.
        :param pulumi.Input['UpgradeNoteArgs'] upgrade: A note describing an upgrade.
        :param pulumi.Input['VulnerabilityTypeArgs'] vulnerability_type: A package vulnerability type of note.
        """
        if attestation_authority is not None:
            pulumi.set(__self__, "attestation_authority", attestation_authority)
        if base_image is not None:
            pulumi.set(__self__, "base_image", base_image)
        if build_type is not None:
            pulumi.set(__self__, "build_type", build_type)
        if compliance is not None:
            pulumi.set(__self__, "compliance", compliance)
        if deployable is not None:
            pulumi.set(__self__, "deployable", deployable)
        if discovery is not None:
            pulumi.set(__self__, "discovery", discovery)
        if dsse_attestation is not None:
            pulumi.set(__self__, "dsse_attestation", dsse_attestation)
        if expiration_time is not None:
            pulumi.set(__self__, "expiration_time", expiration_time)
        if long_description is not None:
            pulumi.set(__self__, "long_description", long_description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if note_id is not None:
            pulumi.set(__self__, "note_id", note_id)
        if package is not None:
            pulumi.set(__self__, "package", package)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if related_url is not None:
            pulumi.set(__self__, "related_url", related_url)
        if sbom is not None:
            pulumi.set(__self__, "sbom", sbom)
        if short_description is not None:
            pulumi.set(__self__, "short_description", short_description)
        if spdx_file is not None:
            pulumi.set(__self__, "spdx_file", spdx_file)
        if spdx_package is not None:
            pulumi.set(__self__, "spdx_package", spdx_package)
        if spdx_relationship is not None:
            pulumi.set(__self__, "spdx_relationship", spdx_relationship)
        if upgrade is not None:
            pulumi.set(__self__, "upgrade", upgrade)
        if vulnerability_type is not None:
            pulumi.set(__self__, "vulnerability_type", vulnerability_type)

    @property
    @pulumi.getter(name="attestationAuthority")
    def attestation_authority(self) -> Optional[pulumi.Input['AttestationAuthorityArgs']]:
        """
        A note describing an attestation role.
        """
        return pulumi.get(self, "attestation_authority")

    @attestation_authority.setter
    def attestation_authority(self, value: Optional[pulumi.Input['AttestationAuthorityArgs']]):
        pulumi.set(self, "attestation_authority", value)

    @property
    @pulumi.getter(name="baseImage")
    def base_image(self) -> Optional[pulumi.Input['BasisArgs']]:
        """
        A note describing a base image.
        """
        return pulumi.get(self, "base_image")

    @base_image.setter
    def base_image(self, value: Optional[pulumi.Input['BasisArgs']]):
        pulumi.set(self, "base_image", value)

    @property
    @pulumi.getter(name="buildType")
    def build_type(self) -> Optional[pulumi.Input['BuildTypeArgs']]:
        """
        Build provenance type for a verifiable build.
        """
        return pulumi.get(self, "build_type")

    @build_type.setter
    def build_type(self, value: Optional[pulumi.Input['BuildTypeArgs']]):
        pulumi.set(self, "build_type", value)

    @property
    @pulumi.getter
    def compliance(self) -> Optional[pulumi.Input['ComplianceNoteArgs']]:
        """
        A note describing a compliance check.
        """
        return pulumi.get(self, "compliance")

    @compliance.setter
    def compliance(self, value: Optional[pulumi.Input['ComplianceNoteArgs']]):
        pulumi.set(self, "compliance", value)

    @property
    @pulumi.getter
    def deployable(self) -> Optional[pulumi.Input['DeployableArgs']]:
        """
        A note describing something that can be deployed.
        """
        return pulumi.get(self, "deployable")

    @deployable.setter
    def deployable(self, value: Optional[pulumi.Input['DeployableArgs']]):
        pulumi.set(self, "deployable", value)

    @property
    @pulumi.getter
    def discovery(self) -> Optional[pulumi.Input['DiscoveryArgs']]:
        """
        A note describing a provider/analysis type.
        """
        return pulumi.get(self, "discovery")

    @discovery.setter
    def discovery(self, value: Optional[pulumi.Input['DiscoveryArgs']]):
        pulumi.set(self, "discovery", value)

    @property
    @pulumi.getter(name="dsseAttestation")
    def dsse_attestation(self) -> Optional[pulumi.Input['DSSEAttestationNoteArgs']]:
        """
        A note describing a dsse attestation note.
        """
        return pulumi.get(self, "dsse_attestation")

    @dsse_attestation.setter
    def dsse_attestation(self, value: Optional[pulumi.Input['DSSEAttestationNoteArgs']]):
        pulumi.set(self, "dsse_attestation", value)

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of expiration for this note, null if note does not expire.
        """
        return pulumi.get(self, "expiration_time")

    @expiration_time.setter
    def expiration_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiration_time", value)

    @property
    @pulumi.getter(name="longDescription")
    def long_description(self) -> Optional[pulumi.Input[str]]:
        """
        A detailed description of this `Note`.
        """
        return pulumi.get(self, "long_description")

    @long_description.setter
    def long_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "long_description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="noteId")
    def note_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for this note.
        """
        return pulumi.get(self, "note_id")

    @note_id.setter
    def note_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "note_id", value)

    @property
    @pulumi.getter
    def package(self) -> Optional[pulumi.Input['PackageArgs']]:
        """
        A note describing a package hosted by various package managers.
        """
        return pulumi.get(self, "package")

    @package.setter
    def package(self, value: Optional[pulumi.Input['PackageArgs']]):
        pulumi.set(self, "package", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="relatedUrl")
    def related_url(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]]]:
        """
        URLs associated with this note
        """
        return pulumi.get(self, "related_url")

    @related_url.setter
    def related_url(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RelatedUrlArgs']]]]):
        pulumi.set(self, "related_url", value)

    @property
    @pulumi.getter
    def sbom(self) -> Optional[pulumi.Input['DocumentNoteArgs']]:
        """
        A note describing a software bill of materials.
        """
        return pulumi.get(self, "sbom")

    @sbom.setter
    def sbom(self, value: Optional[pulumi.Input['DocumentNoteArgs']]):
        pulumi.set(self, "sbom", value)

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> Optional[pulumi.Input[str]]:
        """
        A one sentence description of this `Note`.
        """
        return pulumi.get(self, "short_description")

    @short_description.setter
    def short_description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "short_description", value)

    @property
    @pulumi.getter(name="spdxFile")
    def spdx_file(self) -> Optional[pulumi.Input['FileNoteArgs']]:
        """
        A note describing an SPDX File.
        """
        return pulumi.get(self, "spdx_file")

    @spdx_file.setter
    def spdx_file(self, value: Optional[pulumi.Input['FileNoteArgs']]):
        pulumi.set(self, "spdx_file", value)

    @property
    @pulumi.getter(name="spdxPackage")
    def spdx_package(self) -> Optional[pulumi.Input['PackageInfoNoteArgs']]:
        """
        A note describing an SPDX Package.
        """
        return pulumi.get(self, "spdx_package")

    @spdx_package.setter
    def spdx_package(self, value: Optional[pulumi.Input['PackageInfoNoteArgs']]):
        pulumi.set(self, "spdx_package", value)

    @property
    @pulumi.getter(name="spdxRelationship")
    def spdx_relationship(self) -> Optional[pulumi.Input['RelationshipNoteArgs']]:
        """
        A note describing a relationship between SPDX elements.
        """
        return pulumi.get(self, "spdx_relationship")

    @spdx_relationship.setter
    def spdx_relationship(self, value: Optional[pulumi.Input['RelationshipNoteArgs']]):
        pulumi.set(self, "spdx_relationship", value)

    @property
    @pulumi.getter
    def upgrade(self) -> Optional[pulumi.Input['UpgradeNoteArgs']]:
        """
        A note describing an upgrade.
        """
        return pulumi.get(self, "upgrade")

    @upgrade.setter
    def upgrade(self, value: Optional[pulumi.Input['UpgradeNoteArgs']]):
        pulumi.set(self, "upgrade", value)

    @property
    @pulumi.getter(name="vulnerabilityType")
    def vulnerability_type(self) -> Optional[pulumi.Input['VulnerabilityTypeArgs']]:
        """
        A package vulnerability type of note.
        """
        return pulumi.get(self, "vulnerability_type")

    @vulnerability_type.setter
    def vulnerability_type(self, value: Optional[pulumi.Input['VulnerabilityTypeArgs']]):
        pulumi.set(self, "vulnerability_type", value)


class Note(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestation_authority: Optional[pulumi.Input[pulumi.InputType['AttestationAuthorityArgs']]] = None,
                 base_image: Optional[pulumi.Input[pulumi.InputType['BasisArgs']]] = None,
                 build_type: Optional[pulumi.Input[pulumi.InputType['BuildTypeArgs']]] = None,
                 compliance: Optional[pulumi.Input[pulumi.InputType['ComplianceNoteArgs']]] = None,
                 deployable: Optional[pulumi.Input[pulumi.InputType['DeployableArgs']]] = None,
                 discovery: Optional[pulumi.Input[pulumi.InputType['DiscoveryArgs']]] = None,
                 dsse_attestation: Optional[pulumi.Input[pulumi.InputType['DSSEAttestationNoteArgs']]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 long_description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 note_id: Optional[pulumi.Input[str]] = None,
                 package: Optional[pulumi.Input[pulumi.InputType['PackageArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 related_url: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RelatedUrlArgs']]]]] = None,
                 sbom: Optional[pulumi.Input[pulumi.InputType['DocumentNoteArgs']]] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 spdx_file: Optional[pulumi.Input[pulumi.InputType['FileNoteArgs']]] = None,
                 spdx_package: Optional[pulumi.Input[pulumi.InputType['PackageInfoNoteArgs']]] = None,
                 spdx_relationship: Optional[pulumi.Input[pulumi.InputType['RelationshipNoteArgs']]] = None,
                 upgrade: Optional[pulumi.Input[pulumi.InputType['UpgradeNoteArgs']]] = None,
                 vulnerability_type: Optional[pulumi.Input[pulumi.InputType['VulnerabilityTypeArgs']]] = None,
                 __props__=None):
        """
        Creates a new `Note`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AttestationAuthorityArgs']] attestation_authority: A note describing an attestation role.
        :param pulumi.Input[pulumi.InputType['BasisArgs']] base_image: A note describing a base image.
        :param pulumi.Input[pulumi.InputType['BuildTypeArgs']] build_type: Build provenance type for a verifiable build.
        :param pulumi.Input[pulumi.InputType['ComplianceNoteArgs']] compliance: A note describing a compliance check.
        :param pulumi.Input[pulumi.InputType['DeployableArgs']] deployable: A note describing something that can be deployed.
        :param pulumi.Input[pulumi.InputType['DiscoveryArgs']] discovery: A note describing a provider/analysis type.
        :param pulumi.Input[pulumi.InputType['DSSEAttestationNoteArgs']] dsse_attestation: A note describing a dsse attestation note.
        :param pulumi.Input[str] expiration_time: Time of expiration for this note, null if note does not expire.
        :param pulumi.Input[str] long_description: A detailed description of this `Note`.
        :param pulumi.Input[str] name: The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        :param pulumi.Input[str] note_id: The ID to use for this note.
        :param pulumi.Input[pulumi.InputType['PackageArgs']] package: A note describing a package hosted by various package managers.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RelatedUrlArgs']]]] related_url: URLs associated with this note
        :param pulumi.Input[pulumi.InputType['DocumentNoteArgs']] sbom: A note describing a software bill of materials.
        :param pulumi.Input[str] short_description: A one sentence description of this `Note`.
        :param pulumi.Input[pulumi.InputType['FileNoteArgs']] spdx_file: A note describing an SPDX File.
        :param pulumi.Input[pulumi.InputType['PackageInfoNoteArgs']] spdx_package: A note describing an SPDX Package.
        :param pulumi.Input[pulumi.InputType['RelationshipNoteArgs']] spdx_relationship: A note describing a relationship between SPDX elements.
        :param pulumi.Input[pulumi.InputType['UpgradeNoteArgs']] upgrade: A note describing an upgrade.
        :param pulumi.Input[pulumi.InputType['VulnerabilityTypeArgs']] vulnerability_type: A package vulnerability type of note.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NoteArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new `Note`.

        :param str resource_name: The name of the resource.
        :param NoteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NoteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestation_authority: Optional[pulumi.Input[pulumi.InputType['AttestationAuthorityArgs']]] = None,
                 base_image: Optional[pulumi.Input[pulumi.InputType['BasisArgs']]] = None,
                 build_type: Optional[pulumi.Input[pulumi.InputType['BuildTypeArgs']]] = None,
                 compliance: Optional[pulumi.Input[pulumi.InputType['ComplianceNoteArgs']]] = None,
                 deployable: Optional[pulumi.Input[pulumi.InputType['DeployableArgs']]] = None,
                 discovery: Optional[pulumi.Input[pulumi.InputType['DiscoveryArgs']]] = None,
                 dsse_attestation: Optional[pulumi.Input[pulumi.InputType['DSSEAttestationNoteArgs']]] = None,
                 expiration_time: Optional[pulumi.Input[str]] = None,
                 long_description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 note_id: Optional[pulumi.Input[str]] = None,
                 package: Optional[pulumi.Input[pulumi.InputType['PackageArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 related_url: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RelatedUrlArgs']]]]] = None,
                 sbom: Optional[pulumi.Input[pulumi.InputType['DocumentNoteArgs']]] = None,
                 short_description: Optional[pulumi.Input[str]] = None,
                 spdx_file: Optional[pulumi.Input[pulumi.InputType['FileNoteArgs']]] = None,
                 spdx_package: Optional[pulumi.Input[pulumi.InputType['PackageInfoNoteArgs']]] = None,
                 spdx_relationship: Optional[pulumi.Input[pulumi.InputType['RelationshipNoteArgs']]] = None,
                 upgrade: Optional[pulumi.Input[pulumi.InputType['UpgradeNoteArgs']]] = None,
                 vulnerability_type: Optional[pulumi.Input[pulumi.InputType['VulnerabilityTypeArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NoteArgs.__new__(NoteArgs)

            __props__.__dict__["attestation_authority"] = attestation_authority
            __props__.__dict__["base_image"] = base_image
            __props__.__dict__["build_type"] = build_type
            __props__.__dict__["compliance"] = compliance
            __props__.__dict__["deployable"] = deployable
            __props__.__dict__["discovery"] = discovery
            __props__.__dict__["dsse_attestation"] = dsse_attestation
            __props__.__dict__["expiration_time"] = expiration_time
            __props__.__dict__["long_description"] = long_description
            __props__.__dict__["name"] = name
            __props__.__dict__["note_id"] = note_id
            __props__.__dict__["package"] = package
            __props__.__dict__["project"] = project
            __props__.__dict__["related_url"] = related_url
            __props__.__dict__["sbom"] = sbom
            __props__.__dict__["short_description"] = short_description
            __props__.__dict__["spdx_file"] = spdx_file
            __props__.__dict__["spdx_package"] = spdx_package
            __props__.__dict__["spdx_relationship"] = spdx_relationship
            __props__.__dict__["upgrade"] = upgrade
            __props__.__dict__["vulnerability_type"] = vulnerability_type
            __props__.__dict__["create_time"] = None
            __props__.__dict__["kind"] = None
            __props__.__dict__["update_time"] = None
        super(Note, __self__).__init__(
            'google-native:containeranalysis/v1alpha1:Note',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Note':
        """
        Get an existing Note resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = NoteArgs.__new__(NoteArgs)

        __props__.__dict__["attestation_authority"] = None
        __props__.__dict__["base_image"] = None
        __props__.__dict__["build_type"] = None
        __props__.__dict__["compliance"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["deployable"] = None
        __props__.__dict__["discovery"] = None
        __props__.__dict__["dsse_attestation"] = None
        __props__.__dict__["expiration_time"] = None
        __props__.__dict__["kind"] = None
        __props__.__dict__["long_description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["package"] = None
        __props__.__dict__["related_url"] = None
        __props__.__dict__["sbom"] = None
        __props__.__dict__["short_description"] = None
        __props__.__dict__["spdx_file"] = None
        __props__.__dict__["spdx_package"] = None
        __props__.__dict__["spdx_relationship"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["upgrade"] = None
        __props__.__dict__["vulnerability_type"] = None
        return Note(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attestationAuthority")
    def attestation_authority(self) -> pulumi.Output['outputs.AttestationAuthorityResponse']:
        """
        A note describing an attestation role.
        """
        return pulumi.get(self, "attestation_authority")

    @property
    @pulumi.getter(name="baseImage")
    def base_image(self) -> pulumi.Output['outputs.BasisResponse']:
        """
        A note describing a base image.
        """
        return pulumi.get(self, "base_image")

    @property
    @pulumi.getter(name="buildType")
    def build_type(self) -> pulumi.Output['outputs.BuildTypeResponse']:
        """
        Build provenance type for a verifiable build.
        """
        return pulumi.get(self, "build_type")

    @property
    @pulumi.getter
    def compliance(self) -> pulumi.Output['outputs.ComplianceNoteResponse']:
        """
        A note describing a compliance check.
        """
        return pulumi.get(self, "compliance")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time this note was created. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def deployable(self) -> pulumi.Output['outputs.DeployableResponse']:
        """
        A note describing something that can be deployed.
        """
        return pulumi.get(self, "deployable")

    @property
    @pulumi.getter
    def discovery(self) -> pulumi.Output['outputs.DiscoveryResponse']:
        """
        A note describing a provider/analysis type.
        """
        return pulumi.get(self, "discovery")

    @property
    @pulumi.getter(name="dsseAttestation")
    def dsse_attestation(self) -> pulumi.Output['outputs.DSSEAttestationNoteResponse']:
        """
        A note describing a dsse attestation note.
        """
        return pulumi.get(self, "dsse_attestation")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> pulumi.Output[str]:
        """
        Time of expiration for this note, null if note does not expire.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[str]:
        """
        This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="longDescription")
    def long_description(self) -> pulumi.Output[str]:
        """
        A detailed description of this `Note`.
        """
        return pulumi.get(self, "long_description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def package(self) -> pulumi.Output['outputs.PackageResponse']:
        """
        A note describing a package hosted by various package managers.
        """
        return pulumi.get(self, "package")

    @property
    @pulumi.getter(name="relatedUrl")
    def related_url(self) -> pulumi.Output[Sequence['outputs.RelatedUrlResponse']]:
        """
        URLs associated with this note
        """
        return pulumi.get(self, "related_url")

    @property
    @pulumi.getter
    def sbom(self) -> pulumi.Output['outputs.DocumentNoteResponse']:
        """
        A note describing a software bill of materials.
        """
        return pulumi.get(self, "sbom")

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> pulumi.Output[str]:
        """
        A one sentence description of this `Note`.
        """
        return pulumi.get(self, "short_description")

    @property
    @pulumi.getter(name="spdxFile")
    def spdx_file(self) -> pulumi.Output['outputs.FileNoteResponse']:
        """
        A note describing an SPDX File.
        """
        return pulumi.get(self, "spdx_file")

    @property
    @pulumi.getter(name="spdxPackage")
    def spdx_package(self) -> pulumi.Output['outputs.PackageInfoNoteResponse']:
        """
        A note describing an SPDX Package.
        """
        return pulumi.get(self, "spdx_package")

    @property
    @pulumi.getter(name="spdxRelationship")
    def spdx_relationship(self) -> pulumi.Output['outputs.RelationshipNoteResponse']:
        """
        A note describing a relationship between SPDX elements.
        """
        return pulumi.get(self, "spdx_relationship")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time this note was last updated. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def upgrade(self) -> pulumi.Output['outputs.UpgradeNoteResponse']:
        """
        A note describing an upgrade.
        """
        return pulumi.get(self, "upgrade")

    @property
    @pulumi.getter(name="vulnerabilityType")
    def vulnerability_type(self) -> pulumi.Output['outputs.VulnerabilityTypeResponse']:
        """
        A package vulnerability type of note.
        """
        return pulumi.get(self, "vulnerability_type")
