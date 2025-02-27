// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkeonprem.V1.Inputs
{

    /// <summary>
    /// Specifies the security related settings for the bare metal user cluster.
    /// </summary>
    public sealed class BareMetalSecurityConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configures user access to the user cluster.
        /// </summary>
        [Input("authorization")]
        public Input<Inputs.AuthorizationArgs>? Authorization { get; set; }

        public BareMetalSecurityConfigArgs()
        {
        }
        public static new BareMetalSecurityConfigArgs Empty => new BareMetalSecurityConfigArgs();
    }
}
