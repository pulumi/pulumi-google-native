// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new data policy under a project with the given `dataPolicyId` (used as the display name), policy tag, and data policy type.
 * Auto-naming is currently not supported for this resource.
 */
export class DataPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DataPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataPolicy {
        return new DataPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:bigquerydatapolicy/v1:DataPolicy';

    /**
     * Returns true if the given object is an instance of DataPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataPolicy.__pulumiType;
    }

    /**
     * The data masking policy that specifies the data masking rule to use.
     */
    public readonly dataMaskingPolicy!: pulumi.Output<outputs.bigquerydatapolicy.v1.DataMaskingPolicyResponse>;
    /**
     * User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {data_policy_id} in part of the resource name.
     */
    public readonly dataPolicyId!: pulumi.Output<string>;
    /**
     * Type of data policy.
     */
    public readonly dataPolicyType!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name of this data policy, in the format of `projects/{project_number}/locations/{location_id}/dataPolicies/{data_policy_id}`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Policy tag resource name, in the format of `projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{policyTag_id}`.
     */
    public readonly policyTag!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a DataPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DataPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["dataMaskingPolicy"] = args ? args.dataMaskingPolicy : undefined;
            resourceInputs["dataPolicyId"] = args ? args.dataPolicyId : undefined;
            resourceInputs["dataPolicyType"] = args ? args.dataPolicyType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["policyTag"] = args ? args.policyTag : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["dataMaskingPolicy"] = undefined /*out*/;
            resourceInputs["dataPolicyId"] = undefined /*out*/;
            resourceInputs["dataPolicyType"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["policyTag"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DataPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataPolicy resource.
 */
export interface DataPolicyArgs {
    /**
     * The data masking policy that specifies the data masking rule to use.
     */
    dataMaskingPolicy?: pulumi.Input<inputs.bigquerydatapolicy.v1.DataMaskingPolicyArgs>;
    /**
     * User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {data_policy_id} in part of the resource name.
     */
    dataPolicyId?: pulumi.Input<string>;
    /**
     * Type of data policy.
     */
    dataPolicyType?: pulumi.Input<enums.bigquerydatapolicy.v1.DataPolicyDataPolicyType>;
    location?: pulumi.Input<string>;
    /**
     * Policy tag resource name, in the format of `projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{policyTag_id}`.
     */
    policyTag?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}