// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new managed folder.
 */
export class ManagedFolder extends pulumi.CustomResource {
    /**
     * Get an existing ManagedFolder resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagedFolder {
        return new ManagedFolder(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storage/v1:ManagedFolder';

    /**
     * Returns true if the given object is an instance of ManagedFolder.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagedFolder {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagedFolder.__pulumiType;
    }

    public readonly bucket!: pulumi.Output<string>;
    /**
     * The creation time of the managed folder in RFC 3339 format.
     */
    public readonly createTime!: pulumi.Output<string>;
    /**
     * The kind of item this is. For managed folders, this is always storage#managedFolder.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * The version of the metadata for this managed folder. Used for preconditions and for detecting changes in metadata.
     */
    public readonly metageneration!: pulumi.Output<string>;
    /**
     * The name of the managed folder. Required if not specified by URL parameter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The link to this managed folder.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * The last update time of the managed folder metadata in RFC 3339 format.
     */
    public readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ManagedFolder resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagedFolderArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["createTime"] = args ? args.createTime : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["metageneration"] = args ? args.metageneration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["selfLink"] = args ? args.selfLink : undefined;
            resourceInputs["updateTime"] = args ? args.updateTime : undefined;
        } else {
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["metageneration"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["bucket"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ManagedFolder.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagedFolder resource.
 */
export interface ManagedFolderArgs {
    /**
     * The name of the bucket containing this managed folder.
     */
    bucket: pulumi.Input<string>;
    /**
     * The creation time of the managed folder in RFC 3339 format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The ID of the managed folder, including the bucket name and managed folder name.
     */
    id?: pulumi.Input<string>;
    /**
     * The kind of item this is. For managed folders, this is always storage#managedFolder.
     */
    kind?: pulumi.Input<string>;
    /**
     * The version of the metadata for this managed folder. Used for preconditions and for detecting changes in metadata.
     */
    metageneration?: pulumi.Input<string>;
    /**
     * The name of the managed folder. Required if not specified by URL parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * The link to this managed folder.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The last update time of the managed folder metadata in RFC 3339 format.
     */
    updateTime?: pulumi.Input<string>;
}
