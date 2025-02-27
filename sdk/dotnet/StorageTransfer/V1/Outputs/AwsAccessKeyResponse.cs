// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Outputs
{

    /// <summary>
    /// AWS access key (see [AWS Security Credentials](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html)). For information on our data retention policy for user credentials, see [User credentials](/storage-transfer/docs/data-retention#user-credentials).
    /// </summary>
    [OutputType]
    public sealed class AwsAccessKeyResponse
    {
        /// <summary>
        /// AWS access key ID.
        /// </summary>
        public readonly string AccessKeyId;
        /// <summary>
        /// AWS secret access key. This field is not returned in RPC responses.
        /// </summary>
        public readonly string SecretAccessKey;

        [OutputConstructor]
        private AwsAccessKeyResponse(
            string accessKeyId,

            string secretAccessKey)
        {
            AccessKeyId = accessKeyId;
            SecretAccessKey = secretAccessKey;
        }
    }
}
