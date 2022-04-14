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
    /// CidrBlock contains an optional name and one CIDR block.
    /// </summary>
    public sealed class CidrBlockArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// cidr_block must be specified in CIDR notation.
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// display_name is an optional field for users to identify CIDR blocks.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        public CidrBlockArgs()
        {
        }
    }
}