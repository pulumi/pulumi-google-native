// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Stores a new object and metadata.
 */
export class BucketObject extends pulumi.CustomResource {
    /**
     * Get an existing BucketObject resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BucketObject {
        return new BucketObject(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:storage/v1:BucketObject';

    /**
     * Returns true if the given object is an instance of BucketObject.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketObject {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketObject.__pulumiType;
    }

    /**
     * Access controls on the object.
     */
    public readonly acl!: pulumi.Output<outputs.storage.v1.ObjectAccessControlResponse[]>;
    /**
     * The name of the bucket containing this object.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Cache-Control directive for the object data. If omitted, and the object is accessible to all anonymous users, the default will be public, max-age=3600.
     */
    public readonly cacheControl!: pulumi.Output<string>;
    /**
     * Number of underlying components that make up this object. Components are accumulated by compose operations.
     */
    public readonly componentCount!: pulumi.Output<number>;
    /**
     * Content-Disposition of the object data.
     */
    public readonly contentDisposition!: pulumi.Output<string>;
    /**
     * Content-Encoding of the object data.
     */
    public readonly contentEncoding!: pulumi.Output<string>;
    /**
     * Content-Language of the object data.
     */
    public readonly contentLanguage!: pulumi.Output<string>;
    /**
     * Content-Type of the object data. If an object is stored without a Content-Type, it is served as application/octet-stream.
     */
    public readonly contentType!: pulumi.Output<string>;
    /**
     * CRC32c checksum, as described in RFC 4960, Appendix B; encoded using base64 in big-endian byte order. For more information about using the CRC32c checksum, see Hashes and ETags: Best Practices.
     */
    public readonly crc32c!: pulumi.Output<string>;
    /**
     * A timestamp in RFC 3339 format specified by the user for an object.
     */
    public readonly customTime!: pulumi.Output<string>;
    /**
     * Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
     */
    public readonly customerEncryption!: pulumi.Output<outputs.storage.v1.BucketObjectCustomerEncryptionResponse>;
    /**
     * HTTP 1.1 Entity tag for the object.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is the loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false.
     */
    public readonly eventBasedHold!: pulumi.Output<boolean>;
    /**
     * The content generation of this object. Used for object versioning.
     */
    public readonly generation!: pulumi.Output<string>;
    /**
     * The kind of item this is. For objects, this is always storage#object.
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Not currently supported. Specifying the parameter causes the request to fail with status code 400 - Bad Request.
     */
    public readonly kmsKeyName!: pulumi.Output<string>;
    /**
     * MD5 hash of the data; encoded using base64. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
     */
    public readonly md5Hash!: pulumi.Output<string>;
    /**
     * Media download link.
     */
    public readonly mediaLink!: pulumi.Output<string>;
    /**
     * User-provided metadata, in key/value pairs.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The version of the metadata for this object at this generation. Used for preconditions and for detecting changes in metadata. A metageneration number is only meaningful in the context of a particular generation of a particular object.
     */
    public readonly metageneration!: pulumi.Output<string>;
    /**
     * The name of the object. Required if not specified by URL parameter.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The owner of the object. This will always be the uploader of the object.
     */
    public readonly owner!: pulumi.Output<outputs.storage.v1.BucketObjectOwnerResponse>;
    /**
     * A server-determined value that specifies the earliest time that the object's retention period expires. This value is in RFC 3339 format. Note 1: This field is not provided for objects with an active event-based hold, since retention expiration is unknown until the hold is removed. Note 2: This value can be provided even when temporary hold is set (so that the user can reason about policy without having to first unset the temporary hold).
     */
    public readonly retentionExpirationTime!: pulumi.Output<string>;
    /**
     * The link to this object.
     */
    public readonly selfLink!: pulumi.Output<string>;
    /**
     * Content-Length of the data in bytes.
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * Storage class of the object.
     */
    public readonly storageClass!: pulumi.Output<string>;
    /**
     * Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites. A common use case of this flag is regulatory investigations where objects need to be retained while the investigation is ongoing. Note that unlike event-based hold, temporary hold does not impact retention expiration time of an object.
     */
    public readonly temporaryHold!: pulumi.Output<boolean>;
    /**
     * The creation time of the object in RFC 3339 format.
     */
    public readonly timeCreated!: pulumi.Output<string>;
    /**
     * The deletion time of the object in RFC 3339 format. Will be returned if and only if this version of the object has been deleted.
     */
    public readonly timeDeleted!: pulumi.Output<string>;
    /**
     * The time at which the object's storage class was last changed. When the object is initially created, it will be set to timeCreated.
     */
    public readonly timeStorageClassUpdated!: pulumi.Output<string>;
    /**
     * The modification time of the object metadata in RFC 3339 format.
     */
    public readonly updated!: pulumi.Output<string>;

