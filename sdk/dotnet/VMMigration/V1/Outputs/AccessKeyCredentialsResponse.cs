// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1.Outputs
{

    /// <summary>
    /// Message describing AWS Credentials using access key id and secret.
    /// </summary>
    [OutputType]
    public sealed class AccessKeyCredentialsResponse
    {
        /// <summary>
        /// AWS access key ID.
        /// </summary>
        public readonly string AccessKeyId;
        /// <summary>
        /// Input only. AWS secret access key.
        /// </summary>
        public readonly string SecretAccessKey;

        [OutputConstructor]
        private AccessKeyCredentialsResponse(
            string accessKeyId,

            string secretAccessKey)
        {
            AccessKeyId = accessKeyId;
            SecretAccessKey = secretAccessKey;
        }
    }
}