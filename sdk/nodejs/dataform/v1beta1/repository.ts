// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new Repository in a given project and location.
 * Auto-naming is currently not supported for this resource.
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataform/v1beta1:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * Optional. The repository's user-friendly name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. If set, configures this repository to be linked to a Git remote.
     */
    public readonly gitRemoteSettings!: pulumi.Output<outputs.dataform.v1beta1.GitRemoteSettingsResponse>;
    /**
     * Optional. Repository user labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * The repository's name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*&#47;secrets/*&#47;versions/*`. The file itself must be in a JSON format.
     */
    public readonly npmrcEnvironmentVariablesSecretVersion!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Required. The ID to use for the repository, which will become the final component of the repository's resource name.
     */
    public readonly repositoryId!: pulumi.Output<string>;
    /**
     * Optional. The service account to run workflow invocations under.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    /**
     * Optional. Input only. If set to true, the authenticated user will be granted the roles/dataform.admin role on the created repository. To modify access to the created repository later apply setIamPolicy from https://cloud.google.com/dataform/reference/rest#rest-resource:-v1beta1.projects.locations.repositories
     */
    public readonly setAuthenticatedUserAdmin!: pulumi.Output<boolean>;
    /**
     * Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
     */
    public readonly workspaceCompilationOverrides!: pulumi.Output<outputs.dataform.v1beta1.WorkspaceCompilationOverridesResponse>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.repositoryId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repositoryId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["gitRemoteSettings"] = args ? args.gitRemoteSettings : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["npmrcEnvironmentVariablesSecretVersion"] = args ? args.npmrcEnvironmentVariablesSecretVersion : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["repositoryId"] = args ? args.repositoryId : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["setAuthenticatedUserAdmin"] = args ? args.setAuthenticatedUserAdmin : undefined;
            resourceInputs["workspaceCompilationOverrides"] = args ? args.workspaceCompilationOverrides : undefined;
            resourceInputs["name"] = undefined /*out*/;
        } else {
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["gitRemoteSettings"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["npmrcEnvironmentVariablesSecretVersion"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["repositoryId"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["setAuthenticatedUserAdmin"] = undefined /*out*/;
            resourceInputs["workspaceCompilationOverrides"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project", "repositoryId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Repository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Optional. The repository's user-friendly name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Optional. If set, configures this repository to be linked to a Git remote.
     */
    gitRemoteSettings?: pulumi.Input<inputs.dataform.v1beta1.GitRemoteSettingsArgs>;
    /**
     * Optional. Repository user labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*&#47;secrets/*&#47;versions/*`. The file itself must be in a JSON format.
     */
    npmrcEnvironmentVariablesSecretVersion?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Required. The ID to use for the repository, which will become the final component of the repository's resource name.
     */
    repositoryId: pulumi.Input<string>;
    /**
     * Optional. The service account to run workflow invocations under.
     */
    serviceAccount?: pulumi.Input<string>;
    /**
     * Optional. Input only. If set to true, the authenticated user will be granted the roles/dataform.admin role on the created repository. To modify access to the created repository later apply setIamPolicy from https://cloud.google.com/dataform/reference/rest#rest-resource:-v1beta1.projects.locations.repositories
     */
    setAuthenticatedUserAdmin?: pulumi.Input<boolean>;
    /**
     * Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
     */
    workspaceCompilationOverrides?: pulumi.Input<inputs.dataform.v1beta1.WorkspaceCompilationOverridesArgs>;
}
