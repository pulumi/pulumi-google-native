// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new ServiceConnectionPolicy in a given project and location.
 */
export class ServiceConnectionPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ServiceConnectionPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceConnectionPolicy {
        return new ServiceConnectionPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkconnectivity/v1:ServiceConnectionPolicy';

    /**
     * Returns true if the given object is an instance of ServiceConnectionPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceConnectionPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceConnectionPolicy.__pulumiType;
    }

    /**
     * Time when the ServiceConnectionMap was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A description of this resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The type of underlying resources used to create the connection.
     */
    public /*out*/ readonly infrastructure!: pulumi.Output<string>;
    /**
     * User-defined labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
     */
    public readonly network!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
     */
    public readonly pscConfig!: pulumi.Output<outputs.networkconnectivity.v1.PscConfigResponse>;
    /**
     * [Output only] Information about each Private Service Connect connection.
     */
    public /*out*/ readonly pscConnections!: pulumi.Output<outputs.networkconnectivity.v1.PscConnectionResponse[]>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass. It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
     */
    public readonly serviceClass!: pulumi.Output<string>;
    /**
     * Optional. Resource ID (i.e. 'foo' in '[...]/projects/p/locations/l/serviceConnectionPolicies/foo') See https://google.aip.dev/122#resource-id-segments Unique per location.
     */
    public readonly serviceConnectionPolicyId!: pulumi.Output<string | undefined>;
    /**
     * Time when the ServiceConnectionMap was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ServiceConnectionPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ServiceConnectionPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["pscConfig"] = args ? args.pscConfig : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["serviceClass"] = args ? args.serviceClass : undefined;
            resourceInputs["serviceConnectionPolicyId"] = args ? args.serviceConnectionPolicyId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["infrastructure"] = undefined /*out*/;
            resourceInputs["pscConnections"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["infrastructure"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["pscConfig"] = undefined /*out*/;
            resourceInputs["pscConnections"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["serviceClass"] = undefined /*out*/;
            resourceInputs["serviceConnectionPolicyId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ServiceConnectionPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceConnectionPolicy resource.
 */
export interface ServiceConnectionPolicyArgs {
    /**
     * A description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * User-defined labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
     */
    name?: pulumi.Input<string>;
    /**
     * The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
     */
    network?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
     */
    pscConfig?: pulumi.Input<inputs.networkconnectivity.v1.PscConfigArgs>;
    /**
     * Optional. An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes since the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass. It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
     */
    serviceClass?: pulumi.Input<string>;
    /**
     * Optional. Resource ID (i.e. 'foo' in '[...]/projects/p/locations/l/serviceConnectionPolicies/foo') See https://google.aip.dev/122#resource-id-segments Unique per location.
     */
    serviceConnectionPolicyId?: pulumi.Input<string>;
}