// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BeyondCorp.V1.Inputs
{

    /// <summary>
    /// The peered VPC owned by the consumer project.
    /// </summary>
    public sealed class PeeredVpcArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the peered VPC owned by the consumer project.
        /// </summary>
        [Input("networkVpc", required: true)]
        public Input<string> NetworkVpc { get; set; } = null!;

        public PeeredVpcArgs()
        {
        }
    }
}