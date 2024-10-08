// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseHosting.V1Beta1
{
    public static class GetVersion
    {
        /// <summary>
        /// Get the specified version that has been created for the specified site. This can include versions that were created for the default `live` channel or for any active preview channels for the specified site.
        /// </summary>
        public static Task<GetVersionResult> InvokeAsync(GetVersionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVersionResult>("google-native:firebasehosting/v1beta1:getVersion", args ?? new GetVersionArgs(), options.WithDefaults());

        /// <summary>
        /// Get the specified version that has been created for the specified site. This can include versions that were created for the default `live` channel or for any active preview channels for the specified site.
        /// </summary>
        public static Output<GetVersionResult> Invoke(GetVersionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVersionResult>("google-native:firebasehosting/v1beta1:getVersion", args ?? new GetVersionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVersionArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public string? Project { get; set; }

        [Input("siteId", required: true)]
        public string SiteId { get; set; } = null!;

        [Input("versionId", required: true)]
        public string VersionId { get; set; } = null!;

        public GetVersionArgs()
        {
        }
        public static new GetVersionArgs Empty => new GetVersionArgs();
    }

    public sealed class GetVersionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("siteId", required: true)]
        public Input<string> SiteId { get; set; } = null!;

        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public GetVersionInvokeArgs()
        {
        }
        public static new GetVersionInvokeArgs Empty => new GetVersionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVersionResult
    {
        /// <summary>
        /// The configuration for the behavior of the site. This configuration exists in the [`firebase.json`](https://firebase.google.com/docs/cli/#the_firebasejson_file) file.
        /// </summary>
        public readonly Outputs.ServingConfigResponse Config;
        /// <summary>
        /// The time at which the version was created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// Identifies the user who created the version.
        /// </summary>
        public readonly Outputs.ActingUserResponse CreateUser;
        /// <summary>
        /// The time at which the version was `DELETED`.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// Identifies the user who `DELETED` the version.
        /// </summary>
        public readonly Outputs.ActingUserResponse DeleteUser;
        /// <summary>
        /// The total number of files associated with the version. This value is calculated after a version is `FINALIZED`.
        /// </summary>
        public readonly string FileCount;
        /// <summary>
        /// The time at which the version was `FINALIZED`.
        /// </summary>
        public readonly string FinalizeTime;
        /// <summary>
        /// Identifies the user who `FINALIZED` the version.
        /// </summary>
        public readonly Outputs.ActingUserResponse FinalizeUser;
        /// <summary>
        /// The labels used for extra metadata and/or filtering.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The fully-qualified resource name for the version, in the format: sites/ SITE_ID/versions/VERSION_ID This name is provided in the response body when you call [`CreateVersion`](sites.versions/create).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The deploy status of the version. For a successful deploy, call [`CreateVersion`](sites.versions/create) to make a new version (`CREATED` status), [upload all desired files](sites.versions/populateFiles) to the version, then [update](sites.versions/patch) the version to the `FINALIZED` status. Note that if you leave the version in the `CREATED` state for more than 12 hours, the system will automatically mark the version as `ABANDONED`. You can also change the status of a version to `DELETED` by calling [`DeleteVersion`](sites.versions/delete).
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The total stored bytesize of the version. This value is calculated after a version is `FINALIZED`.
        /// </summary>
        public readonly string VersionBytes;

        [OutputConstructor]
        private GetVersionResult(
            Outputs.ServingConfigResponse config,

            string createTime,

            Outputs.ActingUserResponse createUser,

            string deleteTime,

            Outputs.ActingUserResponse deleteUser,

            string fileCount,

            string finalizeTime,

            Outputs.ActingUserResponse finalizeUser,

            ImmutableDictionary<string, string> labels,

            string name,

            string status,

            string versionBytes)
        {
            Config = config;
            CreateTime = createTime;
            CreateUser = createUser;
            DeleteTime = deleteTime;
            DeleteUser = deleteUser;
            FileCount = fileCount;
            FinalizeTime = finalizeTime;
            FinalizeUser = finalizeUser;
            Labels = labels;
            Name = name;
            Status = status;
            VersionBytes = versionBytes;
        }
    }
}
