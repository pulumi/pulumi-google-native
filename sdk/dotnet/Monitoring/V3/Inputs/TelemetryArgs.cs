// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// Configuration for how to query telemetry on a Service.
    /// </summary>
    public sealed class TelemetryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full name of the resource that defines this service. Formatted as described in https://cloud.google.com/apis/design/resource_names.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        public TelemetryArgs()
        {
        }
        public static new TelemetryArgs Empty => new TelemetryArgs();
    }
}
