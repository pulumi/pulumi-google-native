// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.WebSecurityScanner.V1Beta.Outputs
{

    /// <summary>
    /// Output only. Defines a warning trace message for ScanRun. Warning traces provide customers with useful information that helps make the scanning process more effective.
    /// </summary>
    [OutputType]
    public sealed class ScanRunWarningTraceResponse
    {
        /// <summary>
        /// Indicates the warning code.
        /// </summary>
        public readonly string Code;

        [OutputConstructor]
        private ScanRunWarningTraceResponse(string code)
        {
            Code = code;
        }
    }
}