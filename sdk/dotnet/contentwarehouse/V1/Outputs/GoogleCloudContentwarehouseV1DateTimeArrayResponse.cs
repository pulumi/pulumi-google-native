// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.contentwarehouse.V1.Outputs
{

    /// <summary>
    /// DateTime values.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudContentwarehouseV1DateTimeArrayResponse
    {
        /// <summary>
        /// List of datetime values. Both OffsetDateTime and ZonedDateTime are supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleTypeDateTimeResponse> Values;

        [OutputConstructor]
        private GoogleCloudContentwarehouseV1DateTimeArrayResponse(ImmutableArray<Outputs.GoogleTypeDateTimeResponse> values)
        {
            Values = values;
        }
    }
}