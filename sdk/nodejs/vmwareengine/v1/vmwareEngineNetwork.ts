// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new VMware Engine network that can be used by a private cloud.
 * Auto-naming is currently not supported for this resource.
 */
export class VmwareEngineNetwork extends pulumi.CustomResource {
    /**
     * Get an existing VmwareEngineNetwork resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VmwareEngineNetwork {
        return new VmwareEngineNetwork(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:vmwareengine/v1:VmwareEngineNetwork';

    /**
     * Returns true if the given object is an instance of VmwareEngineNetwork.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VmwareEngineNetwork {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VmwareEngineNetwork.__pulumiType;
    }

    /**
     * Creation time of this resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * User-provided description for this VMware Engine network.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Checksum that may be sent on update and delete requests to ensure that the user-provided value is up to date before the server processes a request. The server computes checksums based on the value of other fields in the request.
     */
    public readonly etag!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The resource name of the VMware Engine network. Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names. For example: `projects/my-project/locations/global/vmwareEngineNetworks/my-network`
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server guarantees that a request doesn't result in creation of duplicate commitments for at least 60 minutes. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * State of the VMware Engine network.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * VMware Engine network type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * System-generated unique identifier for the resource.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Last update time of this resource.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Required. The user-provided identifier of the new VMware Engine network. This identifier must be unique among VMware Engine network resources within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * For networks of type LEGACY, adheres to the format: `{region-id}-default`. Replace `{region-id}` with the region where you want to create the VMware Engine network. For example, "us-central1-default". * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
     */
    public readonly vmwareEngineNetworkId!: pulumi.Output<string>;
    /**
     * VMware Engine service VPC networks that provide connectivity from a private cloud to customer projects, the internet, and other Google Cloud services.
     */
    public /*out*/ readonly vpcNetworks!: pulumi.Output<outputs.vmwareengine.v1.VpcNetworkResponse[]>;

    /**
     * Create a VmwareEngineNetwork resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VmwareEngineNetworkArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.vmwareEngineNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vmwareEngineNetworkId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vmwareEngineNetworkId"] = args ? args.vmwareEngineNetworkId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vpcNetworks"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["vmwareEngineNetworkId"] = undefined /*out*/;
            resourceInputs["vpcNetworks"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project", "vmwareEngineNetworkId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(VmwareEngineNetwork.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a VmwareEngineNetwork resource.
 */
export interface VmwareEngineNetworkArgs {
    /**
     * User-provided description for this VMware Engine network.
     */
    description?: pulumi.Input<string>;
    /**
     * Checksum that may be sent on update and delete requests to ensure that the user-provided value is up to date before the server processes a request. The server computes checksums based on the value of other fields in the request.
     */
    etag?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server guarantees that a request doesn't result in creation of duplicate commitments for at least 60 minutes. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * VMware Engine network type.
     */
    type: pulumi.Input<enums.vmwareengine.v1.VmwareEngineNetworkType>;
    /**
     * Required. The user-provided identifier of the new VMware Engine network. This identifier must be unique among VMware Engine network resources within the parent and becomes the final token in the name URI. The identifier must meet the following requirements: * For networks of type LEGACY, adheres to the format: `{region-id}-default`. Replace `{region-id}` with the region where you want to create the VMware Engine network. For example, "us-central1-default". * Only contains 1-63 alphanumeric characters and hyphens * Begins with an alphabetical character * Ends with a non-hyphen character * Not formatted as a UUID * Complies with [RFC 1034](https://datatracker.ietf.org/doc/html/rfc1034) (section 3.5)
     */
    vmwareEngineNetworkId: pulumi.Input<string>;
}
