// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DLP.V2.Inputs
{

    /// <summary>
    /// The cadence at which to update data profiles when a schema is modified.
    /// </summary>
    public sealed class GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// How frequently profiles may be updated when schemas are modified. Defaults to monthly.
        /// </summary>
        [Input("frequency")]
        public Input<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceFrequency>? Frequency { get; set; }

        [Input("types")]
        private InputList<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceTypesItem>? _types;

        /// <summary>
        /// The type of events to consider when deciding if the table's schema has been modified and should have the profile updated. Defaults to NEW_COLUMNS.
        /// </summary>
        public InputList<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceTypesItem> Types
        {
            get => _types ?? (_types = new InputList<Pulumi.GoogleNative.DLP.V2.GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceTypesItem>());
            set => _types = value;
        }

        public GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceArgs()
        {
        }
        public static new GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceArgs Empty => new GooglePrivacyDlpV2DiscoverySchemaModifiedCadenceArgs();
    }
}
