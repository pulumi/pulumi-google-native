// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// Field represents either the key or value in an entry.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudIntegrationsV1alphaParameterMapFieldResponse
    {
        /// <summary>
        /// Passing a literal value.
        /// </summary>
        public readonly Outputs.GoogleCloudIntegrationsV1alphaValueTypeResponse LiteralValue;
        /// <summary>
        /// Referencing one of the Integration variables.
        /// </summary>
        public readonly string ReferenceKey;

        [OutputConstructor]
        private GoogleCloudIntegrationsV1alphaParameterMapFieldResponse(
            Outputs.GoogleCloudIntegrationsV1alphaValueTypeResponse literalValue,

            string referenceKey)
        {
            LiteralValue = literalValue;
            ReferenceKey = referenceKey;
        }
    }
}
