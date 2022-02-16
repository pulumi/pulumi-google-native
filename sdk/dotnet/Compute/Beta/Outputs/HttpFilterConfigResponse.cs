// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Outputs
{

    /// <summary>
    /// HttpFilterConfiguration supplies additional contextual settings for networkservices.HttpFilter resources enabled by Traffic Director.
    /// </summary>
    [OutputType]
    public sealed class HttpFilterConfigResponse
    {
        /// <summary>
        /// The configuration needed to enable the networkservices.HttpFilter resource. The configuration must be YAML formatted and only contain fields defined in the protobuf identified in configTypeUrl
        /// </summary>
        public readonly string Config;
        /// <summary>
        /// The fully qualified versioned proto3 type url of the protobuf that the filter expects for its contextual settings, for example: type.googleapis.com/google.protobuf.Struct
        /// </summary>
        public readonly string ConfigTypeUrl;
        /// <summary>
        /// Name of the networkservices.HttpFilter resource this configuration belongs to. This name must be known to the xDS client. Example: envoy.wasm
        /// </summary>
        public readonly string FilterName;

        [OutputConstructor]
        private HttpFilterConfigResponse(
            string config,

            string configTypeUrl,

            string filterName)
        {
            Config = config;
            ConfigTypeUrl = configTypeUrl;
            FilterName = filterName;
        }
    }
}