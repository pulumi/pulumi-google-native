// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2
{
    /// <summary>
    /// Create an NFS share.
    /// </summary>
    [GoogleNativeResourceType("google-native:baremetalsolution/v2:NfsShare")]
    public partial class NfsShare : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of allowed access points.
        /// </summary>
        [Output("allowedClients")]
        public Output<ImmutableArray<Outputs.AllowedClientResponse>> AllowedClients { get; private set; } = null!;

        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Immutable. The name of the NFS share.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// An identifier for the NFS share, generated by the backend. This field will be deprecated in the future, use `id` instead.
        /// </summary>
        [Output("nfsShareId")]
        public Output<string> NfsShareId { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The requested size, in GiB.
        /// </summary>
        [Output("requestedSizeGib")]
        public Output<string> RequestedSizeGib { get; private set; } = null!;

        /// <summary>
        /// The state of the NFS share.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Immutable. The storage type of the underlying volume.
        /// </summary>
        [Output("storageType")]
        public Output<string> StorageType { get; private set; } = null!;

        /// <summary>
        /// The underlying volume of the share. Created automatically during provisioning.
        /// </summary>
        [Output("volume")]
        public Output<string> Volume { get; private set; } = null!;


        /// <summary>
        /// Create a NfsShare resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NfsShare(string name, NfsShareArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:baremetalsolution/v2:NfsShare", name, args ?? new NfsShareArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NfsShare(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:baremetalsolution/v2:NfsShare", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "location",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NfsShare resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NfsShare Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NfsShare(name, id, options);
        }
    }

    public sealed class NfsShareArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedClients")]
        private InputList<Inputs.AllowedClientArgs>? _allowedClients;

        /// <summary>
        /// List of allowed access points.
        /// </summary>
        public InputList<Inputs.AllowedClientArgs> AllowedClients
        {
            get => _allowedClients ?? (_allowedClients = new InputList<Inputs.AllowedClientArgs>());
            set => _allowedClients = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. The name of the NFS share.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The requested size, in GiB.
        /// </summary>
        [Input("requestedSizeGib")]
        public Input<string>? RequestedSizeGib { get; set; }

        /// <summary>
        /// Immutable. The storage type of the underlying volume.
        /// </summary>
        [Input("storageType")]
        public Input<Pulumi.GoogleNative.BareMetalSolution.V2.NfsShareStorageType>? StorageType { get; set; }

        public NfsShareArgs()
        {
        }
        public static new NfsShareArgs Empty => new NfsShareArgs();
    }
}