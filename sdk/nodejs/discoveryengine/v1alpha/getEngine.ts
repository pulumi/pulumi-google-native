// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets a Engine.
 */
export function getEngine(args: GetEngineArgs, opts?: pulumi.InvokeOptions): Promise<GetEngineResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:discoveryengine/v1alpha:getEngine", {
        "collectionId": args.collectionId,
        "engineId": args.engineId,
        "location": args.location,
        "project": args.project,
    }, opts);
}

export interface GetEngineArgs {
    collectionId: string;
    engineId: string;
    location: string;
    project?: string;
}

export interface GetEngineResult {
    /**
     * Configurations for the Chat Engine. Only applicable if solution_type is SOLUTION_TYPE_CHAT.
     */
    readonly chatEngineConfig: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigResponse;
    /**
     * Additional information of the Chat Engine. Only applicable if solution_type is SOLUTION_TYPE_CHAT.
     */
    readonly chatEngineMetadata: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineChatEngineMetadataResponse;
    /**
     * Common config spec that specifies the metadata of the engine.
     */
    readonly commonConfig: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineCommonConfigResponse;
    /**
     * Timestamp the Recommendation Engine was created at.
     */
    readonly createTime: string;
    /**
     * The data stores associated with this engine. For SOLUTION_TYPE_SEARCH and SOLUTION_TYPE_RECOMMENDATION type of engines, they can only associate with at most one data store. If solution_type is SOLUTION_TYPE_CHAT, multiple DataStores in the same Collection can be associated here. Note that when used in CreateEngineRequest, one DataStore id must be provided as the system will use it for necessary intializations.
     */
    readonly dataStoreIds: string[];
    /**
     * The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
     */
    readonly displayName: string;
    /**
     * The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to `GENERIC`. Vertical on Engine has to match vertical of the DataStore liniked to the engine.
     */
    readonly industryVertical: string;
    /**
     * Configurations for the Media Engine. Only applicable on the data stores with solution_type SOLUTION_TYPE_RECOMMENDATION and IndustryVertical.MEDIA vertical.
     */
    readonly mediaRecommendationEngineConfig: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineMediaRecommendationEngineConfigResponse;
    /**
     * Immutable. The fully qualified resource name of the engine. This field must be a UTF-8 encoded string with a length limit of 1024 characters. Format: `projects/{project_number}/locations/{location}/collections/{collection}/engines/{engine}` engine should be 1-63 characters, and valid characters are /a-z0-9*&#47;. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    readonly name: string;
    /**
     * Additional information of a recommendation engine. Only applicable if solution_type is SOLUTION_TYPE_RECOMMENDATION.
     */
    readonly recommendationMetadata: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineRecommendationMetadataResponse;
    /**
     * Configurations for the Search Engine. Only applicable if solution_type is SOLUTION_TYPE_SEARCH.
     */
    readonly searchEngineConfig: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineSearchEngineConfigResponse;
    /**
     * Additional config specs for a `similar-items` engine.
     */
    readonly similarDocumentsConfig: outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineSimilarDocumentsEngineConfigResponse;
    /**
     * The solutions of the engine.
     */
    readonly solutionType: string;
    /**
     * Timestamp the Recommendation Engine was last updated.
     */
    readonly updateTime: string;
}
/**
 * Gets a Engine.
 */
export function getEngineOutput(args: GetEngineOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEngineResult> {
    return pulumi.output(args).apply((a: any) => getEngine(a, opts))
}

export interface GetEngineOutputArgs {
    collectionId: pulumi.Input<string>;
    engineId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
