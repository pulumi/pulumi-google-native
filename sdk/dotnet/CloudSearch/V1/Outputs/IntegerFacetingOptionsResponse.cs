// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudSearch.V1.Outputs
{

    /// <summary>
    /// Used to specify integer faceting options.
    /// </summary>
    [OutputType]
    public sealed class IntegerFacetingOptionsResponse
    {
        /// <summary>
        /// Buckets for given integer values should be in strictly ascending order. For example, if values supplied are (1,5,10,100), the following facet buckets will be formed {&lt;1, [1,5), [5-10), [10-100), &gt;=100}.
        /// </summary>
        public readonly ImmutableArray<string> IntegerBuckets;

        [OutputConstructor]
        private IntegerFacetingOptionsResponse(ImmutableArray<string> integerBuckets)
        {
            IntegerBuckets = integerBuckets;
        }
    }
}