    /**
     * Create a BucketObject resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketObjectArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.bucket === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucket'");
            }
            resourceInputs["acl"] = args ? args.acl : undefined;
            resourceInputs["bucket"] = args ? args.bucket : undefined;
            resourceInputs["cacheControl"] = args ? args.cacheControl : undefined;
            resourceInputs["componentCount"] = args ? args.componentCount : undefined;
            resourceInputs["contentDisposition"] = args ? args.contentDisposition : undefined;
            resourceInputs["contentEncoding"] = args ? args.contentEncoding : undefined;
            resourceInputs["contentLanguage"] = args ? args.contentLanguage : undefined;
            resourceInputs["contentType"] = args ? args.contentType : undefined;
            resourceInputs["crc32c"] = args ? args.crc32c : undefined;
            resourceInputs["customTime"] = args ? args.customTime : undefined;
            resourceInputs["customerEncryption"] = args ? args.customerEncryption : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["eventBasedHold"] = args ? args.eventBasedHold : undefined;
            resourceInputs["generation"] = args ? args.generation : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["ifGenerationMatch"] = args ? args.ifGenerationMatch : undefined;
            resourceInputs["ifGenerationNotMatch"] = args ? args.ifGenerationNotMatch : undefined;
            resourceInputs["ifMetagenerationMatch"] = args ? args.ifMetagenerationMatch : undefined;
            resourceInputs["ifMetagenerationNotMatch"] = args ? args.ifMetagenerationNotMatch : undefined;
            resourceInputs["kind"] = args ? args.kind : undefined;
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["md5Hash"] = args ? args.md5Hash : undefined;
            resourceInputs["mediaLink"] = args ? args.mediaLink : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metageneration"] = args ? args.metageneration : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["predefinedAcl"] = args ? args.predefinedAcl : undefined;
            resourceInputs["projection"] = args ? args.projection : undefined;
            resourceInputs["provisionalUserProject"] = args ? args.provisionalUserProject : undefined;
            resourceInputs["retentionExpirationTime"] = args ? args.retentionExpirationTime : undefined;
            resourceInputs["selfLink"] = args ? args.selfLink : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["source"] = args ? args.source : undefined;
            resourceInputs["storageClass"] = args ? args.storageClass : undefined;
            resourceInputs["temporaryHold"] = args ? args.temporaryHold : undefined;
            resourceInputs["timeCreated"] = args ? args.timeCreated : undefined;
            resourceInputs["timeDeleted"] = args ? args.timeDeleted : undefined;
            resourceInputs["timeStorageClassUpdated"] = args ? args.timeStorageClassUpdated : undefined;
            resourceInputs["updated"] = args ? args.updated : undefined;
            resourceInputs["userProject"] = args ? args.userProject : undefined;
        } else {
            resourceInputs["acl"] = undefined /*out*/;
            resourceInputs["bucket"] = undefined /*out*/;
            resourceInputs["cacheControl"] = undefined /*out*/;
            resourceInputs["componentCount"] = undefined /*out*/;
            resourceInputs["contentDisposition"] = undefined /*out*/;
            resourceInputs["contentEncoding"] = undefined /*out*/;
            resourceInputs["contentLanguage"] = undefined /*out*/;
            resourceInputs["contentType"] = undefined /*out*/;
            resourceInputs["crc32c"] = undefined /*out*/;
            resourceInputs["customTime"] = undefined /*out*/;
            resourceInputs["customerEncryption"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["eventBasedHold"] = undefined /*out*/;
            resourceInputs["generation"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["kmsKeyName"] = undefined /*out*/;
            resourceInputs["md5Hash"] = undefined /*out*/;
            resourceInputs["mediaLink"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["metageneration"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["owner"] = undefined /*out*/;
            resourceInputs["retentionExpirationTime"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["storageClass"] = undefined /*out*/;
            resourceInputs["temporaryHold"] = undefined /*out*/;
            resourceInputs["timeCreated"] = undefined /*out*/;
            resourceInputs["timeDeleted"] = undefined /*out*/;
            resourceInputs["timeStorageClassUpdated"] = undefined /*out*/;
            resourceInputs["updated"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketObject.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a BucketObject resource.
 */
export interface BucketObjectArgs {
    /**
     * Access controls on the object.
     */
    acl?: pulumi.Input<pulumi.Input<inputs.storage.v1.ObjectAccessControlArgs>[]>;
    /**
     * The name of the bucket containing this object.
     */
    bucket: pulumi.Input<string>;
    /**
     * Cache-Control directive for the object data. If omitted, and the object is accessible to all anonymous users, the default will be public, max-age=3600.
     */
    cacheControl?: pulumi.Input<string>;
    /**
     * Number of underlying components that make up this object. Components are accumulated by compose operations.
     */
    componentCount?: pulumi.Input<number>;
    /**
     * Content-Disposition of the object data.
     */
    contentDisposition?: pulumi.Input<string>;
    /**
     * Content-Encoding of the object data.
     */
    contentEncoding?: pulumi.Input<string>;
    /**
     * Content-Language of the object data.
     */
    contentLanguage?: pulumi.Input<string>;
    /**
     * Content-Type of the object data. If an object is stored without a Content-Type, it is served as application/octet-stream.
     */
    contentType?: pulumi.Input<string>;
    /**
     * CRC32c checksum, as described in RFC 4960, Appendix B; encoded using base64 in big-endian byte order. For more information about using the CRC32c checksum, see Hashes and ETags: Best Practices.
     */
    crc32c?: pulumi.Input<string>;
    /**
     * A timestamp in RFC 3339 format specified by the user for an object.
     */
    customTime?: pulumi.Input<string>;
    /**
     * Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
     */
    customerEncryption?: pulumi.Input<inputs.storage.v1.BucketObjectCustomerEncryptionArgs>;
    /**
     * HTTP 1.1 Entity tag for the object.
     */
    etag?: pulumi.Input<string>;
    /**
     * Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is the loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false.
     */
    eventBasedHold?: pulumi.Input<boolean>;
    /**
     * The content generation of this object. Used for object versioning.
     */
    generation?: pulumi.Input<string>;
    /**
     * The ID of the object, including the bucket name, object name, and generation number.
     */
    id?: pulumi.Input<string>;
    /**
     * Makes the operation conditional on whether the object's current generation matches the given value. Setting to 0 makes the operation succeed only if there are no live versions of the object.
     */
    ifGenerationMatch?: pulumi.Input<string>;
    /**
     * Makes the operation conditional on whether the object's current generation does not match the given value. If no live object exists, the precondition fails. Setting to 0 makes the operation succeed only if there is a live version of the object.
     */
    ifGenerationNotMatch?: pulumi.Input<string>;
    /**
     * Makes the operation conditional on whether the object's current metageneration matches the given value.
     */
    ifMetagenerationMatch?: pulumi.Input<string>;
    /**
     * Makes the operation conditional on whether the object's current metageneration does not match the given value.
     */
    ifMetagenerationNotMatch?: pulumi.Input<string>;
    /**
     * The kind of item this is. For objects, this is always storage#object.
     */
    kind?: pulumi.Input<string>;
    /**
     * Not currently supported. Specifying the parameter causes the request to fail with status code 400 - Bad Request.
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * MD5 hash of the data; encoded using base64. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
     */
    md5Hash?: pulumi.Input<string>;
    /**
     * Media download link.
     */
    mediaLink?: pulumi.Input<string>;
    /**
     * User-provided metadata, in key/value pairs.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The version of the metadata for this object at this generation. Used for preconditions and for detecting changes in metadata. A metageneration number is only meaningful in the context of a particular generation of a particular object.
     */
    metageneration?: pulumi.Input<string>;
    /**
     * The name of the object. Required if not specified by URL parameter.
     */
    name?: pulumi.Input<string>;
    /**
     * The owner of the object. This will always be the uploader of the object.
     */
    owner?: pulumi.Input<inputs.storage.v1.BucketObjectOwnerArgs>;
    /**
     * Apply a predefined set of access controls to this object.
     */
    predefinedAcl?: pulumi.Input<string>;
    /**
     * Set of properties to return. Defaults to noAcl, unless the object resource specifies the acl property, when it defaults to full.
     */
    projection?: pulumi.Input<string>;
    /**
     * The project to be billed for this request if the target bucket is requester-pays bucket.
     */
    provisionalUserProject?: pulumi.Input<string>;
    /**
     * A server-determined value that specifies the earliest time that the object's retention period expires. This value is in RFC 3339 format. Note 1: This field is not provided for objects with an active event-based hold, since retention expiration is unknown until the hold is removed. Note 2: This value can be provided even when temporary hold is set (so that the user can reason about policy without having to first unset the temporary hold).
     */
    retentionExpirationTime?: pulumi.Input<string>;
    /**
     * The link to this object.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Content-Length of the data in bytes.
     */
    size?: pulumi.Input<string>;
    source?: pulumi.Input<pulumi.asset.Asset | pulumi.asset.Archive>;
    /**
     * Storage class of the object.
     */
    storageClass?: pulumi.Input<string>;
    /**
     * Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites. A common use case of this flag is regulatory investigations where objects need to be retained while the investigation is ongoing. Note that unlike event-based hold, temporary hold does not impact retention expiration time of an object.
     */
    temporaryHold?: pulumi.Input<boolean>;
    /**
     * The creation time of the object in RFC 3339 format.
     */
    timeCreated?: pulumi.Input<string>;
    /**
     * The deletion time of the object in RFC 3339 format. Will be returned if and only if this version of the object has been deleted.
     */
    timeDeleted?: pulumi.Input<string>;
    /**
     * The time at which the object's storage class was last changed. When the object is initially created, it will be set to timeCreated.
     */
    timeStorageClassUpdated?: pulumi.Input<string>;
    /**
     * The modification time of the object metadata in RFC 3339 format.
     */
    updated?: pulumi.Input<string>;
    /**
     * The project to be billed for this request. Required for Requester Pays buckets.
     */
    userProject?: pulumi.Input<string>;
}