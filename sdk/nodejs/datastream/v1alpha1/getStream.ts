// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Use this method to get details about a stream.
 */
export function getStream(args: GetStreamArgs, opts?: pulumi.InvokeOptions): Promise<GetStreamResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:datastream/v1alpha1:getStream", {
        "location": args.location,
        "project": args.project,
        "streamId": args.streamId,
    }, opts);
}

export interface GetStreamArgs {
    location: string;
    project?: string;
    streamId: string;
}

export interface GetStreamResult {
    /**
     * Automatically backfill objects included in the stream source configuration. Specific objects can be excluded.
     */
    readonly backfillAll: outputs.datastream.v1alpha1.BackfillAllStrategyResponse;
    /**
     * Do not automatically backfill any objects.
     */
    readonly backfillNone: outputs.datastream.v1alpha1.BackfillNoneStrategyResponse;
    /**
     * The creation time of the stream.
     */
    readonly createTime: string;
    /**
     * Immutable. A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
     */
    readonly customerManagedEncryptionKey: string;
    /**
     * Destination connection profile configuration.
     */
    readonly destinationConfig: outputs.datastream.v1alpha1.DestinationConfigResponse;
    /**
     * Display name.
     */
    readonly displayName: string;
    /**
     * Errors on the Stream.
     */
    readonly errors: outputs.datastream.v1alpha1.ErrorResponse[];
    /**
     * Labels.
     */
    readonly labels: {[key: string]: string};
    /**
     * The stream's name.
     */
    readonly name: string;
    /**
     * Source connection profile configuration.
     */
    readonly sourceConfig: outputs.datastream.v1alpha1.SourceConfigResponse;
    /**
     * The state of the stream.
     */
    readonly state: string;
    /**
     * The last update time of the stream.
     */
    readonly updateTime: string;
}

export function getStreamOutput(args: GetStreamOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetStreamResult> {
    return pulumi.output(args).apply(a => getStream(a, opts))
}

export interface GetStreamOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    streamId: pulumi.Input<string>;
}