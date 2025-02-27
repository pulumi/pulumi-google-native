// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a metastore federation in a project and location.
 */
export class Federation extends pulumi.CustomResource {
    /**
     * Get an existing Federation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Federation {
        return new Federation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:metastore/v1beta:Federation';

    /**
     * Returns true if the given object is an instance of Federation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Federation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Federation.__pulumiType;
    }

    /**
     * A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
     */
    public readonly backendMetastores!: pulumi.Output<{[key: string]: outputs.metastore.v1beta.BackendMetastoreResponse}>;
    /**
     * The time when the metastore federation was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The federation endpoint.
     */
    public /*out*/ readonly endpointUri!: pulumi.Output<string>;
    /**
     * Required. The ID of the metastore federation, which is used as the final component of the metastore federation's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
     */
    public readonly federationId!: pulumi.Output<string>;
    /**
     * User-defined labels for the metastore federation.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. The relative resource name of the federation, of the form: projects/{project_number}/locations/{location_id}/federations/{federation_id}`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The current state of the federation.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Additional information about the current state of the metastore federation, if available.
     */
    public /*out*/ readonly stateMessage!: pulumi.Output<string>;
    /**
     * The globally unique resource identifier of the metastore federation.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the metastore federation was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Immutable. The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Federation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FederationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.federationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'federationId'");
            }
            resourceInputs["backendMetastores"] = args ? args.backendMetastores : undefined;
            resourceInputs["federationId"] = args ? args.federationId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpointUri"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["backendMetastores"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endpointUri"] = undefined /*out*/;
            resourceInputs["federationId"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateMessage"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["federationId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Federation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Federation resource.
 */
export interface FederationArgs {
    /**
     * A map from BackendMetastore rank to BackendMetastores from which the federation service serves metadata at query time. The map key represents the order in which BackendMetastores should be evaluated to resolve database names at query time and should be greater than or equal to zero. A BackendMetastore with a lower number will be evaluated before a BackendMetastore with a higher number.
     */
    backendMetastores?: pulumi.Input<{[key: string]: pulumi.Input<inputs.metastore.v1beta.BackendMetastoreArgs>}>;
    /**
     * Required. The ID of the metastore federation, which is used as the final component of the metastore federation's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
     */
    federationId: pulumi.Input<string>;
    /**
     * User-defined labels for the metastore federation.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Immutable. The relative resource name of the federation, of the form: projects/{project_number}/locations/{location_id}/federations/{federation_id}`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Immutable. The Apache Hive metastore version of the federation. All backend metastore versions must be compatible with the federation version.
     */
    version?: pulumi.Input<string>;
}
