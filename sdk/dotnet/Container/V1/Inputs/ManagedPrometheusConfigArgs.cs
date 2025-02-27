// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1.Inputs
{

    /// <summary>
    /// ManagedPrometheusConfig defines the configuration for Google Cloud Managed Service for Prometheus.
    /// </summary>
    public sealed class ManagedPrometheusConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable Managed Collection.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        public ManagedPrometheusConfigArgs()
        {
        }
        public static new ManagedPrometheusConfigArgs Empty => new ManagedPrometheusConfigArgs();
    }
}
