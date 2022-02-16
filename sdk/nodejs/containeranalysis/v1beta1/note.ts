// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new note.
 * Auto-naming is currently not supported for this resource.
 */
export class Note extends pulumi.CustomResource {
    /**
     * Get an existing Note resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Note {
        return new Note(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:containeranalysis/v1beta1:Note';

    /**
     * Returns true if the given object is an instance of Note.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Note {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Note.__pulumiType;
    }

    /**
     * A note describing an attestation role.
     */
    public readonly attestationAuthority!: pulumi.Output<outputs.containeranalysis.v1beta1.AuthorityResponse>;
    /**
     * A note describing a base image.
     */
    public readonly baseImage!: pulumi.Output<outputs.containeranalysis.v1beta1.BasisResponse>;
    /**
     * A note describing build provenance for a verifiable build.
     */
    public readonly build!: pulumi.Output<outputs.containeranalysis.v1beta1.BuildResponse>;
    /**
     * The time this note was created. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A note describing something that can be deployed.
     */
    public readonly deployable!: pulumi.Output<outputs.containeranalysis.v1beta1.DeployableResponse>;
    /**
     * A note describing the initial analysis of a resource.
     */
    public readonly discovery!: pulumi.Output<outputs.containeranalysis.v1beta1.DiscoveryResponse>;
    /**
     * Time of expiration for this note. Empty if note does not expire.
     */
    public readonly expirationTime!: pulumi.Output<string>;
    /**
     * A note describing an in-toto link.
     */
    public readonly intoto!: pulumi.Output<outputs.containeranalysis.v1beta1.InTotoResponse>;
    /**
     * The type of analysis. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A detailed description of this note.
     */
    public readonly longDescription!: pulumi.Output<string>;
    /**
     * The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A note describing a package hosted by various package managers.
     */
    public readonly package!: pulumi.Output<outputs.containeranalysis.v1beta1.PackageResponse>;
    /**
     * Other notes related to this note.
     */
    public readonly relatedNoteNames!: pulumi.Output<string[]>;
    /**
     * URLs associated with this note.
     */
    public readonly relatedUrl!: pulumi.Output<outputs.containeranalysis.v1beta1.RelatedUrlResponse[]>;
    /**
     * A note describing a software bill of materials.
     */
    public readonly sbom!: pulumi.Output<outputs.containeranalysis.v1beta1.DocumentNoteResponse>;
    /**
     * A one sentence description of this note.
     */
    public readonly shortDescription!: pulumi.Output<string>;
    /**
     * A note describing an SPDX File.
     */
    public readonly spdxFile!: pulumi.Output<outputs.containeranalysis.v1beta1.FileNoteResponse>;
    /**
     * A note describing an SPDX Package.
     */
    public readonly spdxPackage!: pulumi.Output<outputs.containeranalysis.v1beta1.PackageInfoNoteResponse>;
    /**
     * A note describing an SPDX File.
     */
    public readonly spdxRelationship!: pulumi.Output<outputs.containeranalysis.v1beta1.RelationshipNoteResponse>;
    /**
     * The time this note was last updated. This field can be used as a filter in list requests.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * A note describing a package vulnerability.
     */
    public readonly vulnerability!: pulumi.Output<outputs.containeranalysis.v1beta1.VulnerabilityResponse>;

