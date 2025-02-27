// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataform.V1Beta1
{
    public static class GetRepository
    {
        /// <summary>
        /// Fetches a single Repository.
        /// </summary>
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("google-native:dataform/v1beta1:getRepository", args ?? new GetRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// Fetches a single Repository.
        /// </summary>
        public static Output<GetRepositoryResult> Invoke(GetRepositoryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryResult>("google-native:dataform/v1beta1:getRepository", args ?? new GetRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("repositoryId", required: true)]
        public string RepositoryId { get; set; } = null!;

        public GetRepositoryArgs()
        {
        }
        public static new GetRepositoryArgs Empty => new GetRepositoryArgs();
    }

    public sealed class GetRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("repositoryId", required: true)]
        public Input<string> RepositoryId { get; set; } = null!;

        public GetRepositoryInvokeArgs()
        {
        }
        public static new GetRepositoryInvokeArgs Empty => new GetRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
        /// <summary>
        /// Optional. The repository's user-friendly name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. If set, configures this repository to be linked to a Git remote.
        /// </summary>
        public readonly Outputs.GitRemoteSettingsResponse GitRemoteSettings;
        /// <summary>
        /// Optional. Repository user labels.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Labels;
        /// <summary>
        /// The repository's name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. The name of the Secret Manager secret version to be used to interpolate variables into the .npmrc file for package installation operations. Must be in the format `projects/*/secrets/*/versions/*`. The file itself must be in a JSON format.
        /// </summary>
        public readonly string NpmrcEnvironmentVariablesSecretVersion;
        /// <summary>
        /// Optional. The service account to run workflow invocations under.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Optional. Input only. If set to true, the authenticated user will be granted the roles/dataform.admin role on the created repository. To modify access to the created repository later apply setIamPolicy from https://cloud.google.com/dataform/reference/rest#rest-resource:-v1beta1.projects.locations.repositories
        /// </summary>
        public readonly bool SetAuthenticatedUserAdmin;
        /// <summary>
        /// Optional. If set, fields of `workspace_compilation_overrides` override the default compilation settings that are specified in dataform.json when creating workspace-scoped compilation results. See documentation for `WorkspaceCompilationOverrides` for more information.
        /// </summary>
        public readonly Outputs.WorkspaceCompilationOverridesResponse WorkspaceCompilationOverrides;

        [OutputConstructor]
        private GetRepositoryResult(
            string displayName,

            Outputs.GitRemoteSettingsResponse gitRemoteSettings,

            ImmutableDictionary<string, string> labels,

            string name,

            string npmrcEnvironmentVariablesSecretVersion,

            string serviceAccount,

            bool setAuthenticatedUserAdmin,

            Outputs.WorkspaceCompilationOverridesResponse workspaceCompilationOverrides)
        {
            DisplayName = displayName;
            GitRemoteSettings = gitRemoteSettings;
            Labels = labels;
            Name = name;
            NpmrcEnvironmentVariablesSecretVersion = npmrcEnvironmentVariablesSecretVersion;
            ServiceAccount = serviceAccount;
            SetAuthenticatedUserAdmin = setAuthenticatedUserAdmin;
            WorkspaceCompilationOverrides = workspaceCompilationOverrides;
        }
    }
}
