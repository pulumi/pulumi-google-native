// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.FirebaseRules.V1.Outputs
{

    /// <summary>
    /// `Source` is one or more `File` messages comprising a logical set of rules.
    /// </summary>
    [OutputType]
    public sealed class SourceResponse
    {
        /// <summary>
        /// `File` set constituting the `Source` bundle.
        /// </summary>
        public readonly ImmutableArray<Outputs.FileResponse> Files;

        [OutputConstructor]
        private SourceResponse(ImmutableArray<Outputs.FileResponse> files)
        {
            Files = files;
        }
    }
}
