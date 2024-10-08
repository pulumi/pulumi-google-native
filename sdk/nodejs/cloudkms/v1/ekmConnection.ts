// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new EkmConnection in a given Project and Location.
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class EkmConnection extends pulumi.CustomResource {
    /**
     * Get an existing EkmConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): EkmConnection {
        return new EkmConnection(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudkms/v1:EkmConnection';

    /**
     * Returns true if the given object is an instance of EkmConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EkmConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EkmConnection.__pulumiType;
    }

    /**
     * The time at which the EkmConnection was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if KeyManagementMode is CLOUD_KMS.
     */
    public readonly cryptoSpacePath!: pulumi.Output<string>;
    /**
     * Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`.
     */
    public readonly ekmConnectionId!: pulumi.Output<string>;
    /**
     * Optional. Etag of the currently stored EkmConnection.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL.
     */
    public readonly keyManagementMode!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name for the EkmConnection in the format `projects/*&#47;locations/*&#47;ekmConnections/*`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
     */
    public readonly serviceResolvers!: pulumi.Output<outputs.cloudkms.v1.ServiceResolverResponse[]>;

    /**
     * Create a EkmConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EkmConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["cryptoSpacePath"] = args ? args.cryptoSpacePath : undefined;
            resourceInputs["ekmConnectionId"] = args ? args.ekmConnectionId : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["keyManagementMode"] = args ? args.keyManagementMode : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceResolvers"] = args ? args.serviceResolvers : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["cryptoSpacePath"] = undefined /*out*/;
            resourceInputs["ekmConnectionId"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["keyManagementMode"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["serviceResolvers"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["ekmConnectionId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(EkmConnection.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a EkmConnection resource.
 */
export interface EkmConnectionArgs {
    /**
     * Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if KeyManagementMode is CLOUD_KMS.
     */
    cryptoSpacePath?: pulumi.Input<string>;
    /**
     * Required. It must be unique within a location and match the regular expression `[a-zA-Z0-9_-]{1,63}`.
     */
    ekmConnectionId?: pulumi.Input<string>;
    /**
     * Optional. Etag of the currently stored EkmConnection.
     */
    etag?: pulumi.Input<string>;
    /**
     * Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL.
     */
    keyManagementMode?: pulumi.Input<enums.cloudkms.v1.EkmConnectionKeyManagementMode>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently, only a single ServiceResolver is supported.
     */
    serviceResolvers?: pulumi.Input<pulumi.Input<inputs.cloudkms.v1.ServiceResolverArgs>[]>;
}
