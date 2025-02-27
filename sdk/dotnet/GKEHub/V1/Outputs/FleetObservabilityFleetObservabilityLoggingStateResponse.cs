// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Outputs
{

    /// <summary>
    /// Feature state for logging feature.
    /// </summary>
    [OutputType]
    public sealed class FleetObservabilityFleetObservabilityLoggingStateResponse
    {
        /// <summary>
        /// The base feature state of fleet default log.
        /// </summary>
        public readonly Outputs.FleetObservabilityFleetObservabilityBaseFeatureStateResponse DefaultLog;
        /// <summary>
        /// The base feature state of fleet scope log.
        /// </summary>
        public readonly Outputs.FleetObservabilityFleetObservabilityBaseFeatureStateResponse ScopeLog;

        [OutputConstructor]
        private FleetObservabilityFleetObservabilityLoggingStateResponse(
            Outputs.FleetObservabilityFleetObservabilityBaseFeatureStateResponse defaultLog,

            Outputs.FleetObservabilityFleetObservabilityBaseFeatureStateResponse scopeLog)
        {
            DefaultLog = defaultLog;
            ScopeLog = scopeLog;
        }
    }
}
