// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a HyperparameterTuningJob
 * Auto-naming is currently not supported for this resource.
 */
export class HyperparameterTuningJob extends pulumi.CustomResource {
    /**
     * Get an existing HyperparameterTuningJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HyperparameterTuningJob {
        return new HyperparameterTuningJob(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:aiplatform/v1:HyperparameterTuningJob';

    /**
     * Returns true if the given object is an instance of HyperparameterTuningJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HyperparameterTuningJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HyperparameterTuningJob.__pulumiType;
    }

    /**
     * Time when the HyperparameterTuningJob was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The display name of the HyperparameterTuningJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Customer-managed encryption key options for a HyperparameterTuningJob. If this is set, then all resources created by the HyperparameterTuningJob will be encrypted with the provided encryption key.
     */
    public readonly encryptionSpec!: pulumi.Output<outputs.aiplatform.v1.GoogleCloudAiplatformV1EncryptionSpecResponse>;
    /**
     * Time when the HyperparameterTuningJob entered any of the following states: `JOB_STATE_SUCCEEDED`, `JOB_STATE_FAILED`, `JOB_STATE_CANCELLED`.
     */
    public /*out*/ readonly endTime!: pulumi.Output<string>;
    /**
     * Only populated when job's state is JOB_STATE_FAILED or JOB_STATE_CANCELLED.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.aiplatform.v1.GoogleRpcStatusResponse>;
    /**
     * The labels with user-defined metadata to organize HyperparameterTuningJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The number of failed Trials that need to be seen before failing the HyperparameterTuningJob. If set to 0, Vertex AI decides how many Trials must fail before the whole job fails.
     */
    public readonly maxFailedTrialCount!: pulumi.Output<number>;
    /**
     * The desired total number of Trials.
     */
    public readonly maxTrialCount!: pulumi.Output<number>;
    /**
     * Resource name of the HyperparameterTuningJob.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The desired number of Trials to run in parallel.
     */
    public readonly parallelTrialCount!: pulumi.Output<number>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Time when the HyperparameterTuningJob for the first time entered the `JOB_STATE_RUNNING` state.
     */
    public /*out*/ readonly startTime!: pulumi.Output<string>;
    /**
     * The detailed state of the job.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Study configuration of the HyperparameterTuningJob.
     */
    public readonly studySpec!: pulumi.Output<outputs.aiplatform.v1.GoogleCloudAiplatformV1StudySpecResponse>;
    /**
     * The spec of a trial job. The same spec applies to the CustomJobs created in all the trials.
     */
    public readonly trialJobSpec!: pulumi.Output<outputs.aiplatform.v1.GoogleCloudAiplatformV1CustomJobSpecResponse>;
    /**
     * Trials of the HyperparameterTuningJob.
     */
    public /*out*/ readonly trials!: pulumi.Output<outputs.aiplatform.v1.GoogleCloudAiplatformV1TrialResponse[]>;
    /**
     * Time when the HyperparameterTuningJob was most recently updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a HyperparameterTuningJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HyperparameterTuningJobArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.maxTrialCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'maxTrialCount'");
            }
            if ((!args || args.parallelTrialCount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parallelTrialCount'");
            }
            if ((!args || args.studySpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'studySpec'");
            }
            if ((!args || args.trialJobSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'trialJobSpec'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["encryptionSpec"] = args ? args.encryptionSpec : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maxFailedTrialCount"] = args ? args.maxFailedTrialCount : undefined;
            resourceInputs["maxTrialCount"] = args ? args.maxTrialCount : undefined;
            resourceInputs["parallelTrialCount"] = args ? args.parallelTrialCount : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["studySpec"] = args ? args.studySpec : undefined;
            resourceInputs["trialJobSpec"] = args ? args.trialJobSpec : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["trials"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["encryptionSpec"] = undefined /*out*/;
            resourceInputs["endTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["maxFailedTrialCount"] = undefined /*out*/;
            resourceInputs["maxTrialCount"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parallelTrialCount"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["startTime"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["studySpec"] = undefined /*out*/;
            resourceInputs["trialJobSpec"] = undefined /*out*/;
            resourceInputs["trials"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(HyperparameterTuningJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a HyperparameterTuningJob resource.
 */
export interface HyperparameterTuningJobArgs {
    /**
     * The display name of the HyperparameterTuningJob. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     */
    displayName: pulumi.Input<string>;
    /**
     * Customer-managed encryption key options for a HyperparameterTuningJob. If this is set, then all resources created by the HyperparameterTuningJob will be encrypted with the provided encryption key.
     */
    encryptionSpec?: pulumi.Input<inputs.aiplatform.v1.GoogleCloudAiplatformV1EncryptionSpecArgs>;
    /**
     * The labels with user-defined metadata to organize HyperparameterTuningJobs. Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * The number of failed Trials that need to be seen before failing the HyperparameterTuningJob. If set to 0, Vertex AI decides how many Trials must fail before the whole job fails.
     */
    maxFailedTrialCount?: pulumi.Input<number>;
    /**
     * The desired total number of Trials.
     */
    maxTrialCount: pulumi.Input<number>;
    /**
     * The desired number of Trials to run in parallel.
     */
    parallelTrialCount: pulumi.Input<number>;
    project?: pulumi.Input<string>;
    /**
     * Study configuration of the HyperparameterTuningJob.
     */
    studySpec: pulumi.Input<inputs.aiplatform.v1.GoogleCloudAiplatformV1StudySpecArgs>;
    /**
     * The spec of a trial job. The same spec applies to the CustomJobs created in all the trials.
     */
    trialJobSpec: pulumi.Input<inputs.aiplatform.v1.GoogleCloudAiplatformV1CustomJobSpecArgs>;
}