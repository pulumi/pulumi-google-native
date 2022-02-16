// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1.Outputs
{

    /// <summary>
    /// The encryption options for the Cloud Composer environment and its dependencies.Supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*.
    /// </summary>
    [OutputType]
    public sealed class EncryptionConfigResponse
    {
        /// <summary>
        /// Optional. Customer-managed Encryption Key available through Google's Key Management Service. Cannot be updated. If not specified, Google-managed key will be used.
        /// </summary>
        public readonly string KmsKeyName;

        [OutputConstructor]
        private EncryptionConfigResponse(string kmsKeyName)
        {
            KmsKeyName = kmsKeyName;
        }
    }
}