// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BareMetalSolution.V2
{
    public static class GetNfsShare
    {
        /// <summary>
        /// Get details of a single NFS share.
        /// </summary>
        public static Task<GetNfsShareResult> InvokeAsync(GetNfsShareArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNfsShareResult>("google-native:baremetalsolution/v2:getNfsShare", args ?? new GetNfsShareArgs(), options.WithDefaults());

        /// <summary>
        /// Get details of a single NFS share.
        /// </summary>
        public static Output<GetNfsShareResult> Invoke(GetNfsShareInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNfsShareResult>("google-native:baremetalsolution/v2:getNfsShare", args ?? new GetNfsShareInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNfsShareArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("nfsShareId", required: true)]
        public string NfsShareId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetNfsShareArgs()
        {
        }
        public static new GetNfsShareArgs Empty => new GetNfsShareArgs();
    }

    public sealed class GetNfsShareInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("nfsShareId", required: true)]
        public Input<string> NfsShareId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetNfsShareInvokeArgs()
        {
        }
        public static new GetNfsShareInvokeArgs Empty => new GetNfsShareInvokeArgs();
    }


    [OutputType]
    public sealed class GetNfsShareResult
    {
        /// <summary>
        /// List of allowed access points.
        /// </summary>
        public readonly ImmutableArray<Outputs.AllowedClientResponse> AllowedClients;
        /// <summary>
        /// Labels as key value pairs.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// Immutable. The name of the NFS share.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// An identifier for the NFS share, generated by the backend. This field will be deprecated in the future, use `id` instead.
        /// </summary>
        public readonly string NfsShareId;
        /// <summary>
        /// The requested size, in GiB.
        /// </summary>
        public readonly string RequestedSizeGib;
        /// <summary>
        /// The state of the NFS share.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Immutable. The storage type of the underlying volume.
        /// </summary>
        public readonly string StorageType;
        /// <summary>
        /// The underlying volume of the share. Created automatically during provisioning.
        /// </summary>
        public readonly string Volume;

        [OutputConstructor]
        private GetNfsShareResult(
            ImmutableArray<Outputs.AllowedClientResponse> allowedClients,

            ImmutableDictionary<string, string> labels,

            string name,

            string nfsShareId,

            string requestedSizeGib,

            string state,

            string storageType,

            string volume)
        {
            AllowedClients = allowedClients;
            Labels = labels;
            Name = name;
            NfsShareId = nfsShareId;
            RequestedSizeGib = requestedSizeGib;
            State = state;
            StorageType = storageType;
            Volume = volume;
        }
    }
}