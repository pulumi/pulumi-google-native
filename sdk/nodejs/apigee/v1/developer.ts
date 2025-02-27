// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a developer. Once created, the developer can register an app and obtain an API key. At creation time, a developer is set as `active`. To change the developer status, use the SetDeveloperStatus API.
 * Auto-naming is currently not supported for this resource.
 */
export class Developer extends pulumi.CustomResource {
    /**
     * Get an existing Developer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Developer {
        return new Developer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Developer';

    /**
     * Returns true if the given object is an instance of Developer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Developer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Developer.__pulumiType;
    }

    /**
     * Access type.
     */
    public readonly accessType!: pulumi.Output<string>;
    /**
     * Developer app family.
     */
    public readonly appFamily!: pulumi.Output<string>;
    /**
     * List of apps associated with the developer.
     */
    public readonly apps!: pulumi.Output<string[]>;
    /**
     * Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
     */
    public readonly attributes!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1AttributeResponse[]>;
    /**
     * List of companies associated with the developer.
     */
    public readonly companies!: pulumi.Output<string[]>;
    /**
     * Time at which the developer was created in milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
     */
    public readonly developerId!: pulumi.Output<string>;
    /**
     * Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * First name of the developer.
     */
    public readonly firstName!: pulumi.Output<string>;
    /**
     * Time at which the developer was last modified in milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Last name of the developer.
     */
    public readonly lastName!: pulumi.Output<string>;
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * Name of the Apigee organization in which the developer resides.
     */
    public /*out*/ readonly organizationName!: pulumi.Output<string>;
    /**
     * Status of the developer. Valid values are `active` and `inactive`.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * User name of the developer. Not used by Apigee hybrid.
     */
    public readonly userName!: pulumi.Output<string>;

    /**
     * Create a Developer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeveloperArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.firstName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'firstName'");
            }
            if ((!args || args.lastName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'lastName'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.userName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userName'");
            }
            resourceInputs["accessType"] = args ? args.accessType : undefined;
            resourceInputs["appFamily"] = args ? args.appFamily : undefined;
            resourceInputs["apps"] = args ? args.apps : undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["companies"] = args ? args.companies : undefined;
            resourceInputs["developerId"] = args ? args.developerId : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["firstName"] = args ? args.firstName : undefined;
            resourceInputs["lastName"] = args ? args.lastName : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["organizationName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        } else {
            resourceInputs["accessType"] = undefined /*out*/;
            resourceInputs["appFamily"] = undefined /*out*/;
            resourceInputs["apps"] = undefined /*out*/;
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["companies"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["developerId"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
            resourceInputs["firstName"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["lastName"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["organizationName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["userName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["organizationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Developer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Developer resource.
 */
export interface DeveloperArgs {
    /**
     * Access type.
     */
    accessType?: pulumi.Input<string>;
    /**
     * Developer app family.
     */
    appFamily?: pulumi.Input<string>;
    /**
     * List of apps associated with the developer.
     */
    apps?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Optional. Developer attributes (name/value pairs). The custom attribute limit is 18.
     */
    attributes?: pulumi.Input<pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1AttributeArgs>[]>;
    /**
     * List of companies associated with the developer.
     */
    companies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * ID of the developer. **Note**: IDs are generated internally by Apigee and are not guaranteed to stay the same over time.
     */
    developerId?: pulumi.Input<string>;
    /**
     * Email address of the developer. This value is used to uniquely identify the developer in Apigee hybrid. Note that the email address has to be in lowercase only.
     */
    email: pulumi.Input<string>;
    /**
     * First name of the developer.
     */
    firstName: pulumi.Input<string>;
    /**
     * Last name of the developer.
     */
    lastName: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * User name of the developer. Not used by Apigee hybrid.
     */
    userName: pulumi.Input<string>;
}
