// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSupport.V2Beta.Outputs
{

    /// <summary>
    /// An object containing information about the effective user and authenticated principal responsible for an action.
    /// </summary>
    [OutputType]
    public sealed class ActorResponse
    {
        /// <summary>
        /// The name to display for the actor. If not provided, it is inferred from credentials supplied during case creation. When an email is provided, a display name must also be provided. This will be obfuscated if the user is a Google Support agent.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The email address of the actor. If not provided, it is inferred from credentials supplied during case creation. If the authenticated principal does not have an email address, one must be provided. When a name is provided, an email must also be provided. This will be obfuscated if the user is a Google Support agent.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Whether the actor is a Google support actor.
        /// </summary>
        public readonly bool GoogleSupport;
        /// <summary>
        /// An ID representing the user that was authenticated when the corresponding action was taken. This will be an email address, if one is available, or some other unique ID. See https://cloud.google.com/docs/authentication for more information on types of authentication.
        /// </summary>
        public readonly string PrincipalId;

        [OutputConstructor]
        private ActorResponse(
            string displayName,

            string email,

            bool googleSupport,

            string principalId)
        {
            DisplayName = displayName;
            Email = email;
            GoogleSupport = googleSupport;
            PrincipalId = principalId;
        }
    }
}