// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Workstations.V1Beta.Outputs
{

    /// <summary>
    /// A customer-specified encryption key for the Compute Engine resources of this workstation configuration.
    /// </summary>
    [OutputType]
    public sealed class CustomerEncryptionKeyResponse
    {
        /// <summary>
        /// The name of the encryption key that is stored in Google Cloud KMS, for example, `projects/PROJECT_ID/locations/REGION/keyRings/KEY_RING/cryptoKeys/KEY_NAME`.
        /// </summary>
        public readonly string KmsKey;
        /// <summary>
        /// The service account being used for the encryption request for the given KMS key. If absent, the Compute Engine default service account is used. However, it is recommended to use a separate service account and to follow KMS best practices mentioned at https://cloud.google.com/kms/docs/separation-of-duties
        /// </summary>
        public readonly string KmsKeyServiceAccount;

        [OutputConstructor]
        private CustomerEncryptionKeyResponse(
            string kmsKey,

            string kmsKeyServiceAccount)
        {
            KmsKey = kmsKey;
            KmsKeyServiceAccount = kmsKeyServiceAccount;
        }
    }
}