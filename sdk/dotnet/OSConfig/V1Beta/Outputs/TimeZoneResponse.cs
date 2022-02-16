// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1Beta.Outputs
{

    /// <summary>
    /// Represents a time zone from the [IANA Time Zone Database](https://www.iana.org/time-zones).
    /// </summary>
    [OutputType]
    public sealed class TimeZoneResponse
    {
        /// <summary>
        /// Optional. IANA Time Zone Database version number, e.g. "2019a".
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private TimeZoneResponse(string version)
        {
            Version = version;
        }
    }
}