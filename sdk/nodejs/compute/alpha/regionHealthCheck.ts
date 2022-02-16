// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a HealthCheck resource in the specified project using the data included in the request.
 */
export class RegionHealthCheck extends pulumi.CustomResource {
    /**
     * Get an existing RegionHealthCheck resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionHealthCheck {
        return new RegionHealthCheck(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:RegionHealthCheck';

    /**
     * Returns true if the given object is an instance of RegionHealthCheck.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionHealthCheck {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionHealthCheck.__pulumiType;
    }

    /**
     * How often (in seconds) to send a health check. The default value is 5 seconds.
     */
    public readonly checkIntervalSec!: pulumi.Output<number>;
    /**
     * Creation timestamp in 3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly grpcHealthCheck!: pulumi.Output<outputs.compute.alpha.GRPCHealthCheckResponse>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
     */
    public readonly healthyThreshold!: pulumi.Output<number>;
    public readonly http2HealthCheck!: pulumi.Output<outputs.compute.alpha.HTTP2HealthCheckResponse>;
    public readonly httpHealthCheck!: pulumi.Output<outputs.compute.alpha.HTTPHealthCheckResponse>;
    public readonly httpsHealthCheck!: pulumi.Output<outputs.compute.alpha.HTTPSHealthCheckResponse>;
    /**
     * Type of the resource.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Configure logging on this health check.
     */
    public readonly logConfig!: pulumi.Output<outputs.compute.alpha.HealthCheckLogConfigResponse>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Region where the health check resides. Not applicable to global health checks.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Server-defined URL for the resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    public readonly sslHealthCheck!: pulumi.Output<outputs.compute.alpha.SSLHealthCheckResponse>;
    public readonly tcpHealthCheck!: pulumi.Output<outputs.compute.alpha.TCPHealthCheckResponse>;
    /**
     * How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
     */
    public readonly timeoutSec!: pulumi.Output<number>;
    /**
     * Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
     */
    public readonly type!: pulumi.Output<string>;
    public readonly udpHealthCheck!: pulumi.Output<outputs.compute.alpha.UDPHealthCheckResponse>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
     */
    public readonly unhealthyThreshold!: pulumi.Output<number>;

    /**
     * Create a RegionHealthCheck resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionHealthCheckArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["checkIntervalSec"] = args ? args.checkIntervalSec : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["grpcHealthCheck"] = args ? args.grpcHealthCheck : undefined;
            resourceInputs["healthyThreshold"] = args ? args.healthyThreshold : undefined;
            resourceInputs["http2HealthCheck"] = args ? args.http2HealthCheck : undefined;
            resourceInputs["httpHealthCheck"] = args ? args.httpHealthCheck : undefined;
            resourceInputs["httpsHealthCheck"] = args ? args.httpsHealthCheck : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["logConfig"] = args ? args.logConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["sslHealthCheck"] = args ? args.sslHealthCheck : undefined;
            resourceInputs["tcpHealthCheck"] = args ? args.tcpHealthCheck : undefined;
            resourceInputs["timeoutSec"] = args ? args.timeoutSec : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["udpHealthCheck"] = args ? args.udpHealthCheck : undefined;
            resourceInputs["unhealthyThreshold"] = args ? args.unhealthyThreshold : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
        } else {
            resourceInputs["checkIntervalSec"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["grpcHealthCheck"] = undefined /*out*/;
            resourceInputs["healthyThreshold"] = undefined /*out*/;
            resourceInputs["http2HealthCheck"] = undefined /*out*/;
            resourceInputs["httpHealthCheck"] = undefined /*out*/;
            resourceInputs["httpsHealthCheck"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["logConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["selfLinkWithId"] = undefined /*out*/;
            resourceInputs["sslHealthCheck"] = undefined /*out*/;
            resourceInputs["tcpHealthCheck"] = undefined /*out*/;
            resourceInputs["timeoutSec"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["udpHealthCheck"] = undefined /*out*/;
            resourceInputs["unhealthyThreshold"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RegionHealthCheck.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionHealthCheck resource.
 */
export interface RegionHealthCheckArgs {
    /**
     * How often (in seconds) to send a health check. The default value is 5 seconds.
     */
    checkIntervalSec?: pulumi.Input<number>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    grpcHealthCheck?: pulumi.Input<inputs.compute.alpha.GRPCHealthCheckArgs>;
    /**
     * A so-far unhealthy instance will be marked healthy after this many consecutive successes. The default value is 2.
     */
    healthyThreshold?: pulumi.Input<number>;
    http2HealthCheck?: pulumi.Input<inputs.compute.alpha.HTTP2HealthCheckArgs>;
    httpHealthCheck?: pulumi.Input<inputs.compute.alpha.HTTPHealthCheckArgs>;
    httpsHealthCheck?: pulumi.Input<inputs.compute.alpha.HTTPSHealthCheckArgs>;
    /**
     * Type of the resource.
     */
    kind?: pulumi.Input<string>;
    /**
     * Configure logging on this health check.
     */
    logConfig?: pulumi.Input<inputs.compute.alpha.HealthCheckLogConfigArgs>;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. For example, a name that is 1-63 characters long, matches the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?`, and otherwise complies with RFC1035. This regular expression describes a name where the first character is a lowercase letter, and all following characters are a dash, lowercase letter, or digit, except the last character, which isn't a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    sslHealthCheck?: pulumi.Input<inputs.compute.alpha.SSLHealthCheckArgs>;
    tcpHealthCheck?: pulumi.Input<inputs.compute.alpha.TCPHealthCheckArgs>;
    /**
     * How long (in seconds) to wait before claiming failure. The default value is 5 seconds. It is invalid for timeoutSec to have greater value than checkIntervalSec.
     */
    timeoutSec?: pulumi.Input<number>;
    /**
     * Specifies the type of the healthCheck, either TCP, SSL, HTTP, HTTPS or HTTP2. Exactly one of the protocol-specific health check field must be specified, which must match type field.
     */
    type?: pulumi.Input<enums.compute.alpha.RegionHealthCheckType>;
    udpHealthCheck?: pulumi.Input<inputs.compute.alpha.UDPHealthCheckArgs>;
    /**
     * A so-far healthy instance will be marked unhealthy after this many consecutive failures. The default value is 2.
     */
    unhealthyThreshold?: pulumi.Input<number>;
}