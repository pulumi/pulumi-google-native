// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Alpha.Outputs
{

    /// <summary>
    /// Describes authentication configuration that uses a Google account.
    /// </summary>
    [OutputType]
    public sealed class GoogleAccountResponse
    {
        /// <summary>
        /// Input only. The password of the Google account. The credential is stored encrypted and not returned in any response nor included in audit logs.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The user name of the Google account.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private GoogleAccountResponse(
            string password,

            string username)
        {
            Password = password;
            Username = username;
        }
    }
}