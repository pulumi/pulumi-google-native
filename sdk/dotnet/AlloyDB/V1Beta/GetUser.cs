// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AlloyDB.V1Beta
{
    public static class GetUser
    {
        /// <summary>
        /// Gets details of a single User.
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("google-native:alloydb/v1beta:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Gets details of a single User.
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("google-native:alloydb/v1beta:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        [Input("userId", required: true)]
        public string UserId { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// Optional. List of database roles this user has. The database role strings are subject to the PostgreSQL naming conventions.
        /// </summary>
        public readonly ImmutableArray<string> DatabaseRoles;
        /// <summary>
        /// Name of the resource in the form of projects/{project}/locations/{location}/cluster/{cluster}/users/{user}.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Input only. Password for the user.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Optional. Type of this user.
        /// </summary>
        public readonly string UserType;

        [OutputConstructor]
        private GetUserResult(
            ImmutableArray<string> databaseRoles,

            string name,

            string password,

            string userType)
        {
            DatabaseRoles = databaseRoles;
            Name = name;
            Password = password;
            UserType = userType;
        }
    }
}
