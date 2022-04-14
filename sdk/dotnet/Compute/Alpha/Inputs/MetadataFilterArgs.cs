// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Opaque filter criteria used by load balancers to restrict routing configuration to a limited set of load balancing proxies. Proxies and sidecars involved in load balancing would typically present metadata to the load balancers that need to match criteria specified here. If a match takes place, the relevant configuration is made available to those proxies. For each metadataFilter in this list, if its filterMatchCriteria is set to MATCH_ANY, at least one of the filterLabels must match the corresponding label provided in the metadata. If its filterMatchCriteria is set to MATCH_ALL, then all of its filterLabels must match with corresponding labels provided in the metadata. An example for using metadataFilters would be: if load balancing involves Envoys, they receive routing configuration when values in metadataFilters match values supplied in of their XDS requests to loadbalancers.
    /// </summary>
    public sealed class MetadataFilterArgs : Pulumi.ResourceArgs
    {
        [Input("filterLabels")]
        private InputList<Inputs.MetadataFilterLabelMatchArgs>? _filterLabels;

        /// <summary>
        /// The list of label value pairs that must match labels in the provided metadata based on filterMatchCriteria This list must not be empty and can have at the most 64 entries.
        /// </summary>
        public InputList<Inputs.MetadataFilterLabelMatchArgs> FilterLabels
        {
            get => _filterLabels ?? (_filterLabels = new InputList<Inputs.MetadataFilterLabelMatchArgs>());
            set => _filterLabels = value;
        }

        /// <summary>
        /// Specifies how individual filter label matches within the list of filterLabels and contributes toward the overall metadataFilter match. Supported values are: - MATCH_ANY: at least one of the filterLabels must have a matching label in the provided metadata. - MATCH_ALL: all filterLabels must have matching labels in the provided metadata. 
        /// </summary>
        [Input("filterMatchCriteria")]
        public Input<Pulumi.GoogleNative.Compute.Alpha.MetadataFilterFilterMatchCriteria>? FilterMatchCriteria { get; set; }

        public MetadataFilterArgs()
        {
        }
    }
}