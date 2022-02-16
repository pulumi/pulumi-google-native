// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Time window specified for hourly operations.
    /// </summary>
    public sealed class ResourcePolicyHourlyCycleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines a schedule with units measured in hours. The value determines how many hours pass between the start of each cycle.
        /// </summary>
        [Input("hoursInCycle")]
        public Input<int>? HoursInCycle { get; set; }

        /// <summary>
        /// Time within the window to start the operations. It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public ResourcePolicyHourlyCycleArgs()
        {
        }
    }
}