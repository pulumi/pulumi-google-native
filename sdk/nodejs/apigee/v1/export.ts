// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Submit a data export job to be processed in the background. If the request is successful, the API returns a 201 status, a URI that can be used to retrieve the status of the export job, and the `state` value of "enqueued".
 * Note - this resource's API doesn't support deletion. When deleted, the resource will persist
 * on Google Cloud even though it will be deleted from Pulumi state.
 */
export class Export extends pulumi.CustomResource {
    /**
     * Get an existing Export resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Export {
        return new Export(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Export';

    /**
     * Returns true if the given object is an instance of Export.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Export {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Export.__pulumiType;
    }

    /**
     * Time the export job was created.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Name of the datastore that is the destination of the export job [datastore]
     */
    public readonly datastoreName!: pulumi.Output<string>;
    /**
     * Description of the export job.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Error is set when export fails
     */
    public /*out*/ readonly error!: pulumi.Output<string>;
    /**
     * Execution time for this export job. If the job is still in progress, it will be set to the amount of time that has elapsed since`created`, in seconds. Else, it will set to (`updated` - `created`), in seconds.
     */
    public /*out*/ readonly executionTime!: pulumi.Output<string>;
    /**
     * Display name of the export job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Self link of the export job. A URI that can be used to retrieve the status of an export job. Example: `/organizations/myorg/environments/myenv/analytics/exports/9cfc0d85-0f30-46d6-ae6f-318d0cb961bd`
     */
    public /*out*/ readonly self!: pulumi.Output<string>;
    /**
     * Status of the export job. Valid values include `enqueued`, `running`, `completed`, and `failed`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Time the export job was last updated.
     */
    public /*out*/ readonly updated!: pulumi.Output<string>;

    /**
     * Create a Export resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExportArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.datastoreName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datastoreName'");
            }
            if ((!args || args.dateRange === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dateRange'");
            }
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["csvDelimiter"] = args ? args.csvDelimiter : undefined;
            resourceInputs["datastoreName"] = args ? args.datastoreName : undefined;
            resourceInputs["dateRange"] = args ? args.dateRange : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["environmentId"] = args ? args.environmentId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["outputFormat"] = args ? args.outputFormat : undefined;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["executionTime"] = undefined /*out*/;
            resourceInputs["self"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
        } else {
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["datastoreName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["executionTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["self"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Export.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Export resource.
 */
export interface ExportArgs {
    /**
     * Optional. Delimiter used in the CSV file, if `outputFormat` is set to `csv`. Defaults to the `,` (comma) character. Supported delimiter characters include comma (`,`), pipe (`|`), and tab (`\t`).
     */
    csvDelimiter?: pulumi.Input<string>;
    /**
     * Name of the preconfigured datastore.
     */
    datastoreName: pulumi.Input<string>;
    /**
     * Date range of the data to export.
     */
    dateRange: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1DateRangeArgs>;
    /**
     * Optional. Description of the export job.
     */
    description?: pulumi.Input<string>;
    environmentId: pulumi.Input<string>;
    /**
     * Display name of the export job.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Optional. Output format of the export. Valid values include: `csv` or `json`. Defaults to `json`. Note: Configure the delimiter for CSV output using the `csvDelimiter` property.
     */
    outputFormat?: pulumi.Input<string>;
}