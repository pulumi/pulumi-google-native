// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new policy in the specified project and region using the data included in the request.
 */
export class RegionSslPolicy extends pulumi.CustomResource {
    /**
     * Get an existing RegionSslPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionSslPolicy {
        return new RegionSslPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/v1:RegionSslPolicy';

    /**
     * Returns true if the given object is an instance of RegionSslPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionSslPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionSslPolicy.__pulumiType;
    }

    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
     */
    public readonly customFeatures!: pulumi.Output<string[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The list of features enabled in the SSL policy.
     */
    public /*out*/ readonly enabledFeatures!: pulumi.Output<string[]>;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a SslPolicy. An up-to-date fingerprint must be provided in order to update the SslPolicy, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an SslPolicy.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * [Output only] Type of the resource. Always compute#sslPolicyfor SSL policies.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
     */
    public readonly minTlsVersion!: pulumi.Output<string>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
     */
    public readonly profile!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    public readonly region!: pulumi.Output<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * If potential misconfigurations are detected for this SSL policy, this field will be populated with warning messages.
     */
    public /*out*/ readonly warnings!: pulumi.Output<outputs.compute.v1.RegionSslPolicyWarningsItemResponse[]>;

    /**
     * Create a RegionSslPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionSslPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["customFeatures"] = args ? args.customFeatures : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["minTlsVersion"] = args ? args.minTlsVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["enabledFeatures"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["warnings"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["customFeatures"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["enabledFeatures"] = undefined /*out*/;
            resourceInputs["fingerprint"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["minTlsVersion"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["profile"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["warnings"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project", "region"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RegionSslPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionSslPolicy resource.
 */
export interface RegionSslPolicyArgs {
    /**
     * A list of features enabled when the selected profile is CUSTOM. The method returns the set of features that can be specified in this list. This field must be empty if the profile is not CUSTOM.
     */
    customFeatures?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * The minimum version of SSL protocol that can be used by the clients to establish a connection with the load balancer. This can be one of TLS_1_0, TLS_1_1, TLS_1_2.
     */
    minTlsVersion?: pulumi.Input<enums.compute.v1.RegionSslPolicyMinTlsVersion>;
    /**
     * Name of the resource. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * Profile specifies the set of SSL features that can be used by the load balancer when negotiating SSL with clients. This can be one of COMPATIBLE, MODERN, RESTRICTED, or CUSTOM. If using CUSTOM, the set of SSL features to enable must be specified in the customFeatures field.
     */
    profile?: pulumi.Input<enums.compute.v1.RegionSslPolicyProfile>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
}