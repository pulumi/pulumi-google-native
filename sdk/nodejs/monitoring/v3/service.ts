// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Create a Service.
 * Auto-naming is currently not supported for this resource.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:monitoring/v3:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * Type used for App Engine services.
     */
    public readonly appEngine!: pulumi.Output<outputs.monitoring.v3.AppEngineResponse>;
    /**
     * Type used for Cloud Endpoints services.
     */
    public readonly cloudEndpoints!: pulumi.Output<outputs.monitoring.v3.CloudEndpointsResponse>;
    /**
     * Type used for Istio services that live in a Kubernetes cluster.
     */
    public readonly clusterIstio!: pulumi.Output<outputs.monitoring.v3.ClusterIstioResponse>;
    /**
     * Custom service type.
     */
    public readonly custom!: pulumi.Output<outputs.monitoring.v3.CustomResponse>;
    /**
     * Name used for UI elements listing this Service.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Type used for canonical services scoped to an Istio mesh. Metrics for Istio are documented here (https://istio.io/latest/docs/reference/config/metrics/)
     */
    public readonly istioCanonicalService!: pulumi.Output<outputs.monitoring.v3.IstioCanonicalServiceResponse>;
    /**
     * Type used for Istio services scoped to an Istio mesh.
     */
    public readonly meshIstio!: pulumi.Output<outputs.monitoring.v3.MeshIstioResponse>;
    /**
     * Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID] 
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for how to query telemetry on a Service.
     */
    public readonly telemetry!: pulumi.Output<outputs.monitoring.v3.TelemetryResponse>;
    /**
     * Labels which have been used to annotate the service. Label keys must start with a letter. Label keys and values may contain lowercase letters, numbers, underscores, and dashes. Label keys and values have a maximum length of 63 characters, and must be less than 128 bytes in size. Up to 64 label entries may be stored. For labels which do not have a semantic value, the empty string may be supplied for the label value.
     */
    public readonly userLabels!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.v3Id === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v3Id'");
            }
            if ((!args || args.v3Id1 === undefined) && !opts.urn) {
                throw new Error("Missing required property 'v3Id1'");
            }
            resourceInputs["appEngine"] = args ? args.appEngine : undefined;
            resourceInputs["cloudEndpoints"] = args ? args.cloudEndpoints : undefined;
            resourceInputs["clusterIstio"] = args ? args.clusterIstio : undefined;
            resourceInputs["custom"] = args ? args.custom : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["istioCanonicalService"] = args ? args.istioCanonicalService : undefined;
            resourceInputs["meshIstio"] = args ? args.meshIstio : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["telemetry"] = args ? args.telemetry : undefined;
            resourceInputs["userLabels"] = args ? args.userLabels : undefined;
            resourceInputs["v3Id"] = args ? args.v3Id : undefined;
            resourceInputs["v3Id1"] = args ? args.v3Id1 : undefined;
        } else {
            resourceInputs["appEngine"] = undefined /*out*/;
            resourceInputs["cloudEndpoints"] = undefined /*out*/;
            resourceInputs["clusterIstio"] = undefined /*out*/;
            resourceInputs["custom"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["istioCanonicalService"] = undefined /*out*/;
            resourceInputs["meshIstio"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["telemetry"] = undefined /*out*/;
            resourceInputs["userLabels"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * Type used for App Engine services.
     */
    appEngine?: pulumi.Input<inputs.monitoring.v3.AppEngineArgs>;
    /**
     * Type used for Cloud Endpoints services.
     */
    cloudEndpoints?: pulumi.Input<inputs.monitoring.v3.CloudEndpointsArgs>;
    /**
     * Type used for Istio services that live in a Kubernetes cluster.
     */
    clusterIstio?: pulumi.Input<inputs.monitoring.v3.ClusterIstioArgs>;
    /**
     * Custom service type.
     */
    custom?: pulumi.Input<inputs.monitoring.v3.CustomArgs>;
    /**
     * Name used for UI elements listing this Service.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Type used for canonical services scoped to an Istio mesh. Metrics for Istio are documented here (https://istio.io/latest/docs/reference/config/metrics/)
     */
    istioCanonicalService?: pulumi.Input<inputs.monitoring.v3.IstioCanonicalServiceArgs>;
    /**
     * Type used for Istio services scoped to an Istio mesh.
     */
    meshIstio?: pulumi.Input<inputs.monitoring.v3.MeshIstioArgs>;
    /**
     * Resource name for this Service. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID] 
     */
    name?: pulumi.Input<string>;
    /**
     * Optional. The Service id to use for this Service. If omitted, an id will be generated instead. Must match the pattern [a-z0-9\-]+
     */
    serviceId?: pulumi.Input<string>;
    /**
     * Configuration for how to query telemetry on a Service.
     */
    telemetry?: pulumi.Input<inputs.monitoring.v3.TelemetryArgs>;
    /**
     * Labels which have been used to annotate the service. Label keys must start with a letter. Label keys and values may contain lowercase letters, numbers, underscores, and dashes. Label keys and values have a maximum length of 63 characters, and must be less than 128 bytes in size. Up to 64 label entries may be stored. For labels which do not have a semantic value, the empty string may be supplied for the label value.
     */
    userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    v3Id: pulumi.Input<string>;
    v3Id1: pulumi.Input<string>;
}