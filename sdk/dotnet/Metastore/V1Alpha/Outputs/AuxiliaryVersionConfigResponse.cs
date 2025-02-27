// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha.Outputs
{

    /// <summary>
    /// Configuration information for the auxiliary service versions.
    /// </summary>
    [OutputType]
    public sealed class AuxiliaryVersionConfigResponse
    {
        /// <summary>
        /// A mapping of Hive metastore configuration key-value pairs to apply to the auxiliary Hive metastore (configured in hive-site.xml) in addition to the primary version's overrides. If keys are present in both the auxiliary version's overrides and the primary version's overrides, the value from the auxiliary version's overrides takes precedence.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ConfigOverrides;
        /// <summary>
        /// The network configuration contains the endpoint URI(s) of the auxiliary Hive metastore service.
        /// </summary>
        public readonly Outputs.NetworkConfigResponse NetworkConfig;
        /// <summary>
        /// The Hive metastore version of the auxiliary service. It must be less than the primary Hive metastore service's version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private AuxiliaryVersionConfigResponse(
            ImmutableDictionary<string, string> configOverrides,

            Outputs.NetworkConfigResponse networkConfig,

            string version)
        {
            ConfigOverrides = configOverrides;
            NetworkConfig = networkConfig;
            Version = version;
        }
    }
}
