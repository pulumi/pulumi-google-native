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
    /// Defines sources for the suggest/search APIs.
    /// </summary>
    [OutputType]
    public sealed class SourceResponse
    {
        /// <summary>
        /// Source name for content indexed by the Indexing API.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Predefined content source for Google Apps.
        /// </summary>
        public readonly string PredefinedSource;

        [OutputConstructor]
        private SourceResponse(
            string name,

            string predefinedSource)
        {
            Name = name;
            PredefinedSource = predefinedSource;
        }
    }
}
