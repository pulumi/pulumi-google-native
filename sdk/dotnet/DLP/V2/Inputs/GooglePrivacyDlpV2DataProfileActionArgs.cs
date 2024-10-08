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
    /// A task to execute when a data profile has been generated.
    /// </summary>
    public sealed class GooglePrivacyDlpV2DataProfileActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Export data profiles into a provided location.
        /// </summary>
        [Input("exportData")]
        public Input<Inputs.GooglePrivacyDlpV2ExportArgs>? ExportData { get; set; }

        /// <summary>
        /// Publish a message into the Pub/Sub topic.
        /// </summary>
        [Input("pubSubNotification")]
        public Input<Inputs.GooglePrivacyDlpV2PubSubNotificationArgs>? PubSubNotification { get; set; }

        public GooglePrivacyDlpV2DataProfileActionArgs()
        {
        }
        public static new GooglePrivacyDlpV2DataProfileActionArgs Empty => new GooglePrivacyDlpV2DataProfileActionArgs();
    }
}
