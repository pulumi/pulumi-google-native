// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Container.V1Beta1.Inputs
{

    /// <summary>
    /// Specifies the NodeAffinity key, values, and affinity operator according to [shared sole tenant node group affinities](https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes#node_affinity_and_anti-affinity).
    /// </summary>
    public sealed class NodeAffinityArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key for NodeAffinity.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Operator for NodeAffinity.
        /// </summary>
        [Input("operator")]
        public Input<Pulumi.GoogleNative.Container.V1Beta1.NodeAffinityOperator>? Operator { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Values for NodeAffinity.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public NodeAffinityArgs()
        {
        }
        public static new NodeAffinityArgs Empty => new NodeAffinityArgs();
    }
}