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
    /// PlacementPolicy defines the placement policy used by the node pool.
    /// </summary>
    public sealed class PlacementPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of placement.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Container.V1.PlacementPolicyType>? Type { get; set; }

        public PlacementPolicyArgs()
        {
        }
        public static new PlacementPolicyArgs Empty => new PlacementPolicyArgs();
    }
}