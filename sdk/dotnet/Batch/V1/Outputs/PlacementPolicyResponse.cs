// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Batch.V1.Outputs
{

    /// <summary>
    /// PlacementPolicy describes a group placement policy for the VMs controlled by this AllocationPolicy.
    /// </summary>
    [OutputType]
    public sealed class PlacementPolicyResponse
    {
        /// <summary>
        /// UNSPECIFIED vs. COLLOCATED (default UNSPECIFIED). Use COLLOCATED when you want VMs to be located close to each other for low network latency between the VMs. No placement policy will be generated when collocation is UNSPECIFIED.
        /// </summary>
        public readonly string Collocation;
        /// <summary>
        /// When specified, causes the job to fail if more than max_distance logical switches are required between VMs. Batch uses the most compact possible placement of VMs even when max_distance is not specified. An explicit max_distance makes that level of compactness a strict requirement. Not yet implemented
        /// </summary>
        public readonly string MaxDistance;

        [OutputConstructor]
        private PlacementPolicyResponse(
            string collocation,

            string maxDistance)
        {
            Collocation = collocation;
            MaxDistance = maxDistance;
        }
    }
}