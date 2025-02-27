// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2Alpha1.Inputs
{

    /// <summary>
    /// Details of the TPU node(s) being requested. Users can request either a single node or multiple nodes. NodeSpec provides the specification for node(s) to be created.
    /// </summary>
    public sealed class NodeSpecArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Fields to specify in case of multi-node request.
        /// </summary>
        [Input("multiNodeParams")]
        public Input<Inputs.MultiNodeParamsArgs>? MultiNodeParams { get; set; }

        /// <summary>
        /// The node.
        /// </summary>
        [Input("node", required: true)]
        public Input<Inputs.NodeArgs> Node { get; set; } = null!;

        /// <summary>
        /// The unqualified resource name. Should follow the `^[A-Za-z0-9_.~+%-]+$` regex format. This is only specified when requesting a single node. In case of multi-node requests, multi_node_params must be populated instead. It's an error to specify both node_id and multi_node_params.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// The parent resource name.
        /// </summary>
        [Input("parent", required: true)]
        public Input<string> Parent { get; set; } = null!;

        public NodeSpecArgs()
        {
        }
        public static new NodeSpecArgs Empty => new NodeSpecArgs();
    }
}
