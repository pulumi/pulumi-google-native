// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AlloyDB.V1Alpha.Outputs
{

    /// <summary>
    /// ContinuousBackupConfig describes the continuous backups recovery configurations of a cluster.
    /// </summary>
    [OutputType]
    public sealed class ContinuousBackupConfigResponse
    {
        /// <summary>
        /// Whether ContinuousBackup is enabled.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The encryption config can be specified to encrypt the backups with a customer-managed encryption key (CMEK). When this field is not specified, the backup will then use default encryption scheme to protect the user data.
        /// </summary>
        public readonly Outputs.EncryptionConfigResponse EncryptionConfig;
        /// <summary>
        /// The number of days that are eligible to restore from using PITR. To support the entire recovery window, backups and logs are retained for one day more than the recovery window. If not set, defaults to 14 days.
        /// </summary>
        public readonly int RecoveryWindowDays;

        [OutputConstructor]
        private ContinuousBackupConfigResponse(
            bool enabled,

            Outputs.EncryptionConfigResponse encryptionConfig,

            int recoveryWindowDays)
        {
            Enabled = enabled;
            EncryptionConfig = encryptionConfig;
            RecoveryWindowDays = recoveryWindowDays;
        }
    }
}