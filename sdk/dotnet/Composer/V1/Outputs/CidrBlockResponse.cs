// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1.Outputs
{

    /// <summary>
    /// CIDR block with an optional name.
    /// </summary>
    [OutputType]
    public sealed class CidrBlockResponse
    {
        /// <summary>
        /// CIDR block that must be specified in CIDR notation.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// User-defined name that identifies the CIDR block.
        /// </summary>
        public readonly string DisplayName;

        [OutputConstructor]
        private CidrBlockResponse(
            string cidrBlock,

            string displayName)
        {
            CidrBlock = cidrBlock;
            DisplayName = displayName;
        }
    }
}