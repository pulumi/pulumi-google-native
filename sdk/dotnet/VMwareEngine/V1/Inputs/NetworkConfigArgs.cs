// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMwareEngine.V1.Inputs
{

    /// <summary>
    /// Network configuration in the consumer project with which the peering has to be done.
    /// </summary>
    public sealed class NetworkConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Management CIDR used by VMware management appliances.
        /// </summary>
        [Input("managementCidr", required: true)]
        public Input<string> ManagementCidr { get; set; } = null!;

        /// <summary>
        /// Optional. The relative resource name of the VMware Engine network attached to the private cloud. Specify the name in the following form: `projects/{project}/locations/{location}/vmwareEngineNetworks/{vmware_engine_network_id}` where `{project}` can either be a project number or a project ID.
        /// </summary>
        [Input("vmwareEngineNetwork")]
        public Input<string>? VmwareEngineNetwork { get; set; }

        public NetworkConfigArgs()
        {
        }
        public static new NetworkConfigArgs Empty => new NetworkConfigArgs();
    }
}