    /**
     * Create a Note resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NoteArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.noteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'noteId'");
            }
            resourceInputs["attestationAuthority"] = args ? args.attestationAuthority : undefined;
            resourceInputs["baseImage"] = args ? args.baseImage : undefined;
            resourceInputs["build"] = args ? args.build : undefined;
            resourceInputs["deployable"] = args ? args.deployable : undefined;
            resourceInputs["discovery"] = args ? args.discovery : undefined;
            resourceInputs["expirationTime"] = args ? args.expirationTime : undefined;
            resourceInputs["intoto"] = args ? args.intoto : undefined;
            resourceInputs["longDescription"] = args ? args.longDescription : undefined;
            resourceInputs["noteId"] = args ? args.noteId : undefined;
            resourceInputs["package"] = args ? args.package : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["relatedNoteNames"] = args ? args.relatedNoteNames : undefined;
            resourceInputs["relatedUrl"] = args ? args.relatedUrl : undefined;
            resourceInputs["sbom"] = args ? args.sbom : undefined;
            resourceInputs["shortDescription"] = args ? args.shortDescription : undefined;
            resourceInputs["spdxFile"] = args ? args.spdxFile : undefined;
            resourceInputs["spdxPackage"] = args ? args.spdxPackage : undefined;
            resourceInputs["spdxRelationship"] = args ? args.spdxRelationship : undefined;
            resourceInputs["vulnerability"] = args ? args.vulnerability : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["attestationAuthority"] = undefined /*out*/;
            resourceInputs["baseImage"] = undefined /*out*/;
            resourceInputs["build"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deployable"] = undefined /*out*/;
            resourceInputs["discovery"] = undefined /*out*/;
            resourceInputs["expirationTime"] = undefined /*out*/;
            resourceInputs["intoto"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["longDescription"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["package"] = undefined /*out*/;
            resourceInputs["relatedNoteNames"] = undefined /*out*/;
            resourceInputs["relatedUrl"] = undefined /*out*/;
            resourceInputs["sbom"] = undefined /*out*/;
            resourceInputs["shortDescription"] = undefined /*out*/;
            resourceInputs["spdxFile"] = undefined /*out*/;
            resourceInputs["spdxPackage"] = undefined /*out*/;
            resourceInputs["spdxRelationship"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vulnerability"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Note.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Note resource.
 */
export interface NoteArgs {
    /**
     * A note describing an attestation role.
     */
    attestationAuthority?: pulumi.Input<inputs.containeranalysis.v1beta1.AuthorityArgs>;
    /**
     * A note describing a base image.
     */
    baseImage?: pulumi.Input<inputs.containeranalysis.v1beta1.BasisArgs>;
    /**
     * A note describing build provenance for a verifiable build.
     */
    build?: pulumi.Input<inputs.containeranalysis.v1beta1.BuildArgs>;
    /**
     * A note describing something that can be deployed.
     */
    deployable?: pulumi.Input<inputs.containeranalysis.v1beta1.DeployableArgs>;
    /**
     * A note describing the initial analysis of a resource.
     */
    discovery?: pulumi.Input<inputs.containeranalysis.v1beta1.DiscoveryArgs>;
    /**
     * Time of expiration for this note. Empty if note does not expire.
     */
    expirationTime?: pulumi.Input<string>;
    /**
     * A note describing an in-toto link.
     */
    intoto?: pulumi.Input<inputs.containeranalysis.v1beta1.InTotoArgs>;
    /**
     * A detailed description of this note.
     */
    longDescription?: pulumi.Input<string>;
    noteId: pulumi.Input<string>;
    /**
     * A note describing a package hosted by various package managers.
     */
    package?: pulumi.Input<inputs.containeranalysis.v1beta1.PackageArgs>;
    project?: pulumi.Input<string>;
    /**
     * Other notes related to this note.
     */
    relatedNoteNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * URLs associated with this note.
     */
    relatedUrl?: pulumi.Input<pulumi.Input<inputs.containeranalysis.v1beta1.RelatedUrlArgs>[]>;
    /**
     * A note describing a software bill of materials.
     */
    sbom?: pulumi.Input<inputs.containeranalysis.v1beta1.DocumentNoteArgs>;
    /**
     * A one sentence description of this note.
     */
    shortDescription?: pulumi.Input<string>;
    /**
     * A note describing an SPDX File.
     */
    spdxFile?: pulumi.Input<inputs.containeranalysis.v1beta1.FileNoteArgs>;
    /**
     * A note describing an SPDX Package.
     */
    spdxPackage?: pulumi.Input<inputs.containeranalysis.v1beta1.PackageInfoNoteArgs>;
    /**
     * A note describing an SPDX File.
     */
    spdxRelationship?: pulumi.Input<inputs.containeranalysis.v1beta1.RelationshipNoteArgs>;
    /**
     * A note describing a package vulnerability.
     */
    vulnerability?: pulumi.Input<inputs.containeranalysis.v1beta1.VulnerabilityArgs>;
}