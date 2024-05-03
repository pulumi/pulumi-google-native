// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigtableAdmin.V2.Outputs
{

    /// <summary>
    /// Rule for determining which cells to delete during garbage collection.
    /// </summary>
    [OutputType]
    public sealed class GcRuleResponse
    {
        /// <summary>
        /// Delete cells that would be deleted by every nested rule.
        /// </summary>
        public readonly Outputs.IntersectionResponse Intersection;
        /// <summary>
        /// Delete cells in a column older than the given age. Values must be at least one millisecond, and will be truncated to microsecond granularity.
        /// </summary>
        public readonly string MaxAge;
        /// <summary>
        /// Delete all cells in a column except the most recent N.
        /// </summary>
        public readonly int MaxNumVersions;
        /// <summary>
        /// Delete cells that would be deleted by any nested rule.
        /// </summary>
        public readonly Outputs.UnionResponse Union;

        [OutputConstructor]
        private GcRuleResponse(
            Outputs.IntersectionResponse intersection,

            string maxAge,

            int maxNumVersions,

            Outputs.UnionResponse union)
        {
            Intersection = intersection;
            MaxAge = maxAge;
            MaxNumVersions = maxNumVersions;
            Union = union;
        }
    }
}