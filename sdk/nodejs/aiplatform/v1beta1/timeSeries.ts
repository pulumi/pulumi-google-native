// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a TensorboardTimeSeries.
 * Auto-naming is currently not supported for this resource.
 */
export class TimeSeries extends pulumi.CustomResource {
    /**
     * Get an existing TimeSeries resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TimeSeries {
        return new TimeSeries(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:aiplatform/v1beta1:TimeSeries';

    /**
     * Returns true if the given object is an instance of TimeSeries.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TimeSeries {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TimeSeries.__pulumiType;
    }

    /**
     * Timestamp when this TensorboardTimeSeries was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of this TensorboardTimeSeries.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * User provided name of this TensorboardTimeSeries. This value should be unique among all TensorboardTimeSeries resources belonging to the same TensorboardRun resource (parent resource).
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Used to perform a consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    public readonly etag!: pulumi.Output<string>;
    public readonly experimentId!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Scalar, Tensor, or Blob metadata for this TensorboardTimeSeries.
     */
    public /*out*/ readonly metadata!: pulumi.Output<outputs.aiplatform.v1beta1.GoogleCloudAiplatformV1beta1TensorboardTimeSeriesMetadataResponse>;
    /**
     * Name of the TensorboardTimeSeries.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Data of the current plugin, with the size limited to 65KB.
     */
    public readonly pluginData!: pulumi.Output<string>;
    /**
     * Immutable. Name of the plugin this time series pertain to. Such as Scalar, Tensor, Blob
     */
    public readonly pluginName!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    public readonly runId!: pulumi.Output<string>;
    public readonly tensorboardId!: pulumi.Output<string>;
    /**
     * Optional. The user specified unique ID to use for the TensorboardTimeSeries, which becomes the final component of the TensorboardTimeSeries's resource name. This value should match "a-z0-9{0, 127}"
     */
    public readonly tensorboardTimeSeriesId!: pulumi.Output<string | undefined>;
    /**
     * Timestamp when this TensorboardTimeSeries was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Immutable. Type of TensorboardTimeSeries value.
     */
    public readonly valueType!: pulumi.Output<string>;

    /**
     * Create a TimeSeries resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TimeSeriesArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.experimentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'experimentId'");
            }
            if ((!args || args.runId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runId'");
            }
            if ((!args || args.tensorboardId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tensorboardId'");
            }
            if ((!args || args.valueType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'valueType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["experimentId"] = args ? args.experimentId : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["pluginData"] = args ? args.pluginData : undefined;
            resourceInputs["pluginName"] = args ? args.pluginName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["runId"] = args ? args.runId : undefined;
            resourceInputs["tensorboardId"] = args ? args.tensorboardId : undefined;
            resourceInputs["tensorboardTimeSeriesId"] = args ? args.tensorboardTimeSeriesId : undefined;
            resourceInputs["valueType"] = args ? args.valueType : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["experimentId"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pluginData"] = undefined /*out*/;
            resourceInputs["pluginName"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["runId"] = undefined /*out*/;
            resourceInputs["tensorboardId"] = undefined /*out*/;
            resourceInputs["tensorboardTimeSeriesId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["valueType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["experimentId", "location", "project", "runId", "tensorboardId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TimeSeries.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TimeSeries resource.
 */
export interface TimeSeriesArgs {
    /**
     * Description of this TensorboardTimeSeries.
     */
    description?: pulumi.Input<string>;
    /**
     * User provided name of this TensorboardTimeSeries. This value should be unique among all TensorboardTimeSeries resources belonging to the same TensorboardRun resource (parent resource).
     */
    displayName: pulumi.Input<string>;
    /**
     * Used to perform a consistent read-modify-write updates. If not set, a blind "overwrite" update happens.
     */
    etag?: pulumi.Input<string>;
    experimentId: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Data of the current plugin, with the size limited to 65KB.
     */
    pluginData?: pulumi.Input<string>;
    /**
     * Immutable. Name of the plugin this time series pertain to. Such as Scalar, Tensor, Blob
     */
    pluginName?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    runId: pulumi.Input<string>;
    tensorboardId: pulumi.Input<string>;
    /**
     * Optional. The user specified unique ID to use for the TensorboardTimeSeries, which becomes the final component of the TensorboardTimeSeries's resource name. This value should match "a-z0-9{0, 127}"
     */
    tensorboardTimeSeriesId?: pulumi.Input<string>;
    /**
     * Immutable. Type of TensorboardTimeSeries value.
     */
    valueType: pulumi.Input<enums.aiplatform.v1beta1.TimeSeriesValueType>;
}
