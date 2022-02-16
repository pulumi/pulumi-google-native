// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new Hosting Site in the specified parent Firebase project. Note that Hosting sites can take several minutes to propagate through Firebase systems.
 * Auto-naming is currently not supported for this resource.
 */
export class Site extends pulumi.CustomResource {
    /**
     * Get an existing Site resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Site {
        return new Site(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firebasehosting/v1beta1:Site';

    /**
     * Returns true if the given object is an instance of Site.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Site {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Site.__pulumiType;
    }

    /**
     * Optional. The [ID of a Web App](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id) associated with the Hosting site.
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * The default URL for the Hosting site.
     */
    public /*out*/ readonly defaultUrl!: pulumi.Output<string>;
    /**
     * Optional. User-specified labels for the Hosting site.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The fully-qualified resource name of the Hosting site, in the format: projects/PROJECT_IDENTIFIER/sites/SITE_ID PROJECT_IDENTIFIER: the Firebase project's [`ProjectNumber`](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The type of Hosting site. Every Firebase project has a `DEFAULT_SITE`, which is created when Hosting is provisioned for the project. All additional sites are `USER_SITE`.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Site resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SiteArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.siteId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'siteId'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["siteId"] = args ? args.siteId : undefined;
            resourceInputs["defaultUrl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["defaultUrl"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Site.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Site resource.
 */
export interface SiteArgs {
    /**
     * Optional. The [ID of a Web App](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id) associated with the Hosting site.
     */
    appId?: pulumi.Input<string>;
    /**
     * Optional. User-specified labels for the Hosting site.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    project?: pulumi.Input<string>;
    siteId: pulumi.Input<string>;
}