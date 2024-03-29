// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Outputs
{

    /// <summary>
    /// A list of flow version variants.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDialogflowCxV3VersionVariantsResponse
    {
        /// <summary>
        /// A list of flow version variants.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudDialogflowCxV3VersionVariantsVariantResponse> Variants;

        [OutputConstructor]
        private GoogleCloudDialogflowCxV3VersionVariantsResponse(ImmutableArray<Outputs.GoogleCloudDialogflowCxV3VersionVariantsVariantResponse> variants)
        {
            Variants = variants;
        }
    }
}
