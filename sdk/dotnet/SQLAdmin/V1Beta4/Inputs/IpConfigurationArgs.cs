// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4.Inputs
{

    /// <summary>
    /// IP Management configuration.
    /// </summary>
    public sealed class IpConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the allocated ip range for the private ip CloudSQL instance. For example: "google-managed-services-default". If set, the instance ip will be created in the allocated range. The range name must comply with [RFC 1035](https://tools.ietf.org/html/rfc1035). Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?.`
        /// </summary>
        [Input("allocatedIpRange")]
        public Input<string>? AllocatedIpRange { get; set; }

        [Input("authorizedNetworks")]
        private InputList<Inputs.AclEntryArgs>? _authorizedNetworks;

        /// <summary>
        /// The list of external networks that are allowed to connect to the instance using the IP. In 'CIDR' notation, also known as 'slash' notation (for example: `157.197.200.0/24`).
        /// </summary>
        public InputList<Inputs.AclEntryArgs> AuthorizedNetworks
        {
            get => _authorizedNetworks ?? (_authorizedNetworks = new InputList<Inputs.AclEntryArgs>());
            set => _authorizedNetworks = value;
        }

        /// <summary>
        /// Whether the instance is assigned a public IP address or not.
        /// </summary>
        [Input("ipv4Enabled")]
        public Input<bool>? Ipv4Enabled { get; set; }

        /// <summary>
        /// The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP. For example, `/projects/myProject/global/networks/default`. This setting can be updated, but it cannot be removed after it is set.
        /// </summary>
        [Input("privateNetwork")]
        public Input<string>? PrivateNetwork { get; set; }

        /// <summary>
        /// Whether SSL connections over IP are enforced or not.
        /// </summary>
        [Input("requireSsl")]
        public Input<bool>? RequireSsl { get; set; }

        public IpConfigurationArgs()
        {
        }
    }
}