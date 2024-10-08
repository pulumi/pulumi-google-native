// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Outputs
{

    /// <summary>
    /// Details for an outcome with a SKIPPED outcome summary.
    /// </summary>
    [OutputType]
    public sealed class SkippedDetailResponse
    {
        /// <summary>
        /// If the App doesn't support the specific API level.
        /// </summary>
        public readonly bool IncompatibleAppVersion;
        /// <summary>
        /// If the App doesn't run on the specific architecture, for example, x86.
        /// </summary>
        public readonly bool IncompatibleArchitecture;
        /// <summary>
        /// If the requested OS version doesn't run on the specific device model.
        /// </summary>
        public readonly bool IncompatibleDevice;

        [OutputConstructor]
        private SkippedDetailResponse(
            bool incompatibleAppVersion,

            bool incompatibleArchitecture,

            bool incompatibleDevice)
        {
            IncompatibleAppVersion = incompatibleAppVersion;
            IncompatibleArchitecture = incompatibleArchitecture;
            IncompatibleDevice = incompatibleDevice;
        }
    }
}
