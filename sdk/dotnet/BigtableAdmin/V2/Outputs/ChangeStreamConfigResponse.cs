// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// Change stream configuration.
    /// </summary>
    [OutputType]
    public sealed class ChangeStreamConfigResponse
    {
        /// <summary>
        /// How long the change stream should be retained. Change stream data older than the retention period will not be returned when reading the change stream from the table. Values must be at least 1 day and at most 7 days, and will be truncated to microsecond granularity.
        /// </summary>
        public readonly string RetentionPeriod;

        [OutputConstructor]
        private ChangeStreamConfigResponse(string retentionPeriod)
        {
            RetentionPeriod = retentionPeriod;
        }
    }
}