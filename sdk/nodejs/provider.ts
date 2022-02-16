// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the Google Cloud package.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'google-native';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * The default project to manage resources in. If another project is specified on a resource, it will take precedence.
     */
    public readonly project!: pulumi.Output<string | undefined>;
    /**
     * The default region to manage resources in. If another region is specified on a regional resource, it will take precedence.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The default zone to manage resources in. Generally, this zone should be within the default region you specified. If another zone is specified on a zonal resource, it will take precedence.
     */
    public readonly zone!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["appendUserAgent"] = (args ? args.appendUserAgent : undefined) ?? utilities.getEnv("GOOGLE_APPEND_USER_AGENT");
            resourceInputs["disablePartnerName"] = pulumi.output((args ? args.disablePartnerName : undefined) ?? utilities.getEnvBoolean("GOOGLE_DISABLE_PARTNER_NAME")).apply(JSON.stringify);
            resourceInputs["partnerName"] = (args ? args.partnerName : undefined) ?? utilities.getEnv("GOOGLE_PARTNER_NAME");
            resourceInputs["project"] = (args ? args.project : undefined) ?? utilities.getEnv("GOOGLE_PROJECT", "GOOGLE_CLOUD_PROJECT", "GCLOUD_PROJECT", "CLOUDSDK_CORE_PROJECT");
            resourceInputs["region"] = (args ? args.region : undefined) ?? utilities.getEnv("GOOGLE_REGION", "GCLOUD_REGION", "CLOUDSDK_COMPUTE_REGION");
            resourceInputs["zone"] = (args ? args.zone : undefined) ?? utilities.getEnv("GOOGLE_ZONE", "GCLOUD_ZONE", "CLOUDSDK_COMPUTE_ZONE");
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Additional user-agent string to append to the default one (<prod_name>/<ver>).
     */
    appendUserAgent?: pulumi.Input<string>;
    /**
     * This will disable the Pulumi Partner Name which is used if a custom `partnerName` isn't specified.
     */
    disablePartnerName?: pulumi.Input<boolean>;
    /**
     * A Google Partner Name to facilitate partner resource usage attribution.
     */
    partnerName?: pulumi.Input<string>;
    /**
     * The default project to manage resources in. If another project is specified on a resource, it will take precedence.
     */
    project?: pulumi.Input<string>;
    /**
     * The default region to manage resources in. If another region is specified on a regional resource, it will take precedence.
     */
    region?: pulumi.Input<string>;
    /**
     * The default zone to manage resources in. Generally, this zone should be within the default region you specified. If another zone is specified on a zonal resource, it will take precedence.
     */
    zone?: pulumi.Input<string>;
}