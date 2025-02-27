// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1Beta1.Outputs
{

    /// <summary>
    /// Contains multiple sensitive information findings for each resource slice.
    /// </summary>
    [OutputType]
    public sealed class DetailResponse
    {
        public readonly ImmutableArray<Outputs.FindingResponse> Findings;

        [OutputConstructor]
        private DetailResponse(ImmutableArray<Outputs.FindingResponse> findings)
        {
            Findings = findings;
        }
    }
}
