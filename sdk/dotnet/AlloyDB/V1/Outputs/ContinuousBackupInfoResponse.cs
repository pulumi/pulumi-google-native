// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AlloyDB.V1.Outputs
{

    /// <summary>
    /// ContinuousBackupInfo describes the continuous backup properties of a cluster.
    /// </summary>
    [OutputType]
    public sealed class ContinuousBackupInfoResponse
    {
        /// <summary>
        /// The earliest restorable time that can be restored to. Output only field.
        /// </summary>
        public readonly string EarliestRestorableTime;
        /// <summary>
        /// When ContinuousBackup was most recently enabled. Set to null if ContinuousBackup is not enabled.
        /// </summary>
        public readonly string EnabledTime;
        /// <summary>
        /// The encryption information for the WALs and backups required for ContinuousBackup.
        /// </summary>
        public readonly Outputs.EncryptionInfoResponse EncryptionInfo;
        /// <summary>
        /// Days of the week on which a continuous backup is taken. Output only field. Ignored if passed into the request.
        /// </summary>
        public readonly ImmutableArray<string> Schedule;

        [OutputConstructor]
        private ContinuousBackupInfoResponse(
            string earliestRestorableTime,

            string enabledTime,

            Outputs.EncryptionInfoResponse encryptionInfo,

            ImmutableArray<string> schedule)
        {
            EarliestRestorableTime = earliestRestorableTime;
            EnabledTime = enabledTime;
            EncryptionInfo = encryptionInfo;
            Schedule = schedule;
        }
    }
}