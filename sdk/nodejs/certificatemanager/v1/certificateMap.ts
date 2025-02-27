// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new CertificateMap in a given project and location.
 */
export class CertificateMap extends pulumi.CustomResource {
    /**
     * Get an existing CertificateMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CertificateMap {
        return new CertificateMap(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:certificatemanager/v1:CertificateMap';

    /**
     * Returns true if the given object is an instance of CertificateMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CertificateMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CertificateMap.__pulumiType;
    }

    /**
     * Required. A user-provided name of the certificate map.
     */
    public readonly certificateMapId!: pulumi.Output<string>;
    /**
     * The creation timestamp of a Certificate Map.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * One or more paragraphs of text description of a certificate map.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * A list of GCLB targets that use this Certificate Map. A Target Proxy is only present on this list if it's attached to a Forwarding Rule.
     */
    public /*out*/ readonly gclbTargets!: pulumi.Output<outputs.certificatemanager.v1.GclbTargetResponse[]>;
    /**
     * Set of labels associated with a Certificate Map.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*&#47;locations/*&#47;certificateMaps/*`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The update timestamp of a Certificate Map.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a CertificateMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CertificateMapArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.certificateMapId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateMapId'");
            }
            resourceInputs["certificateMapId"] = args ? args.certificateMapId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["gclbTargets"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["certificateMapId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["gclbTargets"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["certificateMapId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(CertificateMap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a CertificateMap resource.
 */
export interface CertificateMapArgs {
    /**
     * Required. A user-provided name of the certificate map.
     */
    certificateMapId: pulumi.Input<string>;
    /**
     * One or more paragraphs of text description of a certificate map.
     */
    description?: pulumi.Input<string>;
    /**
     * Set of labels associated with a Certificate Map.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match pattern `projects/*&#47;locations/*&#47;certificateMaps/*`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
