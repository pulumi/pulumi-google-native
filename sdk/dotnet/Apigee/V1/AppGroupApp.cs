// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Apigee.V1
{
    /// <summary>
    /// Creates an app and associates it with an AppGroup. This API associates the AppGroup app with the specified API product and auto-generates an API key for the app to use in calls to API proxies inside that API product. The `name` is the unique ID of the app that you can use in API calls.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:apigee/v1:AppGroupApp")]
    public partial class AppGroupApp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of API products associated with the AppGroup app.
        /// </summary>
        [Output("apiProducts")]
        public Output<ImmutableArray<string>> ApiProducts { get; private set; } = null!;

        /// <summary>
        /// Immutable. Name of the parent AppGroup whose resource name format is of syntax (organizations/*/appgroups/*).
        /// </summary>
        [Output("appGroup")]
        public Output<string> AppGroup { get; private set; } = null!;

        /// <summary>
        /// Immutable. ID of the AppGroup app.
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("appgroupId")]
        public Output<string> AppgroupId { get; private set; } = null!;

        /// <summary>
        /// List of attributes for the AppGroup app.
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<Outputs.GoogleCloudApigeeV1AttributeResponse>> Attributes { get; private set; } = null!;

        /// <summary>
        /// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to AppGroup apps.
        /// </summary>
        [Output("callbackUrl")]
        public Output<string> CallbackUrl { get; private set; } = null!;

        /// <summary>
        /// Time the AppGroup app was created in milliseconds since epoch.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Set of credentials for the AppGroup app consisting of the consumer key/secret pairs associated with the API products.
        /// </summary>
        [Output("credentials")]
        public Output<ImmutableArray<Outputs.GoogleCloudApigeeV1CredentialResponse>> Credentials { get; private set; } = null!;

        /// <summary>
        /// Immutable. Expiration time, in seconds, for the consumer key that is generated for the AppGroup app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
        /// </summary>
        [Output("keyExpiresIn")]
        public Output<string> KeyExpiresIn { get; private set; } = null!;

        /// <summary>
        /// Time the AppGroup app was modified in milliseconds since epoch.
        /// </summary>
        [Output("lastModifiedAt")]
        public Output<string> LastModifiedAt { get; private set; } = null!;

        /// <summary>
        /// Immutable. Name of the AppGroup app whose resource name format is of syntax (organizations/*/appgroups/*/apps/*).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("organizationId")]
        public Output<string> OrganizationId { get; private set; } = null!;

        /// <summary>
        /// Scopes to apply to the AppGroup app. The specified scopes must already exist for the API product that you associate with the AppGroup app.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Status of the App. Valid values include `approved` or `revoked`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a AppGroupApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppGroupApp(string name, AppGroupAppArgs args, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:AppGroupApp", name, args ?? new AppGroupAppArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppGroupApp(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:apigee/v1:AppGroupApp", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "appgroupId",
                    "organizationId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AppGroupApp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppGroupApp Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppGroupApp(name, id, options);
        }
    }

    public sealed class AppGroupAppArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiProducts")]
        private InputList<string>? _apiProducts;

        /// <summary>
        /// List of API products associated with the AppGroup app.
        /// </summary>
        public InputList<string> ApiProducts
        {
            get => _apiProducts ?? (_apiProducts = new InputList<string>());
            set => _apiProducts = value;
        }

        /// <summary>
        /// Immutable. Name of the parent AppGroup whose resource name format is of syntax (organizations/*/appgroups/*).
        /// </summary>
        [Input("appGroup")]
        public Input<string>? AppGroup { get; set; }

        /// <summary>
        /// Immutable. ID of the AppGroup app.
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("appgroupId", required: true)]
        public Input<string> AppgroupId { get; set; } = null!;

        [Input("attributes")]
        private InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>? _attributes;

        /// <summary>
        /// List of attributes for the AppGroup app.
        /// </summary>
        public InputList<Inputs.GoogleCloudApigeeV1AttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.GoogleCloudApigeeV1AttributeArgs>());
            set => _attributes = value;
        }

        /// <summary>
        /// Callback URL used by OAuth 2.0 authorization servers to communicate authorization codes back to AppGroup apps.
        /// </summary>
        [Input("callbackUrl")]
        public Input<string>? CallbackUrl { get; set; }

        /// <summary>
        /// Immutable. Expiration time, in seconds, for the consumer key that is generated for the AppGroup app. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
        /// </summary>
        [Input("keyExpiresIn")]
        public Input<string>? KeyExpiresIn { get; set; }

        /// <summary>
        /// Immutable. Name of the AppGroup app whose resource name format is of syntax (organizations/*/appgroups/*/apps/*).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("organizationId", required: true)]
        public Input<string> OrganizationId { get; set; } = null!;

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// Scopes to apply to the AppGroup app. The specified scopes must already exist for the API product that you associate with the AppGroup app.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Status of the App. Valid values include `approved` or `revoked`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AppGroupAppArgs()
        {
        }
        public static new AppGroupAppArgs Empty => new AppGroupAppArgs();
    }
}