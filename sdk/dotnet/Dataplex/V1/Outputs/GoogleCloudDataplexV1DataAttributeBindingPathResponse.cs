// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Outputs
{

    /// <summary>
    /// Represents a subresource of a given resource, and associated bindings with it.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDataplexV1DataAttributeBindingPathResponse
    {
        /// <summary>
        /// Optional. List of attributes to be associated with the path of the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
        /// </summary>
        public readonly ImmutableArray<string> Attributes;
        /// <summary>
        /// The name identifier of the path. Nested columns should be of the form: 'country.state.city'.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GoogleCloudDataplexV1DataAttributeBindingPathResponse(
            ImmutableArray<string> attributes,

            string name)
        {
            Attributes = attributes;
            Name = name;
        }
    }
}