// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a Engine.
 */
export class Engine extends pulumi.CustomResource {
    /**
     * Get an existing Engine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Engine {
        return new Engine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:discoveryengine/v1alpha:Engine';

    /**
     * Returns true if the given object is an instance of Engine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Engine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Engine.__pulumiType;
    }

    /**
     * Configurations for the Chat Engine. Only applicable if solution_type is SOLUTION_TYPE_CHAT.
     */
    public readonly chatEngineConfig!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigResponse>;
    /**
     * Additional information of the Chat Engine. Only applicable if solution_type is SOLUTION_TYPE_CHAT.
     */
    public /*out*/ readonly chatEngineMetadata!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineChatEngineMetadataResponse>;
    public readonly collectionId!: pulumi.Output<string>;
    /**
     * Common config spec that specifies the metadata of the engine.
     */
    public readonly commonConfig!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineCommonConfigResponse>;
    /**
     * Timestamp the Recommendation Engine was created at.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The data stores associated with this engine. For SOLUTION_TYPE_SEARCH and SOLUTION_TYPE_RECOMMENDATION type of engines, they can only associate with at most one data store. If solution_type is SOLUTION_TYPE_CHAT, multiple DataStores in the same Collection can be associated here. Note that when used in CreateEngineRequest, one DataStore id must be provided as the system will use it for necessary intializations.
     */
    public readonly dataStoreIds!: pulumi.Output<string[]>;
    /**
     * The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Required. The ID to use for the Engine, which will become the final component of the Engine's resource name. This field must conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    public readonly engineId!: pulumi.Output<string>;
    /**
     * The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to `GENERIC`. Vertical on Engine has to match vertical of the DataStore liniked to the engine.
     */
    public readonly industryVertical!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Configurations for the Media Engine. Only applicable on the data stores with solution_type SOLUTION_TYPE_RECOMMENDATION and IndustryVertical.MEDIA vertical.
     */
    public readonly mediaRecommendationEngineConfig!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineMediaRecommendationEngineConfigResponse>;
    /**
     * Immutable. The fully qualified resource name of the engine. This field must be a UTF-8 encoded string with a length limit of 1024 characters. Format: `projects/{project_number}/locations/{location}/collections/{collection}/engines/{engine}` engine should be 1-63 characters, and valid characters are /a-z0-9*&#47;. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Additional information of a recommendation engine. Only applicable if solution_type is SOLUTION_TYPE_RECOMMENDATION.
     */
    public /*out*/ readonly recommendationMetadata!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineRecommendationMetadataResponse>;
    /**
     * Configurations for the Search Engine. Only applicable if solution_type is SOLUTION_TYPE_SEARCH.
     */
    public readonly searchEngineConfig!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineSearchEngineConfigResponse>;
    /**
     * Additional config specs for a `similar-items` engine.
     */
    public readonly similarDocumentsConfig!: pulumi.Output<outputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineSimilarDocumentsEngineConfigResponse>;
    /**
     * The solutions of the engine.
     */
    public readonly solutionType!: pulumi.Output<string>;
    /**
     * Timestamp the Recommendation Engine was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Engine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EngineArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.collectionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'collectionId'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.engineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineId'");
            }
            if ((!args || args.solutionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'solutionType'");
            }
            resourceInputs["chatEngineConfig"] = args ? args.chatEngineConfig : undefined;
            resourceInputs["collectionId"] = args ? args.collectionId : undefined;
            resourceInputs["commonConfig"] = args ? args.commonConfig : undefined;
            resourceInputs["dataStoreIds"] = args ? args.dataStoreIds : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["engineId"] = args ? args.engineId : undefined;
            resourceInputs["industryVertical"] = args ? args.industryVertical : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["mediaRecommendationEngineConfig"] = args ? args.mediaRecommendationEngineConfig : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["searchEngineConfig"] = args ? args.searchEngineConfig : undefined;
            resourceInputs["similarDocumentsConfig"] = args ? args.similarDocumentsConfig : undefined;
            resourceInputs["solutionType"] = args ? args.solutionType : undefined;
            resourceInputs["chatEngineMetadata"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["recommendationMetadata"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["chatEngineConfig"] = undefined /*out*/;
            resourceInputs["chatEngineMetadata"] = undefined /*out*/;
            resourceInputs["collectionId"] = undefined /*out*/;
            resourceInputs["commonConfig"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dataStoreIds"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["engineId"] = undefined /*out*/;
            resourceInputs["industryVertical"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["mediaRecommendationEngineConfig"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["recommendationMetadata"] = undefined /*out*/;
            resourceInputs["searchEngineConfig"] = undefined /*out*/;
            resourceInputs["similarDocumentsConfig"] = undefined /*out*/;
            resourceInputs["solutionType"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["collectionId", "engineId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Engine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Engine resource.
 */
export interface EngineArgs {
    /**
     * Configurations for the Chat Engine. Only applicable if solution_type is SOLUTION_TYPE_CHAT.
     */
    chatEngineConfig?: pulumi.Input<inputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineChatEngineConfigArgs>;
    collectionId: pulumi.Input<string>;
    /**
     * Common config spec that specifies the metadata of the engine.
     */
    commonConfig?: pulumi.Input<inputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineCommonConfigArgs>;
    /**
     * The data stores associated with this engine. For SOLUTION_TYPE_SEARCH and SOLUTION_TYPE_RECOMMENDATION type of engines, they can only associate with at most one data store. If solution_type is SOLUTION_TYPE_CHAT, multiple DataStores in the same Collection can be associated here. Note that when used in CreateEngineRequest, one DataStore id must be provided as the system will use it for necessary intializations.
     */
    dataStoreIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * Required. The ID to use for the Engine, which will become the final component of the Engine's resource name. This field must conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    engineId: pulumi.Input<string>;
    /**
     * The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to `GENERIC`. Vertical on Engine has to match vertical of the DataStore liniked to the engine.
     */
    industryVertical?: pulumi.Input<enums.discoveryengine.v1alpha.EngineIndustryVertical>;
    location?: pulumi.Input<string>;
    /**
     * Configurations for the Media Engine. Only applicable on the data stores with solution_type SOLUTION_TYPE_RECOMMENDATION and IndustryVertical.MEDIA vertical.
     */
    mediaRecommendationEngineConfig?: pulumi.Input<inputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineMediaRecommendationEngineConfigArgs>;
    /**
     * Immutable. The fully qualified resource name of the engine. This field must be a UTF-8 encoded string with a length limit of 1024 characters. Format: `projects/{project_number}/locations/{location}/collections/{collection}/engines/{engine}` engine should be 1-63 characters, and valid characters are /a-z0-9*&#47;. Otherwise, an INVALID_ARGUMENT error is returned.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Configurations for the Search Engine. Only applicable if solution_type is SOLUTION_TYPE_SEARCH.
     */
    searchEngineConfig?: pulumi.Input<inputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineSearchEngineConfigArgs>;
    /**
     * Additional config specs for a `similar-items` engine.
     */
    similarDocumentsConfig?: pulumi.Input<inputs.discoveryengine.v1alpha.GoogleCloudDiscoveryengineV1alphaEngineSimilarDocumentsEngineConfigArgs>;
    /**
     * The solutions of the engine.
     */
    solutionType: pulumi.Input<enums.discoveryengine.v1alpha.EngineSolutionType>;
}