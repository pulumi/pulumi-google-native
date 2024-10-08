// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V1.Outputs
{

    /// <summary>
    /// Preview: Parameter value applied to the aggregation function. This is a preview feature and may be subject to change before final release.
    /// </summary>
    [OutputType]
    public sealed class ParameterResponse
    {
        /// <summary>
        /// A floating-point parameter value.
        /// </summary>
        public readonly double DoubleValue;
        /// <summary>
        /// An integer parameter value.
        /// </summary>
        public readonly string IntValue;

        [OutputConstructor]
        private ParameterResponse(
            double doubleValue,

            string intValue)
        {
            DoubleValue = doubleValue;
            IntValue = intValue;
        }
    }
}
