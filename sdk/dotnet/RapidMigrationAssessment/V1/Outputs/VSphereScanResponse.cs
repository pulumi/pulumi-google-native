// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.RapidMigrationAssessment.V1.Outputs
{

    /// <summary>
    /// Message describing a MC Source of type VSphere Scan.
    /// </summary>
    [OutputType]
    public sealed class VSphereScanResponse
    {
        /// <summary>
        /// reference to the corresponding VSphere Scan in MC Source.
        /// </summary>
        public readonly string CoreSource;

        [OutputConstructor]
        private VSphereScanResponse(string coreSource)
        {
            CoreSource = coreSource;
        }
    }
}