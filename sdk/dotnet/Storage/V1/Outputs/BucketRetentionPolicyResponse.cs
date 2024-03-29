// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Storage.V1.Outputs
{

    /// <summary>
    /// The bucket's retention policy. The retention policy enforces a minimum retention time for all objects contained in the bucket, based on their creation time. Any attempt to overwrite or delete objects younger than the retention period will result in a PERMISSION_DENIED error. An unlocked retention policy can be modified or removed from the bucket via a storage.buckets.update operation. A locked retention policy cannot be removed or shortened in duration for the lifetime of the bucket. Attempting to remove or decrease period of a locked retention policy will result in a PERMISSION_DENIED error.
    /// </summary>
    [OutputType]
    public sealed class BucketRetentionPolicyResponse
    {
        /// <summary>
        /// Server-determined value that indicates the time from which policy was enforced and effective. This value is in RFC 3339 format.
        /// </summary>
        public readonly string EffectiveTime;
        /// <summary>
        /// Once locked, an object retention policy cannot be modified.
        /// </summary>
        public readonly bool IsLocked;
        /// <summary>
        /// The duration in seconds that objects need to be retained. Retention duration must be greater than zero and less than 100 years. Note that enforcement of retention periods less than a day is not guaranteed. Such periods should only be used for testing purposes.
        /// </summary>
        public readonly string RetentionPeriod;

        [OutputConstructor]
        private BucketRetentionPolicyResponse(
            string effectiveTime,

            bool isLocked,

            string retentionPeriod)
        {
            EffectiveTime = effectiveTime;
            IsLocked = isLocked;
            RetentionPeriod = retentionPeriod;
        }
    }
}
