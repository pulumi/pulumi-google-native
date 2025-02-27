// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Node group identification and configuration information.
    /// </summary>
    [OutputType]
    public sealed class AuxiliaryNodeGroupResponse
    {
        /// <summary>
        /// Node group configuration.
        /// </summary>
        public readonly Outputs.NodeGroupResponse NodeGroup;
        /// <summary>
        /// Optional. A node group ID. Generated if not specified.The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of from 3 to 33 characters.
        /// </summary>
        public readonly string NodeGroupId;

        [OutputConstructor]
        private AuxiliaryNodeGroupResponse(
            Outputs.NodeGroupResponse nodeGroup,

            string nodeGroupId)
        {
            NodeGroup = nodeGroup;
            NodeGroupId = nodeGroupId;
        }
    }
}
